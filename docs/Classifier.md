# Classifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Type** | **string** |  | 
**Instructions** | **string** |  | 
**Criteria** | Pointer to **interface{}** | Type-specific criteria. An enum classifier uses a map of option keys to descriptions. A float classifier uses an ordered list of level labels. A bool classifier uses optional true and false descriptions.  | [optional] 

## Methods

### NewClassifier

`func NewClassifier(key string, type_ string, instructions string, ) *Classifier`

NewClassifier instantiates a new Classifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifierWithDefaults

`func NewClassifierWithDefaults() *Classifier`

NewClassifierWithDefaults instantiates a new Classifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Classifier) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Classifier) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Classifier) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *Classifier) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Classifier) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Classifier) SetType(v string)`

SetType sets Type field to given value.


### GetInstructions

`func (o *Classifier) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Classifier) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Classifier) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetCriteria

`func (o *Classifier) GetCriteria() interface{}`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *Classifier) GetCriteriaOk() (*interface{}, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *Classifier) SetCriteria(v interface{})`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *Classifier) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### SetCriteriaNil

`func (o *Classifier) SetCriteriaNil(b bool)`

 SetCriteriaNil sets the value for Criteria to be an explicit nil

### UnsetCriteria
`func (o *Classifier) UnsetCriteria()`

UnsetCriteria ensures that no value is present for Criteria, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


