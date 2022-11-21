# RecentTriggerBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int64** |  | [optional] 
**JsonBody** | Pointer to **map[string]interface{}** | The marshalled JSON request body for the incoming trigger webhook. If this is empty or contains invalid JSON, the timestamp is recorded but this field will be empty. | [optional] 

## Methods

### NewRecentTriggerBody

`func NewRecentTriggerBody() *RecentTriggerBody`

NewRecentTriggerBody instantiates a new RecentTriggerBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecentTriggerBodyWithDefaults

`func NewRecentTriggerBodyWithDefaults() *RecentTriggerBody`

NewRecentTriggerBodyWithDefaults instantiates a new RecentTriggerBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *RecentTriggerBody) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RecentTriggerBody) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RecentTriggerBody) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RecentTriggerBody) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetJsonBody

`func (o *RecentTriggerBody) GetJsonBody() map[string]interface{}`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *RecentTriggerBody) GetJsonBodyOk() (*map[string]interface{}, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *RecentTriggerBody) SetJsonBody(v map[string]interface{})`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *RecentTriggerBody) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


