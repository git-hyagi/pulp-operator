# DebSourcePackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**VulnReport** | Pointer to **string** |  | [optional] [readonly] 
**Artifacts** | **map[string]interface{}** | A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {&#39;relative/path&#39;: &#39;/artifacts/1/&#39; | [readonly] 
**RelativePath** | Pointer to **string** | Relative path of the Debian Source Control (dsc) file.It is normally advised to let Pulp generate this. | [optional] 
**Sha256** | Pointer to **string** | sha256 digest of the dsc file. | [optional] [readonly] 
**Format** | Pointer to **string** |  | [optional] [readonly] 
**Source** | Pointer to **string** |  | [optional] [readonly] 
**Binary** | Pointer to **string** |  | [optional] [readonly] 
**Architecture** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**Maintainer** | Pointer to **string** |  | [optional] [readonly] 
**Uploaders** | Pointer to **string** |  | [optional] [readonly] 
**Homepage** | Pointer to **string** |  | [optional] [readonly] 
**VcsBrowser** | Pointer to **string** |  | [optional] [readonly] 
**VcsArch** | Pointer to **string** |  | [optional] [readonly] 
**VcsBzr** | Pointer to **string** |  | [optional] [readonly] 
**VcsCvs** | Pointer to **string** |  | [optional] [readonly] 
**VcsDarcs** | Pointer to **string** |  | [optional] [readonly] 
**VcsGit** | Pointer to **string** |  | [optional] [readonly] 
**VcsHg** | Pointer to **string** |  | [optional] [readonly] 
**VcsMtn** | Pointer to **string** |  | [optional] [readonly] 
**VcsSnv** | Pointer to **string** |  | [optional] [readonly] 
**Testsuite** | Pointer to **string** |  | [optional] [readonly] 
**Dgit** | Pointer to **string** |  | [optional] [readonly] 
**StandardsVersion** | Pointer to **string** |  | [optional] [readonly] 
**BuildDepends** | Pointer to **string** |  | [optional] [readonly] 
**BuildDependsIndep** | Pointer to **string** |  | [optional] [readonly] 
**BuildDependsArch** | Pointer to **string** |  | [optional] [readonly] 
**BuildConflicts** | Pointer to **string** |  | [optional] [readonly] 
**BuildConflictsIndep** | Pointer to **string** |  | [optional] [readonly] 
**BuildConflictsArch** | Pointer to **string** |  | [optional] [readonly] 
**PackageList** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDebSourcePackageResponse

`func NewDebSourcePackageResponse(artifacts map[string]interface{}, ) *DebSourcePackageResponse`

NewDebSourcePackageResponse instantiates a new DebSourcePackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebSourcePackageResponseWithDefaults

`func NewDebSourcePackageResponseWithDefaults() *DebSourcePackageResponse`

NewDebSourcePackageResponseWithDefaults instantiates a new DebSourcePackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebSourcePackageResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebSourcePackageResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebSourcePackageResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebSourcePackageResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *DebSourcePackageResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *DebSourcePackageResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *DebSourcePackageResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *DebSourcePackageResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebSourcePackageResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebSourcePackageResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebSourcePackageResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebSourcePackageResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *DebSourcePackageResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *DebSourcePackageResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *DebSourcePackageResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *DebSourcePackageResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *DebSourcePackageResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *DebSourcePackageResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *DebSourcePackageResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *DebSourcePackageResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetVulnReport

`func (o *DebSourcePackageResponse) GetVulnReport() string`

GetVulnReport returns the VulnReport field if non-nil, zero value otherwise.

### GetVulnReportOk

`func (o *DebSourcePackageResponse) GetVulnReportOk() (*string, bool)`

GetVulnReportOk returns a tuple with the VulnReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnReport

`func (o *DebSourcePackageResponse) SetVulnReport(v string)`

SetVulnReport sets VulnReport field to given value.

### HasVulnReport

`func (o *DebSourcePackageResponse) HasVulnReport() bool`

HasVulnReport returns a boolean if a field has been set.

### GetArtifacts

`func (o *DebSourcePackageResponse) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *DebSourcePackageResponse) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *DebSourcePackageResponse) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.


### GetRelativePath

`func (o *DebSourcePackageResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DebSourcePackageResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DebSourcePackageResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DebSourcePackageResponse) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetSha256

`func (o *DebSourcePackageResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *DebSourcePackageResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *DebSourcePackageResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *DebSourcePackageResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetFormat

`func (o *DebSourcePackageResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DebSourcePackageResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DebSourcePackageResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DebSourcePackageResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSource

`func (o *DebSourcePackageResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DebSourcePackageResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DebSourcePackageResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DebSourcePackageResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetBinary

`func (o *DebSourcePackageResponse) GetBinary() string`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *DebSourcePackageResponse) GetBinaryOk() (*string, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *DebSourcePackageResponse) SetBinary(v string)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *DebSourcePackageResponse) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetArchitecture

`func (o *DebSourcePackageResponse) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebSourcePackageResponse) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebSourcePackageResponse) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *DebSourcePackageResponse) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetVersion

`func (o *DebSourcePackageResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DebSourcePackageResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DebSourcePackageResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DebSourcePackageResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMaintainer

`func (o *DebSourcePackageResponse) GetMaintainer() string`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *DebSourcePackageResponse) GetMaintainerOk() (*string, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *DebSourcePackageResponse) SetMaintainer(v string)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *DebSourcePackageResponse) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetUploaders

`func (o *DebSourcePackageResponse) GetUploaders() string`

GetUploaders returns the Uploaders field if non-nil, zero value otherwise.

### GetUploadersOk

`func (o *DebSourcePackageResponse) GetUploadersOk() (*string, bool)`

GetUploadersOk returns a tuple with the Uploaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaders

`func (o *DebSourcePackageResponse) SetUploaders(v string)`

SetUploaders sets Uploaders field to given value.

### HasUploaders

`func (o *DebSourcePackageResponse) HasUploaders() bool`

HasUploaders returns a boolean if a field has been set.

### GetHomepage

`func (o *DebSourcePackageResponse) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *DebSourcePackageResponse) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *DebSourcePackageResponse) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *DebSourcePackageResponse) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetVcsBrowser

`func (o *DebSourcePackageResponse) GetVcsBrowser() string`

GetVcsBrowser returns the VcsBrowser field if non-nil, zero value otherwise.

### GetVcsBrowserOk

`func (o *DebSourcePackageResponse) GetVcsBrowserOk() (*string, bool)`

GetVcsBrowserOk returns a tuple with the VcsBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsBrowser

`func (o *DebSourcePackageResponse) SetVcsBrowser(v string)`

SetVcsBrowser sets VcsBrowser field to given value.

### HasVcsBrowser

`func (o *DebSourcePackageResponse) HasVcsBrowser() bool`

HasVcsBrowser returns a boolean if a field has been set.

### GetVcsArch

`func (o *DebSourcePackageResponse) GetVcsArch() string`

GetVcsArch returns the VcsArch field if non-nil, zero value otherwise.

### GetVcsArchOk

`func (o *DebSourcePackageResponse) GetVcsArchOk() (*string, bool)`

GetVcsArchOk returns a tuple with the VcsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsArch

`func (o *DebSourcePackageResponse) SetVcsArch(v string)`

SetVcsArch sets VcsArch field to given value.

### HasVcsArch

`func (o *DebSourcePackageResponse) HasVcsArch() bool`

HasVcsArch returns a boolean if a field has been set.

### GetVcsBzr

`func (o *DebSourcePackageResponse) GetVcsBzr() string`

GetVcsBzr returns the VcsBzr field if non-nil, zero value otherwise.

### GetVcsBzrOk

`func (o *DebSourcePackageResponse) GetVcsBzrOk() (*string, bool)`

GetVcsBzrOk returns a tuple with the VcsBzr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsBzr

`func (o *DebSourcePackageResponse) SetVcsBzr(v string)`

SetVcsBzr sets VcsBzr field to given value.

### HasVcsBzr

`func (o *DebSourcePackageResponse) HasVcsBzr() bool`

HasVcsBzr returns a boolean if a field has been set.

### GetVcsCvs

`func (o *DebSourcePackageResponse) GetVcsCvs() string`

GetVcsCvs returns the VcsCvs field if non-nil, zero value otherwise.

### GetVcsCvsOk

`func (o *DebSourcePackageResponse) GetVcsCvsOk() (*string, bool)`

GetVcsCvsOk returns a tuple with the VcsCvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsCvs

`func (o *DebSourcePackageResponse) SetVcsCvs(v string)`

SetVcsCvs sets VcsCvs field to given value.

### HasVcsCvs

`func (o *DebSourcePackageResponse) HasVcsCvs() bool`

HasVcsCvs returns a boolean if a field has been set.

### GetVcsDarcs

`func (o *DebSourcePackageResponse) GetVcsDarcs() string`

GetVcsDarcs returns the VcsDarcs field if non-nil, zero value otherwise.

### GetVcsDarcsOk

`func (o *DebSourcePackageResponse) GetVcsDarcsOk() (*string, bool)`

GetVcsDarcsOk returns a tuple with the VcsDarcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsDarcs

`func (o *DebSourcePackageResponse) SetVcsDarcs(v string)`

SetVcsDarcs sets VcsDarcs field to given value.

### HasVcsDarcs

`func (o *DebSourcePackageResponse) HasVcsDarcs() bool`

HasVcsDarcs returns a boolean if a field has been set.

### GetVcsGit

`func (o *DebSourcePackageResponse) GetVcsGit() string`

GetVcsGit returns the VcsGit field if non-nil, zero value otherwise.

### GetVcsGitOk

`func (o *DebSourcePackageResponse) GetVcsGitOk() (*string, bool)`

GetVcsGitOk returns a tuple with the VcsGit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsGit

`func (o *DebSourcePackageResponse) SetVcsGit(v string)`

SetVcsGit sets VcsGit field to given value.

### HasVcsGit

`func (o *DebSourcePackageResponse) HasVcsGit() bool`

HasVcsGit returns a boolean if a field has been set.

### GetVcsHg

`func (o *DebSourcePackageResponse) GetVcsHg() string`

GetVcsHg returns the VcsHg field if non-nil, zero value otherwise.

### GetVcsHgOk

`func (o *DebSourcePackageResponse) GetVcsHgOk() (*string, bool)`

GetVcsHgOk returns a tuple with the VcsHg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsHg

`func (o *DebSourcePackageResponse) SetVcsHg(v string)`

SetVcsHg sets VcsHg field to given value.

### HasVcsHg

`func (o *DebSourcePackageResponse) HasVcsHg() bool`

HasVcsHg returns a boolean if a field has been set.

### GetVcsMtn

`func (o *DebSourcePackageResponse) GetVcsMtn() string`

GetVcsMtn returns the VcsMtn field if non-nil, zero value otherwise.

### GetVcsMtnOk

`func (o *DebSourcePackageResponse) GetVcsMtnOk() (*string, bool)`

GetVcsMtnOk returns a tuple with the VcsMtn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsMtn

`func (o *DebSourcePackageResponse) SetVcsMtn(v string)`

SetVcsMtn sets VcsMtn field to given value.

### HasVcsMtn

`func (o *DebSourcePackageResponse) HasVcsMtn() bool`

HasVcsMtn returns a boolean if a field has been set.

### GetVcsSnv

`func (o *DebSourcePackageResponse) GetVcsSnv() string`

GetVcsSnv returns the VcsSnv field if non-nil, zero value otherwise.

### GetVcsSnvOk

`func (o *DebSourcePackageResponse) GetVcsSnvOk() (*string, bool)`

GetVcsSnvOk returns a tuple with the VcsSnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsSnv

`func (o *DebSourcePackageResponse) SetVcsSnv(v string)`

SetVcsSnv sets VcsSnv field to given value.

### HasVcsSnv

`func (o *DebSourcePackageResponse) HasVcsSnv() bool`

HasVcsSnv returns a boolean if a field has been set.

### GetTestsuite

`func (o *DebSourcePackageResponse) GetTestsuite() string`

GetTestsuite returns the Testsuite field if non-nil, zero value otherwise.

### GetTestsuiteOk

`func (o *DebSourcePackageResponse) GetTestsuiteOk() (*string, bool)`

GetTestsuiteOk returns a tuple with the Testsuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestsuite

`func (o *DebSourcePackageResponse) SetTestsuite(v string)`

SetTestsuite sets Testsuite field to given value.

### HasTestsuite

`func (o *DebSourcePackageResponse) HasTestsuite() bool`

HasTestsuite returns a boolean if a field has been set.

### GetDgit

`func (o *DebSourcePackageResponse) GetDgit() string`

GetDgit returns the Dgit field if non-nil, zero value otherwise.

### GetDgitOk

`func (o *DebSourcePackageResponse) GetDgitOk() (*string, bool)`

GetDgitOk returns a tuple with the Dgit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDgit

`func (o *DebSourcePackageResponse) SetDgit(v string)`

SetDgit sets Dgit field to given value.

### HasDgit

`func (o *DebSourcePackageResponse) HasDgit() bool`

HasDgit returns a boolean if a field has been set.

### GetStandardsVersion

`func (o *DebSourcePackageResponse) GetStandardsVersion() string`

GetStandardsVersion returns the StandardsVersion field if non-nil, zero value otherwise.

### GetStandardsVersionOk

`func (o *DebSourcePackageResponse) GetStandardsVersionOk() (*string, bool)`

GetStandardsVersionOk returns a tuple with the StandardsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardsVersion

`func (o *DebSourcePackageResponse) SetStandardsVersion(v string)`

SetStandardsVersion sets StandardsVersion field to given value.

### HasStandardsVersion

`func (o *DebSourcePackageResponse) HasStandardsVersion() bool`

HasStandardsVersion returns a boolean if a field has been set.

### GetBuildDepends

`func (o *DebSourcePackageResponse) GetBuildDepends() string`

GetBuildDepends returns the BuildDepends field if non-nil, zero value otherwise.

### GetBuildDependsOk

`func (o *DebSourcePackageResponse) GetBuildDependsOk() (*string, bool)`

GetBuildDependsOk returns a tuple with the BuildDepends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDepends

`func (o *DebSourcePackageResponse) SetBuildDepends(v string)`

SetBuildDepends sets BuildDepends field to given value.

### HasBuildDepends

`func (o *DebSourcePackageResponse) HasBuildDepends() bool`

HasBuildDepends returns a boolean if a field has been set.

### GetBuildDependsIndep

`func (o *DebSourcePackageResponse) GetBuildDependsIndep() string`

GetBuildDependsIndep returns the BuildDependsIndep field if non-nil, zero value otherwise.

### GetBuildDependsIndepOk

`func (o *DebSourcePackageResponse) GetBuildDependsIndepOk() (*string, bool)`

GetBuildDependsIndepOk returns a tuple with the BuildDependsIndep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDependsIndep

`func (o *DebSourcePackageResponse) SetBuildDependsIndep(v string)`

SetBuildDependsIndep sets BuildDependsIndep field to given value.

### HasBuildDependsIndep

`func (o *DebSourcePackageResponse) HasBuildDependsIndep() bool`

HasBuildDependsIndep returns a boolean if a field has been set.

### GetBuildDependsArch

`func (o *DebSourcePackageResponse) GetBuildDependsArch() string`

GetBuildDependsArch returns the BuildDependsArch field if non-nil, zero value otherwise.

### GetBuildDependsArchOk

`func (o *DebSourcePackageResponse) GetBuildDependsArchOk() (*string, bool)`

GetBuildDependsArchOk returns a tuple with the BuildDependsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDependsArch

`func (o *DebSourcePackageResponse) SetBuildDependsArch(v string)`

SetBuildDependsArch sets BuildDependsArch field to given value.

### HasBuildDependsArch

`func (o *DebSourcePackageResponse) HasBuildDependsArch() bool`

HasBuildDependsArch returns a boolean if a field has been set.

### GetBuildConflicts

`func (o *DebSourcePackageResponse) GetBuildConflicts() string`

GetBuildConflicts returns the BuildConflicts field if non-nil, zero value otherwise.

### GetBuildConflictsOk

`func (o *DebSourcePackageResponse) GetBuildConflictsOk() (*string, bool)`

GetBuildConflictsOk returns a tuple with the BuildConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildConflicts

`func (o *DebSourcePackageResponse) SetBuildConflicts(v string)`

SetBuildConflicts sets BuildConflicts field to given value.

### HasBuildConflicts

`func (o *DebSourcePackageResponse) HasBuildConflicts() bool`

HasBuildConflicts returns a boolean if a field has been set.

### GetBuildConflictsIndep

`func (o *DebSourcePackageResponse) GetBuildConflictsIndep() string`

GetBuildConflictsIndep returns the BuildConflictsIndep field if non-nil, zero value otherwise.

### GetBuildConflictsIndepOk

`func (o *DebSourcePackageResponse) GetBuildConflictsIndepOk() (*string, bool)`

GetBuildConflictsIndepOk returns a tuple with the BuildConflictsIndep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildConflictsIndep

`func (o *DebSourcePackageResponse) SetBuildConflictsIndep(v string)`

SetBuildConflictsIndep sets BuildConflictsIndep field to given value.

### HasBuildConflictsIndep

`func (o *DebSourcePackageResponse) HasBuildConflictsIndep() bool`

HasBuildConflictsIndep returns a boolean if a field has been set.

### GetBuildConflictsArch

`func (o *DebSourcePackageResponse) GetBuildConflictsArch() string`

GetBuildConflictsArch returns the BuildConflictsArch field if non-nil, zero value otherwise.

### GetBuildConflictsArchOk

`func (o *DebSourcePackageResponse) GetBuildConflictsArchOk() (*string, bool)`

GetBuildConflictsArchOk returns a tuple with the BuildConflictsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildConflictsArch

`func (o *DebSourcePackageResponse) SetBuildConflictsArch(v string)`

SetBuildConflictsArch sets BuildConflictsArch field to given value.

### HasBuildConflictsArch

`func (o *DebSourcePackageResponse) HasBuildConflictsArch() bool`

HasBuildConflictsArch returns a boolean if a field has been set.

### GetPackageList

`func (o *DebSourcePackageResponse) GetPackageList() string`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *DebSourcePackageResponse) GetPackageListOk() (*string, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *DebSourcePackageResponse) SetPackageList(v string)`

SetPackageList sets PackageList field to given value.

### HasPackageList

`func (o *DebSourcePackageResponse) HasPackageList() bool`

HasPackageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


