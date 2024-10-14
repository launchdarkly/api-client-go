# LayerSnapshotRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key of the layer the experiment was part of | 
**Name** | **string** | Layer name at the time this experiment iteration was stopped | 
**ReservationPercent** | **int32** | Percent of layer traffic that was reserved in the layer for this experiment iteration | 
**OtherReservationPercent** | **int32** | Percent of layer traffic that was reserved for other experiments in the same environment, when this experiment iteration was stopped | 

## Methods

### NewLayerSnapshotRep

`func NewLayerSnapshotRep(key string, name string, reservationPercent int32, otherReservationPercent int32, ) *LayerSnapshotRep`

NewLayerSnapshotRep instantiates a new LayerSnapshotRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerSnapshotRepWithDefaults

`func NewLayerSnapshotRepWithDefaults() *LayerSnapshotRep`

NewLayerSnapshotRepWithDefaults instantiates a new LayerSnapshotRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LayerSnapshotRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LayerSnapshotRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LayerSnapshotRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *LayerSnapshotRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayerSnapshotRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayerSnapshotRep) SetName(v string)`

SetName sets Name field to given value.


### GetReservationPercent

`func (o *LayerSnapshotRep) GetReservationPercent() int32`

GetReservationPercent returns the ReservationPercent field if non-nil, zero value otherwise.

### GetReservationPercentOk

`func (o *LayerSnapshotRep) GetReservationPercentOk() (*int32, bool)`

GetReservationPercentOk returns a tuple with the ReservationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationPercent

`func (o *LayerSnapshotRep) SetReservationPercent(v int32)`

SetReservationPercent sets ReservationPercent field to given value.


### GetOtherReservationPercent

`func (o *LayerSnapshotRep) GetOtherReservationPercent() int32`

GetOtherReservationPercent returns the OtherReservationPercent field if non-nil, zero value otherwise.

### GetOtherReservationPercentOk

`func (o *LayerSnapshotRep) GetOtherReservationPercentOk() (*int32, bool)`

GetOtherReservationPercentOk returns a tuple with the OtherReservationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherReservationPercent

`func (o *LayerSnapshotRep) SetOtherReservationPercent(v int32)`

SetOtherReservationPercent sets OtherReservationPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


