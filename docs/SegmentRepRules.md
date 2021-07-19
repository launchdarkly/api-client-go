# SegmentRepRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Clauses** | Pointer to [**[]FlagConfigurationRepClauses**](FlagConfigurationRepClauses.md) |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**BucketBy** | Pointer to **string** |  | [optional] 

## Methods

### NewSegmentRepRules

`func NewSegmentRepRules() *SegmentRepRules`

NewSegmentRepRules instantiates a new SegmentRepRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentRepRulesWithDefaults

`func NewSegmentRepRulesWithDefaults() *SegmentRepRules`

NewSegmentRepRulesWithDefaults instantiates a new SegmentRepRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SegmentRepRules) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentRepRules) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentRepRules) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SegmentRepRules) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClauses

`func (o *SegmentRepRules) GetClauses() []FlagConfigurationRepClauses`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *SegmentRepRules) GetClausesOk() (*[]FlagConfigurationRepClauses, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *SegmentRepRules) SetClauses(v []FlagConfigurationRepClauses)`

SetClauses sets Clauses field to given value.

### HasClauses

`func (o *SegmentRepRules) HasClauses() bool`

HasClauses returns a boolean if a field has been set.

### GetWeight

`func (o *SegmentRepRules) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SegmentRepRules) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SegmentRepRules) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *SegmentRepRules) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetBucketBy

`func (o *SegmentRepRules) GetBucketBy() string`

GetBucketBy returns the BucketBy field if non-nil, zero value otherwise.

### GetBucketByOk

`func (o *SegmentRepRules) GetBucketByOk() (*string, bool)`

GetBucketByOk returns a tuple with the BucketBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketBy

`func (o *SegmentRepRules) SetBucketBy(v string)`

SetBucketBy sets BucketBy field to given value.

### HasBucketBy

`func (o *SegmentRepRules) HasBucketBy() bool`

HasBucketBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


