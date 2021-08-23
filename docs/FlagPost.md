# FlagPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the feature flag | 
**Key** | **string** | A unique key to reference the flag in your code | 
**Description** | Pointer to **string** | Description of the feature flag | [optional] 
**IncludeInSnippet** | Pointer to **bool** | Deprecated, use clientSideAvailability. Whether or not this flag should be made available to the client-side JavaScript SDK | [optional] 
**ClientSideAvailability** | Pointer to [**ClientSideAvailabilityPost**](ClientSideAvailabilityPost.md) |  | [optional] 
**Variations** | Pointer to [**[]Variate**](Variate.md) | An array of possible variations for the flag | [optional] 
**VariationJsonSchema** | Pointer to **interface{}** |  | [optional] 
**Temporary** | Pointer to **bool** | Whether or not the flag is a temporary flag | [optional] 
**Tags** | Pointer to **[]string** | Tags for the feature flag | [optional] 
**CustomProperties** | Pointer to [**map[string]CustomProperty**](CustomProperty.md) |  | [optional] 
**Defaults** | Pointer to [**Defaults**](Defaults.md) |  | [optional] 

## Methods

### NewFlagPost

`func NewFlagPost(name string, key string, ) *FlagPost`

NewFlagPost instantiates a new FlagPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagPostWithDefaults

`func NewFlagPostWithDefaults() *FlagPost`

NewFlagPostWithDefaults instantiates a new FlagPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlagPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagPost) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *FlagPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *FlagPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlagPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludeInSnippet

`func (o *FlagPost) GetIncludeInSnippet() bool`

GetIncludeInSnippet returns the IncludeInSnippet field if non-nil, zero value otherwise.

### GetIncludeInSnippetOk

`func (o *FlagPost) GetIncludeInSnippetOk() (*bool, bool)`

GetIncludeInSnippetOk returns a tuple with the IncludeInSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippet

`func (o *FlagPost) SetIncludeInSnippet(v bool)`

SetIncludeInSnippet sets IncludeInSnippet field to given value.

### HasIncludeInSnippet

`func (o *FlagPost) HasIncludeInSnippet() bool`

HasIncludeInSnippet returns a boolean if a field has been set.

### GetClientSideAvailability

`func (o *FlagPost) GetClientSideAvailability() ClientSideAvailabilityPost`

GetClientSideAvailability returns the ClientSideAvailability field if non-nil, zero value otherwise.

### GetClientSideAvailabilityOk

`func (o *FlagPost) GetClientSideAvailabilityOk() (*ClientSideAvailabilityPost, bool)`

GetClientSideAvailabilityOk returns a tuple with the ClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideAvailability

`func (o *FlagPost) SetClientSideAvailability(v ClientSideAvailabilityPost)`

SetClientSideAvailability sets ClientSideAvailability field to given value.

### HasClientSideAvailability

`func (o *FlagPost) HasClientSideAvailability() bool`

HasClientSideAvailability returns a boolean if a field has been set.

### GetVariations

`func (o *FlagPost) GetVariations() []Variate`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *FlagPost) GetVariationsOk() (*[]Variate, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *FlagPost) SetVariations(v []Variate)`

SetVariations sets Variations field to given value.

### HasVariations

`func (o *FlagPost) HasVariations() bool`

HasVariations returns a boolean if a field has been set.

### GetVariationJsonSchema

`func (o *FlagPost) GetVariationJsonSchema() interface{}`

GetVariationJsonSchema returns the VariationJsonSchema field if non-nil, zero value otherwise.

### GetVariationJsonSchemaOk

`func (o *FlagPost) GetVariationJsonSchemaOk() (*interface{}, bool)`

GetVariationJsonSchemaOk returns a tuple with the VariationJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationJsonSchema

`func (o *FlagPost) SetVariationJsonSchema(v interface{})`

SetVariationJsonSchema sets VariationJsonSchema field to given value.

### HasVariationJsonSchema

`func (o *FlagPost) HasVariationJsonSchema() bool`

HasVariationJsonSchema returns a boolean if a field has been set.

### SetVariationJsonSchemaNil

`func (o *FlagPost) SetVariationJsonSchemaNil(b bool)`

 SetVariationJsonSchemaNil sets the value for VariationJsonSchema to be an explicit nil

### UnsetVariationJsonSchema
`func (o *FlagPost) UnsetVariationJsonSchema()`

UnsetVariationJsonSchema ensures that no value is present for VariationJsonSchema, not even an explicit nil
### GetTemporary

`func (o *FlagPost) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *FlagPost) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *FlagPost) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *FlagPost) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetTags

`func (o *FlagPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlagPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlagPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlagPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomProperties

`func (o *FlagPost) GetCustomProperties() map[string]CustomProperty`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *FlagPost) GetCustomPropertiesOk() (*map[string]CustomProperty, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *FlagPost) SetCustomProperties(v map[string]CustomProperty)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *FlagPost) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDefaults

`func (o *FlagPost) GetDefaults() Defaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *FlagPost) GetDefaultsOk() (*Defaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *FlagPost) SetDefaults(v Defaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *FlagPost) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


