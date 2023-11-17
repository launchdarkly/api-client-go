# MigrationSafetyIssueRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CausingRuleId** | Pointer to **string** | The ID of the rule which caused this issue | [optional] 
**AffectedRuleIds** | Pointer to **[]string** | A list of the IDs of the rules which are affected by this issue. &lt;code&gt;fallthrough&lt;/code&gt; is a sentinel value for the default rule. | [optional] 
**Issue** | Pointer to **string** | A description of the issue that &lt;code&gt;causingRuleId&lt;/code&gt; has caused for &lt;code&gt;affectedRuleIds&lt;/code&gt;. | [optional] 
**OldSystemAffected** | Pointer to **bool** | Whether the changes caused by &lt;code&gt;causingRuleId&lt;/code&gt; bring inconsistency to the old system | [optional] 

## Methods

### NewMigrationSafetyIssueRep

`func NewMigrationSafetyIssueRep() *MigrationSafetyIssueRep`

NewMigrationSafetyIssueRep instantiates a new MigrationSafetyIssueRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationSafetyIssueRepWithDefaults

`func NewMigrationSafetyIssueRepWithDefaults() *MigrationSafetyIssueRep`

NewMigrationSafetyIssueRepWithDefaults instantiates a new MigrationSafetyIssueRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCausingRuleId

`func (o *MigrationSafetyIssueRep) GetCausingRuleId() string`

GetCausingRuleId returns the CausingRuleId field if non-nil, zero value otherwise.

### GetCausingRuleIdOk

`func (o *MigrationSafetyIssueRep) GetCausingRuleIdOk() (*string, bool)`

GetCausingRuleIdOk returns a tuple with the CausingRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausingRuleId

`func (o *MigrationSafetyIssueRep) SetCausingRuleId(v string)`

SetCausingRuleId sets CausingRuleId field to given value.

### HasCausingRuleId

`func (o *MigrationSafetyIssueRep) HasCausingRuleId() bool`

HasCausingRuleId returns a boolean if a field has been set.

### GetAffectedRuleIds

`func (o *MigrationSafetyIssueRep) GetAffectedRuleIds() []string`

GetAffectedRuleIds returns the AffectedRuleIds field if non-nil, zero value otherwise.

### GetAffectedRuleIdsOk

`func (o *MigrationSafetyIssueRep) GetAffectedRuleIdsOk() (*[]string, bool)`

GetAffectedRuleIdsOk returns a tuple with the AffectedRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedRuleIds

`func (o *MigrationSafetyIssueRep) SetAffectedRuleIds(v []string)`

SetAffectedRuleIds sets AffectedRuleIds field to given value.

### HasAffectedRuleIds

`func (o *MigrationSafetyIssueRep) HasAffectedRuleIds() bool`

HasAffectedRuleIds returns a boolean if a field has been set.

### GetIssue

`func (o *MigrationSafetyIssueRep) GetIssue() string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *MigrationSafetyIssueRep) GetIssueOk() (*string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *MigrationSafetyIssueRep) SetIssue(v string)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *MigrationSafetyIssueRep) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetOldSystemAffected

`func (o *MigrationSafetyIssueRep) GetOldSystemAffected() bool`

GetOldSystemAffected returns the OldSystemAffected field if non-nil, zero value otherwise.

### GetOldSystemAffectedOk

`func (o *MigrationSafetyIssueRep) GetOldSystemAffectedOk() (*bool, bool)`

GetOldSystemAffectedOk returns a tuple with the OldSystemAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSystemAffected

`func (o *MigrationSafetyIssueRep) SetOldSystemAffected(v bool)`

SetOldSystemAffected sets OldSystemAffected field to given value.

### HasOldSystemAffected

`func (o *MigrationSafetyIssueRep) HasOldSystemAffected() bool`

HasOldSystemAffected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


