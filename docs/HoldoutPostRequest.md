# HoldoutPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-friendly name for the holdout | [optional] 
**Key** | Pointer to **string** | A key that identifies the holdout | [optional] 
**Description** | Pointer to **string** | Description of the holdout | [optional] 
**Randomizationunit** | Pointer to **string** | The chosen randomization unit for the holdout base experiment | [optional] 
**Attributes** | Pointer to **[]string** | The attributes that the holdout iteration&#39;s results can be sliced by | [optional] 
**Holdoutamount** | Pointer to **string** | Audience allocation for the holdout | [optional] 
**Primarymetrickey** | Pointer to **string** | The key of the primary metric for this holdout | [optional] 
**Metrics** | Pointer to [**[]MetricInput**](MetricInput.md) | Details on the metrics for this experiment | [optional] 
**Prerequisiteflagkey** | Pointer to **string** | The key of the flag that the holdout is dependent on | [optional] 

## Methods

### NewHoldoutPostRequest

`func NewHoldoutPostRequest() *HoldoutPostRequest`

NewHoldoutPostRequest instantiates a new HoldoutPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldoutPostRequestWithDefaults

`func NewHoldoutPostRequestWithDefaults() *HoldoutPostRequest`

NewHoldoutPostRequestWithDefaults instantiates a new HoldoutPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HoldoutPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HoldoutPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HoldoutPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HoldoutPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *HoldoutPostRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *HoldoutPostRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *HoldoutPostRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *HoldoutPostRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *HoldoutPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldoutPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldoutPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HoldoutPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRandomizationunit

`func (o *HoldoutPostRequest) GetRandomizationunit() string`

GetRandomizationunit returns the Randomizationunit field if non-nil, zero value otherwise.

### GetRandomizationunitOk

`func (o *HoldoutPostRequest) GetRandomizationunitOk() (*string, bool)`

GetRandomizationunitOk returns a tuple with the Randomizationunit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationunit

`func (o *HoldoutPostRequest) SetRandomizationunit(v string)`

SetRandomizationunit sets Randomizationunit field to given value.

### HasRandomizationunit

`func (o *HoldoutPostRequest) HasRandomizationunit() bool`

HasRandomizationunit returns a boolean if a field has been set.

### GetAttributes

`func (o *HoldoutPostRequest) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *HoldoutPostRequest) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *HoldoutPostRequest) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *HoldoutPostRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetHoldoutamount

`func (o *HoldoutPostRequest) GetHoldoutamount() string`

GetHoldoutamount returns the Holdoutamount field if non-nil, zero value otherwise.

### GetHoldoutamountOk

`func (o *HoldoutPostRequest) GetHoldoutamountOk() (*string, bool)`

GetHoldoutamountOk returns a tuple with the Holdoutamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutamount

`func (o *HoldoutPostRequest) SetHoldoutamount(v string)`

SetHoldoutamount sets Holdoutamount field to given value.

### HasHoldoutamount

`func (o *HoldoutPostRequest) HasHoldoutamount() bool`

HasHoldoutamount returns a boolean if a field has been set.

### GetPrimarymetrickey

`func (o *HoldoutPostRequest) GetPrimarymetrickey() string`

GetPrimarymetrickey returns the Primarymetrickey field if non-nil, zero value otherwise.

### GetPrimarymetrickeyOk

`func (o *HoldoutPostRequest) GetPrimarymetrickeyOk() (*string, bool)`

GetPrimarymetrickeyOk returns a tuple with the Primarymetrickey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarymetrickey

`func (o *HoldoutPostRequest) SetPrimarymetrickey(v string)`

SetPrimarymetrickey sets Primarymetrickey field to given value.

### HasPrimarymetrickey

`func (o *HoldoutPostRequest) HasPrimarymetrickey() bool`

HasPrimarymetrickey returns a boolean if a field has been set.

### GetMetrics

`func (o *HoldoutPostRequest) GetMetrics() []MetricInput`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *HoldoutPostRequest) GetMetricsOk() (*[]MetricInput, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *HoldoutPostRequest) SetMetrics(v []MetricInput)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *HoldoutPostRequest) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetPrerequisiteflagkey

`func (o *HoldoutPostRequest) GetPrerequisiteflagkey() string`

GetPrerequisiteflagkey returns the Prerequisiteflagkey field if non-nil, zero value otherwise.

### GetPrerequisiteflagkeyOk

`func (o *HoldoutPostRequest) GetPrerequisiteflagkeyOk() (*string, bool)`

GetPrerequisiteflagkeyOk returns a tuple with the Prerequisiteflagkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteflagkey

`func (o *HoldoutPostRequest) SetPrerequisiteflagkey(v string)`

SetPrerequisiteflagkey sets Prerequisiteflagkey field to given value.

### HasPrerequisiteflagkey

`func (o *HoldoutPostRequest) HasPrerequisiteflagkey() bool`

HasPrerequisiteflagkey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


