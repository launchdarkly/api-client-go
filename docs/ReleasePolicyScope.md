# ReleasePolicyScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentKeys** | Pointer to **[]string** | List of environment keys this policy applies to | [optional] 
**FlagTagKeys** | Pointer to **[]string** | List of flag tag keys this policy applies to | [optional] 

## Methods

### NewReleasePolicyScope

`func NewReleasePolicyScope() *ReleasePolicyScope`

NewReleasePolicyScope instantiates a new ReleasePolicyScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasePolicyScopeWithDefaults

`func NewReleasePolicyScopeWithDefaults() *ReleasePolicyScope`

NewReleasePolicyScopeWithDefaults instantiates a new ReleasePolicyScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentKeys

`func (o *ReleasePolicyScope) GetEnvironmentKeys() []string`

GetEnvironmentKeys returns the EnvironmentKeys field if non-nil, zero value otherwise.

### GetEnvironmentKeysOk

`func (o *ReleasePolicyScope) GetEnvironmentKeysOk() (*[]string, bool)`

GetEnvironmentKeysOk returns a tuple with the EnvironmentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKeys

`func (o *ReleasePolicyScope) SetEnvironmentKeys(v []string)`

SetEnvironmentKeys sets EnvironmentKeys field to given value.

### HasEnvironmentKeys

`func (o *ReleasePolicyScope) HasEnvironmentKeys() bool`

HasEnvironmentKeys returns a boolean if a field has been set.

### GetFlagTagKeys

`func (o *ReleasePolicyScope) GetFlagTagKeys() []string`

GetFlagTagKeys returns the FlagTagKeys field if non-nil, zero value otherwise.

### GetFlagTagKeysOk

`func (o *ReleasePolicyScope) GetFlagTagKeysOk() (*[]string, bool)`

GetFlagTagKeysOk returns a tuple with the FlagTagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagTagKeys

`func (o *ReleasePolicyScope) SetFlagTagKeys(v []string)`

SetFlagTagKeys sets FlagTagKeys field to given value.

### HasFlagTagKeys

`func (o *ReleasePolicyScope) HasFlagTagKeys() bool`

HasFlagTagKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


