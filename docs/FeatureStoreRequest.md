# FeatureStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**PayloadPrefix** | Pointer to **string** |  | [optional] 
**PayloadSuffix** | Pointer to **string** |  | [optional] 

## Methods

### NewFeatureStoreRequest

`func NewFeatureStoreRequest() *FeatureStoreRequest`

NewFeatureStoreRequest instantiates a new FeatureStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureStoreRequestWithDefaults

`func NewFeatureStoreRequestWithDefaults() *FeatureStoreRequest`

NewFeatureStoreRequestWithDefaults instantiates a new FeatureStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *FeatureStoreRequest) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *FeatureStoreRequest) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *FeatureStoreRequest) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *FeatureStoreRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetPayloadPrefix

`func (o *FeatureStoreRequest) GetPayloadPrefix() string`

GetPayloadPrefix returns the PayloadPrefix field if non-nil, zero value otherwise.

### GetPayloadPrefixOk

`func (o *FeatureStoreRequest) GetPayloadPrefixOk() (*string, bool)`

GetPayloadPrefixOk returns a tuple with the PayloadPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadPrefix

`func (o *FeatureStoreRequest) SetPayloadPrefix(v string)`

SetPayloadPrefix sets PayloadPrefix field to given value.

### HasPayloadPrefix

`func (o *FeatureStoreRequest) HasPayloadPrefix() bool`

HasPayloadPrefix returns a boolean if a field has been set.

### GetPayloadSuffix

`func (o *FeatureStoreRequest) GetPayloadSuffix() string`

GetPayloadSuffix returns the PayloadSuffix field if non-nil, zero value otherwise.

### GetPayloadSuffixOk

`func (o *FeatureStoreRequest) GetPayloadSuffixOk() (*string, bool)`

GetPayloadSuffixOk returns a tuple with the PayloadSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadSuffix

`func (o *FeatureStoreRequest) SetPayloadSuffix(v string)`

SetPayloadSuffix sets PayloadSuffix field to given value.

### HasPayloadSuffix

`func (o *FeatureStoreRequest) HasPayloadSuffix() bool`

HasPayloadSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


