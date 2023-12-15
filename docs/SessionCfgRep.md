# SessionCfgRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMins** | Pointer to **int32** |  | [optional] 
**Refresh** | Pointer to **bool** |  | [optional] 

## Methods

### NewSessionCfgRep

`func NewSessionCfgRep() *SessionCfgRep`

NewSessionCfgRep instantiates a new SessionCfgRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionCfgRepWithDefaults

`func NewSessionCfgRepWithDefaults() *SessionCfgRep`

NewSessionCfgRepWithDefaults instantiates a new SessionCfgRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMins

`func (o *SessionCfgRep) GetDurationMins() int32`

GetDurationMins returns the DurationMins field if non-nil, zero value otherwise.

### GetDurationMinsOk

`func (o *SessionCfgRep) GetDurationMinsOk() (*int32, bool)`

GetDurationMinsOk returns a tuple with the DurationMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMins

`func (o *SessionCfgRep) SetDurationMins(v int32)`

SetDurationMins sets DurationMins field to given value.

### HasDurationMins

`func (o *SessionCfgRep) HasDurationMins() bool`

HasDurationMins returns a boolean if a field has been set.

### GetRefresh

`func (o *SessionCfgRep) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *SessionCfgRep) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *SessionCfgRep) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *SessionCfgRep) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


