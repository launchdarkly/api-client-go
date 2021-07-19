# FlagStatusesRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**LastRequested** | Pointer to **time.Time** |  | [optional] 
**Default** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFlagStatusesRep

`func NewFlagStatusesRep() *FlagStatusesRep`

NewFlagStatusesRep instantiates a new FlagStatusesRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagStatusesRepWithDefaults

`func NewFlagStatusesRepWithDefaults() *FlagStatusesRep`

NewFlagStatusesRepWithDefaults instantiates a new FlagStatusesRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlagStatusesRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagStatusesRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagStatusesRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlagStatusesRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastRequested

`func (o *FlagStatusesRep) GetLastRequested() time.Time`

GetLastRequested returns the LastRequested field if non-nil, zero value otherwise.

### GetLastRequestedOk

`func (o *FlagStatusesRep) GetLastRequestedOk() (*time.Time, bool)`

GetLastRequestedOk returns a tuple with the LastRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRequested

`func (o *FlagStatusesRep) SetLastRequested(v time.Time)`

SetLastRequested sets LastRequested field to given value.

### HasLastRequested

`func (o *FlagStatusesRep) HasLastRequested() bool`

HasLastRequested returns a boolean if a field has been set.

### GetDefault

`func (o *FlagStatusesRep) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FlagStatusesRep) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FlagStatusesRep) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FlagStatusesRep) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *FlagStatusesRep) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *FlagStatusesRep) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


