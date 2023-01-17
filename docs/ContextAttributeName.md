# ContextAttributeName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A context&#39;s attribute name. | 
**Weight** | **int32** | A relative estimate of the number of contexts seen recently that have an attribute with the associated name. | 
**Redacted** | Pointer to **bool** | Whether or not the attribute has one or more redacted values. | [optional] 

## Methods

### NewContextAttributeName

`func NewContextAttributeName(name string, weight int32, ) *ContextAttributeName`

NewContextAttributeName instantiates a new ContextAttributeName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextAttributeNameWithDefaults

`func NewContextAttributeNameWithDefaults() *ContextAttributeName`

NewContextAttributeNameWithDefaults instantiates a new ContextAttributeName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContextAttributeName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextAttributeName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextAttributeName) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *ContextAttributeName) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ContextAttributeName) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ContextAttributeName) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetRedacted

`func (o *ContextAttributeName) GetRedacted() bool`

GetRedacted returns the Redacted field if non-nil, zero value otherwise.

### GetRedactedOk

`func (o *ContextAttributeName) GetRedactedOk() (*bool, bool)`

GetRedactedOk returns a tuple with the Redacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedacted

`func (o *ContextAttributeName) SetRedacted(v bool)`

SetRedacted sets Redacted field to given value.

### HasRedacted

`func (o *ContextAttributeName) HasRedacted() bool`

HasRedacted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


