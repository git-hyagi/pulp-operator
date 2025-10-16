# DebSourcePackageReleaseComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**SourcePackage** | **string** | Source package that is contained in release_component. | 
**ReleaseComponent** | **string** | ReleaseComponent this source package is contained in. | 

## Methods

### NewDebSourcePackageReleaseComponent

`func NewDebSourcePackageReleaseComponent(sourcePackage string, releaseComponent string, ) *DebSourcePackageReleaseComponent`

NewDebSourcePackageReleaseComponent instantiates a new DebSourcePackageReleaseComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebSourcePackageReleaseComponentWithDefaults

`func NewDebSourcePackageReleaseComponentWithDefaults() *DebSourcePackageReleaseComponent`

NewDebSourcePackageReleaseComponentWithDefaults instantiates a new DebSourcePackageReleaseComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebSourcePackageReleaseComponent) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebSourcePackageReleaseComponent) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebSourcePackageReleaseComponent) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebSourcePackageReleaseComponent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPulpLabels

`func (o *DebSourcePackageReleaseComponent) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *DebSourcePackageReleaseComponent) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *DebSourcePackageReleaseComponent) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *DebSourcePackageReleaseComponent) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetSourcePackage

`func (o *DebSourcePackageReleaseComponent) GetSourcePackage() string`

GetSourcePackage returns the SourcePackage field if non-nil, zero value otherwise.

### GetSourcePackageOk

`func (o *DebSourcePackageReleaseComponent) GetSourcePackageOk() (*string, bool)`

GetSourcePackageOk returns a tuple with the SourcePackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePackage

`func (o *DebSourcePackageReleaseComponent) SetSourcePackage(v string)`

SetSourcePackage sets SourcePackage field to given value.


### GetReleaseComponent

`func (o *DebSourcePackageReleaseComponent) GetReleaseComponent() string`

GetReleaseComponent returns the ReleaseComponent field if non-nil, zero value otherwise.

### GetReleaseComponentOk

`func (o *DebSourcePackageReleaseComponent) GetReleaseComponentOk() (*string, bool)`

GetReleaseComponentOk returns a tuple with the ReleaseComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseComponent

`func (o *DebSourcePackageReleaseComponent) SetReleaseComponent(v string)`

SetReleaseComponent sets ReleaseComponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


