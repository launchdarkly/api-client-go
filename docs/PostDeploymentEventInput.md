# PostDeploymentEventInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectKey** | **string** | The project key | 
**EnvironmentKey** | **string** | The environment key | 
**ApplicationKey** | **string** | The application key. This defines the granularity at which you want to view your insights metrics. Typically it is the name of one of the GitHub repositories that you use in this project.&lt;br/&gt;&lt;br/&gt;LaunchDarkly automatically creates a new application each time you send a unique application key. | 
**ApplicationName** | Pointer to **string** | The application name. This defines how the application is displayed | [optional] 
**ApplicationKind** | Pointer to **string** | The kind of application. Default: &lt;code&gt;server&lt;/code&gt; | [optional] 
**Version** | **string** | The application version. You can set the application version to any string that includes only letters, numbers, periods (&lt;code&gt;.&lt;/code&gt;), hyphens (&lt;code&gt;-&lt;/code&gt;), or underscores (&lt;code&gt;_&lt;/code&gt;).&lt;br/&gt;&lt;br/&gt;We recommend setting the application version to at least the first seven characters of the SHA or to the tag of the GitHub commit for this deployment. | 
**VersionName** | Pointer to **string** | The version name. This defines how the version is displayed | [optional] 
**EventType** | **string** | The event type | 
**EventTime** | Pointer to **int64** |  | [optional] 
**EventMetadata** | Pointer to **map[string]interface{}** | A JSON object containing metadata about the event | [optional] 
**DeploymentMetadata** | Pointer to **map[string]interface{}** | A JSON object containing metadata about the deployment | [optional] 

## Methods

### NewPostDeploymentEventInput

`func NewPostDeploymentEventInput(projectKey string, environmentKey string, applicationKey string, version string, eventType string, ) *PostDeploymentEventInput`

NewPostDeploymentEventInput instantiates a new PostDeploymentEventInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDeploymentEventInputWithDefaults

`func NewPostDeploymentEventInputWithDefaults() *PostDeploymentEventInput`

NewPostDeploymentEventInputWithDefaults instantiates a new PostDeploymentEventInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectKey

`func (o *PostDeploymentEventInput) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *PostDeploymentEventInput) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *PostDeploymentEventInput) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetEnvironmentKey

`func (o *PostDeploymentEventInput) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *PostDeploymentEventInput) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *PostDeploymentEventInput) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetApplicationKey

`func (o *PostDeploymentEventInput) GetApplicationKey() string`

GetApplicationKey returns the ApplicationKey field if non-nil, zero value otherwise.

### GetApplicationKeyOk

`func (o *PostDeploymentEventInput) GetApplicationKeyOk() (*string, bool)`

GetApplicationKeyOk returns a tuple with the ApplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKey

`func (o *PostDeploymentEventInput) SetApplicationKey(v string)`

SetApplicationKey sets ApplicationKey field to given value.


### GetApplicationName

`func (o *PostDeploymentEventInput) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *PostDeploymentEventInput) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *PostDeploymentEventInput) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *PostDeploymentEventInput) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationKind

`func (o *PostDeploymentEventInput) GetApplicationKind() string`

GetApplicationKind returns the ApplicationKind field if non-nil, zero value otherwise.

### GetApplicationKindOk

`func (o *PostDeploymentEventInput) GetApplicationKindOk() (*string, bool)`

GetApplicationKindOk returns a tuple with the ApplicationKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKind

`func (o *PostDeploymentEventInput) SetApplicationKind(v string)`

SetApplicationKind sets ApplicationKind field to given value.

### HasApplicationKind

`func (o *PostDeploymentEventInput) HasApplicationKind() bool`

HasApplicationKind returns a boolean if a field has been set.

### GetVersion

`func (o *PostDeploymentEventInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostDeploymentEventInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostDeploymentEventInput) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionName

`func (o *PostDeploymentEventInput) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *PostDeploymentEventInput) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *PostDeploymentEventInput) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *PostDeploymentEventInput) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetEventType

`func (o *PostDeploymentEventInput) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PostDeploymentEventInput) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PostDeploymentEventInput) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventTime

`func (o *PostDeploymentEventInput) GetEventTime() int64`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *PostDeploymentEventInput) GetEventTimeOk() (*int64, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *PostDeploymentEventInput) SetEventTime(v int64)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *PostDeploymentEventInput) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventMetadata

`func (o *PostDeploymentEventInput) GetEventMetadata() map[string]interface{}`

GetEventMetadata returns the EventMetadata field if non-nil, zero value otherwise.

### GetEventMetadataOk

`func (o *PostDeploymentEventInput) GetEventMetadataOk() (*map[string]interface{}, bool)`

GetEventMetadataOk returns a tuple with the EventMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMetadata

`func (o *PostDeploymentEventInput) SetEventMetadata(v map[string]interface{})`

SetEventMetadata sets EventMetadata field to given value.

### HasEventMetadata

`func (o *PostDeploymentEventInput) HasEventMetadata() bool`

HasEventMetadata returns a boolean if a field has been set.

### GetDeploymentMetadata

`func (o *PostDeploymentEventInput) GetDeploymentMetadata() map[string]interface{}`

GetDeploymentMetadata returns the DeploymentMetadata field if non-nil, zero value otherwise.

### GetDeploymentMetadataOk

`func (o *PostDeploymentEventInput) GetDeploymentMetadataOk() (*map[string]interface{}, bool)`

GetDeploymentMetadataOk returns a tuple with the DeploymentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMetadata

`func (o *PostDeploymentEventInput) SetDeploymentMetadata(v map[string]interface{})`

SetDeploymentMetadata sets DeploymentMetadata field to given value.

### HasDeploymentMetadata

`func (o *PostDeploymentEventInput) HasDeploymentMetadata() bool`

HasDeploymentMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


