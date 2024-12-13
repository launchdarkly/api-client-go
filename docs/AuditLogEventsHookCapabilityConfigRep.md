# AuditLogEventsHookCapabilityConfigRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statements** | Pointer to [**[]Statement**](Statement.md) | The set of resources you wish to subscribe to audit log notifications for. | [optional] 

## Methods

### NewAuditLogEventsHookCapabilityConfigRep

`func NewAuditLogEventsHookCapabilityConfigRep() *AuditLogEventsHookCapabilityConfigRep`

NewAuditLogEventsHookCapabilityConfigRep instantiates a new AuditLogEventsHookCapabilityConfigRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEventsHookCapabilityConfigRepWithDefaults

`func NewAuditLogEventsHookCapabilityConfigRepWithDefaults() *AuditLogEventsHookCapabilityConfigRep`

NewAuditLogEventsHookCapabilityConfigRepWithDefaults instantiates a new AuditLogEventsHookCapabilityConfigRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatements

`func (o *AuditLogEventsHookCapabilityConfigRep) GetStatements() []Statement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AuditLogEventsHookCapabilityConfigRep) GetStatementsOk() (*[]Statement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AuditLogEventsHookCapabilityConfigRep) SetStatements(v []Statement)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AuditLogEventsHookCapabilityConfigRep) HasStatements() bool`

HasStatements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


