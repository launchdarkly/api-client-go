# TriggerPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional comment describing the trigger | [optional] 
**Instructions** | Pointer to **[]map[string]interface{}** | The action to perform when triggering. This should be an array with a single object that looks like &lt;code&gt;{\&quot;kind\&quot;: \&quot;flag_action\&quot;}&lt;/code&gt;. Supported flag actions are &lt;code&gt;turnFlagOn&lt;/code&gt; and &lt;code&gt;turnFlagOff&lt;/code&gt;. | [optional] 
**IntegrationKey** | **string** | The unique identifier of the integration for your trigger. Use &lt;code&gt;generic-trigger&lt;/code&gt; for integrations not explicitly supported. | 

## Methods

### NewTriggerPost

`func NewTriggerPost(integrationKey string, ) *TriggerPost`

NewTriggerPost instantiates a new TriggerPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerPostWithDefaults

`func NewTriggerPostWithDefaults() *TriggerPost`

NewTriggerPostWithDefaults instantiates a new TriggerPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *TriggerPost) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TriggerPost) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TriggerPost) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TriggerPost) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInstructions

`func (o *TriggerPost) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *TriggerPost) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *TriggerPost) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *TriggerPost) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetIntegrationKey

`func (o *TriggerPost) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *TriggerPost) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *TriggerPost) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


