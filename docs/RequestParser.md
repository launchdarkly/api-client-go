# RequestParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddMemberArrayPath** | Pointer to **string** |  | [optional] 
**ArrayInclusion** | Pointer to [**ArrayInclusionMatcher**](ArrayInclusionMatcher.md) |  | [optional] 
**CohortIdPath** | Pointer to **string** |  | [optional] 
**CohortNamePath** | Pointer to **string** |  | [optional] 
**CohortUrlPath** | Pointer to **string** |  | [optional] 
**ContextKindPath** | Pointer to **string** |  | [optional] 
**EnvironmentIdPath** | Pointer to **string** |  | [optional] 
**MemberArrayParser** | Pointer to [**MemberArrayParser**](MemberArrayParser.md) |  | [optional] 
**MemberArrayPath** | Pointer to **string** |  | [optional] 
**RemoveMemberArrayPath** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestParser

`func NewRequestParser() *RequestParser`

NewRequestParser instantiates a new RequestParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestParserWithDefaults

`func NewRequestParserWithDefaults() *RequestParser`

NewRequestParserWithDefaults instantiates a new RequestParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddMemberArrayPath

`func (o *RequestParser) GetAddMemberArrayPath() string`

GetAddMemberArrayPath returns the AddMemberArrayPath field if non-nil, zero value otherwise.

### GetAddMemberArrayPathOk

`func (o *RequestParser) GetAddMemberArrayPathOk() (*string, bool)`

GetAddMemberArrayPathOk returns a tuple with the AddMemberArrayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddMemberArrayPath

`func (o *RequestParser) SetAddMemberArrayPath(v string)`

SetAddMemberArrayPath sets AddMemberArrayPath field to given value.

### HasAddMemberArrayPath

`func (o *RequestParser) HasAddMemberArrayPath() bool`

HasAddMemberArrayPath returns a boolean if a field has been set.

### GetArrayInclusion

`func (o *RequestParser) GetArrayInclusion() ArrayInclusionMatcher`

GetArrayInclusion returns the ArrayInclusion field if non-nil, zero value otherwise.

### GetArrayInclusionOk

`func (o *RequestParser) GetArrayInclusionOk() (*ArrayInclusionMatcher, bool)`

GetArrayInclusionOk returns a tuple with the ArrayInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayInclusion

`func (o *RequestParser) SetArrayInclusion(v ArrayInclusionMatcher)`

SetArrayInclusion sets ArrayInclusion field to given value.

### HasArrayInclusion

`func (o *RequestParser) HasArrayInclusion() bool`

HasArrayInclusion returns a boolean if a field has been set.

### GetCohortIdPath

`func (o *RequestParser) GetCohortIdPath() string`

GetCohortIdPath returns the CohortIdPath field if non-nil, zero value otherwise.

### GetCohortIdPathOk

`func (o *RequestParser) GetCohortIdPathOk() (*string, bool)`

GetCohortIdPathOk returns a tuple with the CohortIdPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortIdPath

`func (o *RequestParser) SetCohortIdPath(v string)`

SetCohortIdPath sets CohortIdPath field to given value.

### HasCohortIdPath

`func (o *RequestParser) HasCohortIdPath() bool`

HasCohortIdPath returns a boolean if a field has been set.

### GetCohortNamePath

`func (o *RequestParser) GetCohortNamePath() string`

GetCohortNamePath returns the CohortNamePath field if non-nil, zero value otherwise.

### GetCohortNamePathOk

`func (o *RequestParser) GetCohortNamePathOk() (*string, bool)`

GetCohortNamePathOk returns a tuple with the CohortNamePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortNamePath

`func (o *RequestParser) SetCohortNamePath(v string)`

SetCohortNamePath sets CohortNamePath field to given value.

### HasCohortNamePath

`func (o *RequestParser) HasCohortNamePath() bool`

HasCohortNamePath returns a boolean if a field has been set.

### GetCohortUrlPath

`func (o *RequestParser) GetCohortUrlPath() string`

GetCohortUrlPath returns the CohortUrlPath field if non-nil, zero value otherwise.

### GetCohortUrlPathOk

`func (o *RequestParser) GetCohortUrlPathOk() (*string, bool)`

GetCohortUrlPathOk returns a tuple with the CohortUrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortUrlPath

`func (o *RequestParser) SetCohortUrlPath(v string)`

SetCohortUrlPath sets CohortUrlPath field to given value.

### HasCohortUrlPath

`func (o *RequestParser) HasCohortUrlPath() bool`

HasCohortUrlPath returns a boolean if a field has been set.

### GetContextKindPath

`func (o *RequestParser) GetContextKindPath() string`

GetContextKindPath returns the ContextKindPath field if non-nil, zero value otherwise.

### GetContextKindPathOk

`func (o *RequestParser) GetContextKindPathOk() (*string, bool)`

GetContextKindPathOk returns a tuple with the ContextKindPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKindPath

`func (o *RequestParser) SetContextKindPath(v string)`

SetContextKindPath sets ContextKindPath field to given value.

### HasContextKindPath

`func (o *RequestParser) HasContextKindPath() bool`

HasContextKindPath returns a boolean if a field has been set.

### GetEnvironmentIdPath

`func (o *RequestParser) GetEnvironmentIdPath() string`

GetEnvironmentIdPath returns the EnvironmentIdPath field if non-nil, zero value otherwise.

### GetEnvironmentIdPathOk

`func (o *RequestParser) GetEnvironmentIdPathOk() (*string, bool)`

GetEnvironmentIdPathOk returns a tuple with the EnvironmentIdPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentIdPath

`func (o *RequestParser) SetEnvironmentIdPath(v string)`

SetEnvironmentIdPath sets EnvironmentIdPath field to given value.

### HasEnvironmentIdPath

`func (o *RequestParser) HasEnvironmentIdPath() bool`

HasEnvironmentIdPath returns a boolean if a field has been set.

### GetMemberArrayParser

`func (o *RequestParser) GetMemberArrayParser() MemberArrayParser`

GetMemberArrayParser returns the MemberArrayParser field if non-nil, zero value otherwise.

### GetMemberArrayParserOk

`func (o *RequestParser) GetMemberArrayParserOk() (*MemberArrayParser, bool)`

GetMemberArrayParserOk returns a tuple with the MemberArrayParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberArrayParser

`func (o *RequestParser) SetMemberArrayParser(v MemberArrayParser)`

SetMemberArrayParser sets MemberArrayParser field to given value.

### HasMemberArrayParser

`func (o *RequestParser) HasMemberArrayParser() bool`

HasMemberArrayParser returns a boolean if a field has been set.

### GetMemberArrayPath

`func (o *RequestParser) GetMemberArrayPath() string`

GetMemberArrayPath returns the MemberArrayPath field if non-nil, zero value otherwise.

### GetMemberArrayPathOk

`func (o *RequestParser) GetMemberArrayPathOk() (*string, bool)`

GetMemberArrayPathOk returns a tuple with the MemberArrayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberArrayPath

`func (o *RequestParser) SetMemberArrayPath(v string)`

SetMemberArrayPath sets MemberArrayPath field to given value.

### HasMemberArrayPath

`func (o *RequestParser) HasMemberArrayPath() bool`

HasMemberArrayPath returns a boolean if a field has been set.

### GetRemoveMemberArrayPath

`func (o *RequestParser) GetRemoveMemberArrayPath() string`

GetRemoveMemberArrayPath returns the RemoveMemberArrayPath field if non-nil, zero value otherwise.

### GetRemoveMemberArrayPathOk

`func (o *RequestParser) GetRemoveMemberArrayPathOk() (*string, bool)`

GetRemoveMemberArrayPathOk returns a tuple with the RemoveMemberArrayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveMemberArrayPath

`func (o *RequestParser) SetRemoveMemberArrayPath(v string)`

SetRemoveMemberArrayPath sets RemoveMemberArrayPath field to given value.

### HasRemoveMemberArrayPath

`func (o *RequestParser) HasRemoveMemberArrayPath() bool`

HasRemoveMemberArrayPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


