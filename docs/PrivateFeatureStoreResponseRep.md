# PrivateFeatureStoreResponseRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** |  | 
**ResponseBody** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewPrivateFeatureStoreResponseRep

`func NewPrivateFeatureStoreResponseRep(statusCode int32, ) *PrivateFeatureStoreResponseRep`

NewPrivateFeatureStoreResponseRep instantiates a new PrivateFeatureStoreResponseRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateFeatureStoreResponseRepWithDefaults

`func NewPrivateFeatureStoreResponseRepWithDefaults() *PrivateFeatureStoreResponseRep`

NewPrivateFeatureStoreResponseRepWithDefaults instantiates a new PrivateFeatureStoreResponseRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *PrivateFeatureStoreResponseRep) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *PrivateFeatureStoreResponseRep) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *PrivateFeatureStoreResponseRep) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetResponseBody

`func (o *PrivateFeatureStoreResponseRep) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *PrivateFeatureStoreResponseRep) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *PrivateFeatureStoreResponseRep) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *PrivateFeatureStoreResponseRep) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetTimestamp

`func (o *PrivateFeatureStoreResponseRep) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PrivateFeatureStoreResponseRep) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PrivateFeatureStoreResponseRep) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PrivateFeatureStoreResponseRep) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


