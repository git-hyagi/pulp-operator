# \PulpAnsiblePluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyPluginAnsibleContentCollectionsIndexVersionsDocsBlobRead**](PulpAnsiblePluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI.md#PulpAnsibleGalaxyPluginAnsibleContentCollectionsIndexVersionsDocsBlobRead) | **Get** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/versions/{version}/docs-blob/ | 



## PulpAnsibleGalaxyPluginAnsibleContentCollectionsIndexVersionsDocsBlobRead

> CollectionVersionDocsResponse PulpAnsibleGalaxyPluginAnsibleContentCollectionsIndexVersionsDocsBlobRead(ctx, distroBasePath, name, namespace, path, version).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()





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
	name := "name_example" // string | 
	namespace := "namespace_example" // string | 
	path := "path_example" // string | 
	version := "version_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PulpAnsiblePluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI.PulpAnsibleGalaxyPluginAnsibleContentCollectionsIndexVersionsDocsBlobRead(context.Background(), distroBasePath, name, namespace, path, version).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsiblePluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI.PulpAnsibleGalaxyPluginAnsibleContentCollectionsIndexVersionsDocsBlobRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PulpAnsibleGalaxyPluginAnsibleContentCollectionsIndexVersionsDocsBlobRead`: CollectionVersionDocsResponse
	fmt.Fprintf(os.Stdout, "Response from `PulpAnsiblePluginAnsibleContentCollectionsIndexVersionsDocsBlobAPI.PulpAnsibleGalaxyPluginAnsibleContentCollectionsIndexVersionsDocsBlobRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 
**path** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyPluginAnsibleContentCollectionsIndexVersionsDocsBlobReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**CollectionVersionDocsResponse**](CollectionVersionDocsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

