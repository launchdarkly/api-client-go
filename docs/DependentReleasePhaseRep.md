# DependentReleasePhaseRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The guarded rollout release phase ID | 
**FlagKey** | **string** | The guarded rollout flag key  | 
**FlagName** | **string** | The guarded rollout flag name  | 
**FlagPurpose** | Pointer to **string** | The guarded rollout flag purpose | [optional] 
**EnvironmentKey** | **string** | The guarded rollout environment key | 
**EnvironmentName** | **string** | The guarded rollout environment name | 
**Status** | **string** | The guarded rollout status | 
**CreationDate** | **int64** |  | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewDependentReleasePhaseRep

`func NewDependentReleasePhaseRep(id string, flagKey string, flagName string, environmentKey string, environmentName string, status string, creationDate int64, links map[string]Link, ) *DependentReleasePhaseRep`

NewDependentReleasePhaseRep instantiates a new DependentReleasePhaseRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentReleasePhaseRepWithDefaults

`func NewDependentReleasePhaseRepWithDefaults() *DependentReleasePhaseRep`

NewDependentReleasePhaseRepWithDefaults instantiates a new DependentReleasePhaseRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DependentReleasePhaseRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DependentReleasePhaseRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DependentReleasePhaseRep) SetId(v string)`

SetId sets Id field to given value.


### GetFlagKey

`func (o *DependentReleasePhaseRep) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *DependentReleasePhaseRep) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *DependentReleasePhaseRep) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetFlagName

`func (o *DependentReleasePhaseRep) GetFlagName() string`

GetFlagName returns the FlagName field if non-nil, zero value otherwise.

### GetFlagNameOk

`func (o *DependentReleasePhaseRep) GetFlagNameOk() (*string, bool)`

GetFlagNameOk returns a tuple with the FlagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagName

`func (o *DependentReleasePhaseRep) SetFlagName(v string)`

SetFlagName sets FlagName field to given value.


### GetFlagPurpose

`func (o *DependentReleasePhaseRep) GetFlagPurpose() string`

GetFlagPurpose returns the FlagPurpose field if non-nil, zero value otherwise.

### GetFlagPurposeOk

`func (o *DependentReleasePhaseRep) GetFlagPurposeOk() (*string, bool)`

GetFlagPurposeOk returns a tuple with the FlagPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagPurpose

`func (o *DependentReleasePhaseRep) SetFlagPurpose(v string)`

SetFlagPurpose sets FlagPurpose field to given value.

### HasFlagPurpose

`func (o *DependentReleasePhaseRep) HasFlagPurpose() bool`

HasFlagPurpose returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *DependentReleasePhaseRep) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *DependentReleasePhaseRep) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *DependentReleasePhaseRep) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetEnvironmentName

`func (o *DependentReleasePhaseRep) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *DependentReleasePhaseRep) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *DependentReleasePhaseRep) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetStatus

`func (o *DependentReleasePhaseRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DependentReleasePhaseRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DependentReleasePhaseRep) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreationDate

`func (o *DependentReleasePhaseRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *DependentReleasePhaseRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *DependentReleasePhaseRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLinks

`func (o *DependentReleasePhaseRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentReleasePhaseRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentReleasePhaseRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


