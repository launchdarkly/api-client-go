# FlagConfigurationRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | **bool** |  | 
**Archived** | **bool** |  | 
**Salt** | **string** |  | 
**Sel** | **string** |  | 
**LastModified** | **int64** |  | 
**Version** | **int32** |  | 
**Targets** | [**[]FlagConfigurationRepTargets**](FlagConfigurationRepTargets.md) |  | 
**Rules** | [**[]RuleRep**](RuleRep.md) |  | 
**Fallthrough** | [**VariationOrRolloutRep**](VariationOrRolloutRep.md) |  | 
**OffVariation** | Pointer to **int32** |  | [optional] 
**Prerequisites** | [**[]FlagConfigurationRepPrerequisites**](FlagConfigurationRepPrerequisites.md) |  | 
**Site** | [**CoreLink**](CoreLink.md) |  | 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 
**EnvironmentName** | **string** |  | 
**TrackEvents** | **bool** |  | 
**TrackEventsFallthrough** | **bool** |  | 
**DebugEventsUntilDate** | Pointer to **int64** |  | [optional] 
**Summary** | Pointer to [**FlagSummary**](FlagSummary.md) |  | [optional] 

## Methods

### NewFlagConfigurationRep

`func NewFlagConfigurationRep(on bool, archived bool, salt string, sel string, lastModified int64, version int32, targets []FlagConfigurationRepTargets, rules []RuleRep, fallthrough_ VariationOrRolloutRep, prerequisites []FlagConfigurationRepPrerequisites, site CoreLink, environmentName string, trackEvents bool, trackEventsFallthrough bool, ) *FlagConfigurationRep`

NewFlagConfigurationRep instantiates a new FlagConfigurationRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagConfigurationRepWithDefaults

`func NewFlagConfigurationRepWithDefaults() *FlagConfigurationRep`

NewFlagConfigurationRepWithDefaults instantiates a new FlagConfigurationRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOn

`func (o *FlagConfigurationRep) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *FlagConfigurationRep) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *FlagConfigurationRep) SetOn(v bool)`

SetOn sets On field to given value.


### GetArchived

`func (o *FlagConfigurationRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FlagConfigurationRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FlagConfigurationRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetSalt

`func (o *FlagConfigurationRep) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *FlagConfigurationRep) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *FlagConfigurationRep) SetSalt(v string)`

SetSalt sets Salt field to given value.


### GetSel

`func (o *FlagConfigurationRep) GetSel() string`

GetSel returns the Sel field if non-nil, zero value otherwise.

### GetSelOk

`func (o *FlagConfigurationRep) GetSelOk() (*string, bool)`

GetSelOk returns a tuple with the Sel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSel

`func (o *FlagConfigurationRep) SetSel(v string)`

SetSel sets Sel field to given value.


### GetLastModified

`func (o *FlagConfigurationRep) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *FlagConfigurationRep) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *FlagConfigurationRep) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.


### GetVersion

`func (o *FlagConfigurationRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FlagConfigurationRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FlagConfigurationRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetTargets

`func (o *FlagConfigurationRep) GetTargets() []FlagConfigurationRepTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *FlagConfigurationRep) GetTargetsOk() (*[]FlagConfigurationRepTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *FlagConfigurationRep) SetTargets(v []FlagConfigurationRepTargets)`

SetTargets sets Targets field to given value.


### GetRules

`func (o *FlagConfigurationRep) GetRules() []RuleRep`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FlagConfigurationRep) GetRulesOk() (*[]RuleRep, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FlagConfigurationRep) SetRules(v []RuleRep)`

SetRules sets Rules field to given value.


### GetFallthrough

`func (o *FlagConfigurationRep) GetFallthrough() VariationOrRolloutRep`

GetFallthrough returns the Fallthrough field if non-nil, zero value otherwise.

### GetFallthroughOk

`func (o *FlagConfigurationRep) GetFallthroughOk() (*VariationOrRolloutRep, bool)`

GetFallthroughOk returns a tuple with the Fallthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallthrough

`func (o *FlagConfigurationRep) SetFallthrough(v VariationOrRolloutRep)`

SetFallthrough sets Fallthrough field to given value.


### GetOffVariation

`func (o *FlagConfigurationRep) GetOffVariation() int32`

GetOffVariation returns the OffVariation field if non-nil, zero value otherwise.

### GetOffVariationOk

`func (o *FlagConfigurationRep) GetOffVariationOk() (*int32, bool)`

GetOffVariationOk returns a tuple with the OffVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffVariation

`func (o *FlagConfigurationRep) SetOffVariation(v int32)`

SetOffVariation sets OffVariation field to given value.

### HasOffVariation

`func (o *FlagConfigurationRep) HasOffVariation() bool`

HasOffVariation returns a boolean if a field has been set.

### GetPrerequisites

`func (o *FlagConfigurationRep) GetPrerequisites() []FlagConfigurationRepPrerequisites`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *FlagConfigurationRep) GetPrerequisitesOk() (*[]FlagConfigurationRepPrerequisites, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *FlagConfigurationRep) SetPrerequisites(v []FlagConfigurationRepPrerequisites)`

SetPrerequisites sets Prerequisites field to given value.


### GetSite

`func (o *FlagConfigurationRep) GetSite() CoreLink`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *FlagConfigurationRep) GetSiteOk() (*CoreLink, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *FlagConfigurationRep) SetSite(v CoreLink)`

SetSite sets Site field to given value.


### GetAccess

`func (o *FlagConfigurationRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FlagConfigurationRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FlagConfigurationRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FlagConfigurationRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *FlagConfigurationRep) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *FlagConfigurationRep) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *FlagConfigurationRep) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetTrackEvents

`func (o *FlagConfigurationRep) GetTrackEvents() bool`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *FlagConfigurationRep) GetTrackEventsOk() (*bool, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *FlagConfigurationRep) SetTrackEvents(v bool)`

SetTrackEvents sets TrackEvents field to given value.


### GetTrackEventsFallthrough

`func (o *FlagConfigurationRep) GetTrackEventsFallthrough() bool`

GetTrackEventsFallthrough returns the TrackEventsFallthrough field if non-nil, zero value otherwise.

### GetTrackEventsFallthroughOk

`func (o *FlagConfigurationRep) GetTrackEventsFallthroughOk() (*bool, bool)`

GetTrackEventsFallthroughOk returns a tuple with the TrackEventsFallthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEventsFallthrough

`func (o *FlagConfigurationRep) SetTrackEventsFallthrough(v bool)`

SetTrackEventsFallthrough sets TrackEventsFallthrough field to given value.


### GetDebugEventsUntilDate

`func (o *FlagConfigurationRep) GetDebugEventsUntilDate() int64`

GetDebugEventsUntilDate returns the DebugEventsUntilDate field if non-nil, zero value otherwise.

### GetDebugEventsUntilDateOk

`func (o *FlagConfigurationRep) GetDebugEventsUntilDateOk() (*int64, bool)`

GetDebugEventsUntilDateOk returns a tuple with the DebugEventsUntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEventsUntilDate

`func (o *FlagConfigurationRep) SetDebugEventsUntilDate(v int64)`

SetDebugEventsUntilDate sets DebugEventsUntilDate field to given value.

### HasDebugEventsUntilDate

`func (o *FlagConfigurationRep) HasDebugEventsUntilDate() bool`

HasDebugEventsUntilDate returns a boolean if a field has been set.

### GetSummary

`func (o *FlagConfigurationRep) GetSummary() FlagSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FlagConfigurationRep) GetSummaryOk() (*FlagSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FlagConfigurationRep) SetSummary(v FlagSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *FlagConfigurationRep) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


