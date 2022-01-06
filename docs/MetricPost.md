# MetricPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Kind** | **string** |  | 
**Selector** | Pointer to **string** | Required for click metrics | [optional] 
**Urls** | Pointer to [**[]UrlPost**](UrlPost.md) | Required for click and pageview metrics | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsNumeric** | Pointer to **bool** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**EventKey** | Pointer to **string** | Required for custom metrics | [optional] 
**SuccessCriteria** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMetricPost

`func NewMetricPost(key string, kind string, ) *MetricPost`

NewMetricPost instantiates a new MetricPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricPostWithDefaults

`func NewMetricPostWithDefaults() *MetricPost`

NewMetricPostWithDefaults instantiates a new MetricPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *MetricPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetricPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKind

`func (o *MetricPost) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricPost) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricPost) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSelector

`func (o *MetricPost) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MetricPost) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MetricPost) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *MetricPost) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetUrls

`func (o *MetricPost) GetUrls() []UrlPost`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *MetricPost) GetUrlsOk() (*[]UrlPost, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *MetricPost) SetUrls(v []UrlPost)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *MetricPost) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetIsActive

`func (o *MetricPost) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MetricPost) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MetricPost) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MetricPost) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsNumeric

`func (o *MetricPost) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *MetricPost) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *MetricPost) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *MetricPost) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetUnit

`func (o *MetricPost) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricPost) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricPost) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricPost) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetEventKey

`func (o *MetricPost) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *MetricPost) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *MetricPost) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *MetricPost) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetSuccessCriteria

`func (o *MetricPost) GetSuccessCriteria() string`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *MetricPost) GetSuccessCriteriaOk() (*string, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *MetricPost) SetSuccessCriteria(v string)`

SetSuccessCriteria sets SuccessCriteria field to given value.

### HasSuccessCriteria

`func (o *MetricPost) HasSuccessCriteria() bool`

HasSuccessCriteria returns a boolean if a field has been set.

### GetTags

`func (o *MetricPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricPost) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


