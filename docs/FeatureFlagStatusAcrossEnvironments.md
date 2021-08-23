# FeatureFlagStatusAcrossEnvironments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to [**map[string]FeatureFlagStatus**](FeatureFlagStatus.md) | Flag status for environment. | [optional] 
**Key** | Pointer to **string** | feature flag key | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewFeatureFlagStatusAcrossEnvironments

`func NewFeatureFlagStatusAcrossEnvironments() *FeatureFlagStatusAcrossEnvironments`

NewFeatureFlagStatusAcrossEnvironments instantiates a new FeatureFlagStatusAcrossEnvironments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagStatusAcrossEnvironmentsWithDefaults

`func NewFeatureFlagStatusAcrossEnvironmentsWithDefaults() *FeatureFlagStatusAcrossEnvironments`

NewFeatureFlagStatusAcrossEnvironmentsWithDefaults instantiates a new FeatureFlagStatusAcrossEnvironments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *FeatureFlagStatusAcrossEnvironments) GetEnvironments() map[string]FeatureFlagStatus`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *FeatureFlagStatusAcrossEnvironments) GetEnvironmentsOk() (*map[string]FeatureFlagStatus, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *FeatureFlagStatusAcrossEnvironments) SetEnvironments(v map[string]FeatureFlagStatus)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *FeatureFlagStatusAcrossEnvironments) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetKey

`func (o *FeatureFlagStatusAcrossEnvironments) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FeatureFlagStatusAcrossEnvironments) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FeatureFlagStatusAcrossEnvironments) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FeatureFlagStatusAcrossEnvironments) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLinks

`func (o *FeatureFlagStatusAcrossEnvironments) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FeatureFlagStatusAcrossEnvironments) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FeatureFlagStatusAcrossEnvironments) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FeatureFlagStatusAcrossEnvironments) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


