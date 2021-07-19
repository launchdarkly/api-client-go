/*
 * LaunchDarkly REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0
 * Contact: support@launchdarkly.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ldapi

import (
	"encoding/json"
)

// GlobalFlagRepEnvironments struct for GlobalFlagRepEnvironments
type GlobalFlagRepEnvironments struct {
	On *bool `json:"on,omitempty"`
	Archived *bool `json:"archived,omitempty"`
	Salt *string `json:"salt,omitempty"`
	Sel *string `json:"sel,omitempty"`
	LastModified *int64 `json:"lastModified,omitempty"`
	Version *int32 `json:"version,omitempty"`
	Targets *[]FlagConfigurationRepTargets `json:"targets,omitempty"`
	Rules *[]FlagConfigurationRepRules `json:"rules,omitempty"`
	Fallthrough *VariationOrRolloutRep `json:"fallthrough,omitempty"`
	OffVariation *int32 `json:"offVariation,omitempty"`
	Prerequisites *[]FlagConfigurationRepPrerequisites `json:"prerequisites,omitempty"`
	Site *CoreLink `json:"_site,omitempty"`
	Access *AccessRep `json:"_access,omitempty"`
	EnvironmentName *string `json:"_environmentName,omitempty"`
	TrackEvents *bool `json:"trackEvents,omitempty"`
	TrackEventsFallthrough *bool `json:"trackEventsFallthrough,omitempty"`
	DebugEventsUntilDate *int64 `json:"_debugEventsUntilDate,omitempty"`
	Summary *FlagSummary `json:"_summary,omitempty"`
}

// NewGlobalFlagRepEnvironments instantiates a new GlobalFlagRepEnvironments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalFlagRepEnvironments() *GlobalFlagRepEnvironments {
	this := GlobalFlagRepEnvironments{}
	return &this
}

// NewGlobalFlagRepEnvironmentsWithDefaults instantiates a new GlobalFlagRepEnvironments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalFlagRepEnvironmentsWithDefaults() *GlobalFlagRepEnvironments {
	this := GlobalFlagRepEnvironments{}
	return &this
}

// GetOn returns the On field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetOn() bool {
	if o == nil || o.On == nil {
		var ret bool
		return ret
	}
	return *o.On
}

// GetOnOk returns a tuple with the On field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetOnOk() (*bool, bool) {
	if o == nil || o.On == nil {
		return nil, false
	}
	return o.On, true
}

// HasOn returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasOn() bool {
	if o != nil && o.On != nil {
		return true
	}

	return false
}

// SetOn gets a reference to the given bool and assigns it to the On field.
func (o *GlobalFlagRepEnvironments) SetOn(v bool) {
	o.On = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *GlobalFlagRepEnvironments) SetArchived(v bool) {
	o.Archived = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetSalt() string {
	if o == nil || o.Salt == nil {
		var ret string
		return ret
	}
	return *o.Salt
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetSaltOk() (*string, bool) {
	if o == nil || o.Salt == nil {
		return nil, false
	}
	return o.Salt, true
}

// HasSalt returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasSalt() bool {
	if o != nil && o.Salt != nil {
		return true
	}

	return false
}

// SetSalt gets a reference to the given string and assigns it to the Salt field.
func (o *GlobalFlagRepEnvironments) SetSalt(v string) {
	o.Salt = &v
}

// GetSel returns the Sel field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetSel() string {
	if o == nil || o.Sel == nil {
		var ret string
		return ret
	}
	return *o.Sel
}

// GetSelOk returns a tuple with the Sel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetSelOk() (*string, bool) {
	if o == nil || o.Sel == nil {
		return nil, false
	}
	return o.Sel, true
}

// HasSel returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasSel() bool {
	if o != nil && o.Sel != nil {
		return true
	}

	return false
}

// SetSel gets a reference to the given string and assigns it to the Sel field.
func (o *GlobalFlagRepEnvironments) SetSel(v string) {
	o.Sel = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetLastModified() int64 {
	if o == nil || o.LastModified == nil {
		var ret int64
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetLastModifiedOk() (*int64, bool) {
	if o == nil || o.LastModified == nil {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given int64 and assigns it to the LastModified field.
func (o *GlobalFlagRepEnvironments) SetLastModified(v int64) {
	o.LastModified = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *GlobalFlagRepEnvironments) SetVersion(v int32) {
	o.Version = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetTargets() []FlagConfigurationRepTargets {
	if o == nil || o.Targets == nil {
		var ret []FlagConfigurationRepTargets
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetTargetsOk() (*[]FlagConfigurationRepTargets, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []FlagConfigurationRepTargets and assigns it to the Targets field.
func (o *GlobalFlagRepEnvironments) SetTargets(v []FlagConfigurationRepTargets) {
	o.Targets = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetRules() []FlagConfigurationRepRules {
	if o == nil || o.Rules == nil {
		var ret []FlagConfigurationRepRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetRulesOk() (*[]FlagConfigurationRepRules, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []FlagConfigurationRepRules and assigns it to the Rules field.
func (o *GlobalFlagRepEnvironments) SetRules(v []FlagConfigurationRepRules) {
	o.Rules = &v
}

// GetFallthrough returns the Fallthrough field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetFallthrough() VariationOrRolloutRep {
	if o == nil || o.Fallthrough == nil {
		var ret VariationOrRolloutRep
		return ret
	}
	return *o.Fallthrough
}

// GetFallthroughOk returns a tuple with the Fallthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetFallthroughOk() (*VariationOrRolloutRep, bool) {
	if o == nil || o.Fallthrough == nil {
		return nil, false
	}
	return o.Fallthrough, true
}

// HasFallthrough returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasFallthrough() bool {
	if o != nil && o.Fallthrough != nil {
		return true
	}

	return false
}

// SetFallthrough gets a reference to the given VariationOrRolloutRep and assigns it to the Fallthrough field.
func (o *GlobalFlagRepEnvironments) SetFallthrough(v VariationOrRolloutRep) {
	o.Fallthrough = &v
}

// GetOffVariation returns the OffVariation field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetOffVariation() int32 {
	if o == nil || o.OffVariation == nil {
		var ret int32
		return ret
	}
	return *o.OffVariation
}

// GetOffVariationOk returns a tuple with the OffVariation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetOffVariationOk() (*int32, bool) {
	if o == nil || o.OffVariation == nil {
		return nil, false
	}
	return o.OffVariation, true
}

// HasOffVariation returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasOffVariation() bool {
	if o != nil && o.OffVariation != nil {
		return true
	}

	return false
}

// SetOffVariation gets a reference to the given int32 and assigns it to the OffVariation field.
func (o *GlobalFlagRepEnvironments) SetOffVariation(v int32) {
	o.OffVariation = &v
}

// GetPrerequisites returns the Prerequisites field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetPrerequisites() []FlagConfigurationRepPrerequisites {
	if o == nil || o.Prerequisites == nil {
		var ret []FlagConfigurationRepPrerequisites
		return ret
	}
	return *o.Prerequisites
}

// GetPrerequisitesOk returns a tuple with the Prerequisites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetPrerequisitesOk() (*[]FlagConfigurationRepPrerequisites, bool) {
	if o == nil || o.Prerequisites == nil {
		return nil, false
	}
	return o.Prerequisites, true
}

// HasPrerequisites returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasPrerequisites() bool {
	if o != nil && o.Prerequisites != nil {
		return true
	}

	return false
}

// SetPrerequisites gets a reference to the given []FlagConfigurationRepPrerequisites and assigns it to the Prerequisites field.
func (o *GlobalFlagRepEnvironments) SetPrerequisites(v []FlagConfigurationRepPrerequisites) {
	o.Prerequisites = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetSite() CoreLink {
	if o == nil || o.Site == nil {
		var ret CoreLink
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetSiteOk() (*CoreLink, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given CoreLink and assigns it to the Site field.
func (o *GlobalFlagRepEnvironments) SetSite(v CoreLink) {
	o.Site = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetAccess() AccessRep {
	if o == nil || o.Access == nil {
		var ret AccessRep
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetAccessOk() (*AccessRep, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AccessRep and assigns it to the Access field.
func (o *GlobalFlagRepEnvironments) SetAccess(v AccessRep) {
	o.Access = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasEnvironmentName() bool {
	if o != nil && o.EnvironmentName != nil {
		return true
	}

	return false
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *GlobalFlagRepEnvironments) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetTrackEvents returns the TrackEvents field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetTrackEvents() bool {
	if o == nil || o.TrackEvents == nil {
		var ret bool
		return ret
	}
	return *o.TrackEvents
}

// GetTrackEventsOk returns a tuple with the TrackEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetTrackEventsOk() (*bool, bool) {
	if o == nil || o.TrackEvents == nil {
		return nil, false
	}
	return o.TrackEvents, true
}

// HasTrackEvents returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasTrackEvents() bool {
	if o != nil && o.TrackEvents != nil {
		return true
	}

	return false
}

// SetTrackEvents gets a reference to the given bool and assigns it to the TrackEvents field.
func (o *GlobalFlagRepEnvironments) SetTrackEvents(v bool) {
	o.TrackEvents = &v
}

// GetTrackEventsFallthrough returns the TrackEventsFallthrough field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetTrackEventsFallthrough() bool {
	if o == nil || o.TrackEventsFallthrough == nil {
		var ret bool
		return ret
	}
	return *o.TrackEventsFallthrough
}

// GetTrackEventsFallthroughOk returns a tuple with the TrackEventsFallthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetTrackEventsFallthroughOk() (*bool, bool) {
	if o == nil || o.TrackEventsFallthrough == nil {
		return nil, false
	}
	return o.TrackEventsFallthrough, true
}

// HasTrackEventsFallthrough returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasTrackEventsFallthrough() bool {
	if o != nil && o.TrackEventsFallthrough != nil {
		return true
	}

	return false
}

// SetTrackEventsFallthrough gets a reference to the given bool and assigns it to the TrackEventsFallthrough field.
func (o *GlobalFlagRepEnvironments) SetTrackEventsFallthrough(v bool) {
	o.TrackEventsFallthrough = &v
}

// GetDebugEventsUntilDate returns the DebugEventsUntilDate field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetDebugEventsUntilDate() int64 {
	if o == nil || o.DebugEventsUntilDate == nil {
		var ret int64
		return ret
	}
	return *o.DebugEventsUntilDate
}

// GetDebugEventsUntilDateOk returns a tuple with the DebugEventsUntilDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetDebugEventsUntilDateOk() (*int64, bool) {
	if o == nil || o.DebugEventsUntilDate == nil {
		return nil, false
	}
	return o.DebugEventsUntilDate, true
}

// HasDebugEventsUntilDate returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasDebugEventsUntilDate() bool {
	if o != nil && o.DebugEventsUntilDate != nil {
		return true
	}

	return false
}

// SetDebugEventsUntilDate gets a reference to the given int64 and assigns it to the DebugEventsUntilDate field.
func (o *GlobalFlagRepEnvironments) SetDebugEventsUntilDate(v int64) {
	o.DebugEventsUntilDate = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *GlobalFlagRepEnvironments) GetSummary() FlagSummary {
	if o == nil || o.Summary == nil {
		var ret FlagSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRepEnvironments) GetSummaryOk() (*FlagSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *GlobalFlagRepEnvironments) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given FlagSummary and assigns it to the Summary field.
func (o *GlobalFlagRepEnvironments) SetSummary(v FlagSummary) {
	o.Summary = &v
}

func (o GlobalFlagRepEnvironments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.On != nil {
		toSerialize["on"] = o.On
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.Salt != nil {
		toSerialize["salt"] = o.Salt
	}
	if o.Sel != nil {
		toSerialize["sel"] = o.Sel
	}
	if o.LastModified != nil {
		toSerialize["lastModified"] = o.LastModified
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Fallthrough != nil {
		toSerialize["fallthrough"] = o.Fallthrough
	}
	if o.OffVariation != nil {
		toSerialize["offVariation"] = o.OffVariation
	}
	if o.Prerequisites != nil {
		toSerialize["prerequisites"] = o.Prerequisites
	}
	if o.Site != nil {
		toSerialize["_site"] = o.Site
	}
	if o.Access != nil {
		toSerialize["_access"] = o.Access
	}
	if o.EnvironmentName != nil {
		toSerialize["_environmentName"] = o.EnvironmentName
	}
	if o.TrackEvents != nil {
		toSerialize["trackEvents"] = o.TrackEvents
	}
	if o.TrackEventsFallthrough != nil {
		toSerialize["trackEventsFallthrough"] = o.TrackEventsFallthrough
	}
	if o.DebugEventsUntilDate != nil {
		toSerialize["_debugEventsUntilDate"] = o.DebugEventsUntilDate
	}
	if o.Summary != nil {
		toSerialize["_summary"] = o.Summary
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalFlagRepEnvironments struct {
	value *GlobalFlagRepEnvironments
	isSet bool
}

func (v NullableGlobalFlagRepEnvironments) Get() *GlobalFlagRepEnvironments {
	return v.value
}

func (v *NullableGlobalFlagRepEnvironments) Set(val *GlobalFlagRepEnvironments) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalFlagRepEnvironments) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalFlagRepEnvironments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalFlagRepEnvironments(val *GlobalFlagRepEnvironments) *NullableGlobalFlagRepEnvironments {
	return &NullableGlobalFlagRepEnvironments{value: val, isSet: true}
}

func (v NullableGlobalFlagRepEnvironments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalFlagRepEnvironments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


