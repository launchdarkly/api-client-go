# AnnouncementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the announcement | 
**IsDismissible** | **bool** | true if the announcement is dismissible | 
**Title** | **string** | The title of the announcement | 
**Message** | **string** | The message of the announcement | 
**StartTime** | **int64** | The start time of the announcement. This is a Unix timestamp in milliseconds. | 
**EndTime** | Pointer to **int64** | The end time of the announcement. This is a Unix timestamp in milliseconds. | [optional] 
**Severity** | **string** | The severity of the announcement | 
**Status** | **string** | The status of the announcement | 
**Access** | Pointer to [**AnnouncementAccessRep**](AnnouncementAccessRep.md) |  | [optional] 
**Links** | [**AnnouncementResponseLinks**](AnnouncementResponseLinks.md) |  | 

## Methods

### NewAnnouncementResponse

`func NewAnnouncementResponse(id string, isDismissible bool, title string, message string, startTime int64, severity string, status string, links AnnouncementResponseLinks, ) *AnnouncementResponse`

NewAnnouncementResponse instantiates a new AnnouncementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementResponseWithDefaults

`func NewAnnouncementResponseWithDefaults() *AnnouncementResponse`

NewAnnouncementResponseWithDefaults instantiates a new AnnouncementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnnouncementResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnnouncementResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnnouncementResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsDismissible

`func (o *AnnouncementResponse) GetIsDismissible() bool`

GetIsDismissible returns the IsDismissible field if non-nil, zero value otherwise.

### GetIsDismissibleOk

`func (o *AnnouncementResponse) GetIsDismissibleOk() (*bool, bool)`

GetIsDismissibleOk returns a tuple with the IsDismissible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDismissible

`func (o *AnnouncementResponse) SetIsDismissible(v bool)`

SetIsDismissible sets IsDismissible field to given value.


### GetTitle

`func (o *AnnouncementResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AnnouncementResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AnnouncementResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *AnnouncementResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AnnouncementResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AnnouncementResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStartTime

`func (o *AnnouncementResponse) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AnnouncementResponse) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AnnouncementResponse) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *AnnouncementResponse) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AnnouncementResponse) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AnnouncementResponse) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AnnouncementResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetSeverity

`func (o *AnnouncementResponse) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AnnouncementResponse) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AnnouncementResponse) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetStatus

`func (o *AnnouncementResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnnouncementResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnnouncementResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccess

`func (o *AnnouncementResponse) GetAccess() AnnouncementAccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AnnouncementResponse) GetAccessOk() (*AnnouncementAccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AnnouncementResponse) SetAccess(v AnnouncementAccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AnnouncementResponse) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *AnnouncementResponse) GetLinks() AnnouncementResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AnnouncementResponse) GetLinksOk() (*AnnouncementResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AnnouncementResponse) SetLinks(v AnnouncementResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


