# UpsertContextKindPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The context kind name | 
**Description** | Pointer to **string** | The context kind description | [optional] 
**HideInTargeting** | Pointer to **bool** | Alias for archived. | [optional] 
**Archived** | Pointer to **bool** | Whether the context kind is archived. Archived context kinds are unavailable for targeting. | [optional] 
**Version** | Pointer to **int32** | The context kind version. If not specified when the context kind is created, defaults to 1. | [optional] 

## Methods

### NewUpsertContextKindPayload

`func NewUpsertContextKindPayload(name string, ) *UpsertContextKindPayload`

NewUpsertContextKindPayload instantiates a new UpsertContextKindPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertContextKindPayloadWithDefaults

`func NewUpsertContextKindPayloadWithDefaults() *UpsertContextKindPayload`

NewUpsertContextKindPayloadWithDefaults instantiates a new UpsertContextKindPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpsertContextKindPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertContextKindPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertContextKindPayload) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpsertContextKindPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpsertContextKindPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpsertContextKindPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpsertContextKindPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHideInTargeting

`func (o *UpsertContextKindPayload) GetHideInTargeting() bool`

GetHideInTargeting returns the HideInTargeting field if non-nil, zero value otherwise.

### GetHideInTargetingOk

`func (o *UpsertContextKindPayload) GetHideInTargetingOk() (*bool, bool)`

GetHideInTargetingOk returns a tuple with the HideInTargeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideInTargeting

`func (o *UpsertContextKindPayload) SetHideInTargeting(v bool)`

SetHideInTargeting sets HideInTargeting field to given value.

### HasHideInTargeting

`func (o *UpsertContextKindPayload) HasHideInTargeting() bool`

HasHideInTargeting returns a boolean if a field has been set.

### GetArchived

`func (o *UpsertContextKindPayload) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpsertContextKindPayload) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpsertContextKindPayload) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpsertContextKindPayload) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetVersion

`func (o *UpsertContextKindPayload) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpsertContextKindPayload) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpsertContextKindPayload) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpsertContextKindPayload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


