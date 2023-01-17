# MemberImportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | An error message, including CSV line number, if the &lt;code&gt;status&lt;/code&gt; is &lt;code&gt;error&lt;/code&gt; | [optional] 
**Status** | **string** | Whether this member can be successfully imported (&lt;code&gt;success&lt;/code&gt;) or not (&lt;code&gt;error&lt;/code&gt;). Even if the status is &lt;code&gt;success&lt;/code&gt;, members are only added to a team on a &lt;code&gt;201&lt;/code&gt; response. | 
**Value** | **string** | The email address for the member requested to be added to this team. May be blank or an error, such as &#39;invalid email format&#39;, if the email address cannot be found or parsed. | 

## Methods

### NewMemberImportItem

`func NewMemberImportItem(status string, value string, ) *MemberImportItem`

NewMemberImportItem instantiates a new MemberImportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberImportItemWithDefaults

`func NewMemberImportItemWithDefaults() *MemberImportItem`

NewMemberImportItemWithDefaults instantiates a new MemberImportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MemberImportItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MemberImportItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MemberImportItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MemberImportItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *MemberImportItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MemberImportItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MemberImportItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetValue

`func (o *MemberImportItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MemberImportItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MemberImportItem) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


