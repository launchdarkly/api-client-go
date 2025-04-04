# CreateAnnouncementBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDismissible** | **bool** | true if the announcement is dismissible | 
**Title** | **string** | The title of the announcement | 
**Message** | **string** | The message of the announcement | 
**StartTime** | **int64** | The start time of the announcement. This is a Unix timestamp in milliseconds. | 
**EndTime** | Pointer to **int64** | The end time of the announcement. This is a Unix timestamp in milliseconds. | [optional] 
**Severity** | **string** | The severity of the announcement | 

## Methods

### NewCreateAnnouncementBody

`func NewCreateAnnouncementBody(isDismissible bool, title string, message string, startTime int64, severity string, ) *CreateAnnouncementBody`

NewCreateAnnouncementBody instantiates a new CreateAnnouncementBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAnnouncementBodyWithDefaults

`func NewCreateAnnouncementBodyWithDefaults() *CreateAnnouncementBody`

NewCreateAnnouncementBodyWithDefaults instantiates a new CreateAnnouncementBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDismissible

`func (o *CreateAnnouncementBody) GetIsDismissible() bool`

GetIsDismissible returns the IsDismissible field if non-nil, zero value otherwise.

### GetIsDismissibleOk

`func (o *CreateAnnouncementBody) GetIsDismissibleOk() (*bool, bool)`

GetIsDismissibleOk returns a tuple with the IsDismissible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDismissible

`func (o *CreateAnnouncementBody) SetIsDismissible(v bool)`

SetIsDismissible sets IsDismissible field to given value.


### GetTitle

`func (o *CreateAnnouncementBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateAnnouncementBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateAnnouncementBody) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *CreateAnnouncementBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateAnnouncementBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateAnnouncementBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStartTime

`func (o *CreateAnnouncementBody) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateAnnouncementBody) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateAnnouncementBody) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CreateAnnouncementBody) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreateAnnouncementBody) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreateAnnouncementBody) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CreateAnnouncementBody) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetSeverity

`func (o *CreateAnnouncementBody) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CreateAnnouncementBody) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CreateAnnouncementBody) SetSeverity(v string)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


