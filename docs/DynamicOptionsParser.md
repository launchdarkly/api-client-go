# DynamicOptionsParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionsItems** | Pointer to [**OptionsArray**](OptionsArray.md) |  | [optional] 
**OptionsPath** | Pointer to **string** |  | [optional] 

## Methods

### NewDynamicOptionsParser

`func NewDynamicOptionsParser() *DynamicOptionsParser`

NewDynamicOptionsParser instantiates a new DynamicOptionsParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicOptionsParserWithDefaults

`func NewDynamicOptionsParserWithDefaults() *DynamicOptionsParser`

NewDynamicOptionsParserWithDefaults instantiates a new DynamicOptionsParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionsItems

`func (o *DynamicOptionsParser) GetOptionsItems() OptionsArray`

GetOptionsItems returns the OptionsItems field if non-nil, zero value otherwise.

### GetOptionsItemsOk

`func (o *DynamicOptionsParser) GetOptionsItemsOk() (*OptionsArray, bool)`

GetOptionsItemsOk returns a tuple with the OptionsItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsItems

`func (o *DynamicOptionsParser) SetOptionsItems(v OptionsArray)`

SetOptionsItems sets OptionsItems field to given value.

### HasOptionsItems

`func (o *DynamicOptionsParser) HasOptionsItems() bool`

HasOptionsItems returns a boolean if a field has been set.

### GetOptionsPath

`func (o *DynamicOptionsParser) GetOptionsPath() string`

GetOptionsPath returns the OptionsPath field if non-nil, zero value otherwise.

### GetOptionsPathOk

`func (o *DynamicOptionsParser) GetOptionsPathOk() (*string, bool)`

GetOptionsPathOk returns a tuple with the OptionsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsPath

`func (o *DynamicOptionsParser) SetOptionsPath(v string)`

SetOptionsPath sets OptionsPath field to given value.

### HasOptionsPath

`func (o *DynamicOptionsParser) HasOptionsPath() bool`

HasOptionsPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


