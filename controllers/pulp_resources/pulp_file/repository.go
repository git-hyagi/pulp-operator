package pulp_file

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"

	pulpv1 "github.com/pulp/pulp-operator/apis/repo-manager.pulpproject.org/v1"
	bindings "github.com/pulp/pulp-operator/bindings/release"
	"github.com/pulp/pulp-operator/controllers/pulp_resources/utils"
)

func manageFileRepository(ctx context.Context, pulpResource *pulpv1.PulpResource, clientBinding *bindings.APIClient) {
	pluginName := "File"
	resourceType := "Repository"
	owner := pulpResource.Name + "." + pulpResource.Namespace
	operatorLabels := map[string]string{
		"owner": owner,
	}

	if pulpResource.Status == nil {
		pulpResource.Status = &pulpv1.PulpResourceStatus{}
	}

	pluginField := reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(pluginName)
	if utils.FieldEmpty(pluginField) {
		return
	}
	if utils.IsStatusPluginEqualsToSpec(pulpResource, pluginName) {
		return
	}
	resourceValue := reflect.Indirect(pluginField).FieldByName(resourceType)
	if utils.FieldEmpty(resourceValue) {
		return
	}

	resourceName := reflect.Indirect(resourceValue).FieldByName("Name").Interface().(string)
	// check if resource already exists
	repositoriesRequest := clientBinding.RepositoriesFileAPI.RepositoriesFileFileList(ctx)
	repositoriesFilterRequest := repositoriesRequest.Name(resourceName)
	repositoriesList, response, err := repositoriesFilterRequest.Execute()
	if err != nil {
		errMessage := fmt.Sprintf("[%v] Failed to create %s %s: %v\n", err, pluginName, resourceType, response)
		slog.Error(errMessage)
	}
	if len(repositoriesList.Results) > 0 {
		// pending implement this!
		/* 		if !utils.OwnedByOperator(response, owner) {
			slog.Error(pluginName + " " + resourceType + " " + resourceName + " already exists, but is not managed by Pulp operator!")
			return
		} */
		repoFoundMsg := fmt.Sprintf("%s %s %s already exists!", pluginName, resourceType, resourceName)
		slog.Info(repoFoundMsg)
		return
	}

	clientBinding.RepositoriesFileAPI.RepositoriesFileFileCreate(ctx)

	// temporarily set the labels to create pulp resoruces
	clientBinding.RepositoriesFileAPI.RepositoriesFileFileSetLabel(ctx, respoitoriesList)
	set_operator_label_to_resource(operatorLabels, &resourceValue)
	if _, err := make_request("POST", pulpClient.Address+endpoint, pulpClient.Username, pulpClient.Password, resourceValue.Interface()); err != nil {
		log.Error(err, "Failed to create "+pluginName+" "+resourceType+" "+resourceName)
		continue
	}
	// unset the labels to not populate status field with it
	unset_operator_label_to_resource(&resourceValue)
	log.Info(pluginName + " " + resourceType + " " + resourceName + " created!")

}
