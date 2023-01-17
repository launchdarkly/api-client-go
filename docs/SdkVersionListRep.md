# SdkVersionListRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | **map[string]interface{}** | The location and content type of related resources | 
**SdkVersions** | [**[]SdkVersionRep**](SdkVersionRep.md) | The list of SDK names and versions | 

## Methods

### NewSdkVersionListRep

`func NewSdkVersionListRep(links map[string]interface{}, sdkVersions []SdkVersionRep, ) *SdkVersionListRep`

NewSdkVersionListRep instantiates a new SdkVersionListRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkVersionListRepWithDefaults

`func NewSdkVersionListRepWithDefaults() *SdkVersionListRep`

NewSdkVersionListRepWithDefaults instantiates a new SdkVersionListRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SdkVersionListRep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SdkVersionListRep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SdkVersionListRep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.


### GetSdkVersions

`func (o *SdkVersionListRep) GetSdkVersions() []SdkVersionRep`

GetSdkVersions returns the SdkVersions field if non-nil, zero value otherwise.

### GetSdkVersionsOk

`func (o *SdkVersionListRep) GetSdkVersionsOk() (*[]SdkVersionRep, bool)`

GetSdkVersionsOk returns a tuple with the SdkVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersions

`func (o *SdkVersionListRep) SetSdkVersions(v []SdkVersionRep)`

SetSdkVersions sets SdkVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


