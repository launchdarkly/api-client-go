# AccessRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denied** | [**[]AccessDeniedRep**](AccessDeniedRep.md) |  | 

## Methods

### NewAccessRep

`func NewAccessRep(denied []AccessDeniedRep, ) *AccessRep`

NewAccessRep instantiates a new AccessRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRepWithDefaults

`func NewAccessRepWithDefaults() *AccessRep`

NewAccessRepWithDefaults instantiates a new AccessRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenied

`func (o *AccessRep) GetDenied() []AccessDeniedRep`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *AccessRep) GetDeniedOk() (*[]AccessDeniedRep, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *AccessRep) SetDenied(v []AccessDeniedRep)`

SetDenied sets Denied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


