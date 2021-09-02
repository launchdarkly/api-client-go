# ResourceAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewResourceAccess

`func NewResourceAccess() *ResourceAccess`

NewResourceAccess instantiates a new ResourceAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAccessWithDefaults

`func NewResourceAccessWithDefaults() *ResourceAccess`

NewResourceAccessWithDefaults instantiates a new ResourceAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ResourceAccess) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ResourceAccess) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ResourceAccess) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ResourceAccess) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResource

`func (o *ResourceAccess) GetResource() interface{}`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceAccess) GetResourceOk() (*interface{}, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceAccess) SetResource(v interface{})`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceAccess) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ResourceAccess) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ResourceAccess) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


