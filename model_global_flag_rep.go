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

// GlobalFlagRep struct for GlobalFlagRep
type GlobalFlagRep struct {
	Name *string `json:"name,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Description *string `json:"description,omitempty"`
	Key *string `json:"key,omitempty"`
	Version *int32 `json:"_version,omitempty"`
	CreationDate *int64 `json:"creationDate,omitempty"`
	IncludeInSnippet *bool `json:"includeInSnippet,omitempty"`
	ClientSideAvailability *AccountsClientSideAvailability `json:"clientSideAvailability,omitempty"`
	Variations *[]VariateRep `json:"variations,omitempty"`
	VariationJsonSchema interface{} `json:"variationJsonSchema,omitempty"`
	Temporary *bool `json:"temporary,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Links *map[string]InlineResponse200 `json:"_links,omitempty"`
	MaintainerId *string `json:"maintainerId,omitempty"`
	Maintainer *MemberSummaryRep `json:"_maintainer,omitempty"`
	GoalIds *[]string `json:"goalIds,omitempty"`
	Experiments *ExperimentInfoRep `json:"experiments,omitempty"`
	CustomProperties *map[string]CustomProperty `json:"customProperties,omitempty"`
	Archived *bool `json:"archived,omitempty"`
	ArchivedDate *int64 `json:"archivedDate,omitempty"`
	Defaults *FlagDefaultsRep `json:"defaults,omitempty"`
	Environments *map[string]GlobalFlagRepEnvironments `json:"environments,omitempty"`
}

// NewGlobalFlagRep instantiates a new GlobalFlagRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalFlagRep() *GlobalFlagRep {
	this := GlobalFlagRep{}
	return &this
}

// NewGlobalFlagRepWithDefaults instantiates a new GlobalFlagRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalFlagRepWithDefaults() *GlobalFlagRep {
	this := GlobalFlagRep{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlobalFlagRep) SetName(v string) {
	o.Name = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *GlobalFlagRep) SetKind(v string) {
	o.Kind = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GlobalFlagRep) SetDescription(v string) {
	o.Description = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GlobalFlagRep) SetKey(v string) {
	o.Key = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *GlobalFlagRep) SetVersion(v int32) {
	o.Version = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetCreationDate() int64 {
	if o == nil || o.CreationDate == nil {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetCreationDateOk() (*int64, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *GlobalFlagRep) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetIncludeInSnippet returns the IncludeInSnippet field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetIncludeInSnippet() bool {
	if o == nil || o.IncludeInSnippet == nil {
		var ret bool
		return ret
	}
	return *o.IncludeInSnippet
}

// GetIncludeInSnippetOk returns a tuple with the IncludeInSnippet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetIncludeInSnippetOk() (*bool, bool) {
	if o == nil || o.IncludeInSnippet == nil {
		return nil, false
	}
	return o.IncludeInSnippet, true
}

// HasIncludeInSnippet returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasIncludeInSnippet() bool {
	if o != nil && o.IncludeInSnippet != nil {
		return true
	}

	return false
}

// SetIncludeInSnippet gets a reference to the given bool and assigns it to the IncludeInSnippet field.
func (o *GlobalFlagRep) SetIncludeInSnippet(v bool) {
	o.IncludeInSnippet = &v
}

// GetClientSideAvailability returns the ClientSideAvailability field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetClientSideAvailability() AccountsClientSideAvailability {
	if o == nil || o.ClientSideAvailability == nil {
		var ret AccountsClientSideAvailability
		return ret
	}
	return *o.ClientSideAvailability
}

// GetClientSideAvailabilityOk returns a tuple with the ClientSideAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetClientSideAvailabilityOk() (*AccountsClientSideAvailability, bool) {
	if o == nil || o.ClientSideAvailability == nil {
		return nil, false
	}
	return o.ClientSideAvailability, true
}

// HasClientSideAvailability returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasClientSideAvailability() bool {
	if o != nil && o.ClientSideAvailability != nil {
		return true
	}

	return false
}

// SetClientSideAvailability gets a reference to the given AccountsClientSideAvailability and assigns it to the ClientSideAvailability field.
func (o *GlobalFlagRep) SetClientSideAvailability(v AccountsClientSideAvailability) {
	o.ClientSideAvailability = &v
}

// GetVariations returns the Variations field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetVariations() []VariateRep {
	if o == nil || o.Variations == nil {
		var ret []VariateRep
		return ret
	}
	return *o.Variations
}

// GetVariationsOk returns a tuple with the Variations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetVariationsOk() (*[]VariateRep, bool) {
	if o == nil || o.Variations == nil {
		return nil, false
	}
	return o.Variations, true
}

// HasVariations returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasVariations() bool {
	if o != nil && o.Variations != nil {
		return true
	}

	return false
}

// SetVariations gets a reference to the given []VariateRep and assigns it to the Variations field.
func (o *GlobalFlagRep) SetVariations(v []VariateRep) {
	o.Variations = &v
}

// GetVariationJsonSchema returns the VariationJsonSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalFlagRep) GetVariationJsonSchema() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.VariationJsonSchema
}

// GetVariationJsonSchemaOk returns a tuple with the VariationJsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalFlagRep) GetVariationJsonSchemaOk() (*interface{}, bool) {
	if o == nil || o.VariationJsonSchema == nil {
		return nil, false
	}
	return &o.VariationJsonSchema, true
}

// HasVariationJsonSchema returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasVariationJsonSchema() bool {
	if o != nil && o.VariationJsonSchema != nil {
		return true
	}

	return false
}

// SetVariationJsonSchema gets a reference to the given interface{} and assigns it to the VariationJsonSchema field.
func (o *GlobalFlagRep) SetVariationJsonSchema(v interface{}) {
	o.VariationJsonSchema = v
}

// GetTemporary returns the Temporary field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetTemporary() bool {
	if o == nil || o.Temporary == nil {
		var ret bool
		return ret
	}
	return *o.Temporary
}

// GetTemporaryOk returns a tuple with the Temporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetTemporaryOk() (*bool, bool) {
	if o == nil || o.Temporary == nil {
		return nil, false
	}
	return o.Temporary, true
}

// HasTemporary returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasTemporary() bool {
	if o != nil && o.Temporary != nil {
		return true
	}

	return false
}

// SetTemporary gets a reference to the given bool and assigns it to the Temporary field.
func (o *GlobalFlagRep) SetTemporary(v bool) {
	o.Temporary = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GlobalFlagRep) SetTags(v []string) {
	o.Tags = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetLinks() map[string]InlineResponse200 {
	if o == nil || o.Links == nil {
		var ret map[string]InlineResponse200
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetLinksOk() (*map[string]InlineResponse200, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]InlineResponse200 and assigns it to the Links field.
func (o *GlobalFlagRep) SetLinks(v map[string]InlineResponse200) {
	o.Links = &v
}

// GetMaintainerId returns the MaintainerId field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetMaintainerId() string {
	if o == nil || o.MaintainerId == nil {
		var ret string
		return ret
	}
	return *o.MaintainerId
}

// GetMaintainerIdOk returns a tuple with the MaintainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetMaintainerIdOk() (*string, bool) {
	if o == nil || o.MaintainerId == nil {
		return nil, false
	}
	return o.MaintainerId, true
}

// HasMaintainerId returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasMaintainerId() bool {
	if o != nil && o.MaintainerId != nil {
		return true
	}

	return false
}

// SetMaintainerId gets a reference to the given string and assigns it to the MaintainerId field.
func (o *GlobalFlagRep) SetMaintainerId(v string) {
	o.MaintainerId = &v
}

// GetMaintainer returns the Maintainer field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetMaintainer() MemberSummaryRep {
	if o == nil || o.Maintainer == nil {
		var ret MemberSummaryRep
		return ret
	}
	return *o.Maintainer
}

// GetMaintainerOk returns a tuple with the Maintainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetMaintainerOk() (*MemberSummaryRep, bool) {
	if o == nil || o.Maintainer == nil {
		return nil, false
	}
	return o.Maintainer, true
}

// HasMaintainer returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasMaintainer() bool {
	if o != nil && o.Maintainer != nil {
		return true
	}

	return false
}

// SetMaintainer gets a reference to the given MemberSummaryRep and assigns it to the Maintainer field.
func (o *GlobalFlagRep) SetMaintainer(v MemberSummaryRep) {
	o.Maintainer = &v
}

// GetGoalIds returns the GoalIds field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetGoalIds() []string {
	if o == nil || o.GoalIds == nil {
		var ret []string
		return ret
	}
	return *o.GoalIds
}

// GetGoalIdsOk returns a tuple with the GoalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetGoalIdsOk() (*[]string, bool) {
	if o == nil || o.GoalIds == nil {
		return nil, false
	}
	return o.GoalIds, true
}

// HasGoalIds returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasGoalIds() bool {
	if o != nil && o.GoalIds != nil {
		return true
	}

	return false
}

// SetGoalIds gets a reference to the given []string and assigns it to the GoalIds field.
func (o *GlobalFlagRep) SetGoalIds(v []string) {
	o.GoalIds = &v
}

// GetExperiments returns the Experiments field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetExperiments() ExperimentInfoRep {
	if o == nil || o.Experiments == nil {
		var ret ExperimentInfoRep
		return ret
	}
	return *o.Experiments
}

// GetExperimentsOk returns a tuple with the Experiments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetExperimentsOk() (*ExperimentInfoRep, bool) {
	if o == nil || o.Experiments == nil {
		return nil, false
	}
	return o.Experiments, true
}

// HasExperiments returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasExperiments() bool {
	if o != nil && o.Experiments != nil {
		return true
	}

	return false
}

// SetExperiments gets a reference to the given ExperimentInfoRep and assigns it to the Experiments field.
func (o *GlobalFlagRep) SetExperiments(v ExperimentInfoRep) {
	o.Experiments = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetCustomProperties() map[string]CustomProperty {
	if o == nil || o.CustomProperties == nil {
		var ret map[string]CustomProperty
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetCustomPropertiesOk() (*map[string]CustomProperty, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]CustomProperty and assigns it to the CustomProperties field.
func (o *GlobalFlagRep) SetCustomProperties(v map[string]CustomProperty) {
	o.CustomProperties = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *GlobalFlagRep) SetArchived(v bool) {
	o.Archived = &v
}

// GetArchivedDate returns the ArchivedDate field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetArchivedDate() int64 {
	if o == nil || o.ArchivedDate == nil {
		var ret int64
		return ret
	}
	return *o.ArchivedDate
}

// GetArchivedDateOk returns a tuple with the ArchivedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetArchivedDateOk() (*int64, bool) {
	if o == nil || o.ArchivedDate == nil {
		return nil, false
	}
	return o.ArchivedDate, true
}

// HasArchivedDate returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasArchivedDate() bool {
	if o != nil && o.ArchivedDate != nil {
		return true
	}

	return false
}

// SetArchivedDate gets a reference to the given int64 and assigns it to the ArchivedDate field.
func (o *GlobalFlagRep) SetArchivedDate(v int64) {
	o.ArchivedDate = &v
}

// GetDefaults returns the Defaults field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetDefaults() FlagDefaultsRep {
	if o == nil || o.Defaults == nil {
		var ret FlagDefaultsRep
		return ret
	}
	return *o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetDefaultsOk() (*FlagDefaultsRep, bool) {
	if o == nil || o.Defaults == nil {
		return nil, false
	}
	return o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasDefaults() bool {
	if o != nil && o.Defaults != nil {
		return true
	}

	return false
}

// SetDefaults gets a reference to the given FlagDefaultsRep and assigns it to the Defaults field.
func (o *GlobalFlagRep) SetDefaults(v FlagDefaultsRep) {
	o.Defaults = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *GlobalFlagRep) GetEnvironments() map[string]GlobalFlagRepEnvironments {
	if o == nil || o.Environments == nil {
		var ret map[string]GlobalFlagRepEnvironments
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalFlagRep) GetEnvironmentsOk() (*map[string]GlobalFlagRepEnvironments, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *GlobalFlagRep) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given map[string]GlobalFlagRepEnvironments and assigns it to the Environments field.
func (o *GlobalFlagRep) SetEnvironments(v map[string]GlobalFlagRepEnvironments) {
	o.Environments = &v
}

func (o GlobalFlagRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Version != nil {
		toSerialize["_version"] = o.Version
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.IncludeInSnippet != nil {
		toSerialize["includeInSnippet"] = o.IncludeInSnippet
	}
	if o.ClientSideAvailability != nil {
		toSerialize["clientSideAvailability"] = o.ClientSideAvailability
	}
	if o.Variations != nil {
		toSerialize["variations"] = o.Variations
	}
	if o.VariationJsonSchema != nil {
		toSerialize["variationJsonSchema"] = o.VariationJsonSchema
	}
	if o.Temporary != nil {
		toSerialize["temporary"] = o.Temporary
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.MaintainerId != nil {
		toSerialize["maintainerId"] = o.MaintainerId
	}
	if o.Maintainer != nil {
		toSerialize["_maintainer"] = o.Maintainer
	}
	if o.GoalIds != nil {
		toSerialize["goalIds"] = o.GoalIds
	}
	if o.Experiments != nil {
		toSerialize["experiments"] = o.Experiments
	}
	if o.CustomProperties != nil {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.ArchivedDate != nil {
		toSerialize["archivedDate"] = o.ArchivedDate
	}
	if o.Defaults != nil {
		toSerialize["defaults"] = o.Defaults
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalFlagRep struct {
	value *GlobalFlagRep
	isSet bool
}

func (v NullableGlobalFlagRep) Get() *GlobalFlagRep {
	return v.value
}

func (v *NullableGlobalFlagRep) Set(val *GlobalFlagRep) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalFlagRep) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalFlagRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalFlagRep(val *GlobalFlagRep) *NullableGlobalFlagRep {
	return &NullableGlobalFlagRep{value: val, isSet: true}
}

func (v NullableGlobalFlagRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalFlagRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


