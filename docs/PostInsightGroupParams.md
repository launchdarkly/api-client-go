# PostInsightGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the insight group | 
**Key** | **string** | The key of the insight group | 
**ProjectKey** | **string** | The projectKey to be associated with the insight group | 
**EnvironmentKey** | **string** | The environmentKey to be associated with the insight group | 
**ApplicationKeys** | Pointer to **[]string** | The application keys to associate with the insight group. If not provided, the insight group will include data from all applications. | [optional] 

## Methods

### NewPostInsightGroupParams

`func NewPostInsightGroupParams(name string, key string, projectKey string, environmentKey string, ) *PostInsightGroupParams`

NewPostInsightGroupParams instantiates a new PostInsightGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInsightGroupParamsWithDefaults

`func NewPostInsightGroupParamsWithDefaults() *PostInsightGroupParams`

NewPostInsightGroupParamsWithDefaults instantiates a new PostInsightGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostInsightGroupParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostInsightGroupParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostInsightGroupParams) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *PostInsightGroupParams) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostInsightGroupParams) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostInsightGroupParams) SetKey(v string)`

SetKey sets Key field to given value.


### GetProjectKey

`func (o *PostInsightGroupParams) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *PostInsightGroupParams) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *PostInsightGroupParams) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetEnvironmentKey

`func (o *PostInsightGroupParams) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *PostInsightGroupParams) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *PostInsightGroupParams) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetApplicationKeys

`func (o *PostInsightGroupParams) GetApplicationKeys() []string`

GetApplicationKeys returns the ApplicationKeys field if non-nil, zero value otherwise.

### GetApplicationKeysOk

`func (o *PostInsightGroupParams) GetApplicationKeysOk() (*[]string, bool)`

GetApplicationKeysOk returns a tuple with the ApplicationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKeys

`func (o *PostInsightGroupParams) SetApplicationKeys(v []string)`

SetApplicationKeys sets ApplicationKeys field to given value.

### HasApplicationKeys

`func (o *PostInsightGroupParams) HasApplicationKeys() bool`

HasApplicationKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


