# BooleanDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrueDisplayName** | Pointer to **string** | The display name for the true variation, displayed in the LaunchDarkly user interface | [optional] 
**FalseDisplayName** | Pointer to **string** | The display name for the false variation, displayed in the LaunchDarkly user interface | [optional] 
**TrueDescription** | Pointer to **string** | The description for the true variation | [optional] 
**FalseDescription** | Pointer to **string** | The description for the false variation | [optional] 
**OnVariation** | Pointer to **int32** | The variation index of the flag variation to use for the default targeting behavior when a flag&#39;s targeting is on and the target did not match any rules | [optional] 
**OffVariation** | Pointer to **int32** | The variation index of the flag variation to use for the default targeting behavior when a flag&#39;s targeting is off | [optional] 

## Methods

### NewBooleanDefaults

`func NewBooleanDefaults() *BooleanDefaults`

NewBooleanDefaults instantiates a new BooleanDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooleanDefaultsWithDefaults

`func NewBooleanDefaultsWithDefaults() *BooleanDefaults`

NewBooleanDefaultsWithDefaults instantiates a new BooleanDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrueDisplayName

`func (o *BooleanDefaults) GetTrueDisplayName() string`

GetTrueDisplayName returns the TrueDisplayName field if non-nil, zero value otherwise.

### GetTrueDisplayNameOk

`func (o *BooleanDefaults) GetTrueDisplayNameOk() (*string, bool)`

GetTrueDisplayNameOk returns a tuple with the TrueDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueDisplayName

`func (o *BooleanDefaults) SetTrueDisplayName(v string)`

SetTrueDisplayName sets TrueDisplayName field to given value.

### HasTrueDisplayName

`func (o *BooleanDefaults) HasTrueDisplayName() bool`

HasTrueDisplayName returns a boolean if a field has been set.

### GetFalseDisplayName

`func (o *BooleanDefaults) GetFalseDisplayName() string`

GetFalseDisplayName returns the FalseDisplayName field if non-nil, zero value otherwise.

### GetFalseDisplayNameOk

`func (o *BooleanDefaults) GetFalseDisplayNameOk() (*string, bool)`

GetFalseDisplayNameOk returns a tuple with the FalseDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseDisplayName

`func (o *BooleanDefaults) SetFalseDisplayName(v string)`

SetFalseDisplayName sets FalseDisplayName field to given value.

### HasFalseDisplayName

`func (o *BooleanDefaults) HasFalseDisplayName() bool`

HasFalseDisplayName returns a boolean if a field has been set.

### GetTrueDescription

`func (o *BooleanDefaults) GetTrueDescription() string`

GetTrueDescription returns the TrueDescription field if non-nil, zero value otherwise.

### GetTrueDescriptionOk

`func (o *BooleanDefaults) GetTrueDescriptionOk() (*string, bool)`

GetTrueDescriptionOk returns a tuple with the TrueDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueDescription

`func (o *BooleanDefaults) SetTrueDescription(v string)`

SetTrueDescription sets TrueDescription field to given value.

### HasTrueDescription

`func (o *BooleanDefaults) HasTrueDescription() bool`

HasTrueDescription returns a boolean if a field has been set.

### GetFalseDescription

`func (o *BooleanDefaults) GetFalseDescription() string`

GetFalseDescription returns the FalseDescription field if non-nil, zero value otherwise.

### GetFalseDescriptionOk

`func (o *BooleanDefaults) GetFalseDescriptionOk() (*string, bool)`

GetFalseDescriptionOk returns a tuple with the FalseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseDescription

`func (o *BooleanDefaults) SetFalseDescription(v string)`

SetFalseDescription sets FalseDescription field to given value.

### HasFalseDescription

`func (o *BooleanDefaults) HasFalseDescription() bool`

HasFalseDescription returns a boolean if a field has been set.

### GetOnVariation

`func (o *BooleanDefaults) GetOnVariation() int32`

GetOnVariation returns the OnVariation field if non-nil, zero value otherwise.

### GetOnVariationOk

`func (o *BooleanDefaults) GetOnVariationOk() (*int32, bool)`

GetOnVariationOk returns a tuple with the OnVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnVariation

`func (o *BooleanDefaults) SetOnVariation(v int32)`

SetOnVariation sets OnVariation field to given value.

### HasOnVariation

`func (o *BooleanDefaults) HasOnVariation() bool`

HasOnVariation returns a boolean if a field has been set.

### GetOffVariation

`func (o *BooleanDefaults) GetOffVariation() int32`

GetOffVariation returns the OffVariation field if non-nil, zero value otherwise.

### GetOffVariationOk

`func (o *BooleanDefaults) GetOffVariationOk() (*int32, bool)`

GetOffVariationOk returns a tuple with the OffVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffVariation

`func (o *BooleanDefaults) SetOffVariation(v int32)`

SetOffVariation sets OffVariation field to given value.

### HasOffVariation

`func (o *BooleanDefaults) HasOffVariation() bool`

HasOffVariation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


