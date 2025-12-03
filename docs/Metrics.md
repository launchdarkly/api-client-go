# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputTokens** | Pointer to **int32** |  | [optional] 
**OutputTokens** | Pointer to **int32** |  | [optional] 
**TotalTokens** | Pointer to **int32** |  | [optional] 
**GenerationCount** | Pointer to **int32** | Number of attempted generations | [optional] 
**GenerationSuccessCount** | Pointer to **int32** | Number of successful generations | [optional] 
**GenerationErrorCount** | Pointer to **int32** | Number of generations with errors | [optional] 
**ThumbsUp** | Pointer to **int32** |  | [optional] 
**ThumbsDown** | Pointer to **int32** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**TimeToFirstTokenMs** | Pointer to **int32** |  | [optional] 
**SatisfactionRating** | Pointer to **float32** | A value between 0 and 1 representing satisfaction rating | [optional] 
**InputCost** | Pointer to **float64** | Cost of input tokens in USD | [optional] 
**OutputCost** | Pointer to **float64** | Cost of output tokens in USD | [optional] 
**JudgeAccuracy** | Pointer to **float32** | Average accuracy judge score (0.0-1.0) | [optional] 
**JudgeRelevance** | Pointer to **float32** | Average relevance judge score (0.0-1.0) | [optional] 
**JudgeToxicity** | Pointer to **float32** | Average toxicity judge score (0.0-1.0) | [optional] 

## Methods

### NewMetrics

`func NewMetrics() *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputTokens

`func (o *Metrics) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *Metrics) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *Metrics) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *Metrics) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetOutputTokens

`func (o *Metrics) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *Metrics) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *Metrics) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *Metrics) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetTotalTokens

`func (o *Metrics) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *Metrics) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *Metrics) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *Metrics) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetGenerationCount

`func (o *Metrics) GetGenerationCount() int32`

GetGenerationCount returns the GenerationCount field if non-nil, zero value otherwise.

### GetGenerationCountOk

`func (o *Metrics) GetGenerationCountOk() (*int32, bool)`

GetGenerationCountOk returns a tuple with the GenerationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationCount

`func (o *Metrics) SetGenerationCount(v int32)`

SetGenerationCount sets GenerationCount field to given value.

### HasGenerationCount

`func (o *Metrics) HasGenerationCount() bool`

HasGenerationCount returns a boolean if a field has been set.

### GetGenerationSuccessCount

`func (o *Metrics) GetGenerationSuccessCount() int32`

GetGenerationSuccessCount returns the GenerationSuccessCount field if non-nil, zero value otherwise.

### GetGenerationSuccessCountOk

`func (o *Metrics) GetGenerationSuccessCountOk() (*int32, bool)`

GetGenerationSuccessCountOk returns a tuple with the GenerationSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationSuccessCount

`func (o *Metrics) SetGenerationSuccessCount(v int32)`

SetGenerationSuccessCount sets GenerationSuccessCount field to given value.

### HasGenerationSuccessCount

`func (o *Metrics) HasGenerationSuccessCount() bool`

HasGenerationSuccessCount returns a boolean if a field has been set.

### GetGenerationErrorCount

`func (o *Metrics) GetGenerationErrorCount() int32`

GetGenerationErrorCount returns the GenerationErrorCount field if non-nil, zero value otherwise.

### GetGenerationErrorCountOk

`func (o *Metrics) GetGenerationErrorCountOk() (*int32, bool)`

GetGenerationErrorCountOk returns a tuple with the GenerationErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationErrorCount

`func (o *Metrics) SetGenerationErrorCount(v int32)`

SetGenerationErrorCount sets GenerationErrorCount field to given value.

### HasGenerationErrorCount

`func (o *Metrics) HasGenerationErrorCount() bool`

HasGenerationErrorCount returns a boolean if a field has been set.

### GetThumbsUp

`func (o *Metrics) GetThumbsUp() int32`

GetThumbsUp returns the ThumbsUp field if non-nil, zero value otherwise.

### GetThumbsUpOk

`func (o *Metrics) GetThumbsUpOk() (*int32, bool)`

GetThumbsUpOk returns a tuple with the ThumbsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbsUp

`func (o *Metrics) SetThumbsUp(v int32)`

SetThumbsUp sets ThumbsUp field to given value.

### HasThumbsUp

`func (o *Metrics) HasThumbsUp() bool`

HasThumbsUp returns a boolean if a field has been set.

### GetThumbsDown

`func (o *Metrics) GetThumbsDown() int32`

GetThumbsDown returns the ThumbsDown field if non-nil, zero value otherwise.

### GetThumbsDownOk

`func (o *Metrics) GetThumbsDownOk() (*int32, bool)`

GetThumbsDownOk returns a tuple with the ThumbsDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbsDown

`func (o *Metrics) SetThumbsDown(v int32)`

SetThumbsDown sets ThumbsDown field to given value.

### HasThumbsDown

`func (o *Metrics) HasThumbsDown() bool`

HasThumbsDown returns a boolean if a field has been set.

### GetDurationMs

`func (o *Metrics) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *Metrics) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *Metrics) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *Metrics) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetTimeToFirstTokenMs

`func (o *Metrics) GetTimeToFirstTokenMs() int32`

GetTimeToFirstTokenMs returns the TimeToFirstTokenMs field if non-nil, zero value otherwise.

### GetTimeToFirstTokenMsOk

`func (o *Metrics) GetTimeToFirstTokenMsOk() (*int32, bool)`

GetTimeToFirstTokenMsOk returns a tuple with the TimeToFirstTokenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToFirstTokenMs

`func (o *Metrics) SetTimeToFirstTokenMs(v int32)`

SetTimeToFirstTokenMs sets TimeToFirstTokenMs field to given value.

### HasTimeToFirstTokenMs

`func (o *Metrics) HasTimeToFirstTokenMs() bool`

HasTimeToFirstTokenMs returns a boolean if a field has been set.

### GetSatisfactionRating

`func (o *Metrics) GetSatisfactionRating() float32`

GetSatisfactionRating returns the SatisfactionRating field if non-nil, zero value otherwise.

### GetSatisfactionRatingOk

`func (o *Metrics) GetSatisfactionRatingOk() (*float32, bool)`

GetSatisfactionRatingOk returns a tuple with the SatisfactionRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfactionRating

`func (o *Metrics) SetSatisfactionRating(v float32)`

SetSatisfactionRating sets SatisfactionRating field to given value.

### HasSatisfactionRating

`func (o *Metrics) HasSatisfactionRating() bool`

HasSatisfactionRating returns a boolean if a field has been set.

### GetInputCost

`func (o *Metrics) GetInputCost() float64`

GetInputCost returns the InputCost field if non-nil, zero value otherwise.

### GetInputCostOk

`func (o *Metrics) GetInputCostOk() (*float64, bool)`

GetInputCostOk returns a tuple with the InputCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCost

`func (o *Metrics) SetInputCost(v float64)`

SetInputCost sets InputCost field to given value.

### HasInputCost

`func (o *Metrics) HasInputCost() bool`

HasInputCost returns a boolean if a field has been set.

### GetOutputCost

`func (o *Metrics) GetOutputCost() float64`

GetOutputCost returns the OutputCost field if non-nil, zero value otherwise.

### GetOutputCostOk

`func (o *Metrics) GetOutputCostOk() (*float64, bool)`

GetOutputCostOk returns a tuple with the OutputCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCost

`func (o *Metrics) SetOutputCost(v float64)`

SetOutputCost sets OutputCost field to given value.

### HasOutputCost

`func (o *Metrics) HasOutputCost() bool`

HasOutputCost returns a boolean if a field has been set.

### GetJudgeAccuracy

`func (o *Metrics) GetJudgeAccuracy() float32`

GetJudgeAccuracy returns the JudgeAccuracy field if non-nil, zero value otherwise.

### GetJudgeAccuracyOk

`func (o *Metrics) GetJudgeAccuracyOk() (*float32, bool)`

GetJudgeAccuracyOk returns a tuple with the JudgeAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeAccuracy

`func (o *Metrics) SetJudgeAccuracy(v float32)`

SetJudgeAccuracy sets JudgeAccuracy field to given value.

### HasJudgeAccuracy

`func (o *Metrics) HasJudgeAccuracy() bool`

HasJudgeAccuracy returns a boolean if a field has been set.

### GetJudgeRelevance

`func (o *Metrics) GetJudgeRelevance() float32`

GetJudgeRelevance returns the JudgeRelevance field if non-nil, zero value otherwise.

### GetJudgeRelevanceOk

`func (o *Metrics) GetJudgeRelevanceOk() (*float32, bool)`

GetJudgeRelevanceOk returns a tuple with the JudgeRelevance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeRelevance

`func (o *Metrics) SetJudgeRelevance(v float32)`

SetJudgeRelevance sets JudgeRelevance field to given value.

### HasJudgeRelevance

`func (o *Metrics) HasJudgeRelevance() bool`

HasJudgeRelevance returns a boolean if a field has been set.

### GetJudgeToxicity

`func (o *Metrics) GetJudgeToxicity() float32`

GetJudgeToxicity returns the JudgeToxicity field if non-nil, zero value otherwise.

### GetJudgeToxicityOk

`func (o *Metrics) GetJudgeToxicityOk() (*float32, bool)`

GetJudgeToxicityOk returns a tuple with the JudgeToxicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeToxicity

`func (o *Metrics) SetJudgeToxicity(v float32)`

SetJudgeToxicity sets JudgeToxicity field to given value.

### HasJudgeToxicity

`func (o *Metrics) HasJudgeToxicity() bool`

HasJudgeToxicity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


