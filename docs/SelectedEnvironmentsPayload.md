# SelectedEnvironmentsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedEnvironmentKey** | **string** | The key for the selected environment in the project | 
**EnvironmentKeys** | **[]string** | An ordered list of keys for selected environments in this project | 

## Methods

### NewSelectedEnvironmentsPayload

`func NewSelectedEnvironmentsPayload(selectedEnvironmentKey string, environmentKeys []string, ) *SelectedEnvironmentsPayload`

NewSelectedEnvironmentsPayload instantiates a new SelectedEnvironmentsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedEnvironmentsPayloadWithDefaults

`func NewSelectedEnvironmentsPayloadWithDefaults() *SelectedEnvironmentsPayload`

NewSelectedEnvironmentsPayloadWithDefaults instantiates a new SelectedEnvironmentsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedEnvironmentKey

`func (o *SelectedEnvironmentsPayload) GetSelectedEnvironmentKey() string`

GetSelectedEnvironmentKey returns the SelectedEnvironmentKey field if non-nil, zero value otherwise.

### GetSelectedEnvironmentKeyOk

`func (o *SelectedEnvironmentsPayload) GetSelectedEnvironmentKeyOk() (*string, bool)`

GetSelectedEnvironmentKeyOk returns a tuple with the SelectedEnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedEnvironmentKey

`func (o *SelectedEnvironmentsPayload) SetSelectedEnvironmentKey(v string)`

SetSelectedEnvironmentKey sets SelectedEnvironmentKey field to given value.


### GetEnvironmentKeys

`func (o *SelectedEnvironmentsPayload) GetEnvironmentKeys() []string`

GetEnvironmentKeys returns the EnvironmentKeys field if non-nil, zero value otherwise.

### GetEnvironmentKeysOk

`func (o *SelectedEnvironmentsPayload) GetEnvironmentKeysOk() (*[]string, bool)`

GetEnvironmentKeysOk returns a tuple with the EnvironmentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKeys

`func (o *SelectedEnvironmentsPayload) SetEnvironmentKeys(v []string)`

SetEnvironmentKeys sets EnvironmentKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


