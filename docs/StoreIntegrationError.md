# StoreIntegrationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewStoreIntegrationError

`func NewStoreIntegrationError() *StoreIntegrationError`

NewStoreIntegrationError instantiates a new StoreIntegrationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIntegrationErrorWithDefaults

`func NewStoreIntegrationErrorWithDefaults() *StoreIntegrationError`

NewStoreIntegrationErrorWithDefaults instantiates a new StoreIntegrationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *StoreIntegrationError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *StoreIntegrationError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *StoreIntegrationError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *StoreIntegrationError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetMessage

`func (o *StoreIntegrationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StoreIntegrationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StoreIntegrationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StoreIntegrationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *StoreIntegrationError) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StoreIntegrationError) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StoreIntegrationError) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StoreIntegrationError) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


