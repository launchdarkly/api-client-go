# VersionsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidVersions** | **[]int32** |  | 
**LatestVersion** | **int32** |  | 
**CurrentVersion** | **int32** |  | 
**Beta** | Pointer to **bool** |  | [optional] 

## Methods

### NewVersionsRep

`func NewVersionsRep(validVersions []int32, latestVersion int32, currentVersion int32, ) *VersionsRep`

NewVersionsRep instantiates a new VersionsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionsRepWithDefaults

`func NewVersionsRepWithDefaults() *VersionsRep`

NewVersionsRepWithDefaults instantiates a new VersionsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidVersions

`func (o *VersionsRep) GetValidVersions() []int32`

GetValidVersions returns the ValidVersions field if non-nil, zero value otherwise.

### GetValidVersionsOk

`func (o *VersionsRep) GetValidVersionsOk() (*[]int32, bool)`

GetValidVersionsOk returns a tuple with the ValidVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidVersions

`func (o *VersionsRep) SetValidVersions(v []int32)`

SetValidVersions sets ValidVersions field to given value.


### GetLatestVersion

`func (o *VersionsRep) GetLatestVersion() int32`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *VersionsRep) GetLatestVersionOk() (*int32, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *VersionsRep) SetLatestVersion(v int32)`

SetLatestVersion sets LatestVersion field to given value.


### GetCurrentVersion

`func (o *VersionsRep) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *VersionsRep) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *VersionsRep) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetBeta

`func (o *VersionsRep) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *VersionsRep) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *VersionsRep) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *VersionsRep) HasBeta() bool`

HasBeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


