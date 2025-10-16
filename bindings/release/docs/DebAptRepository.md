# DebAptRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**PublishUpstreamReleaseFields** | Pointer to **bool** | Previously, pulp_deb only synced the Release file fields codename and suite, now version, origin, label, and description are also synced. Setting this setting to False will make Pulp revert to the old behaviour of using it&#39;s own internal values for the new fields during publish. This is primarily intended to avoid a sudden change in behaviour for existing Pulp repositories, since many Release file field changes need to be accepted by hosts consuming the published repository. The default for new repositories is True. | [optional] 
**SigningService** | Pointer to **NullableString** | A reference to an associated signing service. Used if AptPublication.signing_service is not set | [optional] 
**SigningServiceReleaseOverrides** | Pointer to **map[string]string** | A dictionary of Release distributions and the Signing Service URLs they should use.Example: {\&quot;bionic\&quot;: \&quot;/pulp/api/v3/signing-services/433a1f70-c589-4413-a803-c50b842ea9b5/\&quot;} | [optional] 

## Methods

### NewDebAptRepository

`func NewDebAptRepository(name string, ) *DebAptRepository`

NewDebAptRepository instantiates a new DebAptRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebAptRepositoryWithDefaults

`func NewDebAptRepositoryWithDefaults() *DebAptRepository`

NewDebAptRepositoryWithDefaults instantiates a new DebAptRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *DebAptRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *DebAptRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *DebAptRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *DebAptRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *DebAptRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DebAptRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DebAptRepository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DebAptRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DebAptRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DebAptRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DebAptRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DebAptRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DebAptRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *DebAptRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *DebAptRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *DebAptRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *DebAptRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *DebAptRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *DebAptRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *DebAptRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *DebAptRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *DebAptRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *DebAptRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *DebAptRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *DebAptRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetPublishUpstreamReleaseFields

`func (o *DebAptRepository) GetPublishUpstreamReleaseFields() bool`

GetPublishUpstreamReleaseFields returns the PublishUpstreamReleaseFields field if non-nil, zero value otherwise.

### GetPublishUpstreamReleaseFieldsOk

`func (o *DebAptRepository) GetPublishUpstreamReleaseFieldsOk() (*bool, bool)`

GetPublishUpstreamReleaseFieldsOk returns a tuple with the PublishUpstreamReleaseFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishUpstreamReleaseFields

`func (o *DebAptRepository) SetPublishUpstreamReleaseFields(v bool)`

SetPublishUpstreamReleaseFields sets PublishUpstreamReleaseFields field to given value.

### HasPublishUpstreamReleaseFields

`func (o *DebAptRepository) HasPublishUpstreamReleaseFields() bool`

HasPublishUpstreamReleaseFields returns a boolean if a field has been set.

### GetSigningService

`func (o *DebAptRepository) GetSigningService() string`

GetSigningService returns the SigningService field if non-nil, zero value otherwise.

### GetSigningServiceOk

`func (o *DebAptRepository) GetSigningServiceOk() (*string, bool)`

GetSigningServiceOk returns a tuple with the SigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningService

`func (o *DebAptRepository) SetSigningService(v string)`

SetSigningService sets SigningService field to given value.

### HasSigningService

`func (o *DebAptRepository) HasSigningService() bool`

HasSigningService returns a boolean if a field has been set.

### SetSigningServiceNil

`func (o *DebAptRepository) SetSigningServiceNil(b bool)`

 SetSigningServiceNil sets the value for SigningService to be an explicit nil

### UnsetSigningService
`func (o *DebAptRepository) UnsetSigningService()`

UnsetSigningService ensures that no value is present for SigningService, not even an explicit nil
### GetSigningServiceReleaseOverrides

`func (o *DebAptRepository) GetSigningServiceReleaseOverrides() map[string]string`

GetSigningServiceReleaseOverrides returns the SigningServiceReleaseOverrides field if non-nil, zero value otherwise.

### GetSigningServiceReleaseOverridesOk

`func (o *DebAptRepository) GetSigningServiceReleaseOverridesOk() (*map[string]string, bool)`

GetSigningServiceReleaseOverridesOk returns a tuple with the SigningServiceReleaseOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningServiceReleaseOverrides

`func (o *DebAptRepository) SetSigningServiceReleaseOverrides(v map[string]string)`

SetSigningServiceReleaseOverrides sets SigningServiceReleaseOverrides field to given value.

### HasSigningServiceReleaseOverrides

`func (o *DebAptRepository) HasSigningServiceReleaseOverrides() bool`

HasSigningServiceReleaseOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


