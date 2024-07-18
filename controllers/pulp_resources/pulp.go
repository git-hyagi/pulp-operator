package repo_manager_pulp

import (
	"reflect"
	"sync"

	repomanagerpulpprojectorgv1beta2 "github.com/pulp/pulp-operator/apis/repo-manager.pulpproject.org/v1beta2"
)

func (r *RepoManagerPulpResourceReconciler) createPulpResources(pulpResource *repomanagerpulpprojectorgv1beta2.PulpResource, pulpClient pulpClient) {

	log := pulpClient.Log
	owner := pulpResourceOwner(*pulpResource)
	operatorLabels := operatorLabels(*pulpResource)

	if pulpResource.Status == nil {
		pulpResource.Status = &repomanagerpulpprojectorgv1beta2.PulpResourceStatus{}
	}

	var wg sync.WaitGroup
	for plugin, resource := range plugins {
		wg.Add(1)

		go func(pluginName string, resource map[string]string) {
			defer wg.Done()
			for resourceType, endpoint := range resource {
				updatePulpResource(pulpResource, pluginName, resourceType, pulpClient, endpoint)
			}
		}(plugin, resource)

		wg.Add(1)
		go func(pluginName string, resource map[string]string) {
			defer wg.Done()
			pluginField := reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(pluginName)
			if field_empty(pluginField) {
				return
			}
			if is_status_plugin_equals_to_spec(pulpResource, pluginName) {
				return
			}
			for resourceType, endpoint := range resource {
				resourceValue := reflect.Indirect(pluginField).FieldByName(resourceType)
				if field_empty(resourceValue) {
					continue
				}

				resourceName := reflect.Indirect(resourceValue).FieldByName("Name").Interface().(string)
				if response, exists := resource_exists(pulpClient, endpoint, resourceName); exists {
					// if already exists and responsebody is equal to what is in the spec, nothing to do
					// if already exists, but what we got from pulp is different from what is in the spec we need to update it
					if !owned_by_operator(response, owner) {
						log.Error(nil, pluginName+" "+resourceType+" "+resourceName+" already exists, but is not managed by Pulp operator!")
						continue
					}
					log.V(1).Info(pluginName + " " + resourceType + " " + resourceName + " already exists!")
					continue
				}

				// temporarily set the labels to create pulp resoruces
				set_operator_label_to_resource(operatorLabels, &resourceValue)
				if _, err := make_request("POST", pulpClient.Address+endpoint, pulpClient.Username, pulpClient.Password, resourceValue.Interface()); err != nil {
					log.Error(err, "Failed to create "+pluginName+" "+resourceType+" "+resourceName)
					continue
				}
				// unset the labels to not populate status field with it
				unset_operator_label_to_resource(&resourceValue)
				log.Info(pluginName + " " + resourceType + " " + resourceName + " created!")
			}
		}(plugin, resource)
	}
	wg.Wait()

}

func deletePulpResource(pulpResource *repomanagerpulpprojectorgv1beta2.PulpResource, plugin, resource, endpoint string, pulpClient pulpClient) {
	log := pulpClient.Log
	resourceNameToBeDeleted := reflect.ValueOf(*pulpResource).FieldByName("Status").Elem().FieldByName(plugin).Elem().FieldByName(resource).Elem().FieldByName("Name").Interface().(string)

	// make a request to get the pulp_href from the resource to be removed
	log.V(1).Info("Making a GET request to " + endpoint)
	response, err := make_request("GET", pulpClient.Address+endpoint, pulpClient.Username, pulpClient.Password, resourceNameToBeDeleted)
	if err != nil {
		log.Error(err, "Failed to make GET request to "+endpoint)
	}
	if response.Count == 0 {
		pulpClient.Log.Error(err, "Failed to find the "+plugin+" "+resource+" "+resourceNameToBeDeleted)
		return
	}

	pulp_href := response.Results[0].(map[string]interface{})["pulp_href"].(string)
	log.V(1).Info("GET response " + pulp_href)

	// if the resource is found, but is not owned by pulp-operator we should not touch it
	owner := pulpResourceOwner(*pulpResource)
	if !owned_by_operator(response, owner) {
		log.Error(nil, plugin+" "+resource+" "+resourceNameToBeDeleted+" already exists, but is not managed by Pulp operator!")
		return
	}

	// delete it
	log.V(1).Info("Making a DELETE request to " + pulp_href)
	_, err = make_request("DELETE", pulpClient.Address+pulp_href, pulpClient.Username, pulpClient.Password, nil)
	if err != nil {
		pulpClient.Log.Error(err, "Failed to make DELETE request to "+pulp_href)
	}
	log.Info(plugin + " " + resource + " " + resourceNameToBeDeleted + " deleted!")
}

func updatePulpResource(pulpResource *repomanagerpulpprojectorgv1beta2.PulpResource, plugin, resource string, pulpClient pulpClient, endpoint string) {
	log := pulpClient.Log

	statusField := reflect.ValueOf(*pulpResource).FieldByName("Status")

	// if status is not defined yet (first operator execution) we just return
	if field_empty(statusField) {
		return
	}

	// if status.plugin is not defined yet (no previous definition of this plugin) there is nothing to be done
	statusPluginField := statusField.Elem().FieldByName(plugin)
	if field_empty(statusPluginField) {
		return
	}

	// if resource exists in status but not in spec, it means the user wants to remove a previously defined resource
	statusResourceField := statusPluginField.Elem().FieldByName(resource)
	specField := reflect.ValueOf(*pulpResource).FieldByName("Spec")
	if !field_empty(statusResourceField) &&
		(field_empty(specField) ||
			specField.FieldByName(plugin).IsNil() ||
			specField.FieldByName(plugin).Elem().FieldByName(resource).IsNil()) {
		deletePulpResource(pulpResource, plugin, resource, endpoint, pulpClient)
		return
	}

	// if what is in status is equal to spec, we are up-to-date, so there is nothing to do
	specPluginField := specField.FieldByName(plugin)
	if !field_empty(specPluginField) {
		specResourceField := specPluginField.Elem().FieldByName(resource)
		if reflect.DeepEqual(statusResourceField.Interface(), specResourceField.Interface()) {
			return
		}
	}

	// in the first execution there is no definition of the status, we just return
	if field_empty(statusPluginField) ||
		field_empty(statusResourceField) {
		return
	}

	resourceNameToBeUpdated := statusResourceField.Elem().FieldByName("Name").Interface().(string)

	// make a request to get the pulp_href from the resource to be removed
	log.V(1).Info("Making a GET request to " + endpoint)
	response, err := make_request("GET", pulpClient.Address+endpoint, pulpClient.Username, pulpClient.Password, resourceNameToBeUpdated)
	if err != nil {
		log.Error(err, "Failed to make a GET request to "+endpoint)
	}

	if response.Count == 0 {
		log.Error(err, "Failed to find the "+plugin+" "+resource+" "+resourceNameToBeUpdated)
		return
	}
	pulp_href := response.Results[0].(map[string]interface{})["pulp_href"].(string)
	log.V(1).Info("GET response " + pulp_href)

	// if the resource is found, but is not owned by pulp-operator we should not touch it
	owner := pulpResourceOwner(*pulpResource)
	if !owned_by_operator(response, owner) {
		log.Error(nil, plugin+" "+resource+" "+resourceNameToBeUpdated+" already exists, but is not managed by Pulp operator!")
		return
	}

	// if the resource name changed we need to delete the resource (a new one will be automatically created by the create resource function)
	specResourceField := specPluginField.Elem().FieldByName(resource)
	if statusResourceField.Elem().FieldByName("Name").String() != specResourceField.Elem().FieldByName("Name").String() {
		// delete it
		log.V(1).Info("Making a DELETE request to " + pulp_href)
		_, err = make_request("DELETE", pulpClient.Address+pulp_href, pulpClient.Username, pulpClient.Password, nil)
		if err != nil {
			log.Error(err, "Failed to make DELETE request to "+pulp_href)
		}
		log.Info(plugin + " " + resource + " " + resourceNameToBeUpdated + " deleted!")
	} else {
		/// patch it
		log.V(1).Info("Making a PATCH request to " + pulp_href)
		response, err = make_request("PATCH", pulpClient.Address+pulp_href, pulpClient.Username, pulpClient.Password, specResourceField.Elem().Interface())
		if err != nil {
			log.Error(err, "Failed to make a PATCH request to "+pulp_href)
		}
		log.Info(plugin + " " + resource + " " + resourceNameToBeUpdated + " updated!")
	}
}

func createSyncResource(pulpResource *repomanagerpulpprojectorgv1beta2.PulpResource, plugin string, pulpClient pulpClient) {
	log, username, password, pulp_address := pulpClient.Log, pulpClient.Username, pulpClient.Password, pulpClient.Address

	specRemoteField := reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(plugin).Elem().FieldByName("Remote")
	if field_empty(specRemoteField) {
		log.Error(nil, "Sync enabled for "+plugin+" plugin without defining a remote.")
		return
	}
	specRepositoryField := reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(plugin).Elem().FieldByName("Repository")
	if field_empty(specRepositoryField) {
		log.Error(nil, "Sync enabled for "+plugin+" plugin without defining a repository.")
		return
	}
	specDistributionField := reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(plugin).Elem().FieldByName("Distribution")
	if field_empty(specDistributionField) {
		log.Error(nil, "Sync enabled for "+plugin+" plugin without defining a distribution.")
		return
	}

	statusPluginField := reflect.ValueOf(*pulpResource).FieldByName("Status").Elem().FieldByName(plugin)
	if !field_empty(statusPluginField) {
		if reflect.DeepEqual(specRemoteField.Interface(), statusPluginField.Elem().FieldByName("Remote").Interface()) &&
			reflect.DeepEqual(specRepositoryField.Interface(), statusPluginField.Elem().FieldByName("Repository").Interface()) &&
			reflect.DeepEqual(specDistributionField.Interface(), statusPluginField.Elem().FieldByName("Distribution").Interface()) {
			log.V(1).Info("Already synced " + plugin + " plugin before, nothing to do")
			return
		}
	}

	specRepositoryName := specRepositoryField.Elem().FieldByName("Name").String()
	specRemoteName := specRemoteField.Elem().FieldByName("Name").String()
	specDistributionName := specDistributionField.Elem().FieldByName("Name").String()

	repositoryHref := get_pulp_href(plugin, "Repository", specRepositoryName, pulpClient)
	if repositoryHref == "" {
		return
	}

	remoteHref := get_pulp_href(plugin, "Remote", specRemoteName, pulpClient)
	if remoteHref == "" {
		return
	}

	distributionHref := get_pulp_href(plugin, "Distribution", specDistributionName, pulpClient)
	if distributionHref == "" {
		return
	}

	repositoryRemoteBody := map[string]string{"remote": remoteHref}
	_, err := make_request("PATCH", pulp_address+repositoryHref, username, password, repositoryRemoteBody)
	if err != nil {
		log.Error(err, "Failed to make a PATCH request to "+repositoryHref)
	}

	sync := repomanagerpulpprojectorgv1beta2.PulpSync{Remote: remoteHref, Mirror: false, SignedOnly: false}
	_, err = make_request("POST", pulp_address+repositoryHref+"sync/", username, password, sync)
	if err != nil {
		log.Error(err, "Failed to make a POST request to "+repositoryHref+"sync/")
	}

	distributionBody := map[string]string{"repository": repositoryHref}
	_, err = make_request("PATCH", pulp_address+distributionHref, username, password, distributionBody)
	if err != nil {
		log.Error(err, "Failed to make a PATCH request to "+distributionHref)
	}

	// todo: update pulpresource status (status.sync) based on the pulp task output
}

func (r *RepoManagerPulpResourceReconciler) syncPulpResources(pulpResource *repomanagerpulpprojectorgv1beta2.PulpResource, pulpClient pulpClient) {
	var wg sync.WaitGroup
	for plugin := range plugins {
		specPluginField := reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(plugin)
		if field_empty(specPluginField) {
			continue
		}

		specSyncField := specPluginField.Elem().FieldByName("Sync")
		if field_empty(specSyncField) ||
			!specSyncField.Elem().Bool() {
			continue
		}

		wg.Add(1)
		go func(pluginName string) {
			defer wg.Done()
			createSyncResource(pulpResource, pluginName, pulpClient)
		}(plugin)
	}
	wg.Wait()
}

func get_pulp_href(plugin, resourceType, resourceName string, pulpClient pulpClient) string {
	log := pulpClient.Log
	// make a request to get the pulp_href from the resource to be synced
	resourceEndpoint := plugins[plugin][resourceType]
	log.V(1).Info("Making a GET request to " + resourceEndpoint)
	response, err := make_request("GET", pulpClient.Address+resourceEndpoint, pulpClient.Username, pulpClient.Password, resourceName)
	if err != nil {
		log.Error(err, "Failed to make a GET request to "+resourceEndpoint)
	}
	if response.Count == 0 {
		log.Error(err, "Failed to find the "+plugin+" "+resourceType+" "+resourceName)
		return ""
	}
	resourceHref := response.Results[0].(map[string]interface{})["pulp_href"].(string)
	log.V(1).Info("GET response " + resourceHref)
	return resourceHref
}
