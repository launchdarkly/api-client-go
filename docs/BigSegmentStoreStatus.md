# BigSegmentStoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Whether the persistent store integration is fully synchronized with the LaunchDarkly environment, and the &lt;code&gt;lastSync&lt;/code&gt; occurred within a few minutes | [optional] 
**PotentiallyStale** | Pointer to **bool** | Whether the persistent store integration may not be fully synchronized with the LaunchDarkly environment. &lt;code&gt;true&lt;/code&gt; if the integration could be stale. | [optional] 
**LastSync** | Pointer to **int64** |  | [optional] 
**LastError** | Pointer to **int64** |  | [optional] 
**Errors** | Pointer to [**[]StoreIntegrationError**](StoreIntegrationError.md) |  | [optional] 

## Methods

### NewBigSegmentStoreStatus

`func NewBigSegmentStoreStatus() *BigSegmentStoreStatus`

NewBigSegmentStoreStatus instantiates a new BigSegmentStoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigSegmentStoreStatusWithDefaults

`func NewBigSegmentStoreStatusWithDefaults() *BigSegmentStoreStatus`

NewBigSegmentStoreStatusWithDefaults instantiates a new BigSegmentStoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *BigSegmentStoreStatus) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *BigSegmentStoreStatus) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *BigSegmentStoreStatus) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *BigSegmentStoreStatus) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetPotentiallyStale

`func (o *BigSegmentStoreStatus) GetPotentiallyStale() bool`

GetPotentiallyStale returns the PotentiallyStale field if non-nil, zero value otherwise.

### GetPotentiallyStaleOk

`func (o *BigSegmentStoreStatus) GetPotentiallyStaleOk() (*bool, bool)`

GetPotentiallyStaleOk returns a tuple with the PotentiallyStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentiallyStale

`func (o *BigSegmentStoreStatus) SetPotentiallyStale(v bool)`

SetPotentiallyStale sets PotentiallyStale field to given value.

### HasPotentiallyStale

`func (o *BigSegmentStoreStatus) HasPotentiallyStale() bool`

HasPotentiallyStale returns a boolean if a field has been set.

### GetLastSync

`func (o *BigSegmentStoreStatus) GetLastSync() int64`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *BigSegmentStoreStatus) GetLastSyncOk() (*int64, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *BigSegmentStoreStatus) SetLastSync(v int64)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *BigSegmentStoreStatus) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastError

`func (o *BigSegmentStoreStatus) GetLastError() int64`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *BigSegmentStoreStatus) GetLastErrorOk() (*int64, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *BigSegmentStoreStatus) SetLastError(v int64)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *BigSegmentStoreStatus) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetErrors

`func (o *BigSegmentStoreStatus) GetErrors() []StoreIntegrationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BigSegmentStoreStatus) GetErrorsOk() (*[]StoreIntegrationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BigSegmentStoreStatus) SetErrors(v []StoreIntegrationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BigSegmentStoreStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


