# TokenDataRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]CoreLink**](CoreLink.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ending** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **bool** |  | [optional] 

## Methods

### NewTokenDataRep

`func NewTokenDataRep() *TokenDataRep`

NewTokenDataRep instantiates a new TokenDataRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDataRepWithDefaults

`func NewTokenDataRepWithDefaults() *TokenDataRep`

NewTokenDataRepWithDefaults instantiates a new TokenDataRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TokenDataRep) GetLinks() []CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TokenDataRep) GetLinksOk() (*[]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TokenDataRep) SetLinks(v []CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TokenDataRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *TokenDataRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenDataRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenDataRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenDataRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TokenDataRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenDataRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenDataRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenDataRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnding

`func (o *TokenDataRep) GetEnding() string`

GetEnding returns the Ending field if non-nil, zero value otherwise.

### GetEndingOk

`func (o *TokenDataRep) GetEndingOk() (*string, bool)`

GetEndingOk returns a tuple with the Ending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnding

`func (o *TokenDataRep) SetEnding(v string)`

SetEnding sets Ending field to given value.

### HasEnding

`func (o *TokenDataRep) HasEnding() bool`

HasEnding returns a boolean if a field has been set.

### GetServiceToken

`func (o *TokenDataRep) GetServiceToken() bool`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *TokenDataRep) GetServiceTokenOk() (*bool, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *TokenDataRep) SetServiceToken(v bool)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *TokenDataRep) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


