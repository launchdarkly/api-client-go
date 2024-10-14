# LayerReservationRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentKey** | **string** | The key of the experiment | 
**FlagKey** | **string** | The key of the flag | 
**ReservationPercent** | **int32** | The percentage of traffic reserved for the experiment | 

## Methods

### NewLayerReservationRep

`func NewLayerReservationRep(experimentKey string, flagKey string, reservationPercent int32, ) *LayerReservationRep`

NewLayerReservationRep instantiates a new LayerReservationRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerReservationRepWithDefaults

`func NewLayerReservationRepWithDefaults() *LayerReservationRep`

NewLayerReservationRepWithDefaults instantiates a new LayerReservationRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentKey

`func (o *LayerReservationRep) GetExperimentKey() string`

GetExperimentKey returns the ExperimentKey field if non-nil, zero value otherwise.

### GetExperimentKeyOk

`func (o *LayerReservationRep) GetExperimentKeyOk() (*string, bool)`

GetExperimentKeyOk returns a tuple with the ExperimentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentKey

`func (o *LayerReservationRep) SetExperimentKey(v string)`

SetExperimentKey sets ExperimentKey field to given value.


### GetFlagKey

`func (o *LayerReservationRep) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *LayerReservationRep) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *LayerReservationRep) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetReservationPercent

`func (o *LayerReservationRep) GetReservationPercent() int32`

GetReservationPercent returns the ReservationPercent field if non-nil, zero value otherwise.

### GetReservationPercentOk

`func (o *LayerReservationRep) GetReservationPercentOk() (*int32, bool)`

GetReservationPercentOk returns a tuple with the ReservationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationPercent

`func (o *LayerReservationRep) SetReservationPercent(v int32)`

SetReservationPercent sets ReservationPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


