# TitleRep

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

### NewTitleRep

`func NewTitleRep() *TitleRep`

NewTitleRep instantiates a new TitleRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTitleRepWithDefaults

`func NewTitleRepWithDefaults() *TitleRep`

NewTitleRepWithDefaults instantiates a new TitleRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *TitleRep) GetSubject() SubjectDataRep`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TitleRep) GetSubjectOk() (*SubjectDataRep, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TitleRep) SetSubject(v SubjectDataRep)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TitleRep) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMember

`func (o *TitleRep) GetMember() MemberDataRep`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TitleRep) GetMemberOk() (*MemberDataRep, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TitleRep) SetMember(v MemberDataRep)`

SetMember sets Member field to given value.

### HasMember

`func (o *TitleRep) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetToken

`func (o *TitleRep) GetToken() TokenDataRep`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TitleRep) GetTokenOk() (*TokenDataRep, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TitleRep) SetToken(v TokenDataRep)`

SetToken sets Token field to given value.

### HasToken

`func (o *TitleRep) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetApp

`func (o *TitleRep) GetApp() AuthorizedAppDataRep`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *TitleRep) GetAppOk() (*AuthorizedAppDataRep, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *TitleRep) SetApp(v AuthorizedAppDataRep)`

SetApp sets App field to given value.

### HasApp

`func (o *TitleRep) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetTitleVerb

`func (o *TitleRep) GetTitleVerb() string`

GetTitleVerb returns the TitleVerb field if non-nil, zero value otherwise.

### GetTitleVerbOk

`func (o *TitleRep) GetTitleVerbOk() (*string, bool)`

GetTitleVerbOk returns a tuple with the TitleVerb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleVerb

`func (o *TitleRep) SetTitleVerb(v string)`

SetTitleVerb sets TitleVerb field to given value.

### HasTitleVerb

`func (o *TitleRep) HasTitleVerb() bool`

HasTitleVerb returns a boolean if a field has been set.

### GetTitle

`func (o *TitleRep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TitleRep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TitleRep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TitleRep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTarget

`func (o *TitleRep) GetTarget() TargetResourceRep`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TitleRep) GetTargetOk() (*TargetResourceRep, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TitleRep) SetTarget(v TargetResourceRep)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TitleRep) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetParent

`func (o *TitleRep) GetParent() ParentResourceRep`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TitleRep) GetParentOk() (*ParentResourceRep, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TitleRep) SetParent(v ParentResourceRep)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TitleRep) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetDelta

`func (o *TitleRep) GetDelta() interface{}`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *TitleRep) GetDeltaOk() (*interface{}, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *TitleRep) SetDelta(v interface{})`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *TitleRep) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### SetDeltaNil

`func (o *TitleRep) SetDeltaNil(b bool)`

 SetDeltaNil sets the value for Delta to be an explicit nil

### UnsetDelta
`func (o *TitleRep) UnsetDelta()`

UnsetDelta ensures that no value is present for Delta, not even an explicit nil
### GetTriggerBody

`func (o *TitleRep) GetTriggerBody() interface{}`

GetTriggerBody returns the TriggerBody field if non-nil, zero value otherwise.

### GetTriggerBodyOk

`func (o *TitleRep) GetTriggerBodyOk() (*interface{}, bool)`

GetTriggerBodyOk returns a tuple with the TriggerBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerBody

`func (o *TitleRep) SetTriggerBody(v interface{})`

SetTriggerBody sets TriggerBody field to given value.

### HasTriggerBody

`func (o *TitleRep) HasTriggerBody() bool`

HasTriggerBody returns a boolean if a field has been set.

### SetTriggerBodyNil

`func (o *TitleRep) SetTriggerBodyNil(b bool)`

 SetTriggerBodyNil sets the value for TriggerBody to be an explicit nil

### UnsetTriggerBody
`func (o *TitleRep) UnsetTriggerBody()`

UnsetTriggerBody ensures that no value is present for TriggerBody, not even an explicit nil
### GetMerge

`func (o *TitleRep) GetMerge() interface{}`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *TitleRep) GetMergeOk() (*interface{}, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *TitleRep) SetMerge(v interface{})`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *TitleRep) HasMerge() bool`

HasMerge returns a boolean if a field has been set.

### SetMergeNil

`func (o *TitleRep) SetMergeNil(b bool)`

 SetMergeNil sets the value for Merge to be an explicit nil

### UnsetMerge
`func (o *TitleRep) UnsetMerge()`

UnsetMerge ensures that no value is present for Merge, not even an explicit nil
### GetPreviousVersion

`func (o *TitleRep) GetPreviousVersion() interface{}`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *TitleRep) GetPreviousVersionOk() (*interface{}, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *TitleRep) SetPreviousVersion(v interface{})`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *TitleRep) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.

### SetPreviousVersionNil

`func (o *TitleRep) SetPreviousVersionNil(b bool)`

 SetPreviousVersionNil sets the value for PreviousVersion to be an explicit nil

### UnsetPreviousVersion
`func (o *TitleRep) UnsetPreviousVersion()`

UnsetPreviousVersion ensures that no value is present for PreviousVersion, not even an explicit nil
### GetCurrentVersion

`func (o *TitleRep) GetCurrentVersion() interface{}`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *TitleRep) GetCurrentVersionOk() (*interface{}, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *TitleRep) SetCurrentVersion(v interface{})`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *TitleRep) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### SetCurrentVersionNil

`func (o *TitleRep) SetCurrentVersionNil(b bool)`

 SetCurrentVersionNil sets the value for CurrentVersion to be an explicit nil

### UnsetCurrentVersion
`func (o *TitleRep) UnsetCurrentVersion()`

UnsetCurrentVersion ensures that no value is present for CurrentVersion, not even an explicit nil
### GetSubentries

`func (o *TitleRep) GetSubentries() []AuditLogEntryListingRep`

GetSubentries returns the Subentries field if non-nil, zero value otherwise.

### GetSubentriesOk

`func (o *TitleRep) GetSubentriesOk() (*[]AuditLogEntryListingRep, bool)`

GetSubentriesOk returns a tuple with the Subentries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubentries

`func (o *TitleRep) SetSubentries(v []AuditLogEntryListingRep)`

SetSubentries sets Subentries field to given value.

### HasSubentries

`func (o *TitleRep) HasSubentries() bool`

HasSubentries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


