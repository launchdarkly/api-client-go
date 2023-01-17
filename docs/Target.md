# Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | **[]string** | A list of the keys for targets that will receive this variation because of individual targeting | 
**Variation** | **int32** | The index, from the array of variations for this flag, of the variation to serve this list of targets | 
**ContextKind** | Pointer to **string** |  | [optional] 

## Methods

### NewTarget

`func NewTarget(values []string, variation int32, ) *Target`

NewTarget instantiates a new Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetWithDefaults

`func NewTargetWithDefaults() *Target`

NewTargetWithDefaults instantiates a new Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *Target) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Target) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Target) SetValues(v []string)`

SetValues sets Values field to given value.


### GetVariation

`func (o *Target) GetVariation() int32`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *Target) GetVariationOk() (*int32, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *Target) SetVariation(v int32)`

SetVariation sets Variation field to given value.


### GetContextKind

`func (o *Target) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *Target) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *Target) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.

### HasContextKind

`func (o *Target) HasContextKind() bool`

HasContextKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


