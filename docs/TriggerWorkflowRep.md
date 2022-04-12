# TriggerWorkflowRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**MaintainerId** | Pointer to **string** |  | [optional] 
**Maintainer** | Pointer to [**MemberSummary**](MemberSummary.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IntegrationKey** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **[]interface{}** |  | [optional] 
**LastTriggeredAt** | Pointer to **int64** |  | [optional] 
**RecentTriggerBodies** | Pointer to [**[]RecentTriggerBody**](RecentTriggerBody.md) |  | [optional] 
**TriggerCount** | Pointer to **int32** |  | [optional] 
**TriggerURL** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewTriggerWorkflowRep

`func NewTriggerWorkflowRep() *TriggerWorkflowRep`

NewTriggerWorkflowRep instantiates a new TriggerWorkflowRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWorkflowRepWithDefaults

`func NewTriggerWorkflowRepWithDefaults() *TriggerWorkflowRep`

NewTriggerWorkflowRepWithDefaults instantiates a new TriggerWorkflowRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerWorkflowRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerWorkflowRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerWorkflowRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TriggerWorkflowRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *TriggerWorkflowRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TriggerWorkflowRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TriggerWorkflowRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TriggerWorkflowRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreationDate

`func (o *TriggerWorkflowRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *TriggerWorkflowRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *TriggerWorkflowRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *TriggerWorkflowRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetMaintainerId

`func (o *TriggerWorkflowRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *TriggerWorkflowRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *TriggerWorkflowRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *TriggerWorkflowRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainer

`func (o *TriggerWorkflowRep) GetMaintainer() MemberSummary`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *TriggerWorkflowRep) GetMaintainerOk() (*MemberSummary, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *TriggerWorkflowRep) SetMaintainer(v MemberSummary)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *TriggerWorkflowRep) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetEnabled

`func (o *TriggerWorkflowRep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TriggerWorkflowRep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TriggerWorkflowRep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TriggerWorkflowRep) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegrationKey

`func (o *TriggerWorkflowRep) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *TriggerWorkflowRep) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *TriggerWorkflowRep) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.

### HasIntegrationKey

`func (o *TriggerWorkflowRep) HasIntegrationKey() bool`

HasIntegrationKey returns a boolean if a field has been set.

### GetInstructions

`func (o *TriggerWorkflowRep) GetInstructions() []interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *TriggerWorkflowRep) GetInstructionsOk() (*[]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *TriggerWorkflowRep) SetInstructions(v []interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *TriggerWorkflowRep) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetLastTriggeredAt

`func (o *TriggerWorkflowRep) GetLastTriggeredAt() int64`

GetLastTriggeredAt returns the LastTriggeredAt field if non-nil, zero value otherwise.

### GetLastTriggeredAtOk

`func (o *TriggerWorkflowRep) GetLastTriggeredAtOk() (*int64, bool)`

GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredAt

`func (o *TriggerWorkflowRep) SetLastTriggeredAt(v int64)`

SetLastTriggeredAt sets LastTriggeredAt field to given value.

### HasLastTriggeredAt

`func (o *TriggerWorkflowRep) HasLastTriggeredAt() bool`

HasLastTriggeredAt returns a boolean if a field has been set.

### GetRecentTriggerBodies

`func (o *TriggerWorkflowRep) GetRecentTriggerBodies() []RecentTriggerBody`

GetRecentTriggerBodies returns the RecentTriggerBodies field if non-nil, zero value otherwise.

### GetRecentTriggerBodiesOk

`func (o *TriggerWorkflowRep) GetRecentTriggerBodiesOk() (*[]RecentTriggerBody, bool)`

GetRecentTriggerBodiesOk returns a tuple with the RecentTriggerBodies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTriggerBodies

`func (o *TriggerWorkflowRep) SetRecentTriggerBodies(v []RecentTriggerBody)`

SetRecentTriggerBodies sets RecentTriggerBodies field to given value.

### HasRecentTriggerBodies

`func (o *TriggerWorkflowRep) HasRecentTriggerBodies() bool`

HasRecentTriggerBodies returns a boolean if a field has been set.

### GetTriggerCount

`func (o *TriggerWorkflowRep) GetTriggerCount() int32`

GetTriggerCount returns the TriggerCount field if non-nil, zero value otherwise.

### GetTriggerCountOk

`func (o *TriggerWorkflowRep) GetTriggerCountOk() (*int32, bool)`

GetTriggerCountOk returns a tuple with the TriggerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCount

`func (o *TriggerWorkflowRep) SetTriggerCount(v int32)`

SetTriggerCount sets TriggerCount field to given value.

### HasTriggerCount

`func (o *TriggerWorkflowRep) HasTriggerCount() bool`

HasTriggerCount returns a boolean if a field has been set.

### GetTriggerURL

`func (o *TriggerWorkflowRep) GetTriggerURL() string`

GetTriggerURL returns the TriggerURL field if non-nil, zero value otherwise.

### GetTriggerURLOk

`func (o *TriggerWorkflowRep) GetTriggerURLOk() (*string, bool)`

GetTriggerURLOk returns a tuple with the TriggerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerURL

`func (o *TriggerWorkflowRep) SetTriggerURL(v string)`

SetTriggerURL sets TriggerURL field to given value.

### HasTriggerURL

`func (o *TriggerWorkflowRep) HasTriggerURL() bool`

HasTriggerURL returns a boolean if a field has been set.

### GetLinks

`func (o *TriggerWorkflowRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TriggerWorkflowRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TriggerWorkflowRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TriggerWorkflowRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


