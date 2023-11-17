# TitleRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to [**SubjectDataRep**](SubjectDataRep.md) |  | [optional] 
**Member** | Pointer to [**MemberDataRep**](MemberDataRep.md) |  | [optional] 
**Token** | Pointer to [**TokenSummary**](TokenSummary.md) |  | [optional] 
**App** | Pointer to [**AuthorizedAppDataRep**](AuthorizedAppDataRep.md) |  | [optional] 
**TitleVerb** | Pointer to **string** | The action and resource recorded in this audit log entry | [optional] 
**Title** | Pointer to **string** | A description of what occurred, in the format &lt;code&gt;member&lt;/code&gt; &lt;code&gt;titleVerb&lt;/code&gt; &lt;code&gt;target&lt;/code&gt; | [optional] 
**Target** | Pointer to [**TargetResourceRep**](TargetResourceRep.md) |  | [optional] 
**Parent** | Pointer to [**ParentResourceRep**](ParentResourceRep.md) |  | [optional] 

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

`func (o *TitleRep) GetToken() TokenSummary`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TitleRep) GetTokenOk() (*TokenSummary, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TitleRep) SetToken(v TokenSummary)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


