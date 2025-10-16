# \PulpAnsibleArtifactsCollectionsAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyArtifactsCollectionsGet**](PulpAnsibleArtifactsCollectionsAPI.md#PulpAnsibleGalaxyArtifactsCollectionsGet) | **Get** /pulp_ansible/galaxy/{path}/api/v3/artifacts/collections/{path}/{filename} | 



## PulpAnsibleGalaxyArtifactsCollectionsGet

> PulpAnsibleGalaxyArtifactsCollectionsGet(ctx, filename, path).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()





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
	filename := "filename_example" // string | 
	path := "path_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PulpAnsibleArtifactsCollectionsAPI.PulpAnsibleGalaxyArtifactsCollectionsGet(context.Background(), filename, path).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleArtifactsCollectionsAPI.PulpAnsibleGalaxyArtifactsCollectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyArtifactsCollectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

