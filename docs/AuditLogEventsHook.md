# AuditLogEventsHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPolicy** | Pointer to [**[]Policy**](Policy.md) |  | [optional] 
**DeliveryMethod** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**IncludeErrorResponseBody** | Pointer to **bool** |  | [optional] 
**Templates** | Pointer to [**WebhookBodyTemplate**](WebhookBodyTemplate.md) |  | [optional] 
**UseStandardWebhookPayload** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuditLogEventsHook

`func NewAuditLogEventsHook() *AuditLogEventsHook`

NewAuditLogEventsHook instantiates a new AuditLogEventsHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEventsHookWithDefaults

`func NewAuditLogEventsHookWithDefaults() *AuditLogEventsHook`

NewAuditLogEventsHookWithDefaults instantiates a new AuditLogEventsHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPolicy

`func (o *AuditLogEventsHook) GetDefaultPolicy() []Policy`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *AuditLogEventsHook) GetDefaultPolicyOk() (*[]Policy, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *AuditLogEventsHook) SetDefaultPolicy(v []Policy)`

SetDefaultPolicy sets DefaultPolicy field to given value.

### HasDefaultPolicy

`func (o *AuditLogEventsHook) HasDefaultPolicy() bool`

HasDefaultPolicy returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *AuditLogEventsHook) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *AuditLogEventsHook) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *AuditLogEventsHook) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *AuditLogEventsHook) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetEndpoint

`func (o *AuditLogEventsHook) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AuditLogEventsHook) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AuditLogEventsHook) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AuditLogEventsHook) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetIncludeErrorResponseBody

`func (o *AuditLogEventsHook) GetIncludeErrorResponseBody() bool`

GetIncludeErrorResponseBody returns the IncludeErrorResponseBody field if non-nil, zero value otherwise.

### GetIncludeErrorResponseBodyOk

`func (o *AuditLogEventsHook) GetIncludeErrorResponseBodyOk() (*bool, bool)`

GetIncludeErrorResponseBodyOk returns a tuple with the IncludeErrorResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeErrorResponseBody

`func (o *AuditLogEventsHook) SetIncludeErrorResponseBody(v bool)`

SetIncludeErrorResponseBody sets IncludeErrorResponseBody field to given value.

### HasIncludeErrorResponseBody

`func (o *AuditLogEventsHook) HasIncludeErrorResponseBody() bool`

HasIncludeErrorResponseBody returns a boolean if a field has been set.

### GetTemplates

`func (o *AuditLogEventsHook) GetTemplates() WebhookBodyTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *AuditLogEventsHook) GetTemplatesOk() (*WebhookBodyTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *AuditLogEventsHook) SetTemplates(v WebhookBodyTemplate)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *AuditLogEventsHook) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetUseStandardWebhookPayload

`func (o *AuditLogEventsHook) GetUseStandardWebhookPayload() bool`

GetUseStandardWebhookPayload returns the UseStandardWebhookPayload field if non-nil, zero value otherwise.

### GetUseStandardWebhookPayloadOk

`func (o *AuditLogEventsHook) GetUseStandardWebhookPayloadOk() (*bool, bool)`

GetUseStandardWebhookPayloadOk returns a tuple with the UseStandardWebhookPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStandardWebhookPayload

`func (o *AuditLogEventsHook) SetUseStandardWebhookPayload(v bool)`

SetUseStandardWebhookPayload sets UseStandardWebhookPayload field to given value.

### HasUseStandardWebhookPayload

`func (o *AuditLogEventsHook) HasUseStandardWebhookPayload() bool`

HasUseStandardWebhookPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


