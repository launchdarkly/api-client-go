# Distribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The type of distribution. | [optional] 
**Parameters** | Pointer to **map[string]interface{}** | The parameters of the distribution. The parameters are different for each distribution type. When &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;normal&lt;/code&gt;, the parameters of the distribution are &#39;mu&#39; and &#39;sigma&#39;. When &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;beta&lt;/code&gt;, the parameters of the distribution are &#39;alpha&#39; and &#39;beta.&#39; | [optional] 

## Methods

### NewDistribution

`func NewDistribution() *Distribution`

NewDistribution instantiates a new Distribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionWithDefaults

`func NewDistributionWithDefaults() *Distribution`

NewDistributionWithDefaults instantiates a new Distribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *Distribution) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Distribution) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Distribution) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Distribution) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetParameters

`func (o *Distribution) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Distribution) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Distribution) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Distribution) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


