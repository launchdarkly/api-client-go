# FlagStatusRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LastRequested** | Pointer to **time.Time** |  | [optional] 
**Default** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFlagStatusRep

`func NewFlagStatusRep() *FlagStatusRep`

NewFlagStatusRep instantiates a new FlagStatusRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagStatusRepWithDefaults

`func NewFlagStatusRepWithDefaults() *FlagStatusRep`

NewFlagStatusRepWithDefaults instantiates a new FlagStatusRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlagStatusRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagStatusRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagStatusRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlagStatusRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *FlagStatusRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagStatusRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagStatusRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlagStatusRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastRequested

`func (o *FlagStatusRep) GetLastRequested() time.Time`

GetLastRequested returns the LastRequested field if non-nil, zero value otherwise.

### GetLastRequestedOk

`func (o *FlagStatusRep) GetLastRequestedOk() (*time.Time, bool)`

GetLastRequestedOk returns a tuple with the LastRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRequested

`func (o *FlagStatusRep) SetLastRequested(v time.Time)`

SetLastRequested sets LastRequested field to given value.

### HasLastRequested

`func (o *FlagStatusRep) HasLastRequested() bool`

HasLastRequested returns a boolean if a field has been set.

### GetDefault

`func (o *FlagStatusRep) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FlagStatusRep) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FlagStatusRep) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FlagStatusRep) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *FlagStatusRep) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *FlagStatusRep) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


