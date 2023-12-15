# FormVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedValues** | Pointer to **[]string** |  | [optional] 
**DefaultValue** | Pointer to **interface{}** |  | [optional] 
**DependsOn** | Pointer to [**[]DependsOnItems**](DependsOnItems.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DynamicOptions** | Pointer to [**DynamicOptions**](DynamicOptions.md) |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsOptional** | Pointer to **bool** |  | [optional] 
**IsSecret** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewFormVariable

`func NewFormVariable() *FormVariable`

NewFormVariable instantiates a new FormVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormVariableWithDefaults

`func NewFormVariableWithDefaults() *FormVariable`

NewFormVariableWithDefaults instantiates a new FormVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedValues

`func (o *FormVariable) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *FormVariable) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *FormVariable) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *FormVariable) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *FormVariable) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FormVariable) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FormVariable) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *FormVariable) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *FormVariable) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *FormVariable) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDependsOn

`func (o *FormVariable) GetDependsOn() []DependsOnItems`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *FormVariable) GetDependsOnOk() (*[]DependsOnItems, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *FormVariable) SetDependsOn(v []DependsOnItems)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *FormVariable) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetDescription

`func (o *FormVariable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormVariable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormVariable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FormVariable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDynamicOptions

`func (o *FormVariable) GetDynamicOptions() DynamicOptions`

GetDynamicOptions returns the DynamicOptions field if non-nil, zero value otherwise.

### GetDynamicOptionsOk

`func (o *FormVariable) GetDynamicOptionsOk() (*DynamicOptions, bool)`

GetDynamicOptionsOk returns a tuple with the DynamicOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicOptions

`func (o *FormVariable) SetDynamicOptions(v DynamicOptions)`

SetDynamicOptions sets DynamicOptions field to given value.

### HasDynamicOptions

`func (o *FormVariable) HasDynamicOptions() bool`

HasDynamicOptions returns a boolean if a field has been set.

### GetIsHidden

`func (o *FormVariable) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *FormVariable) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *FormVariable) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *FormVariable) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsOptional

`func (o *FormVariable) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *FormVariable) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *FormVariable) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *FormVariable) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### GetIsSecret

`func (o *FormVariable) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *FormVariable) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *FormVariable) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *FormVariable) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### GetKey

`func (o *FormVariable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormVariable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormVariable) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormVariable) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *FormVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormVariable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormVariable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaceholder

`func (o *FormVariable) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *FormVariable) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *FormVariable) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *FormVariable) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetType

`func (o *FormVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormVariable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormVariable) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


