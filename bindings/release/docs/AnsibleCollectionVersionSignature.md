# AnsibleCollectionVersionSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the content unit. | [optional] 
**FileUrl** | Pointer to **string** | A url that Pulp can download and turn into the content unit. | [optional] 
**SignedCollection** | **string** | The content this signature is pointing to. | 

## Methods

### NewAnsibleCollectionVersionSignature

`func NewAnsibleCollectionVersionSignature(signedCollection string, ) *AnsibleCollectionVersionSignature`

NewAnsibleCollectionVersionSignature instantiates a new AnsibleCollectionVersionSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleCollectionVersionSignatureWithDefaults

`func NewAnsibleCollectionVersionSignatureWithDefaults() *AnsibleCollectionVersionSignature`

NewAnsibleCollectionVersionSignatureWithDefaults instantiates a new AnsibleCollectionVersionSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *AnsibleCollectionVersionSignature) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AnsibleCollectionVersionSignature) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AnsibleCollectionVersionSignature) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AnsibleCollectionVersionSignature) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPulpLabels

`func (o *AnsibleCollectionVersionSignature) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *AnsibleCollectionVersionSignature) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *AnsibleCollectionVersionSignature) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *AnsibleCollectionVersionSignature) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetFile

`func (o *AnsibleCollectionVersionSignature) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AnsibleCollectionVersionSignature) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AnsibleCollectionVersionSignature) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *AnsibleCollectionVersionSignature) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *AnsibleCollectionVersionSignature) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *AnsibleCollectionVersionSignature) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *AnsibleCollectionVersionSignature) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *AnsibleCollectionVersionSignature) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetFileUrl

`func (o *AnsibleCollectionVersionSignature) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *AnsibleCollectionVersionSignature) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *AnsibleCollectionVersionSignature) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *AnsibleCollectionVersionSignature) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetSignedCollection

`func (o *AnsibleCollectionVersionSignature) GetSignedCollection() string`

GetSignedCollection returns the SignedCollection field if non-nil, zero value otherwise.

### GetSignedCollectionOk

`func (o *AnsibleCollectionVersionSignature) GetSignedCollectionOk() (*string, bool)`

GetSignedCollectionOk returns a tuple with the SignedCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedCollection

`func (o *AnsibleCollectionVersionSignature) SetSignedCollection(v string)`

SetSignedCollection sets SignedCollection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


