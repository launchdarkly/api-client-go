# ViewLinkRequestKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | **[]string** | Keys of the resources (flags, segments) to link/unlink | 
**Filter** | Pointer to **string** | Optional filter string to determine which resources should be linked. Resources only need to match either the filter or explicitly-listed keys to be linked (union). Uses the same syntax as list endpoints: flags use comma-separated field:value filters, segments use queryfilter syntax.  Supported filters by resource type: - flags: query, tags, maintainerId, maintainerTeamKey, type, status, state, staleState, sdkAvailability, targeting, hasExperiment, hasDataExport, evaluated, creationDate, contextKindTargeted, contextKindsEvaluated, filterEnv, segmentTargeted, codeReferences.min, codeReferences.max, excludeSettings, releasePipeline, applicationEvaluated, purpose, guardedRollout, view, key, name, archived, followerId - segments (queryfilter): query, tags, keys, excludedKeys, unbounded, external, view, type Some filters are only available when the corresponding feature is enabled on your account.  | [optional] 
**Comment** | Pointer to **string** | Optional comment for the link/unlink operation | [optional] [default to ""]

## Methods

### NewViewLinkRequestKeys

`func NewViewLinkRequestKeys(keys []string, ) *ViewLinkRequestKeys`

NewViewLinkRequestKeys instantiates a new ViewLinkRequestKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewLinkRequestKeysWithDefaults

`func NewViewLinkRequestKeysWithDefaults() *ViewLinkRequestKeys`

NewViewLinkRequestKeysWithDefaults instantiates a new ViewLinkRequestKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ViewLinkRequestKeys) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ViewLinkRequestKeys) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ViewLinkRequestKeys) SetKeys(v []string)`

SetKeys sets Keys field to given value.


### GetFilter

`func (o *ViewLinkRequestKeys) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ViewLinkRequestKeys) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ViewLinkRequestKeys) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ViewLinkRequestKeys) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetComment

`func (o *ViewLinkRequestKeys) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ViewLinkRequestKeys) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ViewLinkRequestKeys) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ViewLinkRequestKeys) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


