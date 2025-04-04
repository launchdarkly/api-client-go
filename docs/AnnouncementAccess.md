# AnnouncementAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denied** | [**[]AnnouncementAccessDenied**](AnnouncementAccessDenied.md) |  | 
**Allowed** | [**[]AnnouncementAccessAllowedRep**](AnnouncementAccessAllowedRep.md) |  | 

## Methods

### NewAnnouncementAccess

`func NewAnnouncementAccess(denied []AnnouncementAccessDenied, allowed []AnnouncementAccessAllowedRep, ) *AnnouncementAccess`

NewAnnouncementAccess instantiates a new AnnouncementAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementAccessWithDefaults

`func NewAnnouncementAccessWithDefaults() *AnnouncementAccess`

NewAnnouncementAccessWithDefaults instantiates a new AnnouncementAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenied

`func (o *AnnouncementAccess) GetDenied() []AnnouncementAccessDenied`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *AnnouncementAccess) GetDeniedOk() (*[]AnnouncementAccessDenied, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *AnnouncementAccess) SetDenied(v []AnnouncementAccessDenied)`

SetDenied sets Denied field to given value.


### GetAllowed

`func (o *AnnouncementAccess) GetAllowed() []AnnouncementAccessAllowedRep`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *AnnouncementAccess) GetAllowedOk() (*[]AnnouncementAccessAllowedRep, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *AnnouncementAccess) SetAllowed(v []AnnouncementAccessAllowedRep)`

SetAllowed sets Allowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


