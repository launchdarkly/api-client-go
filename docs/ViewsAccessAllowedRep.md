# ViewsAccessAllowedRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Reason** | [**ViewsAccessAllowedReason**](ViewsAccessAllowedReason.md) |  | 

## Methods

### NewViewsAccessAllowedRep

`func NewViewsAccessAllowedRep(action string, reason ViewsAccessAllowedReason, ) *ViewsAccessAllowedRep`

NewViewsAccessAllowedRep instantiates a new ViewsAccessAllowedRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsAccessAllowedRepWithDefaults

`func NewViewsAccessAllowedRepWithDefaults() *ViewsAccessAllowedRep`

NewViewsAccessAllowedRepWithDefaults instantiates a new ViewsAccessAllowedRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ViewsAccessAllowedRep) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ViewsAccessAllowedRep) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ViewsAccessAllowedRep) SetAction(v string)`

SetAction sets Action field to given value.


### GetReason

`func (o *ViewsAccessAllowedRep) GetReason() ViewsAccessAllowedReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ViewsAccessAllowedRep) GetReasonOk() (*ViewsAccessAllowedReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ViewsAccessAllowedRep) SetReason(v ViewsAccessAllowedReason)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


