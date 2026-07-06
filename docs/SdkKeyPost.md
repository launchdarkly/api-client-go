# SdkKeyPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind of SDK key. Can be either \&quot;sdk\&quot; (server-side) or \&quot;mobile\&quot; (mobile). Defaults to \&quot;sdk\&quot; when not explicitly defined. | [optional] [default to "sdk"]
**Key** | **string** | The user-defined key of the SDK key. | 
**Name** | **string** | The human-readable name of the SDK key. | 
**Description** | Pointer to **string** | The optional description of the SDK key. | [optional] 
**Expiry** | Pointer to **int64** |  | [optional] 

## Methods

### NewSdkKeyPost

`func NewSdkKeyPost(key string, name string, ) *SdkKeyPost`

NewSdkKeyPost instantiates a new SdkKeyPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkKeyPostWithDefaults

`func NewSdkKeyPostWithDefaults() *SdkKeyPost`

NewSdkKeyPostWithDefaults instantiates a new SdkKeyPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SdkKeyPost) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SdkKeyPost) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SdkKeyPost) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SdkKeyPost) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetKey

`func (o *SdkKeyPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SdkKeyPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SdkKeyPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *SdkKeyPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkKeyPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkKeyPost) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SdkKeyPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdkKeyPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdkKeyPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdkKeyPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiry

`func (o *SdkKeyPost) GetExpiry() int64`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SdkKeyPost) GetExpiryOk() (*int64, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SdkKeyPost) SetExpiry(v int64)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SdkKeyPost) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


