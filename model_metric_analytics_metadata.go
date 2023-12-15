/*
LaunchDarkly REST API

# Overview  ## Authentication  All REST API resources are authenticated with either [personal or service access tokens](https://docs.launchdarkly.com/home/account-security/api-access-tokens), or session cookies. Other authentication mechanisms are not supported. You can manage personal access tokens on your [**Account settings**](https://app.launchdarkly.com/settings/tokens) page.  LaunchDarkly also has SDK keys, mobile keys, and client-side IDs that are used by our server-side SDKs, mobile SDKs, and JavaScript-based SDKs, respectively. **These keys cannot be used to access our REST API**. These keys are environment-specific, and can only perform read-only operations such as fetching feature flag settings.  | Auth mechanism                                                                                  | Allowed resources                                                                                     | Use cases                                          | | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | -------------------------------------------------- | | [Personal or service access tokens](https://docs.launchdarkly.com/home/account-security/api-access-tokens) | Can be customized on a per-token basis                                                                | Building scripts, custom integrations, data export. | | SDK keys                                                                                        | Can only access read-only resources specific to server-side SDKs. Restricted to a single environment. | Server-side SDKs                     | | Mobile keys                                                                                     | Can only access read-only resources specific to mobile SDKs, and only for flags marked available to mobile keys. Restricted to a single environment.           | Mobile SDKs                                        | | Client-side ID                                                                                  | Can only access read-only resources specific to JavaScript-based client-side SDKs, and only for flags marked available to client-side. Restricted to a single environment.           | Client-side JavaScript                             |  > #### Keep your access tokens and SDK keys private > > Access tokens should _never_ be exposed in untrusted contexts. Never put an access token in client-side JavaScript, or embed it in a mobile application. LaunchDarkly has special mobile keys that you can embed in mobile apps. If you accidentally expose an access token or SDK key, you can reset it from your [**Account settings**](https://app.launchdarkly.com/settings/tokens) page. > > The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.  ### Authentication using request header  The preferred way to authenticate with the API is by adding an `Authorization` header containing your access token to your requests. The value of the `Authorization` header must be your access token.  Manage personal access tokens from the [**Account settings**](https://app.launchdarkly.com/settings/tokens) page.  ### Authentication using session cookie  For testing purposes, you can make API calls directly from your web browser. If you are logged in to the LaunchDarkly application, the API will use your existing session to authenticate calls.  If you have a [role](https://docs.launchdarkly.com/home/team/built-in-roles) other than Admin, or have a [custom role](https://docs.launchdarkly.com/home/team/custom-roles) defined, you may not have permission to perform some API calls. You will receive a `401` response code in that case.  > ### Modifying the Origin header causes an error > > LaunchDarkly validates that the Origin header for any API request authenticated by a session cookie matches the expected Origin header. The expected Origin header is `https://app.launchdarkly.com`. > > If the Origin header does not match what's expected, LaunchDarkly returns an error. This error can prevent the LaunchDarkly app from working correctly. > > Any browser extension that intentionally changes the Origin header can cause this problem. For example, the `Allow-Control-Allow-Origin: *` Chrome extension changes the Origin header to `http://evil.com` and causes the app to fail. > > To prevent this error, do not modify your Origin header. > > LaunchDarkly does not require origin matching when authenticating with an access token, so this issue does not affect normal API usage.  ## Representations  All resources expect and return JSON response bodies. Error responses also send a JSON body. To learn more about the error format of the API, read [Errors](/#section/Overview/Errors).  In practice this means that you always get a response with a `Content-Type` header set to `application/json`.  In addition, request bodies for `PATCH`, `POST`, and `PUT` requests must be encoded as JSON with a `Content-Type` header set to `application/json`.  ### Summary and detailed representations  When you fetch a list of resources, the response includes only the most important attributes of each resource. This is a _summary representation_ of the resource. When you fetch an individual resource, such as a single feature flag, you receive a _detailed representation_ of the resource.  The best way to find a detailed representation is to follow links. Every summary representation includes a link to its detailed representation.  ### Expanding responses  Sometimes the detailed representation of a resource does not include all of the attributes of the resource by default. If this is the case, the request method will clearly document this and describe which attributes you can include in an expanded response.  To include the additional attributes, append the `expand` request parameter to your request and add a comma-separated list of the attributes to include. For example, when you append `?expand=members,roles` to the [Get team](/tag/Teams#operation/getTeam) endpoint, the expanded response includes both of these attributes.  ### Links and addressability  The best way to navigate the API is by following links. These are attributes in representations that link to other resources. The API always uses the same format for links:  - Links to other resources within the API are encapsulated in a `_links` object - If the resource has a corresponding link to HTML content on the site, it is stored in a special `_site` link  Each link has two attributes:  - An `href`, which contains the URL - A `type`, which describes the content type  For example, a feature resource might return the following:  ```json {   \"_links\": {     \"parent\": {       \"href\": \"/api/features\",       \"type\": \"application/json\"     },     \"self\": {       \"href\": \"/api/features/sort.order\",       \"type\": \"application/json\"     }   },   \"_site\": {     \"href\": \"/features/sort.order\",     \"type\": \"text/html\"   } } ```  From this, you can navigate to the parent collection of features by following the `parent` link, or navigate to the site page for the feature by following the `_site` link.  Collections are always represented as a JSON object with an `items` attribute containing an array of representations. Like all other representations, collections have `_links` defined at the top level.  Paginated collections include `first`, `last`, `next`, and `prev` links containing a URL with the respective set of elements in the collection.  ## Updates  Resources that accept partial updates use the `PATCH` verb. Most resources support the [JSON patch](/reference#updates-using-json-patch) format. Some resources also support the [JSON merge patch](/reference#updates-using-json-merge-patch) format, and some resources support the [semantic patch](/reference#updates-using-semantic-patch) format, which is a way to specify the modifications to perform as a set of executable instructions. Each resource supports optional [comments](/reference#updates-with-comments) that you can submit with updates. Comments appear in outgoing webhooks, the audit log, and other integrations.  When a resource supports both JSON patch and semantic patch, we document both in the request method. However, the specific request body fields and descriptions included in our documentation only match one type of patch or the other.  ### Updates using JSON patch  [JSON patch](https://datatracker.ietf.org/doc/html/rfc6902) is a way to specify the modifications to perform on a resource. JSON patch uses paths and a limited set of operations to describe how to transform the current state of the resource into a new state. JSON patch documents are always arrays, where each element contains an operation, a path to the field to update, and the new value.  For example, in this feature flag representation:  ```json {     \"name\": \"New recommendations engine\",     \"key\": \"engine.enable\",     \"description\": \"This is the description\",     ... } ``` You can change the feature flag's description with the following patch document:  ```json [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"This is the new description\" }] ```  You can specify multiple modifications to perform in a single request. You can also test that certain preconditions are met before applying the patch:  ```json [   { \"op\": \"test\", \"path\": \"/version\", \"value\": 10 },   { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" } ] ```  The above patch request tests whether the feature flag's `version` is `10`, and if so, changes the feature flag's description.  Attributes that are not editable, such as a resource's `_links`, have names that start with an underscore.  ### Updates using JSON merge patch  [JSON merge patch](https://datatracker.ietf.org/doc/html/rfc7386) is another format for specifying the modifications to perform on a resource. JSON merge patch is less expressive than JSON patch. However, in many cases it is simpler to construct a merge patch document. For example, you can change a feature flag's description with the following merge patch document:  ```json {   \"description\": \"New flag description\" } ```  ### Updates using semantic patch  Some resources support the semantic patch format. A semantic patch is a way to specify the modifications to perform on a resource as a set of executable instructions.  Semantic patch allows you to be explicit about intent using precise, custom instructions. In many cases, you can define semantic patch instructions independently of the current state of the resource. This can be useful when defining a change that may be applied at a future date.  To make a semantic patch request, you must append `domain-model=launchdarkly.semanticpatch` to your `Content-Type` header.  Here's how:  ``` Content-Type: application/json; domain-model=launchdarkly.semanticpatch ```  If you call a semantic patch resource without this header, you will receive a `400` response because your semantic patch will be interpreted as a JSON patch.  The body of a semantic patch request takes the following properties:  * `comment` (string): (Optional) A description of the update. * `environmentKey` (string): (Required for some resources only) The environment key. * `instructions` (array): (Required) A list of actions the update should perform. Each action in the list must be an object with a `kind` property that indicates the instruction. If the instruction requires parameters, you must include those parameters as additional fields in the object. The documentation for each resource that supports semantic patch includes the available instructions and any additional parameters.  For example:  ```json {   \"comment\": \"optional comment\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  If any instruction in the patch encounters an error, the endpoint returns an error and will not change the resource. In general, each instruction silently does nothing if the resource is already in the state you request.  ### Updates with comments  You can submit optional comments with `PATCH` changes.  To submit a comment along with a JSON patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"patch\": [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" }] } ```  To submit a comment along with a JSON merge patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"merge\": { \"description\": \"New flag description\" } } ```  To submit a comment along with a semantic patch, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  ## Errors  The API always returns errors in a common format. Here's an example:  ```json {   \"code\": \"invalid_request\",   \"message\": \"A feature with that key already exists\",   \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\" } ```  The `code` indicates the general class of error. The `message` is a human-readable explanation of what went wrong. The `id` is a unique identifier. Use it when you're working with LaunchDarkly Support to debug a problem with a specific API call.  ### HTTP status error response codes  | Code | Definition        | Description                                                                                       | Possible Solution                                                | | ---- | ----------------- | ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- | | 400  | Invalid request       | The request cannot be understood.                                    | Ensure JSON syntax in request body is correct.                   | | 401  | Invalid access token      | Requestor is unauthorized or does not have permission for this API call.                                                | Ensure your API access token is valid and has the appropriate permissions.                                     | | 403  | Forbidden         | Requestor does not have access to this resource.                                                | Ensure that the account member or access token has proper permissions set. | | 404  | Invalid resource identifier | The requested resource is not valid. | Ensure that the resource is correctly identified by ID or key. | | 405  | Method not allowed | The request method is not allowed on this resource. | Ensure that the HTTP verb is correct. | | 409  | Conflict          | The API request can not be completed because it conflicts with a concurrent API request. | Retry your request.                                              | | 422  | Unprocessable entity | The API request can not be completed because the update description can not be understood. | Ensure that the request body is correct for the type of patch you are using, either JSON patch or semantic patch. | 429  | Too many requests | Read [Rate limiting](/#section/Overview/Rate-limiting).                                               | Wait and try again later.                                        |  ## CORS  The LaunchDarkly API supports Cross Origin Resource Sharing (CORS) for AJAX requests from any origin. If an `Origin` header is given in a request, it will be echoed as an explicitly allowed origin. Otherwise the request returns a wildcard, `Access-Control-Allow-Origin: *`. For more information on CORS, read the [CORS W3C Recommendation](http://www.w3.org/TR/cors). Example CORS headers might look like:  ```http Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, Authorization Access-Control-Allow-Methods: OPTIONS, GET, DELETE, PATCH Access-Control-Allow-Origin: * Access-Control-Max-Age: 300 ```  You can make authenticated CORS calls just as you would make same-origin calls, using either [token or session-based authentication](/#section/Overview/Authentication). If you are using session authentication, you should set the `withCredentials` property for your `xhr` request to `true`. You should never expose your access tokens to untrusted entities.  ## Rate limiting  We use several rate limiting strategies to ensure the availability of our APIs. Rate-limited calls to our APIs return a `429` status code. Calls to our APIs include headers indicating the current rate limit status. The specific headers returned depend on the API route being called. The limits differ based on the route, authentication mechanism, and other factors. Routes that are not rate limited may not contain any of the headers described below.  > ### Rate limiting and SDKs > > LaunchDarkly SDKs are never rate limited and do not use the API endpoints defined here. LaunchDarkly uses a different set of approaches, including streaming/server-sent events and a global CDN, to ensure availability to the routes used by LaunchDarkly SDKs.  ### Global rate limits  Authenticated requests are subject to a global limit. This is the maximum number of calls that your account can make to the API per ten seconds. All service and personal access tokens on the account share this limit, so exceeding the limit with one access token will impact other tokens. Calls that are subject to global rate limits may return the headers below:  | Header name                    | Description                                                                      | | ------------------------------ | -------------------------------------------------------------------------------- | | `X-Ratelimit-Global-Remaining` | The maximum number of requests the account is permitted to make per ten seconds. | | `X-Ratelimit-Reset`            | The time at which the current rate limit window resets in epoch milliseconds.    |  We do not publicly document the specific number of calls that can be made globally. This limit may change, and we encourage clients to program against the specification, relying on the two headers defined above, rather than hardcoding to the current limit.  ### Route-level rate limits  Some authenticated routes have custom rate limits. These also reset every ten seconds. Any service or personal access tokens hitting the same route share this limit, so exceeding the limit with one access token may impact other tokens. Calls that are subject to route-level rate limits return the headers below:  | Header name                   | Description                                                                                           | | ----------------------------- | ----------------------------------------------------------------------------------------------------- | | `X-Ratelimit-Route-Remaining` | The maximum number of requests to the current route the account is permitted to make per ten seconds. | | `X-Ratelimit-Reset`           | The time at which the current rate limit window resets in epoch milliseconds.                         |  A _route_ represents a specific URL pattern and verb. For example, the [Delete environment](/tag/Environments#operation/deleteEnvironment) endpoint is considered a single route, and each call to delete an environment counts against your route-level rate limit for that route.  We do not publicly document the specific number of calls that an account can make to each endpoint per ten seconds. These limits may change, and we encourage clients to program against the specification, relying on the two headers defined above, rather than hardcoding to the current limits.  ### IP-based rate limiting  We also employ IP-based rate limiting on some API routes. If you hit an IP-based rate limit, your API response will include a `Retry-After` header indicating how long to wait before re-trying the call. Clients must wait at least `Retry-After` seconds before making additional calls to our API, and should employ jitter and backoff strategies to avoid triggering rate limits again.  ## OpenAPI (Swagger) and client libraries  We have a [complete OpenAPI (Swagger) specification](https://app.launchdarkly.com/api/v2/openapi.json) for our API.  We auto-generate multiple client libraries based on our OpenAPI specification. To learn more, visit the [collection of client libraries on GitHub](https://github.com/search?q=topic%3Alaunchdarkly-api+org%3Alaunchdarkly&type=Repositories). You can also use this specification to generate client libraries to interact with our REST API in your language of choice.  Our OpenAPI specification is supported by several API-based tools such as Postman and Insomnia. In many cases, you can directly import our specification to explore our APIs.  ## Method overriding  Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those environments, our API endpoints that use `DELETE`, `PATCH`, and `PUT` verbs are inaccessible.  To avoid this issue, our API supports the `X-HTTP-Method-Override` header, allowing clients to \"tunnel\" `DELETE`, `PATCH`, and `PUT` requests using a `POST` request.  For example, to call a `PATCH` endpoint using a `POST` request, you can include `X-HTTP-Method-Override:PATCH` as a header.  ## Beta resources  We sometimes release new API resources in **beta** status before we release them with general availability.  Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  We try to promote resources into general availability as quickly as possible. This happens after sufficient testing and when we're satisfied that we no longer need to make backwards-incompatible changes.  We mark beta resources with a \"Beta\" callout in our documentation, pictured below:  > ### This feature is in beta > > To use this feature, pass in a header including the `LD-API-Version` key with value set to `beta`. Use this header with each call. To learn more, read [Beta resources](/#section/Overview/Beta-resources). > > Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  ### Using beta resources  To use a beta resource, you must include a header in the request. If you call a beta resource without this header, you receive a `403` response.  Use this header:  ``` LD-API-Version: beta ```  ## Federal environments  The version of LaunchDarkly that is available on domains controlled by the United States government is different from the version of LaunchDarkly available to the general public. If you are an employee or contractor for a United States federal agency and use LaunchDarkly in your work, you likely use the federal instance of LaunchDarkly.  If you are working in the federal instance of LaunchDarkly, the base URI for each request is `https://app.launchdarkly.us`. In the \"Try it\" sandbox for each request, click the request path to view the complete resource path for the federal environment.  To learn more, read [LaunchDarkly in federal environments](https://docs.launchdarkly.com/home/advanced/federal).  ## Versioning  We try hard to keep our REST API backwards compatible, but we occasionally have to make backwards-incompatible changes in the process of shipping new features. These breaking changes can cause unexpected behavior if you don't prepare for them accordingly.  Updates to our REST API include support for the latest features in LaunchDarkly. We also release a new version of our REST API every time we make a breaking change. We provide simultaneous support for multiple API versions so you can migrate from your current API version to a new version at your own pace.  ### Setting the API version per request  You can set the API version on a specific request by sending an `LD-API-Version` header, as shown in the example below:  ``` LD-API-Version: 20220603 ```  The header value is the version number of the API version you would like to request. The number for each version corresponds to the date the version was released in `yyyymmdd` format. In the example above the version `20220603` corresponds to June 03, 2022.  ### Setting the API version per access token  When you create an access token, you must specify a specific version of the API to use. This ensures that integrations using this token cannot be broken by version changes.  Tokens created before versioning was released have their version set to `20160426`, which is the version of the API that existed before the current versioning scheme, so that they continue working the same way they did before versioning.  If you would like to upgrade your integration to use a new API version, you can explicitly set the header described above.  > ### Best practice: Set the header for every client or integration > > We recommend that you set the API version header explicitly in any client or integration you build. > > Only rely on the access token API version during manual testing.  ### API version changelog  |<div style=\"width:75px\">Version</div> | Changes | End of life (EOL) |---|---|---| | `20220603` | <ul><li>Changed the [list projects](/tag/Projects#operation/getProjects) return value:<ul><li>Response is now paginated with a default limit of `20`.</li><li>Added support for filter and sort.</li><li>The project `environments` field is now expandable. This field is omitted by default.</li></ul></li><li>Changed the [get project](/tag/Projects#operation/getProject) return value:<ul><li>The `environments` field is now expandable. This field is omitted by default.</li></ul></li></ul> | Current | | `20210729` | <ul><li>Changed the [create approval request](/tag/Approvals#operation/postApprovalRequest) return value. It now returns HTTP Status Code `201` instead of `200`.</li><li> Changed the [get users](/tag/Users#operation/getUser) return value. It now returns a user record, not a user. </li><li>Added additional optional fields to environment, segments, flags, members, and segments, including the ability to create Big Segments. </li><li> Added default values for flag variations when new environments are created. </li><li>Added filtering and pagination for getting flags and members, including `limit`, `number`, `filter`, and `sort` query parameters. </li><li>Added endpoints for expiring user targets for flags and segments, scheduled changes, access tokens, Relay Proxy configuration, integrations and subscriptions, and approvals. </li></ul> | 2023-06-03 | | `20191212` | <ul><li>[List feature flags](/tag/Feature-flags#operation/getFeatureFlags) now defaults to sending summaries of feature flag configurations, equivalent to setting the query parameter `summary=true`. Summaries omit flag targeting rules and individual user targets from the payload. </li><li> Added endpoints for flags, flag status, projects, environments, audit logs, members, users, custom roles, segments, usage, streams, events, and data export. </li></ul> | 2022-07-29 | | `20160426` | <ul><li>Initial versioning of API. Tokens created before versioning have their version set to this.</li></ul> | 2020-12-12 | 

API version: 2.0
Contact: support@launchdarkly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ldapi

import (
	"encoding/json"
)

// MetricAnalyticsMetadata struct for MetricAnalyticsMetadata
type MetricAnalyticsMetadata struct {
	MetricKey string `json:"metricKey"`
	MetricName string `json:"metricName"`
	GroupBy *MetricQueryGroupByRep `json:"groupBy,omitempty"`
	GroupByValue *string `json:"groupByValue,omitempty"`
	Filters []MetricQueryFilterRep `json:"filters,omitempty"`
	UnitOfAnalysis string `json:"unitOfAnalysis"`
	Aggregation string `json:"aggregation"`
	AnalysisType string `json:"analysisType"`
	PercentileValue *int32 `json:"percentileValue,omitempty"`
	AnalysisValue *float32 `json:"analysisValue,omitempty"`
	NumberOfEvents int32 `json:"numberOfEvents"`
	IsHigherBetter bool `json:"isHigherBetter"`
	IsNumeric bool `json:"isNumeric"`
	Unit *string `json:"unit,omitempty"`
	DefaultMissingDataToZero bool `json:"defaultMissingDataToZero"`
}

// NewMetricAnalyticsMetadata instantiates a new MetricAnalyticsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricAnalyticsMetadata(metricKey string, metricName string, unitOfAnalysis string, aggregation string, analysisType string, numberOfEvents int32, isHigherBetter bool, isNumeric bool, defaultMissingDataToZero bool) *MetricAnalyticsMetadata {
	this := MetricAnalyticsMetadata{}
	this.MetricKey = metricKey
	this.MetricName = metricName
	this.UnitOfAnalysis = unitOfAnalysis
	this.Aggregation = aggregation
	this.AnalysisType = analysisType
	this.NumberOfEvents = numberOfEvents
	this.IsHigherBetter = isHigherBetter
	this.IsNumeric = isNumeric
	this.DefaultMissingDataToZero = defaultMissingDataToZero
	return &this
}

// NewMetricAnalyticsMetadataWithDefaults instantiates a new MetricAnalyticsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricAnalyticsMetadataWithDefaults() *MetricAnalyticsMetadata {
	this := MetricAnalyticsMetadata{}
	return &this
}

// GetMetricKey returns the MetricKey field value
func (o *MetricAnalyticsMetadata) GetMetricKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricKey
}

// GetMetricKeyOk returns a tuple with the MetricKey field value
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetMetricKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricKey, true
}

// SetMetricKey sets field value
func (o *MetricAnalyticsMetadata) SetMetricKey(v string) {
	o.MetricKey = v
}

// GetMetricName returns the MetricName field value
func (o *MetricAnalyticsMetadata) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *MetricAnalyticsMetadata) SetMetricName(v string) {
	o.MetricName = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *MetricAnalyticsMetadata) GetGroupBy() MetricQueryGroupByRep {
	if o == nil || o.GroupBy == nil {
		var ret MetricQueryGroupByRep
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetGroupByOk() (*MetricQueryGroupByRep, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *MetricAnalyticsMetadata) HasGroupBy() bool {
	if o != nil && o.GroupBy != nil {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given MetricQueryGroupByRep and assigns it to the GroupBy field.
func (o *MetricAnalyticsMetadata) SetGroupBy(v MetricQueryGroupByRep) {
	o.GroupBy = &v
}

// GetGroupByValue returns the GroupByValue field value if set, zero value otherwise.
func (o *MetricAnalyticsMetadata) GetGroupByValue() string {
	if o == nil || o.GroupByValue == nil {
		var ret string
		return ret
	}
	return *o.GroupByValue
}

// GetGroupByValueOk returns a tuple with the GroupByValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetGroupByValueOk() (*string, bool) {
	if o == nil || o.GroupByValue == nil {
		return nil, false
	}
	return o.GroupByValue, true
}

// HasGroupByValue returns a boolean if a field has been set.
func (o *MetricAnalyticsMetadata) HasGroupByValue() bool {
	if o != nil && o.GroupByValue != nil {
		return true
	}

	return false
}

// SetGroupByValue gets a reference to the given string and assigns it to the GroupByValue field.
func (o *MetricAnalyticsMetadata) SetGroupByValue(v string) {
	o.GroupByValue = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *MetricAnalyticsMetadata) GetFilters() []MetricQueryFilterRep {
	if o == nil || o.Filters == nil {
		var ret []MetricQueryFilterRep
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetFiltersOk() ([]MetricQueryFilterRep, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *MetricAnalyticsMetadata) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []MetricQueryFilterRep and assigns it to the Filters field.
func (o *MetricAnalyticsMetadata) SetFilters(v []MetricQueryFilterRep) {
	o.Filters = v
}

// GetUnitOfAnalysis returns the UnitOfAnalysis field value
func (o *MetricAnalyticsMetadata) GetUnitOfAnalysis() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitOfAnalysis
}

// GetUnitOfAnalysisOk returns a tuple with the UnitOfAnalysis field value
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetUnitOfAnalysisOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfAnalysis, true
}

// SetUnitOfAnalysis sets field value
func (o *MetricAnalyticsMetadata) SetUnitOfAnalysis(v string) {
	o.UnitOfAnalysis = v
}

// GetAggregation returns the Aggregation field value
func (o *MetricAnalyticsMetadata) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *MetricAnalyticsMetadata) SetAggregation(v string) {
	o.Aggregation = v
}

// GetAnalysisType returns the AnalysisType field value
func (o *MetricAnalyticsMetadata) GetAnalysisType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnalysisType
}

// GetAnalysisTypeOk returns a tuple with the AnalysisType field value
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetAnalysisTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalysisType, true
}

// SetAnalysisType sets field value
func (o *MetricAnalyticsMetadata) SetAnalysisType(v string) {
	o.AnalysisType = v
}

// GetPercentileValue returns the PercentileValue field value if set, zero value otherwise.
func (o *MetricAnalyticsMetadata) GetPercentileValue() int32 {
	if o == nil || o.PercentileValue == nil {
		var ret int32
		return ret
	}
	return *o.PercentileValue
}

// GetPercentileValueOk returns a tuple with the PercentileValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetPercentileValueOk() (*int32, bool) {
	if o == nil || o.PercentileValue == nil {
		return nil, false
	}
	return o.PercentileValue, true
}

// HasPercentileValue returns a boolean if a field has been set.
func (o *MetricAnalyticsMetadata) HasPercentileValue() bool {
	if o != nil && o.PercentileValue != nil {
		return true
	}

	return false
}

// SetPercentileValue gets a reference to the given int32 and assigns it to the PercentileValue field.
func (o *MetricAnalyticsMetadata) SetPercentileValue(v int32) {
	o.PercentileValue = &v
}

// GetAnalysisValue returns the AnalysisValue field value if set, zero value otherwise.
func (o *MetricAnalyticsMetadata) GetAnalysisValue() float32 {
	if o == nil || o.AnalysisValue == nil {
		var ret float32
		return ret
	}
	return *o.AnalysisValue
}

// GetAnalysisValueOk returns a tuple with the AnalysisValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetAnalysisValueOk() (*float32, bool) {
	if o == nil || o.AnalysisValue == nil {
		return nil, false
	}
	return o.AnalysisValue, true
}

// HasAnalysisValue returns a boolean if a field has been set.
func (o *MetricAnalyticsMetadata) HasAnalysisValue() bool {
	if o != nil && o.AnalysisValue != nil {
		return true
	}

	return false
}

// SetAnalysisValue gets a reference to the given float32 and assigns it to the AnalysisValue field.
func (o *MetricAnalyticsMetadata) SetAnalysisValue(v float32) {
	o.AnalysisValue = &v
}

// GetNumberOfEvents returns the NumberOfEvents field value
func (o *MetricAnalyticsMetadata) GetNumberOfEvents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfEvents
}

// GetNumberOfEventsOk returns a tuple with the NumberOfEvents field value
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetNumberOfEventsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfEvents, true
}

// SetNumberOfEvents sets field value
func (o *MetricAnalyticsMetadata) SetNumberOfEvents(v int32) {
	o.NumberOfEvents = v
}

// GetIsHigherBetter returns the IsHigherBetter field value
func (o *MetricAnalyticsMetadata) GetIsHigherBetter() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHigherBetter
}

// GetIsHigherBetterOk returns a tuple with the IsHigherBetter field value
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetIsHigherBetterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHigherBetter, true
}

// SetIsHigherBetter sets field value
func (o *MetricAnalyticsMetadata) SetIsHigherBetter(v bool) {
	o.IsHigherBetter = v
}

// GetIsNumeric returns the IsNumeric field value
func (o *MetricAnalyticsMetadata) GetIsNumeric() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNumeric
}

// GetIsNumericOk returns a tuple with the IsNumeric field value
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetIsNumericOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsNumeric, true
}

// SetIsNumeric sets field value
func (o *MetricAnalyticsMetadata) SetIsNumeric(v bool) {
	o.IsNumeric = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MetricAnalyticsMetadata) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricAnalyticsMetadata) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *MetricAnalyticsMetadata) SetUnit(v string) {
	o.Unit = &v
}

// GetDefaultMissingDataToZero returns the DefaultMissingDataToZero field value
func (o *MetricAnalyticsMetadata) GetDefaultMissingDataToZero() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultMissingDataToZero
}

// GetDefaultMissingDataToZeroOk returns a tuple with the DefaultMissingDataToZero field value
// and a boolean to check if the value has been set.
func (o *MetricAnalyticsMetadata) GetDefaultMissingDataToZeroOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMissingDataToZero, true
}

// SetDefaultMissingDataToZero sets field value
func (o *MetricAnalyticsMetadata) SetDefaultMissingDataToZero(v bool) {
	o.DefaultMissingDataToZero = v
}

func (o MetricAnalyticsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metricKey"] = o.MetricKey
	}
	if true {
		toSerialize["metricName"] = o.MetricName
	}
	if o.GroupBy != nil {
		toSerialize["groupBy"] = o.GroupBy
	}
	if o.GroupByValue != nil {
		toSerialize["groupByValue"] = o.GroupByValue
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if true {
		toSerialize["unitOfAnalysis"] = o.UnitOfAnalysis
	}
	if true {
		toSerialize["aggregation"] = o.Aggregation
	}
	if true {
		toSerialize["analysisType"] = o.AnalysisType
	}
	if o.PercentileValue != nil {
		toSerialize["percentileValue"] = o.PercentileValue
	}
	if o.AnalysisValue != nil {
		toSerialize["analysisValue"] = o.AnalysisValue
	}
	if true {
		toSerialize["numberOfEvents"] = o.NumberOfEvents
	}
	if true {
		toSerialize["isHigherBetter"] = o.IsHigherBetter
	}
	if true {
		toSerialize["isNumeric"] = o.IsNumeric
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["defaultMissingDataToZero"] = o.DefaultMissingDataToZero
	}
	return json.Marshal(toSerialize)
}

type NullableMetricAnalyticsMetadata struct {
	value *MetricAnalyticsMetadata
	isSet bool
}

func (v NullableMetricAnalyticsMetadata) Get() *MetricAnalyticsMetadata {
	return v.value
}

func (v *NullableMetricAnalyticsMetadata) Set(val *MetricAnalyticsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricAnalyticsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricAnalyticsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricAnalyticsMetadata(val *MetricAnalyticsMetadata) *NullableMetricAnalyticsMetadata {
	return &NullableMetricAnalyticsMetadata{value: val, isSet: true}
}

func (v NullableMetricAnalyticsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricAnalyticsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


