# CreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**JsonBody** | Pointer to **string** |  | [optional] 
**Parser** | Pointer to [**ApprovalParser**](ApprovalParser.md) |  | [optional] 

## Methods

### NewCreationRequest

`func NewCreationRequest() *CreationRequest`

NewCreationRequest instantiates a new CreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreationRequestWithDefaults

`func NewCreationRequestWithDefaults() *CreationRequest`

NewCreationRequestWithDefaults instantiates a new CreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *CreationRequest) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreationRequest) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreationRequest) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *CreationRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetJsonBody

`func (o *CreationRequest) GetJsonBody() string`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *CreationRequest) GetJsonBodyOk() (*string, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *CreationRequest) SetJsonBody(v string)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *CreationRequest) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetParser

`func (o *CreationRequest) GetParser() ApprovalParser`

GetParser returns the Parser field if non-nil, zero value otherwise.

### GetParserOk

`func (o *CreationRequest) GetParserOk() (*ApprovalParser, bool)`

GetParserOk returns a tuple with the Parser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParser

`func (o *CreationRequest) SetParser(v ApprovalParser)`

SetParser sets Parser field to given value.

### HasParser

`func (o *CreationRequest) HasParser() bool`

HasParser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


