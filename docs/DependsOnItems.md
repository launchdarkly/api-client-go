# DependsOnItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]ConditionsItems**](ConditionsItems.md) |  | [optional] 
**VariableKey** | Pointer to **string** |  | [optional] 
**VariableLocation** | Pointer to **string** |  | [optional] 

## Methods

### NewDependsOnItems

`func NewDependsOnItems() *DependsOnItems`

NewDependsOnItems instantiates a new DependsOnItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependsOnItemsWithDefaults

`func NewDependsOnItemsWithDefaults() *DependsOnItems`

NewDependsOnItemsWithDefaults instantiates a new DependsOnItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DependsOnItems) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DependsOnItems) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DependsOnItems) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DependsOnItems) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetConditions

`func (o *DependsOnItems) GetConditions() []ConditionsItems`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DependsOnItems) GetConditionsOk() (*[]ConditionsItems, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DependsOnItems) SetConditions(v []ConditionsItems)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DependsOnItems) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetVariableKey

`func (o *DependsOnItems) GetVariableKey() string`

GetVariableKey returns the VariableKey field if non-nil, zero value otherwise.

### GetVariableKeyOk

`func (o *DependsOnItems) GetVariableKeyOk() (*string, bool)`

GetVariableKeyOk returns a tuple with the VariableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableKey

`func (o *DependsOnItems) SetVariableKey(v string)`

SetVariableKey sets VariableKey field to given value.

### HasVariableKey

`func (o *DependsOnItems) HasVariableKey() bool`

HasVariableKey returns a boolean if a field has been set.

### GetVariableLocation

`func (o *DependsOnItems) GetVariableLocation() string`

GetVariableLocation returns the VariableLocation field if non-nil, zero value otherwise.

### GetVariableLocationOk

`func (o *DependsOnItems) GetVariableLocationOk() (*string, bool)`

GetVariableLocationOk returns a tuple with the VariableLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableLocation

`func (o *DependsOnItems) SetVariableLocation(v string)`

SetVariableLocation sets VariableLocation field to given value.

### HasVariableLocation

`func (o *DependsOnItems) HasVariableLocation() bool`

HasVariableLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


