# DynamicOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUnknownValue** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**Filters** | Pointer to [**[]DynamicOptionFiltersItems**](DynamicOptionFiltersItems.md) |  | [optional] 
**Parser** | Pointer to [**DynamicOptionsParser**](DynamicOptionsParser.md) |  | [optional] 
**Source** | Pointer to [**DynamicOptionsSource**](DynamicOptionsSource.md) |  | [optional] 

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

### GetAllowUnknownValue

`func (o *DynamicOptions) GetAllowUnknownValue() bool`

GetAllowUnknownValue returns the AllowUnknownValue field if non-nil, zero value otherwise.

### GetAllowUnknownValueOk

`func (o *DynamicOptions) GetAllowUnknownValueOk() (*bool, bool)`

GetAllowUnknownValueOk returns a tuple with the AllowUnknownValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknownValue

`func (o *DynamicOptions) SetAllowUnknownValue(v bool)`

SetAllowUnknownValue sets AllowUnknownValue field to given value.

### HasAllowUnknownValue

`func (o *DynamicOptions) HasAllowUnknownValue() bool`

HasAllowUnknownValue returns a boolean if a field has been set.

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

### GetFilters

`func (o *DynamicOptions) GetFilters() []DynamicOptionFiltersItems`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DynamicOptions) GetFiltersOk() (*[]DynamicOptionFiltersItems, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DynamicOptions) SetFilters(v []DynamicOptionFiltersItems)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DynamicOptions) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

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

### GetSource

`func (o *DynamicOptions) GetSource() DynamicOptionsSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DynamicOptions) GetSourceOk() (*DynamicOptionsSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DynamicOptions) SetSource(v DynamicOptionsSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *DynamicOptions) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


