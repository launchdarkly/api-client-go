# SlicedResultsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **string** | An attribute that results are sliced by | [optional] 
**AttributeValue** | Pointer to **string** | Attribute Value for &#39;attribute&#39; | [optional] 
**TreatmentResults** | Pointer to [**[]TreatmentResultRep**](TreatmentResultRep.md) | A list of the results for each treatment | [optional] 

## Methods

### NewSlicedResultsRep

`func NewSlicedResultsRep() *SlicedResultsRep`

NewSlicedResultsRep instantiates a new SlicedResultsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlicedResultsRepWithDefaults

`func NewSlicedResultsRepWithDefaults() *SlicedResultsRep`

NewSlicedResultsRepWithDefaults instantiates a new SlicedResultsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *SlicedResultsRep) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *SlicedResultsRep) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *SlicedResultsRep) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *SlicedResultsRep) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetAttributeValue

`func (o *SlicedResultsRep) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *SlicedResultsRep) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *SlicedResultsRep) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *SlicedResultsRep) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.

### GetTreatmentResults

`func (o *SlicedResultsRep) GetTreatmentResults() []TreatmentResultRep`

GetTreatmentResults returns the TreatmentResults field if non-nil, zero value otherwise.

### GetTreatmentResultsOk

`func (o *SlicedResultsRep) GetTreatmentResultsOk() (*[]TreatmentResultRep, bool)`

GetTreatmentResultsOk returns a tuple with the TreatmentResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatmentResults

`func (o *SlicedResultsRep) SetTreatmentResults(v []TreatmentResultRep)`

SetTreatmentResults sets TreatmentResults field to given value.

### HasTreatmentResults

`func (o *SlicedResultsRep) HasTreatmentResults() bool`

HasTreatmentResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


