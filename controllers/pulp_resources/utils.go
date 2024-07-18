package repo_manager_pulp

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sync"

	"github.com/go-logr/logr"
	repomanagerpulpprojectorgv1beta2 "github.com/pulp/pulp-operator/apis/repo-manager.pulpproject.org/v1beta2"
	"github.com/pulp/pulp-operator/controllers/settings"
	corev1 "k8s.io/api/core/v1"
	k8s_errors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	HTTP_NOT_FOUND = errors.New("404: Not Found")

	plugins = map[string]map[string]string{
		"Container": {
			"Repository":   PULP_CONTAINER_REPO_ENDPOINT,
			"Distribution": PULP_CONTAINER_DISTRIBUTION_ENDPOINT,
			"Remote":       PULP_CONTAINER_REMOTE_ENDPOINT,
		},
		"File": {
			"Repository":   PULP_FILE_REPO_ENDPOINT,
			"Distribution": PULP_FILE_DISTRIBUTION_ENDPOINT,
			"Remote":       PULP_FILE_REMOTE_ENDPOINT,
		},
		"RPM": {
			"Repository":   PULP_RPM_REPO_ENDPOINT,
			"Distribution": PULP_RPM_DISTRIBUTION_ENDPOINT,
			"Remote":       PULP_RPM_REMOTE_ENDPOINT,
		},
	}
)

type pulpClient struct {
	Address  string
	Username string
	Password string
	Log      logr.Logger
	Context  context.Context
}

type responseBody struct {
	Count    int
	Next     any
	Previous any
	Results  []any
}

func get_api_svc(pulpName string) string {
	return fmt.Sprintf("http://%s:24817", settings.ApiService(pulpName))
}

func get_credentials(ctx context.Context, r RepoManagerPulpResourceReconciler, pulpName, namespace string) string {
	pulp := &repomanagerpulpprojectorgv1beta2.Pulp{}
	r.Get(ctx, types.NamespacedName{Name: pulpName, Namespace: namespace}, pulp)

	adminSecret := pulp.Spec.AdminPasswordSecret
	secret := &corev1.Secret{}
	r.Get(ctx, types.NamespacedName{Name: adminSecret, Namespace: namespace}, secret)

	return string(secret.Data["password"])
}

func make_request(method, addr, user, pass string, body any) (responseBody, error) {
	// todo: this config should come from pulpClient
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	var req *http.Request
	if method == "POST" || method == "PATCH" || method == "PUT" {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return responseBody{}, err
		}
		postBody := bytes.NewBuffer(jsonBody)
		req, err = http.NewRequest(method, addr, postBody)
		if err != nil {
			return responseBody{}, err
		}
	} else if method == "GET" {
		req, _ = http.NewRequest(method, addr+"?name="+body.(string), nil)
	} else if method == "DELETE" {
		req, _ = http.NewRequest(method, addr, nil)
	}

	req.SetBasicAuth(user, pass)
	req.Header.Set("Content-Type", `application/json`)
	req.Header.Add("Accept", `application/json`)

	resp, err := client.Do(req)
	if err != nil {
		return responseBody{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	response := &responseBody{}
	json.Unmarshal(respBody, response)
	if err != nil {
		return responseBody{}, err
	}

	if resp.StatusCode/100 != 2 {
		return responseBody{}, fmt.Errorf(string(respBody))
	}

	return *response, nil
}

// is_status_plugin_equals_to_spec compares the definitions of pulpresource.status and pulpresource.spec
func is_status_plugin_equals_to_spec(pulpResource *repomanagerpulpprojectorgv1beta2.PulpResource, plugin string) bool {
	return reflect.DeepEqual(
		reflect.ValueOf(*pulpResource).FieldByName("Status").Elem().FieldByName(plugin).Interface(),
		reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(plugin).Interface())
}

// field_empty checks is a pulpresource field is defined
func field_empty(field reflect.Value) bool {
	return !field.IsValid() || field.IsZero()
}

// resource_exists verifies if pulp resource already exists
func resource_exists(pulpClient pulpClient, endpoint string, resource any) (responseBody, bool) {
	// todo: pending handle errors (like 401, 403, etc)
	response, err := make_request("GET", pulpClient.Address+endpoint, pulpClient.Username, pulpClient.Password, resource)
	if err != nil {
		pulpClient.Log.Error(err, "Failed to make GET request to "+endpoint)
	}
	if response.Count > 0 {
		return response, true
	}
	return response, false
}

// set_operator_label_to_resource will set pulp resource with a label to help identifying it is an operator managed resource
func set_operator_label_to_resource(operatorLabels map[string]string, resourceValue *reflect.Value) {
	reflect.Indirect(*resourceValue).FieldByName("Labels").Set(reflect.ValueOf(operatorLabels))
}

// unset_operator_label_to_resource will remove the resource labels definitions
func unset_operator_label_to_resource(resourceValue *reflect.Value) {
	reflect.Indirect(*resourceValue).FieldByName("Labels").Set(reflect.ValueOf(map[string]string{}))
}

// create_status_fields creates the plugin and resource fields to be able to set them dinamically
func create_status_fields(pulpResource *repomanagerpulpprojectorgv1beta2.PulpResource, plugin string, resource string) {
	pluginField := reflect.ValueOf(*pulpResource).FieldByName("Status").Elem().FieldByName(plugin)
	if pluginField.IsNil() {
		newField := reflect.New(reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(plugin).Elem().Type())
		reflect.ValueOf(*pulpResource).FieldByName("Status").Elem().FieldByName(plugin).Set(newField)
	}
	resourceField := reflect.ValueOf(*pulpResource).FieldByName("Status").Elem().FieldByName(plugin).Elem().FieldByName(resource)
	if resourceField.IsNil() {
		newField := reflect.New(reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(plugin).Elem().FieldByName(resource).Elem().Type())
		reflect.ValueOf(*pulpResource).FieldByName("Status").Elem().FieldByName(plugin).Elem().FieldByName(resource).Set(newField)
	}
}

// owned_by_operator will check, from a response body, if the resource is managed by pulp operator
func owned_by_operator(response responseBody, owner string) bool {
	return response.Results[0].(map[string]interface{})["pulp_labels"].(map[string]interface{})["owner"] == owner
}

func getClientConfig(ctx context.Context, r RepoManagerPulpResourceReconciler, pulpResource repomanagerpulpprojectorgv1beta2.PulpResource) pulpClient {
	pulpClient := pulpClient{
		Address:  get_api_svc(pulpResource.Name),
		Username: PULP_ADMIN_USER,
		Password: get_credentials(ctx, r, pulpResource.Name, pulpResource.Namespace),
		Log:      r.RawLogger,
		Context:  ctx,
	}
	if pulpResource.Spec.ClientConfig != nil {
		clientConfig := &corev1.Secret{}
		r.Get(ctx, types.NamespacedName{Name: *pulpResource.Spec.ClientConfig, Namespace: pulpResource.Namespace}, clientConfig)
		pulpClient.Address = string(clientConfig.Data["address"])
		pulpClient.Username = string(clientConfig.Data["username"])
		pulpClient.Password = string(clientConfig.Data["password"])
	}

	return pulpClient
}

func (r RepoManagerPulpResourceReconciler) update_status_fields(pulpResource *repomanagerpulpprojectorgv1beta2.PulpResource, pulpClient pulpClient) *ctrl.Result {
	var wg sync.WaitGroup
	for plugin, resource := range plugins {
		wg.Add(1)
		go func(pluginName string, resource map[string]string) {
			defer wg.Done()
			for resourceType := range resource {
				// remove the resource from status.plugin.resource if it was removed from spec.plugin.resource
				specPluginField := reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(plugin)
				statusPluginField := reflect.ValueOf(*pulpResource).FieldByName("Status").Elem().FieldByName(plugin)
				if field_empty(specPluginField) {
					statusPluginField.Set(reflect.Zero(statusPluginField.Type()))
					continue
				}

				// remove the resource from status.plugin.resource if it was removed from spec.plugin.resource
				specResourceField := specPluginField.Elem().FieldByName(resourceType)
				if field_empty(specResourceField) {
					statusResourceField := statusPluginField.Elem().FieldByName(resourceType)
					statusResourceField.Set(reflect.Zero(statusResourceField.Type()))
					continue
				}

				create_status_fields(pulpResource, pluginName, resourceType)
				statusPluginField.Elem().FieldByName(resourceType).Set(specResourceField)
			}
		}(plugin, resource)
	}
	wg.Wait()
	if err := r.Status().Update(pulpClient.Context, pulpResource); k8s_errors.IsConflict(err) {
		return &ctrl.Result{Requeue: true}
	}
	return nil
}

func pulpResourceOwner(pulpResource repomanagerpulpprojectorgv1beta2.PulpResource) string {
	return pulpResource.Namespace + "." + pulpResource.Name
}
func operatorLabels(pulpResource repomanagerpulpprojectorgv1beta2.PulpResource) map[string]string {
	return map[string]string{
		"owner": pulpResourceOwner(pulpResource),
	}
}

func (r *RepoManagerPulpResourceReconciler) pulpResourceFinalizer(ctx context.Context, pulpResource repomanagerpulpprojectorgv1beta2.PulpResource, pulpClient pulpClient) (ctrl.Result, error) {
	if pulpResource.GetDeletionTimestamp() != nil {
		if controllerutil.ContainsFinalizer(&pulpResource, PULP_RESOURCE_FINALIZER) {
			if err := r.finalizePulpResource(&pulpResource, pulpClient); err != nil {
				return ctrl.Result{}, err
			}

			controllerutil.RemoveFinalizer(&pulpResource, PULP_RESOURCE_FINALIZER)
			err := r.Update(ctx, &pulpResource)
			if err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(&pulpResource, PULP_RESOURCE_FINALIZER) {
		controllerutil.AddFinalizer(&pulpResource, PULP_RESOURCE_FINALIZER)
		if err := r.Update(ctx, &pulpResource); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *RepoManagerPulpResourceReconciler) finalizePulpResource(pulpResource *repomanagerpulpprojectorgv1beta2.PulpResource, pulpClient pulpClient) error {

	for plugin, resource := range plugins {
		// nothing to do if there is nothing defined for this plugin
		pluginField := reflect.ValueOf(*pulpResource).FieldByName("Spec").FieldByName(plugin)
		if field_empty(pluginField) {
			continue
		}
		for resourceType, endpoint := range resource {
			// nothing to do if there is nothing defined for this resource
			resourceValue := reflect.Indirect(pluginField).FieldByName(resourceType)
			if field_empty(resourceValue) {
				continue
			}
			deletePulpResource(pulpResource, plugin, resourceType, endpoint, pulpClient)
		}
	}
	return nil
}

func (r *RepoManagerPulpResourceReconciler) findPulpResourceDependentObjects(ctx context.Context, obj client.Object) []reconcile.Request {

	associatedPulp := repomanagerpulpprojectorgv1beta2.PulpResourceList{}
	opts := &client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("pulpresources-objects", obj.GetName()),
		Namespace:     obj.GetNamespace(),
	}
	if err := r.List(context.TODO(), &associatedPulp, opts); err != nil {
		return []reconcile.Request{}
	}

	requests := make([]reconcile.Request, len(associatedPulp.Items))
	for i, item := range associatedPulp.Items {
		requests[i] = reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      item.GetName(),
				Namespace: item.GetNamespace(),
			},
		}
	}
	return requests
}

func indexerFunc(obj client.Object) []string {
	pulpResource := obj.(*repomanagerpulpprojectorgv1beta2.PulpResource)
	if pulpResource.Spec.ClientConfig != nil {
		return []string{*pulpResource.Spec.ClientConfig}
	}
	return []string{}
}

func validatePulpConnection(pulpClient pulpClient) error {

	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	var req *http.Request
	req, _ = http.NewRequest("GET", pulpClient.Address+PULP_STATUS_ENDPOINT, nil)

	req.SetBasicAuth(pulpClient.Username, pulpClient.Password)
	req.Header.Set("Content-Type", `application/json`)
	req.Header.Add("Accept", `application/json`)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	response := &responseBody{}
	err = json.Unmarshal(respBody, response)
	if err != nil {
		return err
	}
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf(string(respBody))
	}

	return nil
}
