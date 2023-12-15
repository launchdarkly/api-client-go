# ApplicationVersionAdoptionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalDevices** | Pointer to **int32** | Total devices with any application version installed | [optional] 
**Items** | Pointer to [**[]ApplicationVersionAdoption**](ApplicationVersionAdoption.md) | List of application versions and their adoption data | [optional] 

## Methods

### NewApplicationVersionAdoptionCollection

`func NewApplicationVersionAdoptionCollection() *ApplicationVersionAdoptionCollection`

NewApplicationVersionAdoptionCollection instantiates a new ApplicationVersionAdoptionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationVersionAdoptionCollectionWithDefaults

`func NewApplicationVersionAdoptionCollectionWithDefaults() *ApplicationVersionAdoptionCollection`

NewApplicationVersionAdoptionCollectionWithDefaults instantiates a new ApplicationVersionAdoptionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalDevices

`func (o *ApplicationVersionAdoptionCollection) GetTotalDevices() int32`

GetTotalDevices returns the TotalDevices field if non-nil, zero value otherwise.

### GetTotalDevicesOk

`func (o *ApplicationVersionAdoptionCollection) GetTotalDevicesOk() (*int32, bool)`

GetTotalDevicesOk returns a tuple with the TotalDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDevices

`func (o *ApplicationVersionAdoptionCollection) SetTotalDevices(v int32)`

SetTotalDevices sets TotalDevices field to given value.

### HasTotalDevices

`func (o *ApplicationVersionAdoptionCollection) HasTotalDevices() bool`

HasTotalDevices returns a boolean if a field has been set.

### GetItems

`func (o *ApplicationVersionAdoptionCollection) GetItems() []ApplicationVersionAdoption`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApplicationVersionAdoptionCollection) GetItemsOk() (*[]ApplicationVersionAdoption, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApplicationVersionAdoptionCollection) SetItems(v []ApplicationVersionAdoption)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApplicationVersionAdoptionCollection) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


