# InsightGroupScores

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overall** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**DeploymentFrequency** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**DeploymentFailureRate** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**LeadTime** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**ImpactSize** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**ExperimentationCoverage** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**FlagHealth** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**Velocity** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**Risk** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**Efficiency** | [**InsightsMetricScore**](InsightsMetricScore.md) |  | 
**CreationRatio** | Pointer to [**InsightsMetricScore**](InsightsMetricScore.md) |  | [optional] 

## Methods

### NewInsightGroupScores

`func NewInsightGroupScores(overall InsightsMetricScore, deploymentFrequency InsightsMetricScore, deploymentFailureRate InsightsMetricScore, leadTime InsightsMetricScore, impactSize InsightsMetricScore, experimentationCoverage InsightsMetricScore, flagHealth InsightsMetricScore, velocity InsightsMetricScore, risk InsightsMetricScore, efficiency InsightsMetricScore, ) *InsightGroupScores`

NewInsightGroupScores instantiates a new InsightGroupScores object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightGroupScoresWithDefaults

`func NewInsightGroupScoresWithDefaults() *InsightGroupScores`

NewInsightGroupScoresWithDefaults instantiates a new InsightGroupScores object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverall

`func (o *InsightGroupScores) GetOverall() InsightsMetricScore`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InsightGroupScores) GetOverallOk() (*InsightsMetricScore, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InsightGroupScores) SetOverall(v InsightsMetricScore)`

SetOverall sets Overall field to given value.


### GetDeploymentFrequency

`func (o *InsightGroupScores) GetDeploymentFrequency() InsightsMetricScore`

GetDeploymentFrequency returns the DeploymentFrequency field if non-nil, zero value otherwise.

### GetDeploymentFrequencyOk

`func (o *InsightGroupScores) GetDeploymentFrequencyOk() (*InsightsMetricScore, bool)`

GetDeploymentFrequencyOk returns a tuple with the DeploymentFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentFrequency

`func (o *InsightGroupScores) SetDeploymentFrequency(v InsightsMetricScore)`

SetDeploymentFrequency sets DeploymentFrequency field to given value.


### GetDeploymentFailureRate

`func (o *InsightGroupScores) GetDeploymentFailureRate() InsightsMetricScore`

GetDeploymentFailureRate returns the DeploymentFailureRate field if non-nil, zero value otherwise.

### GetDeploymentFailureRateOk

`func (o *InsightGroupScores) GetDeploymentFailureRateOk() (*InsightsMetricScore, bool)`

GetDeploymentFailureRateOk returns a tuple with the DeploymentFailureRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentFailureRate

`func (o *InsightGroupScores) SetDeploymentFailureRate(v InsightsMetricScore)`

SetDeploymentFailureRate sets DeploymentFailureRate field to given value.


### GetLeadTime

`func (o *InsightGroupScores) GetLeadTime() InsightsMetricScore`

GetLeadTime returns the LeadTime field if non-nil, zero value otherwise.

### GetLeadTimeOk

`func (o *InsightGroupScores) GetLeadTimeOk() (*InsightsMetricScore, bool)`

GetLeadTimeOk returns a tuple with the LeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadTime

`func (o *InsightGroupScores) SetLeadTime(v InsightsMetricScore)`

SetLeadTime sets LeadTime field to given value.


### GetImpactSize

`func (o *InsightGroupScores) GetImpactSize() InsightsMetricScore`

GetImpactSize returns the ImpactSize field if non-nil, zero value otherwise.

### GetImpactSizeOk

`func (o *InsightGroupScores) GetImpactSizeOk() (*InsightsMetricScore, bool)`

GetImpactSizeOk returns a tuple with the ImpactSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactSize

`func (o *InsightGroupScores) SetImpactSize(v InsightsMetricScore)`

SetImpactSize sets ImpactSize field to given value.


### GetExperimentationCoverage

`func (o *InsightGroupScores) GetExperimentationCoverage() InsightsMetricScore`

GetExperimentationCoverage returns the ExperimentationCoverage field if non-nil, zero value otherwise.

### GetExperimentationCoverageOk

`func (o *InsightGroupScores) GetExperimentationCoverageOk() (*InsightsMetricScore, bool)`

GetExperimentationCoverageOk returns a tuple with the ExperimentationCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentationCoverage

`func (o *InsightGroupScores) SetExperimentationCoverage(v InsightsMetricScore)`

SetExperimentationCoverage sets ExperimentationCoverage field to given value.


### GetFlagHealth

`func (o *InsightGroupScores) GetFlagHealth() InsightsMetricScore`

GetFlagHealth returns the FlagHealth field if non-nil, zero value otherwise.

### GetFlagHealthOk

`func (o *InsightGroupScores) GetFlagHealthOk() (*InsightsMetricScore, bool)`

GetFlagHealthOk returns a tuple with the FlagHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagHealth

`func (o *InsightGroupScores) SetFlagHealth(v InsightsMetricScore)`

SetFlagHealth sets FlagHealth field to given value.


### GetVelocity

`func (o *InsightGroupScores) GetVelocity() InsightsMetricScore`

GetVelocity returns the Velocity field if non-nil, zero value otherwise.

### GetVelocityOk

`func (o *InsightGroupScores) GetVelocityOk() (*InsightsMetricScore, bool)`

GetVelocityOk returns a tuple with the Velocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocity

`func (o *InsightGroupScores) SetVelocity(v InsightsMetricScore)`

SetVelocity sets Velocity field to given value.


### GetRisk

`func (o *InsightGroupScores) GetRisk() InsightsMetricScore`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *InsightGroupScores) GetRiskOk() (*InsightsMetricScore, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *InsightGroupScores) SetRisk(v InsightsMetricScore)`

SetRisk sets Risk field to given value.


### GetEfficiency

`func (o *InsightGroupScores) GetEfficiency() InsightsMetricScore`

GetEfficiency returns the Efficiency field if non-nil, zero value otherwise.

### GetEfficiencyOk

`func (o *InsightGroupScores) GetEfficiencyOk() (*InsightsMetricScore, bool)`

GetEfficiencyOk returns a tuple with the Efficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfficiency

`func (o *InsightGroupScores) SetEfficiency(v InsightsMetricScore)`

SetEfficiency sets Efficiency field to given value.


### GetCreationRatio

`func (o *InsightGroupScores) GetCreationRatio() InsightsMetricScore`

GetCreationRatio returns the CreationRatio field if non-nil, zero value otherwise.

### GetCreationRatioOk

`func (o *InsightGroupScores) GetCreationRatioOk() (*InsightsMetricScore, bool)`

GetCreationRatioOk returns a tuple with the CreationRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationRatio

`func (o *InsightGroupScores) SetCreationRatio(v InsightsMetricScore)`

SetCreationRatio sets CreationRatio field to given value.

### HasCreationRatio

`func (o *InsightGroupScores) HasCreationRatio() bool`

HasCreationRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


