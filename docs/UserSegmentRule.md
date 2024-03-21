# UserSegmentRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Clauses** | [**[]Clause**](Clause.md) |  | 
**Weight** | Pointer to **int32** |  | [optional] 
**RolloutContextKind** | Pointer to **string** |  | [optional] 
**BucketBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewUserSegmentRule

`func NewUserSegmentRule(clauses []Clause, ) *UserSegmentRule`

NewUserSegmentRule instantiates a new UserSegmentRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSegmentRuleWithDefaults

`func NewUserSegmentRuleWithDefaults() *UserSegmentRule`

NewUserSegmentRuleWithDefaults instantiates a new UserSegmentRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSegmentRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSegmentRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSegmentRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserSegmentRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClauses

`func (o *UserSegmentRule) GetClauses() []Clause`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *UserSegmentRule) GetClausesOk() (*[]Clause, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *UserSegmentRule) SetClauses(v []Clause)`

SetClauses sets Clauses field to given value.


### GetWeight

`func (o *UserSegmentRule) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UserSegmentRule) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UserSegmentRule) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *UserSegmentRule) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetRolloutContextKind

`func (o *UserSegmentRule) GetRolloutContextKind() string`

GetRolloutContextKind returns the RolloutContextKind field if non-nil, zero value otherwise.

### GetRolloutContextKindOk

`func (o *UserSegmentRule) GetRolloutContextKindOk() (*string, bool)`

GetRolloutContextKindOk returns a tuple with the RolloutContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutContextKind

`func (o *UserSegmentRule) SetRolloutContextKind(v string)`

SetRolloutContextKind sets RolloutContextKind field to given value.

### HasRolloutContextKind

`func (o *UserSegmentRule) HasRolloutContextKind() bool`

HasRolloutContextKind returns a boolean if a field has been set.

### GetBucketBy

`func (o *UserSegmentRule) GetBucketBy() string`

GetBucketBy returns the BucketBy field if non-nil, zero value otherwise.

### GetBucketByOk

`func (o *UserSegmentRule) GetBucketByOk() (*string, bool)`

GetBucketByOk returns a tuple with the BucketBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketBy

`func (o *UserSegmentRule) SetBucketBy(v string)`

SetBucketBy sets BucketBy field to given value.

### HasBucketBy

`func (o *UserSegmentRule) HasBucketBy() bool`

HasBucketBy returns a boolean if a field has been set.

### GetDescription

`func (o *UserSegmentRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserSegmentRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserSegmentRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserSegmentRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


