# ReviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The approval request ID | 
**Kind** | **string** | The type of review action to take | 
**CreationDate** | Pointer to **int64** |  | [optional] 
**Comment** | Pointer to **string** | A comment describing the approval response | [optional] 
**MemberId** | Pointer to **string** | ID of account member that reviewed request | [optional] 
**ServiceTokenId** | Pointer to **string** | ID of account service token that reviewed request | [optional] 

## Methods

### NewReviewResponse

`func NewReviewResponse(id string, kind string, ) *ReviewResponse`

NewReviewResponse instantiates a new ReviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewResponseWithDefaults

`func NewReviewResponseWithDefaults() *ReviewResponse`

NewReviewResponseWithDefaults instantiates a new ReviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ReviewResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReviewResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReviewResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetCreationDate

`func (o *ReviewResponse) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ReviewResponse) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ReviewResponse) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ReviewResponse) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetComment

`func (o *ReviewResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ReviewResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ReviewResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ReviewResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetMemberId

`func (o *ReviewResponse) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *ReviewResponse) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *ReviewResponse) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *ReviewResponse) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetServiceTokenId

`func (o *ReviewResponse) GetServiceTokenId() string`

GetServiceTokenId returns the ServiceTokenId field if non-nil, zero value otherwise.

### GetServiceTokenIdOk

`func (o *ReviewResponse) GetServiceTokenIdOk() (*string, bool)`

GetServiceTokenIdOk returns a tuple with the ServiceTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenId

`func (o *ReviewResponse) SetServiceTokenId(v string)`

SetServiceTokenId sets ServiceTokenId field to given value.

### HasServiceTokenId

`func (o *ReviewResponse) HasServiceTokenId() bool`

HasServiceTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


