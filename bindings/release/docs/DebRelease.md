# DebRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Codename** | **string** |  | 
**Suite** | **string** |  | 
**Distribution** | **string** |  | 
**Version** | Pointer to **NullableString** |  | [optional] 
**Origin** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Architectures** | Pointer to **[]string** |  | [optional] 
**Components** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDebRelease

`func NewDebRelease(codename string, suite string, distribution string, ) *DebRelease`

NewDebRelease instantiates a new DebRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebReleaseWithDefaults

`func NewDebReleaseWithDefaults() *DebRelease`

NewDebReleaseWithDefaults instantiates a new DebRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebRelease) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebRelease) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebRelease) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebRelease) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPulpLabels

`func (o *DebRelease) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *DebRelease) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *DebRelease) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *DebRelease) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetCodename

`func (o *DebRelease) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *DebRelease) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *DebRelease) SetCodename(v string)`

SetCodename sets Codename field to given value.


### GetSuite

`func (o *DebRelease) GetSuite() string`

GetSuite returns the Suite field if non-nil, zero value otherwise.

### GetSuiteOk

`func (o *DebRelease) GetSuiteOk() (*string, bool)`

GetSuiteOk returns a tuple with the Suite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuite

`func (o *DebRelease) SetSuite(v string)`

SetSuite sets Suite field to given value.


### GetDistribution

`func (o *DebRelease) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebRelease) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebRelease) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetVersion

`func (o *DebRelease) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DebRelease) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DebRelease) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DebRelease) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DebRelease) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DebRelease) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetOrigin

`func (o *DebRelease) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DebRelease) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DebRelease) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *DebRelease) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *DebRelease) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *DebRelease) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetLabel

`func (o *DebRelease) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DebRelease) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DebRelease) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DebRelease) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *DebRelease) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *DebRelease) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDescription

`func (o *DebRelease) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DebRelease) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DebRelease) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DebRelease) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DebRelease) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DebRelease) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetArchitectures

`func (o *DebRelease) GetArchitectures() []string`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *DebRelease) GetArchitecturesOk() (*[]string, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *DebRelease) SetArchitectures(v []string)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *DebRelease) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetComponents

`func (o *DebRelease) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DebRelease) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DebRelease) SetComponents(v []string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DebRelease) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


