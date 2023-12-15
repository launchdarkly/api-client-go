# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to [**[]HeaderItems**](HeaderItems.md) |  | [optional] 
**HmacSignature** | Pointer to [**HMACSignature**](HMACSignature.md) |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewEndpoint

`func NewEndpoint() *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *Endpoint) GetHeaders() []HeaderItems`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Endpoint) GetHeadersOk() (*[]HeaderItems, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Endpoint) SetHeaders(v []HeaderItems)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Endpoint) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHmacSignature

`func (o *Endpoint) GetHmacSignature() HMACSignature`

GetHmacSignature returns the HmacSignature field if non-nil, zero value otherwise.

### GetHmacSignatureOk

`func (o *Endpoint) GetHmacSignatureOk() (*HMACSignature, bool)`

GetHmacSignatureOk returns a tuple with the HmacSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSignature

`func (o *Endpoint) SetHmacSignature(v HMACSignature)`

SetHmacSignature sets HmacSignature field to given value.

### HasHmacSignature

`func (o *Endpoint) HasHmacSignature() bool`

HasHmacSignature returns a boolean if a field has been set.

### GetMethod

`func (o *Endpoint) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Endpoint) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Endpoint) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Endpoint) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *Endpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Endpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Endpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Endpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


