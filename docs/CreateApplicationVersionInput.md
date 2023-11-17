# CreateApplicationVersionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The unique identifier of this application version | 
**Name** | **string** | The name of this version | 
**Supported** | Pointer to **bool** | Whether this version is supported. Only applicable if the application &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;mobile&lt;/code&gt;. Defaults to &lt;code&gt;true&lt;/code&gt;. | [optional] 

## Methods

### NewCreateApplicationVersionInput

`func NewCreateApplicationVersionInput(key string, name string, ) *CreateApplicationVersionInput`

NewCreateApplicationVersionInput instantiates a new CreateApplicationVersionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationVersionInputWithDefaults

`func NewCreateApplicationVersionInputWithDefaults() *CreateApplicationVersionInput`

NewCreateApplicationVersionInputWithDefaults instantiates a new CreateApplicationVersionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateApplicationVersionInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateApplicationVersionInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateApplicationVersionInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateApplicationVersionInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApplicationVersionInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApplicationVersionInput) SetName(v string)`

SetName sets Name field to given value.


### GetSupported

`func (o *CreateApplicationVersionInput) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *CreateApplicationVersionInput) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *CreateApplicationVersionInput) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *CreateApplicationVersionInput) HasSupported() bool`

HasSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


