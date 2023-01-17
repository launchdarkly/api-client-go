# ContextInstanceRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSeen** | Pointer to **time.Time** | Timestamp of the last time an evaluation occurred for this context instance | [optional] 
**Id** | **string** | The context instance ID | 
**ApplicationId** | Pointer to **string** | An identifier representing the application where the LaunchDarkly SDK is running | [optional] 
**AnonymousKinds** | Pointer to **[]string** | A list of the context kinds this context was associated with that the SDK removed because they were marked as anonymous at flag evaluation | [optional] 
**Context** | **interface{}** | The context, including its kind and attributes | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 

## Methods

### NewContextInstanceRecord

`func NewContextInstanceRecord(id string, context interface{}, ) *ContextInstanceRecord`

NewContextInstanceRecord instantiates a new ContextInstanceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextInstanceRecordWithDefaults

`func NewContextInstanceRecordWithDefaults() *ContextInstanceRecord`

NewContextInstanceRecordWithDefaults instantiates a new ContextInstanceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSeen

`func (o *ContextInstanceRecord) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ContextInstanceRecord) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ContextInstanceRecord) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ContextInstanceRecord) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetId

`func (o *ContextInstanceRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContextInstanceRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContextInstanceRecord) SetId(v string)`

SetId sets Id field to given value.


### GetApplicationId

`func (o *ContextInstanceRecord) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ContextInstanceRecord) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ContextInstanceRecord) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ContextInstanceRecord) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAnonymousKinds

`func (o *ContextInstanceRecord) GetAnonymousKinds() []string`

GetAnonymousKinds returns the AnonymousKinds field if non-nil, zero value otherwise.

### GetAnonymousKindsOk

`func (o *ContextInstanceRecord) GetAnonymousKindsOk() (*[]string, bool)`

GetAnonymousKindsOk returns a tuple with the AnonymousKinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousKinds

`func (o *ContextInstanceRecord) SetAnonymousKinds(v []string)`

SetAnonymousKinds sets AnonymousKinds field to given value.

### HasAnonymousKinds

`func (o *ContextInstanceRecord) HasAnonymousKinds() bool`

HasAnonymousKinds returns a boolean if a field has been set.

### GetContext

`func (o *ContextInstanceRecord) GetContext() interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ContextInstanceRecord) GetContextOk() (*interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ContextInstanceRecord) SetContext(v interface{})`

SetContext sets Context field to given value.


### SetContextNil

`func (o *ContextInstanceRecord) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *ContextInstanceRecord) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetLinks

`func (o *ContextInstanceRecord) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContextInstanceRecord) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContextInstanceRecord) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContextInstanceRecord) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccess

`func (o *ContextInstanceRecord) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ContextInstanceRecord) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ContextInstanceRecord) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ContextInstanceRecord) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


