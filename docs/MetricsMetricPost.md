# MetricsMetricPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | One of: custom, click, pageview | [optional] 
**Selector** | Pointer to **string** | Required for click metrics | [optional] 
**Urls** | Pointer to [**[]MetricsMetricPostUrls**](MetricsMetricPostUrls.md) | Required for click and pageview metrics | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsNumeric** | Pointer to **bool** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**EventKey** | Pointer to **string** | Required for custom metrics | [optional] 
**SuccessCriteria** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMetricsMetricPost

`func NewMetricsMetricPost() *MetricsMetricPost`

NewMetricsMetricPost instantiates a new MetricsMetricPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMetricPostWithDefaults

`func NewMetricsMetricPostWithDefaults() *MetricsMetricPost`

NewMetricsMetricPostWithDefaults instantiates a new MetricsMetricPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricsMetricPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricsMetricPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricsMetricPost) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetricsMetricPost) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *MetricsMetricPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsMetricPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsMetricPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsMetricPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetricsMetricPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsMetricPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsMetricPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsMetricPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKind

`func (o *MetricsMetricPost) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricsMetricPost) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricsMetricPost) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MetricsMetricPost) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSelector

`func (o *MetricsMetricPost) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MetricsMetricPost) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MetricsMetricPost) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *MetricsMetricPost) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetUrls

`func (o *MetricsMetricPost) GetUrls() []MetricsMetricPostUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *MetricsMetricPost) GetUrlsOk() (*[]MetricsMetricPostUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *MetricsMetricPost) SetUrls(v []MetricsMetricPostUrls)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *MetricsMetricPost) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetIsActive

`func (o *MetricsMetricPost) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MetricsMetricPost) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MetricsMetricPost) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MetricsMetricPost) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsNumeric

`func (o *MetricsMetricPost) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *MetricsMetricPost) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *MetricsMetricPost) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *MetricsMetricPost) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetUnit

`func (o *MetricsMetricPost) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricsMetricPost) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricsMetricPost) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricsMetricPost) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetEventKey

`func (o *MetricsMetricPost) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *MetricsMetricPost) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *MetricsMetricPost) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *MetricsMetricPost) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetSuccessCriteria

`func (o *MetricsMetricPost) GetSuccessCriteria() int32`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *MetricsMetricPost) GetSuccessCriteriaOk() (*int32, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *MetricsMetricPost) SetSuccessCriteria(v int32)`

SetSuccessCriteria sets SuccessCriteria field to given value.

### HasSuccessCriteria

`func (o *MetricsMetricPost) HasSuccessCriteria() bool`

HasSuccessCriteria returns a boolean if a field has been set.

### GetTags

`func (o *MetricsMetricPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricsMetricPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricsMetricPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricsMetricPost) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


