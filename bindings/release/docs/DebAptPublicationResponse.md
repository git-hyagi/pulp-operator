# DebAptPublicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**RepositoryVersion** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | A URI of the repository to be published. | [optional] 
**Simple** | Pointer to **bool** | Activate simple publishing mode (all packages in one release component). | [optional] [default to false]
**Structured** | Pointer to **bool** | Activate structured publishing mode. | [optional] [default to true]
**Checkpoint** | Pointer to **bool** |  | [optional] 
**SigningService** | Pointer to **string** | Sign Release files with this signing key | [optional] 
**PublishUpstreamReleaseFields** | Pointer to **bool** |  | [optional] 

## Methods

### NewDebAptPublicationResponse

`func NewDebAptPublicationResponse() *DebAptPublicationResponse`

NewDebAptPublicationResponse instantiates a new DebAptPublicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebAptPublicationResponseWithDefaults

`func NewDebAptPublicationResponseWithDefaults() *DebAptPublicationResponse`

NewDebAptPublicationResponseWithDefaults instantiates a new DebAptPublicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebAptPublicationResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebAptPublicationResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebAptPublicationResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebAptPublicationResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *DebAptPublicationResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *DebAptPublicationResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *DebAptPublicationResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *DebAptPublicationResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebAptPublicationResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebAptPublicationResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebAptPublicationResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebAptPublicationResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *DebAptPublicationResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *DebAptPublicationResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *DebAptPublicationResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *DebAptPublicationResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetRepositoryVersion

`func (o *DebAptPublicationResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *DebAptPublicationResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *DebAptPublicationResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *DebAptPublicationResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetRepository

`func (o *DebAptPublicationResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebAptPublicationResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebAptPublicationResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebAptPublicationResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSimple

`func (o *DebAptPublicationResponse) GetSimple() bool`

GetSimple returns the Simple field if non-nil, zero value otherwise.

### GetSimpleOk

`func (o *DebAptPublicationResponse) GetSimpleOk() (*bool, bool)`

GetSimpleOk returns a tuple with the Simple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimple

`func (o *DebAptPublicationResponse) SetSimple(v bool)`

SetSimple sets Simple field to given value.

### HasSimple

`func (o *DebAptPublicationResponse) HasSimple() bool`

HasSimple returns a boolean if a field has been set.

### GetStructured

`func (o *DebAptPublicationResponse) GetStructured() bool`

GetStructured returns the Structured field if non-nil, zero value otherwise.

### GetStructuredOk

`func (o *DebAptPublicationResponse) GetStructuredOk() (*bool, bool)`

GetStructuredOk returns a tuple with the Structured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructured

`func (o *DebAptPublicationResponse) SetStructured(v bool)`

SetStructured sets Structured field to given value.

### HasStructured

`func (o *DebAptPublicationResponse) HasStructured() bool`

HasStructured returns a boolean if a field has been set.

### GetCheckpoint

`func (o *DebAptPublicationResponse) GetCheckpoint() bool`

GetCheckpoint returns the Checkpoint field if non-nil, zero value otherwise.

### GetCheckpointOk

`func (o *DebAptPublicationResponse) GetCheckpointOk() (*bool, bool)`

GetCheckpointOk returns a tuple with the Checkpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoint

`func (o *DebAptPublicationResponse) SetCheckpoint(v bool)`

SetCheckpoint sets Checkpoint field to given value.

### HasCheckpoint

`func (o *DebAptPublicationResponse) HasCheckpoint() bool`

HasCheckpoint returns a boolean if a field has been set.

### GetSigningService

`func (o *DebAptPublicationResponse) GetSigningService() string`

GetSigningService returns the SigningService field if non-nil, zero value otherwise.

### GetSigningServiceOk

`func (o *DebAptPublicationResponse) GetSigningServiceOk() (*string, bool)`

GetSigningServiceOk returns a tuple with the SigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningService

`func (o *DebAptPublicationResponse) SetSigningService(v string)`

SetSigningService sets SigningService field to given value.

### HasSigningService

`func (o *DebAptPublicationResponse) HasSigningService() bool`

HasSigningService returns a boolean if a field has been set.

### GetPublishUpstreamReleaseFields

`func (o *DebAptPublicationResponse) GetPublishUpstreamReleaseFields() bool`

GetPublishUpstreamReleaseFields returns the PublishUpstreamReleaseFields field if non-nil, zero value otherwise.

### GetPublishUpstreamReleaseFieldsOk

`func (o *DebAptPublicationResponse) GetPublishUpstreamReleaseFieldsOk() (*bool, bool)`

GetPublishUpstreamReleaseFieldsOk returns a tuple with the PublishUpstreamReleaseFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishUpstreamReleaseFields

`func (o *DebAptPublicationResponse) SetPublishUpstreamReleaseFields(v bool)`

SetPublishUpstreamReleaseFields sets PublishUpstreamReleaseFields field to given value.

### HasPublishUpstreamReleaseFields

`func (o *DebAptPublicationResponse) HasPublishUpstreamReleaseFields() bool`

HasPublishUpstreamReleaseFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


