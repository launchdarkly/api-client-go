# FlagEventRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The flag event ID | 
**ProjectId** | **string** | The project ID | 
**ProjectKey** | **string** | The project key | 
**EnvironmentId** | Pointer to **string** | The environment ID | [optional] 
**EnvironmentKey** | Pointer to **string** | The environment key | [optional] 
**FlagKey** | **string** | The flag key | 
**EventType** | **string** |  | 
**EventTime** | **int64** |  | 
**Description** | **string** | The event description | 
**AuditLogEntryId** | Pointer to **string** | The audit log entry ID | [optional] 
**Member** | Pointer to [**FlagEventMemberRep**](FlagEventMemberRep.md) |  | [optional] 
**Actions** | Pointer to **[]string** | The resource actions | [optional] 
**Impact** | [**FlagEventImpactRep**](FlagEventImpactRep.md) |  | 
**Experiments** | Pointer to [**FlagEventExperimentCollection**](FlagEventExperimentCollection.md) |  | [optional] 

## Methods

### NewFlagEventRep

`func NewFlagEventRep(id string, projectId string, projectKey string, flagKey string, eventType string, eventTime int64, description string, impact FlagEventImpactRep, ) *FlagEventRep`

NewFlagEventRep instantiates a new FlagEventRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagEventRepWithDefaults

`func NewFlagEventRepWithDefaults() *FlagEventRep`

NewFlagEventRepWithDefaults instantiates a new FlagEventRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlagEventRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagEventRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagEventRep) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *FlagEventRep) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FlagEventRep) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FlagEventRep) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProjectKey

`func (o *FlagEventRep) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *FlagEventRep) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *FlagEventRep) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetEnvironmentId

`func (o *FlagEventRep) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FlagEventRep) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FlagEventRep) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *FlagEventRep) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *FlagEventRep) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *FlagEventRep) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *FlagEventRep) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *FlagEventRep) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetFlagKey

`func (o *FlagEventRep) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *FlagEventRep) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *FlagEventRep) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetEventType

`func (o *FlagEventRep) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *FlagEventRep) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *FlagEventRep) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventTime

`func (o *FlagEventRep) GetEventTime() int64`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *FlagEventRep) GetEventTimeOk() (*int64, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *FlagEventRep) SetEventTime(v int64)`

SetEventTime sets EventTime field to given value.


### GetDescription

`func (o *FlagEventRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagEventRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagEventRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAuditLogEntryId

`func (o *FlagEventRep) GetAuditLogEntryId() string`

GetAuditLogEntryId returns the AuditLogEntryId field if non-nil, zero value otherwise.

### GetAuditLogEntryIdOk

`func (o *FlagEventRep) GetAuditLogEntryIdOk() (*string, bool)`

GetAuditLogEntryIdOk returns a tuple with the AuditLogEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogEntryId

`func (o *FlagEventRep) SetAuditLogEntryId(v string)`

SetAuditLogEntryId sets AuditLogEntryId field to given value.

### HasAuditLogEntryId

`func (o *FlagEventRep) HasAuditLogEntryId() bool`

HasAuditLogEntryId returns a boolean if a field has been set.

### GetMember

`func (o *FlagEventRep) GetMember() FlagEventMemberRep`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *FlagEventRep) GetMemberOk() (*FlagEventMemberRep, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *FlagEventRep) SetMember(v FlagEventMemberRep)`

SetMember sets Member field to given value.

### HasMember

`func (o *FlagEventRep) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetActions

`func (o *FlagEventRep) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *FlagEventRep) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *FlagEventRep) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *FlagEventRep) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetImpact

`func (o *FlagEventRep) GetImpact() FlagEventImpactRep`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *FlagEventRep) GetImpactOk() (*FlagEventImpactRep, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *FlagEventRep) SetImpact(v FlagEventImpactRep)`

SetImpact sets Impact field to given value.


### GetExperiments

`func (o *FlagEventRep) GetExperiments() FlagEventExperimentCollection`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *FlagEventRep) GetExperimentsOk() (*FlagEventExperimentCollection, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *FlagEventRep) SetExperiments(v FlagEventExperimentCollection)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *FlagEventRep) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


