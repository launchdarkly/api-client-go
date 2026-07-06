# SdkVersionDetailsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ProjectKey** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**EnvironmentKey** | Pointer to **string** |  | [optional] 
**EnvironmentName** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**LdLatestVersion** | Pointer to **string** |  | [optional] 
**EolStatus** | Pointer to **string** | The end of life (EOL) status of the SDK version. Possible values are: &lt;br/&gt;- &lt;code&gt;EolAllClear&lt;/code&gt;: the SDK version is current&lt;br/&gt;- &lt;code&gt;EolNear&lt;/code&gt;: the SDK version is approaching EOL&lt;br/&gt;- &lt;code&gt;EolPast&lt;/code&gt;: the SDK version is past EOL&lt;br/&gt;- &lt;code&gt;MajorVersionAvailable&lt;/code&gt;: a new major version is available but the current version is not near EOL&lt;br/&gt;- &lt;code&gt;EolUnknown&lt;/code&gt;: the EOL status cannot be determined. | [optional] 
**LatestReleaseUrl** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**RelayVersion** | Pointer to **string** |  | [optional] 
**RelayEolStatus** | Pointer to **string** | The end of life status of the Relay Proxy version. Only present when the SDK connects through a Relay Proxy. Uses the same values as &lt;code&gt;eolStatus&lt;/code&gt;. | [optional] 
**RelayLatestVersion** | Pointer to **string** |  | [optional] 
**RelayLatestReleaseUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewSdkVersionDetailsRep

`func NewSdkVersionDetailsRep() *SdkVersionDetailsRep`

NewSdkVersionDetailsRep instantiates a new SdkVersionDetailsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkVersionDetailsRepWithDefaults

`func NewSdkVersionDetailsRepWithDefaults() *SdkVersionDetailsRep`

NewSdkVersionDetailsRepWithDefaults instantiates a new SdkVersionDetailsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SdkVersionDetailsRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkVersionDetailsRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkVersionDetailsRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdkVersionDetailsRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *SdkVersionDetailsRep) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SdkVersionDetailsRep) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SdkVersionDetailsRep) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SdkVersionDetailsRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *SdkVersionDetailsRep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SdkVersionDetailsRep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SdkVersionDetailsRep) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SdkVersionDetailsRep) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProjectId

`func (o *SdkVersionDetailsRep) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SdkVersionDetailsRep) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SdkVersionDetailsRep) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SdkVersionDetailsRep) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectKey

`func (o *SdkVersionDetailsRep) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *SdkVersionDetailsRep) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *SdkVersionDetailsRep) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *SdkVersionDetailsRep) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetProjectName

`func (o *SdkVersionDetailsRep) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *SdkVersionDetailsRep) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *SdkVersionDetailsRep) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *SdkVersionDetailsRep) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *SdkVersionDetailsRep) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SdkVersionDetailsRep) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SdkVersionDetailsRep) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *SdkVersionDetailsRep) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *SdkVersionDetailsRep) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *SdkVersionDetailsRep) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *SdkVersionDetailsRep) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *SdkVersionDetailsRep) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *SdkVersionDetailsRep) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *SdkVersionDetailsRep) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *SdkVersionDetailsRep) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *SdkVersionDetailsRep) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetApplicationId

`func (o *SdkVersionDetailsRep) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *SdkVersionDetailsRep) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *SdkVersionDetailsRep) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *SdkVersionDetailsRep) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetLdLatestVersion

`func (o *SdkVersionDetailsRep) GetLdLatestVersion() string`

GetLdLatestVersion returns the LdLatestVersion field if non-nil, zero value otherwise.

### GetLdLatestVersionOk

`func (o *SdkVersionDetailsRep) GetLdLatestVersionOk() (*string, bool)`

GetLdLatestVersionOk returns a tuple with the LdLatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdLatestVersion

`func (o *SdkVersionDetailsRep) SetLdLatestVersion(v string)`

SetLdLatestVersion sets LdLatestVersion field to given value.

### HasLdLatestVersion

`func (o *SdkVersionDetailsRep) HasLdLatestVersion() bool`

HasLdLatestVersion returns a boolean if a field has been set.

### GetEolStatus

`func (o *SdkVersionDetailsRep) GetEolStatus() string`

GetEolStatus returns the EolStatus field if non-nil, zero value otherwise.

### GetEolStatusOk

`func (o *SdkVersionDetailsRep) GetEolStatusOk() (*string, bool)`

GetEolStatusOk returns a tuple with the EolStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEolStatus

`func (o *SdkVersionDetailsRep) SetEolStatus(v string)`

SetEolStatus sets EolStatus field to given value.

### HasEolStatus

`func (o *SdkVersionDetailsRep) HasEolStatus() bool`

HasEolStatus returns a boolean if a field has been set.

### GetLatestReleaseUrl

`func (o *SdkVersionDetailsRep) GetLatestReleaseUrl() string`

GetLatestReleaseUrl returns the LatestReleaseUrl field if non-nil, zero value otherwise.

### GetLatestReleaseUrlOk

`func (o *SdkVersionDetailsRep) GetLatestReleaseUrlOk() (*string, bool)`

GetLatestReleaseUrlOk returns a tuple with the LatestReleaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReleaseUrl

`func (o *SdkVersionDetailsRep) SetLatestReleaseUrl(v string)`

SetLatestReleaseUrl sets LatestReleaseUrl field to given value.

### HasLatestReleaseUrl

`func (o *SdkVersionDetailsRep) HasLatestReleaseUrl() bool`

HasLatestReleaseUrl returns a boolean if a field has been set.

### GetConnectionType

`func (o *SdkVersionDetailsRep) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SdkVersionDetailsRep) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SdkVersionDetailsRep) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *SdkVersionDetailsRep) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetRelayVersion

`func (o *SdkVersionDetailsRep) GetRelayVersion() string`

GetRelayVersion returns the RelayVersion field if non-nil, zero value otherwise.

### GetRelayVersionOk

`func (o *SdkVersionDetailsRep) GetRelayVersionOk() (*string, bool)`

GetRelayVersionOk returns a tuple with the RelayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayVersion

`func (o *SdkVersionDetailsRep) SetRelayVersion(v string)`

SetRelayVersion sets RelayVersion field to given value.

### HasRelayVersion

`func (o *SdkVersionDetailsRep) HasRelayVersion() bool`

HasRelayVersion returns a boolean if a field has been set.

### GetRelayEolStatus

`func (o *SdkVersionDetailsRep) GetRelayEolStatus() string`

GetRelayEolStatus returns the RelayEolStatus field if non-nil, zero value otherwise.

### GetRelayEolStatusOk

`func (o *SdkVersionDetailsRep) GetRelayEolStatusOk() (*string, bool)`

GetRelayEolStatusOk returns a tuple with the RelayEolStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayEolStatus

`func (o *SdkVersionDetailsRep) SetRelayEolStatus(v string)`

SetRelayEolStatus sets RelayEolStatus field to given value.

### HasRelayEolStatus

`func (o *SdkVersionDetailsRep) HasRelayEolStatus() bool`

HasRelayEolStatus returns a boolean if a field has been set.

### GetRelayLatestVersion

`func (o *SdkVersionDetailsRep) GetRelayLatestVersion() string`

GetRelayLatestVersion returns the RelayLatestVersion field if non-nil, zero value otherwise.

### GetRelayLatestVersionOk

`func (o *SdkVersionDetailsRep) GetRelayLatestVersionOk() (*string, bool)`

GetRelayLatestVersionOk returns a tuple with the RelayLatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayLatestVersion

`func (o *SdkVersionDetailsRep) SetRelayLatestVersion(v string)`

SetRelayLatestVersion sets RelayLatestVersion field to given value.

### HasRelayLatestVersion

`func (o *SdkVersionDetailsRep) HasRelayLatestVersion() bool`

HasRelayLatestVersion returns a boolean if a field has been set.

### GetRelayLatestReleaseUrl

`func (o *SdkVersionDetailsRep) GetRelayLatestReleaseUrl() string`

GetRelayLatestReleaseUrl returns the RelayLatestReleaseUrl field if non-nil, zero value otherwise.

### GetRelayLatestReleaseUrlOk

`func (o *SdkVersionDetailsRep) GetRelayLatestReleaseUrlOk() (*string, bool)`

GetRelayLatestReleaseUrlOk returns a tuple with the RelayLatestReleaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayLatestReleaseUrl

`func (o *SdkVersionDetailsRep) SetRelayLatestReleaseUrl(v string)`

SetRelayLatestReleaseUrl sets RelayLatestReleaseUrl field to given value.

### HasRelayLatestReleaseUrl

`func (o *SdkVersionDetailsRep) HasRelayLatestReleaseUrl() bool`

HasRelayLatestReleaseUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


