# LayerConfigurationRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservations** | [**[]LayerReservationRep**](LayerReservationRep.md) | The experiment reservations for the layer | 

## Methods

### NewLayerConfigurationRep

`func NewLayerConfigurationRep(reservations []LayerReservationRep, ) *LayerConfigurationRep`

NewLayerConfigurationRep instantiates a new LayerConfigurationRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerConfigurationRepWithDefaults

`func NewLayerConfigurationRepWithDefaults() *LayerConfigurationRep`

NewLayerConfigurationRepWithDefaults instantiates a new LayerConfigurationRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservations

`func (o *LayerConfigurationRep) GetReservations() []LayerReservationRep`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *LayerConfigurationRep) GetReservationsOk() (*[]LayerReservationRep, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *LayerConfigurationRep) SetReservations(v []LayerReservationRep)`

SetReservations sets Reservations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


