# ExpiringUserTargetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Version** | **int32** |  | 
**ExpirationDate** | **int64** |  | 
**UserKey** | **string** | A unique key used to represent the user | 
**TargetType** | Pointer to **string** | A segment&#39;s target type. Included when expiring user targets are updated on a user segment. | [optional] 
**VariationId** | Pointer to **string** | A unique key used to represent the flag variation. Included when expiring user targets are updated on a feature flag. | [optional] 
**ResourceId** | [**ResourceIDResponse**](ResourceIDResponse.md) |  | 

## Methods

### NewExpiringUserTargetItem

`func NewExpiringUserTargetItem(id string, version int32, expirationDate int64, userKey string, resourceId ResourceIDResponse, ) *ExpiringUserTargetItem`

NewExpiringUserTargetItem instantiates a new ExpiringUserTargetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringUserTargetItemWithDefaults

`func NewExpiringUserTargetItemWithDefaults() *ExpiringUserTargetItem`

NewExpiringUserTargetItemWithDefaults instantiates a new ExpiringUserTargetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpiringUserTargetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpiringUserTargetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpiringUserTargetItem) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *ExpiringUserTargetItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExpiringUserTargetItem) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExpiringUserTargetItem) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetExpirationDate

`func (o *ExpiringUserTargetItem) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ExpiringUserTargetItem) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ExpiringUserTargetItem) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.


### GetUserKey

`func (o *ExpiringUserTargetItem) GetUserKey() string`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *ExpiringUserTargetItem) GetUserKeyOk() (*string, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *ExpiringUserTargetItem) SetUserKey(v string)`

SetUserKey sets UserKey field to given value.


### GetTargetType

`func (o *ExpiringUserTargetItem) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ExpiringUserTargetItem) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ExpiringUserTargetItem) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ExpiringUserTargetItem) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetVariationId

`func (o *ExpiringUserTargetItem) GetVariationId() string`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *ExpiringUserTargetItem) GetVariationIdOk() (*string, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *ExpiringUserTargetItem) SetVariationId(v string)`

SetVariationId sets VariationId field to given value.

### HasVariationId

`func (o *ExpiringUserTargetItem) HasVariationId() bool`

HasVariationId returns a boolean if a field has been set.

### GetResourceId

`func (o *ExpiringUserTargetItem) GetResourceId() ResourceIDResponse`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ExpiringUserTargetItem) GetResourceIdOk() (*ResourceIDResponse, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ExpiringUserTargetItem) SetResourceId(v ResourceIDResponse)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


