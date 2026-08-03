# SdkKeysEnvironmentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Key** | **string** | A project-unique key for the environment | 
**Name** | **string** | A human-friendly name for the environment | 
**Color** | **string** | The color used to indicate this environment in the UI | 

## Methods

### NewSdkKeysEnvironmentSummary

`func NewSdkKeysEnvironmentSummary(links map[string]Link, key string, name string, color string, ) *SdkKeysEnvironmentSummary`

NewSdkKeysEnvironmentSummary instantiates a new SdkKeysEnvironmentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkKeysEnvironmentSummaryWithDefaults

`func NewSdkKeysEnvironmentSummaryWithDefaults() *SdkKeysEnvironmentSummary`

NewSdkKeysEnvironmentSummaryWithDefaults instantiates a new SdkKeysEnvironmentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SdkKeysEnvironmentSummary) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SdkKeysEnvironmentSummary) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SdkKeysEnvironmentSummary) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetKey

`func (o *SdkKeysEnvironmentSummary) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SdkKeysEnvironmentSummary) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SdkKeysEnvironmentSummary) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *SdkKeysEnvironmentSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkKeysEnvironmentSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkKeysEnvironmentSummary) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *SdkKeysEnvironmentSummary) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *SdkKeysEnvironmentSummary) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *SdkKeysEnvironmentSummary) SetColor(v string)`

SetColor sets Color field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


