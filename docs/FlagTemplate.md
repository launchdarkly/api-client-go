# FlagTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Key** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ClientSideAvailability** | Pointer to [**ClientSideAvailability**](ClientSideAvailability.md) |  | [optional] 
**VariationType** | Pointer to **string** |  | [optional] 
**Variations** | Pointer to [**[]Variation**](Variation.md) |  | [optional] 
**VariationJsonSchema** | Pointer to **interface{}** |  | [optional] 
**Temporary** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CustomProperties** | Pointer to [**[]CustomProperty**](CustomProperty.md) |  | [optional] 
**DefaultVariations** | Pointer to [**Defaults**](Defaults.md) |  | [optional] 

## Methods

### NewFlagTemplate

`func NewFlagTemplate(id string, name string, key string, ) *FlagTemplate`

NewFlagTemplate instantiates a new FlagTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagTemplateWithDefaults

`func NewFlagTemplateWithDefaults() *FlagTemplate`

NewFlagTemplateWithDefaults instantiates a new FlagTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlagTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FlagTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *FlagTemplate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagTemplate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagTemplate) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *FlagTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlagTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClientSideAvailability

`func (o *FlagTemplate) GetClientSideAvailability() ClientSideAvailability`

GetClientSideAvailability returns the ClientSideAvailability field if non-nil, zero value otherwise.

### GetClientSideAvailabilityOk

`func (o *FlagTemplate) GetClientSideAvailabilityOk() (*ClientSideAvailability, bool)`

GetClientSideAvailabilityOk returns a tuple with the ClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideAvailability

`func (o *FlagTemplate) SetClientSideAvailability(v ClientSideAvailability)`

SetClientSideAvailability sets ClientSideAvailability field to given value.

### HasClientSideAvailability

`func (o *FlagTemplate) HasClientSideAvailability() bool`

HasClientSideAvailability returns a boolean if a field has been set.

### GetVariationType

`func (o *FlagTemplate) GetVariationType() string`

GetVariationType returns the VariationType field if non-nil, zero value otherwise.

### GetVariationTypeOk

`func (o *FlagTemplate) GetVariationTypeOk() (*string, bool)`

GetVariationTypeOk returns a tuple with the VariationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationType

`func (o *FlagTemplate) SetVariationType(v string)`

SetVariationType sets VariationType field to given value.

### HasVariationType

`func (o *FlagTemplate) HasVariationType() bool`

HasVariationType returns a boolean if a field has been set.

### GetVariations

`func (o *FlagTemplate) GetVariations() []Variation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *FlagTemplate) GetVariationsOk() (*[]Variation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *FlagTemplate) SetVariations(v []Variation)`

SetVariations sets Variations field to given value.

### HasVariations

`func (o *FlagTemplate) HasVariations() bool`

HasVariations returns a boolean if a field has been set.

### GetVariationJsonSchema

`func (o *FlagTemplate) GetVariationJsonSchema() interface{}`

GetVariationJsonSchema returns the VariationJsonSchema field if non-nil, zero value otherwise.

### GetVariationJsonSchemaOk

`func (o *FlagTemplate) GetVariationJsonSchemaOk() (*interface{}, bool)`

GetVariationJsonSchemaOk returns a tuple with the VariationJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationJsonSchema

`func (o *FlagTemplate) SetVariationJsonSchema(v interface{})`

SetVariationJsonSchema sets VariationJsonSchema field to given value.

### HasVariationJsonSchema

`func (o *FlagTemplate) HasVariationJsonSchema() bool`

HasVariationJsonSchema returns a boolean if a field has been set.

### SetVariationJsonSchemaNil

`func (o *FlagTemplate) SetVariationJsonSchemaNil(b bool)`

 SetVariationJsonSchemaNil sets the value for VariationJsonSchema to be an explicit nil

### UnsetVariationJsonSchema
`func (o *FlagTemplate) UnsetVariationJsonSchema()`

UnsetVariationJsonSchema ensures that no value is present for VariationJsonSchema, not even an explicit nil
### GetTemporary

`func (o *FlagTemplate) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *FlagTemplate) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *FlagTemplate) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *FlagTemplate) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetTags

`func (o *FlagTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlagTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlagTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlagTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomProperties

`func (o *FlagTemplate) GetCustomProperties() []CustomProperty`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *FlagTemplate) GetCustomPropertiesOk() (*[]CustomProperty, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *FlagTemplate) SetCustomProperties(v []CustomProperty)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *FlagTemplate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDefaultVariations

`func (o *FlagTemplate) GetDefaultVariations() Defaults`

GetDefaultVariations returns the DefaultVariations field if non-nil, zero value otherwise.

### GetDefaultVariationsOk

`func (o *FlagTemplate) GetDefaultVariationsOk() (*Defaults, bool)`

GetDefaultVariationsOk returns a tuple with the DefaultVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVariations

`func (o *FlagTemplate) SetDefaultVariations(v Defaults)`

SetDefaultVariations sets DefaultVariations field to given value.

### HasDefaultVariations

`func (o *FlagTemplate) HasDefaultVariations() bool`

HasDefaultVariations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


