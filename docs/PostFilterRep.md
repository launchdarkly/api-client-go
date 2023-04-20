# PostFilterRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The filter key | 
**Name** | **string** | The filter name | 
**Description** | **string** | Textual description of the filter | 
**Rules** | **[]map[string]interface{}** |  | 

## Methods

### NewPostFilterRep

`func NewPostFilterRep(key string, name string, description string, rules []map[string]interface{}, ) *PostFilterRep`

NewPostFilterRep instantiates a new PostFilterRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFilterRepWithDefaults

`func NewPostFilterRepWithDefaults() *PostFilterRep`

NewPostFilterRepWithDefaults instantiates a new PostFilterRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PostFilterRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostFilterRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostFilterRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *PostFilterRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostFilterRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostFilterRep) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostFilterRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostFilterRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostFilterRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRules

`func (o *PostFilterRep) GetRules() []map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PostFilterRep) GetRulesOk() (*[]map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PostFilterRep) SetRules(v []map[string]interface{})`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


