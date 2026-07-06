# ViewSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 
**Key** | **string** | The key of the view. | 
**Name** | **string** | The human-readable name of the view. | 
**ResourceSummary** | [**ViewResourceSummary**](ViewResourceSummary.md) |  | 

## Methods

### NewViewSummary

`func NewViewSummary(key string, name string, resourceSummary ViewResourceSummary, ) *ViewSummary`

NewViewSummary instantiates a new ViewSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewSummaryWithDefaults

`func NewViewSummaryWithDefaults() *ViewSummary`

NewViewSummaryWithDefaults instantiates a new ViewSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ViewSummary) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ViewSummary) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ViewSummary) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ViewSummary) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetKey

`func (o *ViewSummary) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ViewSummary) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ViewSummary) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ViewSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewSummary) SetName(v string)`

SetName sets Name field to given value.


### GetResourceSummary

`func (o *ViewSummary) GetResourceSummary() ViewResourceSummary`

GetResourceSummary returns the ResourceSummary field if non-nil, zero value otherwise.

### GetResourceSummaryOk

`func (o *ViewSummary) GetResourceSummaryOk() (*ViewResourceSummary, bool)`

GetResourceSummaryOk returns a tuple with the ResourceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSummary

`func (o *ViewSummary) SetResourceSummary(v ViewResourceSummary)`

SetResourceSummary sets ResourceSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


