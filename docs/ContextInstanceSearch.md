# ContextInstanceSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | A collection of context instance filters | [optional] 
**Sort** | Pointer to **string** | Specifies a field by which to sort. LaunchDarkly supports sorting by timestamp in ascending order by specifying &lt;code&gt;ts&lt;/code&gt; for this value, or descending order by specifying &lt;code&gt;-ts&lt;/code&gt;. | [optional] 
**Limit** | Pointer to **int32** | Specifies the maximum number of items in the collection to return (max: 50, default: 20) | [optional] 
**ContinuationToken** | Pointer to **string** | Limits results to context instances with sort values after the value specified. You can use this for pagination, however, we recommend using the &lt;code&gt;next&lt;/code&gt; link instead, because this value is an obfuscated string. | [optional] 

## Methods

### NewContextInstanceSearch

`func NewContextInstanceSearch() *ContextInstanceSearch`

NewContextInstanceSearch instantiates a new ContextInstanceSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextInstanceSearchWithDefaults

`func NewContextInstanceSearchWithDefaults() *ContextInstanceSearch`

NewContextInstanceSearchWithDefaults instantiates a new ContextInstanceSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ContextInstanceSearch) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ContextInstanceSearch) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ContextInstanceSearch) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ContextInstanceSearch) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *ContextInstanceSearch) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ContextInstanceSearch) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ContextInstanceSearch) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ContextInstanceSearch) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetLimit

`func (o *ContextInstanceSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ContextInstanceSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ContextInstanceSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ContextInstanceSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetContinuationToken

`func (o *ContextInstanceSearch) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ContextInstanceSearch) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ContextInstanceSearch) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ContextInstanceSearch) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


