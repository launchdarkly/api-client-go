# ValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**Parser** | Pointer to [**FeatureStoreValidationParser**](FeatureStoreValidationParser.md) |  | [optional] 

## Methods

### NewValidationRequest

`func NewValidationRequest() *ValidationRequest`

NewValidationRequest instantiates a new ValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationRequestWithDefaults

`func NewValidationRequestWithDefaults() *ValidationRequest`

NewValidationRequestWithDefaults instantiates a new ValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *ValidationRequest) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ValidationRequest) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ValidationRequest) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ValidationRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetParser

`func (o *ValidationRequest) GetParser() FeatureStoreValidationParser`

GetParser returns the Parser field if non-nil, zero value otherwise.

### GetParserOk

`func (o *ValidationRequest) GetParserOk() (*FeatureStoreValidationParser, bool)`

GetParserOk returns a tuple with the Parser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParser

`func (o *ValidationRequest) SetParser(v FeatureStoreValidationParser)`

SetParser sets Parser field to given value.

### HasParser

`func (o *ValidationRequest) HasParser() bool`

HasParser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


