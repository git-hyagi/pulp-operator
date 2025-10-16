# \PulpAnsibleDefaultPluginAnsibleContentCollectionsAllVersionsAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsAllVersionsList**](PulpAnsibleDefaultPluginAnsibleContentCollectionsAllVersionsAPI.md#PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsAllVersionsList) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/all-versions/ | 



## PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsAllVersionsList

> []UnpaginatedCollectionVersionResponse PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsAllVersionsList(ctx, distroBasePath).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pulp/pulp-operator/bindings/release"
)

func main() {
	distroBasePath := "distroBasePath_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PulpAnsibleDefaultPluginAnsibleContentCollectionsAllVersionsAPI.PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsAllVersionsList(context.Background(), distroBasePath).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultPluginAnsibleContentCollectionsAllVersionsAPI.PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsAllVersionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsAllVersionsList`: []UnpaginatedCollectionVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultPluginAnsibleContentCollectionsAllVersionsAPI.PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsAllVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsAllVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**[]UnpaginatedCollectionVersionResponse**](UnpaginatedCollectionVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

