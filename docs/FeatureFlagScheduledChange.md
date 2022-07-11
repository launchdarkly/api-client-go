# FeatureFlagScheduledChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreationDate** | **int64** |  | 
**MaintainerId** | **string** | The ID of the scheduled change maintainer | 
**Version** | **int32** | Version of the scheduled change | 
**ExecutionDate** | **int64** |  | 
**Instructions** | **[]map[string]interface{}** |  | 
**Conflicts** | Pointer to **interface{}** | Details on any conflicting scheduled changes | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | Links to other resources within the API. Includes the URL and content type of those resources. | [optional] 

## Methods

### NewFeatureFlagScheduledChange

`func NewFeatureFlagScheduledChange(id string, creationDate int64, maintainerId string, version int32, executionDate int64, instructions []map[string]interface{}, ) *FeatureFlagScheduledChange`

NewFeatureFlagScheduledChange instantiates a new FeatureFlagScheduledChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagScheduledChangeWithDefaults

`func NewFeatureFlagScheduledChangeWithDefaults() *FeatureFlagScheduledChange`

NewFeatureFlagScheduledChangeWithDefaults instantiates a new FeatureFlagScheduledChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeatureFlagScheduledChange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureFlagScheduledChange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureFlagScheduledChange) SetId(v string)`

SetId sets Id field to given value.


### GetCreationDate

`func (o *FeatureFlagScheduledChange) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *FeatureFlagScheduledChange) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *FeatureFlagScheduledChange) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetMaintainerId

`func (o *FeatureFlagScheduledChange) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *FeatureFlagScheduledChange) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *FeatureFlagScheduledChange) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.


### GetVersion

`func (o *FeatureFlagScheduledChange) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeatureFlagScheduledChange) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeatureFlagScheduledChange) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetExecutionDate

`func (o *FeatureFlagScheduledChange) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *FeatureFlagScheduledChange) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *FeatureFlagScheduledChange) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.


### GetInstructions

`func (o *FeatureFlagScheduledChange) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *FeatureFlagScheduledChange) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *FeatureFlagScheduledChange) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.


### GetConflicts

`func (o *FeatureFlagScheduledChange) GetConflicts() interface{}`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *FeatureFlagScheduledChange) GetConflictsOk() (*interface{}, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *FeatureFlagScheduledChange) SetConflicts(v interface{})`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *FeatureFlagScheduledChange) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### SetConflictsNil

`func (o *FeatureFlagScheduledChange) SetConflictsNil(b bool)`

 SetConflictsNil sets the value for Conflicts to be an explicit nil

### UnsetConflicts
`func (o *FeatureFlagScheduledChange) UnsetConflicts()`

UnsetConflicts ensures that no value is present for Conflicts, not even an explicit nil
### GetLinks

`func (o *FeatureFlagScheduledChange) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FeatureFlagScheduledChange) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FeatureFlagScheduledChange) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FeatureFlagScheduledChange) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


