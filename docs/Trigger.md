# Trigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to **string** |  | [optional] 
**DefaultEventName** | Pointer to **string** |  | [optional] 
**Documentation** | Pointer to **string** |  | [optional] 
**Parser** | Pointer to [**TriggerParser**](TriggerParser.md) |  | [optional] 
**TestEventNameRegexp** | Pointer to **string** |  | [optional] 

## Methods

### NewTrigger

`func NewTrigger() *Trigger`

NewTrigger instantiates a new Trigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWithDefaults

`func NewTriggerWithDefaults() *Trigger`

NewTriggerWithDefaults instantiates a new Trigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *Trigger) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Trigger) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Trigger) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *Trigger) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetDefaultEventName

`func (o *Trigger) GetDefaultEventName() string`

GetDefaultEventName returns the DefaultEventName field if non-nil, zero value otherwise.

### GetDefaultEventNameOk

`func (o *Trigger) GetDefaultEventNameOk() (*string, bool)`

GetDefaultEventNameOk returns a tuple with the DefaultEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEventName

`func (o *Trigger) SetDefaultEventName(v string)`

SetDefaultEventName sets DefaultEventName field to given value.

### HasDefaultEventName

`func (o *Trigger) HasDefaultEventName() bool`

HasDefaultEventName returns a boolean if a field has been set.

### GetDocumentation

`func (o *Trigger) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *Trigger) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *Trigger) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *Trigger) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetParser

`func (o *Trigger) GetParser() TriggerParser`

GetParser returns the Parser field if non-nil, zero value otherwise.

### GetParserOk

`func (o *Trigger) GetParserOk() (*TriggerParser, bool)`

GetParserOk returns a tuple with the Parser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParser

`func (o *Trigger) SetParser(v TriggerParser)`

SetParser sets Parser field to given value.

### HasParser

`func (o *Trigger) HasParser() bool`

HasParser returns a boolean if a field has been set.

### GetTestEventNameRegexp

`func (o *Trigger) GetTestEventNameRegexp() string`

GetTestEventNameRegexp returns the TestEventNameRegexp field if non-nil, zero value otherwise.

### GetTestEventNameRegexpOk

`func (o *Trigger) GetTestEventNameRegexpOk() (*string, bool)`

GetTestEventNameRegexpOk returns a tuple with the TestEventNameRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestEventNameRegexp

`func (o *Trigger) SetTestEventNameRegexp(v string)`

SetTestEventNameRegexp sets TestEventNameRegexp field to given value.

### HasTestEventNameRegexp

`func (o *Trigger) HasTestEventNameRegexp() bool`

HasTestEventNameRegexp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


