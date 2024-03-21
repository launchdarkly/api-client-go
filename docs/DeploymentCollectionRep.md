# DeploymentCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of deployments | 
**Items** | [**[]DeploymentRep**](DeploymentRep.md) | A list of deployments | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewDeploymentCollectionRep

`func NewDeploymentCollectionRep(totalCount int32, items []DeploymentRep, ) *DeploymentCollectionRep`

NewDeploymentCollectionRep instantiates a new DeploymentCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCollectionRepWithDefaults

`func NewDeploymentCollectionRepWithDefaults() *DeploymentCollectionRep`

NewDeploymentCollectionRepWithDefaults instantiates a new DeploymentCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *DeploymentCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DeploymentCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DeploymentCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetItems

`func (o *DeploymentCollectionRep) GetItems() []DeploymentRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeploymentCollectionRep) GetItemsOk() (*[]DeploymentRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeploymentCollectionRep) SetItems(v []DeploymentRep)`

SetItems sets Items field to given value.


### GetLinks

`func (o *DeploymentCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeploymentCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeploymentCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeploymentCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


