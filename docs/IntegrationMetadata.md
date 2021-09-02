# IntegrationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** |  | 
**ExternalStatus** | [**IntegrationStatus**](IntegrationStatus.md) |  | 
**ExternalUrl** | **string** |  | 
**LastChecked** | **int64** |  | 

## Methods

### NewIntegrationMetadata

`func NewIntegrationMetadata(externalId string, externalStatus IntegrationStatus, externalUrl string, lastChecked int64, ) *IntegrationMetadata`

NewIntegrationMetadata instantiates a new IntegrationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationMetadataWithDefaults

`func NewIntegrationMetadataWithDefaults() *IntegrationMetadata`

NewIntegrationMetadataWithDefaults instantiates a new IntegrationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *IntegrationMetadata) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IntegrationMetadata) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IntegrationMetadata) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetExternalStatus

`func (o *IntegrationMetadata) GetExternalStatus() IntegrationStatus`

GetExternalStatus returns the ExternalStatus field if non-nil, zero value otherwise.

### GetExternalStatusOk

`func (o *IntegrationMetadata) GetExternalStatusOk() (*IntegrationStatus, bool)`

GetExternalStatusOk returns a tuple with the ExternalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStatus

`func (o *IntegrationMetadata) SetExternalStatus(v IntegrationStatus)`

SetExternalStatus sets ExternalStatus field to given value.


### GetExternalUrl

`func (o *IntegrationMetadata) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *IntegrationMetadata) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *IntegrationMetadata) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.


### GetLastChecked

`func (o *IntegrationMetadata) GetLastChecked() int64`

GetLastChecked returns the LastChecked field if non-nil, zero value otherwise.

### GetLastCheckedOk

`func (o *IntegrationMetadata) GetLastCheckedOk() (*int64, bool)`

GetLastCheckedOk returns a tuple with the LastChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChecked

`func (o *IntegrationMetadata) SetLastChecked(v int64)`

SetLastChecked sets LastChecked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


