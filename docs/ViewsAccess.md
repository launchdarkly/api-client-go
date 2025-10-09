# ViewsAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denied** | [**[]ViewsAccessDenied**](ViewsAccessDenied.md) |  | 
**Allowed** | [**[]ViewsAccessAllowedRep**](ViewsAccessAllowedRep.md) |  | 

## Methods

### NewViewsAccess

`func NewViewsAccess(denied []ViewsAccessDenied, allowed []ViewsAccessAllowedRep, ) *ViewsAccess`

NewViewsAccess instantiates a new ViewsAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsAccessWithDefaults

`func NewViewsAccessWithDefaults() *ViewsAccess`

NewViewsAccessWithDefaults instantiates a new ViewsAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenied

`func (o *ViewsAccess) GetDenied() []ViewsAccessDenied`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *ViewsAccess) GetDeniedOk() (*[]ViewsAccessDenied, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *ViewsAccess) SetDenied(v []ViewsAccessDenied)`

SetDenied sets Denied field to given value.


### GetAllowed

`func (o *ViewsAccess) GetAllowed() []ViewsAccessAllowedRep`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ViewsAccess) GetAllowedOk() (*[]ViewsAccessAllowedRep, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ViewsAccess) SetAllowed(v []ViewsAccessAllowedRep)`

SetAllowed sets Allowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


