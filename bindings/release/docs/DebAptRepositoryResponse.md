# DebAptRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**VersionsHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**LatestVersionHref** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**PublishUpstreamReleaseFields** | Pointer to **bool** | Previously, pulp_deb only synced the Release file fields codename and suite, now version, origin, label, and description are also synced. Setting this setting to False will make Pulp revert to the old behaviour of using it&#39;s own internal values for the new fields during publish. This is primarily intended to avoid a sudden change in behaviour for existing Pulp repositories, since many Release file field changes need to be accepted by hosts consuming the published repository. The default for new repositories is True. | [optional] 
**SigningService** | Pointer to **NullableString** | A reference to an associated signing service. Used if AptPublication.signing_service is not set | [optional] 
**SigningServiceReleaseOverrides** | Pointer to **map[string]string** | A dictionary of Release distributions and the Signing Service URLs they should use.Example: {\&quot;bionic\&quot;: \&quot;/pulp/api/v3/signing-services/433a1f70-c589-4413-a803-c50b842ea9b5/\&quot;} | [optional] 

## Methods

### NewDebAptRepositoryResponse

`func NewDebAptRepositoryResponse(name string, ) *DebAptRepositoryResponse`

NewDebAptRepositoryResponse instantiates a new DebAptRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebAptRepositoryResponseWithDefaults

`func NewDebAptRepositoryResponseWithDefaults() *DebAptRepositoryResponse`

NewDebAptRepositoryResponseWithDefaults instantiates a new DebAptRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebAptRepositoryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebAptRepositoryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebAptRepositoryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebAptRepositoryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *DebAptRepositoryResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *DebAptRepositoryResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *DebAptRepositoryResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *DebAptRepositoryResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebAptRepositoryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebAptRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebAptRepositoryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebAptRepositoryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *DebAptRepositoryResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *DebAptRepositoryResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *DebAptRepositoryResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *DebAptRepositoryResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetVersionsHref

`func (o *DebAptRepositoryResponse) GetVersionsHref() string`

GetVersionsHref returns the VersionsHref field if non-nil, zero value otherwise.

### GetVersionsHrefOk

`func (o *DebAptRepositoryResponse) GetVersionsHrefOk() (*string, bool)`

GetVersionsHrefOk returns a tuple with the VersionsHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsHref

`func (o *DebAptRepositoryResponse) SetVersionsHref(v string)`

SetVersionsHref sets VersionsHref field to given value.

### HasVersionsHref

`func (o *DebAptRepositoryResponse) HasVersionsHref() bool`

HasVersionsHref returns a boolean if a field has been set.

### GetPulpLabels

`func (o *DebAptRepositoryResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *DebAptRepositoryResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *DebAptRepositoryResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *DebAptRepositoryResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetLatestVersionHref

`func (o *DebAptRepositoryResponse) GetLatestVersionHref() string`

GetLatestVersionHref returns the LatestVersionHref field if non-nil, zero value otherwise.

### GetLatestVersionHrefOk

`func (o *DebAptRepositoryResponse) GetLatestVersionHrefOk() (*string, bool)`

GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionHref

`func (o *DebAptRepositoryResponse) SetLatestVersionHref(v string)`

SetLatestVersionHref sets LatestVersionHref field to given value.

### HasLatestVersionHref

`func (o *DebAptRepositoryResponse) HasLatestVersionHref() bool`

HasLatestVersionHref returns a boolean if a field has been set.

### GetName

`func (o *DebAptRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DebAptRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DebAptRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DebAptRepositoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DebAptRepositoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DebAptRepositoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DebAptRepositoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DebAptRepositoryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DebAptRepositoryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *DebAptRepositoryResponse) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *DebAptRepositoryResponse) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *DebAptRepositoryResponse) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *DebAptRepositoryResponse) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *DebAptRepositoryResponse) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *DebAptRepositoryResponse) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *DebAptRepositoryResponse) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *DebAptRepositoryResponse) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *DebAptRepositoryResponse) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *DebAptRepositoryResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *DebAptRepositoryResponse) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *DebAptRepositoryResponse) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetPublishUpstreamReleaseFields

`func (o *DebAptRepositoryResponse) GetPublishUpstreamReleaseFields() bool`

GetPublishUpstreamReleaseFields returns the PublishUpstreamReleaseFields field if non-nil, zero value otherwise.

### GetPublishUpstreamReleaseFieldsOk

`func (o *DebAptRepositoryResponse) GetPublishUpstreamReleaseFieldsOk() (*bool, bool)`

GetPublishUpstreamReleaseFieldsOk returns a tuple with the PublishUpstreamReleaseFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishUpstreamReleaseFields

`func (o *DebAptRepositoryResponse) SetPublishUpstreamReleaseFields(v bool)`

SetPublishUpstreamReleaseFields sets PublishUpstreamReleaseFields field to given value.

### HasPublishUpstreamReleaseFields

`func (o *DebAptRepositoryResponse) HasPublishUpstreamReleaseFields() bool`

HasPublishUpstreamReleaseFields returns a boolean if a field has been set.

### GetSigningService

`func (o *DebAptRepositoryResponse) GetSigningService() string`

GetSigningService returns the SigningService field if non-nil, zero value otherwise.

### GetSigningServiceOk

`func (o *DebAptRepositoryResponse) GetSigningServiceOk() (*string, bool)`

GetSigningServiceOk returns a tuple with the SigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningService

`func (o *DebAptRepositoryResponse) SetSigningService(v string)`

SetSigningService sets SigningService field to given value.

### HasSigningService

`func (o *DebAptRepositoryResponse) HasSigningService() bool`

HasSigningService returns a boolean if a field has been set.

### SetSigningServiceNil

`func (o *DebAptRepositoryResponse) SetSigningServiceNil(b bool)`

 SetSigningServiceNil sets the value for SigningService to be an explicit nil

### UnsetSigningService
`func (o *DebAptRepositoryResponse) UnsetSigningService()`

UnsetSigningService ensures that no value is present for SigningService, not even an explicit nil
### GetSigningServiceReleaseOverrides

`func (o *DebAptRepositoryResponse) GetSigningServiceReleaseOverrides() map[string]string`

GetSigningServiceReleaseOverrides returns the SigningServiceReleaseOverrides field if non-nil, zero value otherwise.

### GetSigningServiceReleaseOverridesOk

`func (o *DebAptRepositoryResponse) GetSigningServiceReleaseOverridesOk() (*map[string]string, bool)`

GetSigningServiceReleaseOverridesOk returns a tuple with the SigningServiceReleaseOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningServiceReleaseOverrides

`func (o *DebAptRepositoryResponse) SetSigningServiceReleaseOverrides(v map[string]string)`

SetSigningServiceReleaseOverrides sets SigningServiceReleaseOverrides field to given value.

### HasSigningServiceReleaseOverrides

`func (o *DebAptRepositoryResponse) HasSigningServiceReleaseOverrides() bool`

HasSigningServiceReleaseOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


