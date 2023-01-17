# IntegrationStatusRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**ResponseBody** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewIntegrationStatusRep

`func NewIntegrationStatusRep() *IntegrationStatusRep`

NewIntegrationStatusRep instantiates a new IntegrationStatusRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationStatusRepWithDefaults

`func NewIntegrationStatusRepWithDefaults() *IntegrationStatusRep`

NewIntegrationStatusRepWithDefaults instantiates a new IntegrationStatusRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *IntegrationStatusRep) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *IntegrationStatusRep) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *IntegrationStatusRep) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *IntegrationStatusRep) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetResponseBody

`func (o *IntegrationStatusRep) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *IntegrationStatusRep) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *IntegrationStatusRep) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *IntegrationStatusRep) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetTimestamp

`func (o *IntegrationStatusRep) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IntegrationStatusRep) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IntegrationStatusRep) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *IntegrationStatusRep) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


