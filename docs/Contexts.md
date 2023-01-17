# Contexts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**TotalCount** | **int32** | The number of contexts | 
**EnvironmentId** | **string** | The environment ID where the context was evaluated | 
**ContinuationToken** | Pointer to **string** | An obfuscated string that references the last context instance on the previous page of results. You can use this for pagination, however, we recommend using the &lt;code&gt;next&lt;/code&gt; link instead. | [optional] 
**Items** | [**[]ContextRecord**](ContextRecord.md) | A collection of contexts. Can include multiple versions of contexts that have the same &lt;code&gt;kind&lt;/code&gt; and &lt;code&gt;key&lt;/code&gt;, but different &lt;code&gt;applicationId&lt;/code&gt;s. | 

## Methods

### NewContexts

`func NewContexts(totalCount int32, environmentId string, items []ContextRecord, ) *Contexts`

NewContexts instantiates a new Contexts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextsWithDefaults

`func NewContextsWithDefaults() *Contexts`

NewContextsWithDefaults instantiates a new Contexts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Contexts) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Contexts) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Contexts) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Contexts) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalCount

`func (o *Contexts) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Contexts) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Contexts) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetEnvironmentId

`func (o *Contexts) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Contexts) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Contexts) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetContinuationToken

`func (o *Contexts) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *Contexts) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *Contexts) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *Contexts) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetItems

`func (o *Contexts) GetItems() []ContextRecord`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Contexts) GetItemsOk() (*[]ContextRecord, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Contexts) SetItems(v []ContextRecord)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


