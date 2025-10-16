# \ContentCollectionSignaturesAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentAnsibleCollectionSignaturesCreate**](ContentCollectionSignaturesAPI.md#ContentAnsibleCollectionSignaturesCreate) | **Post** /api/v3/content/ansible/collection_signatures/ | Create a collection version signature
[**ContentAnsibleCollectionSignaturesList**](ContentCollectionSignaturesAPI.md#ContentAnsibleCollectionSignaturesList) | **Get** /api/v3/content/ansible/collection_signatures/ | List collection version signatures
[**ContentAnsibleCollectionSignaturesRead**](ContentCollectionSignaturesAPI.md#ContentAnsibleCollectionSignaturesRead) | **Get** /{ansible_collection_version_signature_href} | Inspect a collection version signature
[**ContentAnsibleCollectionSignaturesSetLabel**](ContentCollectionSignaturesAPI.md#ContentAnsibleCollectionSignaturesSetLabel) | **Post** /{ansible_collection_version_signature_href}set_label/ | Set a label
[**ContentAnsibleCollectionSignaturesUnsetLabel**](ContentCollectionSignaturesAPI.md#ContentAnsibleCollectionSignaturesUnsetLabel) | **Post** /{ansible_collection_version_signature_href}unset_label/ | Unset a label



## ContentAnsibleCollectionSignaturesCreate

> AsyncOperationResponse ContentAnsibleCollectionSignaturesCreate(ctx).SignedCollection(signedCollection).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).PulpLabels(pulpLabels).File(file).Upload(upload).FileUrl(fileUrl).Execute()

Create a collection version signature



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
	signedCollection := "signedCollection_example" // string | The content this signature is pointing to.
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
	pulpLabels := map[string]string{"key": "Inner_example"} // map[string]string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesCreate(context.Background()).SignedCollection(signedCollection).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).PulpLabels(pulpLabels).File(file).Upload(upload).FileUrl(fileUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentAnsibleCollectionSignaturesCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signedCollection** | **string** | The content this signature is pointing to. | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **file** | ***os.File** | An uploaded file that may be turned into the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the content unit. | 
 **fileUrl** | **string** | A url that Pulp can download and turn into the content unit. | 

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


## ContentAnsibleCollectionSignaturesList

> PaginatedansibleCollectionVersionSignatureResponseList ContentAnsibleCollectionSignaturesList(ctx).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PubkeyFingerprint(pubkeyFingerprint).PubkeyFingerprintIn(pubkeyFingerprintIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).SignedCollection(signedCollection).SigningService(signingService).Fields(fields).ExcludeFields(excludeFields).Execute()

List collection version signatures



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
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `data` - Data* `-data` - Data (descending)* `digest` - Digest* `-digest` - Digest (descending)* `pubkey_fingerprint` - Pubkey fingerprint* `-pubkey_fingerprint` - Pubkey fingerprint (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pubkeyFingerprint := "pubkeyFingerprint_example" // string | Filter results where pubkey_fingerprint matches value (optional)
	pubkeyFingerprintIn := []string{"Inner_example"} // []string | Filter results where pubkey_fingerprint is in a comma-separated list of values (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	signedCollection := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter signatures for collection version (optional)
	signingService := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter signatures produced by signature service (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesList(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PubkeyFingerprint(pubkeyFingerprint).PubkeyFingerprintIn(pubkeyFingerprintIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).SignedCollection(signedCollection).SigningService(signingService).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentAnsibleCollectionSignaturesList`: PaginatedansibleCollectionVersionSignatureResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;data&#x60; - Data* &#x60;-data&#x60; - Data (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;pubkey_fingerprint&#x60; - Pubkey fingerprint* &#x60;-pubkey_fingerprint&#x60; - Pubkey fingerprint (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pubkeyFingerprint** | **string** | Filter results where pubkey_fingerprint matches value | 
 **pubkeyFingerprintIn** | **[]string** | Filter results where pubkey_fingerprint is in a comma-separated list of values | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **signedCollection** | **string** | Filter signatures for collection version | 
 **signingService** | **string** | Filter signatures produced by signature service | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleCollectionVersionSignatureResponseList**](PaginatedansibleCollectionVersionSignatureResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAnsibleCollectionSignaturesRead

> AnsibleCollectionVersionSignatureResponse ContentAnsibleCollectionSignaturesRead(ctx, ansibleCollectionVersionSignatureHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a collection version signature



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
	ansibleCollectionVersionSignatureHref := "ansibleCollectionVersionSignatureHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesRead(context.Background(), ansibleCollectionVersionSignatureHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentAnsibleCollectionSignaturesRead`: AnsibleCollectionVersionSignatureResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionVersionSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleCollectionVersionSignatureResponse**](AnsibleCollectionVersionSignatureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAnsibleCollectionSignaturesSetLabel

> SetLabelResponse ContentAnsibleCollectionSignaturesSetLabel(ctx, ansibleCollectionVersionSignatureHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	ansibleCollectionVersionSignatureHref := "ansibleCollectionVersionSignatureHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesSetLabel(context.Background(), ansibleCollectionVersionSignatureHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentAnsibleCollectionSignaturesSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionVersionSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesSetLabelRequest struct via the builder pattern


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


## ContentAnsibleCollectionSignaturesUnsetLabel

> UnsetLabelResponse ContentAnsibleCollectionSignaturesUnsetLabel(ctx, ansibleCollectionVersionSignatureHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	ansibleCollectionVersionSignatureHref := "ansibleCollectionVersionSignatureHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesUnsetLabel(context.Background(), ansibleCollectionVersionSignatureHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentAnsibleCollectionSignaturesUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionVersionSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesUnsetLabelRequest struct via the builder pattern


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

