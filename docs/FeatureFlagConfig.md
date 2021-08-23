# FeatureFlagConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | **bool** |  | 
**Archived** | **bool** |  | 
**Salt** | **string** |  | 
**Sel** | **string** |  | 
**LastModified** | **int64** |  | 
**Version** | **int32** |  | 
**Targets** | [**[]Target**](Target.md) |  | 
**Rules** | [**[]Rule**](Rule.md) |  | 
**Fallthrough** | [**VariationOrRolloutRep**](VariationOrRolloutRep.md) |  | 
**OffVariation** | Pointer to **int32** |  | [optional] 
**Prerequisites** | [**[]Prerequisite**](Prerequisite.md) |  | 
**Site** | [**Link**](Link.md) |  | 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 
**EnvironmentName** | **string** |  | 
**TrackEvents** | **bool** |  | 
**TrackEventsFallthrough** | **bool** |  | 
**DebugEventsUntilDate** | Pointer to **int64** |  | [optional] 
**Summary** | Pointer to [**FlagSummary**](FlagSummary.md) |  | [optional] 

## Methods

### NewFeatureFlagConfig

`func NewFeatureFlagConfig(on bool, archived bool, salt string, sel string, lastModified int64, version int32, targets []Target, rules []Rule, fallthrough_ VariationOrRolloutRep, prerequisites []Prerequisite, site Link, environmentName string, trackEvents bool, trackEventsFallthrough bool, ) *FeatureFlagConfig`

NewFeatureFlagConfig instantiates a new FeatureFlagConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagConfigWithDefaults

`func NewFeatureFlagConfigWithDefaults() *FeatureFlagConfig`

NewFeatureFlagConfigWithDefaults instantiates a new FeatureFlagConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOn

`func (o *FeatureFlagConfig) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *FeatureFlagConfig) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *FeatureFlagConfig) SetOn(v bool)`

SetOn sets On field to given value.


### GetArchived

`func (o *FeatureFlagConfig) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FeatureFlagConfig) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FeatureFlagConfig) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetSalt

`func (o *FeatureFlagConfig) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *FeatureFlagConfig) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *FeatureFlagConfig) SetSalt(v string)`

SetSalt sets Salt field to given value.


### GetSel

`func (o *FeatureFlagConfig) GetSel() string`

GetSel returns the Sel field if non-nil, zero value otherwise.

### GetSelOk

`func (o *FeatureFlagConfig) GetSelOk() (*string, bool)`

GetSelOk returns a tuple with the Sel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSel

`func (o *FeatureFlagConfig) SetSel(v string)`

SetSel sets Sel field to given value.


### GetLastModified

`func (o *FeatureFlagConfig) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *FeatureFlagConfig) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *FeatureFlagConfig) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.


### GetVersion

`func (o *FeatureFlagConfig) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeatureFlagConfig) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeatureFlagConfig) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetTargets

`func (o *FeatureFlagConfig) GetTargets() []Target`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *FeatureFlagConfig) GetTargetsOk() (*[]Target, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *FeatureFlagConfig) SetTargets(v []Target)`

SetTargets sets Targets field to given value.


### GetRules

`func (o *FeatureFlagConfig) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FeatureFlagConfig) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FeatureFlagConfig) SetRules(v []Rule)`

SetRules sets Rules field to given value.


### GetFallthrough

`func (o *FeatureFlagConfig) GetFallthrough() VariationOrRolloutRep`

GetFallthrough returns the Fallthrough field if non-nil, zero value otherwise.

### GetFallthroughOk

`func (o *FeatureFlagConfig) GetFallthroughOk() (*VariationOrRolloutRep, bool)`

GetFallthroughOk returns a tuple with the Fallthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallthrough

`func (o *FeatureFlagConfig) SetFallthrough(v VariationOrRolloutRep)`

SetFallthrough sets Fallthrough field to given value.


### GetOffVariation

`func (o *FeatureFlagConfig) GetOffVariation() int32`

GetOffVariation returns the OffVariation field if non-nil, zero value otherwise.

### GetOffVariationOk

`func (o *FeatureFlagConfig) GetOffVariationOk() (*int32, bool)`

GetOffVariationOk returns a tuple with the OffVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffVariation

`func (o *FeatureFlagConfig) SetOffVariation(v int32)`

SetOffVariation sets OffVariation field to given value.

### HasOffVariation

`func (o *FeatureFlagConfig) HasOffVariation() bool`

HasOffVariation returns a boolean if a field has been set.

### GetPrerequisites

`func (o *FeatureFlagConfig) GetPrerequisites() []Prerequisite`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *FeatureFlagConfig) GetPrerequisitesOk() (*[]Prerequisite, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *FeatureFlagConfig) SetPrerequisites(v []Prerequisite)`

SetPrerequisites sets Prerequisites field to given value.


### GetSite

`func (o *FeatureFlagConfig) GetSite() Link`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *FeatureFlagConfig) GetSiteOk() (*Link, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *FeatureFlagConfig) SetSite(v Link)`

SetSite sets Site field to given value.


### GetAccess

`func (o *FeatureFlagConfig) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FeatureFlagConfig) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FeatureFlagConfig) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FeatureFlagConfig) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *FeatureFlagConfig) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *FeatureFlagConfig) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *FeatureFlagConfig) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetTrackEvents

`func (o *FeatureFlagConfig) GetTrackEvents() bool`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *FeatureFlagConfig) GetTrackEventsOk() (*bool, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *FeatureFlagConfig) SetTrackEvents(v bool)`

SetTrackEvents sets TrackEvents field to given value.


### GetTrackEventsFallthrough

`func (o *FeatureFlagConfig) GetTrackEventsFallthrough() bool`

GetTrackEventsFallthrough returns the TrackEventsFallthrough field if non-nil, zero value otherwise.

### GetTrackEventsFallthroughOk

`func (o *FeatureFlagConfig) GetTrackEventsFallthroughOk() (*bool, bool)`

GetTrackEventsFallthroughOk returns a tuple with the TrackEventsFallthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEventsFallthrough

`func (o *FeatureFlagConfig) SetTrackEventsFallthrough(v bool)`

SetTrackEventsFallthrough sets TrackEventsFallthrough field to given value.


### GetDebugEventsUntilDate

`func (o *FeatureFlagConfig) GetDebugEventsUntilDate() int64`

GetDebugEventsUntilDate returns the DebugEventsUntilDate field if non-nil, zero value otherwise.

### GetDebugEventsUntilDateOk

`func (o *FeatureFlagConfig) GetDebugEventsUntilDateOk() (*int64, bool)`

GetDebugEventsUntilDateOk returns a tuple with the DebugEventsUntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEventsUntilDate

`func (o *FeatureFlagConfig) SetDebugEventsUntilDate(v int64)`

SetDebugEventsUntilDate sets DebugEventsUntilDate field to given value.

### HasDebugEventsUntilDate

`func (o *FeatureFlagConfig) HasDebugEventsUntilDate() bool`

HasDebugEventsUntilDate returns a boolean if a field has been set.

### GetSummary

`func (o *FeatureFlagConfig) GetSummary() FlagSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FeatureFlagConfig) GetSummaryOk() (*FlagSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FeatureFlagConfig) SetSummary(v FlagSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *FeatureFlagConfig) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


