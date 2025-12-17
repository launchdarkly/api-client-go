# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The flag rule ID | [optional] 
**Disabled** | Pointer to **bool** | Whether the rule is disabled | [optional] 
**Variation** | Pointer to **int32** | The index of the variation, from the array of variations for this flag | [optional] 
**Rollout** | Pointer to [**Rollout**](Rollout.md) |  | [optional] 
**Clauses** | [**[]Clause**](Clause.md) | An array of clauses used for individual targeting based on attributes | 
**TrackEvents** | **bool** | Whether LaunchDarkly tracks events for this rule | 
**Description** | Pointer to **string** | The rule description | [optional] 
**Ref** | Pointer to **string** |  | [optional] 

## Methods

### NewRule

`func NewRule(clauses []Clause, trackEvents bool, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Rule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisabled

`func (o *Rule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Rule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Rule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Rule) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetVariation

`func (o *Rule) GetVariation() int32`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *Rule) GetVariationOk() (*int32, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *Rule) SetVariation(v int32)`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *Rule) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetRollout

`func (o *Rule) GetRollout() Rollout`

GetRollout returns the Rollout field if non-nil, zero value otherwise.

### GetRolloutOk

`func (o *Rule) GetRolloutOk() (*Rollout, bool)`

GetRolloutOk returns a tuple with the Rollout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollout

`func (o *Rule) SetRollout(v Rollout)`

SetRollout sets Rollout field to given value.

### HasRollout

`func (o *Rule) HasRollout() bool`

HasRollout returns a boolean if a field has been set.

### GetClauses

`func (o *Rule) GetClauses() []Clause`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *Rule) GetClausesOk() (*[]Clause, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *Rule) SetClauses(v []Clause)`

SetClauses sets Clauses field to given value.


### GetTrackEvents

`func (o *Rule) GetTrackEvents() bool`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *Rule) GetTrackEventsOk() (*bool, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *Rule) SetTrackEvents(v bool)`

SetTrackEvents sets TrackEvents field to given value.


### GetDescription

`func (o *Rule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Rule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Rule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Rule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRef

`func (o *Rule) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Rule) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Rule) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Rule) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


