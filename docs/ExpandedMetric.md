# ExpandedMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | A unique key used to reference the metric | [optional] 
**Name** | Pointer to **string** | A human-friendly name for the metric | [optional] 
**CreationDate** | Pointer to **int64** | Creation date in milliseconds | [optional] 
**LastModified** | Pointer to **int64** | Last modification date in milliseconds | [optional] 
**IsActive** | Pointer to **bool** | Whether the metric is active | [optional] 
**EventKey** | Pointer to **string** | Event key for the metric | [optional] 
**Id** | Pointer to **string** | ID of the metric | [optional] 
**VersionId** | Pointer to **string** | Version ID of the metric | [optional] 
**Kind** | Pointer to **string** | Kind of the Metric | [optional] 
**Category** | Pointer to **string** | Category of the Metric | [optional] 
**Description** | Pointer to **string** | Description of the Metric | [optional] 
**IsNumeric** | Pointer to **bool** |  | [optional] 
**LastSeen** | Pointer to **int64** | Last seen date in milliseconds | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 

## Methods

### NewExpandedMetric

`func NewExpandedMetric() *ExpandedMetric`

NewExpandedMetric instantiates a new ExpandedMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedMetricWithDefaults

`func NewExpandedMetricWithDefaults() *ExpandedMetric`

NewExpandedMetricWithDefaults instantiates a new ExpandedMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ExpandedMetric) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExpandedMetric) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExpandedMetric) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ExpandedMetric) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ExpandedMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpandedMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpandedMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExpandedMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationDate

`func (o *ExpandedMetric) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ExpandedMetric) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ExpandedMetric) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ExpandedMetric) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastModified

`func (o *ExpandedMetric) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ExpandedMetric) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ExpandedMetric) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ExpandedMetric) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetIsActive

`func (o *ExpandedMetric) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ExpandedMetric) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ExpandedMetric) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ExpandedMetric) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEventKey

`func (o *ExpandedMetric) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *ExpandedMetric) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *ExpandedMetric) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *ExpandedMetric) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetId

`func (o *ExpandedMetric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpandedMetric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpandedMetric) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExpandedMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionId

`func (o *ExpandedMetric) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ExpandedMetric) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ExpandedMetric) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ExpandedMetric) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetKind

`func (o *ExpandedMetric) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExpandedMetric) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExpandedMetric) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExpandedMetric) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCategory

`func (o *ExpandedMetric) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ExpandedMetric) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ExpandedMetric) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ExpandedMetric) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ExpandedMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpandedMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpandedMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpandedMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsNumeric

`func (o *ExpandedMetric) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *ExpandedMetric) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *ExpandedMetric) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *ExpandedMetric) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetLastSeen

`func (o *ExpandedMetric) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ExpandedMetric) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ExpandedMetric) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ExpandedMetric) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLinks

`func (o *ExpandedMetric) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpandedMetric) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpandedMetric) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExpandedMetric) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


