# InternalApplicationVersionAdoption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The application version number | [optional] 
**Name** | Pointer to **string** | The application version name | [optional] 
**DeviceCount** | Pointer to **int32** | The number of devices that have adopted this version | [optional] 
**AdoptionPercentage** | Pointer to **float32** | The percentage of total devices that have adopted this version | [optional] 
**Supported** | Pointer to **bool** | Whether the application version is supported | [optional] 

## Methods

### NewInternalApplicationVersionAdoption

`func NewInternalApplicationVersionAdoption() *InternalApplicationVersionAdoption`

NewInternalApplicationVersionAdoption instantiates a new InternalApplicationVersionAdoption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalApplicationVersionAdoptionWithDefaults

`func NewInternalApplicationVersionAdoptionWithDefaults() *InternalApplicationVersionAdoption`

NewInternalApplicationVersionAdoptionWithDefaults instantiates a new InternalApplicationVersionAdoption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *InternalApplicationVersionAdoption) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InternalApplicationVersionAdoption) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InternalApplicationVersionAdoption) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InternalApplicationVersionAdoption) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *InternalApplicationVersionAdoption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalApplicationVersionAdoption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalApplicationVersionAdoption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalApplicationVersionAdoption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeviceCount

`func (o *InternalApplicationVersionAdoption) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *InternalApplicationVersionAdoption) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *InternalApplicationVersionAdoption) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *InternalApplicationVersionAdoption) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetAdoptionPercentage

`func (o *InternalApplicationVersionAdoption) GetAdoptionPercentage() float32`

GetAdoptionPercentage returns the AdoptionPercentage field if non-nil, zero value otherwise.

### GetAdoptionPercentageOk

`func (o *InternalApplicationVersionAdoption) GetAdoptionPercentageOk() (*float32, bool)`

GetAdoptionPercentageOk returns a tuple with the AdoptionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptionPercentage

`func (o *InternalApplicationVersionAdoption) SetAdoptionPercentage(v float32)`

SetAdoptionPercentage sets AdoptionPercentage field to given value.

### HasAdoptionPercentage

`func (o *InternalApplicationVersionAdoption) HasAdoptionPercentage() bool`

HasAdoptionPercentage returns a boolean if a field has been set.

### GetSupported

`func (o *InternalApplicationVersionAdoption) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *InternalApplicationVersionAdoption) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *InternalApplicationVersionAdoption) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *InternalApplicationVersionAdoption) HasSupported() bool`

HasSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


