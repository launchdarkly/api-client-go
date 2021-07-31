# ApiStatisticRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SourceLink** | **string** |  | 
**DefaultBranch** | **string** |  | 
**Enabled** | **bool** |  | 
**Version** | **int32** |  | 
**HunkCount** | **int32** |  | 
**Links** | [**map[string]CoreLink**](CoreLink.md) |  | 

## Methods

### NewApiStatisticRep

`func NewApiStatisticRep(name string, sourceLink string, defaultBranch string, enabled bool, version int32, hunkCount int32, links map[string]CoreLink, ) *ApiStatisticRep`

NewApiStatisticRep instantiates a new ApiStatisticRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStatisticRepWithDefaults

`func NewApiStatisticRepWithDefaults() *ApiStatisticRep`

NewApiStatisticRepWithDefaults instantiates a new ApiStatisticRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiStatisticRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiStatisticRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiStatisticRep) SetName(v string)`

SetName sets Name field to given value.


### GetSourceLink

`func (o *ApiStatisticRep) GetSourceLink() string`

GetSourceLink returns the SourceLink field if non-nil, zero value otherwise.

### GetSourceLinkOk

`func (o *ApiStatisticRep) GetSourceLinkOk() (*string, bool)`

GetSourceLinkOk returns a tuple with the SourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLink

`func (o *ApiStatisticRep) SetSourceLink(v string)`

SetSourceLink sets SourceLink field to given value.


### GetDefaultBranch

`func (o *ApiStatisticRep) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ApiStatisticRep) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ApiStatisticRep) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetEnabled

`func (o *ApiStatisticRep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiStatisticRep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiStatisticRep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVersion

`func (o *ApiStatisticRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiStatisticRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiStatisticRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetHunkCount

`func (o *ApiStatisticRep) GetHunkCount() int32`

GetHunkCount returns the HunkCount field if non-nil, zero value otherwise.

### GetHunkCountOk

`func (o *ApiStatisticRep) GetHunkCountOk() (*int32, bool)`

GetHunkCountOk returns a tuple with the HunkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunkCount

`func (o *ApiStatisticRep) SetHunkCount(v int32)`

SetHunkCount sets HunkCount field to given value.


### GetLinks

`func (o *ApiStatisticRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiStatisticRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiStatisticRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


