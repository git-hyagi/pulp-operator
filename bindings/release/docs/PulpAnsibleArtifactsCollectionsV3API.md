# \PulpAnsibleArtifactsCollectionsV3API

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyArtifactsCollectionsCreate**](PulpAnsibleArtifactsCollectionsV3API.md#PulpAnsibleGalaxyArtifactsCollectionsCreate) | **Post** /pulp_ansible/galaxy/{path}/api/v3/artifacts/collections/ | Upload a collection
[**PulpAnsibleGalaxyDefaultArtifactsCollectionsCreate**](PulpAnsibleArtifactsCollectionsV3API.md#PulpAnsibleGalaxyDefaultArtifactsCollectionsCreate) | **Post** /pulp_ansible/galaxy/default/api/v3/artifacts/collections/ | Upload a collection
[**PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsArtifactsCreate**](PulpAnsibleArtifactsCollectionsV3API.md#PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsArtifactsCreate) | **Post** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/ | Upload a collection
[**PulpAnsibleGalaxyPluginAnsibleContentCollectionsArtifactsCreate**](PulpAnsibleArtifactsCollectionsV3API.md#PulpAnsibleGalaxyPluginAnsibleContentCollectionsArtifactsCreate) | **Post** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/ | Upload a collection



## PulpAnsibleGalaxyArtifactsCollectionsCreate

> AsyncOperationResponse PulpAnsibleGalaxyArtifactsCollectionsCreate(ctx, path).File(file).XTaskDiagnostics(xTaskDiagnostics).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
	file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
	expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
	expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
	expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyArtifactsCollectionsCreate(context.Background(), path).File(file).XTaskDiagnostics(xTaskDiagnostics).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyArtifactsCollectionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PulpAnsibleGalaxyArtifactsCollectionsCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyArtifactsCollectionsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyArtifactsCollectionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The Collection tarball. | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultArtifactsCollectionsCreate

> AsyncOperationResponse PulpAnsibleGalaxyDefaultArtifactsCollectionsCreate(ctx).File(file).XTaskDiagnostics(xTaskDiagnostics).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
	file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
	expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
	expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
	expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultArtifactsCollectionsCreate(context.Background()).File(file).XTaskDiagnostics(xTaskDiagnostics).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultArtifactsCollectionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PulpAnsibleGalaxyDefaultArtifactsCollectionsCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultArtifactsCollectionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultArtifactsCollectionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The Collection tarball. | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsArtifactsCreate

> AsyncOperationResponse PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsArtifactsCreate(ctx, distroBasePath).File(file).XTaskDiagnostics(xTaskDiagnostics).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
	file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
	expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
	expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
	expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsArtifactsCreate(context.Background(), distroBasePath).File(file).XTaskDiagnostics(xTaskDiagnostics).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsArtifactsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsArtifactsCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsArtifactsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultPluginAnsibleContentCollectionsArtifactsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The Collection tarball. | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyPluginAnsibleContentCollectionsArtifactsCreate

> AsyncOperationResponse PulpAnsibleGalaxyPluginAnsibleContentCollectionsArtifactsCreate(ctx, distroBasePath, path).File(file).XTaskDiagnostics(xTaskDiagnostics).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
	path := "path_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
	expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
	expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
	expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyPluginAnsibleContentCollectionsArtifactsCreate(context.Background(), distroBasePath, path).File(file).XTaskDiagnostics(xTaskDiagnostics).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyPluginAnsibleContentCollectionsArtifactsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PulpAnsibleGalaxyPluginAnsibleContentCollectionsArtifactsCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleArtifactsCollectionsV3API.PulpAnsibleGalaxyPluginAnsibleContentCollectionsArtifactsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyPluginAnsibleContentCollectionsArtifactsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The Collection tarball. | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

