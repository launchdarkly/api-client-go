# ContextRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSeen** | Pointer to **time.Time** | Timestamp of the last time an evaluation occurred for this context | [optional] 
**ApplicationId** | Pointer to **string** | An identifier representing the application where the LaunchDarkly SDK is running | [optional] 
**Context** | **interface{}** | The context, including its kind and attributes | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**AssociatedContexts** | Pointer to **int32** | The total number of associated contexts. Associated contexts are contexts that have appeared in the same context instance, that is, they were part of the same flag evaluation. | [optional] 

## Methods

### NewContextRecord

`func NewContextRecord(context interface{}, ) *ContextRecord`

NewContextRecord instantiates a new ContextRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextRecordWithDefaults

`func NewContextRecordWithDefaults() *ContextRecord`

NewContextRecordWithDefaults instantiates a new ContextRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSeen

`func (o *ContextRecord) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ContextRecord) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ContextRecord) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ContextRecord) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetApplicationId

`func (o *ContextRecord) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ContextRecord) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ContextRecord) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ContextRecord) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetContext

`func (o *ContextRecord) GetContext() interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ContextRecord) GetContextOk() (*interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ContextRecord) SetContext(v interface{})`

SetContext sets Context field to given value.


### SetContextNil

`func (o *ContextRecord) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *ContextRecord) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetLinks

`func (o *ContextRecord) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContextRecord) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContextRecord) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContextRecord) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccess

`func (o *ContextRecord) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ContextRecord) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ContextRecord) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ContextRecord) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAssociatedContexts

`func (o *ContextRecord) GetAssociatedContexts() int32`

GetAssociatedContexts returns the AssociatedContexts field if non-nil, zero value otherwise.

### GetAssociatedContextsOk

`func (o *ContextRecord) GetAssociatedContextsOk() (*int32, bool)`

GetAssociatedContextsOk returns a tuple with the AssociatedContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedContexts

`func (o *ContextRecord) SetAssociatedContexts(v int32)`

SetAssociatedContexts sets AssociatedContexts field to given value.

### HasAssociatedContexts

`func (o *ContextRecord) HasAssociatedContexts() bool`

HasAssociatedContexts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


