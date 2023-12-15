# ApplicationVersionAdoption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The application version number | [optional] 
**DeviceCount** | Pointer to **int32** | The number of devices that have adopted this version | [optional] 
**AdoptionPercentage** | Pointer to **float32** | The percentage of total devices that have adopted this version | [optional] 

## Methods

### NewApplicationVersionAdoption

`func NewApplicationVersionAdoption() *ApplicationVersionAdoption`

NewApplicationVersionAdoption instantiates a new ApplicationVersionAdoption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationVersionAdoptionWithDefaults

`func NewApplicationVersionAdoptionWithDefaults() *ApplicationVersionAdoption`

NewApplicationVersionAdoptionWithDefaults instantiates a new ApplicationVersionAdoption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApplicationVersionAdoption) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationVersionAdoption) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationVersionAdoption) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApplicationVersionAdoption) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDeviceCount

`func (o *ApplicationVersionAdoption) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *ApplicationVersionAdoption) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *ApplicationVersionAdoption) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *ApplicationVersionAdoption) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetAdoptionPercentage

`func (o *ApplicationVersionAdoption) GetAdoptionPercentage() float32`

GetAdoptionPercentage returns the AdoptionPercentage field if non-nil, zero value otherwise.

### GetAdoptionPercentageOk

`func (o *ApplicationVersionAdoption) GetAdoptionPercentageOk() (*float32, bool)`

GetAdoptionPercentageOk returns a tuple with the AdoptionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptionPercentage

`func (o *ApplicationVersionAdoption) SetAdoptionPercentage(v float32)`

SetAdoptionPercentage sets AdoptionPercentage field to given value.

### HasAdoptionPercentage

`func (o *ApplicationVersionAdoption) HasAdoptionPercentage() bool`

HasAdoptionPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


