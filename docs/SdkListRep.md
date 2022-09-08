# SdkListRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | **map[string]interface{}** | The location and content type of related resources | 
**Sdks** | **[]string** | The list of SDK names | 

## Methods

### NewSdkListRep

`func NewSdkListRep(links map[string]interface{}, sdks []string, ) *SdkListRep`

NewSdkListRep instantiates a new SdkListRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkListRepWithDefaults

`func NewSdkListRepWithDefaults() *SdkListRep`

NewSdkListRepWithDefaults instantiates a new SdkListRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SdkListRep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SdkListRep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SdkListRep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.


### GetSdks

`func (o *SdkListRep) GetSdks() []string`

GetSdks returns the Sdks field if non-nil, zero value otherwise.

### GetSdksOk

`func (o *SdkListRep) GetSdksOk() (*[]string, bool)`

GetSdksOk returns a tuple with the Sdks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdks

`func (o *SdkListRep) SetSdks(v []string)`

SetSdks sets Sdks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


