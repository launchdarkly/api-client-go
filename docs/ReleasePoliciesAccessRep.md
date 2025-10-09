# ReleasePoliciesAccessRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denied** | [**[]ReleasePoliciesAccessDenied**](ReleasePoliciesAccessDenied.md) |  | 
**Allowed** | [**[]ReleasePoliciesAccessAllowedRep**](ReleasePoliciesAccessAllowedRep.md) |  | 

## Methods

### NewReleasePoliciesAccessRep

`func NewReleasePoliciesAccessRep(denied []ReleasePoliciesAccessDenied, allowed []ReleasePoliciesAccessAllowedRep, ) *ReleasePoliciesAccessRep`

NewReleasePoliciesAccessRep instantiates a new ReleasePoliciesAccessRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasePoliciesAccessRepWithDefaults

`func NewReleasePoliciesAccessRepWithDefaults() *ReleasePoliciesAccessRep`

NewReleasePoliciesAccessRepWithDefaults instantiates a new ReleasePoliciesAccessRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenied

`func (o *ReleasePoliciesAccessRep) GetDenied() []ReleasePoliciesAccessDenied`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *ReleasePoliciesAccessRep) GetDeniedOk() (*[]ReleasePoliciesAccessDenied, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *ReleasePoliciesAccessRep) SetDenied(v []ReleasePoliciesAccessDenied)`

SetDenied sets Denied field to given value.


### GetAllowed

`func (o *ReleasePoliciesAccessRep) GetAllowed() []ReleasePoliciesAccessAllowedRep`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ReleasePoliciesAccessRep) GetAllowedOk() (*[]ReleasePoliciesAccessAllowedRep, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ReleasePoliciesAccessRep) SetAllowed(v []ReleasePoliciesAccessAllowedRep)`

SetAllowed sets Allowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


