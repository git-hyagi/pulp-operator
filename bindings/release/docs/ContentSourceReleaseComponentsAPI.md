# \ContentSourceReleaseComponentsAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebSourceReleaseComponentsCreate**](ContentSourceReleaseComponentsAPI.md#ContentDebSourceReleaseComponentsCreate) | **Post** /api/v3/content/deb/source_release_components/ | Create a source package release component
[**ContentDebSourceReleaseComponentsList**](ContentSourceReleaseComponentsAPI.md#ContentDebSourceReleaseComponentsList) | **Get** /api/v3/content/deb/source_release_components/ | List source package release components
[**ContentDebSourceReleaseComponentsRead**](ContentSourceReleaseComponentsAPI.md#ContentDebSourceReleaseComponentsRead) | **Get** /{deb_source_package_release_component_href} | Inspect a source package release component
[**ContentDebSourceReleaseComponentsSetLabel**](ContentSourceReleaseComponentsAPI.md#ContentDebSourceReleaseComponentsSetLabel) | **Post** /{deb_source_package_release_component_href}set_label/ | Set a label
[**ContentDebSourceReleaseComponentsUnsetLabel**](ContentSourceReleaseComponentsAPI.md#ContentDebSourceReleaseComponentsUnsetLabel) | **Post** /{deb_source_package_release_component_href}unset_label/ | Unset a label



## ContentDebSourceReleaseComponentsCreate

> DebSourcePackageReleaseComponentResponse ContentDebSourceReleaseComponentsCreate(ctx).DebSourcePackageReleaseComponent(debSourcePackageReleaseComponent).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a source package release component



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
	debSourcePackageReleaseComponent := *openapiclient.NewDebSourcePackageReleaseComponent("SourcePackage_example", "ReleaseComponent_example") // DebSourcePackageReleaseComponent | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsCreate(context.Background()).DebSourcePackageReleaseComponent(debSourcePackageReleaseComponent).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourceReleaseComponentsCreate`: DebSourcePackageReleaseComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourceReleaseComponentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debSourcePackageReleaseComponent** | [**DebSourcePackageReleaseComponent**](DebSourcePackageReleaseComponent.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**DebSourcePackageReleaseComponentResponse**](DebSourcePackageReleaseComponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebSourceReleaseComponentsList

> PaginateddebSourcePackageReleaseComponentResponseList ContentDebSourceReleaseComponentsList(ctx).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).ReleaseComponent(releaseComponent).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).SourcePackage(sourcePackage).Fields(fields).ExcludeFields(excludeFields).Execute()

List source package release components



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
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	releaseComponent := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where release_component matches value (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	sourcePackage := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where source_package matches value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsList(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).ReleaseComponent(releaseComponent).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).SourcePackage(sourcePackage).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourceReleaseComponentsList`: PaginateddebSourcePackageReleaseComponentResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourceReleaseComponentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **releaseComponent** | **string** | Filter results where release_component matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **sourcePackage** | **string** | Filter results where source_package matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebSourcePackageReleaseComponentResponseList**](PaginateddebSourcePackageReleaseComponentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebSourceReleaseComponentsRead

> DebSourcePackageReleaseComponentResponse ContentDebSourceReleaseComponentsRead(ctx, debSourcePackageReleaseComponentHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a source package release component



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
	debSourcePackageReleaseComponentHref := "debSourcePackageReleaseComponentHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsRead(context.Background(), debSourcePackageReleaseComponentHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourceReleaseComponentsRead`: DebSourcePackageReleaseComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debSourcePackageReleaseComponentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourceReleaseComponentsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebSourcePackageReleaseComponentResponse**](DebSourcePackageReleaseComponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebSourceReleaseComponentsSetLabel

> SetLabelResponse ContentDebSourceReleaseComponentsSetLabel(ctx, debSourcePackageReleaseComponentHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

Set a label



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
	debSourcePackageReleaseComponentHref := "debSourcePackageReleaseComponentHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsSetLabel(context.Background(), debSourcePackageReleaseComponentHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourceReleaseComponentsSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debSourcePackageReleaseComponentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourceReleaseComponentsSetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setLabel** | [**SetLabel**](SetLabel.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**SetLabelResponse**](SetLabelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebSourceReleaseComponentsUnsetLabel

> UnsetLabelResponse ContentDebSourceReleaseComponentsUnsetLabel(ctx, debSourcePackageReleaseComponentHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

Unset a label



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
	debSourcePackageReleaseComponentHref := "debSourcePackageReleaseComponentHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsUnsetLabel(context.Background(), debSourcePackageReleaseComponentHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourceReleaseComponentsUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSourceReleaseComponentsAPI.ContentDebSourceReleaseComponentsUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debSourcePackageReleaseComponentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourceReleaseComponentsUnsetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unsetLabel** | [**UnsetLabel**](UnsetLabel.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**UnsetLabelResponse**](UnsetLabelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

