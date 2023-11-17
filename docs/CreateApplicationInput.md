# CreateApplicationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The application description | [optional] 
**Key** | **string** | The unique identifier of this application | 
**Kind** | **string** | To distinguish the kind of application | 
**Maintainer** | Pointer to [**MaintainerInput**](MaintainerInput.md) |  | [optional] 
**Name** | **string** | The name of the application | 

## Methods

### NewCreateApplicationInput

`func NewCreateApplicationInput(key string, kind string, name string, ) *CreateApplicationInput`

NewCreateApplicationInput instantiates a new CreateApplicationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationInputWithDefaults

`func NewCreateApplicationInputWithDefaults() *CreateApplicationInput`

NewCreateApplicationInputWithDefaults instantiates a new CreateApplicationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateApplicationInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateApplicationInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateApplicationInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateApplicationInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *CreateApplicationInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateApplicationInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateApplicationInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetKind

`func (o *CreateApplicationInput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateApplicationInput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateApplicationInput) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMaintainer

`func (o *CreateApplicationInput) GetMaintainer() MaintainerInput`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *CreateApplicationInput) GetMaintainerOk() (*MaintainerInput, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *CreateApplicationInput) SetMaintainer(v MaintainerInput)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *CreateApplicationInput) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetName

`func (o *CreateApplicationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApplicationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApplicationInput) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


