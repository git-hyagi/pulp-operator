# \ContentInstallerPackagesAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebInstallerPackagesCreate**](ContentInstallerPackagesAPI.md#ContentDebInstallerPackagesCreate) | **Post** /api/v3/content/deb/installer_packages/ | Create an installer package
[**ContentDebInstallerPackagesList**](ContentInstallerPackagesAPI.md#ContentDebInstallerPackagesList) | **Get** /api/v3/content/deb/installer_packages/ | List installer packages
[**ContentDebInstallerPackagesRead**](ContentInstallerPackagesAPI.md#ContentDebInstallerPackagesRead) | **Get** /{deb_installer_package_href} | Inspect an installer package
[**ContentDebInstallerPackagesSetLabel**](ContentInstallerPackagesAPI.md#ContentDebInstallerPackagesSetLabel) | **Post** /{deb_installer_package_href}set_label/ | Set a label
[**ContentDebInstallerPackagesUnsetLabel**](ContentInstallerPackagesAPI.md#ContentDebInstallerPackagesUnsetLabel) | **Post** /{deb_installer_package_href}unset_label/ | Unset a label



## ContentDebInstallerPackagesCreate

> AsyncOperationResponse ContentDebInstallerPackagesCreate(ctx).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).RelativePath(relativePath).File(file).Upload(upload).FileUrl(fileUrl).Execute()

Create an installer package



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
	repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
	pulpLabels := map[string]string{"key": "Inner_example"} // map[string]string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
	relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentInstallerPackagesAPI.ContentDebInstallerPackagesCreate(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).RelativePath(relativePath).File(file).Upload(upload).FileUrl(fileUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerPackagesAPI.ContentDebInstallerPackagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebInstallerPackagesCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentInstallerPackagesAPI.ContentDebInstallerPackagesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerPackagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
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


## ContentDebInstallerPackagesList

> PaginateddebInstallerPackageResponseList ContentDebInstallerPackagesList(ctx).XTaskDiagnostics(xTaskDiagnostics).Architecture(architecture).AutoBuiltPackage(autoBuiltPackage).BuildEssential(buildEssential).BuiltUsing(builtUsing).Essential(essential).InstalledSize(installedSize).Limit(limit).Maintainer(maintainer).MultiArch(multiArch).Offset(offset).Ordering(ordering).Origin(origin).OriginalMaintainer(originalMaintainer).OrphanedFor(orphanedFor).Package_(package_).Priority(priority).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Section(section).Sha256(sha256).Source(source).Tag(tag).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()

List installer packages



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
	architecture := "architecture_example" // string | Filter results where architecture matches value (optional)
	autoBuiltPackage := "autoBuiltPackage_example" // string | Filter results where auto_built_package matches value (optional)
	buildEssential := true // bool | Filter results where build_essential matches value* `True` - yes* `False` - no (optional)
	builtUsing := "builtUsing_example" // string | Filter results where built_using matches value (optional)
	essential := true // bool | Filter results where essential matches value* `True` - yes* `False` - no (optional)
	installedSize := int32(56) // int32 | Filter results where installed_size matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	maintainer := "maintainer_example" // string | Filter results where maintainer matches value (optional)
	multiArch := "multiArch_example" // string | Filter results where multi_arch matches value* `no` - no* `same` - same* `foreign` - foreign* `allowed` - allowed (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `package` - Package* `-package` - Package (descending)* `source` - Source* `-source` - Source (descending)* `version` - Version* `-version` - Version (descending)* `architecture` - Architecture* `-architecture` - Architecture (descending)* `section` - Section* `-section` - Section (descending)* `priority` - Priority* `-priority` - Priority (descending)* `origin` - Origin* `-origin` - Origin (descending)* `tag` - Tag* `-tag` - Tag (descending)* `bugs` - Bugs* `-bugs` - Bugs (descending)* `essential` - Essential* `-essential` - Essential (descending)* `build_essential` - Build essential* `-build_essential` - Build essential (descending)* `installed_size` - Installed size* `-installed_size` - Installed size (descending)* `maintainer` - Maintainer* `-maintainer` - Maintainer (descending)* `original_maintainer` - Original maintainer* `-original_maintainer` - Original maintainer (descending)* `description` - Description* `-description` - Description (descending)* `description_md5` - Description md5* `-description_md5` - Description md5 (descending)* `homepage` - Homepage* `-homepage` - Homepage (descending)* `built_using` - Built using* `-built_using` - Built using (descending)* `auto_built_package` - Auto built package* `-auto_built_package` - Auto built package (descending)* `multi_arch` - Multi arch* `-multi_arch` - Multi arch (descending)* `breaks` - Breaks* `-breaks` - Breaks (descending)* `conflicts` - Conflicts* `-conflicts` - Conflicts (descending)* `depends` - Depends* `-depends` - Depends (descending)* `recommends` - Recommends* `-recommends` - Recommends (descending)* `suggests` - Suggests* `-suggests` - Suggests (descending)* `enhances` - Enhances* `-enhances` - Enhances (descending)* `pre_depends` - Pre depends* `-pre_depends` - Pre depends (descending)* `provides` - Provides* `-provides` - Provides (descending)* `replaces` - Replaces* `-replaces` - Replaces (descending)* `relative_path` - Relative path* `-relative_path` - Relative path (descending)* `sha256` - Sha256* `-sha256` - Sha256 (descending)* `custom_fields` - Custom fields* `-custom_fields` - Custom fields (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	origin := "origin_example" // string | Filter results where origin matches value (optional)
	originalMaintainer := "originalMaintainer_example" // string | Filter results where original_maintainer matches value (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	package_ := "package__example" // string | Filter results where package matches value (optional)
	priority := "priority_example" // string | Filter results where priority matches value (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	section := "section_example" // string | Filter results where section matches value (optional)
	sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
	source := "source_example" // string | Filter results where source matches value (optional)
	tag := "tag_example" // string | Filter results where tag matches value (optional)
	version := "version_example" // string | Filter results where version matches value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentInstallerPackagesAPI.ContentDebInstallerPackagesList(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Architecture(architecture).AutoBuiltPackage(autoBuiltPackage).BuildEssential(buildEssential).BuiltUsing(builtUsing).Essential(essential).InstalledSize(installedSize).Limit(limit).Maintainer(maintainer).MultiArch(multiArch).Offset(offset).Ordering(ordering).Origin(origin).OriginalMaintainer(originalMaintainer).OrphanedFor(orphanedFor).Package_(package_).Priority(priority).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Section(section).Sha256(sha256).Source(source).Tag(tag).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerPackagesAPI.ContentDebInstallerPackagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebInstallerPackagesList`: PaginateddebInstallerPackageResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentInstallerPackagesAPI.ContentDebInstallerPackagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **architecture** | **string** | Filter results where architecture matches value | 
 **autoBuiltPackage** | **string** | Filter results where auto_built_package matches value | 
 **buildEssential** | **bool** | Filter results where build_essential matches value* &#x60;True&#x60; - yes* &#x60;False&#x60; - no | 
 **builtUsing** | **string** | Filter results where built_using matches value | 
 **essential** | **bool** | Filter results where essential matches value* &#x60;True&#x60; - yes* &#x60;False&#x60; - no | 
 **installedSize** | **int32** | Filter results where installed_size matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **maintainer** | **string** | Filter results where maintainer matches value | 
 **multiArch** | **string** | Filter results where multi_arch matches value* &#x60;no&#x60; - no* &#x60;same&#x60; - same* &#x60;foreign&#x60; - foreign* &#x60;allowed&#x60; - allowed | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;package&#x60; - Package* &#x60;-package&#x60; - Package (descending)* &#x60;source&#x60; - Source* &#x60;-source&#x60; - Source (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;architecture&#x60; - Architecture* &#x60;-architecture&#x60; - Architecture (descending)* &#x60;section&#x60; - Section* &#x60;-section&#x60; - Section (descending)* &#x60;priority&#x60; - Priority* &#x60;-priority&#x60; - Priority (descending)* &#x60;origin&#x60; - Origin* &#x60;-origin&#x60; - Origin (descending)* &#x60;tag&#x60; - Tag* &#x60;-tag&#x60; - Tag (descending)* &#x60;bugs&#x60; - Bugs* &#x60;-bugs&#x60; - Bugs (descending)* &#x60;essential&#x60; - Essential* &#x60;-essential&#x60; - Essential (descending)* &#x60;build_essential&#x60; - Build essential* &#x60;-build_essential&#x60; - Build essential (descending)* &#x60;installed_size&#x60; - Installed size* &#x60;-installed_size&#x60; - Installed size (descending)* &#x60;maintainer&#x60; - Maintainer* &#x60;-maintainer&#x60; - Maintainer (descending)* &#x60;original_maintainer&#x60; - Original maintainer* &#x60;-original_maintainer&#x60; - Original maintainer (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;description_md5&#x60; - Description md5* &#x60;-description_md5&#x60; - Description md5 (descending)* &#x60;homepage&#x60; - Homepage* &#x60;-homepage&#x60; - Homepage (descending)* &#x60;built_using&#x60; - Built using* &#x60;-built_using&#x60; - Built using (descending)* &#x60;auto_built_package&#x60; - Auto built package* &#x60;-auto_built_package&#x60; - Auto built package (descending)* &#x60;multi_arch&#x60; - Multi arch* &#x60;-multi_arch&#x60; - Multi arch (descending)* &#x60;breaks&#x60; - Breaks* &#x60;-breaks&#x60; - Breaks (descending)* &#x60;conflicts&#x60; - Conflicts* &#x60;-conflicts&#x60; - Conflicts (descending)* &#x60;depends&#x60; - Depends* &#x60;-depends&#x60; - Depends (descending)* &#x60;recommends&#x60; - Recommends* &#x60;-recommends&#x60; - Recommends (descending)* &#x60;suggests&#x60; - Suggests* &#x60;-suggests&#x60; - Suggests (descending)* &#x60;enhances&#x60; - Enhances* &#x60;-enhances&#x60; - Enhances (descending)* &#x60;pre_depends&#x60; - Pre depends* &#x60;-pre_depends&#x60; - Pre depends (descending)* &#x60;provides&#x60; - Provides* &#x60;-provides&#x60; - Provides (descending)* &#x60;replaces&#x60; - Replaces* &#x60;-replaces&#x60; - Replaces (descending)* &#x60;relative_path&#x60; - Relative path* &#x60;-relative_path&#x60; - Relative path (descending)* &#x60;sha256&#x60; - Sha256* &#x60;-sha256&#x60; - Sha256 (descending)* &#x60;custom_fields&#x60; - Custom fields* &#x60;-custom_fields&#x60; - Custom fields (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **origin** | **string** | Filter results where origin matches value | 
 **originalMaintainer** | **string** | Filter results where original_maintainer matches value | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **package_** | **string** | Filter results where package matches value | 
 **priority** | **string** | Filter results where priority matches value | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **section** | **string** | Filter results where section matches value | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **source** | **string** | Filter results where source matches value | 
 **tag** | **string** | Filter results where tag matches value | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebInstallerPackageResponseList**](PaginateddebInstallerPackageResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebInstallerPackagesRead

> DebInstallerPackageResponse ContentDebInstallerPackagesRead(ctx, debInstallerPackageHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an installer package



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
	debInstallerPackageHref := "debInstallerPackageHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentInstallerPackagesAPI.ContentDebInstallerPackagesRead(context.Background(), debInstallerPackageHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerPackagesAPI.ContentDebInstallerPackagesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebInstallerPackagesRead`: DebInstallerPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentInstallerPackagesAPI.ContentDebInstallerPackagesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debInstallerPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerPackagesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebInstallerPackageResponse**](DebInstallerPackageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebInstallerPackagesSetLabel

> SetLabelResponse ContentDebInstallerPackagesSetLabel(ctx, debInstallerPackageHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	debInstallerPackageHref := "debInstallerPackageHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentInstallerPackagesAPI.ContentDebInstallerPackagesSetLabel(context.Background(), debInstallerPackageHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerPackagesAPI.ContentDebInstallerPackagesSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebInstallerPackagesSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentInstallerPackagesAPI.ContentDebInstallerPackagesSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debInstallerPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerPackagesSetLabelRequest struct via the builder pattern


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


## ContentDebInstallerPackagesUnsetLabel

> UnsetLabelResponse ContentDebInstallerPackagesUnsetLabel(ctx, debInstallerPackageHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	debInstallerPackageHref := "debInstallerPackageHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentInstallerPackagesAPI.ContentDebInstallerPackagesUnsetLabel(context.Background(), debInstallerPackageHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentInstallerPackagesAPI.ContentDebInstallerPackagesUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebInstallerPackagesUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentInstallerPackagesAPI.ContentDebInstallerPackagesUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debInstallerPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebInstallerPackagesUnsetLabelRequest struct via the builder pattern


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

