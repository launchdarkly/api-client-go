# ContextAttributeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | A value for a context&#39;s attribute. | 
**Weight** | **int32** | A relative estimate of the number of contexts seen recently that have a matching value for a given attribute. | 

## Methods

### NewContextAttributeValue

`func NewContextAttributeValue(name interface{}, weight int32, ) *ContextAttributeValue`

NewContextAttributeValue instantiates a new ContextAttributeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextAttributeValueWithDefaults

`func NewContextAttributeValueWithDefaults() *ContextAttributeValue`

NewContextAttributeValueWithDefaults instantiates a new ContextAttributeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContextAttributeValue) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextAttributeValue) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextAttributeValue) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *ContextAttributeValue) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ContextAttributeValue) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetWeight

`func (o *ContextAttributeValue) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ContextAttributeValue) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ContextAttributeValue) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


