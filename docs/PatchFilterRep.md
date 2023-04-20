# PatchFilterRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The filter name | [optional] 
**Description** | Pointer to **string** | Textual description of the filter | [optional] 
**Rules** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** | True if the filter&#39;s rules are active | [optional] 
**Archived** | Pointer to **bool** | True if filter is archived | [optional] 

## Methods

### NewPatchFilterRep

`func NewPatchFilterRep() *PatchFilterRep`

NewPatchFilterRep instantiates a new PatchFilterRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchFilterRepWithDefaults

`func NewPatchFilterRepWithDefaults() *PatchFilterRep`

NewPatchFilterRepWithDefaults instantiates a new PatchFilterRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchFilterRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchFilterRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchFilterRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchFilterRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchFilterRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchFilterRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchFilterRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchFilterRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *PatchFilterRep) GetRules() []map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PatchFilterRep) GetRulesOk() (*[]map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PatchFilterRep) SetRules(v []map[string]interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *PatchFilterRep) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchFilterRep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchFilterRep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchFilterRep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchFilterRep) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetArchived

`func (o *PatchFilterRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PatchFilterRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PatchFilterRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PatchFilterRep) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


