# MemberListParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberArrayPath** | Pointer to **string** |  | [optional] 
**MemberItems** | Pointer to [**MemberItemsArray**](MemberItemsArray.md) |  | [optional] 

## Methods

### NewMemberListParser

`func NewMemberListParser() *MemberListParser`

NewMemberListParser instantiates a new MemberListParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberListParserWithDefaults

`func NewMemberListParserWithDefaults() *MemberListParser`

NewMemberListParserWithDefaults instantiates a new MemberListParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberArrayPath

`func (o *MemberListParser) GetMemberArrayPath() string`

GetMemberArrayPath returns the MemberArrayPath field if non-nil, zero value otherwise.

### GetMemberArrayPathOk

`func (o *MemberListParser) GetMemberArrayPathOk() (*string, bool)`

GetMemberArrayPathOk returns a tuple with the MemberArrayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberArrayPath

`func (o *MemberListParser) SetMemberArrayPath(v string)`

SetMemberArrayPath sets MemberArrayPath field to given value.

### HasMemberArrayPath

`func (o *MemberListParser) HasMemberArrayPath() bool`

HasMemberArrayPath returns a boolean if a field has been set.

### GetMemberItems

`func (o *MemberListParser) GetMemberItems() MemberItemsArray`

GetMemberItems returns the MemberItems field if non-nil, zero value otherwise.

### GetMemberItemsOk

`func (o *MemberListParser) GetMemberItemsOk() (*MemberItemsArray, bool)`

GetMemberItemsOk returns a tuple with the MemberItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberItems

`func (o *MemberListParser) SetMemberItems(v MemberItemsArray)`

SetMemberItems sets MemberItems field to given value.

### HasMemberItems

`func (o *MemberListParser) HasMemberItems() bool`

HasMemberItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


