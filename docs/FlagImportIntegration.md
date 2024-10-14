# FlagImportIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**FlagImportIntegrationLinks**](FlagImportIntegrationLinks.md) |  | 
**Id** | **string** | The integration ID | 
**IntegrationKey** | **string** | The integration key | 
**ProjectKey** | **string** | The project key | 
**Config** | **map[string]interface{}** |  | 
**Tags** | **[]string** | List of tags for this configuration | 
**Name** | **string** | Name of the configuration | 
**Version** | **int32** | Version of the current configuration | 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**Status** | [**FlagImportStatus**](FlagImportStatus.md) |  | 

## Methods

### NewFlagImportIntegration

`func NewFlagImportIntegration(links FlagImportIntegrationLinks, id string, integrationKey string, projectKey string, config map[string]interface{}, tags []string, name string, version int32, status FlagImportStatus, ) *FlagImportIntegration`

NewFlagImportIntegration instantiates a new FlagImportIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagImportIntegrationWithDefaults

`func NewFlagImportIntegrationWithDefaults() *FlagImportIntegration`

NewFlagImportIntegrationWithDefaults instantiates a new FlagImportIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlagImportIntegration) GetLinks() FlagImportIntegrationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagImportIntegration) GetLinksOk() (*FlagImportIntegrationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagImportIntegration) SetLinks(v FlagImportIntegrationLinks)`

SetLinks sets Links field to given value.


### GetId

`func (o *FlagImportIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagImportIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagImportIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationKey

`func (o *FlagImportIntegration) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *FlagImportIntegration) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *FlagImportIntegration) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.


### GetProjectKey

`func (o *FlagImportIntegration) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *FlagImportIntegration) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *FlagImportIntegration) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetConfig

`func (o *FlagImportIntegration) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FlagImportIntegration) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FlagImportIntegration) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetTags

`func (o *FlagImportIntegration) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlagImportIntegration) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlagImportIntegration) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetName

`func (o *FlagImportIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagImportIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagImportIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *FlagImportIntegration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FlagImportIntegration) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FlagImportIntegration) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetAccess

`func (o *FlagImportIntegration) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FlagImportIntegration) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FlagImportIntegration) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FlagImportIntegration) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetStatus

`func (o *FlagImportIntegration) GetStatus() FlagImportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlagImportIntegration) GetStatusOk() (*FlagImportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlagImportIntegration) SetStatus(v FlagImportStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


