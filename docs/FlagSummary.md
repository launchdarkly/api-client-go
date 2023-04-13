# FlagSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variations** | [**map[string]VariationSummary**](VariationSummary.md) |  | 
**Prerequisites** | **int32** | The number of prerequisites for this flag | 

## Methods

### NewFlagSummary

`func NewFlagSummary(variations map[string]VariationSummary, prerequisites int32, ) *FlagSummary`

NewFlagSummary instantiates a new FlagSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagSummaryWithDefaults

`func NewFlagSummaryWithDefaults() *FlagSummary`

NewFlagSummaryWithDefaults instantiates a new FlagSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariations

`func (o *FlagSummary) GetVariations() map[string]VariationSummary`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *FlagSummary) GetVariationsOk() (*map[string]VariationSummary, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *FlagSummary) SetVariations(v map[string]VariationSummary)`

SetVariations sets Variations field to given value.


### GetPrerequisites

`func (o *FlagSummary) GetPrerequisites() int32`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *FlagSummary) GetPrerequisitesOk() (*int32, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *FlagSummary) SetPrerequisites(v int32)`

SetPrerequisites sets Prerequisites field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


