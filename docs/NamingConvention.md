# NamingConvention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Case** | Pointer to **string** | The casing convention to enforce for new flag keys in this project | [optional] 
**Prefix** | Pointer to **string** | The prefix to enforce for new flag keys in this project | [optional] 

## Methods

### NewNamingConvention

`func NewNamingConvention() *NamingConvention`

NewNamingConvention instantiates a new NamingConvention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingConventionWithDefaults

`func NewNamingConventionWithDefaults() *NamingConvention`

NewNamingConventionWithDefaults instantiates a new NamingConvention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCase

`func (o *NamingConvention) GetCase() string`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *NamingConvention) GetCaseOk() (*string, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *NamingConvention) SetCase(v string)`

SetCase sets Case field to given value.

### HasCase

`func (o *NamingConvention) HasCase() bool`

HasCase returns a boolean if a field has been set.

### GetPrefix

`func (o *NamingConvention) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *NamingConvention) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *NamingConvention) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *NamingConvention) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


