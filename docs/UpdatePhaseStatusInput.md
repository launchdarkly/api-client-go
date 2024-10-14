# UpdatePhaseStatusInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Audiences** | Pointer to [**[]ReleaserAudienceConfigInput**](ReleaserAudienceConfigInput.md) | Extra configuration for audiences required upon phase initialization. | [optional] 

## Methods

### NewUpdatePhaseStatusInput

`func NewUpdatePhaseStatusInput() *UpdatePhaseStatusInput`

NewUpdatePhaseStatusInput instantiates a new UpdatePhaseStatusInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePhaseStatusInputWithDefaults

`func NewUpdatePhaseStatusInputWithDefaults() *UpdatePhaseStatusInput`

NewUpdatePhaseStatusInputWithDefaults instantiates a new UpdatePhaseStatusInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdatePhaseStatusInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdatePhaseStatusInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdatePhaseStatusInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdatePhaseStatusInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAudiences

`func (o *UpdatePhaseStatusInput) GetAudiences() []ReleaserAudienceConfigInput`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *UpdatePhaseStatusInput) GetAudiencesOk() (*[]ReleaserAudienceConfigInput, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *UpdatePhaseStatusInput) SetAudiences(v []ReleaserAudienceConfigInput)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *UpdatePhaseStatusInput) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


