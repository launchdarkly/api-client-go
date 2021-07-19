# AuditLogEntryListingRepCollectionItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 
**Accesses** | Pointer to [**[]AuditLogEntryListingRepAccesses**](AuditLogEntryListingRepAccesses.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
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

### NewAuditLogEntryListingRepCollectionItems

`func NewAuditLogEntryListingRepCollectionItems() *AuditLogEntryListingRepCollectionItems`

NewAuditLogEntryListingRepCollectionItems instantiates a new AuditLogEntryListingRepCollectionItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEntryListingRepCollectionItemsWithDefaults

`func NewAuditLogEntryListingRepCollectionItemsWithDefaults() *AuditLogEntryListingRepCollectionItems`

NewAuditLogEntryListingRepCollectionItemsWithDefaults instantiates a new AuditLogEntryListingRepCollectionItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuditLogEntryListingRepCollectionItems) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuditLogEntryListingRepCollectionItems) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuditLogEntryListingRepCollectionItems) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuditLogEntryListingRepCollectionItems) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *AuditLogEntryListingRepCollectionItems) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogEntryListingRepCollectionItems) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogEntryListingRepCollectionItems) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogEntryListingRepCollectionItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *AuditLogEntryListingRepCollectionItems) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AuditLogEntryListingRepCollectionItems) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AuditLogEntryListingRepCollectionItems) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AuditLogEntryListingRepCollectionItems) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDate

`func (o *AuditLogEntryListingRepCollectionItems) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AuditLogEntryListingRepCollectionItems) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AuditLogEntryListingRepCollectionItems) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *AuditLogEntryListingRepCollectionItems) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAccesses

`func (o *AuditLogEntryListingRepCollectionItems) GetAccesses() []AuditLogEntryListingRepAccesses`

GetAccesses returns the Accesses field if non-nil, zero value otherwise.

### GetAccessesOk

`func (o *AuditLogEntryListingRepCollectionItems) GetAccessesOk() (*[]AuditLogEntryListingRepAccesses, bool)`

GetAccessesOk returns a tuple with the Accesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccesses

`func (o *AuditLogEntryListingRepCollectionItems) SetAccesses(v []AuditLogEntryListingRepAccesses)`

SetAccesses sets Accesses field to given value.

### HasAccesses

`func (o *AuditLogEntryListingRepCollectionItems) HasAccesses() bool`

HasAccesses returns a boolean if a field has been set.

### GetKind

`func (o *AuditLogEntryListingRepCollectionItems) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuditLogEntryListingRepCollectionItems) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuditLogEntryListingRepCollectionItems) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AuditLogEntryListingRepCollectionItems) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *AuditLogEntryListingRepCollectionItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditLogEntryListingRepCollectionItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditLogEntryListingRepCollectionItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditLogEntryListingRepCollectionItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AuditLogEntryListingRepCollectionItems) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditLogEntryListingRepCollectionItems) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditLogEntryListingRepCollectionItems) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuditLogEntryListingRepCollectionItems) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *AuditLogEntryListingRepCollectionItems) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *AuditLogEntryListingRepCollectionItems) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *AuditLogEntryListingRepCollectionItems) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *AuditLogEntryListingRepCollectionItems) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetComment

`func (o *AuditLogEntryListingRepCollectionItems) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AuditLogEntryListingRepCollectionItems) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AuditLogEntryListingRepCollectionItems) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AuditLogEntryListingRepCollectionItems) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSubject

`func (o *AuditLogEntryListingRepCollectionItems) GetSubject() SubjectDataRep`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuditLogEntryListingRepCollectionItems) GetSubjectOk() (*SubjectDataRep, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuditLogEntryListingRepCollectionItems) SetSubject(v SubjectDataRep)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AuditLogEntryListingRepCollectionItems) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMember

`func (o *AuditLogEntryListingRepCollectionItems) GetMember() MemberDataRep`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *AuditLogEntryListingRepCollectionItems) GetMemberOk() (*MemberDataRep, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *AuditLogEntryListingRepCollectionItems) SetMember(v MemberDataRep)`

SetMember sets Member field to given value.

### HasMember

`func (o *AuditLogEntryListingRepCollectionItems) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetToken

`func (o *AuditLogEntryListingRepCollectionItems) GetToken() TokenDataRep`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuditLogEntryListingRepCollectionItems) GetTokenOk() (*TokenDataRep, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuditLogEntryListingRepCollectionItems) SetToken(v TokenDataRep)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuditLogEntryListingRepCollectionItems) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetApp

`func (o *AuditLogEntryListingRepCollectionItems) GetApp() AuthorizedAppDataRep`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AuditLogEntryListingRepCollectionItems) GetAppOk() (*AuthorizedAppDataRep, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AuditLogEntryListingRepCollectionItems) SetApp(v AuthorizedAppDataRep)`

SetApp sets App field to given value.

### HasApp

`func (o *AuditLogEntryListingRepCollectionItems) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetTitleVerb

`func (o *AuditLogEntryListingRepCollectionItems) GetTitleVerb() string`

GetTitleVerb returns the TitleVerb field if non-nil, zero value otherwise.

### GetTitleVerbOk

`func (o *AuditLogEntryListingRepCollectionItems) GetTitleVerbOk() (*string, bool)`

GetTitleVerbOk returns a tuple with the TitleVerb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleVerb

`func (o *AuditLogEntryListingRepCollectionItems) SetTitleVerb(v string)`

SetTitleVerb sets TitleVerb field to given value.

### HasTitleVerb

`func (o *AuditLogEntryListingRepCollectionItems) HasTitleVerb() bool`

HasTitleVerb returns a boolean if a field has been set.

### GetTitle

`func (o *AuditLogEntryListingRepCollectionItems) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuditLogEntryListingRepCollectionItems) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuditLogEntryListingRepCollectionItems) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuditLogEntryListingRepCollectionItems) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTarget

`func (o *AuditLogEntryListingRepCollectionItems) GetTarget() TargetResourceRep`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AuditLogEntryListingRepCollectionItems) GetTargetOk() (*TargetResourceRep, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AuditLogEntryListingRepCollectionItems) SetTarget(v TargetResourceRep)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AuditLogEntryListingRepCollectionItems) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetParent

`func (o *AuditLogEntryListingRepCollectionItems) GetParent() ParentResourceRep`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuditLogEntryListingRepCollectionItems) GetParentOk() (*ParentResourceRep, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuditLogEntryListingRepCollectionItems) SetParent(v ParentResourceRep)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuditLogEntryListingRepCollectionItems) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


