# FeatureFlagStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Status of the flag | 
**LastRequested** | Pointer to **time.Time** | Timestamp of last time flag was requested | [optional] 
**Default** | Pointer to **interface{}** | Default value seen from code | [optional] 

## Methods

### NewFeatureFlagStatus

`func NewFeatureFlagStatus(name string, ) *FeatureFlagStatus`

NewFeatureFlagStatus instantiates a new FeatureFlagStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagStatusWithDefaults

`func NewFeatureFlagStatusWithDefaults() *FeatureFlagStatus`

NewFeatureFlagStatusWithDefaults instantiates a new FeatureFlagStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureFlagStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureFlagStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureFlagStatus) SetName(v string)`

SetName sets Name field to given value.


### GetLastRequested

`func (o *FeatureFlagStatus) GetLastRequested() time.Time`

GetLastRequested returns the LastRequested field if non-nil, zero value otherwise.

### GetLastRequestedOk

`func (o *FeatureFlagStatus) GetLastRequestedOk() (*time.Time, bool)`

GetLastRequestedOk returns a tuple with the LastRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRequested

`func (o *FeatureFlagStatus) SetLastRequested(v time.Time)`

SetLastRequested sets LastRequested field to given value.

### HasLastRequested

`func (o *FeatureFlagStatus) HasLastRequested() bool`

HasLastRequested returns a boolean if a field has been set.

### GetDefault

`func (o *FeatureFlagStatus) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FeatureFlagStatus) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FeatureFlagStatus) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FeatureFlagStatus) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *FeatureFlagStatus) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *FeatureFlagStatus) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


