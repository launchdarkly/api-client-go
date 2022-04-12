# FlagRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetingRule** | Pointer to **string** |  | [optional] 
**FlagConfigVersion** | Pointer to **int32** |  | [optional] 
**Links** | [**map[string]Link**](Link.md) |  | 

## Methods

### NewFlagRep

`func NewFlagRep(links map[string]Link, ) *FlagRep`

NewFlagRep instantiates a new FlagRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagRepWithDefaults

`func NewFlagRepWithDefaults() *FlagRep`

NewFlagRepWithDefaults instantiates a new FlagRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetingRule

`func (o *FlagRep) GetTargetingRule() string`

GetTargetingRule returns the TargetingRule field if non-nil, zero value otherwise.

### GetTargetingRuleOk

`func (o *FlagRep) GetTargetingRuleOk() (*string, bool)`

GetTargetingRuleOk returns a tuple with the TargetingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRule

`func (o *FlagRep) SetTargetingRule(v string)`

SetTargetingRule sets TargetingRule field to given value.

### HasTargetingRule

`func (o *FlagRep) HasTargetingRule() bool`

HasTargetingRule returns a boolean if a field has been set.

### GetFlagConfigVersion

`func (o *FlagRep) GetFlagConfigVersion() int32`

GetFlagConfigVersion returns the FlagConfigVersion field if non-nil, zero value otherwise.

### GetFlagConfigVersionOk

`func (o *FlagRep) GetFlagConfigVersionOk() (*int32, bool)`

GetFlagConfigVersionOk returns a tuple with the FlagConfigVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagConfigVersion

`func (o *FlagRep) SetFlagConfigVersion(v int32)`

SetFlagConfigVersion sets FlagConfigVersion field to given value.

### HasFlagConfigVersion

`func (o *FlagRep) HasFlagConfigVersion() bool`

HasFlagConfigVersion returns a boolean if a field has been set.

### GetLinks

`func (o *FlagRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


