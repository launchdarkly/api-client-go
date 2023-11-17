# MemberPermissionGrantSummaryRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionSet** | Pointer to **string** | The name of the group of related actions to allow. A permission grant may have either an &lt;code&gt;actionSet&lt;/code&gt; or a list of &lt;code&gt;actions&lt;/code&gt; but not both at the same time. | [optional] 
**Actions** | Pointer to **[]string** | A list of actions to allow. A permission grant may have either an &lt;code&gt;actionSet&lt;/code&gt; or a list of &lt;code&gt;actions&lt;/code&gt; but not both at the same time. | [optional] 
**Resource** | **string** | The resource for which the actions are allowed | 

## Methods

### NewMemberPermissionGrantSummaryRep

`func NewMemberPermissionGrantSummaryRep(resource string, ) *MemberPermissionGrantSummaryRep`

NewMemberPermissionGrantSummaryRep instantiates a new MemberPermissionGrantSummaryRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberPermissionGrantSummaryRepWithDefaults

`func NewMemberPermissionGrantSummaryRepWithDefaults() *MemberPermissionGrantSummaryRep`

NewMemberPermissionGrantSummaryRepWithDefaults instantiates a new MemberPermissionGrantSummaryRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionSet

`func (o *MemberPermissionGrantSummaryRep) GetActionSet() string`

GetActionSet returns the ActionSet field if non-nil, zero value otherwise.

### GetActionSetOk

`func (o *MemberPermissionGrantSummaryRep) GetActionSetOk() (*string, bool)`

GetActionSetOk returns a tuple with the ActionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSet

`func (o *MemberPermissionGrantSummaryRep) SetActionSet(v string)`

SetActionSet sets ActionSet field to given value.

### HasActionSet

`func (o *MemberPermissionGrantSummaryRep) HasActionSet() bool`

HasActionSet returns a boolean if a field has been set.

### GetActions

`func (o *MemberPermissionGrantSummaryRep) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MemberPermissionGrantSummaryRep) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MemberPermissionGrantSummaryRep) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MemberPermissionGrantSummaryRep) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetResource

`func (o *MemberPermissionGrantSummaryRep) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MemberPermissionGrantSummaryRep) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MemberPermissionGrantSummaryRep) SetResource(v string)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


