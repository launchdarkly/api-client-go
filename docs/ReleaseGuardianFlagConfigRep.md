# ReleaseGuardianFlagConfigRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**[]MetricV2Rep**](MetricV2Rep.md) |  | [optional] 
**RandomizationUnit** | Pointer to **string** |  | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewReleaseGuardianFlagConfigRep

`func NewReleaseGuardianFlagConfigRep(links map[string]Link, ) *ReleaseGuardianFlagConfigRep`

NewReleaseGuardianFlagConfigRep instantiates a new ReleaseGuardianFlagConfigRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseGuardianFlagConfigRepWithDefaults

`func NewReleaseGuardianFlagConfigRepWithDefaults() *ReleaseGuardianFlagConfigRep`

NewReleaseGuardianFlagConfigRepWithDefaults instantiates a new ReleaseGuardianFlagConfigRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *ReleaseGuardianFlagConfigRep) GetMetrics() []MetricV2Rep`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ReleaseGuardianFlagConfigRep) GetMetricsOk() (*[]MetricV2Rep, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ReleaseGuardianFlagConfigRep) SetMetrics(v []MetricV2Rep)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ReleaseGuardianFlagConfigRep) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetRandomizationUnit

`func (o *ReleaseGuardianFlagConfigRep) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *ReleaseGuardianFlagConfigRep) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *ReleaseGuardianFlagConfigRep) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *ReleaseGuardianFlagConfigRep) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.

### GetLinks

`func (o *ReleaseGuardianFlagConfigRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReleaseGuardianFlagConfigRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReleaseGuardianFlagConfigRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


