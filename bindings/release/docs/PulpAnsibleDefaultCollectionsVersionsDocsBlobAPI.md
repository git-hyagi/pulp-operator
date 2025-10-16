# \PulpAnsibleDefaultCollectionsVersionsDocsBlobAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultCollectionsVersionsDocsBlobRead**](PulpAnsibleDefaultCollectionsVersionsDocsBlobAPI.md#PulpAnsibleGalaxyDefaultCollectionsVersionsDocsBlobRead) | **Get** /pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/versions/{version}/docs-blob/ | 



## PulpAnsibleGalaxyDefaultCollectionsVersionsDocsBlobRead

> CollectionVersionDocsResponse PulpAnsibleGalaxyDefaultCollectionsVersionsDocsBlobRead(ctx, name, namespace, version).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()





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
	name := "name_example" // string | 
	namespace := "namespace_example" // string | 
	version := "version_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PulpAnsibleDefaultCollectionsVersionsDocsBlobAPI.PulpAnsibleGalaxyDefaultCollectionsVersionsDocsBlobRead(context.Background(), name, namespace, version).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultCollectionsVersionsDocsBlobAPI.PulpAnsibleGalaxyDefaultCollectionsVersionsDocsBlobRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PulpAnsibleGalaxyDefaultCollectionsVersionsDocsBlobRead`: CollectionVersionDocsResponse
	fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultCollectionsVersionsDocsBlobAPI.PulpAnsibleGalaxyDefaultCollectionsVersionsDocsBlobRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultCollectionsVersionsDocsBlobReadRequest struct via the builder pattern


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

