# IpAllowlistResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionAllowlistEnabled** | **bool** |  | 
**ApiTokenAllowlistEnabled** | **bool** |  | 
**Entries** | [**[]IpAllowlistEntryResponse**](IpAllowlistEntryResponse.md) |  | 

## Methods

### NewIpAllowlistResponse

`func NewIpAllowlistResponse(sessionAllowlistEnabled bool, apiTokenAllowlistEnabled bool, entries []IpAllowlistEntryResponse, ) *IpAllowlistResponse`

NewIpAllowlistResponse instantiates a new IpAllowlistResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAllowlistResponseWithDefaults

`func NewIpAllowlistResponseWithDefaults() *IpAllowlistResponse`

NewIpAllowlistResponseWithDefaults instantiates a new IpAllowlistResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionAllowlistEnabled

`func (o *IpAllowlistResponse) GetSessionAllowlistEnabled() bool`

GetSessionAllowlistEnabled returns the SessionAllowlistEnabled field if non-nil, zero value otherwise.

### GetSessionAllowlistEnabledOk

`func (o *IpAllowlistResponse) GetSessionAllowlistEnabledOk() (*bool, bool)`

GetSessionAllowlistEnabledOk returns a tuple with the SessionAllowlistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAllowlistEnabled

`func (o *IpAllowlistResponse) SetSessionAllowlistEnabled(v bool)`

SetSessionAllowlistEnabled sets SessionAllowlistEnabled field to given value.


### GetApiTokenAllowlistEnabled

`func (o *IpAllowlistResponse) GetApiTokenAllowlistEnabled() bool`

GetApiTokenAllowlistEnabled returns the ApiTokenAllowlistEnabled field if non-nil, zero value otherwise.

### GetApiTokenAllowlistEnabledOk

`func (o *IpAllowlistResponse) GetApiTokenAllowlistEnabledOk() (*bool, bool)`

GetApiTokenAllowlistEnabledOk returns a tuple with the ApiTokenAllowlistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenAllowlistEnabled

`func (o *IpAllowlistResponse) SetApiTokenAllowlistEnabled(v bool)`

SetApiTokenAllowlistEnabled sets ApiTokenAllowlistEnabled field to given value.


### GetEntries

`func (o *IpAllowlistResponse) GetEntries() []IpAllowlistEntryResponse`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *IpAllowlistResponse) GetEntriesOk() (*[]IpAllowlistEntryResponse, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *IpAllowlistResponse) SetEntries(v []IpAllowlistEntryResponse)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


