# FlagsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**LinkedFlags** | Pointer to [**ExpandedDirectlyLinkedFlags**](ExpandedDirectlyLinkedFlags.md) |  | [optional] 

## Methods

### NewFlagsSummary

`func NewFlagsSummary(count int32, ) *FlagsSummary`

NewFlagsSummary instantiates a new FlagsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagsSummaryWithDefaults

`func NewFlagsSummaryWithDefaults() *FlagsSummary`

NewFlagsSummaryWithDefaults instantiates a new FlagsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FlagsSummary) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FlagsSummary) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FlagsSummary) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLinkedFlags

`func (o *FlagsSummary) GetLinkedFlags() ExpandedDirectlyLinkedFlags`

GetLinkedFlags returns the LinkedFlags field if non-nil, zero value otherwise.

### GetLinkedFlagsOk

`func (o *FlagsSummary) GetLinkedFlagsOk() (*ExpandedDirectlyLinkedFlags, bool)`

GetLinkedFlagsOk returns a tuple with the LinkedFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFlags

`func (o *FlagsSummary) SetLinkedFlags(v ExpandedDirectlyLinkedFlags)`

SetLinkedFlags sets LinkedFlags field to given value.

### HasLinkedFlags

`func (o *FlagsSummary) HasLinkedFlags() bool`

HasLinkedFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


