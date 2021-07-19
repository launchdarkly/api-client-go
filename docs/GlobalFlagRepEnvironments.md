# GlobalFlagRepEnvironments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | Pointer to **bool** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Salt** | Pointer to **string** |  | [optional] 
**Sel** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Targets** | Pointer to [**[]FlagConfigurationRepTargets**](FlagConfigurationRepTargets.md) |  | [optional] 
**Rules** | Pointer to [**[]FlagConfigurationRepRules**](FlagConfigurationRepRules.md) |  | [optional] 
**Fallthrough** | Pointer to [**VariationOrRolloutRep**](VariationOrRolloutRep.md) |  | [optional] 
**OffVariation** | Pointer to **int32** |  | [optional] 
**Prerequisites** | Pointer to [**[]FlagConfigurationRepPrerequisites**](FlagConfigurationRepPrerequisites.md) |  | [optional] 
**Site** | Pointer to [**CoreLink**](CoreLink.md) |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 
**EnvironmentName** | Pointer to **string** |  | [optional] 
**TrackEvents** | Pointer to **bool** |  | [optional] 
**TrackEventsFallthrough** | Pointer to **bool** |  | [optional] 
**DebugEventsUntilDate** | Pointer to **int64** |  | [optional] 
**Summary** | Pointer to [**FlagSummary**](FlagSummary.md) |  | [optional] 

## Methods

### NewGlobalFlagRepEnvironments

`func NewGlobalFlagRepEnvironments() *GlobalFlagRepEnvironments`

NewGlobalFlagRepEnvironments instantiates a new GlobalFlagRepEnvironments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalFlagRepEnvironmentsWithDefaults

`func NewGlobalFlagRepEnvironmentsWithDefaults() *GlobalFlagRepEnvironments`

NewGlobalFlagRepEnvironmentsWithDefaults instantiates a new GlobalFlagRepEnvironments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOn

`func (o *GlobalFlagRepEnvironments) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *GlobalFlagRepEnvironments) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *GlobalFlagRepEnvironments) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *GlobalFlagRepEnvironments) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetArchived

`func (o *GlobalFlagRepEnvironments) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *GlobalFlagRepEnvironments) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *GlobalFlagRepEnvironments) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *GlobalFlagRepEnvironments) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetSalt

`func (o *GlobalFlagRepEnvironments) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *GlobalFlagRepEnvironments) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *GlobalFlagRepEnvironments) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *GlobalFlagRepEnvironments) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetSel

`func (o *GlobalFlagRepEnvironments) GetSel() string`

GetSel returns the Sel field if non-nil, zero value otherwise.

### GetSelOk

`func (o *GlobalFlagRepEnvironments) GetSelOk() (*string, bool)`

GetSelOk returns a tuple with the Sel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSel

`func (o *GlobalFlagRepEnvironments) SetSel(v string)`

SetSel sets Sel field to given value.

### HasSel

`func (o *GlobalFlagRepEnvironments) HasSel() bool`

HasSel returns a boolean if a field has been set.

### GetLastModified

`func (o *GlobalFlagRepEnvironments) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *GlobalFlagRepEnvironments) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *GlobalFlagRepEnvironments) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *GlobalFlagRepEnvironments) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetVersion

`func (o *GlobalFlagRepEnvironments) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GlobalFlagRepEnvironments) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GlobalFlagRepEnvironments) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GlobalFlagRepEnvironments) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTargets

`func (o *GlobalFlagRepEnvironments) GetTargets() []FlagConfigurationRepTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *GlobalFlagRepEnvironments) GetTargetsOk() (*[]FlagConfigurationRepTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *GlobalFlagRepEnvironments) SetTargets(v []FlagConfigurationRepTargets)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *GlobalFlagRepEnvironments) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetRules

`func (o *GlobalFlagRepEnvironments) GetRules() []FlagConfigurationRepRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GlobalFlagRepEnvironments) GetRulesOk() (*[]FlagConfigurationRepRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GlobalFlagRepEnvironments) SetRules(v []FlagConfigurationRepRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GlobalFlagRepEnvironments) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetFallthrough

`func (o *GlobalFlagRepEnvironments) GetFallthrough() VariationOrRolloutRep`

GetFallthrough returns the Fallthrough field if non-nil, zero value otherwise.

### GetFallthroughOk

`func (o *GlobalFlagRepEnvironments) GetFallthroughOk() (*VariationOrRolloutRep, bool)`

GetFallthroughOk returns a tuple with the Fallthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallthrough

`func (o *GlobalFlagRepEnvironments) SetFallthrough(v VariationOrRolloutRep)`

SetFallthrough sets Fallthrough field to given value.

### HasFallthrough

`func (o *GlobalFlagRepEnvironments) HasFallthrough() bool`

HasFallthrough returns a boolean if a field has been set.

### GetOffVariation

`func (o *GlobalFlagRepEnvironments) GetOffVariation() int32`

GetOffVariation returns the OffVariation field if non-nil, zero value otherwise.

### GetOffVariationOk

`func (o *GlobalFlagRepEnvironments) GetOffVariationOk() (*int32, bool)`

GetOffVariationOk returns a tuple with the OffVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffVariation

`func (o *GlobalFlagRepEnvironments) SetOffVariation(v int32)`

SetOffVariation sets OffVariation field to given value.

### HasOffVariation

`func (o *GlobalFlagRepEnvironments) HasOffVariation() bool`

HasOffVariation returns a boolean if a field has been set.

### GetPrerequisites

`func (o *GlobalFlagRepEnvironments) GetPrerequisites() []FlagConfigurationRepPrerequisites`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *GlobalFlagRepEnvironments) GetPrerequisitesOk() (*[]FlagConfigurationRepPrerequisites, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *GlobalFlagRepEnvironments) SetPrerequisites(v []FlagConfigurationRepPrerequisites)`

SetPrerequisites sets Prerequisites field to given value.

### HasPrerequisites

`func (o *GlobalFlagRepEnvironments) HasPrerequisites() bool`

HasPrerequisites returns a boolean if a field has been set.

### GetSite

`func (o *GlobalFlagRepEnvironments) GetSite() CoreLink`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GlobalFlagRepEnvironments) GetSiteOk() (*CoreLink, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GlobalFlagRepEnvironments) SetSite(v CoreLink)`

SetSite sets Site field to given value.

### HasSite

`func (o *GlobalFlagRepEnvironments) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAccess

`func (o *GlobalFlagRepEnvironments) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *GlobalFlagRepEnvironments) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *GlobalFlagRepEnvironments) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *GlobalFlagRepEnvironments) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *GlobalFlagRepEnvironments) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *GlobalFlagRepEnvironments) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *GlobalFlagRepEnvironments) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *GlobalFlagRepEnvironments) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetTrackEvents

`func (o *GlobalFlagRepEnvironments) GetTrackEvents() bool`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *GlobalFlagRepEnvironments) GetTrackEventsOk() (*bool, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *GlobalFlagRepEnvironments) SetTrackEvents(v bool)`

SetTrackEvents sets TrackEvents field to given value.

### HasTrackEvents

`func (o *GlobalFlagRepEnvironments) HasTrackEvents() bool`

HasTrackEvents returns a boolean if a field has been set.

### GetTrackEventsFallthrough

`func (o *GlobalFlagRepEnvironments) GetTrackEventsFallthrough() bool`

GetTrackEventsFallthrough returns the TrackEventsFallthrough field if non-nil, zero value otherwise.

### GetTrackEventsFallthroughOk

`func (o *GlobalFlagRepEnvironments) GetTrackEventsFallthroughOk() (*bool, bool)`

GetTrackEventsFallthroughOk returns a tuple with the TrackEventsFallthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEventsFallthrough

`func (o *GlobalFlagRepEnvironments) SetTrackEventsFallthrough(v bool)`

SetTrackEventsFallthrough sets TrackEventsFallthrough field to given value.

### HasTrackEventsFallthrough

`func (o *GlobalFlagRepEnvironments) HasTrackEventsFallthrough() bool`

HasTrackEventsFallthrough returns a boolean if a field has been set.

### GetDebugEventsUntilDate

`func (o *GlobalFlagRepEnvironments) GetDebugEventsUntilDate() int64`

GetDebugEventsUntilDate returns the DebugEventsUntilDate field if non-nil, zero value otherwise.

### GetDebugEventsUntilDateOk

`func (o *GlobalFlagRepEnvironments) GetDebugEventsUntilDateOk() (*int64, bool)`

GetDebugEventsUntilDateOk returns a tuple with the DebugEventsUntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEventsUntilDate

`func (o *GlobalFlagRepEnvironments) SetDebugEventsUntilDate(v int64)`

SetDebugEventsUntilDate sets DebugEventsUntilDate field to given value.

### HasDebugEventsUntilDate

`func (o *GlobalFlagRepEnvironments) HasDebugEventsUntilDate() bool`

HasDebugEventsUntilDate returns a boolean if a field has been set.

### GetSummary

`func (o *GlobalFlagRepEnvironments) GetSummary() FlagSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GlobalFlagRepEnvironments) GetSummaryOk() (*FlagSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GlobalFlagRepEnvironments) SetSummary(v FlagSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GlobalFlagRepEnvironments) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


