# PatchSegmentInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of change to make to the user&#39;s removal date from this segment | 
**UserKey** | **string** | A unique key used to represent the user | 
**TargetType** | **string** | The segment&#39;s target type | 
**Value** | Pointer to **int32** | The time, in Unix milliseconds, when the user should be removed from this segment. Required if &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;addExpireUserTargetDate&lt;/code&gt; or &lt;code&gt;updateExpireUserTargetDate&lt;/code&gt;. | [optional] 
**Version** | Pointer to **int32** | The version of the segment to update. Required if &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;updateExpireUserTargetDate&lt;/code&gt;. | [optional] 

## Methods

### NewPatchSegmentInstruction

`func NewPatchSegmentInstruction(kind string, userKey string, targetType string, ) *PatchSegmentInstruction`

NewPatchSegmentInstruction instantiates a new PatchSegmentInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSegmentInstructionWithDefaults

`func NewPatchSegmentInstructionWithDefaults() *PatchSegmentInstruction`

NewPatchSegmentInstructionWithDefaults instantiates a new PatchSegmentInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PatchSegmentInstruction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatchSegmentInstruction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatchSegmentInstruction) SetKind(v string)`

SetKind sets Kind field to given value.


### GetUserKey

`func (o *PatchSegmentInstruction) GetUserKey() string`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *PatchSegmentInstruction) GetUserKeyOk() (*string, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *PatchSegmentInstruction) SetUserKey(v string)`

SetUserKey sets UserKey field to given value.


### GetTargetType

`func (o *PatchSegmentInstruction) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *PatchSegmentInstruction) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *PatchSegmentInstruction) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetValue

`func (o *PatchSegmentInstruction) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchSegmentInstruction) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchSegmentInstruction) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchSegmentInstruction) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVersion

`func (o *PatchSegmentInstruction) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchSegmentInstruction) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchSegmentInstruction) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchSegmentInstruction) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


