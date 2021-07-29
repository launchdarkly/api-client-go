# MetricRepAttachedFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Key** | **string** |  | 
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Site** | Pointer to [**CoreLink**](CoreLink.md) |  | [optional] 

## Methods

### NewMetricRepAttachedFeatures

`func NewMetricRepAttachedFeatures(name string, key string, ) *MetricRepAttachedFeatures`

NewMetricRepAttachedFeatures instantiates a new MetricRepAttachedFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRepAttachedFeaturesWithDefaults

`func NewMetricRepAttachedFeaturesWithDefaults() *MetricRepAttachedFeatures`

NewMetricRepAttachedFeaturesWithDefaults instantiates a new MetricRepAttachedFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricRepAttachedFeatures) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricRepAttachedFeatures) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricRepAttachedFeatures) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *MetricRepAttachedFeatures) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricRepAttachedFeatures) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricRepAttachedFeatures) SetKey(v string)`

SetKey sets Key field to given value.


### GetLinks

`func (o *MetricRepAttachedFeatures) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricRepAttachedFeatures) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricRepAttachedFeatures) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MetricRepAttachedFeatures) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSite

`func (o *MetricRepAttachedFeatures) GetSite() CoreLink`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MetricRepAttachedFeatures) GetSiteOk() (*CoreLink, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MetricRepAttachedFeatures) SetSite(v CoreLink)`

SetSite sets Site field to given value.

### HasSite

`func (o *MetricRepAttachedFeatures) HasSite() bool`

HasSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


