# \ContentSourcePackagesAPI

All URIs are relative to *http://pulp-web-svc-pulp.apps-crc.testing*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebSourcePackagesCreate**](ContentSourcePackagesAPI.md#ContentDebSourcePackagesCreate) | **Post** /api/v3/content/deb/source_packages/ | Create a source package
[**ContentDebSourcePackagesList**](ContentSourcePackagesAPI.md#ContentDebSourcePackagesList) | **Get** /api/v3/content/deb/source_packages/ | List source packages
[**ContentDebSourcePackagesRead**](ContentSourcePackagesAPI.md#ContentDebSourcePackagesRead) | **Get** /{deb_source_package_href} | Inspect a source package
[**ContentDebSourcePackagesSetLabel**](ContentSourcePackagesAPI.md#ContentDebSourcePackagesSetLabel) | **Post** /{deb_source_package_href}set_label/ | Set a label
[**ContentDebSourcePackagesUnsetLabel**](ContentSourcePackagesAPI.md#ContentDebSourcePackagesUnsetLabel) | **Post** /{deb_source_package_href}unset_label/ | Unset a label



## ContentDebSourcePackagesCreate

> AsyncOperationResponse ContentDebSourcePackagesCreate(ctx).DebSourcePackage(debSourcePackage).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a source package



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
	debSourcePackage := *openapiclient.NewDebSourcePackage("Artifact_example") // DebSourcePackage | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourcePackagesAPI.ContentDebSourcePackagesCreate(context.Background()).DebSourcePackage(debSourcePackage).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourcePackagesAPI.ContentDebSourcePackagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourcePackagesCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSourcePackagesAPI.ContentDebSourcePackagesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourcePackagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debSourcePackage** | [**DebSourcePackage**](DebSourcePackage.md) |  | 
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


## ContentDebSourcePackagesList

> PaginateddebSourcePackageResponseList ContentDebSourcePackagesList(ctx).XTaskDiagnostics(xTaskDiagnostics).Architecture(architecture).Binary(binary).BuildConflicts(buildConflicts).BuildConflictsArch(buildConflictsArch).BuildConflictsIndep(buildConflictsIndep).BuildDepends(buildDepends).BuildDependsArch(buildDependsArch).BuildDependsIndep(buildDependsIndep).Dgit(dgit).Format(format).Homepage(homepage).Limit(limit).Maintainer(maintainer).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PackageList(packageList).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RelativePath(relativePath).Release(release).ReleaseComponent(releaseComponent).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Source(source).StandardsVersion(standardsVersion).Testsuite(testsuite).Uploaders(uploaders).VcsArch(vcsArch).VcsBrowser(vcsBrowser).VcsBzr(vcsBzr).VcsCvs(vcsCvs).VcsDarcs(vcsDarcs).VcsGit(vcsGit).VcsHg(vcsHg).VcsMtn(vcsMtn).VcsSnv(vcsSnv).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()

List source packages



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
	binary := "binary_example" // string | Filter results where binary matches value (optional)
	buildConflicts := "buildConflicts_example" // string | Filter results where build_conflicts matches value (optional)
	buildConflictsArch := "buildConflictsArch_example" // string | Filter results where build_conflicts_arch matches value (optional)
	buildConflictsIndep := "buildConflictsIndep_example" // string | Filter results where build_conflicts_indep matches value (optional)
	buildDepends := "buildDepends_example" // string | Filter results where build_depends matches value (optional)
	buildDependsArch := "buildDependsArch_example" // string | Filter results where build_depends_arch matches value (optional)
	buildDependsIndep := "buildDependsIndep_example" // string | Filter results where build_depends_indep matches value (optional)
	dgit := "dgit_example" // string | Filter results where dgit matches value (optional)
	format := "format_example" // string | Filter results where format matches value (optional)
	homepage := "homepage_example" // string | Filter results where homepage matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	maintainer := "maintainer_example" // string | Filter results where maintainer matches value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `relative_path` - Relative path* `-relative_path` - Relative path (descending)* `format` - Format* `-format` - Format (descending)* `source` - Source* `-source` - Source (descending)* `binary` - Binary* `-binary` - Binary (descending)* `architecture` - Architecture* `-architecture` - Architecture (descending)* `version` - Version* `-version` - Version (descending)* `maintainer` - Maintainer* `-maintainer` - Maintainer (descending)* `uploaders` - Uploaders* `-uploaders` - Uploaders (descending)* `homepage` - Homepage* `-homepage` - Homepage (descending)* `vcs_browser` - Vcs browser* `-vcs_browser` - Vcs browser (descending)* `vcs_arch` - Vcs arch* `-vcs_arch` - Vcs arch (descending)* `vcs_bzr` - Vcs bzr* `-vcs_bzr` - Vcs bzr (descending)* `vcs_cvs` - Vcs cvs* `-vcs_cvs` - Vcs cvs (descending)* `vcs_darcs` - Vcs darcs* `-vcs_darcs` - Vcs darcs (descending)* `vcs_git` - Vcs git* `-vcs_git` - Vcs git (descending)* `vcs_hg` - Vcs hg* `-vcs_hg` - Vcs hg (descending)* `vcs_mtn` - Vcs mtn* `-vcs_mtn` - Vcs mtn (descending)* `vcs_snv` - Vcs snv* `-vcs_snv` - Vcs snv (descending)* `testsuite` - Testsuite* `-testsuite` - Testsuite (descending)* `dgit` - Dgit* `-dgit` - Dgit (descending)* `standards_version` - Standards version* `-standards_version` - Standards version (descending)* `build_depends` - Build depends* `-build_depends` - Build depends (descending)* `build_depends_indep` - Build depends indep* `-build_depends_indep` - Build depends indep (descending)* `build_depends_arch` - Build depends arch* `-build_depends_arch` - Build depends arch (descending)* `build_conflicts` - Build conflicts* `-build_conflicts` - Build conflicts (descending)* `build_conflicts_indep` - Build conflicts indep* `-build_conflicts_indep` - Build conflicts indep (descending)* `build_conflicts_arch` - Build conflicts arch* `-build_conflicts_arch` - Build conflicts arch (descending)* `package_list` - Package list* `-package_list` - Package list (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	packageList := "packageList_example" // string | Filter results where package_list matches value (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	relativePath := "relativePath_example" // string | Filter results where relative_path matches value (optional)
	release := "release_example" // string |     Must be a comma-separated string: \"release_href,repository_or_repository_version_href\"    release_href: Filter results where SourcePackage in Release    repository_or_repository_version_href: The RepositoryVersion href to filter by, or Repository        href (assume latest version)     (optional)
	releaseComponent := "releaseComponent_example" // string |     Must be a comma-separated string: \"release_component_href,repository_or_repository_version_href\"    release_component_href: Filter results where SourcePackage in ReleaseComponent    repository_or_repository_version_href: The RepositoryVersion href to filter by, or Repository        href (assume latest version)     (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	source := "source_example" // string | Filter results where source matches value (optional)
	standardsVersion := "standardsVersion_example" // string | Filter results where standards_version matches value (optional)
	testsuite := "testsuite_example" // string | Filter results where testsuite matches value (optional)
	uploaders := "uploaders_example" // string | Filter results where uploaders matches value (optional)
	vcsArch := "vcsArch_example" // string | Filter results where vcs_arch matches value (optional)
	vcsBrowser := "vcsBrowser_example" // string | Filter results where vcs_browser matches value (optional)
	vcsBzr := "vcsBzr_example" // string | Filter results where vcs_bzr matches value (optional)
	vcsCvs := "vcsCvs_example" // string | Filter results where vcs_cvs matches value (optional)
	vcsDarcs := "vcsDarcs_example" // string | Filter results where vcs_darcs matches value (optional)
	vcsGit := "vcsGit_example" // string | Filter results where vcs_git matches value (optional)
	vcsHg := "vcsHg_example" // string | Filter results where vcs_hg matches value (optional)
	vcsMtn := "vcsMtn_example" // string | Filter results where vcs_mtn matches value (optional)
	vcsSnv := "vcsSnv_example" // string | Filter results where vcs_snv matches value (optional)
	version := "version_example" // string | Filter results where version matches value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourcePackagesAPI.ContentDebSourcePackagesList(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Architecture(architecture).Binary(binary).BuildConflicts(buildConflicts).BuildConflictsArch(buildConflictsArch).BuildConflictsIndep(buildConflictsIndep).BuildDepends(buildDepends).BuildDependsArch(buildDependsArch).BuildDependsIndep(buildDependsIndep).Dgit(dgit).Format(format).Homepage(homepage).Limit(limit).Maintainer(maintainer).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PackageList(packageList).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RelativePath(relativePath).Release(release).ReleaseComponent(releaseComponent).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Source(source).StandardsVersion(standardsVersion).Testsuite(testsuite).Uploaders(uploaders).VcsArch(vcsArch).VcsBrowser(vcsBrowser).VcsBzr(vcsBzr).VcsCvs(vcsCvs).VcsDarcs(vcsDarcs).VcsGit(vcsGit).VcsHg(vcsHg).VcsMtn(vcsMtn).VcsSnv(vcsSnv).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourcePackagesAPI.ContentDebSourcePackagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourcePackagesList`: PaginateddebSourcePackageResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentSourcePackagesAPI.ContentDebSourcePackagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourcePackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **architecture** | **string** | Filter results where architecture matches value | 
 **binary** | **string** | Filter results where binary matches value | 
 **buildConflicts** | **string** | Filter results where build_conflicts matches value | 
 **buildConflictsArch** | **string** | Filter results where build_conflicts_arch matches value | 
 **buildConflictsIndep** | **string** | Filter results where build_conflicts_indep matches value | 
 **buildDepends** | **string** | Filter results where build_depends matches value | 
 **buildDependsArch** | **string** | Filter results where build_depends_arch matches value | 
 **buildDependsIndep** | **string** | Filter results where build_depends_indep matches value | 
 **dgit** | **string** | Filter results where dgit matches value | 
 **format** | **string** | Filter results where format matches value | 
 **homepage** | **string** | Filter results where homepage matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **maintainer** | **string** | Filter results where maintainer matches value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;relative_path&#x60; - Relative path* &#x60;-relative_path&#x60; - Relative path (descending)* &#x60;format&#x60; - Format* &#x60;-format&#x60; - Format (descending)* &#x60;source&#x60; - Source* &#x60;-source&#x60; - Source (descending)* &#x60;binary&#x60; - Binary* &#x60;-binary&#x60; - Binary (descending)* &#x60;architecture&#x60; - Architecture* &#x60;-architecture&#x60; - Architecture (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;maintainer&#x60; - Maintainer* &#x60;-maintainer&#x60; - Maintainer (descending)* &#x60;uploaders&#x60; - Uploaders* &#x60;-uploaders&#x60; - Uploaders (descending)* &#x60;homepage&#x60; - Homepage* &#x60;-homepage&#x60; - Homepage (descending)* &#x60;vcs_browser&#x60; - Vcs browser* &#x60;-vcs_browser&#x60; - Vcs browser (descending)* &#x60;vcs_arch&#x60; - Vcs arch* &#x60;-vcs_arch&#x60; - Vcs arch (descending)* &#x60;vcs_bzr&#x60; - Vcs bzr* &#x60;-vcs_bzr&#x60; - Vcs bzr (descending)* &#x60;vcs_cvs&#x60; - Vcs cvs* &#x60;-vcs_cvs&#x60; - Vcs cvs (descending)* &#x60;vcs_darcs&#x60; - Vcs darcs* &#x60;-vcs_darcs&#x60; - Vcs darcs (descending)* &#x60;vcs_git&#x60; - Vcs git* &#x60;-vcs_git&#x60; - Vcs git (descending)* &#x60;vcs_hg&#x60; - Vcs hg* &#x60;-vcs_hg&#x60; - Vcs hg (descending)* &#x60;vcs_mtn&#x60; - Vcs mtn* &#x60;-vcs_mtn&#x60; - Vcs mtn (descending)* &#x60;vcs_snv&#x60; - Vcs snv* &#x60;-vcs_snv&#x60; - Vcs snv (descending)* &#x60;testsuite&#x60; - Testsuite* &#x60;-testsuite&#x60; - Testsuite (descending)* &#x60;dgit&#x60; - Dgit* &#x60;-dgit&#x60; - Dgit (descending)* &#x60;standards_version&#x60; - Standards version* &#x60;-standards_version&#x60; - Standards version (descending)* &#x60;build_depends&#x60; - Build depends* &#x60;-build_depends&#x60; - Build depends (descending)* &#x60;build_depends_indep&#x60; - Build depends indep* &#x60;-build_depends_indep&#x60; - Build depends indep (descending)* &#x60;build_depends_arch&#x60; - Build depends arch* &#x60;-build_depends_arch&#x60; - Build depends arch (descending)* &#x60;build_conflicts&#x60; - Build conflicts* &#x60;-build_conflicts&#x60; - Build conflicts (descending)* &#x60;build_conflicts_indep&#x60; - Build conflicts indep* &#x60;-build_conflicts_indep&#x60; - Build conflicts indep (descending)* &#x60;build_conflicts_arch&#x60; - Build conflicts arch* &#x60;-build_conflicts_arch&#x60; - Build conflicts arch (descending)* &#x60;package_list&#x60; - Package list* &#x60;-package_list&#x60; - Package list (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **packageList** | **string** | Filter results where package_list matches value | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **relativePath** | **string** | Filter results where relative_path matches value | 
 **release** | **string** |     Must be a comma-separated string: \&quot;release_href,repository_or_repository_version_href\&quot;    release_href: Filter results where SourcePackage in Release    repository_or_repository_version_href: The RepositoryVersion href to filter by, or Repository        href (assume latest version)     | 
 **releaseComponent** | **string** |     Must be a comma-separated string: \&quot;release_component_href,repository_or_repository_version_href\&quot;    release_component_href: Filter results where SourcePackage in ReleaseComponent    repository_or_repository_version_href: The RepositoryVersion href to filter by, or Repository        href (assume latest version)     | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **source** | **string** | Filter results where source matches value | 
 **standardsVersion** | **string** | Filter results where standards_version matches value | 
 **testsuite** | **string** | Filter results where testsuite matches value | 
 **uploaders** | **string** | Filter results where uploaders matches value | 
 **vcsArch** | **string** | Filter results where vcs_arch matches value | 
 **vcsBrowser** | **string** | Filter results where vcs_browser matches value | 
 **vcsBzr** | **string** | Filter results where vcs_bzr matches value | 
 **vcsCvs** | **string** | Filter results where vcs_cvs matches value | 
 **vcsDarcs** | **string** | Filter results where vcs_darcs matches value | 
 **vcsGit** | **string** | Filter results where vcs_git matches value | 
 **vcsHg** | **string** | Filter results where vcs_hg matches value | 
 **vcsMtn** | **string** | Filter results where vcs_mtn matches value | 
 **vcsSnv** | **string** | Filter results where vcs_snv matches value | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebSourcePackageResponseList**](PaginateddebSourcePackageResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebSourcePackagesRead

> DebSourcePackageResponse ContentDebSourcePackagesRead(ctx, debSourcePackageHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a source package



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
	debSourcePackageHref := "debSourcePackageHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourcePackagesAPI.ContentDebSourcePackagesRead(context.Background(), debSourcePackageHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourcePackagesAPI.ContentDebSourcePackagesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourcePackagesRead`: DebSourcePackageResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSourcePackagesAPI.ContentDebSourcePackagesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debSourcePackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourcePackagesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebSourcePackageResponse**](DebSourcePackageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebSourcePackagesSetLabel

> SetLabelResponse ContentDebSourcePackagesSetLabel(ctx, debSourcePackageHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	debSourcePackageHref := "debSourcePackageHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourcePackagesAPI.ContentDebSourcePackagesSetLabel(context.Background(), debSourcePackageHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourcePackagesAPI.ContentDebSourcePackagesSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourcePackagesSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSourcePackagesAPI.ContentDebSourcePackagesSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debSourcePackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourcePackagesSetLabelRequest struct via the builder pattern


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


## ContentDebSourcePackagesUnsetLabel

> UnsetLabelResponse ContentDebSourcePackagesUnsetLabel(ctx, debSourcePackageHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	debSourcePackageHref := "debSourcePackageHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSourcePackagesAPI.ContentDebSourcePackagesUnsetLabel(context.Background(), debSourcePackageHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSourcePackagesAPI.ContentDebSourcePackagesUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentDebSourcePackagesUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSourcePackagesAPI.ContentDebSourcePackagesUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debSourcePackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebSourcePackagesUnsetLabelRequest struct via the builder pattern


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

