# ViewsAccessRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denied** | [**[]ViewsAccessDenied**](ViewsAccessDenied.md) |  | 
**Allowed** | [**[]ViewsAccessAllowedRep**](ViewsAccessAllowedRep.md) |  | 

## Methods

### NewViewsAccessRep

`func NewViewsAccessRep(denied []ViewsAccessDenied, allowed []ViewsAccessAllowedRep, ) *ViewsAccessRep`

NewViewsAccessRep instantiates a new ViewsAccessRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsAccessRepWithDefaults

`func NewViewsAccessRepWithDefaults() *ViewsAccessRep`

NewViewsAccessRepWithDefaults instantiates a new ViewsAccessRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenied

`func (o *ViewsAccessRep) GetDenied() []ViewsAccessDenied`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *ViewsAccessRep) GetDeniedOk() (*[]ViewsAccessDenied, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *ViewsAccessRep) SetDenied(v []ViewsAccessDenied)`

SetDenied sets Denied field to given value.


### GetAllowed

`func (o *ViewsAccessRep) GetAllowed() []ViewsAccessAllowedRep`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ViewsAccessRep) GetAllowedOk() (*[]ViewsAccessAllowedRep, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ViewsAccessRep) SetAllowed(v []ViewsAccessAllowedRep)`

SetAllowed sets Allowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


