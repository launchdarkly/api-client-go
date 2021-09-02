# AuditLogEntryListingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**Id** | **string** |  | 
**AccountId** | **string** |  | 
**Date** | **int64** |  | 
**Accesses** | [**[]ResourceAccess**](ResourceAccess.md) |  | 
**Kind** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**ShortDescription** | **string** |  | 
**Comment** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to [**SubjectDataRep**](SubjectDataRep.md) |  | [optional] 
**Member** | Pointer to [**MemberDataRep**](MemberDataRep.md) |  | [optional] 
**Token** | Pointer to [**TokenDataRep**](TokenDataRep.md) |  | [optional] 
**App** | Pointer to [**AuthorizedAppDataRep**](AuthorizedAppDataRep.md) |  | [optional] 
**TitleVerb** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**TargetResourceRep**](TargetResourceRep.md) |  | [optional] 
**Parent** | Pointer to [**ParentResourceRep**](ParentResourceRep.md) |  | [optional] 

## Methods

### NewAuditLogEntryListingRep

`func NewAuditLogEntryListingRep(links map[string]Link, id string, accountId string, date int64, accesses []ResourceAccess, kind string, name string, description string, shortDescription string, ) *AuditLogEntryListingRep`

NewAuditLogEntryListingRep instantiates a new AuditLogEntryListingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEntryListingRepWithDefaults

`func NewAuditLogEntryListingRepWithDefaults() *AuditLogEntryListingRep`

NewAuditLogEntryListingRepWithDefaults instantiates a new AuditLogEntryListingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuditLogEntryListingRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuditLogEntryListingRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuditLogEntryListingRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetId

`func (o *AuditLogEntryListingRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogEntryListingRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogEntryListingRep) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *AuditLogEntryListingRep) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AuditLogEntryListingRep) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AuditLogEntryListingRep) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetDate

`func (o *AuditLogEntryListingRep) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AuditLogEntryListingRep) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AuditLogEntryListingRep) SetDate(v int64)`

SetDate sets Date field to given value.


### GetAccesses

`func (o *AuditLogEntryListingRep) GetAccesses() []ResourceAccess`

GetAccesses returns the Accesses field if non-nil, zero value otherwise.

### GetAccessesOk

`func (o *AuditLogEntryListingRep) GetAccessesOk() (*[]ResourceAccess, bool)`

GetAccessesOk returns a tuple with the Accesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccesses

`func (o *AuditLogEntryListingRep) SetAccesses(v []ResourceAccess)`

SetAccesses sets Accesses field to given value.


### GetKind

`func (o *AuditLogEntryListingRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuditLogEntryListingRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuditLogEntryListingRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *AuditLogEntryListingRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditLogEntryListingRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditLogEntryListingRep) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuditLogEntryListingRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditLogEntryListingRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditLogEntryListingRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetShortDescription

`func (o *AuditLogEntryListingRep) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *AuditLogEntryListingRep) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *AuditLogEntryListingRep) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.


### GetComment

`func (o *AuditLogEntryListingRep) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AuditLogEntryListingRep) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AuditLogEntryListingRep) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AuditLogEntryListingRep) HasComment() bool`

HasComment returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


