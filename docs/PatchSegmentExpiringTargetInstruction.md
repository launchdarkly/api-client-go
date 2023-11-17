# PatchSegmentExpiringTargetInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of change to make to the context&#39;s removal date from this segment | 
**ContextKey** | **string** | A unique key used to represent the context | 
**ContextKind** | **string** | The kind of context | 
**TargetType** | **string** | The segment&#39;s target type | 
**Value** | Pointer to **int32** | The time, in Unix milliseconds, when the context should be removed from this segment. Required if &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;addExpiringTarget&lt;/code&gt; or &lt;code&gt;updateExpiringTarget&lt;/code&gt;. | [optional] 
**Version** | Pointer to **int32** | The version of the expiring target to update. Optional and only used if &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;updateExpiringTarget&lt;/code&gt;. If included, update will fail if version doesn&#39;t match current version of the expiring target. | [optional] 

## Methods

### NewPatchSegmentExpiringTargetInstruction

`func NewPatchSegmentExpiringTargetInstruction(kind string, contextKey string, contextKind string, targetType string, ) *PatchSegmentExpiringTargetInstruction`

NewPatchSegmentExpiringTargetInstruction instantiates a new PatchSegmentExpiringTargetInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSegmentExpiringTargetInstructionWithDefaults

`func NewPatchSegmentExpiringTargetInstructionWithDefaults() *PatchSegmentExpiringTargetInstruction`

NewPatchSegmentExpiringTargetInstructionWithDefaults instantiates a new PatchSegmentExpiringTargetInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PatchSegmentExpiringTargetInstruction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatchSegmentExpiringTargetInstruction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatchSegmentExpiringTargetInstruction) SetKind(v string)`

SetKind sets Kind field to given value.


### GetContextKey

`func (o *PatchSegmentExpiringTargetInstruction) GetContextKey() string`

GetContextKey returns the ContextKey field if non-nil, zero value otherwise.

### GetContextKeyOk

`func (o *PatchSegmentExpiringTargetInstruction) GetContextKeyOk() (*string, bool)`

GetContextKeyOk returns a tuple with the ContextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKey

`func (o *PatchSegmentExpiringTargetInstruction) SetContextKey(v string)`

SetContextKey sets ContextKey field to given value.


### GetContextKind

`func (o *PatchSegmentExpiringTargetInstruction) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *PatchSegmentExpiringTargetInstruction) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *PatchSegmentExpiringTargetInstruction) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.


### GetTargetType

`func (o *PatchSegmentExpiringTargetInstruction) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *PatchSegmentExpiringTargetInstruction) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *PatchSegmentExpiringTargetInstruction) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetValue

`func (o *PatchSegmentExpiringTargetInstruction) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchSegmentExpiringTargetInstruction) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchSegmentExpiringTargetInstruction) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchSegmentExpiringTargetInstruction) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVersion

`func (o *PatchSegmentExpiringTargetInstruction) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchSegmentExpiringTargetInstruction) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchSegmentExpiringTargetInstruction) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchSegmentExpiringTargetInstruction) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


