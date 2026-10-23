# EnvironmentReleasesAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denied** | [**[]EnvironmentReleasesAccessDenied**](EnvironmentReleasesAccessDenied.md) |  | 
**Allowed** | [**[]EnvironmentReleasesAccessAllowedRep**](EnvironmentReleasesAccessAllowedRep.md) |  | 

## Methods

### NewEnvironmentReleasesAccess

`func NewEnvironmentReleasesAccess(denied []EnvironmentReleasesAccessDenied, allowed []EnvironmentReleasesAccessAllowedRep, ) *EnvironmentReleasesAccess`

NewEnvironmentReleasesAccess instantiates a new EnvironmentReleasesAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasesAccessWithDefaults

`func NewEnvironmentReleasesAccessWithDefaults() *EnvironmentReleasesAccess`

NewEnvironmentReleasesAccessWithDefaults instantiates a new EnvironmentReleasesAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenied

`func (o *EnvironmentReleasesAccess) GetDenied() []EnvironmentReleasesAccessDenied`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *EnvironmentReleasesAccess) GetDeniedOk() (*[]EnvironmentReleasesAccessDenied, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *EnvironmentReleasesAccess) SetDenied(v []EnvironmentReleasesAccessDenied)`

SetDenied sets Denied field to given value.


### GetAllowed

`func (o *EnvironmentReleasesAccess) GetAllowed() []EnvironmentReleasesAccessAllowedRep`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *EnvironmentReleasesAccess) GetAllowedOk() (*[]EnvironmentReleasesAccessAllowedRep, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *EnvironmentReleasesAccess) SetAllowed(v []EnvironmentReleasesAccessAllowedRep)`

SetAllowed sets Allowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


