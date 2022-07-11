# Defaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnVariation** | **int32** | The index, from the array of variations for this flag, of the variation to serve by default when targeting is on. | 
**OffVariation** | **int32** | The index, from the array of variations for this flag, of the variation to serve by default when targeting is off. | 

## Methods

### NewDefaults

`func NewDefaults(onVariation int32, offVariation int32, ) *Defaults`

NewDefaults instantiates a new Defaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultsWithDefaults

`func NewDefaultsWithDefaults() *Defaults`

NewDefaultsWithDefaults instantiates a new Defaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnVariation

`func (o *Defaults) GetOnVariation() int32`

GetOnVariation returns the OnVariation field if non-nil, zero value otherwise.

### GetOnVariationOk

`func (o *Defaults) GetOnVariationOk() (*int32, bool)`

GetOnVariationOk returns a tuple with the OnVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnVariation

`func (o *Defaults) SetOnVariation(v int32)`

SetOnVariation sets OnVariation field to given value.


### GetOffVariation

`func (o *Defaults) GetOffVariation() int32`

GetOffVariation returns the OffVariation field if non-nil, zero value otherwise.

### GetOffVariationOk

`func (o *Defaults) GetOffVariationOk() (*int32, bool)`

GetOffVariationOk returns a tuple with the OffVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffVariation

`func (o *Defaults) SetOffVariation(v int32)`

SetOffVariation sets OffVariation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


