# SdkKeyForPutSdkKeyViews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 
**Kind** | [**SdkKeyKind**](SdkKeyKind.md) |  | 
**Key** | **string** | The user-defined identifying key of the SDK key. This is used solely to identify an SDK key and is distinct from the value field, which is the actual SDK key value. | 
**Name** | **string** | The human-readable name of the SDK key. | 
**Description** | Pointer to **string** | The optional description of the SDK key. | [optional] 
**Expiry** | Pointer to **int64** |  | [optional] 
**Value** | **string** | The string value of the SDK key. Use this when configuring your SDK. | 
**IsDefault** | **bool** | Indicates if this SDK key is the system-defined default for the environment. There may also be an expiring default SDK key for the environment (not possible with mobile keys). | 
**CreatedByMemberId** | Pointer to **string** | The ID of the member who created the SDK key. This field is immutable. | [optional] 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 
**Version** | **int32** | The auto-incremented version number of the SDK key. | 
**ViewSummaries** | Pointer to [**[]ViewSummary**](ViewSummary.md) | Summaries of views associated with the SDK key. | [optional] 

## Methods

### NewSdkKeyForPutSdkKeyViews

`func NewSdkKeyForPutSdkKeyViews(kind SdkKeyKind, key string, name string, value string, isDefault bool, createdAt int64, updatedAt int64, version int32, ) *SdkKeyForPutSdkKeyViews`

NewSdkKeyForPutSdkKeyViews instantiates a new SdkKeyForPutSdkKeyViews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkKeyForPutSdkKeyViewsWithDefaults

`func NewSdkKeyForPutSdkKeyViewsWithDefaults() *SdkKeyForPutSdkKeyViews`

NewSdkKeyForPutSdkKeyViewsWithDefaults instantiates a new SdkKeyForPutSdkKeyViews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SdkKeyForPutSdkKeyViews) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SdkKeyForPutSdkKeyViews) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SdkKeyForPutSdkKeyViews) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SdkKeyForPutSdkKeyViews) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetKind

`func (o *SdkKeyForPutSdkKeyViews) GetKind() SdkKeyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SdkKeyForPutSdkKeyViews) GetKindOk() (*SdkKeyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SdkKeyForPutSdkKeyViews) SetKind(v SdkKeyKind)`

SetKind sets Kind field to given value.


### GetKey

`func (o *SdkKeyForPutSdkKeyViews) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SdkKeyForPutSdkKeyViews) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SdkKeyForPutSdkKeyViews) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *SdkKeyForPutSdkKeyViews) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkKeyForPutSdkKeyViews) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkKeyForPutSdkKeyViews) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SdkKeyForPutSdkKeyViews) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdkKeyForPutSdkKeyViews) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdkKeyForPutSdkKeyViews) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdkKeyForPutSdkKeyViews) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiry

`func (o *SdkKeyForPutSdkKeyViews) GetExpiry() int64`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SdkKeyForPutSdkKeyViews) GetExpiryOk() (*int64, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SdkKeyForPutSdkKeyViews) SetExpiry(v int64)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SdkKeyForPutSdkKeyViews) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetValue

`func (o *SdkKeyForPutSdkKeyViews) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SdkKeyForPutSdkKeyViews) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SdkKeyForPutSdkKeyViews) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsDefault

`func (o *SdkKeyForPutSdkKeyViews) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SdkKeyForPutSdkKeyViews) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SdkKeyForPutSdkKeyViews) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCreatedByMemberId

`func (o *SdkKeyForPutSdkKeyViews) GetCreatedByMemberId() string`

GetCreatedByMemberId returns the CreatedByMemberId field if non-nil, zero value otherwise.

### GetCreatedByMemberIdOk

`func (o *SdkKeyForPutSdkKeyViews) GetCreatedByMemberIdOk() (*string, bool)`

GetCreatedByMemberIdOk returns a tuple with the CreatedByMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByMemberId

`func (o *SdkKeyForPutSdkKeyViews) SetCreatedByMemberId(v string)`

SetCreatedByMemberId sets CreatedByMemberId field to given value.

### HasCreatedByMemberId

`func (o *SdkKeyForPutSdkKeyViews) HasCreatedByMemberId() bool`

HasCreatedByMemberId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SdkKeyForPutSdkKeyViews) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SdkKeyForPutSdkKeyViews) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SdkKeyForPutSdkKeyViews) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SdkKeyForPutSdkKeyViews) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SdkKeyForPutSdkKeyViews) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SdkKeyForPutSdkKeyViews) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *SdkKeyForPutSdkKeyViews) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SdkKeyForPutSdkKeyViews) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SdkKeyForPutSdkKeyViews) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetViewSummaries

`func (o *SdkKeyForPutSdkKeyViews) GetViewSummaries() []ViewSummary`

GetViewSummaries returns the ViewSummaries field if non-nil, zero value otherwise.

### GetViewSummariesOk

`func (o *SdkKeyForPutSdkKeyViews) GetViewSummariesOk() (*[]ViewSummary, bool)`

GetViewSummariesOk returns a tuple with the ViewSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSummaries

`func (o *SdkKeyForPutSdkKeyViews) SetViewSummaries(v []ViewSummary)`

SetViewSummaries sets ViewSummaries field to given value.

### HasViewSummaries

`func (o *SdkKeyForPutSdkKeyViews) HasViewSummaries() bool`

HasViewSummaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


