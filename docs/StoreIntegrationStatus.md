# StoreIntegrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Whether the persistent store integration is fully synchronized with the LaunchDarkly environment, and the &lt;code&gt;lastSync&lt;/code&gt; occurred within a few minutes | [optional] 
**PotentiallyStale** | Pointer to **bool** | Whether the persistent store integration may not be fully synchronized with the LaunchDarkly environment. &lt;code&gt;true&lt;/code&gt; if the integration could be stale. | [optional] 
**LastSync** | Pointer to **int64** |  | [optional] 
**LastError** | Pointer to **int64** |  | [optional] 

## Methods

### NewStoreIntegrationStatus

`func NewStoreIntegrationStatus() *StoreIntegrationStatus`

NewStoreIntegrationStatus instantiates a new StoreIntegrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIntegrationStatusWithDefaults

`func NewStoreIntegrationStatusWithDefaults() *StoreIntegrationStatus`

NewStoreIntegrationStatusWithDefaults instantiates a new StoreIntegrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *StoreIntegrationStatus) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *StoreIntegrationStatus) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *StoreIntegrationStatus) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *StoreIntegrationStatus) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetPotentiallyStale

`func (o *StoreIntegrationStatus) GetPotentiallyStale() bool`

GetPotentiallyStale returns the PotentiallyStale field if non-nil, zero value otherwise.

### GetPotentiallyStaleOk

`func (o *StoreIntegrationStatus) GetPotentiallyStaleOk() (*bool, bool)`

GetPotentiallyStaleOk returns a tuple with the PotentiallyStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentiallyStale

`func (o *StoreIntegrationStatus) SetPotentiallyStale(v bool)`

SetPotentiallyStale sets PotentiallyStale field to given value.

### HasPotentiallyStale

`func (o *StoreIntegrationStatus) HasPotentiallyStale() bool`

HasPotentiallyStale returns a boolean if a field has been set.

### GetLastSync

`func (o *StoreIntegrationStatus) GetLastSync() int64`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *StoreIntegrationStatus) GetLastSyncOk() (*int64, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *StoreIntegrationStatus) SetLastSync(v int64)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *StoreIntegrationStatus) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastError

`func (o *StoreIntegrationStatus) GetLastError() int64`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *StoreIntegrationStatus) GetLastErrorOk() (*int64, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *StoreIntegrationStatus) SetLastError(v int64)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *StoreIntegrationStatus) HasLastError() bool`

HasLastError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


