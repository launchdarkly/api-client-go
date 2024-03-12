# MetricPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key to reference the metric | 
**Name** | Pointer to **string** | A human-friendly name for the metric | [optional] 
**Description** | Pointer to **string** | Description of the metric | [optional] 
**Kind** | **string** | The kind of event your metric will track | 
**Selector** | Pointer to **string** | One or more CSS selectors. Required for click metrics only. | [optional] 
**Urls** | Pointer to [**[]UrlPost**](UrlPost.md) | One or more target URLs. Required for click and pageview metrics only. | [optional] 
**IsActive** | Pointer to **bool** | Whether the metric is active. Set to &lt;code&gt;true&lt;/code&gt; to record click or pageview metrics. Not applicable for custom metrics. | [optional] 
**IsNumeric** | Pointer to **bool** | Whether to track numeric changes in value against a baseline (&lt;code&gt;true&lt;/code&gt;) or to track a conversion when an end user takes an action (&lt;code&gt;false&lt;/code&gt;). Required for custom metrics only. | [optional] 
**Unit** | Pointer to **string** | The unit of measure. Applicable for numeric custom metrics only. | [optional] 
**EventKey** | Pointer to **string** | The event key to use in your code. Required for custom conversion/binary and custom numeric metrics only. | [optional] 
**SuccessCriteria** | Pointer to **string** | Success criteria. Required for custom numeric metrics, optional for custom conversion metrics. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the metric | [optional] 
**RandomizationUnits** | Pointer to **[]string** | An array of randomization units allowed for this metric | [optional] 
**UnitAggregationType** | Pointer to **string** | The method in which multiple unit event values are aggregated | [optional] 

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

### GetRandomizationUnits

`func (o *MetricPost) GetRandomizationUnits() []string`

GetRandomizationUnits returns the RandomizationUnits field if non-nil, zero value otherwise.

### GetRandomizationUnitsOk

`func (o *MetricPost) GetRandomizationUnitsOk() (*[]string, bool)`

GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnits

`func (o *MetricPost) SetRandomizationUnits(v []string)`

SetRandomizationUnits sets RandomizationUnits field to given value.

### HasRandomizationUnits

`func (o *MetricPost) HasRandomizationUnits() bool`

HasRandomizationUnits returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *MetricPost) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *MetricPost) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *MetricPost) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *MetricPost) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


