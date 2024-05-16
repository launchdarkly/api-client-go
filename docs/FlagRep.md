# FlagRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetingRule** | Pointer to **string** | The targeting rule | [optional] 
**TargetingRuleDescription** | Pointer to **string** | The rule description | [optional] 
**TargetingRuleClauses** | Pointer to **[]interface{}** | An array of clauses used for individual targeting based on attributes | [optional] 
**FlagConfigVersion** | Pointer to **int32** | The flag version | [optional] 
**NotInExperimentVariationId** | Pointer to **string** | The ID of the variation to route traffic not part of the experiment analysis to | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

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

### GetTargetingRuleDescription

`func (o *FlagRep) GetTargetingRuleDescription() string`

GetTargetingRuleDescription returns the TargetingRuleDescription field if non-nil, zero value otherwise.

### GetTargetingRuleDescriptionOk

`func (o *FlagRep) GetTargetingRuleDescriptionOk() (*string, bool)`

GetTargetingRuleDescriptionOk returns a tuple with the TargetingRuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRuleDescription

`func (o *FlagRep) SetTargetingRuleDescription(v string)`

SetTargetingRuleDescription sets TargetingRuleDescription field to given value.

### HasTargetingRuleDescription

`func (o *FlagRep) HasTargetingRuleDescription() bool`

HasTargetingRuleDescription returns a boolean if a field has been set.

### GetTargetingRuleClauses

`func (o *FlagRep) GetTargetingRuleClauses() []interface{}`

GetTargetingRuleClauses returns the TargetingRuleClauses field if non-nil, zero value otherwise.

### GetTargetingRuleClausesOk

`func (o *FlagRep) GetTargetingRuleClausesOk() (*[]interface{}, bool)`

GetTargetingRuleClausesOk returns a tuple with the TargetingRuleClauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRuleClauses

`func (o *FlagRep) SetTargetingRuleClauses(v []interface{})`

SetTargetingRuleClauses sets TargetingRuleClauses field to given value.

### HasTargetingRuleClauses

`func (o *FlagRep) HasTargetingRuleClauses() bool`

HasTargetingRuleClauses returns a boolean if a field has been set.

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

### GetNotInExperimentVariationId

`func (o *FlagRep) GetNotInExperimentVariationId() string`

GetNotInExperimentVariationId returns the NotInExperimentVariationId field if non-nil, zero value otherwise.

### GetNotInExperimentVariationIdOk

`func (o *FlagRep) GetNotInExperimentVariationIdOk() (*string, bool)`

GetNotInExperimentVariationIdOk returns a tuple with the NotInExperimentVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInExperimentVariationId

`func (o *FlagRep) SetNotInExperimentVariationId(v string)`

SetNotInExperimentVariationId sets NotInExperimentVariationId field to given value.

### HasNotInExperimentVariationId

`func (o *FlagRep) HasNotInExperimentVariationId() bool`

HasNotInExperimentVariationId returns a boolean if a field has been set.

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


