# AuditLogEntryRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Id** | **string** | The ID of the audit log entry | 
**AccountId** | **string** | The ID of the account to which this audit log entry belongs | 
**Date** | **int64** |  | 
**Accesses** | [**[]ResourceAccess**](ResourceAccess.md) | Details on the actions performed and resources acted on in this audit log entry | 
**Kind** | **string** |  | 
**Name** | **string** | The name of the resource this audit log entry refers to | 
**Description** | **string** | Description of the change recorded in the audit log entry | 
**ShortDescription** | **string** | Shorter version of the change recorded in the audit log entry | 
**Comment** | Pointer to **string** | Optional comment for the audit log entry | [optional] 
**Subject** | Pointer to [**SubjectDataRep**](SubjectDataRep.md) |  | [optional] 
**Member** | Pointer to [**MemberDataRep**](MemberDataRep.md) |  | [optional] 
**Token** | Pointer to [**TokenDataRep**](TokenDataRep.md) |  | [optional] 
**App** | Pointer to [**AuthorizedAppDataRep**](AuthorizedAppDataRep.md) |  | [optional] 
**TitleVerb** | Pointer to **string** | The action and resource recorded in this audit log entry | [optional] 
**Title** | Pointer to **string** | A description of what occurred, in the format &lt;code&gt;member&lt;/code&gt; &lt;code&gt;titleVerb&lt;/code&gt; &lt;code&gt;target&lt;/code&gt; | [optional] 
**Target** | Pointer to [**TargetResourceRep**](TargetResourceRep.md) |  | [optional] 
**Parent** | Pointer to [**ParentResourceRep**](ParentResourceRep.md) |  | [optional] 
**Delta** | Pointer to **interface{}** | If the audit log entry has been updated, this is the JSON patch body that was used in the request to update the entity | [optional] 
**TriggerBody** | Pointer to **interface{}** | A JSON representation of the external trigger for this audit log entry, if any | [optional] 
**Merge** | Pointer to **interface{}** | A JSON representation of the merge information for this audit log entry, if any | [optional] 
**PreviousVersion** | Pointer to **interface{}** | If the audit log entry has been updated, this is a JSON representation of the previous version of the entity | [optional] 
**CurrentVersion** | Pointer to **interface{}** | If the audit log entry has been updated, this is a JSON representation of the current version of the entity | [optional] 
**Subentries** | Pointer to [**[]AuditLogEntryListingRep**](AuditLogEntryListingRep.md) |  | [optional] 

## Methods

### NewAuditLogEntryRep

`func NewAuditLogEntryRep(links map[string]Link, id string, accountId string, date int64, accesses []ResourceAccess, kind string, name string, description string, shortDescription string, ) *AuditLogEntryRep`

NewAuditLogEntryRep instantiates a new AuditLogEntryRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEntryRepWithDefaults

`func NewAuditLogEntryRepWithDefaults() *AuditLogEntryRep`

NewAuditLogEntryRepWithDefaults instantiates a new AuditLogEntryRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuditLogEntryRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuditLogEntryRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuditLogEntryRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetId

`func (o *AuditLogEntryRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogEntryRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogEntryRep) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *AuditLogEntryRep) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AuditLogEntryRep) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AuditLogEntryRep) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetDate

`func (o *AuditLogEntryRep) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AuditLogEntryRep) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AuditLogEntryRep) SetDate(v int64)`

SetDate sets Date field to given value.


### GetAccesses

`func (o *AuditLogEntryRep) GetAccesses() []ResourceAccess`

GetAccesses returns the Accesses field if non-nil, zero value otherwise.

### GetAccessesOk

`func (o *AuditLogEntryRep) GetAccessesOk() (*[]ResourceAccess, bool)`

GetAccessesOk returns a tuple with the Accesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccesses

`func (o *AuditLogEntryRep) SetAccesses(v []ResourceAccess)`

SetAccesses sets Accesses field to given value.


### GetKind

`func (o *AuditLogEntryRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuditLogEntryRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuditLogEntryRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *AuditLogEntryRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditLogEntryRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditLogEntryRep) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuditLogEntryRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditLogEntryRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditLogEntryRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetShortDescription

`func (o *AuditLogEntryRep) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *AuditLogEntryRep) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *AuditLogEntryRep) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.


### GetComment

`func (o *AuditLogEntryRep) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AuditLogEntryRep) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AuditLogEntryRep) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AuditLogEntryRep) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSubject

`func (o *AuditLogEntryRep) GetSubject() SubjectDataRep`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuditLogEntryRep) GetSubjectOk() (*SubjectDataRep, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuditLogEntryRep) SetSubject(v SubjectDataRep)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AuditLogEntryRep) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMember

`func (o *AuditLogEntryRep) GetMember() MemberDataRep`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *AuditLogEntryRep) GetMemberOk() (*MemberDataRep, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *AuditLogEntryRep) SetMember(v MemberDataRep)`

SetMember sets Member field to given value.

### HasMember

`func (o *AuditLogEntryRep) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetToken

`func (o *AuditLogEntryRep) GetToken() TokenDataRep`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuditLogEntryRep) GetTokenOk() (*TokenDataRep, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuditLogEntryRep) SetToken(v TokenDataRep)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuditLogEntryRep) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetApp

`func (o *AuditLogEntryRep) GetApp() AuthorizedAppDataRep`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AuditLogEntryRep) GetAppOk() (*AuthorizedAppDataRep, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AuditLogEntryRep) SetApp(v AuthorizedAppDataRep)`

SetApp sets App field to given value.

### HasApp

`func (o *AuditLogEntryRep) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetTitleVerb

`func (o *AuditLogEntryRep) GetTitleVerb() string`

GetTitleVerb returns the TitleVerb field if non-nil, zero value otherwise.

### GetTitleVerbOk

`func (o *AuditLogEntryRep) GetTitleVerbOk() (*string, bool)`

GetTitleVerbOk returns a tuple with the TitleVerb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleVerb

`func (o *AuditLogEntryRep) SetTitleVerb(v string)`

SetTitleVerb sets TitleVerb field to given value.

### HasTitleVerb

`func (o *AuditLogEntryRep) HasTitleVerb() bool`

HasTitleVerb returns a boolean if a field has been set.

### GetTitle

`func (o *AuditLogEntryRep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuditLogEntryRep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuditLogEntryRep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuditLogEntryRep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTarget

`func (o *AuditLogEntryRep) GetTarget() TargetResourceRep`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AuditLogEntryRep) GetTargetOk() (*TargetResourceRep, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AuditLogEntryRep) SetTarget(v TargetResourceRep)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AuditLogEntryRep) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetParent

`func (o *AuditLogEntryRep) GetParent() ParentResourceRep`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuditLogEntryRep) GetParentOk() (*ParentResourceRep, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuditLogEntryRep) SetParent(v ParentResourceRep)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuditLogEntryRep) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetDelta

`func (o *AuditLogEntryRep) GetDelta() interface{}`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *AuditLogEntryRep) GetDeltaOk() (*interface{}, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *AuditLogEntryRep) SetDelta(v interface{})`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *AuditLogEntryRep) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### SetDeltaNil

`func (o *AuditLogEntryRep) SetDeltaNil(b bool)`

 SetDeltaNil sets the value for Delta to be an explicit nil

### UnsetDelta
`func (o *AuditLogEntryRep) UnsetDelta()`

UnsetDelta ensures that no value is present for Delta, not even an explicit nil
### GetTriggerBody

`func (o *AuditLogEntryRep) GetTriggerBody() interface{}`

GetTriggerBody returns the TriggerBody field if non-nil, zero value otherwise.

### GetTriggerBodyOk

`func (o *AuditLogEntryRep) GetTriggerBodyOk() (*interface{}, bool)`

GetTriggerBodyOk returns a tuple with the TriggerBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerBody

`func (o *AuditLogEntryRep) SetTriggerBody(v interface{})`

SetTriggerBody sets TriggerBody field to given value.

### HasTriggerBody

`func (o *AuditLogEntryRep) HasTriggerBody() bool`

HasTriggerBody returns a boolean if a field has been set.

### SetTriggerBodyNil

`func (o *AuditLogEntryRep) SetTriggerBodyNil(b bool)`

 SetTriggerBodyNil sets the value for TriggerBody to be an explicit nil

### UnsetTriggerBody
`func (o *AuditLogEntryRep) UnsetTriggerBody()`

UnsetTriggerBody ensures that no value is present for TriggerBody, not even an explicit nil
### GetMerge

`func (o *AuditLogEntryRep) GetMerge() interface{}`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *AuditLogEntryRep) GetMergeOk() (*interface{}, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *AuditLogEntryRep) SetMerge(v interface{})`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *AuditLogEntryRep) HasMerge() bool`

HasMerge returns a boolean if a field has been set.

### SetMergeNil

`func (o *AuditLogEntryRep) SetMergeNil(b bool)`

 SetMergeNil sets the value for Merge to be an explicit nil

### UnsetMerge
`func (o *AuditLogEntryRep) UnsetMerge()`

UnsetMerge ensures that no value is present for Merge, not even an explicit nil
### GetPreviousVersion

`func (o *AuditLogEntryRep) GetPreviousVersion() interface{}`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *AuditLogEntryRep) GetPreviousVersionOk() (*interface{}, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *AuditLogEntryRep) SetPreviousVersion(v interface{})`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *AuditLogEntryRep) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.

### SetPreviousVersionNil

`func (o *AuditLogEntryRep) SetPreviousVersionNil(b bool)`

 SetPreviousVersionNil sets the value for PreviousVersion to be an explicit nil

### UnsetPreviousVersion
`func (o *AuditLogEntryRep) UnsetPreviousVersion()`

UnsetPreviousVersion ensures that no value is present for PreviousVersion, not even an explicit nil
### GetCurrentVersion

`func (o *AuditLogEntryRep) GetCurrentVersion() interface{}`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *AuditLogEntryRep) GetCurrentVersionOk() (*interface{}, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *AuditLogEntryRep) SetCurrentVersion(v interface{})`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *AuditLogEntryRep) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### SetCurrentVersionNil

`func (o *AuditLogEntryRep) SetCurrentVersionNil(b bool)`

 SetCurrentVersionNil sets the value for CurrentVersion to be an explicit nil

### UnsetCurrentVersion
`func (o *AuditLogEntryRep) UnsetCurrentVersion()`

UnsetCurrentVersion ensures that no value is present for CurrentVersion, not even an explicit nil
### GetSubentries

`func (o *AuditLogEntryRep) GetSubentries() []AuditLogEntryListingRep`

GetSubentries returns the Subentries field if non-nil, zero value otherwise.

### GetSubentriesOk

`func (o *AuditLogEntryRep) GetSubentriesOk() (*[]AuditLogEntryListingRep, bool)`

GetSubentriesOk returns a tuple with the Subentries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubentries

`func (o *AuditLogEntryRep) SetSubentries(v []AuditLogEntryListingRep)`

SetSubentries sets Subentries field to given value.

### HasSubentries

`func (o *AuditLogEntryRep) HasSubentries() bool`

HasSubentries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


