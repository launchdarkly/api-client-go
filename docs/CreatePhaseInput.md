# CreatePhaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | [**[]AudiencePost**](AudiencePost.md) | An ordered list of the audiences for this release phase. Each audience corresponds to a LaunchDarkly environment. | 
**Name** | **string** | The release phase name | 
**Configuration** | Pointer to [**PhaseConfiguration**](PhaseConfiguration.md) |  | [optional] 

## Methods

### NewCreatePhaseInput

`func NewCreatePhaseInput(audiences []AudiencePost, name string, ) *CreatePhaseInput`

NewCreatePhaseInput instantiates a new CreatePhaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePhaseInputWithDefaults

`func NewCreatePhaseInputWithDefaults() *CreatePhaseInput`

NewCreatePhaseInputWithDefaults instantiates a new CreatePhaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *CreatePhaseInput) GetAudiences() []AudiencePost`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *CreatePhaseInput) GetAudiencesOk() (*[]AudiencePost, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *CreatePhaseInput) SetAudiences(v []AudiencePost)`

SetAudiences sets Audiences field to given value.


### GetName

`func (o *CreatePhaseInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePhaseInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePhaseInput) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *CreatePhaseInput) GetConfiguration() PhaseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreatePhaseInput) GetConfigurationOk() (*PhaseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreatePhaseInput) SetConfiguration(v PhaseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreatePhaseInput) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


