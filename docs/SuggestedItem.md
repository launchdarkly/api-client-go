# SuggestedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key of the flag | 
**Name** | **string** | Name of the flag | 
**Similarity** | **float32** | Similarity (-1 is not similar at all, 1 is extremely similar) of the flag to the generated flag | 

## Methods

### NewSuggestedItem

`func NewSuggestedItem(key string, name string, similarity float32, ) *SuggestedItem`

NewSuggestedItem instantiates a new SuggestedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedItemWithDefaults

`func NewSuggestedItemWithDefaults() *SuggestedItem`

NewSuggestedItemWithDefaults instantiates a new SuggestedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SuggestedItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SuggestedItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SuggestedItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *SuggestedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SuggestedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SuggestedItem) SetName(v string)`

SetName sets Name field to given value.


### GetSimilarity

`func (o *SuggestedItem) GetSimilarity() float32`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *SuggestedItem) GetSimilarityOk() (*float32, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *SuggestedItem) SetSimilarity(v float32)`

SetSimilarity sets Similarity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


