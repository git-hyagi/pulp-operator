# \OrphansCleanupAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrphansCleanupCleanup**](OrphansCleanupAPI.md#OrphansCleanupCleanup) | **Post** /api/v3/orphans/cleanup/ | 



## OrphansCleanupCleanup

> AsyncOperationResponse OrphansCleanupCleanup(ctx).OrphansCleanup(orphansCleanup).XTaskDiagnostics(xTaskDiagnostics).Execute()





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
	orphansCleanup := *openapiclient.NewOrphansCleanup() // OrphansCleanup | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrphansCleanupAPI.OrphansCleanupCleanup(context.Background()).OrphansCleanup(orphansCleanup).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrphansCleanupAPI.OrphansCleanupCleanup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrphansCleanupCleanup`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrphansCleanupAPI.OrphansCleanupCleanup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrphansCleanupCleanupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orphansCleanup** | [**OrphansCleanup**](OrphansCleanup.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

