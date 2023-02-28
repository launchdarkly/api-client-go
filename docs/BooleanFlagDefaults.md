# BooleanFlagDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrueDisplayName** | **string** | The display name for the true variation, displayed in the LaunchDarkly user interface | 
**FalseDisplayName** | **string** | The display name for the false variation, displayed in the LaunchDarkly user interface | 
**TrueDescription** | **string** | The description for the true variation | 
**FalseDescription** | **string** | The description for the false variation | 
**OnVariation** | **int32** | The variation index of the flag variation to use for the default targeting behavior when a flag&#39;s targeting is on and the target did not match any rules | 
**OffVariation** | **int32** | The variation index of the flag variation to use for the default targeting behavior when a flag&#39;s targeting is off | 

## Methods

### NewBooleanFlagDefaults

`func NewBooleanFlagDefaults(trueDisplayName string, falseDisplayName string, trueDescription string, falseDescription string, onVariation int32, offVariation int32, ) *BooleanFlagDefaults`

NewBooleanFlagDefaults instantiates a new BooleanFlagDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooleanFlagDefaultsWithDefaults

`func NewBooleanFlagDefaultsWithDefaults() *BooleanFlagDefaults`

NewBooleanFlagDefaultsWithDefaults instantiates a new BooleanFlagDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrueDisplayName

`func (o *BooleanFlagDefaults) GetTrueDisplayName() string`

GetTrueDisplayName returns the TrueDisplayName field if non-nil, zero value otherwise.

### GetTrueDisplayNameOk

`func (o *BooleanFlagDefaults) GetTrueDisplayNameOk() (*string, bool)`

GetTrueDisplayNameOk returns a tuple with the TrueDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueDisplayName

`func (o *BooleanFlagDefaults) SetTrueDisplayName(v string)`

SetTrueDisplayName sets TrueDisplayName field to given value.


### GetFalseDisplayName

`func (o *BooleanFlagDefaults) GetFalseDisplayName() string`

GetFalseDisplayName returns the FalseDisplayName field if non-nil, zero value otherwise.

### GetFalseDisplayNameOk

`func (o *BooleanFlagDefaults) GetFalseDisplayNameOk() (*string, bool)`

GetFalseDisplayNameOk returns a tuple with the FalseDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseDisplayName

`func (o *BooleanFlagDefaults) SetFalseDisplayName(v string)`

SetFalseDisplayName sets FalseDisplayName field to given value.


### GetTrueDescription

`func (o *BooleanFlagDefaults) GetTrueDescription() string`

GetTrueDescription returns the TrueDescription field if non-nil, zero value otherwise.

### GetTrueDescriptionOk

`func (o *BooleanFlagDefaults) GetTrueDescriptionOk() (*string, bool)`

GetTrueDescriptionOk returns a tuple with the TrueDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueDescription

`func (o *BooleanFlagDefaults) SetTrueDescription(v string)`

SetTrueDescription sets TrueDescription field to given value.


### GetFalseDescription

`func (o *BooleanFlagDefaults) GetFalseDescription() string`

GetFalseDescription returns the FalseDescription field if non-nil, zero value otherwise.

### GetFalseDescriptionOk

`func (o *BooleanFlagDefaults) GetFalseDescriptionOk() (*string, bool)`

GetFalseDescriptionOk returns a tuple with the FalseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseDescription

`func (o *BooleanFlagDefaults) SetFalseDescription(v string)`

SetFalseDescription sets FalseDescription field to given value.


### GetOnVariation

`func (o *BooleanFlagDefaults) GetOnVariation() int32`

GetOnVariation returns the OnVariation field if non-nil, zero value otherwise.

### GetOnVariationOk

`func (o *BooleanFlagDefaults) GetOnVariationOk() (*int32, bool)`

GetOnVariationOk returns a tuple with the OnVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnVariation

`func (o *BooleanFlagDefaults) SetOnVariation(v int32)`

SetOnVariation sets OnVariation field to given value.


### GetOffVariation

`func (o *BooleanFlagDefaults) GetOffVariation() int32`

GetOffVariation returns the OffVariation field if non-nil, zero value otherwise.

### GetOffVariationOk

`func (o *BooleanFlagDefaults) GetOffVariationOk() (*int32, bool)`

GetOffVariationOk returns a tuple with the OffVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffVariation

`func (o *BooleanFlagDefaults) SetOffVariation(v int32)`

SetOffVariation sets OffVariation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


