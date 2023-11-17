# ContextKind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The context kind key | 
**Name** | **string** | The context kind name | 
**Description** | **string** | The context kind description | 
**Version** | **int32** | The context kind version | 
**CreationDate** | **int64** |  | 
**LastModified** | **int64** |  | 
**LastSeen** | Pointer to **int64** |  | [optional] 
**CreatedFrom** | **string** |  | 
**HideInTargeting** | Pointer to **bool** | Alias for archived. | [optional] 
**Archived** | Pointer to **bool** | Whether the context kind is archived. Archived context kinds are unavailable for targeting. | [optional] 

## Methods

### NewContextKind

`func NewContextKind(key string, name string, description string, version int32, creationDate int64, lastModified int64, createdFrom string, ) *ContextKind`

NewContextKind instantiates a new ContextKind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextKindWithDefaults

`func NewContextKindWithDefaults() *ContextKind`

NewContextKindWithDefaults instantiates a new ContextKind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ContextKind) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContextKind) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContextKind) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ContextKind) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextKind) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextKind) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ContextKind) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContextKind) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContextKind) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVersion

`func (o *ContextKind) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContextKind) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContextKind) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreationDate

`func (o *ContextKind) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ContextKind) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ContextKind) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLastModified

`func (o *ContextKind) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ContextKind) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ContextKind) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.


### GetLastSeen

`func (o *ContextKind) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ContextKind) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ContextKind) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ContextKind) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetCreatedFrom

`func (o *ContextKind) GetCreatedFrom() string`

GetCreatedFrom returns the CreatedFrom field if non-nil, zero value otherwise.

### GetCreatedFromOk

`func (o *ContextKind) GetCreatedFromOk() (*string, bool)`

GetCreatedFromOk returns a tuple with the CreatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFrom

`func (o *ContextKind) SetCreatedFrom(v string)`

SetCreatedFrom sets CreatedFrom field to given value.


### GetHideInTargeting

`func (o *ContextKind) GetHideInTargeting() bool`

GetHideInTargeting returns the HideInTargeting field if non-nil, zero value otherwise.

### GetHideInTargetingOk

`func (o *ContextKind) GetHideInTargetingOk() (*bool, bool)`

GetHideInTargetingOk returns a tuple with the HideInTargeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideInTargeting

`func (o *ContextKind) SetHideInTargeting(v bool)`

SetHideInTargeting sets HideInTargeting field to given value.

### HasHideInTargeting

`func (o *ContextKind) HasHideInTargeting() bool`

HasHideInTargeting returns a boolean if a field has been set.

### GetArchived

`func (o *ContextKind) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ContextKind) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ContextKind) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ContextKind) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


