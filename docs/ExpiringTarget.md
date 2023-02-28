# ExpiringTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of this expiring target | 
**Version** | **int32** | The version of this expiring target | 
**ExpirationDate** | **int64** |  | 
**ContextKind** | **string** | The context kind of the context to be removed | 
**ContextKey** | **string** | A unique key used to represent the context to be removed | 
**TargetType** | Pointer to **string** | A segment&#39;s target type, &lt;code&gt;included&lt;/code&gt; or &lt;code&gt;excluded&lt;/code&gt;. Included when expiring targets are updated on a segment. | [optional] 
**VariationId** | Pointer to **string** | A unique ID used to represent the flag variation. Included when expiring targets are updated on a feature flag. | [optional] 
**ResourceId** | [**ResourceId**](ResourceId.md) |  | 

## Methods

### NewExpiringTarget

`func NewExpiringTarget(id string, version int32, expirationDate int64, contextKind string, contextKey string, resourceId ResourceId, ) *ExpiringTarget`

NewExpiringTarget instantiates a new ExpiringTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringTargetWithDefaults

`func NewExpiringTargetWithDefaults() *ExpiringTarget`

NewExpiringTargetWithDefaults instantiates a new ExpiringTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpiringTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpiringTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpiringTarget) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *ExpiringTarget) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExpiringTarget) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExpiringTarget) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetExpirationDate

`func (o *ExpiringTarget) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ExpiringTarget) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ExpiringTarget) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.


### GetContextKind

`func (o *ExpiringTarget) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *ExpiringTarget) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *ExpiringTarget) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.


### GetContextKey

`func (o *ExpiringTarget) GetContextKey() string`

GetContextKey returns the ContextKey field if non-nil, zero value otherwise.

### GetContextKeyOk

`func (o *ExpiringTarget) GetContextKeyOk() (*string, bool)`

GetContextKeyOk returns a tuple with the ContextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKey

`func (o *ExpiringTarget) SetContextKey(v string)`

SetContextKey sets ContextKey field to given value.


### GetTargetType

`func (o *ExpiringTarget) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ExpiringTarget) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ExpiringTarget) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ExpiringTarget) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetVariationId

`func (o *ExpiringTarget) GetVariationId() string`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *ExpiringTarget) GetVariationIdOk() (*string, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *ExpiringTarget) SetVariationId(v string)`

SetVariationId sets VariationId field to given value.

### HasVariationId

`func (o *ExpiringTarget) HasVariationId() bool`

HasVariationId returns a boolean if a field has been set.

### GetResourceId

`func (o *ExpiringTarget) GetResourceId() ResourceId`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ExpiringTarget) GetResourceIdOk() (*ResourceId, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ExpiringTarget) SetResourceId(v ResourceId)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


