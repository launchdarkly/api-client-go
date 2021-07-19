# SegmentRuleRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Clauses** | Pointer to [**[]FlagConfigurationRepClauses**](FlagConfigurationRepClauses.md) |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**BucketBy** | Pointer to **string** |  | [optional] 

## Methods

### NewSegmentRuleRep

`func NewSegmentRuleRep() *SegmentRuleRep`

NewSegmentRuleRep instantiates a new SegmentRuleRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentRuleRepWithDefaults

`func NewSegmentRuleRepWithDefaults() *SegmentRuleRep`

NewSegmentRuleRepWithDefaults instantiates a new SegmentRuleRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SegmentRuleRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentRuleRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentRuleRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SegmentRuleRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClauses

`func (o *SegmentRuleRep) GetClauses() []FlagConfigurationRepClauses`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *SegmentRuleRep) GetClausesOk() (*[]FlagConfigurationRepClauses, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *SegmentRuleRep) SetClauses(v []FlagConfigurationRepClauses)`

SetClauses sets Clauses field to given value.

### HasClauses

`func (o *SegmentRuleRep) HasClauses() bool`

HasClauses returns a boolean if a field has been set.

### GetWeight

`func (o *SegmentRuleRep) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SegmentRuleRep) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SegmentRuleRep) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *SegmentRuleRep) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetBucketBy

`func (o *SegmentRuleRep) GetBucketBy() string`

GetBucketBy returns the BucketBy field if non-nil, zero value otherwise.

### GetBucketByOk

`func (o *SegmentRuleRep) GetBucketByOk() (*string, bool)`

GetBucketByOk returns a tuple with the BucketBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketBy

`func (o *SegmentRuleRep) SetBucketBy(v string)`

SetBucketBy sets BucketBy field to given value.

### HasBucketBy

`func (o *SegmentRuleRep) HasBucketBy() bool`

HasBucketBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


