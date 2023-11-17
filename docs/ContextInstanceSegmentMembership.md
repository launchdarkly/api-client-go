# ContextInstanceSegmentMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the segment | 
**Key** | **string** | A unique key used to reference the segment | 
**Description** | **string** | A description of the segment&#39;s purpose | 
**Unbounded** | **bool** | Whether this is an unbounded segment. Unbounded segments, also called Big Segments, may be list-based segments with more than 15,000 entries, or synced segments. | 
**External** | **string** | If the segment is a synced segment, the name of the external source | 
**IsMember** | **bool** | Whether the context is a member of this segment, either by explicit inclusion or by rule matching | 
**IsIndividuallyTargeted** | **bool** | Whether the context is explicitly included in this segment | 
**IsRuleTargeted** | **bool** | Whether the context is captured by this segment&#39;s rules. The value of this field is undefined if the context is also explicitly included (&lt;code&gt;isIndividuallyTargeted&lt;/code&gt; is &lt;code&gt;true&lt;/code&gt;). | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewContextInstanceSegmentMembership

`func NewContextInstanceSegmentMembership(name string, key string, description string, unbounded bool, external string, isMember bool, isIndividuallyTargeted bool, isRuleTargeted bool, links map[string]Link, ) *ContextInstanceSegmentMembership`

NewContextInstanceSegmentMembership instantiates a new ContextInstanceSegmentMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextInstanceSegmentMembershipWithDefaults

`func NewContextInstanceSegmentMembershipWithDefaults() *ContextInstanceSegmentMembership`

NewContextInstanceSegmentMembershipWithDefaults instantiates a new ContextInstanceSegmentMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContextInstanceSegmentMembership) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextInstanceSegmentMembership) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextInstanceSegmentMembership) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ContextInstanceSegmentMembership) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContextInstanceSegmentMembership) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContextInstanceSegmentMembership) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *ContextInstanceSegmentMembership) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContextInstanceSegmentMembership) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContextInstanceSegmentMembership) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUnbounded

`func (o *ContextInstanceSegmentMembership) GetUnbounded() bool`

GetUnbounded returns the Unbounded field if non-nil, zero value otherwise.

### GetUnboundedOk

`func (o *ContextInstanceSegmentMembership) GetUnboundedOk() (*bool, bool)`

GetUnboundedOk returns a tuple with the Unbounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbounded

`func (o *ContextInstanceSegmentMembership) SetUnbounded(v bool)`

SetUnbounded sets Unbounded field to given value.


### GetExternal

`func (o *ContextInstanceSegmentMembership) GetExternal() string`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *ContextInstanceSegmentMembership) GetExternalOk() (*string, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *ContextInstanceSegmentMembership) SetExternal(v string)`

SetExternal sets External field to given value.


### GetIsMember

`func (o *ContextInstanceSegmentMembership) GetIsMember() bool`

GetIsMember returns the IsMember field if non-nil, zero value otherwise.

### GetIsMemberOk

`func (o *ContextInstanceSegmentMembership) GetIsMemberOk() (*bool, bool)`

GetIsMemberOk returns a tuple with the IsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMember

`func (o *ContextInstanceSegmentMembership) SetIsMember(v bool)`

SetIsMember sets IsMember field to given value.


### GetIsIndividuallyTargeted

`func (o *ContextInstanceSegmentMembership) GetIsIndividuallyTargeted() bool`

GetIsIndividuallyTargeted returns the IsIndividuallyTargeted field if non-nil, zero value otherwise.

### GetIsIndividuallyTargetedOk

`func (o *ContextInstanceSegmentMembership) GetIsIndividuallyTargetedOk() (*bool, bool)`

GetIsIndividuallyTargetedOk returns a tuple with the IsIndividuallyTargeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndividuallyTargeted

`func (o *ContextInstanceSegmentMembership) SetIsIndividuallyTargeted(v bool)`

SetIsIndividuallyTargeted sets IsIndividuallyTargeted field to given value.


### GetIsRuleTargeted

`func (o *ContextInstanceSegmentMembership) GetIsRuleTargeted() bool`

GetIsRuleTargeted returns the IsRuleTargeted field if non-nil, zero value otherwise.

### GetIsRuleTargetedOk

`func (o *ContextInstanceSegmentMembership) GetIsRuleTargetedOk() (*bool, bool)`

GetIsRuleTargetedOk returns a tuple with the IsRuleTargeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRuleTargeted

`func (o *ContextInstanceSegmentMembership) SetIsRuleTargeted(v bool)`

SetIsRuleTargeted sets IsRuleTargeted field to given value.


### GetLinks

`func (o *ContextInstanceSegmentMembership) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContextInstanceSegmentMembership) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContextInstanceSegmentMembership) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


