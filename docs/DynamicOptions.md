# DynamicOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**Parser** | Pointer to [**DynamicOptionsParser**](DynamicOptionsParser.md) |  | [optional] 

## Methods

### NewDynamicOptions

`func NewDynamicOptions() *DynamicOptions`

NewDynamicOptions instantiates a new DynamicOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicOptionsWithDefaults

`func NewDynamicOptionsWithDefaults() *DynamicOptions`

NewDynamicOptionsWithDefaults instantiates a new DynamicOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *DynamicOptions) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DynamicOptions) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DynamicOptions) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *DynamicOptions) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetParser

`func (o *DynamicOptions) GetParser() DynamicOptionsParser`

GetParser returns the Parser field if non-nil, zero value otherwise.

### GetParserOk

`func (o *DynamicOptions) GetParserOk() (*DynamicOptionsParser, bool)`

GetParserOk returns a tuple with the Parser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParser

`func (o *DynamicOptions) SetParser(v DynamicOptionsParser)`

SetParser sets Parser field to given value.

### HasParser

`func (o *DynamicOptions) HasParser() bool`

HasParser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


