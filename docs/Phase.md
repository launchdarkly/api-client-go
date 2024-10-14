# Phase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The phase ID | 
**Audiences** | [**[]Audience**](Audience.md) |  | 
**Name** | **string** | The release phase name | 
**Configuration** | Pointer to [**PhaseConfiguration**](PhaseConfiguration.md) |  | [optional] 

## Methods

### NewPhase

`func NewPhase(id string, audiences []Audience, name string, ) *Phase`

NewPhase instantiates a new Phase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhaseWithDefaults

`func NewPhaseWithDefaults() *Phase`

NewPhaseWithDefaults instantiates a new Phase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Phase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Phase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Phase) SetId(v string)`

SetId sets Id field to given value.


### GetAudiences

`func (o *Phase) GetAudiences() []Audience`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *Phase) GetAudiencesOk() (*[]Audience, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *Phase) SetAudiences(v []Audience)`

SetAudiences sets Audiences field to given value.


### GetName

`func (o *Phase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Phase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Phase) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *Phase) GetConfiguration() PhaseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Phase) GetConfigurationOk() (*PhaseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Phase) SetConfiguration(v PhaseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Phase) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


