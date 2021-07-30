# AuditLogEntryListingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to [**SubjectDataRep**](SubjectDataRep.md) |  | [optional] 
**Member** | Pointer to [**MemberDataRep**](MemberDataRep.md) |  | [optional] 
**Token** | Pointer to [**TokenDataRep**](TokenDataRep.md) |  | [optional] 
**App** | Pointer to [**AuthorizedAppDataRep**](AuthorizedAppDataRep.md) |  | [optional] 
**TitleVerb** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**TargetResourceRep**](TargetResourceRep.md) |  | [optional] 
**Parent** | Pointer to [**ParentResourceRep**](ParentResourceRep.md) |  | [optional] 
**Delta** | Pointer to **interface{}** |  | [optional] 
**TriggerBody** | Pointer to **interface{}** |  | [optional] 
**Merge** | Pointer to **interface{}** |  | [optional] 
**PreviousVersion** | Pointer to **interface{}** |  | [optional] 
**CurrentVersion** | Pointer to **interface{}** |  | [optional] 
**Subentries** | Pointer to [**[]AuditLogEntryListingRep**](AuditLogEntryListingRep.md) |  | [optional] 

## Methods

### NewAuditLogEntryListingRep

`func NewAuditLogEntryListingRep() *AuditLogEntryListingRep`

NewAuditLogEntryListingRep instantiates a new AuditLogEntryListingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEntryListingRepWithDefaults

`func NewAuditLogEntryListingRepWithDefaults() *AuditLogEntryListingRep`

NewAuditLogEntryListingRepWithDefaults instantiates a new AuditLogEntryListingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *AuditLogEntryListingRep) GetSubject() SubjectDataRep`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuditLogEntryListingRep) GetSubjectOk() (*SubjectDataRep, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuditLogEntryListingRep) SetSubject(v SubjectDataRep)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AuditLogEntryListingRep) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMember

`func (o *AuditLogEntryListingRep) GetMember() MemberDataRep`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *AuditLogEntryListingRep) GetMemberOk() (*MemberDataRep, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *AuditLogEntryListingRep) SetMember(v MemberDataRep)`

SetMember sets Member field to given value.

### HasMember

`func (o *AuditLogEntryListingRep) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetToken

`func (o *AuditLogEntryListingRep) GetToken() TokenDataRep`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuditLogEntryListingRep) GetTokenOk() (*TokenDataRep, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuditLogEntryListingRep) SetToken(v TokenDataRep)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuditLogEntryListingRep) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetApp

`func (o *AuditLogEntryListingRep) GetApp() AuthorizedAppDataRep`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AuditLogEntryListingRep) GetAppOk() (*AuthorizedAppDataRep, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AuditLogEntryListingRep) SetApp(v AuthorizedAppDataRep)`

SetApp sets App field to given value.

### HasApp

`func (o *AuditLogEntryListingRep) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetTitleVerb

`func (o *AuditLogEntryListingRep) GetTitleVerb() string`

GetTitleVerb returns the TitleVerb field if non-nil, zero value otherwise.

### GetTitleVerbOk

`func (o *AuditLogEntryListingRep) GetTitleVerbOk() (*string, bool)`

GetTitleVerbOk returns a tuple with the TitleVerb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleVerb

`func (o *AuditLogEntryListingRep) SetTitleVerb(v string)`

SetTitleVerb sets TitleVerb field to given value.

### HasTitleVerb

`func (o *AuditLogEntryListingRep) HasTitleVerb() bool`

HasTitleVerb returns a boolean if a field has been set.

### GetTitle

`func (o *AuditLogEntryListingRep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuditLogEntryListingRep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuditLogEntryListingRep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuditLogEntryListingRep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTarget

`func (o *AuditLogEntryListingRep) GetTarget() TargetResourceRep`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AuditLogEntryListingRep) GetTargetOk() (*TargetResourceRep, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AuditLogEntryListingRep) SetTarget(v TargetResourceRep)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AuditLogEntryListingRep) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetParent

`func (o *AuditLogEntryListingRep) GetParent() ParentResourceRep`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuditLogEntryListingRep) GetParentOk() (*ParentResourceRep, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuditLogEntryListingRep) SetParent(v ParentResourceRep)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuditLogEntryListingRep) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetDelta

`func (o *AuditLogEntryListingRep) GetDelta() interface{}`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *AuditLogEntryListingRep) GetDeltaOk() (*interface{}, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *AuditLogEntryListingRep) SetDelta(v interface{})`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *AuditLogEntryListingRep) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### SetDeltaNil

`func (o *AuditLogEntryListingRep) SetDeltaNil(b bool)`

 SetDeltaNil sets the value for Delta to be an explicit nil

### UnsetDelta
`func (o *AuditLogEntryListingRep) UnsetDelta()`

UnsetDelta ensures that no value is present for Delta, not even an explicit nil
### GetTriggerBody

`func (o *AuditLogEntryListingRep) GetTriggerBody() interface{}`

GetTriggerBody returns the TriggerBody field if non-nil, zero value otherwise.

### GetTriggerBodyOk

`func (o *AuditLogEntryListingRep) GetTriggerBodyOk() (*interface{}, bool)`

GetTriggerBodyOk returns a tuple with the TriggerBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerBody

`func (o *AuditLogEntryListingRep) SetTriggerBody(v interface{})`

SetTriggerBody sets TriggerBody field to given value.

### HasTriggerBody

`func (o *AuditLogEntryListingRep) HasTriggerBody() bool`

HasTriggerBody returns a boolean if a field has been set.

### SetTriggerBodyNil

`func (o *AuditLogEntryListingRep) SetTriggerBodyNil(b bool)`

 SetTriggerBodyNil sets the value for TriggerBody to be an explicit nil

### UnsetTriggerBody
`func (o *AuditLogEntryListingRep) UnsetTriggerBody()`

UnsetTriggerBody ensures that no value is present for TriggerBody, not even an explicit nil
### GetMerge

`func (o *AuditLogEntryListingRep) GetMerge() interface{}`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *AuditLogEntryListingRep) GetMergeOk() (*interface{}, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *AuditLogEntryListingRep) SetMerge(v interface{})`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *AuditLogEntryListingRep) HasMerge() bool`

HasMerge returns a boolean if a field has been set.

### SetMergeNil

`func (o *AuditLogEntryListingRep) SetMergeNil(b bool)`

 SetMergeNil sets the value for Merge to be an explicit nil

### UnsetMerge
`func (o *AuditLogEntryListingRep) UnsetMerge()`

UnsetMerge ensures that no value is present for Merge, not even an explicit nil
### GetPreviousVersion

`func (o *AuditLogEntryListingRep) GetPreviousVersion() interface{}`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *AuditLogEntryListingRep) GetPreviousVersionOk() (*interface{}, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *AuditLogEntryListingRep) SetPreviousVersion(v interface{})`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *AuditLogEntryListingRep) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.

### SetPreviousVersionNil

`func (o *AuditLogEntryListingRep) SetPreviousVersionNil(b bool)`

 SetPreviousVersionNil sets the value for PreviousVersion to be an explicit nil

### UnsetPreviousVersion
`func (o *AuditLogEntryListingRep) UnsetPreviousVersion()`

UnsetPreviousVersion ensures that no value is present for PreviousVersion, not even an explicit nil
### GetCurrentVersion

`func (o *AuditLogEntryListingRep) GetCurrentVersion() interface{}`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *AuditLogEntryListingRep) GetCurrentVersionOk() (*interface{}, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *AuditLogEntryListingRep) SetCurrentVersion(v interface{})`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *AuditLogEntryListingRep) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### SetCurrentVersionNil

`func (o *AuditLogEntryListingRep) SetCurrentVersionNil(b bool)`

 SetCurrentVersionNil sets the value for CurrentVersion to be an explicit nil

### UnsetCurrentVersion
`func (o *AuditLogEntryListingRep) UnsetCurrentVersion()`

UnsetCurrentVersion ensures that no value is present for CurrentVersion, not even an explicit nil
### GetSubentries

`func (o *AuditLogEntryListingRep) GetSubentries() []AuditLogEntryListingRep`

GetSubentries returns the Subentries field if non-nil, zero value otherwise.

### GetSubentriesOk

`func (o *AuditLogEntryListingRep) GetSubentriesOk() (*[]AuditLogEntryListingRep, bool)`

GetSubentriesOk returns a tuple with the Subentries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubentries

`func (o *AuditLogEntryListingRep) SetSubentries(v []AuditLogEntryListingRep)`

SetSubentries sets Subentries field to given value.

### HasSubentries

`func (o *AuditLogEntryListingRep) HasSubentries() bool`

HasSubentries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


