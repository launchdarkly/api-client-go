# ParentResourceRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]CoreLink**](CoreLink.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewParentResourceRep

`func NewParentResourceRep() *ParentResourceRep`

NewParentResourceRep instantiates a new ParentResourceRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentResourceRepWithDefaults

`func NewParentResourceRepWithDefaults() *ParentResourceRep`

NewParentResourceRepWithDefaults instantiates a new ParentResourceRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ParentResourceRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ParentResourceRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ParentResourceRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ParentResourceRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *ParentResourceRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParentResourceRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParentResourceRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParentResourceRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResource

`func (o *ParentResourceRep) GetResource() interface{}`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ParentResourceRep) GetResourceOk() (*interface{}, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ParentResourceRep) SetResource(v interface{})`

SetResource sets Resource field to given value.

### HasResource

`func (o *ParentResourceRep) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ParentResourceRep) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ParentResourceRep) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


