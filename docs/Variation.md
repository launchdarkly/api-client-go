# Variation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Value** | **interface{}** | The value of the variation. For boolean flags, this must be &lt;code&gt;true&lt;/code&gt; or &lt;code&gt;false&lt;/code&gt;. For multivariate flags, this may be a string, number, or JSON object. | 
**Description** | Pointer to **string** | Description of the variation | [optional] 
**Name** | Pointer to **string** | A human-friendly name for the variation | [optional] 

## Methods

### NewVariation

`func NewVariation(value interface{}, ) *Variation`

NewVariation instantiates a new Variation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationWithDefaults

`func NewVariationWithDefaults() *Variation`

NewVariationWithDefaults instantiates a new Variation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Variation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Variation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Variation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Variation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *Variation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Variation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Variation) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Variation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Variation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDescription

`func (o *Variation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Variation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Variation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Variation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Variation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Variation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Variation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Variation) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


