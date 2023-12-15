# UserSegmentMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the segment | 
**Key** | **string** | A unique key used to reference the segment | 
**Description** | **string** | A description of the segment&#39;s purpose | 
**Unbounded** | **bool** | Whether this is a Big Segment | 
**External** | **string** | If the segment is an externally-synced Big Segment, the name of the external source&#39; | 
**IsMember** | **bool** | Whether the user is a member of this segment, either by explicit inclusion or by rule matching | 
**IsUserTargeted** | **bool** | Whether the user is explicitly included in this segment | 
**IsRuleTargeted** | **bool** | Whether the user is captured by this segment&#39;s rules. The value of this field is undefined if the user is also explicitly included (isUserTargeted &#x3D; true). | 
**Links** | [**map[string]Link**](Link.md) |  | 

## Methods

### NewUserSegmentMembership

`func NewUserSegmentMembership(name string, key string, description string, unbounded bool, external string, isMember bool, isUserTargeted bool, isRuleTargeted bool, links map[string]Link, ) *UserSegmentMembership`

NewUserSegmentMembership instantiates a new UserSegmentMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSegmentMembershipWithDefaults

`func NewUserSegmentMembershipWithDefaults() *UserSegmentMembership`

NewUserSegmentMembershipWithDefaults instantiates a new UserSegmentMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserSegmentMembership) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSegmentMembership) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSegmentMembership) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *UserSegmentMembership) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserSegmentMembership) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserSegmentMembership) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *UserSegmentMembership) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserSegmentMembership) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserSegmentMembership) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUnbounded

`func (o *UserSegmentMembership) GetUnbounded() bool`

GetUnbounded returns the Unbounded field if non-nil, zero value otherwise.

### GetUnboundedOk

`func (o *UserSegmentMembership) GetUnboundedOk() (*bool, bool)`

GetUnboundedOk returns a tuple with the Unbounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbounded

`func (o *UserSegmentMembership) SetUnbounded(v bool)`

SetUnbounded sets Unbounded field to given value.


### GetExternal

`func (o *UserSegmentMembership) GetExternal() string`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *UserSegmentMembership) GetExternalOk() (*string, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *UserSegmentMembership) SetExternal(v string)`

SetExternal sets External field to given value.


### GetIsMember

`func (o *UserSegmentMembership) GetIsMember() bool`

GetIsMember returns the IsMember field if non-nil, zero value otherwise.

### GetIsMemberOk

`func (o *UserSegmentMembership) GetIsMemberOk() (*bool, bool)`

GetIsMemberOk returns a tuple with the IsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMember

`func (o *UserSegmentMembership) SetIsMember(v bool)`

SetIsMember sets IsMember field to given value.


### GetIsUserTargeted

`func (o *UserSegmentMembership) GetIsUserTargeted() bool`

GetIsUserTargeted returns the IsUserTargeted field if non-nil, zero value otherwise.

### GetIsUserTargetedOk

`func (o *UserSegmentMembership) GetIsUserTargetedOk() (*bool, bool)`

GetIsUserTargetedOk returns a tuple with the IsUserTargeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserTargeted

`func (o *UserSegmentMembership) SetIsUserTargeted(v bool)`

SetIsUserTargeted sets IsUserTargeted field to given value.


### GetIsRuleTargeted

`func (o *UserSegmentMembership) GetIsRuleTargeted() bool`

GetIsRuleTargeted returns the IsRuleTargeted field if non-nil, zero value otherwise.

### GetIsRuleTargetedOk

`func (o *UserSegmentMembership) GetIsRuleTargetedOk() (*bool, bool)`

GetIsRuleTargetedOk returns a tuple with the IsRuleTargeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRuleTargeted

`func (o *UserSegmentMembership) SetIsRuleTargeted(v bool)`

SetIsRuleTargeted sets IsRuleTargeted field to given value.


### GetLinks

`func (o *UserSegmentMembership) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSegmentMembership) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSegmentMembership) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


