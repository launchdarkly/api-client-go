# TargetResourceRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewTargetResourceRep

`func NewTargetResourceRep() *TargetResourceRep`

NewTargetResourceRep instantiates a new TargetResourceRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetResourceRepWithDefaults

`func NewTargetResourceRepWithDefaults() *TargetResourceRep`

NewTargetResourceRepWithDefaults instantiates a new TargetResourceRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TargetResourceRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TargetResourceRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TargetResourceRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TargetResourceRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *TargetResourceRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetResourceRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetResourceRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TargetResourceRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *TargetResourceRep) GetResources() []interface{}`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TargetResourceRep) GetResourcesOk() (*[]interface{}, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TargetResourceRep) SetResources(v []interface{})`

SetResources sets Resources field to given value.

### HasResources

`func (o *TargetResourceRep) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


