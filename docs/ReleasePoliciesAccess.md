# ReleasePoliciesAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denied** | [**[]ReleasePoliciesAccessDenied**](ReleasePoliciesAccessDenied.md) |  | 
**Allowed** | [**[]ReleasePoliciesAccessAllowedRep**](ReleasePoliciesAccessAllowedRep.md) |  | 

## Methods

### NewReleasePoliciesAccess

`func NewReleasePoliciesAccess(denied []ReleasePoliciesAccessDenied, allowed []ReleasePoliciesAccessAllowedRep, ) *ReleasePoliciesAccess`

NewReleasePoliciesAccess instantiates a new ReleasePoliciesAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasePoliciesAccessWithDefaults

`func NewReleasePoliciesAccessWithDefaults() *ReleasePoliciesAccess`

NewReleasePoliciesAccessWithDefaults instantiates a new ReleasePoliciesAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenied

`func (o *ReleasePoliciesAccess) GetDenied() []ReleasePoliciesAccessDenied`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *ReleasePoliciesAccess) GetDeniedOk() (*[]ReleasePoliciesAccessDenied, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *ReleasePoliciesAccess) SetDenied(v []ReleasePoliciesAccessDenied)`

SetDenied sets Denied field to given value.


### GetAllowed

`func (o *ReleasePoliciesAccess) GetAllowed() []ReleasePoliciesAccessAllowedRep`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ReleasePoliciesAccess) GetAllowedOk() (*[]ReleasePoliciesAccessAllowedRep, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ReleasePoliciesAccess) SetAllowed(v []ReleasePoliciesAccessAllowedRep)`

SetAllowed sets Allowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


