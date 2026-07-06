# CreateIpAllowlistEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateIpAllowlistEntryRequest

`func NewCreateIpAllowlistEntryRequest(ipAddress string, ) *CreateIpAllowlistEntryRequest`

NewCreateIpAllowlistEntryRequest instantiates a new CreateIpAllowlistEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIpAllowlistEntryRequestWithDefaults

`func NewCreateIpAllowlistEntryRequestWithDefaults() *CreateIpAllowlistEntryRequest`

NewCreateIpAllowlistEntryRequestWithDefaults instantiates a new CreateIpAllowlistEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *CreateIpAllowlistEntryRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateIpAllowlistEntryRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateIpAllowlistEntryRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetDescription

`func (o *CreateIpAllowlistEntryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIpAllowlistEntryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIpAllowlistEntryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIpAllowlistEntryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


