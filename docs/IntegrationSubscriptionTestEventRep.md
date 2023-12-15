# IntegrationSubscriptionTestEventRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**ResponseBody** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegrationSubscriptionTestEventRep

`func NewIntegrationSubscriptionTestEventRep() *IntegrationSubscriptionTestEventRep`

NewIntegrationSubscriptionTestEventRep instantiates a new IntegrationSubscriptionTestEventRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSubscriptionTestEventRepWithDefaults

`func NewIntegrationSubscriptionTestEventRepWithDefaults() *IntegrationSubscriptionTestEventRep`

NewIntegrationSubscriptionTestEventRepWithDefaults instantiates a new IntegrationSubscriptionTestEventRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *IntegrationSubscriptionTestEventRep) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *IntegrationSubscriptionTestEventRep) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *IntegrationSubscriptionTestEventRep) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *IntegrationSubscriptionTestEventRep) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *IntegrationSubscriptionTestEventRep) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IntegrationSubscriptionTestEventRep) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IntegrationSubscriptionTestEventRep) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *IntegrationSubscriptionTestEventRep) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTimestamp

`func (o *IntegrationSubscriptionTestEventRep) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IntegrationSubscriptionTestEventRep) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IntegrationSubscriptionTestEventRep) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *IntegrationSubscriptionTestEventRep) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetResponseBody

`func (o *IntegrationSubscriptionTestEventRep) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *IntegrationSubscriptionTestEventRep) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *IntegrationSubscriptionTestEventRep) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *IntegrationSubscriptionTestEventRep) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


