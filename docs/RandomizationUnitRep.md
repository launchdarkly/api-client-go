# RandomizationUnitRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RandomizationUnit** | Pointer to **string** | The unit of randomization. Defaults to user. | [optional] 
**Default** | Pointer to **bool** | Whether this randomization unit is the default for experiments | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** | The display name for the randomization unit, displayed in the LaunchDarkly user interface. | [optional] 

## Methods

### NewRandomizationUnitRep

`func NewRandomizationUnitRep() *RandomizationUnitRep`

NewRandomizationUnitRep instantiates a new RandomizationUnitRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandomizationUnitRepWithDefaults

`func NewRandomizationUnitRepWithDefaults() *RandomizationUnitRep`

NewRandomizationUnitRepWithDefaults instantiates a new RandomizationUnitRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRandomizationUnit

`func (o *RandomizationUnitRep) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *RandomizationUnitRep) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *RandomizationUnitRep) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *RandomizationUnitRep) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.

### GetDefault

`func (o *RandomizationUnitRep) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RandomizationUnitRep) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RandomizationUnitRep) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RandomizationUnitRep) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetHidden

`func (o *RandomizationUnitRep) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *RandomizationUnitRep) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *RandomizationUnitRep) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *RandomizationUnitRep) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetDisplayName

`func (o *RandomizationUnitRep) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RandomizationUnitRep) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RandomizationUnitRep) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RandomizationUnitRep) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


