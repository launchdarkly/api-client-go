# FlagStatusesQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentKeys** | **[]string** | Keys of environments to fetch statuses for. Must include one key, can include up to 10 keys. | 
**FlagKeys** | **[]string** | Keys of flags to fetch statuses for. Must include one key, can include up to 100 keys. | 

## Methods

### NewFlagStatusesQuery

`func NewFlagStatusesQuery(environmentKeys []string, flagKeys []string, ) *FlagStatusesQuery`

NewFlagStatusesQuery instantiates a new FlagStatusesQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagStatusesQueryWithDefaults

`func NewFlagStatusesQueryWithDefaults() *FlagStatusesQuery`

NewFlagStatusesQueryWithDefaults instantiates a new FlagStatusesQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentKeys

`func (o *FlagStatusesQuery) GetEnvironmentKeys() []string`

GetEnvironmentKeys returns the EnvironmentKeys field if non-nil, zero value otherwise.

### GetEnvironmentKeysOk

`func (o *FlagStatusesQuery) GetEnvironmentKeysOk() (*[]string, bool)`

GetEnvironmentKeysOk returns a tuple with the EnvironmentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKeys

`func (o *FlagStatusesQuery) SetEnvironmentKeys(v []string)`

SetEnvironmentKeys sets EnvironmentKeys field to given value.


### GetFlagKeys

`func (o *FlagStatusesQuery) GetFlagKeys() []string`

GetFlagKeys returns the FlagKeys field if non-nil, zero value otherwise.

### GetFlagKeysOk

`func (o *FlagStatusesQuery) GetFlagKeysOk() (*[]string, bool)`

GetFlagKeysOk returns a tuple with the FlagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKeys

`func (o *FlagStatusesQuery) SetFlagKeys(v []string)`

SetFlagKeys sets FlagKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


