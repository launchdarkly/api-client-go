# DependentMeasuredRolloutRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The guarded rollout measured rollout Id | 
**FlagKey** | **string** | The guarded rollout flag key  | 
**FlagName** | **string** | The guarded rollout flag name  | 
**FlagPurpose** | Pointer to **string** | The guarded rollout flag purpose | [optional] 
**EnvironmentKey** | **string** | The guarded rollout environment key | 
**EnvironmentName** | **string** | The guarded rollout environment name | 
**Status** | **string** | The guarded rollout status | 
**CreationDate** | **int64** |  | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewDependentMeasuredRolloutRep

`func NewDependentMeasuredRolloutRep(id string, flagKey string, flagName string, environmentKey string, environmentName string, status string, creationDate int64, links map[string]Link, ) *DependentMeasuredRolloutRep`

NewDependentMeasuredRolloutRep instantiates a new DependentMeasuredRolloutRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentMeasuredRolloutRepWithDefaults

`func NewDependentMeasuredRolloutRepWithDefaults() *DependentMeasuredRolloutRep`

NewDependentMeasuredRolloutRepWithDefaults instantiates a new DependentMeasuredRolloutRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DependentMeasuredRolloutRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DependentMeasuredRolloutRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DependentMeasuredRolloutRep) SetId(v string)`

SetId sets Id field to given value.


### GetFlagKey

`func (o *DependentMeasuredRolloutRep) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *DependentMeasuredRolloutRep) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *DependentMeasuredRolloutRep) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetFlagName

`func (o *DependentMeasuredRolloutRep) GetFlagName() string`

GetFlagName returns the FlagName field if non-nil, zero value otherwise.

### GetFlagNameOk

`func (o *DependentMeasuredRolloutRep) GetFlagNameOk() (*string, bool)`

GetFlagNameOk returns a tuple with the FlagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagName

`func (o *DependentMeasuredRolloutRep) SetFlagName(v string)`

SetFlagName sets FlagName field to given value.


### GetFlagPurpose

`func (o *DependentMeasuredRolloutRep) GetFlagPurpose() string`

GetFlagPurpose returns the FlagPurpose field if non-nil, zero value otherwise.

### GetFlagPurposeOk

`func (o *DependentMeasuredRolloutRep) GetFlagPurposeOk() (*string, bool)`

GetFlagPurposeOk returns a tuple with the FlagPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagPurpose

`func (o *DependentMeasuredRolloutRep) SetFlagPurpose(v string)`

SetFlagPurpose sets FlagPurpose field to given value.

### HasFlagPurpose

`func (o *DependentMeasuredRolloutRep) HasFlagPurpose() bool`

HasFlagPurpose returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *DependentMeasuredRolloutRep) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *DependentMeasuredRolloutRep) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *DependentMeasuredRolloutRep) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetEnvironmentName

`func (o *DependentMeasuredRolloutRep) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *DependentMeasuredRolloutRep) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *DependentMeasuredRolloutRep) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetStatus

`func (o *DependentMeasuredRolloutRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DependentMeasuredRolloutRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DependentMeasuredRolloutRep) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreationDate

`func (o *DependentMeasuredRolloutRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *DependentMeasuredRolloutRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *DependentMeasuredRolloutRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLinks

`func (o *DependentMeasuredRolloutRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentMeasuredRolloutRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentMeasuredRolloutRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


