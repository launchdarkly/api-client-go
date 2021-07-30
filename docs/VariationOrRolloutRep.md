# VariationOrRolloutRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variation** | Pointer to **int32** |  | [optional] 
**Rollout** | Pointer to [**RolloutRep**](RolloutRep.md) |  | [optional] 
**Clauses** | Pointer to [**[]RuleRepClauses**](RuleRepClauses.md) |  | [optional] 
**TrackEvents** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewVariationOrRolloutRep

`func NewVariationOrRolloutRep() *VariationOrRolloutRep`

NewVariationOrRolloutRep instantiates a new VariationOrRolloutRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationOrRolloutRepWithDefaults

`func NewVariationOrRolloutRepWithDefaults() *VariationOrRolloutRep`

NewVariationOrRolloutRepWithDefaults instantiates a new VariationOrRolloutRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariation

`func (o *VariationOrRolloutRep) GetVariation() int32`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *VariationOrRolloutRep) GetVariationOk() (*int32, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *VariationOrRolloutRep) SetVariation(v int32)`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *VariationOrRolloutRep) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetRollout

`func (o *VariationOrRolloutRep) GetRollout() RolloutRep`

GetRollout returns the Rollout field if non-nil, zero value otherwise.

### GetRolloutOk

`func (o *VariationOrRolloutRep) GetRolloutOk() (*RolloutRep, bool)`

GetRolloutOk returns a tuple with the Rollout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollout

`func (o *VariationOrRolloutRep) SetRollout(v RolloutRep)`

SetRollout sets Rollout field to given value.

### HasRollout

`func (o *VariationOrRolloutRep) HasRollout() bool`

HasRollout returns a boolean if a field has been set.

### GetClauses

`func (o *VariationOrRolloutRep) GetClauses() []RuleRepClauses`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *VariationOrRolloutRep) GetClausesOk() (*[]RuleRepClauses, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *VariationOrRolloutRep) SetClauses(v []RuleRepClauses)`

SetClauses sets Clauses field to given value.

### HasClauses

`func (o *VariationOrRolloutRep) HasClauses() bool`

HasClauses returns a boolean if a field has been set.

### GetTrackEvents

`func (o *VariationOrRolloutRep) GetTrackEvents() bool`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *VariationOrRolloutRep) GetTrackEventsOk() (*bool, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *VariationOrRolloutRep) SetTrackEvents(v bool)`

SetTrackEvents sets TrackEvents field to given value.

### HasTrackEvents

`func (o *VariationOrRolloutRep) HasTrackEvents() bool`

HasTrackEvents returns a boolean if a field has been set.

### GetDescription

`func (o *VariationOrRolloutRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariationOrRolloutRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariationOrRolloutRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariationOrRolloutRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


