# IntegrationDeliveryConfigurationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Whether the latest sync succeeded | 
**PotentiallyStale** | **bool** | Whether the feature store may be missing flag data from a failed or in-progress sync | 
**LastSync** | Pointer to **int64** |  | [optional] 
**LastError** | Pointer to **int64** |  | [optional] 
**Errors** | Pointer to [**[]StoreIntegrationError**](StoreIntegrationError.md) | Recent failed sync responses. Successful (200) and in-progress (202) responses are omitted. Included only when fetching a single configuration by ID. | [optional] 

## Methods

### NewIntegrationDeliveryConfigurationStatus

`func NewIntegrationDeliveryConfigurationStatus(available bool, potentiallyStale bool, ) *IntegrationDeliveryConfigurationStatus`

NewIntegrationDeliveryConfigurationStatus instantiates a new IntegrationDeliveryConfigurationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationDeliveryConfigurationStatusWithDefaults

`func NewIntegrationDeliveryConfigurationStatusWithDefaults() *IntegrationDeliveryConfigurationStatus`

NewIntegrationDeliveryConfigurationStatusWithDefaults instantiates a new IntegrationDeliveryConfigurationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *IntegrationDeliveryConfigurationStatus) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *IntegrationDeliveryConfigurationStatus) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *IntegrationDeliveryConfigurationStatus) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetPotentiallyStale

`func (o *IntegrationDeliveryConfigurationStatus) GetPotentiallyStale() bool`

GetPotentiallyStale returns the PotentiallyStale field if non-nil, zero value otherwise.

### GetPotentiallyStaleOk

`func (o *IntegrationDeliveryConfigurationStatus) GetPotentiallyStaleOk() (*bool, bool)`

GetPotentiallyStaleOk returns a tuple with the PotentiallyStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentiallyStale

`func (o *IntegrationDeliveryConfigurationStatus) SetPotentiallyStale(v bool)`

SetPotentiallyStale sets PotentiallyStale field to given value.


### GetLastSync

`func (o *IntegrationDeliveryConfigurationStatus) GetLastSync() int64`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *IntegrationDeliveryConfigurationStatus) GetLastSyncOk() (*int64, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *IntegrationDeliveryConfigurationStatus) SetLastSync(v int64)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *IntegrationDeliveryConfigurationStatus) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastError

`func (o *IntegrationDeliveryConfigurationStatus) GetLastError() int64`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *IntegrationDeliveryConfigurationStatus) GetLastErrorOk() (*int64, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *IntegrationDeliveryConfigurationStatus) SetLastError(v int64)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *IntegrationDeliveryConfigurationStatus) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetErrors

`func (o *IntegrationDeliveryConfigurationStatus) GetErrors() []StoreIntegrationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *IntegrationDeliveryConfigurationStatus) GetErrorsOk() (*[]StoreIntegrationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *IntegrationDeliveryConfigurationStatus) SetErrors(v []StoreIntegrationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *IntegrationDeliveryConfigurationStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


