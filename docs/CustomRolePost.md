# CustomRolePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the custom role | 
**Key** | **string** | The custom role key | 
**Description** | Pointer to **string** | Description of custom role | [optional] 
**Policy** | [**[]StatementPost**](StatementPost.md) |  | 

## Methods

### NewCustomRolePost

`func NewCustomRolePost(name string, key string, policy []StatementPost, ) *CustomRolePost`

NewCustomRolePost instantiates a new CustomRolePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRolePostWithDefaults

`func NewCustomRolePostWithDefaults() *CustomRolePost`

NewCustomRolePostWithDefaults instantiates a new CustomRolePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomRolePost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomRolePost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomRolePost) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *CustomRolePost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomRolePost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomRolePost) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *CustomRolePost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomRolePost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomRolePost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomRolePost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicy

`func (o *CustomRolePost) GetPolicy() []StatementPost`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CustomRolePost) GetPolicyOk() (*[]StatementPost, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CustomRolePost) SetPolicy(v []StatementPost)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


