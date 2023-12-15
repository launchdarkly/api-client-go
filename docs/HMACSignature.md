# HMACSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderName** | Pointer to **string** |  | [optional] 
**HmacSecretFormVariableKey** | Pointer to **string** |  | [optional] 

## Methods

### NewHMACSignature

`func NewHMACSignature() *HMACSignature`

NewHMACSignature instantiates a new HMACSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHMACSignatureWithDefaults

`func NewHMACSignatureWithDefaults() *HMACSignature`

NewHMACSignatureWithDefaults instantiates a new HMACSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderName

`func (o *HMACSignature) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *HMACSignature) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *HMACSignature) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *HMACSignature) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### GetHmacSecretFormVariableKey

`func (o *HMACSignature) GetHmacSecretFormVariableKey() string`

GetHmacSecretFormVariableKey returns the HmacSecretFormVariableKey field if non-nil, zero value otherwise.

### GetHmacSecretFormVariableKeyOk

`func (o *HMACSignature) GetHmacSecretFormVariableKeyOk() (*string, bool)`

GetHmacSecretFormVariableKeyOk returns a tuple with the HmacSecretFormVariableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSecretFormVariableKey

`func (o *HMACSignature) SetHmacSecretFormVariableKey(v string)`

SetHmacSecretFormVariableKey sets HmacSecretFormVariableKey field to given value.

### HasHmacSecretFormVariableKey

`func (o *HMACSignature) HasHmacSecretFormVariableKey() bool`

HasHmacSecretFormVariableKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


