# SdkKeyListItem

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
**EnvironmentSummary** | Pointer to [**SdkKeysEnvironmentSummary**](SdkKeysEnvironmentSummary.md) |  | [optional] 

## Methods

### NewSdkKeyListItem

`func NewSdkKeyListItem(kind SdkKeyKind, key string, name string, value string, isDefault bool, createdAt int64, updatedAt int64, version int32, ) *SdkKeyListItem`

NewSdkKeyListItem instantiates a new SdkKeyListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkKeyListItemWithDefaults

`func NewSdkKeyListItemWithDefaults() *SdkKeyListItem`

NewSdkKeyListItemWithDefaults instantiates a new SdkKeyListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SdkKeyListItem) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SdkKeyListItem) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SdkKeyListItem) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SdkKeyListItem) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetKind

`func (o *SdkKeyListItem) GetKind() SdkKeyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SdkKeyListItem) GetKindOk() (*SdkKeyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SdkKeyListItem) SetKind(v SdkKeyKind)`

SetKind sets Kind field to given value.


### GetKey

`func (o *SdkKeyListItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SdkKeyListItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SdkKeyListItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *SdkKeyListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkKeyListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkKeyListItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SdkKeyListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdkKeyListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdkKeyListItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdkKeyListItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiry

`func (o *SdkKeyListItem) GetExpiry() int64`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SdkKeyListItem) GetExpiryOk() (*int64, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SdkKeyListItem) SetExpiry(v int64)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SdkKeyListItem) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetValue

`func (o *SdkKeyListItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SdkKeyListItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SdkKeyListItem) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsDefault

`func (o *SdkKeyListItem) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SdkKeyListItem) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SdkKeyListItem) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCreatedByMemberId

`func (o *SdkKeyListItem) GetCreatedByMemberId() string`

GetCreatedByMemberId returns the CreatedByMemberId field if non-nil, zero value otherwise.

### GetCreatedByMemberIdOk

`func (o *SdkKeyListItem) GetCreatedByMemberIdOk() (*string, bool)`

GetCreatedByMemberIdOk returns a tuple with the CreatedByMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByMemberId

`func (o *SdkKeyListItem) SetCreatedByMemberId(v string)`

SetCreatedByMemberId sets CreatedByMemberId field to given value.

### HasCreatedByMemberId

`func (o *SdkKeyListItem) HasCreatedByMemberId() bool`

HasCreatedByMemberId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SdkKeyListItem) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SdkKeyListItem) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SdkKeyListItem) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SdkKeyListItem) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SdkKeyListItem) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SdkKeyListItem) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *SdkKeyListItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SdkKeyListItem) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SdkKeyListItem) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetEnvironmentSummary

`func (o *SdkKeyListItem) GetEnvironmentSummary() SdkKeysEnvironmentSummary`

GetEnvironmentSummary returns the EnvironmentSummary field if non-nil, zero value otherwise.

### GetEnvironmentSummaryOk

`func (o *SdkKeyListItem) GetEnvironmentSummaryOk() (*SdkKeysEnvironmentSummary, bool)`

GetEnvironmentSummaryOk returns a tuple with the EnvironmentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentSummary

`func (o *SdkKeyListItem) SetEnvironmentSummary(v SdkKeysEnvironmentSummary)`

SetEnvironmentSummary sets EnvironmentSummary field to given value.

### HasEnvironmentSummary

`func (o *SdkKeyListItem) HasEnvironmentSummary() bool`

HasEnvironmentSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


