# AiConfigsMetricEventDefaultRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Whether to disable defaulting missing unit events when calculating results. Defaults to false | [optional] 
**Value** | Pointer to **float32** | The default value applied to missing unit events. Set to 0 when &lt;code&gt;disabled&lt;/code&gt; is false. No other values are currently supported. | [optional] 

## Methods

### NewAiConfigsMetricEventDefaultRep

`func NewAiConfigsMetricEventDefaultRep() *AiConfigsMetricEventDefaultRep`

NewAiConfigsMetricEventDefaultRep instantiates a new AiConfigsMetricEventDefaultRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiConfigsMetricEventDefaultRepWithDefaults

`func NewAiConfigsMetricEventDefaultRepWithDefaults() *AiConfigsMetricEventDefaultRep`

NewAiConfigsMetricEventDefaultRepWithDefaults instantiates a new AiConfigsMetricEventDefaultRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *AiConfigsMetricEventDefaultRep) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AiConfigsMetricEventDefaultRep) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AiConfigsMetricEventDefaultRep) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AiConfigsMetricEventDefaultRep) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetValue

`func (o *AiConfigsMetricEventDefaultRep) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AiConfigsMetricEventDefaultRep) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AiConfigsMetricEventDefaultRep) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *AiConfigsMetricEventDefaultRep) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


