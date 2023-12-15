# InternalEnvironmentCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Items** | [**[]InternalEnvironment**](InternalEnvironment.md) | A list of requested environments in this project | 

## Methods

### NewInternalEnvironmentCollection

`func NewInternalEnvironmentCollection(items []InternalEnvironment, ) *InternalEnvironmentCollection`

NewInternalEnvironmentCollection instantiates a new InternalEnvironmentCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalEnvironmentCollectionWithDefaults

`func NewInternalEnvironmentCollectionWithDefaults() *InternalEnvironmentCollection`

NewInternalEnvironmentCollectionWithDefaults instantiates a new InternalEnvironmentCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *InternalEnvironmentCollection) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InternalEnvironmentCollection) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InternalEnvironmentCollection) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InternalEnvironmentCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *InternalEnvironmentCollection) GetItems() []InternalEnvironment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InternalEnvironmentCollection) GetItemsOk() (*[]InternalEnvironment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InternalEnvironmentCollection) SetItems(v []InternalEnvironment)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


