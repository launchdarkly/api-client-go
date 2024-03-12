# FailureReasonRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | The attribute that failed validation | 
**Reason** | **string** | The reason the attribute failed validation | 

## Methods

### NewFailureReasonRep

`func NewFailureReasonRep(attribute string, reason string, ) *FailureReasonRep`

NewFailureReasonRep instantiates a new FailureReasonRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureReasonRepWithDefaults

`func NewFailureReasonRepWithDefaults() *FailureReasonRep`

NewFailureReasonRepWithDefaults instantiates a new FailureReasonRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *FailureReasonRep) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *FailureReasonRep) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *FailureReasonRep) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetReason

`func (o *FailureReasonRep) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FailureReasonRep) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FailureReasonRep) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


