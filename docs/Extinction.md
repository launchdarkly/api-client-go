# Extinction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Time** | **int64** |  | 
**FlagKey** | **string** |  | 
**ProjectKey** | **string** |  | 

## Methods

### NewExtinction

`func NewExtinction(revision string, time int64, flagKey string, projectKey string, ) *Extinction`

NewExtinction instantiates a new Extinction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtinctionWithDefaults

`func NewExtinctionWithDefaults() *Extinction`

NewExtinctionWithDefaults instantiates a new Extinction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *Extinction) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Extinction) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Extinction) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetMessage

`func (o *Extinction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Extinction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Extinction) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Extinction) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTime

`func (o *Extinction) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Extinction) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Extinction) SetTime(v int64)`

SetTime sets Time field to given value.


### GetFlagKey

`func (o *Extinction) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *Extinction) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *Extinction) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetProjectKey

`func (o *Extinction) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *Extinction) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *Extinction) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


