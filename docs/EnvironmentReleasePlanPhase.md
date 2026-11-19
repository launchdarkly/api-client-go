# EnvironmentReleasePlanPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The phase&#39;s server-generated identifier | [optional] [readonly] 
**Description** | Pointer to **string** | Human-readable description of this phase | [optional] 
**Steps** | [**[]EnvironmentReleasePlanStep**](EnvironmentReleasePlanStep.md) | Ordered allocation increments within this phase | 
**Clauses** | Pointer to [**[]EnvironmentReleasePlanClause**](EnvironmentReleasePlanClause.md) | Targeting clauses scoping this phase&#39;s audience. Each phase&#39;s clauses must include the randomization unit clause {\&quot;contextKind\&quot;: \&quot;&lt;randomizationUnit&gt;\&quot;, \&quot;attribute\&quot;: \&quot;kind\&quot;, \&quot;op\&quot;: \&quot;in\&quot;, \&quot;values\&quot;: [\&quot;&lt;randomizationUnit&gt;\&quot;]}, so the release only reaches contexts with that context kind. To release to every context of the randomization unit, provide exactly that clause. | [optional] 

## Methods

### NewEnvironmentReleasePlanPhase

`func NewEnvironmentReleasePlanPhase(steps []EnvironmentReleasePlanStep, ) *EnvironmentReleasePlanPhase`

NewEnvironmentReleasePlanPhase instantiates a new EnvironmentReleasePlanPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePlanPhaseWithDefaults

`func NewEnvironmentReleasePlanPhaseWithDefaults() *EnvironmentReleasePlanPhase`

NewEnvironmentReleasePlanPhaseWithDefaults instantiates a new EnvironmentReleasePlanPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentReleasePlanPhase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentReleasePlanPhase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentReleasePlanPhase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentReleasePlanPhase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentReleasePlanPhase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentReleasePlanPhase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentReleasePlanPhase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentReleasePlanPhase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSteps

`func (o *EnvironmentReleasePlanPhase) GetSteps() []EnvironmentReleasePlanStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *EnvironmentReleasePlanPhase) GetStepsOk() (*[]EnvironmentReleasePlanStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *EnvironmentReleasePlanPhase) SetSteps(v []EnvironmentReleasePlanStep)`

SetSteps sets Steps field to given value.


### GetClauses

`func (o *EnvironmentReleasePlanPhase) GetClauses() []EnvironmentReleasePlanClause`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *EnvironmentReleasePlanPhase) GetClausesOk() (*[]EnvironmentReleasePlanClause, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *EnvironmentReleasePlanPhase) SetClauses(v []EnvironmentReleasePlanClause)`

SetClauses sets Clauses field to given value.

### HasClauses

`func (o *EnvironmentReleasePlanPhase) HasClauses() bool`

HasClauses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


