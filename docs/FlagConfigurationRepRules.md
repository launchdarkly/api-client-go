# FlagConfigurationRepRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Variation** | Pointer to **int32** |  | [optional] 
**Rollout** | Pointer to [**RolloutRep**](RolloutRep.md) |  | [optional] 
**Clauses** | Pointer to [**[]FlagConfigurationRepClauses**](FlagConfigurationRepClauses.md) |  | [optional] 
**TrackEvents** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewFlagConfigurationRepRules

`func NewFlagConfigurationRepRules() *FlagConfigurationRepRules`

NewFlagConfigurationRepRules instantiates a new FlagConfigurationRepRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagConfigurationRepRulesWithDefaults

`func NewFlagConfigurationRepRulesWithDefaults() *FlagConfigurationRepRules`

NewFlagConfigurationRepRulesWithDefaults instantiates a new FlagConfigurationRepRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlagConfigurationRepRules) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagConfigurationRepRules) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagConfigurationRepRules) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlagConfigurationRepRules) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVariation

`func (o *FlagConfigurationRepRules) GetVariation() int32`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *FlagConfigurationRepRules) GetVariationOk() (*int32, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *FlagConfigurationRepRules) SetVariation(v int32)`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *FlagConfigurationRepRules) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetRollout

`func (o *FlagConfigurationRepRules) GetRollout() RolloutRep`

GetRollout returns the Rollout field if non-nil, zero value otherwise.

### GetRolloutOk

`func (o *FlagConfigurationRepRules) GetRolloutOk() (*RolloutRep, bool)`

GetRolloutOk returns a tuple with the Rollout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollout

`func (o *FlagConfigurationRepRules) SetRollout(v RolloutRep)`

SetRollout sets Rollout field to given value.

### HasRollout

`func (o *FlagConfigurationRepRules) HasRollout() bool`

HasRollout returns a boolean if a field has been set.

### GetClauses

`func (o *FlagConfigurationRepRules) GetClauses() []FlagConfigurationRepClauses`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *FlagConfigurationRepRules) GetClausesOk() (*[]FlagConfigurationRepClauses, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *FlagConfigurationRepRules) SetClauses(v []FlagConfigurationRepClauses)`

SetClauses sets Clauses field to given value.

### HasClauses

`func (o *FlagConfigurationRepRules) HasClauses() bool`

HasClauses returns a boolean if a field has been set.

### GetTrackEvents

`func (o *FlagConfigurationRepRules) GetTrackEvents() bool`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *FlagConfigurationRepRules) GetTrackEventsOk() (*bool, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *FlagConfigurationRepRules) SetTrackEvents(v bool)`

SetTrackEvents sets TrackEvents field to given value.

### HasTrackEvents

`func (o *FlagConfigurationRepRules) HasTrackEvents() bool`

HasTrackEvents returns a boolean if a field has been set.

### GetDescription

`func (o *FlagConfigurationRepRules) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagConfigurationRepRules) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagConfigurationRepRules) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlagConfigurationRepRules) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


