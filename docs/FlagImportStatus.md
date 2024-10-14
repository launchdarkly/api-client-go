# FlagImportStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The current status of the import integrations related import job | [optional] 
**LastImport** | Pointer to **int64** |  | [optional] 
**LastError** | Pointer to **int64** |  | [optional] 
**Errors** | Pointer to [**[]StatusResponse**](StatusResponse.md) |  | [optional] 

## Methods

### NewFlagImportStatus

`func NewFlagImportStatus() *FlagImportStatus`

NewFlagImportStatus instantiates a new FlagImportStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagImportStatusWithDefaults

`func NewFlagImportStatusWithDefaults() *FlagImportStatus`

NewFlagImportStatusWithDefaults instantiates a new FlagImportStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FlagImportStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlagImportStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlagImportStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlagImportStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastImport

`func (o *FlagImportStatus) GetLastImport() int64`

GetLastImport returns the LastImport field if non-nil, zero value otherwise.

### GetLastImportOk

`func (o *FlagImportStatus) GetLastImportOk() (*int64, bool)`

GetLastImportOk returns a tuple with the LastImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastImport

`func (o *FlagImportStatus) SetLastImport(v int64)`

SetLastImport sets LastImport field to given value.

### HasLastImport

`func (o *FlagImportStatus) HasLastImport() bool`

HasLastImport returns a boolean if a field has been set.

### GetLastError

`func (o *FlagImportStatus) GetLastError() int64`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *FlagImportStatus) GetLastErrorOk() (*int64, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *FlagImportStatus) SetLastError(v int64)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *FlagImportStatus) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetErrors

`func (o *FlagImportStatus) GetErrors() []StatusResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FlagImportStatus) GetErrorsOk() (*[]StatusResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FlagImportStatus) SetErrors(v []StatusResponse)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FlagImportStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


