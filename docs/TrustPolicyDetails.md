# TrustPolicyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version of the trust policy | [optional] 
**Statement** | Pointer to [**[]TrustPolicyStatement**](TrustPolicyStatement.md) | The statements of the trust policy | [optional] 

## Methods

### NewTrustPolicyDetails

`func NewTrustPolicyDetails() *TrustPolicyDetails`

NewTrustPolicyDetails instantiates a new TrustPolicyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustPolicyDetailsWithDefaults

`func NewTrustPolicyDetailsWithDefaults() *TrustPolicyDetails`

NewTrustPolicyDetailsWithDefaults instantiates a new TrustPolicyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *TrustPolicyDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrustPolicyDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrustPolicyDetails) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrustPolicyDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatement

`func (o *TrustPolicyDetails) GetStatement() []TrustPolicyStatement`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *TrustPolicyDetails) GetStatementOk() (*[]TrustPolicyStatement, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *TrustPolicyDetails) SetStatement(v []TrustPolicyStatement)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *TrustPolicyDetails) HasStatement() bool`

HasStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


