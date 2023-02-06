# ContextInstanceEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the flag. | 
**Key** | **string** | Key of the flag. | 
**Value** | **interface{}** | The value of the flag variation that the context receives. If there is no defined default rule, this is null. | 
**Reason** | Pointer to [**ContextInstanceEvaluationReason**](ContextInstanceEvaluationReason.md) |  | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewContextInstanceEvaluation

`func NewContextInstanceEvaluation(name string, key string, value interface{}, links map[string]Link, ) *ContextInstanceEvaluation`

NewContextInstanceEvaluation instantiates a new ContextInstanceEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextInstanceEvaluationWithDefaults

`func NewContextInstanceEvaluationWithDefaults() *ContextInstanceEvaluation`

NewContextInstanceEvaluationWithDefaults instantiates a new ContextInstanceEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContextInstanceEvaluation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextInstanceEvaluation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextInstanceEvaluation) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ContextInstanceEvaluation) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContextInstanceEvaluation) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContextInstanceEvaluation) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *ContextInstanceEvaluation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContextInstanceEvaluation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContextInstanceEvaluation) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ContextInstanceEvaluation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ContextInstanceEvaluation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetReason

`func (o *ContextInstanceEvaluation) GetReason() ContextInstanceEvaluationReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ContextInstanceEvaluation) GetReasonOk() (*ContextInstanceEvaluationReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ContextInstanceEvaluation) SetReason(v ContextInstanceEvaluationReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ContextInstanceEvaluation) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetLinks

`func (o *ContextInstanceEvaluation) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContextInstanceEvaluation) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContextInstanceEvaluation) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


