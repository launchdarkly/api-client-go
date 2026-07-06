# IpAllowlistEntryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the allowlist entry | 
**IpAddress** | **string** | IP address or CIDR block | 
**Description** | Pointer to **string** |  | [optional] 
**CreatedByMemberId** | Pointer to **string** |  | [optional] 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewIpAllowlistEntryResponse

`func NewIpAllowlistEntryResponse(id string, ipAddress string, createdAt int64, updatedAt int64, ) *IpAllowlistEntryResponse`

NewIpAllowlistEntryResponse instantiates a new IpAllowlistEntryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAllowlistEntryResponseWithDefaults

`func NewIpAllowlistEntryResponseWithDefaults() *IpAllowlistEntryResponse`

NewIpAllowlistEntryResponseWithDefaults instantiates a new IpAllowlistEntryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpAllowlistEntryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpAllowlistEntryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpAllowlistEntryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *IpAllowlistEntryResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IpAllowlistEntryResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IpAllowlistEntryResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetDescription

`func (o *IpAllowlistEntryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpAllowlistEntryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpAllowlistEntryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpAllowlistEntryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedByMemberId

`func (o *IpAllowlistEntryResponse) GetCreatedByMemberId() string`

GetCreatedByMemberId returns the CreatedByMemberId field if non-nil, zero value otherwise.

### GetCreatedByMemberIdOk

`func (o *IpAllowlistEntryResponse) GetCreatedByMemberIdOk() (*string, bool)`

GetCreatedByMemberIdOk returns a tuple with the CreatedByMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByMemberId

`func (o *IpAllowlistEntryResponse) SetCreatedByMemberId(v string)`

SetCreatedByMemberId sets CreatedByMemberId field to given value.

### HasCreatedByMemberId

`func (o *IpAllowlistEntryResponse) HasCreatedByMemberId() bool`

HasCreatedByMemberId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpAllowlistEntryResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpAllowlistEntryResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpAllowlistEntryResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IpAllowlistEntryResponse) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpAllowlistEntryResponse) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpAllowlistEntryResponse) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


