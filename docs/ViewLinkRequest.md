# ViewLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | **[]string** | Keys of the resources (flags, segments) to link/unlink | 
**Filter** | **string** | Filter string to match resources for linking. Uses the same syntax as list endpoints: flags use comma-separated field:value filters, segments use queryfilter syntax.  Supported filters by resource type: - flags: query, tags, maintainerId, maintainerTeamKey, type, status, state, staleState, sdkAvailability, targeting, hasExperiment, hasDataExport, evaluated, creationDate, contextKindTargeted, contextKindsEvaluated, filterEnv, segmentTargeted, codeReferences.min, codeReferences.max, excludeSettings, releasePipeline, applicationEvaluated, purpose, guardedRollout, view, key, name, archived, followerId - segments (queryfilter): query, tags, keys, excludedKeys, unbounded, external, view, type Some filters are only available when the corresponding feature is enabled on your account.  | 
**Comment** | Pointer to **string** | Optional comment for the link/unlink operation | [optional] [default to ""]
**SegmentIdentifiers** | [**[]ViewLinkRequestSegmentIdentifier**](ViewLinkRequestSegmentIdentifier.md) | Identifiers of the segments to link/unlink (environmentId and segmentKey) | 
**EnvironmentId** | Pointer to **string** | Required when using filter for segment resources. Specifies which environment to query for segments matching the filter. Ignored for flag resources (flags are global across environments).  | [optional] 

## Methods

### NewViewLinkRequest

`func NewViewLinkRequest(keys []string, filter string, segmentIdentifiers []ViewLinkRequestSegmentIdentifier, ) *ViewLinkRequest`

NewViewLinkRequest instantiates a new ViewLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewLinkRequestWithDefaults

`func NewViewLinkRequestWithDefaults() *ViewLinkRequest`

NewViewLinkRequestWithDefaults instantiates a new ViewLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ViewLinkRequest) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ViewLinkRequest) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ViewLinkRequest) SetKeys(v []string)`

SetKeys sets Keys field to given value.


### GetFilter

`func (o *ViewLinkRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ViewLinkRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ViewLinkRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetComment

`func (o *ViewLinkRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ViewLinkRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ViewLinkRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ViewLinkRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSegmentIdentifiers

`func (o *ViewLinkRequest) GetSegmentIdentifiers() []ViewLinkRequestSegmentIdentifier`

GetSegmentIdentifiers returns the SegmentIdentifiers field if non-nil, zero value otherwise.

### GetSegmentIdentifiersOk

`func (o *ViewLinkRequest) GetSegmentIdentifiersOk() (*[]ViewLinkRequestSegmentIdentifier, bool)`

GetSegmentIdentifiersOk returns a tuple with the SegmentIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIdentifiers

`func (o *ViewLinkRequest) SetSegmentIdentifiers(v []ViewLinkRequestSegmentIdentifier)`

SetSegmentIdentifiers sets SegmentIdentifiers field to given value.


### GetEnvironmentId

`func (o *ViewLinkRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ViewLinkRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ViewLinkRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ViewLinkRequest) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


