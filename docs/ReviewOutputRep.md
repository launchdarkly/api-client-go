# ReviewOutputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**CreationDate** | Pointer to **int64** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**MemberId** | Pointer to **string** |  | [optional] 

## Methods

### NewReviewOutputRep

`func NewReviewOutputRep(id string, kind string, ) *ReviewOutputRep`

NewReviewOutputRep instantiates a new ReviewOutputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewOutputRepWithDefaults

`func NewReviewOutputRepWithDefaults() *ReviewOutputRep`

NewReviewOutputRepWithDefaults instantiates a new ReviewOutputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewOutputRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewOutputRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewOutputRep) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ReviewOutputRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReviewOutputRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReviewOutputRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetCreationDate

`func (o *ReviewOutputRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ReviewOutputRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ReviewOutputRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ReviewOutputRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetComment

`func (o *ReviewOutputRep) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ReviewOutputRep) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ReviewOutputRep) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ReviewOutputRep) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetMemberId

`func (o *ReviewOutputRep) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *ReviewOutputRep) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *ReviewOutputRep) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *ReviewOutputRep) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


