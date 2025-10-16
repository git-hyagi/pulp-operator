# AnsibleAnsibleCollectionDeprecated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Name** | **string** | The name of the Collection. | 
**Namespace** | **string** | The namespace of the Collection. | 

## Methods

### NewAnsibleAnsibleCollectionDeprecated

`func NewAnsibleAnsibleCollectionDeprecated(name string, namespace string, ) *AnsibleAnsibleCollectionDeprecated`

NewAnsibleAnsibleCollectionDeprecated instantiates a new AnsibleAnsibleCollectionDeprecated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleAnsibleCollectionDeprecatedWithDefaults

`func NewAnsibleAnsibleCollectionDeprecatedWithDefaults() *AnsibleAnsibleCollectionDeprecated`

NewAnsibleAnsibleCollectionDeprecatedWithDefaults instantiates a new AnsibleAnsibleCollectionDeprecated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *AnsibleAnsibleCollectionDeprecated) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AnsibleAnsibleCollectionDeprecated) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AnsibleAnsibleCollectionDeprecated) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AnsibleAnsibleCollectionDeprecated) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPulpLabels

`func (o *AnsibleAnsibleCollectionDeprecated) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *AnsibleAnsibleCollectionDeprecated) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *AnsibleAnsibleCollectionDeprecated) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *AnsibleAnsibleCollectionDeprecated) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *AnsibleAnsibleCollectionDeprecated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleAnsibleCollectionDeprecated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleAnsibleCollectionDeprecated) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AnsibleAnsibleCollectionDeprecated) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AnsibleAnsibleCollectionDeprecated) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AnsibleAnsibleCollectionDeprecated) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


