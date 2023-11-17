# InstructionUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of change to make to the removal date for this user from individual targeting for this flag. | 
**FlagKey** | **string** | The flag key | 
**VariationId** | **string** | ID of a variation on the flag | 
**Value** | Pointer to **int32** | The time, in Unix milliseconds, when LaunchDarkly should remove the user from individual targeting for this flag. Required if &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;addExpireUserTargetDate&lt;/code&gt; or &lt;code&gt;updateExpireUserTargetDate&lt;/code&gt;. | [optional] 
**Version** | Pointer to **int32** | The version of the expiring user target to update. Optional and only used if &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;updateExpireUserTargetDate&lt;/code&gt;. If included, update will fail if version doesn&#39;t match current version of the expiring user target. | [optional] 

## Methods

### NewInstructionUserRequest

`func NewInstructionUserRequest(kind string, flagKey string, variationId string, ) *InstructionUserRequest`

NewInstructionUserRequest instantiates a new InstructionUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstructionUserRequestWithDefaults

`func NewInstructionUserRequestWithDefaults() *InstructionUserRequest`

NewInstructionUserRequestWithDefaults instantiates a new InstructionUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *InstructionUserRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InstructionUserRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InstructionUserRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetFlagKey

`func (o *InstructionUserRequest) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *InstructionUserRequest) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *InstructionUserRequest) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetVariationId

`func (o *InstructionUserRequest) GetVariationId() string`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *InstructionUserRequest) GetVariationIdOk() (*string, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *InstructionUserRequest) SetVariationId(v string)`

SetVariationId sets VariationId field to given value.


### GetValue

`func (o *InstructionUserRequest) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InstructionUserRequest) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InstructionUserRequest) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *InstructionUserRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVersion

`func (o *InstructionUserRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstructionUserRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstructionUserRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstructionUserRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


