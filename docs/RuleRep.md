# RuleRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variation** | Pointer to **int32** |  | [optional] 
**Rollout** | Pointer to [**RolloutRep**](RolloutRep.md) |  | [optional] 
**Clauses** | [**[]RuleRepClauses**](RuleRepClauses.md) |  | 
**TrackEvents** | **bool** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewRuleRep

`func NewRuleRep(clauses []RuleRepClauses, trackEvents bool, ) *RuleRep`

NewRuleRep instantiates a new RuleRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleRepWithDefaults

`func NewRuleRepWithDefaults() *RuleRep`

NewRuleRepWithDefaults instantiates a new RuleRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariation

`func (o *RuleRep) GetVariation() int32`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *RuleRep) GetVariationOk() (*int32, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *RuleRep) SetVariation(v int32)`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *RuleRep) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetRollout

`func (o *RuleRep) GetRollout() RolloutRep`

GetRollout returns the Rollout field if non-nil, zero value otherwise.

### GetRolloutOk

`func (o *RuleRep) GetRolloutOk() (*RolloutRep, bool)`

GetRolloutOk returns a tuple with the Rollout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollout

`func (o *RuleRep) SetRollout(v RolloutRep)`

SetRollout sets Rollout field to given value.

### HasRollout

`func (o *RuleRep) HasRollout() bool`

HasRollout returns a boolean if a field has been set.

### GetClauses

`func (o *RuleRep) GetClauses() []RuleRepClauses`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *RuleRep) GetClausesOk() (*[]RuleRepClauses, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *RuleRep) SetClauses(v []RuleRepClauses)`

SetClauses sets Clauses field to given value.


### GetTrackEvents

`func (o *RuleRep) GetTrackEvents() bool`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *RuleRep) GetTrackEventsOk() (*bool, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *RuleRep) SetTrackEvents(v bool)`

SetTrackEvents sets TrackEvents field to given value.


### GetDescription

`func (o *RuleRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


