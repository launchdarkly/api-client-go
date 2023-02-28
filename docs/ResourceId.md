# ResourceId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentKey** | Pointer to **string** | The environment key | [optional] 
**FlagKey** | Pointer to **string** | Deprecated, use &lt;code&gt;key&lt;/code&gt; instead | [optional] 
**Key** | Pointer to **string** | The key of the flag or segment | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ProjectKey** | Pointer to **string** | The project key | [optional] 

## Methods

### NewResourceId

`func NewResourceId() *ResourceId`

NewResourceId instantiates a new ResourceId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceIdWithDefaults

`func NewResourceIdWithDefaults() *ResourceId`

NewResourceIdWithDefaults instantiates a new ResourceId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentKey

`func (o *ResourceId) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *ResourceId) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *ResourceId) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *ResourceId) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetFlagKey

`func (o *ResourceId) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *ResourceId) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *ResourceId) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.

### HasFlagKey

`func (o *ResourceId) HasFlagKey() bool`

HasFlagKey returns a boolean if a field has been set.

### GetKey

`func (o *ResourceId) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResourceId) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResourceId) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ResourceId) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKind

`func (o *ResourceId) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResourceId) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResourceId) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ResourceId) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProjectKey

`func (o *ResourceId) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *ResourceId) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *ResourceId) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *ResourceId) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


