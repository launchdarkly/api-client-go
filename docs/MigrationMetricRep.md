# MigrationMetricRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Series** | Pointer to **[]map[string]float32** |  | [optional] 
**Metadata** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewMigrationMetricRep

`func NewMigrationMetricRep() *MigrationMetricRep`

NewMigrationMetricRep instantiates a new MigrationMetricRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationMetricRepWithDefaults

`func NewMigrationMetricRepWithDefaults() *MigrationMetricRep`

NewMigrationMetricRepWithDefaults instantiates a new MigrationMetricRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MigrationMetricRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MigrationMetricRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MigrationMetricRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MigrationMetricRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSeries

`func (o *MigrationMetricRep) GetSeries() []map[string]float32`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *MigrationMetricRep) GetSeriesOk() (*[]map[string]float32, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *MigrationMetricRep) SetSeries(v []map[string]float32)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *MigrationMetricRep) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetMetadata

`func (o *MigrationMetricRep) GetMetadata() []map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MigrationMetricRep) GetMetadataOk() (*[]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MigrationMetricRep) SetMetadata(v []map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MigrationMetricRep) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


