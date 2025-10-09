# AIConfigTargetingEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextTargets** | [**[]AIConfigTargetingEnvironmentTarget**](AIConfigTargetingEnvironmentTarget.md) |  | 
**Enabled** | **bool** |  | 
**Fallthrough** | [**AIConfigTargetingEnvironmentFallthrough**](AIConfigTargetingEnvironmentFallthrough.md) |  | 
**LastModified** | **int64** |  | 
**OffVariation** | Pointer to **int32** |  | [optional] 
**Rules** | [**[]AIConfigTargetingEnvironmentRule**](AIConfigTargetingEnvironmentRule.md) |  | 
**Targets** | [**[]AIConfigTargetingEnvironmentTarget**](AIConfigTargetingEnvironmentTarget.md) |  | 
**TrackEvents** | **bool** |  | 
**TrackEventsFallthrough** | **bool** |  | 
**EnvironmentName** | **string** |  | 
**Version** | **int32** |  | 

## Methods

### NewAIConfigTargetingEnvironment

`func NewAIConfigTargetingEnvironment(contextTargets []AIConfigTargetingEnvironmentTarget, enabled bool, fallthrough_ AIConfigTargetingEnvironmentFallthrough, lastModified int64, rules []AIConfigTargetingEnvironmentRule, targets []AIConfigTargetingEnvironmentTarget, trackEvents bool, trackEventsFallthrough bool, environmentName string, version int32, ) *AIConfigTargetingEnvironment`

NewAIConfigTargetingEnvironment instantiates a new AIConfigTargetingEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigTargetingEnvironmentWithDefaults

`func NewAIConfigTargetingEnvironmentWithDefaults() *AIConfigTargetingEnvironment`

NewAIConfigTargetingEnvironmentWithDefaults instantiates a new AIConfigTargetingEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextTargets

`func (o *AIConfigTargetingEnvironment) GetContextTargets() []AIConfigTargetingEnvironmentTarget`

GetContextTargets returns the ContextTargets field if non-nil, zero value otherwise.

### GetContextTargetsOk

`func (o *AIConfigTargetingEnvironment) GetContextTargetsOk() (*[]AIConfigTargetingEnvironmentTarget, bool)`

GetContextTargetsOk returns a tuple with the ContextTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextTargets

`func (o *AIConfigTargetingEnvironment) SetContextTargets(v []AIConfigTargetingEnvironmentTarget)`

SetContextTargets sets ContextTargets field to given value.


### GetEnabled

`func (o *AIConfigTargetingEnvironment) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AIConfigTargetingEnvironment) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AIConfigTargetingEnvironment) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFallthrough

`func (o *AIConfigTargetingEnvironment) GetFallthrough() AIConfigTargetingEnvironmentFallthrough`

GetFallthrough returns the Fallthrough field if non-nil, zero value otherwise.

### GetFallthroughOk

`func (o *AIConfigTargetingEnvironment) GetFallthroughOk() (*AIConfigTargetingEnvironmentFallthrough, bool)`

GetFallthroughOk returns a tuple with the Fallthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallthrough

`func (o *AIConfigTargetingEnvironment) SetFallthrough(v AIConfigTargetingEnvironmentFallthrough)`

SetFallthrough sets Fallthrough field to given value.


### GetLastModified

`func (o *AIConfigTargetingEnvironment) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AIConfigTargetingEnvironment) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AIConfigTargetingEnvironment) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.


### GetOffVariation

`func (o *AIConfigTargetingEnvironment) GetOffVariation() int32`

GetOffVariation returns the OffVariation field if non-nil, zero value otherwise.

### GetOffVariationOk

`func (o *AIConfigTargetingEnvironment) GetOffVariationOk() (*int32, bool)`

GetOffVariationOk returns a tuple with the OffVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffVariation

`func (o *AIConfigTargetingEnvironment) SetOffVariation(v int32)`

SetOffVariation sets OffVariation field to given value.

### HasOffVariation

`func (o *AIConfigTargetingEnvironment) HasOffVariation() bool`

HasOffVariation returns a boolean if a field has been set.

### GetRules

`func (o *AIConfigTargetingEnvironment) GetRules() []AIConfigTargetingEnvironmentRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AIConfigTargetingEnvironment) GetRulesOk() (*[]AIConfigTargetingEnvironmentRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AIConfigTargetingEnvironment) SetRules(v []AIConfigTargetingEnvironmentRule)`

SetRules sets Rules field to given value.


### GetTargets

`func (o *AIConfigTargetingEnvironment) GetTargets() []AIConfigTargetingEnvironmentTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AIConfigTargetingEnvironment) GetTargetsOk() (*[]AIConfigTargetingEnvironmentTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AIConfigTargetingEnvironment) SetTargets(v []AIConfigTargetingEnvironmentTarget)`

SetTargets sets Targets field to given value.


### GetTrackEvents

`func (o *AIConfigTargetingEnvironment) GetTrackEvents() bool`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *AIConfigTargetingEnvironment) GetTrackEventsOk() (*bool, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *AIConfigTargetingEnvironment) SetTrackEvents(v bool)`

SetTrackEvents sets TrackEvents field to given value.


### GetTrackEventsFallthrough

`func (o *AIConfigTargetingEnvironment) GetTrackEventsFallthrough() bool`

GetTrackEventsFallthrough returns the TrackEventsFallthrough field if non-nil, zero value otherwise.

### GetTrackEventsFallthroughOk

`func (o *AIConfigTargetingEnvironment) GetTrackEventsFallthroughOk() (*bool, bool)`

GetTrackEventsFallthroughOk returns a tuple with the TrackEventsFallthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEventsFallthrough

`func (o *AIConfigTargetingEnvironment) SetTrackEventsFallthrough(v bool)`

SetTrackEventsFallthrough sets TrackEventsFallthrough field to given value.


### GetEnvironmentName

`func (o *AIConfigTargetingEnvironment) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *AIConfigTargetingEnvironment) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *AIConfigTargetingEnvironment) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetVersion

`func (o *AIConfigTargetingEnvironment) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AIConfigTargetingEnvironment) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AIConfigTargetingEnvironment) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


