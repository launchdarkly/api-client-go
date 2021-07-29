/*
 * LaunchDarkly REST API
 *
 * # Authentication  All REST API resources are authenticated with [personal access tokens](https://docs.launchdarkly.com/home/account-security/api-access-tokens) or session cookies. Other authentication mechanisms are not supported. You can manage personal access tokens on your [Account settings](https://app.launchdarkly.com/settings/tokens) page.   LaunchDarkly also has SDK keys, mobile keys, and client-side IDs that are used by our server-side SDKs, mobile SDKs, and client-side JavaScript SDKs, respectively. **These keys cannot be used to access our REST API**. These keys are environment-specific, and can only perform read-only operations (fetching feature flag settings).  | Auth mechanism | Allowed resources | Use cases | | -------------- | ----------------- | --------- | | [Personal access tokens](https://docs.launchdarkly.com/home/account-security/api-access-tokens) | Can be customized on a per-token basis | Building scripts, custom integrations, data export | | SDK keys | Can only access read-only SDK-specific resources and the firehose, restricted to a single environment | Server-side SDKs, Firehose API | | Mobile keys | Can only access read-only mobile SDK-specific resources, restricted to a single environment | Mobile SDKs | | Client-side ID | Single environment, only flags marked available to client-side | Client-side JavaScript |  <blockquote>     <h3><span>❗️</span>Keep your access tokens and SDK keys private</h3>     <p>Access tokens should *never* be exposed in untrusted contexts. Never put an access token in client-side JavaScript, or embed it in a mobile application. LaunchDarkly has special mobile keys that you can embed in mobile apps. If you accidentally expose an access token or SDK key, you can reset it from your [Account Settings](https://app.launchdarkly.com/settings#/tokens) page.</p>     <p>The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.</p> </blockquote>  ## Via request header  The preferred way to authenticate with the API is by adding an `Authorization` header containing your access token to your requests. The value of the `Authorization` header must be your access token.  Manage personal access tokens from the [Account Settings](https://app.launchdarkly.com/settings/tokens) page.  ## Via session cookie  For testing purposes, you can make API calls directly from your web browser. If you're logged in to the application, the API will use your existing session to authenticate calls.  If you have a [role](https://docs.launchdarkly.com/home/team/built-in-roles) other than Admin, or have a [custom role](https://docs.launchdarkly.com/home/team/custom-roles) defined, you may not have permission to perform some API calls. You will receive a `401` response code in that case.  <blockquote>     <h3><span>❗️</span>Modifying the Origin header causes an error</h3>     <p>We validate that the Origin header for any API request authenticated by a session cookie matches the expected Origin header. The expected Origin header is `https://app.launchdarkly.com`.</p>     <p>If the Origin header does not match what's expected, LaunchDarkly returns an error. This error can prevent the LaunchDarkly app from working correctly. </p>     <p>Any browser extension that intentionally changes the Origin header can cause this problem. For example, the `Allow-Control-Allow-Origin: *` Chrome extension changes the Origin header to http://evil.com and causes the app to fail.</p>     <p>To prevent this error, do not modify your Origin header.</p>     <p>LaunchDarkly does not require origin matching when authenticating with an Access Token, so this issue does not affect normal API usage.</p> </blockquote>  # Representations  All resources expect and return JSON response bodies. Error responses will also send a JSON body. Read [Errors](#section/Errors) for a more detailed description of the error format used by the API.   In practice this means that you always get a response with a `Content-Type` header set to `application/json`.  In addition, request bodies for `PUT`, `POST`, `REPORT` and `PATCH` requests must be encoded as JSON with a `Content-Type` header set to `application/json`.  ## Summary and detailed representations  When you fetch a list of resources, the response includes only the most important attributes of each resource. This is a *summary representation* of the resource. When you fetch an individual resource (for example, a single feature flag), you receive a *detailed representation* containing all of the attributes of the resource.  The best way to find a detailed representation is to follow links. Every summary representation includes a link to its detailed representation.  ## Links and addressability  The best way to navigate the API is by following links. These are attributes in representations that link to other resources. The API always uses the same format for links:  * Links to other resources within the API are encapsulated in a `_links` object. * If the resource has a corresponding link to HTML content on the site, it is stored in a special `_site` link.  Each link has two attributes: an href (the URL) and a type (the content type). For example, a feature resource might return the following:  ```json {     \"_links\": {         \"parent\": {             \"href\": \"/api/features\",             \"type\": \"application/json\"         },         \"self\": {             \"href\":\"/api/features/sort.order\",             \"type\":\"application/json\"         }     },     \"_site\":{         \"href\":\"/features/sort.order\",         \"type\":\"text/html\"     } } ```  From this, you can navigate to the parent collection of features by following the `parent` link, or navigate to the site page for the feature by following the `_site` link.  Collections are always represented as a JSON object with an `items` attribute containing an array of representations. Like all other representations, collections have `_links` defined at the top level.  Paginated collections include `first`, `last`, `next`, and `prev` links containing a URL with the respective set of elements in the collection.  # Updates  Resources that accept partial updates use the `PATCH` verb, and support the [JSON Patch](http://tools.ietf.org/html/rfc6902) format. Some resources also support the [JSON Merge Patch](https://tools.ietf.org/html/rfc7386) format. In addition, some resources support optional comments that can be submitted with updates. Comments appear in outgoing webhooks, the audit log, and other integrations.  ## Updates via JSON Patch  [JSON Patch](http://tools.ietf.org/html/rfc6902) is a way to specify the modifications to perform on a resource. For example, in this feature flag representation:  ```json {     \"name\": \"New recommendations engine\",     \"key\": \"engine.enable\",     \"description\": \"This is the description\",     ... } ```  You can change the feature flag's description with the following patch document:  ```json [     { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"This is the new description\"} ] ```  JSON Patch documents are always arrays. You can specify multiple modifications to perform in a single request. You can also test that certain preconditions are met before applying the patch:  ```json [     { \"op\": \"test\", \"path\": \"/version\", \"value\": 10 },     { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" } ] ```  The above patch request tests whether the feature flag's `version` is `10`, and if so, changes the feature flag's description.  Attributes that aren't editable, like a resource's `_links`, have names that start with an underscore.  ## Updates via JSON Merge Patch  The API also supports the [JSON Merge Patch](https://tools.ietf.org/html/rfc7386) format, as well as the [Update feature flag](#operation/patchFeatureFlag) resource.   JSON Merge Patch is less expressive than JSON Patch but in many cases, it is simpler to construct a merge patch document. For example, you can change a feature flag's description with the following merge patch document:  ```json {     \"description\": \"New flag description\" } ```  ## Updates with comments  You can submit optional comments with `PATCH` changes. The [Update feature flag](#operation/patchFeatureFlag) resource supports comments.  To submit a comment along with a JSON Patch document, use the following format:  ```json {     \"comment\": \"This is a comment string\",     \"patch\": [ {\"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" } ] } ```  To submit a comment along with a JSON Merge Patch document, use the following format:  ```json {     \"comment\": \"This is a comment string\",     \"merge\": { \"description\": \"New flag description\"} } ```  ## Updates via semantic patches  The API also supports the Semantic patch format. A semantic `PATCH` is a way to specify the modifications to perform on a resource as a set of executable instructions.   JSON Patch uses paths and a limited set of operations to describe how to transform the current state of the resource into a new state. Semantic patch allows you to be explicit about intent using precise, custom instructions. In many cases, semantic patch instructions can also be defined independently of the current state of the resource. This can be useful when defining a change that may be applied at a future date.  For example, in this feature flag configuration in environment Production:  ```json {     \"name\": \"Alternate sort order\",     \"kind\": \"boolean\",     \"key\": \"sort.order\",    ...     \"environments\": {         \"production\": {             \"on\": true,             \"archived\": false,             \"salt\": \"c29ydC5vcmRlcg==\",             \"sel\": \"8de1085cb7354b0ab41c0e778376dfd3\",             \"lastModified\": 1469131558260,             \"version\": 81,             \"targets\": [                 {                     \"values\": [                         \"Gerhard.Little@yahoo.com\"                     ],                     \"variation\": 0                 },                 {                     \"values\": [                         \"1461797806429-33-861961230\",                         \"438580d8-02ee-418d-9eec-0085cab2bdf0\"                     ],                     \"variation\": 1                 }             ],             \"rules\": [],             \"fallthrough\": {                 \"variation\": 0             },             \"offVariation\": 1,             \"prerequisites\": [],             \"_site\": {                 \"href\": \"/default/production/features/sort.order\",                 \"type\": \"text/html\"             }        }     } } ```  You can add a date you want a user to be removed from the feature flag's user targets. For example, “remove user 1461797806429-33-861961230 from the user target for variation 0 on the Alternate sort order flag in the production environment on Wed Jul 08 2020 at 15:27:41 pm”. This is done using the following:  ```json {     \"comment\": \"update expiring user targets\",     \"instructions\": [         {             \"kind\": \"removeExpireUserTargetDate\",             \"userKey\": \"userKey\",             \"variationId\": \"978d53f9-7fe3-4a63-992d-97bcb4535dc8\",         },         {             \"kind\": \"updateExpireUserTargetDate\",             \"userKey\": \"userKey2\",             \"variationId\": \"978d53f9-7fe3-4a63-992d-97bcb4535dc8\",             \"value\": 1587582000000         },         {             \"kind\": \"addExpireUserTargetDate\",             \"userKey\": \"userKey3\",             \"variationId\": \"978d53f9-7fe3-4a63-992d-97bcb4535dc8\",             \"value\": 1594247266386         }     ] } ```  Here is another example. In this feature flag configuration:  ```json {     \"name\": \"New recommendations engine\",     \"key\": \"engine.enable\",     \"environments\": {         \"test\": {             \"on\": true         }     } } ```  You can change the feature flag's description with the following patch document as a set of executable instructions. For example, “add user X to targets for variation Y and remove user A from targets for variation B for test flag”:  ```json {     \"comment\": \"\",     \"instructions\": [         {             \"kind\": \"removeUserTargets\",             \"values\": [\"438580d8-02ee-418d-9eec-0085cab2bdf0\"],             \"variationId\": \"852cb784-54ff-46b9-8c35-5498d2e4f270\"         },         {             \"kind\": \"addUserTargets\",             \"values\": [\"438580d8-02ee-418d-9eec-0085cab2bdf0\"],             \"variationId\": \"1bb18465-33b6-49aa-a3bd-eeb6650b33ad\"         }     ] } ```  <blockquote>     <h3><span>📘</span>Supported semantic patch API endpoints</h3>     <p>TODO: update these links</p>     <p><a href=\"#operation/patchFeatureFlag\">Update feature flag</a></p>     <p><a href=\"\">Update expiring user targets on feature flag</a></p>     <p><a href=\"\">Update expiring user target for flags</a></p> </blockquote>  # Errors  The API always returns errors in a common format. Here's an example:  ```json {     \"code\": \"invalid_request\",     \"message\": \"A feature with that key already exists\",     \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\" } ```  The general class of error is indicated by the `code`. The `message` is a human-readable explanation of what went wrong. The `id` is a unique identifier. Use it when you're working with LaunchDarkly support to debug a problem with a specific API call.  ## HTTP Status - Error Response Codes  | Code | Definition | Desc. | Possible Solution | | ---- | ---------- | ----- | ----------------- | | 400  | Bad Request | A request that fails may return this HTTP response code. | Ensure JSON syntax in request body is correct. | | 401  | Unauthorized | User doesn't have permission to an API call. | Ensure your SDK key is good. | | 403  | Forbidden | User does not have permission for operation. | Ensure that the user or access token has proper permissions set. | | 409  | Conflict | The API request could not be completed because it conflicted with a concurrent API request. | Retry your request. | | 429  | Too many requests | See [Rate limiting](ref:rate-limiting). | Wait and try again later. |  # CORS  The LaunchDarkly API supports Cross Origin Resource Sharing (CORS) for AJAX requests from any origin. If an `Origin` header is given in a request, it will be echoed as an explicitly allowed origin. Otherwise, a wildcard is returned: `Access-Control-Allow-Origin: *`. For more information on CORS, see the [CORS W3C Recommendation](http://www.w3.org/TR/cors). Example CORS headers might look like:  ```http Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, Authorization Access-Control-Allow-Methods: OPTIONS, GET, DELETE, PATCH Access-Control-Allow-Origin: * Access-Control-Max-Age: 300 ```  You can make authenticated CORS calls just as you would make same-origin calls, using either [token or session-based authentication](#section/Authentication). If you’re using session auth, you should set the `withCredentials` property for your `xhr` request to `true`. You should never expose your access tokens to untrusted users.  # Rate limiting  We use several rate limiting strategies to ensure the availability of our APIs. Rate-limited calls to our APIs will return a `429` status code. Calls to our APIs will include headers indicating the current rate limit status. The specific headers returned depend on the API route being called-- the limits differ based on the route, authentication mechanism, and other factors. Routes that are not rate limited may not contain any of the headers described below.  <blockquote>     <h3><span>❗️</span>Rate limiting and SDKs</h3>     <p>Our SDKs are never rate limited. Our SDKs do not use the API endpoints defined here. We use a different set of approaches, including streaming/server-sent events and a global CDN, to ensure availability to the routes used by our SDKs.</p>     <p>The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.</p> </blockquote>  ## Global rate limits  Authenticated requests are subject to a global limit. This is the maximum number of calls that can be made to the API per ten seconds. All personal access tokens on the account share this limit, so exceeding the limit with one access token will impact other tokens. Calls that are subject to global rate limits will return the headers below:  | Header name | Description | | ----------- | ----------- | | `X-Ratelimit-Global-Remaining` | The maximum number of requests the account is permitted to make per ten seconds. | | `X-Ratelimit-Reset` | The time at which the current rate limit window resets in epoch milliseconds. |  We do not publicly document the specific number of calls that can be made globally. This limit may change, and we encourage clients to program against the specification, relying on the two headers defined above, rather than hardcoding to the current limit.  ## Route-level rate limits  Some authenticated routes have custom rate limits. These also reset every ten seconds. Any access tokens hitting the same route share this limit, so exceeding the limit with one access token may impact other tokens. Calls that are subject to route-level rate limits will return the headers below:  | Header name | Description | | ----------- | ----------- | | `X-Ratelimit-Route-Remaining` | The maximum number of requests to the current route the account is permitted to make per ten seconds. | | `X-Ratelimit-Reset` | The time at which the current rate limit window resets in epoch milliseconds. |  A *route* represents a specific URL pattern and verb. For example, the [Delete environment](#operation/deleteEnvironment) endpoint is considered a single route, and each call to delete an environment counts against your route-level rate limit for that route.   We do not publicly document the specific number of calls that can be made to each endpoint per ten seconds. These limits may change, and we encourage clients to program against the specification, relying on the two headers defined above, rather than hardcoding to the current limits.  ## IP-based rate limiting  We also employ IP-based rate limiting on some API routes. If you hit an IP-based rate limit, your API response will include a `Retry-After` header indicating how long to wait before re-trying the call. Clients must wait at least `Retry-After` seconds before making additional calls to our API, and should employ jitter and backoff strategies to avoid triggering rate limits again.  # OpenAPI (Swagger)  We have a [complete OpenAPI (Swagger) specification](https://app.launchdarkly.com/api/v2/openapi.json) for our API.  You can use this specification to generate client libraries to interact with our REST API in your language of choice.   # Client libraries  We auto-generate multiple client libraries based on our OpenAPI specification. To learn more, visit [GitHub](https://github.com/search?q=topic%3Alaunchdarkly-api+org%3Alaunchdarkly&type=Repositories).  # Method Overriding  Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those environments, our API endpoints that use `PUT`, `PATCH`, and `DELETE` verbs will be inaccessible.  To avoid this issue, our API supports the `X-HTTP-Method-Override` header, allowing clients to \"tunnel\" `PUT`, `PATCH`, and `DELETE` requests via a `POST` request.   For example, if you wish to call one of our `PATCH` resources via a `POST` request, you can include `X-HTTP-Method-Override:PATCH` as a header.  # Beta resources  We sometimes release new API resources in **beta** status before we release them with general availability.   Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.   We try to promote resources into general availability as quickly as possible. This happens after sufficient testing and when we're satisfied that we no longer need to make backwards-incompatible changes.  We mark beta resources with a \"Beta\" callout in our documentation, pictured below: <blockquote>     <h3><span>📘</span>Beta</h3>     <p>**This feature is in beta.** You must include a specific header to use it.\\n\\nTo learn more, read [Beta resources](ref:beta-resources).</p> </blockquote>  ## Using beta resources  To use a beta resource, you must include a header in the request. If you call a beta resource without this header, you'll receive a `403` response.  Use this header:   ``` LD-API-Version: beta ```  # Versioning  We try hard to keep our REST API backwards compatible, but we occasionally have to make backwards-incompatible changes in the process of shipping new features. These breaking changes can cause unexpected behavior if you don't prepare for them accordingly.   Updates to our REST API include support for the latest features in LaunchDarkly. We also release a new version of our REST API every time we make a breaking change. We provide simultaneous support for multiple API versions so you can migrate from your current API version to a new version at your own pace.   See new versions in the [Changelog](ref:changelog).  ## Setting the API version per request  You can set the API version on a specific request by sending an `LD-API-Version` header, as shown in the example below:  ``` LD-API-Version: 20191212 ```  The header value is the version number of the API version you'd like to request. The number for each version corresponds to the date the version was released. In the example above the version `20191212` corresponds to December 12, 2019.   ## Setting the API version per access token  When creating an access token, you must specify a specific version of the API to use. This ensures that integrations using this token cannot be broken by version changes.  Tokens created before versioning was released have their version set to `20160426` (the version of the API that existed before versioning) so that they continue working the same way they did before versioning.  If you would like to upgrade your integration to use a new API version, you can explicitly set the header described above.  <blockquote>     <h3><span>👍</span>Best practice: Set the header for every client or integration</h3>     <p>We recommend that you set the API version header explicitly in any client or integration you build.</p>     <p>Only rely on the access token API version during manual testing.</p> </blockquote>  <blockquote>     <h3><span>🚧</span>API Path Versioning</h3>     <p>In the past, we've used path-based API versioning. For example, versioning resources by adding `v2` to endpoint URLs. We don't foresee the need to do this again, but may do so if we need to make major revisions to the API.</p> </blockquote> 
 *
 * API version: 2.0
 * Contact: support@launchdarkly.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ldapi

import (
	"encoding/json"
)

// WebFlagConfigApprovalRequestResponse struct for WebFlagConfigApprovalRequestResponse
type WebFlagConfigApprovalRequestResponse struct {
	Id *string `json:"_id,omitempty"`
	Version *int32 `json:"_version,omitempty"`
	CreationDate *int64 `json:"creationDate,omitempty"`
	ServiceKind *string `json:"serviceKind,omitempty"`
	RequestorId *string `json:"requestorId,omitempty"`
	Description *string `json:"description,omitempty"`
	ReviewStatus *string `json:"reviewStatus,omitempty"`
	AllReviews *[]WebFlagConfigApprovalRequestResponseAllReviews `json:"allReviews,omitempty"`
	NotifyMemberIds *[]string `json:"notifyMemberIds,omitempty"`
	AppliedDate *int64 `json:"appliedDate,omitempty"`
	AppliedByMemberId *string `json:"appliedByMemberId,omitempty"`
	Status *string `json:"status,omitempty"`
	Instructions *[]map[string]interface{} `json:"instructions,omitempty"`
	Conflicts *[]WebFlagConfigApprovalRequestResponseConflicts `json:"conflicts,omitempty"`
	Links *map[string]InlineResponse200 `json:"_links,omitempty"`
	ExecutionDate *int64 `json:"executionDate,omitempty"`
	OperatingOnId *string `json:"operatingOnId,omitempty"`
	IntegrationMetadata *WebIntegrationMetadata `json:"integrationMetadata,omitempty"`
	Source *WebCopiedFromEnv `json:"source,omitempty"`
}

// NewWebFlagConfigApprovalRequestResponse instantiates a new WebFlagConfigApprovalRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebFlagConfigApprovalRequestResponse() *WebFlagConfigApprovalRequestResponse {
	this := WebFlagConfigApprovalRequestResponse{}
	return &this
}

// NewWebFlagConfigApprovalRequestResponseWithDefaults instantiates a new WebFlagConfigApprovalRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebFlagConfigApprovalRequestResponseWithDefaults() *WebFlagConfigApprovalRequestResponse {
	this := WebFlagConfigApprovalRequestResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebFlagConfigApprovalRequestResponse) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *WebFlagConfigApprovalRequestResponse) SetVersion(v int32) {
	o.Version = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetCreationDate() int64 {
	if o == nil || o.CreationDate == nil {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetCreationDateOk() (*int64, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *WebFlagConfigApprovalRequestResponse) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetServiceKind returns the ServiceKind field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetServiceKind() string {
	if o == nil || o.ServiceKind == nil {
		var ret string
		return ret
	}
	return *o.ServiceKind
}

// GetServiceKindOk returns a tuple with the ServiceKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetServiceKindOk() (*string, bool) {
	if o == nil || o.ServiceKind == nil {
		return nil, false
	}
	return o.ServiceKind, true
}

// HasServiceKind returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasServiceKind() bool {
	if o != nil && o.ServiceKind != nil {
		return true
	}

	return false
}

// SetServiceKind gets a reference to the given string and assigns it to the ServiceKind field.
func (o *WebFlagConfigApprovalRequestResponse) SetServiceKind(v string) {
	o.ServiceKind = &v
}

// GetRequestorId returns the RequestorId field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetRequestorId() string {
	if o == nil || o.RequestorId == nil {
		var ret string
		return ret
	}
	return *o.RequestorId
}

// GetRequestorIdOk returns a tuple with the RequestorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetRequestorIdOk() (*string, bool) {
	if o == nil || o.RequestorId == nil {
		return nil, false
	}
	return o.RequestorId, true
}

// HasRequestorId returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasRequestorId() bool {
	if o != nil && o.RequestorId != nil {
		return true
	}

	return false
}

// SetRequestorId gets a reference to the given string and assigns it to the RequestorId field.
func (o *WebFlagConfigApprovalRequestResponse) SetRequestorId(v string) {
	o.RequestorId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WebFlagConfigApprovalRequestResponse) SetDescription(v string) {
	o.Description = &v
}

// GetReviewStatus returns the ReviewStatus field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetReviewStatus() string {
	if o == nil || o.ReviewStatus == nil {
		var ret string
		return ret
	}
	return *o.ReviewStatus
}

// GetReviewStatusOk returns a tuple with the ReviewStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetReviewStatusOk() (*string, bool) {
	if o == nil || o.ReviewStatus == nil {
		return nil, false
	}
	return o.ReviewStatus, true
}

// HasReviewStatus returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasReviewStatus() bool {
	if o != nil && o.ReviewStatus != nil {
		return true
	}

	return false
}

// SetReviewStatus gets a reference to the given string and assigns it to the ReviewStatus field.
func (o *WebFlagConfigApprovalRequestResponse) SetReviewStatus(v string) {
	o.ReviewStatus = &v
}

// GetAllReviews returns the AllReviews field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetAllReviews() []WebFlagConfigApprovalRequestResponseAllReviews {
	if o == nil || o.AllReviews == nil {
		var ret []WebFlagConfigApprovalRequestResponseAllReviews
		return ret
	}
	return *o.AllReviews
}

// GetAllReviewsOk returns a tuple with the AllReviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetAllReviewsOk() (*[]WebFlagConfigApprovalRequestResponseAllReviews, bool) {
	if o == nil || o.AllReviews == nil {
		return nil, false
	}
	return o.AllReviews, true
}

// HasAllReviews returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasAllReviews() bool {
	if o != nil && o.AllReviews != nil {
		return true
	}

	return false
}

// SetAllReviews gets a reference to the given []WebFlagConfigApprovalRequestResponseAllReviews and assigns it to the AllReviews field.
func (o *WebFlagConfigApprovalRequestResponse) SetAllReviews(v []WebFlagConfigApprovalRequestResponseAllReviews) {
	o.AllReviews = &v
}

// GetNotifyMemberIds returns the NotifyMemberIds field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetNotifyMemberIds() []string {
	if o == nil || o.NotifyMemberIds == nil {
		var ret []string
		return ret
	}
	return *o.NotifyMemberIds
}

// GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetNotifyMemberIdsOk() (*[]string, bool) {
	if o == nil || o.NotifyMemberIds == nil {
		return nil, false
	}
	return o.NotifyMemberIds, true
}

// HasNotifyMemberIds returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasNotifyMemberIds() bool {
	if o != nil && o.NotifyMemberIds != nil {
		return true
	}

	return false
}

// SetNotifyMemberIds gets a reference to the given []string and assigns it to the NotifyMemberIds field.
func (o *WebFlagConfigApprovalRequestResponse) SetNotifyMemberIds(v []string) {
	o.NotifyMemberIds = &v
}

// GetAppliedDate returns the AppliedDate field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetAppliedDate() int64 {
	if o == nil || o.AppliedDate == nil {
		var ret int64
		return ret
	}
	return *o.AppliedDate
}

// GetAppliedDateOk returns a tuple with the AppliedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetAppliedDateOk() (*int64, bool) {
	if o == nil || o.AppliedDate == nil {
		return nil, false
	}
	return o.AppliedDate, true
}

// HasAppliedDate returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasAppliedDate() bool {
	if o != nil && o.AppliedDate != nil {
		return true
	}

	return false
}

// SetAppliedDate gets a reference to the given int64 and assigns it to the AppliedDate field.
func (o *WebFlagConfigApprovalRequestResponse) SetAppliedDate(v int64) {
	o.AppliedDate = &v
}

// GetAppliedByMemberId returns the AppliedByMemberId field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetAppliedByMemberId() string {
	if o == nil || o.AppliedByMemberId == nil {
		var ret string
		return ret
	}
	return *o.AppliedByMemberId
}

// GetAppliedByMemberIdOk returns a tuple with the AppliedByMemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetAppliedByMemberIdOk() (*string, bool) {
	if o == nil || o.AppliedByMemberId == nil {
		return nil, false
	}
	return o.AppliedByMemberId, true
}

// HasAppliedByMemberId returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasAppliedByMemberId() bool {
	if o != nil && o.AppliedByMemberId != nil {
		return true
	}

	return false
}

// SetAppliedByMemberId gets a reference to the given string and assigns it to the AppliedByMemberId field.
func (o *WebFlagConfigApprovalRequestResponse) SetAppliedByMemberId(v string) {
	o.AppliedByMemberId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WebFlagConfigApprovalRequestResponse) SetStatus(v string) {
	o.Status = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetInstructions() []map[string]interface{} {
	if o == nil || o.Instructions == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetInstructionsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Instructions == nil {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasInstructions() bool {
	if o != nil && o.Instructions != nil {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given []map[string]interface{} and assigns it to the Instructions field.
func (o *WebFlagConfigApprovalRequestResponse) SetInstructions(v []map[string]interface{}) {
	o.Instructions = &v
}

// GetConflicts returns the Conflicts field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetConflicts() []WebFlagConfigApprovalRequestResponseConflicts {
	if o == nil || o.Conflicts == nil {
		var ret []WebFlagConfigApprovalRequestResponseConflicts
		return ret
	}
	return *o.Conflicts
}

// GetConflictsOk returns a tuple with the Conflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetConflictsOk() (*[]WebFlagConfigApprovalRequestResponseConflicts, bool) {
	if o == nil || o.Conflicts == nil {
		return nil, false
	}
	return o.Conflicts, true
}

// HasConflicts returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasConflicts() bool {
	if o != nil && o.Conflicts != nil {
		return true
	}

	return false
}

// SetConflicts gets a reference to the given []WebFlagConfigApprovalRequestResponseConflicts and assigns it to the Conflicts field.
func (o *WebFlagConfigApprovalRequestResponse) SetConflicts(v []WebFlagConfigApprovalRequestResponseConflicts) {
	o.Conflicts = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetLinks() map[string]InlineResponse200 {
	if o == nil || o.Links == nil {
		var ret map[string]InlineResponse200
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetLinksOk() (*map[string]InlineResponse200, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]InlineResponse200 and assigns it to the Links field.
func (o *WebFlagConfigApprovalRequestResponse) SetLinks(v map[string]InlineResponse200) {
	o.Links = &v
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetExecutionDate() int64 {
	if o == nil || o.ExecutionDate == nil {
		var ret int64
		return ret
	}
	return *o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetExecutionDateOk() (*int64, bool) {
	if o == nil || o.ExecutionDate == nil {
		return nil, false
	}
	return o.ExecutionDate, true
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate != nil {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given int64 and assigns it to the ExecutionDate field.
func (o *WebFlagConfigApprovalRequestResponse) SetExecutionDate(v int64) {
	o.ExecutionDate = &v
}

// GetOperatingOnId returns the OperatingOnId field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetOperatingOnId() string {
	if o == nil || o.OperatingOnId == nil {
		var ret string
		return ret
	}
	return *o.OperatingOnId
}

// GetOperatingOnIdOk returns a tuple with the OperatingOnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetOperatingOnIdOk() (*string, bool) {
	if o == nil || o.OperatingOnId == nil {
		return nil, false
	}
	return o.OperatingOnId, true
}

// HasOperatingOnId returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasOperatingOnId() bool {
	if o != nil && o.OperatingOnId != nil {
		return true
	}

	return false
}

// SetOperatingOnId gets a reference to the given string and assigns it to the OperatingOnId field.
func (o *WebFlagConfigApprovalRequestResponse) SetOperatingOnId(v string) {
	o.OperatingOnId = &v
}

// GetIntegrationMetadata returns the IntegrationMetadata field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetIntegrationMetadata() WebIntegrationMetadata {
	if o == nil || o.IntegrationMetadata == nil {
		var ret WebIntegrationMetadata
		return ret
	}
	return *o.IntegrationMetadata
}

// GetIntegrationMetadataOk returns a tuple with the IntegrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetIntegrationMetadataOk() (*WebIntegrationMetadata, bool) {
	if o == nil || o.IntegrationMetadata == nil {
		return nil, false
	}
	return o.IntegrationMetadata, true
}

// HasIntegrationMetadata returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasIntegrationMetadata() bool {
	if o != nil && o.IntegrationMetadata != nil {
		return true
	}

	return false
}

// SetIntegrationMetadata gets a reference to the given WebIntegrationMetadata and assigns it to the IntegrationMetadata field.
func (o *WebFlagConfigApprovalRequestResponse) SetIntegrationMetadata(v WebIntegrationMetadata) {
	o.IntegrationMetadata = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *WebFlagConfigApprovalRequestResponse) GetSource() WebCopiedFromEnv {
	if o == nil || o.Source == nil {
		var ret WebCopiedFromEnv
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebFlagConfigApprovalRequestResponse) GetSourceOk() (*WebCopiedFromEnv, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *WebFlagConfigApprovalRequestResponse) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given WebCopiedFromEnv and assigns it to the Source field.
func (o *WebFlagConfigApprovalRequestResponse) SetSource(v WebCopiedFromEnv) {
	o.Source = &v
}

func (o WebFlagConfigApprovalRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.Version != nil {
		toSerialize["_version"] = o.Version
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.ServiceKind != nil {
		toSerialize["serviceKind"] = o.ServiceKind
	}
	if o.RequestorId != nil {
		toSerialize["requestorId"] = o.RequestorId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ReviewStatus != nil {
		toSerialize["reviewStatus"] = o.ReviewStatus
	}
	if o.AllReviews != nil {
		toSerialize["allReviews"] = o.AllReviews
	}
	if o.NotifyMemberIds != nil {
		toSerialize["notifyMemberIds"] = o.NotifyMemberIds
	}
	if o.AppliedDate != nil {
		toSerialize["appliedDate"] = o.AppliedDate
	}
	if o.AppliedByMemberId != nil {
		toSerialize["appliedByMemberId"] = o.AppliedByMemberId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Instructions != nil {
		toSerialize["instructions"] = o.Instructions
	}
	if o.Conflicts != nil {
		toSerialize["conflicts"] = o.Conflicts
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.ExecutionDate != nil {
		toSerialize["executionDate"] = o.ExecutionDate
	}
	if o.OperatingOnId != nil {
		toSerialize["operatingOnId"] = o.OperatingOnId
	}
	if o.IntegrationMetadata != nil {
		toSerialize["integrationMetadata"] = o.IntegrationMetadata
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableWebFlagConfigApprovalRequestResponse struct {
	value *WebFlagConfigApprovalRequestResponse
	isSet bool
}

func (v NullableWebFlagConfigApprovalRequestResponse) Get() *WebFlagConfigApprovalRequestResponse {
	return v.value
}

func (v *NullableWebFlagConfigApprovalRequestResponse) Set(val *WebFlagConfigApprovalRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebFlagConfigApprovalRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebFlagConfigApprovalRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebFlagConfigApprovalRequestResponse(val *WebFlagConfigApprovalRequestResponse) *NullableWebFlagConfigApprovalRequestResponse {
	return &NullableWebFlagConfigApprovalRequestResponse{value: val, isSet: true}
}

func (v NullableWebFlagConfigApprovalRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebFlagConfigApprovalRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


