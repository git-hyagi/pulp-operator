# \PypiLegacyAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PypiLegacyCreate**](PypiLegacyAPI.md#PypiLegacyCreate) | **Post** /pypi/{path}/legacy/ | Upload a package



## PypiLegacyCreate

> PackageUploadTaskResponse PypiLegacyCreate(ctx, path).XTaskDiagnostics(xTaskDiagnostics).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).Execute()

Upload a package



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
	path := "path_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	pulpLabels := map[string]string{"key": "Inner_example"} // map[string]string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PypiLegacyAPI.PypiLegacyCreate(context.Background(), path).XTaskDiagnostics(xTaskDiagnostics).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PypiLegacyAPI.PypiLegacyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PypiLegacyCreate`: PackageUploadTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `PypiLegacyAPI.PypiLegacyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiLegacyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that may be turned into the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the content unit. | 
 **fileUrl** | **string** | A url that Pulp can download and turn into the content unit. | 

### Return type

[**PackageUploadTaskResponse**](PackageUploadTaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

