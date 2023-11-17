# TokenSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the token | [optional] 
**Ending** | Pointer to **string** | The last few characters of the token | [optional] 
**ServiceToken** | Pointer to **bool** | Whether this is a service token | [optional] 

## Methods

### NewTokenSummary

`func NewTokenSummary() *TokenSummary`

NewTokenSummary instantiates a new TokenSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenSummaryWithDefaults

`func NewTokenSummaryWithDefaults() *TokenSummary`

NewTokenSummaryWithDefaults instantiates a new TokenSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TokenSummary) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TokenSummary) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TokenSummary) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TokenSummary) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *TokenSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TokenSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnding

`func (o *TokenSummary) GetEnding() string`

GetEnding returns the Ending field if non-nil, zero value otherwise.

### GetEndingOk

`func (o *TokenSummary) GetEndingOk() (*string, bool)`

GetEndingOk returns a tuple with the Ending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnding

`func (o *TokenSummary) SetEnding(v string)`

SetEnding sets Ending field to given value.

### HasEnding

`func (o *TokenSummary) HasEnding() bool`

HasEnding returns a boolean if a field has been set.

### GetServiceToken

`func (o *TokenSummary) GetServiceToken() bool`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *TokenSummary) GetServiceTokenOk() (*bool, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *TokenSummary) SetServiceToken(v bool)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *TokenSummary) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


