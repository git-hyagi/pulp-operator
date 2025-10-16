# \PulpAnsibleDefaultPluginAnsibleClientConfigurationAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultPluginAnsibleClientConfigurationRead**](PulpAnsibleDefaultPluginAnsibleClientConfigurationAPI.md#PulpAnsibleGalaxyDefaultPluginAnsibleClientConfigurationRead) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/client-configuration/ | 



## PulpAnsibleGalaxyDefaultPluginAnsibleClientConfigurationRead

> ClientConfigurationResponse PulpAnsibleGalaxyDefaultPluginAnsibleClientConfigurationRead(ctx).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()





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
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PulpAnsibleDefaultPluginAnsibleClientConfigurationAPI.PulpAnsibleGalaxyDefaultPluginAnsibleClientConfigurationRead(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultPluginAnsibleClientConfigurationAPI.PulpAnsibleGalaxyDefaultPluginAnsibleClientConfigurationRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PulpAnsibleGalaxyDefaultPluginAnsibleClientConfigurationRead`: ClientConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultPluginAnsibleClientConfigurationAPI.PulpAnsibleGalaxyDefaultPluginAnsibleClientConfigurationRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultPluginAnsibleClientConfigurationReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ClientConfigurationResponse**](ClientConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

