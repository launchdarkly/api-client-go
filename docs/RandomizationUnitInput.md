# RandomizationUnitInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RandomizationUnit** | **string** | The unit of randomization. | 
**Default** | **bool** | If true, any experiment iterations created within this project will default to using this randomization unit. A project can only have one default randomization unit. | 
**StandardRandomizationUnit** | **string** | One of LaunchDarkly&#39;s fixed set of standard randomization units. | 

## Methods

### NewRandomizationUnitInput

`func NewRandomizationUnitInput(randomizationUnit string, default_ bool, standardRandomizationUnit string, ) *RandomizationUnitInput`

NewRandomizationUnitInput instantiates a new RandomizationUnitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandomizationUnitInputWithDefaults

`func NewRandomizationUnitInputWithDefaults() *RandomizationUnitInput`

NewRandomizationUnitInputWithDefaults instantiates a new RandomizationUnitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRandomizationUnit

`func (o *RandomizationUnitInput) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *RandomizationUnitInput) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *RandomizationUnitInput) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.


### GetDefault

`func (o *RandomizationUnitInput) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RandomizationUnitInput) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RandomizationUnitInput) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetStandardRandomizationUnit

`func (o *RandomizationUnitInput) GetStandardRandomizationUnit() string`

GetStandardRandomizationUnit returns the StandardRandomizationUnit field if non-nil, zero value otherwise.

### GetStandardRandomizationUnitOk

`func (o *RandomizationUnitInput) GetStandardRandomizationUnitOk() (*string, bool)`

GetStandardRandomizationUnitOk returns a tuple with the StandardRandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardRandomizationUnit

`func (o *RandomizationUnitInput) SetStandardRandomizationUnit(v string)`

SetStandardRandomizationUnit sets StandardRandomizationUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


