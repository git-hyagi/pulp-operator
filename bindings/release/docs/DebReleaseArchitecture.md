# DebReleaseArchitecture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Architecture** | **string** | Name of the architecture. | 
**Distribution** | **string** | Name of the distribution. | 

## Methods

### NewDebReleaseArchitecture

`func NewDebReleaseArchitecture(architecture string, distribution string, ) *DebReleaseArchitecture`

NewDebReleaseArchitecture instantiates a new DebReleaseArchitecture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebReleaseArchitectureWithDefaults

`func NewDebReleaseArchitectureWithDefaults() *DebReleaseArchitecture`

NewDebReleaseArchitectureWithDefaults instantiates a new DebReleaseArchitecture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebReleaseArchitecture) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebReleaseArchitecture) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebReleaseArchitecture) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebReleaseArchitecture) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPulpLabels

`func (o *DebReleaseArchitecture) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *DebReleaseArchitecture) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *DebReleaseArchitecture) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *DebReleaseArchitecture) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetArchitecture

`func (o *DebReleaseArchitecture) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebReleaseArchitecture) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebReleaseArchitecture) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetDistribution

`func (o *DebReleaseArchitecture) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebReleaseArchitecture) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebReleaseArchitecture) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


