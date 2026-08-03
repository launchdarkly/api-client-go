/*
LaunchDarkly REST API

This documentation describes LaunchDarkly's REST API. To access the complete OpenAPI spec directly, use [Get OpenAPI spec](https://launchdarkly.com/docs/api/other/get-openapi-spec).  To learn how to use LaunchDarkly using the user interface (UI) instead, read our [product documentation](https://launchdarkly.com/docs/home).  ## Authentication  LaunchDarkly's REST API uses the HTTPS protocol with a minimum TLS version of 1.2.  All REST API resources are authenticated with either [personal or service access tokens](https://launchdarkly.com/docs/home/account/api), or session cookies. Other authentication mechanisms are not supported. You can manage personal access tokens on your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page in the LaunchDarkly UI.  LaunchDarkly also has SDK keys, mobile keys, and client-side IDs that are used by our server-side SDKs, mobile SDKs, and JavaScript-based SDKs, respectively. **These keys cannot be used to access our REST API**. These keys are environment-specific, and can only perform read-only operations such as fetching feature flag settings.  | Auth mechanism                                                                                  | Allowed resources                                                                                     | Use cases                                          | | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | -------------------------------------------------- | | [Personal or service access tokens](https://launchdarkly.com/docs/home/account/api) | Can be customized on a per-token basis                                                                | Building scripts, custom integrations, data export. | | SDK keys                                                                                        | Can only access read-only resources specific to server-side SDKs. Restricted to a single environment. | Server-side SDKs                     | | Mobile keys                                                                                     | Can only access read-only resources specific to mobile SDKs, and only for flags marked available to mobile keys. Restricted to a single environment.           | Mobile SDKs                                        | | Client-side ID                                                                                  | Can only access read-only resources specific to JavaScript-based client-side SDKs, and only for flags marked available to client-side. Restricted to a single environment.           | Client-side JavaScript                             |  > #### Keep your access tokens and SDK keys private > > Access tokens should _never_ be exposed in untrusted contexts. Never put an access token in client-side JavaScript, or embed it in a mobile application. LaunchDarkly has special mobile keys that you can embed in mobile apps. If you accidentally expose an access token or SDK key, you can reset it from your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page. > > The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.  ### Authentication using request header  The preferred way to authenticate with the API is by adding an `Authorization` header containing your access token to your requests. The value of the `Authorization` header must be your access token.  Manage personal access tokens from the [**Authorization**](https://app.launchdarkly.com/settings/authorization) page.  ### Authentication using session cookie  For testing purposes, you can make API calls directly from your web browser. If you are logged in to the LaunchDarkly application, the API will use your existing session to authenticate calls.  Depending on the permissions granted as part of your [role](https://launchdarkly.com/docs/home/account/roles), you may not have permission to perform some API calls. You will receive a `401` response code in that case.  > ### Modifying the Origin header causes an error > > LaunchDarkly validates that the Origin header for any API request authenticated by a session cookie matches the expected Origin header. The expected Origin header is `https://app.launchdarkly.com`. > > If the Origin header does not match what's expected, LaunchDarkly returns an error. This error can prevent the LaunchDarkly app from working correctly. > > Any browser extension that intentionally changes the Origin header can cause this problem. For example, Cross-Origin Resource Sharing (CORS) extensions used during development can modify the Origin header and cause the app to fail. > > To prevent this error, do not modify your Origin header. > > LaunchDarkly does not require origin matching when authenticating with an access token, so this issue does not affect normal API usage.  ## Representations  All resources expect and return JSON response bodies. Error responses also send a JSON body. To learn more about the error format of the API, read [Errors](https://launchdarkly.com/docs/api#errors).  In practice this means that you always get a response with a `Content-Type` header set to `application/json`.  In addition, request bodies for `PATCH`, `POST`, and `PUT` requests must be encoded as JSON with a `Content-Type` header set to `application/json`.  ### Summary and detailed representations  When you fetch a list of resources, the response includes only the most important attributes of each resource. This is a _summary representation_ of the resource. When you fetch an individual resource, such as a single feature flag, you receive a _detailed representation_ of the resource.  The best way to find a detailed representation is to follow links. Every summary representation includes a link to its detailed representation.  ### Expanding responses  Sometimes the detailed representation of a resource does not include all of the attributes of the resource by default. If this is the case, the request method will clearly document this and describe which attributes you can include in an expanded response.  To include the additional attributes, append the `expand` request parameter to your request and add a comma-separated list of the attributes to include. For example, when you append `?expand=members,maintainers` to the [Get team](https://launchdarkly.com/docs/api/teams/get-team) endpoint, the expanded response includes both of these attributes.  ### Links and addressability  The best way to navigate the API is by following links. These are attributes in representations that link to other resources. The API always uses the same format for links:  - Links to other resources within the API are encapsulated in a `_links` object - If the resource has a corresponding link to HTML content on the site, it is stored in a special `_site` link  Each link has two attributes:  - An `href`, which contains the URL - A `type`, which describes the content type  For example, a feature resource might return the following:  ```json {   \"_links\": {     \"parent\": {       \"href\": \"/api/features\",       \"type\": \"application/json\"     },     \"self\": {       \"href\": \"/api/features/sort.order\",       \"type\": \"application/json\"     }   },   \"_site\": {     \"href\": \"/features/sort.order\",     \"type\": \"text/html\"   } } ```  From this, you can navigate to the parent collection of features by following the `parent` link, or navigate to the site page for the feature by following the `_site` link.  Collections are always represented as a JSON object with an `items` attribute containing an array of representations. Like all other representations, collections have `_links` defined at the top level.  Paginated collections include `first`, `last`, `next`, and `prev` links containing a URL with the respective set of elements in the collection.  ## Updates  Resources that accept partial updates use the `PATCH` verb. Most resources support the [JSON patch](https://launchdarkly.com/docs/api#updates-using-json-patch) format. Some resources also support the [JSON merge patch](https://launchdarkly.com/docs/api#updates-using-json-merge-patch) format, and some resources support the [semantic patch](https://launchdarkly.com/docs/api#updates-using-semantic-patch) format, which is a way to specify the modifications to perform as a set of executable instructions. Each resource supports optional [comments](https://launchdarkly.com/docs/api#updates-with-comments) that you can submit with updates. Comments appear in outgoing webhooks, the audit log, and other integrations.  When a resource supports both JSON patch and semantic patch, we document both in the request method. However, the specific request body fields and descriptions included in our documentation only match one type of patch or the other.  ### Updates using JSON patch  [JSON patch](https://datatracker.ietf.org/doc/html/rfc6902) is a way to specify the modifications to perform on a resource. JSON patch uses paths and a limited set of operations to describe how to transform the current state of the resource into a new state. JSON patch documents are always arrays, where each element contains an operation, a path to the field to update, and the new value.  For example, in this feature flag representation:  ```json {     \"name\": \"New recommendations engine\",     \"key\": \"engine.enable\",     \"description\": \"This is the description\",     ... } ``` You can change the feature flag's description with the following patch document:  ```json [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"This is the new description\" }] ```  You can specify multiple modifications to perform in a single request. You can also test that certain preconditions are met before applying the patch:  ```json [   { \"op\": \"test\", \"path\": \"/version\", \"value\": 10 },   { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" } ] ```  The above patch request tests whether the feature flag's `version` is `10`, and if so, changes the feature flag's description.  Attributes that are not editable, such as a resource's `_links`, have names that start with an underscore.  ### Updates using JSON merge patch  [JSON merge patch](https://datatracker.ietf.org/doc/html/rfc7386) is another format for specifying the modifications to perform on a resource. JSON merge patch is less expressive than JSON patch. However, in many cases it is simpler to construct a merge patch document. For example, you can change a feature flag's description with the following merge patch document:  ```json {   \"description\": \"New flag description\" } ```  ### Updates using semantic patch  Some resources support the semantic patch format. A semantic patch is a way to specify the modifications to perform on a resource as a set of executable instructions.  Semantic patch allows you to be explicit about intent using precise, custom instructions. In many cases, you can define semantic patch instructions independently of the current state of the resource. This can be useful when defining a change that may be applied at a future date.  To make a semantic patch request, you must append `domain-model=launchdarkly.semanticpatch` to your `Content-Type` header.  Here's how:  ``` Content-Type: application/json; domain-model=launchdarkly.semanticpatch ```  If you call a semantic patch resource without this header, you will receive a `400` response because your semantic patch will be interpreted as a JSON patch.  The body of a semantic patch request takes the following properties:  * `comment` (string): (Optional) A description of the update. * `environmentKey` (string): (Required for some resources only) The environment key. * `instructions` (array): (Required) A list of actions the update should perform. Each action in the list must be an object with a `kind` property that indicates the instruction. If the instruction requires parameters, you must include those parameters as additional fields in the object. The documentation for each resource that supports semantic patch includes the available instructions and any additional parameters.  For example:  ```json {   \"comment\": \"optional comment\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  Semantic patches are not applied partially; either all of the instructions are applied or none of them are. If **any** instruction is invalid, the endpoint returns an error and will not change the resource. If all instructions are valid, the request succeeds and the resources are updated if necessary, or left unchanged if they are already in the state you request.  ### Updates with comments  You can submit optional comments with `PATCH` changes.  To submit a comment along with a JSON patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"patch\": [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" }] } ```  To submit a comment along with a JSON merge patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"merge\": { \"description\": \"New flag description\" } } ```  To submit a comment along with a semantic patch, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  ## Errors  The API always returns errors in a common format. Here's an example:  ```json {   \"code\": \"invalid_request\",   \"message\": \"A feature with that key already exists\",   \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\" } ```  The `code` indicates the general class of error. The `message` is a human-readable explanation of what went wrong. The `id` is a unique identifier. Use it when you're working with LaunchDarkly Support to debug a problem with a specific API call.  ### HTTP status error response codes  | Code | Definition        | Description                                                                                       | Possible Solution                                                | | ---- | ----------------- | ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- | | 400  | Invalid request       | The request cannot be understood.                                    | Ensure JSON syntax in request body is correct.                   | | 401  | Invalid access token      | Requestor is unauthorized or does not have permission for this API call.                                                | Ensure your API access token is valid and has the appropriate permissions.                                     | | 403  | Forbidden         | Requestor does not have access to this resource.                                                | Ensure that the account member or access token has proper permissions set. | | 404  | Invalid resource identifier | The requested resource is not valid. | Ensure that the resource is correctly identified by ID or key. | | 405  | Method not allowed | The request method is not allowed on this resource. | Ensure that the HTTP verb is correct. | | 409  | Conflict          | The API request can not be completed because it conflicts with a concurrent API request. | Retry your request.                                              | | 422  | Unprocessable entity | The API request can not be completed because the update description can not be understood. | Ensure that the request body is correct for the type of patch you are using, either JSON patch or semantic patch. | 429  | Too many requests | Read [Rate limiting](https://launchdarkly.com/docs/api#rate-limiting).                                               | Wait and try again later.                                        |  ## CORS  The LaunchDarkly API supports Cross Origin Resource Sharing (CORS) for AJAX requests from any origin. If an `Origin` header is given in a request, it will be echoed as an explicitly allowed origin. Otherwise the request returns a wildcard, `Access-Control-Allow-Origin: *`. For more information on CORS, read the [CORS W3C Recommendation](http://www.w3.org/TR/cors). Example CORS headers might look like:  ```http Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, Authorization Access-Control-Allow-Methods: OPTIONS, GET, DELETE, PATCH Access-Control-Allow-Origin: * Access-Control-Max-Age: 300 ```  You can make authenticated CORS calls just as you would make same-origin calls, using either [token or session-based authentication](https://launchdarkly.com/docs/api#authentication). If you are using session authentication, you should set the `withCredentials` property for your `xhr` request to `true`. You should never expose your access tokens to untrusted entities.  ## Rate limiting  We use several rate-limiting strategies to ensure the availability of our APIs. Rate-limited calls to our APIs return a `429` status code and include headers to indicate the current rate limit status. The specific headers returned depend on the API route that was called. Limits differ based on the route, authentication mechanism, and other factors.  Each set of headers below appears only when the corresponding limit is being enforced for your call. A given route may be subject to any combination of these limits, so a response can include one, several, or none of these headers. A missing header indicates that the limit was not applied to this specific call; it does not necessarily indicate that the limit does not exist. To reduce usage before hitting a `429` status, program against whichever rate limit headers are present rather than expecting a specific header.  We do not publicly document the specific number of calls permitted by any of these limits, and these limits may change. We encourage clients to program against the specification and rely on the headers described below, rather than hardcoding the current limits.  > ### Rate limiting and SDKs > > LaunchDarkly SDKs are never rate limited and do not use the API endpoints defined here. LaunchDarkly uses a different set of approaches, including streaming/server-sent events and a global CDN, to ensure availability to the routes used by LaunchDarkly SDKs.  ### Global rate limits  Authenticated requests are subject to a global limit. This is the maximum number of calls that your account can make to the API per ten seconds. All service and personal access tokens on the account share this limit, so exceeding the limit with one access token will impact other tokens. Calls that are subject to global rate limits may return the headers below:  | Header name                    | Description                                                                      | | ------------------------------ | -------------------------------------------------------------------------------- | | `X-Ratelimit-Global-Limit`     | The maximum number of requests the account is permitted to make per ten seconds. | | `X-Ratelimit-Global-Remaining` | The number of requests remaining in the current global rate limit window.        | | `X-Ratelimit-Reset`            | The time at which the current rate limit window resets in epoch milliseconds.    |  ### Route-level rate limits  Some authenticated routes have custom rate limits. These also reset every ten seconds. Any service or personal access tokens hitting the same route share this limit, so exceeding the limit with one access token may impact other tokens. Calls that are subject to route-level rate limits return the headers below:  | Header name                   | Description                                                                                           | | ----------------------------- | ----------------------------------------------------------------------------------------------------- | | `X-Ratelimit-Route-Limit`     | The maximum number of requests to the current route permitted per ten seconds.           | | `X-Ratelimit-Route-Remaining` | The number of requests remaining for the current route in the current rate limit window. | | `X-Ratelimit-Reset`           | The time at which the current rate limit window resets in epoch milliseconds.            |  A _route_ represents a specific URL pattern and verb. For example, the [Delete environment](https://launchdarkly.com/docs/api/environments/delete-environment) endpoint is considered a single route, and each call to delete an environment counts against your route-level rate limit for that route.  ### Access token rate limits  Some calls are rate limited per access token. Unlike the global and route-level limits, this limit applies to a single service or personal access token on its own. Exceeding a limit with one access token does not affect other tokens on the account. Calls that are subject to access token rate limits return these headers:  | Header name                        | Description                                                                             | | ---------------------------------- | --------------------------------------------------------------------------------------- | | `X-Ratelimit-Auth-Token-Limit`     | The maximum number of requests the access token can make per ten seconds.               | | `X-Ratelimit-Auth-Token-Remaining` | The number of requests remaining for the access token in the current rate limit window. | | `X-Ratelimit-Auth-Token-Reset`     | The time at which the current rate limit window resets in epoch milliseconds.           |  Unlike the other rate limits, access token rate limits report their own reset time in the `X-Ratelimit-Auth-Token-Reset` header instead of in `X-Ratelimit-Reset`.  ### IP-based rate limiting  We also employ IP-based rate limiting on some API routes. If you hit an IP-based rate limit, your API response will include a `Retry-After` header indicating how long to wait before re-trying the call. Clients must wait at least `Retry-After` seconds before making additional calls to our API, and should employ jitter and backoff strategies to avoid triggering rate limits again.  ## OpenAPI (Swagger) and client libraries  We have a [complete OpenAPI (Swagger) specification](https://app.launchdarkly.com/api/v2/openapi.json) for our API.  We auto-generate multiple client libraries based on our OpenAPI specification. To learn more, visit the [collection of client libraries on GitHub](https://github.com/search?q=topic%3Alaunchdarkly-api+org%3Alaunchdarkly&type=Repositories). Alternatively, you can use the specification to generate client libraries to interact with our REST API in your language of choice. Or, you can refer to our API endpoints' documentation for guidance on how to make requests with a common HTTP library in your language of choice.  Our OpenAPI specification is supported by several API-based tools such as Postman and Insomnia. In many cases, you can directly import our specification to explore our APIs.  ## Method overriding  Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those environments, our API endpoints that use `DELETE`, `PATCH`, and `PUT` verbs are inaccessible.  To avoid this issue, our API supports the `X-HTTP-Method-Override` header, allowing clients to \"tunnel\" `DELETE`, `PATCH`, and `PUT` requests using a `POST` request.  For example, to call a `PATCH` endpoint using a `POST` request, you can include `X-HTTP-Method-Override:PATCH` as a header.  ## Beta resources  We sometimes release new API resources in **beta** status before we release them with general availability.  Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  We try to promote resources into general availability as quickly as possible. This happens after sufficient testing and when we're satisfied that we no longer need to make backwards-incompatible changes.  We mark beta resources with a \"Beta\" callout in our documentation, pictured below:  > ### This feature is in beta > > To use this feature, pass in a header including the `LD-API-Version` key with value set to `beta`. Use this header with each call. To learn more, read [Beta resources](https://launchdarkly.com/docs/api#beta-resources). > > Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  ### Using beta resources  To use a beta resource, you must include a header in the request. If you call a beta resource without this header, you receive a `403` response.  Use this header:  ``` LD-API-Version: beta ```  ## Federal and EU environments  In addition to the commercial versions, LaunchDarkly offers instances for federal agencies and those based in the European Union (EU).  ### Federal environments  The version of LaunchDarkly that is available on domains controlled by the United States government is different from the version of LaunchDarkly available to the general public. If you are an employee or contractor for a United States federal agency and use LaunchDarkly in your work, you likely use the federal instance of LaunchDarkly.  If you are working in the federal instance of LaunchDarkly, the base URI for each request is `https://app.launchdarkly.us`.  To learn more, read [LaunchDarkly in federal environments](https://launchdarkly.com/docs/home/infrastructure/federal).  ### EU environments  The version of LaunchDarkly that is available in the EU is different from the version of LaunchDarkly available to other regions. If you are based in the EU, you likely use the EU instance of LaunchDarkly. The LaunchDarkly EU instance complies with EU data residency principles, including the protection and confidentiality of EU customer information.  If you are working in the EU instance of LaunchDarkly, the base URI for each request is `https://app.eu.launchdarkly.com`.  To learn more, read [LaunchDarkly in the European Union (EU)](https://launchdarkly.com/docs/home/infrastructure/eu).  ## Versioning  We try hard to keep our REST API backwards compatible, but we occasionally have to make backwards-incompatible changes in the process of shipping new features. These breaking changes can cause unexpected behavior if you don't prepare for them accordingly.  Updates to our REST API include support for the latest features in LaunchDarkly. We also release a new version of our REST API every time we make a breaking change. We provide simultaneous support for multiple API versions so you can migrate from your current API version to a new version at your own pace.  ### Setting the API version per request  You can set the API version on a specific request by sending an `LD-API-Version` header, as shown in the example below:  ``` LD-API-Version: 20240415 ```  The header value is the version number of the API version you would like to request. The number for each version corresponds to the date the version was released in `yyyymmdd` format. In the example above the version `20240415` corresponds to April 15, 2024.  ### Setting the API version per access token  When you create an access token, you must specify a specific version of the API to use. This ensures that integrations using this token cannot be broken by version changes.  Tokens created before versioning was released have their version set to `20160426`, which is the version of the API that existed before the current versioning scheme, so that they continue working the same way they did before versioning.  If you would like to upgrade your integration to use a new API version, you can explicitly set the header described above.  > ### Best practice: Set the header for every client or integration > > We recommend that you set the API version header explicitly in any client or integration you build. > > Only rely on the access token API version during manual testing.  ### API version changelog  <table>   <tr>     <th>Version</th>     <th>Changes</th>     <th>End of life (EOL)</th>   </tr>   <tr>     <td>`20240415`</td>     <td>       <ul><li>Changed several endpoints from unpaginated to paginated. Use the `limit` and `offset` query parameters to page through the results.</li> <li>Changed the [list access tokens](https://launchdarkly.com/docs/api/access-tokens/get-tokens) endpoint: <ul><li>Response is now paginated with a default limit of `25`</li></ul></li> <li>Changed the [list account members](https://launchdarkly.com/docs/api/account-members/get-members) endpoint: <ul><li>The `accessCheck` filter is no longer available</li></ul></li> <li>Changed the [list custom roles](https://launchdarkly.com/docs/api/custom-roles/get-custom-roles) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `environments` field is now only returned if the request is filtered by environment, using the `filterEnv` query parameter</li><li>The `followerId`, `hasDataExport`, `status`, `contextKindTargeted`, and `segmentTargeted` filters are no longer available</li><li>The `compare` query parameter is no longer available</li></ul></li> <li>Changed the [list segments](https://launchdarkly.com/docs/api/segments/get-segments) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list teams](https://launchdarkly.com/docs/api/teams/get-teams) endpoint: <ul><li>The `expand` parameter no longer supports including `projects` or `roles`</li><li>In paginated results, the maximum page size is now 100</li></ul></li> <li>Changed the [get workflows](https://launchdarkly.com/docs/api/workflows/get-workflows) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `_conflicts` field in the response is no longer available</li></ul></li> </ul>     </td>     <td>Current</td>   </tr>   <tr>     <td>`20220603`</td>     <td>       <ul><li>Changed the [list projects](https://launchdarkly.com/docs/api/projects/get-projects) return value:<ul><li>Response is now paginated with a default limit of `20`.</li><li>Added support for filter and sort.</li><li>The project `environments` field is now expandable. This field is omitted by default.</li></ul></li><li>Changed the [get project](https://launchdarkly.com/docs/api/projects/get-project) return value:<ul><li>The `environments` field is now expandable. This field is omitted by default.</li></ul></li></ul>     </td>     <td>2025-04-15</td>   </tr>   <tr>     <td>`20210729`</td>     <td>       <ul><li>Changed the [create approval request](https://launchdarkly.com/docs/api/approvals/post-approval-request) return value. It now returns HTTP Status Code `201` instead of `200`.</li><li> Changed the [get user](https://launchdarkly.com/docs/api/users/get-user) return value. It now returns a user record, not a user. </li><li>Added additional optional fields to environment, segments, flags, members, and segments, including the ability to create big segments. </li><li> Added default values for flag variations when new environments are created. </li><li>Added filtering and pagination for getting flags and members, including `limit`, `number`, `filter`, and `sort` query parameters. </li><li>Added endpoints for expiring user targets for flags and segments, scheduled changes, access tokens, Relay Proxy configuration, integrations and subscriptions, and approvals. </li></ul>     </td>     <td>2023-06-03</td>   </tr>   <tr>     <td>`20191212`</td>     <td>       <ul><li>[List feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) now defaults to sending summaries of feature flag configurations, equivalent to setting the query parameter `summary=true`. Summaries omit flag targeting rules and individual user targets from the payload. </li><li> Added endpoints for flags, flag status, projects, environments, audit logs, members, users, custom roles, segments, usage, streams, events, and data export. </li></ul>     </td>     <td>2022-07-29</td>   </tr>   <tr>     <td>`20160426`</td>     <td>       <ul><li>Initial versioning of API. Tokens created before versioning have their version set to this.</li></ul>     </td>     <td>2020-12-12</td>   </tr> </table>  To learn more about how EOL is determined, read LaunchDarkly's [End of Life (EOL) Policy](https://launchdarkly.com/policies/end-of-life-policy/). 

API version: 2.0
Contact: support@launchdarkly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ldapi

import (
	"encoding/json"
	"fmt"
)

// checks if the MetricRep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricRep{}

// MetricRep struct for MetricRep
type MetricRep struct {
	// The number of experiments using this metric
	ExperimentCount *int32 `json:"experimentCount,omitempty"`
	// The number of metric groups using this metric
	MetricGroupCount *int32 `json:"metricGroupCount,omitempty"`
	// The number of active experiments using this metric
	ActiveExperimentCount *int32 `json:"activeExperimentCount,omitempty"`
	// The number of active guarded rollouts using this metric
	ActiveGuardedRolloutCount *int32 `json:"activeGuardedRolloutCount,omitempty"`
	// The ID of this metric
	Id string `json:"_id"`
	// The version ID of the metric
	VersionId string `json:"_versionId"`
	// Version of the metric
	Version *int32 `json:"_version,omitempty"`
	// A unique key to reference the metric
	Key string `json:"key"`
	// A human-friendly name for the metric
	Name string `json:"name"`
	// The kind of event the metric tracks
	Kind string `json:"kind"`
	// The number of feature flags currently attached to this metric
	AttachedFlagCount *int32 `json:"_attachedFlagCount,omitempty"`
	// The location and content type of related resources
	Links map[string]Link `json:"_links"`
	Site *Link `json:"_site,omitempty"`
	Access *Access `json:"_access,omitempty"`
	// Tags for the metric
	Tags []string `json:"tags"`
	CreationDate int64 `json:"_creationDate"`
	LastModified *Modification `json:"lastModified,omitempty"`
	// The ID of the member who maintains this metric
	MaintainerId *string `json:"maintainerId,omitempty"`
	Maintainer *MemberSummary `json:"_maintainer,omitempty"`
	// Description of the metric
	Description *string `json:"description,omitempty"`
	// The category of the metric
	Category *string `json:"category,omitempty"`
	// For custom and trace metrics, whether to track numeric changes in value against a baseline (<code>true</code>) or to track a conversion when an end user takes an action (<code>false</code>).
	IsNumeric *bool `json:"isNumeric,omitempty"`
	// For custom and trace metrics, the success criteria
	SuccessCriteria *string `json:"successCriteria,omitempty"`
	// For numeric custom and trace metrics, the unit of measure
	Unit *string `json:"unit,omitempty"`
	// For custom metrics, the event key to use in your code
	EventKey *string `json:"eventKey,omitempty"`
	// Deprecated, use <code>analysisUnits</code> instead.
	// Deprecated
	RandomizationUnits []string `json:"randomizationUnits,omitempty"`
	// An array of analysis units allowed for this metric.
	AnalysisUnits []string `json:"analysisUnits,omitempty"`
	Filters *Filter `json:"filters,omitempty"`
	// The method by which multiple unit event values are aggregated
	UnitAggregationType *string `json:"unitAggregationType,omitempty"`
	// The method for analyzing metric events
	AnalysisType *string `json:"analysisType,omitempty"`
	// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when <code>analysisType</code> is <code>percentile</code>.
	PercentileValue *int32 `json:"percentileValue,omitempty"`
	EventDefault *MetricEventDefaultRep `json:"eventDefault,omitempty"`
	DataSource MetricDataSourceRefRep `json:"dataSource"`
	LastSeen *int64 `json:"lastSeen,omitempty"`
	// Whether the metric version is archived
	Archived *bool `json:"archived,omitempty"`
	ArchivedAt *int64 `json:"archivedAt,omitempty"`
	// For click metrics, the CSS selectors
	Selector *string `json:"selector,omitempty"`
	Urls []map[string]interface{} `json:"urls,omitempty"`
	// Not yet implemented - The start of the measurement window, in milliseconds relative to the unit's first exposure to a flag variation
	WindowStartOffset *int64 `json:"windowStartOffset,omitempty"`
	// Not yet implemented - The end of the measurement window, in milliseconds relative to the unit's first exposure to a flag variation
	WindowEndOffset *int64 `json:"windowEndOffset,omitempty"`
	// Lower winsorization percentile, expressed as a percent in the open interval (0, 100). When both bounds are set, defines a two-sided clamp range. Otherwise lower-only winsorization.
	WinsorLowerPercentile *float32 `json:"winsorLowerPercentile,omitempty"`
	// Upper winsorization percentile, expressed as a percent in the open interval (0, 100). When both bounds are set, must be greater than winsorLowerPercentile.
	WinsorUpperPercentile *float32 `json:"winsorUpperPercentile,omitempty"`
	// When true, the percentile bound calculation includes imputed zeros. Only meaningful when at least one bound is set and the metric includes units that didn't send events.
	WinsorIncludeImputed *bool `json:"winsorIncludeImputed,omitempty"`
	// For trace metrics, the trace query to use for the metric.
	TraceQuery *string `json:"traceQuery,omitempty"`
	// For trace metrics, the location in the trace to use for numeric values.
	TraceValueLocation *string `json:"traceValueLocation,omitempty"`
	Denominator *MetricDenominatorRep `json:"denominator,omitempty"`
	Experiments []DependentExperimentRep `json:"experiments,omitempty"`
	// Metric groups that use this metric
	MetricGroups []DependentMetricGroupRep `json:"metricGroups,omitempty"`
	LastUsedInExperiment *DependentExperimentRep `json:"lastUsedInExperiment,omitempty"`
	LastUsedInGuardedRollout *DependentMeasuredRolloutRep `json:"lastUsedInGuardedRollout,omitempty"`
	// Whether the metric is active
	IsActive *bool `json:"isActive,omitempty"`
	// Details on the flags attached to this metric
	AttachedFeatures []FlagListingRep `json:"_attachedFeatures,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetricRep MetricRep

// NewMetricRep instantiates a new MetricRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricRep(id string, versionId string, key string, name string, kind string, links map[string]Link, tags []string, creationDate int64, dataSource MetricDataSourceRefRep) *MetricRep {
	this := MetricRep{}
	this.Id = id
	this.VersionId = versionId
	this.Key = key
	this.Name = name
	this.Kind = kind
	this.Links = links
	this.Tags = tags
	this.CreationDate = creationDate
	this.DataSource = dataSource
	return &this
}

// NewMetricRepWithDefaults instantiates a new MetricRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricRepWithDefaults() *MetricRep {
	this := MetricRep{}
	return &this
}

// GetExperimentCount returns the ExperimentCount field value if set, zero value otherwise.
func (o *MetricRep) GetExperimentCount() int32 {
	if o == nil || IsNil(o.ExperimentCount) {
		var ret int32
		return ret
	}
	return *o.ExperimentCount
}

// GetExperimentCountOk returns a tuple with the ExperimentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetExperimentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ExperimentCount) {
		return nil, false
	}
	return o.ExperimentCount, true
}

// HasExperimentCount returns a boolean if a field has been set.
func (o *MetricRep) HasExperimentCount() bool {
	if o != nil && !IsNil(o.ExperimentCount) {
		return true
	}

	return false
}

// SetExperimentCount gets a reference to the given int32 and assigns it to the ExperimentCount field.
func (o *MetricRep) SetExperimentCount(v int32) {
	o.ExperimentCount = &v
}

// GetMetricGroupCount returns the MetricGroupCount field value if set, zero value otherwise.
func (o *MetricRep) GetMetricGroupCount() int32 {
	if o == nil || IsNil(o.MetricGroupCount) {
		var ret int32
		return ret
	}
	return *o.MetricGroupCount
}

// GetMetricGroupCountOk returns a tuple with the MetricGroupCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetMetricGroupCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MetricGroupCount) {
		return nil, false
	}
	return o.MetricGroupCount, true
}

// HasMetricGroupCount returns a boolean if a field has been set.
func (o *MetricRep) HasMetricGroupCount() bool {
	if o != nil && !IsNil(o.MetricGroupCount) {
		return true
	}

	return false
}

// SetMetricGroupCount gets a reference to the given int32 and assigns it to the MetricGroupCount field.
func (o *MetricRep) SetMetricGroupCount(v int32) {
	o.MetricGroupCount = &v
}

// GetActiveExperimentCount returns the ActiveExperimentCount field value if set, zero value otherwise.
func (o *MetricRep) GetActiveExperimentCount() int32 {
	if o == nil || IsNil(o.ActiveExperimentCount) {
		var ret int32
		return ret
	}
	return *o.ActiveExperimentCount
}

// GetActiveExperimentCountOk returns a tuple with the ActiveExperimentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetActiveExperimentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveExperimentCount) {
		return nil, false
	}
	return o.ActiveExperimentCount, true
}

// HasActiveExperimentCount returns a boolean if a field has been set.
func (o *MetricRep) HasActiveExperimentCount() bool {
	if o != nil && !IsNil(o.ActiveExperimentCount) {
		return true
	}

	return false
}

// SetActiveExperimentCount gets a reference to the given int32 and assigns it to the ActiveExperimentCount field.
func (o *MetricRep) SetActiveExperimentCount(v int32) {
	o.ActiveExperimentCount = &v
}

// GetActiveGuardedRolloutCount returns the ActiveGuardedRolloutCount field value if set, zero value otherwise.
func (o *MetricRep) GetActiveGuardedRolloutCount() int32 {
	if o == nil || IsNil(o.ActiveGuardedRolloutCount) {
		var ret int32
		return ret
	}
	return *o.ActiveGuardedRolloutCount
}

// GetActiveGuardedRolloutCountOk returns a tuple with the ActiveGuardedRolloutCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetActiveGuardedRolloutCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveGuardedRolloutCount) {
		return nil, false
	}
	return o.ActiveGuardedRolloutCount, true
}

// HasActiveGuardedRolloutCount returns a boolean if a field has been set.
func (o *MetricRep) HasActiveGuardedRolloutCount() bool {
	if o != nil && !IsNil(o.ActiveGuardedRolloutCount) {
		return true
	}

	return false
}

// SetActiveGuardedRolloutCount gets a reference to the given int32 and assigns it to the ActiveGuardedRolloutCount field.
func (o *MetricRep) SetActiveGuardedRolloutCount(v int32) {
	o.ActiveGuardedRolloutCount = &v
}

// GetId returns the Id field value
func (o *MetricRep) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetricRep) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MetricRep) SetId(v string) {
	o.Id = v
}

// GetVersionId returns the VersionId field value
func (o *MetricRep) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *MetricRep) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *MetricRep) SetVersionId(v string) {
	o.VersionId = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MetricRep) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MetricRep) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *MetricRep) SetVersion(v int32) {
	o.Version = &v
}

// GetKey returns the Key field value
func (o *MetricRep) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MetricRep) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MetricRep) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *MetricRep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricRep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetricRep) SetName(v string) {
	o.Name = v
}

// GetKind returns the Kind field value
func (o *MetricRep) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *MetricRep) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *MetricRep) SetKind(v string) {
	o.Kind = v
}

// GetAttachedFlagCount returns the AttachedFlagCount field value if set, zero value otherwise.
func (o *MetricRep) GetAttachedFlagCount() int32 {
	if o == nil || IsNil(o.AttachedFlagCount) {
		var ret int32
		return ret
	}
	return *o.AttachedFlagCount
}

// GetAttachedFlagCountOk returns a tuple with the AttachedFlagCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetAttachedFlagCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AttachedFlagCount) {
		return nil, false
	}
	return o.AttachedFlagCount, true
}

// HasAttachedFlagCount returns a boolean if a field has been set.
func (o *MetricRep) HasAttachedFlagCount() bool {
	if o != nil && !IsNil(o.AttachedFlagCount) {
		return true
	}

	return false
}

// SetAttachedFlagCount gets a reference to the given int32 and assigns it to the AttachedFlagCount field.
func (o *MetricRep) SetAttachedFlagCount(v int32) {
	o.AttachedFlagCount = &v
}

// GetLinks returns the Links field value
func (o *MetricRep) GetLinks() map[string]Link {
	if o == nil {
		var ret map[string]Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *MetricRep) GetLinksOk() (*map[string]Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *MetricRep) SetLinks(v map[string]Link) {
	o.Links = v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *MetricRep) GetSite() Link {
	if o == nil || IsNil(o.Site) {
		var ret Link
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetSiteOk() (*Link, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *MetricRep) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given Link and assigns it to the Site field.
func (o *MetricRep) SetSite(v Link) {
	o.Site = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *MetricRep) GetAccess() Access {
	if o == nil || IsNil(o.Access) {
		var ret Access
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetAccessOk() (*Access, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *MetricRep) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given Access and assigns it to the Access field.
func (o *MetricRep) SetAccess(v Access) {
	o.Access = &v
}

// GetTags returns the Tags field value
func (o *MetricRep) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *MetricRep) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *MetricRep) SetTags(v []string) {
	o.Tags = v
}

// GetCreationDate returns the CreationDate field value
func (o *MetricRep) GetCreationDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value
// and a boolean to check if the value has been set.
func (o *MetricRep) GetCreationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationDate, true
}

// SetCreationDate sets field value
func (o *MetricRep) SetCreationDate(v int64) {
	o.CreationDate = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *MetricRep) GetLastModified() Modification {
	if o == nil || IsNil(o.LastModified) {
		var ret Modification
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetLastModifiedOk() (*Modification, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *MetricRep) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given Modification and assigns it to the LastModified field.
func (o *MetricRep) SetLastModified(v Modification) {
	o.LastModified = &v
}

// GetMaintainerId returns the MaintainerId field value if set, zero value otherwise.
func (o *MetricRep) GetMaintainerId() string {
	if o == nil || IsNil(o.MaintainerId) {
		var ret string
		return ret
	}
	return *o.MaintainerId
}

// GetMaintainerIdOk returns a tuple with the MaintainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetMaintainerIdOk() (*string, bool) {
	if o == nil || IsNil(o.MaintainerId) {
		return nil, false
	}
	return o.MaintainerId, true
}

// HasMaintainerId returns a boolean if a field has been set.
func (o *MetricRep) HasMaintainerId() bool {
	if o != nil && !IsNil(o.MaintainerId) {
		return true
	}

	return false
}

// SetMaintainerId gets a reference to the given string and assigns it to the MaintainerId field.
func (o *MetricRep) SetMaintainerId(v string) {
	o.MaintainerId = &v
}

// GetMaintainer returns the Maintainer field value if set, zero value otherwise.
func (o *MetricRep) GetMaintainer() MemberSummary {
	if o == nil || IsNil(o.Maintainer) {
		var ret MemberSummary
		return ret
	}
	return *o.Maintainer
}

// GetMaintainerOk returns a tuple with the Maintainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetMaintainerOk() (*MemberSummary, bool) {
	if o == nil || IsNil(o.Maintainer) {
		return nil, false
	}
	return o.Maintainer, true
}

// HasMaintainer returns a boolean if a field has been set.
func (o *MetricRep) HasMaintainer() bool {
	if o != nil && !IsNil(o.Maintainer) {
		return true
	}

	return false
}

// SetMaintainer gets a reference to the given MemberSummary and assigns it to the Maintainer field.
func (o *MetricRep) SetMaintainer(v MemberSummary) {
	o.Maintainer = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetricRep) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricRep) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetricRep) SetDescription(v string) {
	o.Description = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *MetricRep) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MetricRep) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *MetricRep) SetCategory(v string) {
	o.Category = &v
}

// GetIsNumeric returns the IsNumeric field value if set, zero value otherwise.
func (o *MetricRep) GetIsNumeric() bool {
	if o == nil || IsNil(o.IsNumeric) {
		var ret bool
		return ret
	}
	return *o.IsNumeric
}

// GetIsNumericOk returns a tuple with the IsNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetIsNumericOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNumeric) {
		return nil, false
	}
	return o.IsNumeric, true
}

// HasIsNumeric returns a boolean if a field has been set.
func (o *MetricRep) HasIsNumeric() bool {
	if o != nil && !IsNil(o.IsNumeric) {
		return true
	}

	return false
}

// SetIsNumeric gets a reference to the given bool and assigns it to the IsNumeric field.
func (o *MetricRep) SetIsNumeric(v bool) {
	o.IsNumeric = &v
}

// GetSuccessCriteria returns the SuccessCriteria field value if set, zero value otherwise.
func (o *MetricRep) GetSuccessCriteria() string {
	if o == nil || IsNil(o.SuccessCriteria) {
		var ret string
		return ret
	}
	return *o.SuccessCriteria
}

// GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetSuccessCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessCriteria) {
		return nil, false
	}
	return o.SuccessCriteria, true
}

// HasSuccessCriteria returns a boolean if a field has been set.
func (o *MetricRep) HasSuccessCriteria() bool {
	if o != nil && !IsNil(o.SuccessCriteria) {
		return true
	}

	return false
}

// SetSuccessCriteria gets a reference to the given string and assigns it to the SuccessCriteria field.
func (o *MetricRep) SetSuccessCriteria(v string) {
	o.SuccessCriteria = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MetricRep) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricRep) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *MetricRep) SetUnit(v string) {
	o.Unit = &v
}

// GetEventKey returns the EventKey field value if set, zero value otherwise.
func (o *MetricRep) GetEventKey() string {
	if o == nil || IsNil(o.EventKey) {
		var ret string
		return ret
	}
	return *o.EventKey
}

// GetEventKeyOk returns a tuple with the EventKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetEventKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EventKey) {
		return nil, false
	}
	return o.EventKey, true
}

// HasEventKey returns a boolean if a field has been set.
func (o *MetricRep) HasEventKey() bool {
	if o != nil && !IsNil(o.EventKey) {
		return true
	}

	return false
}

// SetEventKey gets a reference to the given string and assigns it to the EventKey field.
func (o *MetricRep) SetEventKey(v string) {
	o.EventKey = &v
}

// GetRandomizationUnits returns the RandomizationUnits field value if set, zero value otherwise.
// Deprecated
func (o *MetricRep) GetRandomizationUnits() []string {
	if o == nil || IsNil(o.RandomizationUnits) {
		var ret []string
		return ret
	}
	return o.RandomizationUnits
}

// GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MetricRep) GetRandomizationUnitsOk() ([]string, bool) {
	if o == nil || IsNil(o.RandomizationUnits) {
		return nil, false
	}
	return o.RandomizationUnits, true
}

// HasRandomizationUnits returns a boolean if a field has been set.
func (o *MetricRep) HasRandomizationUnits() bool {
	if o != nil && !IsNil(o.RandomizationUnits) {
		return true
	}

	return false
}

// SetRandomizationUnits gets a reference to the given []string and assigns it to the RandomizationUnits field.
// Deprecated
func (o *MetricRep) SetRandomizationUnits(v []string) {
	o.RandomizationUnits = v
}

// GetAnalysisUnits returns the AnalysisUnits field value if set, zero value otherwise.
func (o *MetricRep) GetAnalysisUnits() []string {
	if o == nil || IsNil(o.AnalysisUnits) {
		var ret []string
		return ret
	}
	return o.AnalysisUnits
}

// GetAnalysisUnitsOk returns a tuple with the AnalysisUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetAnalysisUnitsOk() ([]string, bool) {
	if o == nil || IsNil(o.AnalysisUnits) {
		return nil, false
	}
	return o.AnalysisUnits, true
}

// HasAnalysisUnits returns a boolean if a field has been set.
func (o *MetricRep) HasAnalysisUnits() bool {
	if o != nil && !IsNil(o.AnalysisUnits) {
		return true
	}

	return false
}

// SetAnalysisUnits gets a reference to the given []string and assigns it to the AnalysisUnits field.
func (o *MetricRep) SetAnalysisUnits(v []string) {
	o.AnalysisUnits = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *MetricRep) GetFilters() Filter {
	if o == nil || IsNil(o.Filters) {
		var ret Filter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetFiltersOk() (*Filter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *MetricRep) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given Filter and assigns it to the Filters field.
func (o *MetricRep) SetFilters(v Filter) {
	o.Filters = &v
}

// GetUnitAggregationType returns the UnitAggregationType field value if set, zero value otherwise.
func (o *MetricRep) GetUnitAggregationType() string {
	if o == nil || IsNil(o.UnitAggregationType) {
		var ret string
		return ret
	}
	return *o.UnitAggregationType
}

// GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetUnitAggregationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UnitAggregationType) {
		return nil, false
	}
	return o.UnitAggregationType, true
}

// HasUnitAggregationType returns a boolean if a field has been set.
func (o *MetricRep) HasUnitAggregationType() bool {
	if o != nil && !IsNil(o.UnitAggregationType) {
		return true
	}

	return false
}

// SetUnitAggregationType gets a reference to the given string and assigns it to the UnitAggregationType field.
func (o *MetricRep) SetUnitAggregationType(v string) {
	o.UnitAggregationType = &v
}

// GetAnalysisType returns the AnalysisType field value if set, zero value otherwise.
func (o *MetricRep) GetAnalysisType() string {
	if o == nil || IsNil(o.AnalysisType) {
		var ret string
		return ret
	}
	return *o.AnalysisType
}

// GetAnalysisTypeOk returns a tuple with the AnalysisType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetAnalysisTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AnalysisType) {
		return nil, false
	}
	return o.AnalysisType, true
}

// HasAnalysisType returns a boolean if a field has been set.
func (o *MetricRep) HasAnalysisType() bool {
	if o != nil && !IsNil(o.AnalysisType) {
		return true
	}

	return false
}

// SetAnalysisType gets a reference to the given string and assigns it to the AnalysisType field.
func (o *MetricRep) SetAnalysisType(v string) {
	o.AnalysisType = &v
}

// GetPercentileValue returns the PercentileValue field value if set, zero value otherwise.
func (o *MetricRep) GetPercentileValue() int32 {
	if o == nil || IsNil(o.PercentileValue) {
		var ret int32
		return ret
	}
	return *o.PercentileValue
}

// GetPercentileValueOk returns a tuple with the PercentileValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetPercentileValueOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentileValue) {
		return nil, false
	}
	return o.PercentileValue, true
}

// HasPercentileValue returns a boolean if a field has been set.
func (o *MetricRep) HasPercentileValue() bool {
	if o != nil && !IsNil(o.PercentileValue) {
		return true
	}

	return false
}

// SetPercentileValue gets a reference to the given int32 and assigns it to the PercentileValue field.
func (o *MetricRep) SetPercentileValue(v int32) {
	o.PercentileValue = &v
}

// GetEventDefault returns the EventDefault field value if set, zero value otherwise.
func (o *MetricRep) GetEventDefault() MetricEventDefaultRep {
	if o == nil || IsNil(o.EventDefault) {
		var ret MetricEventDefaultRep
		return ret
	}
	return *o.EventDefault
}

// GetEventDefaultOk returns a tuple with the EventDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetEventDefaultOk() (*MetricEventDefaultRep, bool) {
	if o == nil || IsNil(o.EventDefault) {
		return nil, false
	}
	return o.EventDefault, true
}

// HasEventDefault returns a boolean if a field has been set.
func (o *MetricRep) HasEventDefault() bool {
	if o != nil && !IsNil(o.EventDefault) {
		return true
	}

	return false
}

// SetEventDefault gets a reference to the given MetricEventDefaultRep and assigns it to the EventDefault field.
func (o *MetricRep) SetEventDefault(v MetricEventDefaultRep) {
	o.EventDefault = &v
}

// GetDataSource returns the DataSource field value
func (o *MetricRep) GetDataSource() MetricDataSourceRefRep {
	if o == nil {
		var ret MetricDataSourceRefRep
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MetricRep) GetDataSourceOk() (*MetricDataSourceRefRep, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *MetricRep) SetDataSource(v MetricDataSourceRefRep) {
	o.DataSource = v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *MetricRep) GetLastSeen() int64 {
	if o == nil || IsNil(o.LastSeen) {
		var ret int64
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetLastSeenOk() (*int64, bool) {
	if o == nil || IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *MetricRep) HasLastSeen() bool {
	if o != nil && !IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given int64 and assigns it to the LastSeen field.
func (o *MetricRep) SetLastSeen(v int64) {
	o.LastSeen = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *MetricRep) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *MetricRep) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *MetricRep) SetArchived(v bool) {
	o.Archived = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *MetricRep) GetArchivedAt() int64 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret int64
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetArchivedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *MetricRep) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given int64 and assigns it to the ArchivedAt field.
func (o *MetricRep) SetArchivedAt(v int64) {
	o.ArchivedAt = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *MetricRep) GetSelector() string {
	if o == nil || IsNil(o.Selector) {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *MetricRep) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *MetricRep) SetSelector(v string) {
	o.Selector = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *MetricRep) GetUrls() []map[string]interface{} {
	if o == nil || IsNil(o.Urls) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetUrlsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *MetricRep) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []map[string]interface{} and assigns it to the Urls field.
func (o *MetricRep) SetUrls(v []map[string]interface{}) {
	o.Urls = v
}

// GetWindowStartOffset returns the WindowStartOffset field value if set, zero value otherwise.
func (o *MetricRep) GetWindowStartOffset() int64 {
	if o == nil || IsNil(o.WindowStartOffset) {
		var ret int64
		return ret
	}
	return *o.WindowStartOffset
}

// GetWindowStartOffsetOk returns a tuple with the WindowStartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetWindowStartOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.WindowStartOffset) {
		return nil, false
	}
	return o.WindowStartOffset, true
}

// HasWindowStartOffset returns a boolean if a field has been set.
func (o *MetricRep) HasWindowStartOffset() bool {
	if o != nil && !IsNil(o.WindowStartOffset) {
		return true
	}

	return false
}

// SetWindowStartOffset gets a reference to the given int64 and assigns it to the WindowStartOffset field.
func (o *MetricRep) SetWindowStartOffset(v int64) {
	o.WindowStartOffset = &v
}

// GetWindowEndOffset returns the WindowEndOffset field value if set, zero value otherwise.
func (o *MetricRep) GetWindowEndOffset() int64 {
	if o == nil || IsNil(o.WindowEndOffset) {
		var ret int64
		return ret
	}
	return *o.WindowEndOffset
}

// GetWindowEndOffsetOk returns a tuple with the WindowEndOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetWindowEndOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.WindowEndOffset) {
		return nil, false
	}
	return o.WindowEndOffset, true
}

// HasWindowEndOffset returns a boolean if a field has been set.
func (o *MetricRep) HasWindowEndOffset() bool {
	if o != nil && !IsNil(o.WindowEndOffset) {
		return true
	}

	return false
}

// SetWindowEndOffset gets a reference to the given int64 and assigns it to the WindowEndOffset field.
func (o *MetricRep) SetWindowEndOffset(v int64) {
	o.WindowEndOffset = &v
}

// GetWinsorLowerPercentile returns the WinsorLowerPercentile field value if set, zero value otherwise.
func (o *MetricRep) GetWinsorLowerPercentile() float32 {
	if o == nil || IsNil(o.WinsorLowerPercentile) {
		var ret float32
		return ret
	}
	return *o.WinsorLowerPercentile
}

// GetWinsorLowerPercentileOk returns a tuple with the WinsorLowerPercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetWinsorLowerPercentileOk() (*float32, bool) {
	if o == nil || IsNil(o.WinsorLowerPercentile) {
		return nil, false
	}
	return o.WinsorLowerPercentile, true
}

// HasWinsorLowerPercentile returns a boolean if a field has been set.
func (o *MetricRep) HasWinsorLowerPercentile() bool {
	if o != nil && !IsNil(o.WinsorLowerPercentile) {
		return true
	}

	return false
}

// SetWinsorLowerPercentile gets a reference to the given float32 and assigns it to the WinsorLowerPercentile field.
func (o *MetricRep) SetWinsorLowerPercentile(v float32) {
	o.WinsorLowerPercentile = &v
}

// GetWinsorUpperPercentile returns the WinsorUpperPercentile field value if set, zero value otherwise.
func (o *MetricRep) GetWinsorUpperPercentile() float32 {
	if o == nil || IsNil(o.WinsorUpperPercentile) {
		var ret float32
		return ret
	}
	return *o.WinsorUpperPercentile
}

// GetWinsorUpperPercentileOk returns a tuple with the WinsorUpperPercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetWinsorUpperPercentileOk() (*float32, bool) {
	if o == nil || IsNil(o.WinsorUpperPercentile) {
		return nil, false
	}
	return o.WinsorUpperPercentile, true
}

// HasWinsorUpperPercentile returns a boolean if a field has been set.
func (o *MetricRep) HasWinsorUpperPercentile() bool {
	if o != nil && !IsNil(o.WinsorUpperPercentile) {
		return true
	}

	return false
}

// SetWinsorUpperPercentile gets a reference to the given float32 and assigns it to the WinsorUpperPercentile field.
func (o *MetricRep) SetWinsorUpperPercentile(v float32) {
	o.WinsorUpperPercentile = &v
}

// GetWinsorIncludeImputed returns the WinsorIncludeImputed field value if set, zero value otherwise.
func (o *MetricRep) GetWinsorIncludeImputed() bool {
	if o == nil || IsNil(o.WinsorIncludeImputed) {
		var ret bool
		return ret
	}
	return *o.WinsorIncludeImputed
}

// GetWinsorIncludeImputedOk returns a tuple with the WinsorIncludeImputed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetWinsorIncludeImputedOk() (*bool, bool) {
	if o == nil || IsNil(o.WinsorIncludeImputed) {
		return nil, false
	}
	return o.WinsorIncludeImputed, true
}

// HasWinsorIncludeImputed returns a boolean if a field has been set.
func (o *MetricRep) HasWinsorIncludeImputed() bool {
	if o != nil && !IsNil(o.WinsorIncludeImputed) {
		return true
	}

	return false
}

// SetWinsorIncludeImputed gets a reference to the given bool and assigns it to the WinsorIncludeImputed field.
func (o *MetricRep) SetWinsorIncludeImputed(v bool) {
	o.WinsorIncludeImputed = &v
}

// GetTraceQuery returns the TraceQuery field value if set, zero value otherwise.
func (o *MetricRep) GetTraceQuery() string {
	if o == nil || IsNil(o.TraceQuery) {
		var ret string
		return ret
	}
	return *o.TraceQuery
}

// GetTraceQueryOk returns a tuple with the TraceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetTraceQueryOk() (*string, bool) {
	if o == nil || IsNil(o.TraceQuery) {
		return nil, false
	}
	return o.TraceQuery, true
}

// HasTraceQuery returns a boolean if a field has been set.
func (o *MetricRep) HasTraceQuery() bool {
	if o != nil && !IsNil(o.TraceQuery) {
		return true
	}

	return false
}

// SetTraceQuery gets a reference to the given string and assigns it to the TraceQuery field.
func (o *MetricRep) SetTraceQuery(v string) {
	o.TraceQuery = &v
}

// GetTraceValueLocation returns the TraceValueLocation field value if set, zero value otherwise.
func (o *MetricRep) GetTraceValueLocation() string {
	if o == nil || IsNil(o.TraceValueLocation) {
		var ret string
		return ret
	}
	return *o.TraceValueLocation
}

// GetTraceValueLocationOk returns a tuple with the TraceValueLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetTraceValueLocationOk() (*string, bool) {
	if o == nil || IsNil(o.TraceValueLocation) {
		return nil, false
	}
	return o.TraceValueLocation, true
}

// HasTraceValueLocation returns a boolean if a field has been set.
func (o *MetricRep) HasTraceValueLocation() bool {
	if o != nil && !IsNil(o.TraceValueLocation) {
		return true
	}

	return false
}

// SetTraceValueLocation gets a reference to the given string and assigns it to the TraceValueLocation field.
func (o *MetricRep) SetTraceValueLocation(v string) {
	o.TraceValueLocation = &v
}

// GetDenominator returns the Denominator field value if set, zero value otherwise.
func (o *MetricRep) GetDenominator() MetricDenominatorRep {
	if o == nil || IsNil(o.Denominator) {
		var ret MetricDenominatorRep
		return ret
	}
	return *o.Denominator
}

// GetDenominatorOk returns a tuple with the Denominator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetDenominatorOk() (*MetricDenominatorRep, bool) {
	if o == nil || IsNil(o.Denominator) {
		return nil, false
	}
	return o.Denominator, true
}

// HasDenominator returns a boolean if a field has been set.
func (o *MetricRep) HasDenominator() bool {
	if o != nil && !IsNil(o.Denominator) {
		return true
	}

	return false
}

// SetDenominator gets a reference to the given MetricDenominatorRep and assigns it to the Denominator field.
func (o *MetricRep) SetDenominator(v MetricDenominatorRep) {
	o.Denominator = &v
}

// GetExperiments returns the Experiments field value if set, zero value otherwise.
func (o *MetricRep) GetExperiments() []DependentExperimentRep {
	if o == nil || IsNil(o.Experiments) {
		var ret []DependentExperimentRep
		return ret
	}
	return o.Experiments
}

// GetExperimentsOk returns a tuple with the Experiments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetExperimentsOk() ([]DependentExperimentRep, bool) {
	if o == nil || IsNil(o.Experiments) {
		return nil, false
	}
	return o.Experiments, true
}

// HasExperiments returns a boolean if a field has been set.
func (o *MetricRep) HasExperiments() bool {
	if o != nil && !IsNil(o.Experiments) {
		return true
	}

	return false
}

// SetExperiments gets a reference to the given []DependentExperimentRep and assigns it to the Experiments field.
func (o *MetricRep) SetExperiments(v []DependentExperimentRep) {
	o.Experiments = v
}

// GetMetricGroups returns the MetricGroups field value if set, zero value otherwise.
func (o *MetricRep) GetMetricGroups() []DependentMetricGroupRep {
	if o == nil || IsNil(o.MetricGroups) {
		var ret []DependentMetricGroupRep
		return ret
	}
	return o.MetricGroups
}

// GetMetricGroupsOk returns a tuple with the MetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetMetricGroupsOk() ([]DependentMetricGroupRep, bool) {
	if o == nil || IsNil(o.MetricGroups) {
		return nil, false
	}
	return o.MetricGroups, true
}

// HasMetricGroups returns a boolean if a field has been set.
func (o *MetricRep) HasMetricGroups() bool {
	if o != nil && !IsNil(o.MetricGroups) {
		return true
	}

	return false
}

// SetMetricGroups gets a reference to the given []DependentMetricGroupRep and assigns it to the MetricGroups field.
func (o *MetricRep) SetMetricGroups(v []DependentMetricGroupRep) {
	o.MetricGroups = v
}

// GetLastUsedInExperiment returns the LastUsedInExperiment field value if set, zero value otherwise.
func (o *MetricRep) GetLastUsedInExperiment() DependentExperimentRep {
	if o == nil || IsNil(o.LastUsedInExperiment) {
		var ret DependentExperimentRep
		return ret
	}
	return *o.LastUsedInExperiment
}

// GetLastUsedInExperimentOk returns a tuple with the LastUsedInExperiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetLastUsedInExperimentOk() (*DependentExperimentRep, bool) {
	if o == nil || IsNil(o.LastUsedInExperiment) {
		return nil, false
	}
	return o.LastUsedInExperiment, true
}

// HasLastUsedInExperiment returns a boolean if a field has been set.
func (o *MetricRep) HasLastUsedInExperiment() bool {
	if o != nil && !IsNil(o.LastUsedInExperiment) {
		return true
	}

	return false
}

// SetLastUsedInExperiment gets a reference to the given DependentExperimentRep and assigns it to the LastUsedInExperiment field.
func (o *MetricRep) SetLastUsedInExperiment(v DependentExperimentRep) {
	o.LastUsedInExperiment = &v
}

// GetLastUsedInGuardedRollout returns the LastUsedInGuardedRollout field value if set, zero value otherwise.
func (o *MetricRep) GetLastUsedInGuardedRollout() DependentMeasuredRolloutRep {
	if o == nil || IsNil(o.LastUsedInGuardedRollout) {
		var ret DependentMeasuredRolloutRep
		return ret
	}
	return *o.LastUsedInGuardedRollout
}

// GetLastUsedInGuardedRolloutOk returns a tuple with the LastUsedInGuardedRollout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetLastUsedInGuardedRolloutOk() (*DependentMeasuredRolloutRep, bool) {
	if o == nil || IsNil(o.LastUsedInGuardedRollout) {
		return nil, false
	}
	return o.LastUsedInGuardedRollout, true
}

// HasLastUsedInGuardedRollout returns a boolean if a field has been set.
func (o *MetricRep) HasLastUsedInGuardedRollout() bool {
	if o != nil && !IsNil(o.LastUsedInGuardedRollout) {
		return true
	}

	return false
}

// SetLastUsedInGuardedRollout gets a reference to the given DependentMeasuredRolloutRep and assigns it to the LastUsedInGuardedRollout field.
func (o *MetricRep) SetLastUsedInGuardedRollout(v DependentMeasuredRolloutRep) {
	o.LastUsedInGuardedRollout = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *MetricRep) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *MetricRep) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *MetricRep) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAttachedFeatures returns the AttachedFeatures field value if set, zero value otherwise.
func (o *MetricRep) GetAttachedFeatures() []FlagListingRep {
	if o == nil || IsNil(o.AttachedFeatures) {
		var ret []FlagListingRep
		return ret
	}
	return o.AttachedFeatures
}

// GetAttachedFeaturesOk returns a tuple with the AttachedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRep) GetAttachedFeaturesOk() ([]FlagListingRep, bool) {
	if o == nil || IsNil(o.AttachedFeatures) {
		return nil, false
	}
	return o.AttachedFeatures, true
}

// HasAttachedFeatures returns a boolean if a field has been set.
func (o *MetricRep) HasAttachedFeatures() bool {
	if o != nil && !IsNil(o.AttachedFeatures) {
		return true
	}

	return false
}

// SetAttachedFeatures gets a reference to the given []FlagListingRep and assigns it to the AttachedFeatures field.
func (o *MetricRep) SetAttachedFeatures(v []FlagListingRep) {
	o.AttachedFeatures = v
}

func (o MetricRep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricRep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExperimentCount) {
		toSerialize["experimentCount"] = o.ExperimentCount
	}
	if !IsNil(o.MetricGroupCount) {
		toSerialize["metricGroupCount"] = o.MetricGroupCount
	}
	if !IsNil(o.ActiveExperimentCount) {
		toSerialize["activeExperimentCount"] = o.ActiveExperimentCount
	}
	if !IsNil(o.ActiveGuardedRolloutCount) {
		toSerialize["activeGuardedRolloutCount"] = o.ActiveGuardedRolloutCount
	}
	toSerialize["_id"] = o.Id
	toSerialize["_versionId"] = o.VersionId
	if !IsNil(o.Version) {
		toSerialize["_version"] = o.Version
	}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["kind"] = o.Kind
	if !IsNil(o.AttachedFlagCount) {
		toSerialize["_attachedFlagCount"] = o.AttachedFlagCount
	}
	toSerialize["_links"] = o.Links
	if !IsNil(o.Site) {
		toSerialize["_site"] = o.Site
	}
	if !IsNil(o.Access) {
		toSerialize["_access"] = o.Access
	}
	toSerialize["tags"] = o.Tags
	toSerialize["_creationDate"] = o.CreationDate
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.MaintainerId) {
		toSerialize["maintainerId"] = o.MaintainerId
	}
	if !IsNil(o.Maintainer) {
		toSerialize["_maintainer"] = o.Maintainer
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.IsNumeric) {
		toSerialize["isNumeric"] = o.IsNumeric
	}
	if !IsNil(o.SuccessCriteria) {
		toSerialize["successCriteria"] = o.SuccessCriteria
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.EventKey) {
		toSerialize["eventKey"] = o.EventKey
	}
	if !IsNil(o.RandomizationUnits) {
		toSerialize["randomizationUnits"] = o.RandomizationUnits
	}
	if !IsNil(o.AnalysisUnits) {
		toSerialize["analysisUnits"] = o.AnalysisUnits
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.UnitAggregationType) {
		toSerialize["unitAggregationType"] = o.UnitAggregationType
	}
	if !IsNil(o.AnalysisType) {
		toSerialize["analysisType"] = o.AnalysisType
	}
	if !IsNil(o.PercentileValue) {
		toSerialize["percentileValue"] = o.PercentileValue
	}
	if !IsNil(o.EventDefault) {
		toSerialize["eventDefault"] = o.EventDefault
	}
	toSerialize["dataSource"] = o.DataSource
	if !IsNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	if !IsNil(o.WindowStartOffset) {
		toSerialize["windowStartOffset"] = o.WindowStartOffset
	}
	if !IsNil(o.WindowEndOffset) {
		toSerialize["windowEndOffset"] = o.WindowEndOffset
	}
	if !IsNil(o.WinsorLowerPercentile) {
		toSerialize["winsorLowerPercentile"] = o.WinsorLowerPercentile
	}
	if !IsNil(o.WinsorUpperPercentile) {
		toSerialize["winsorUpperPercentile"] = o.WinsorUpperPercentile
	}
	if !IsNil(o.WinsorIncludeImputed) {
		toSerialize["winsorIncludeImputed"] = o.WinsorIncludeImputed
	}
	if !IsNil(o.TraceQuery) {
		toSerialize["traceQuery"] = o.TraceQuery
	}
	if !IsNil(o.TraceValueLocation) {
		toSerialize["traceValueLocation"] = o.TraceValueLocation
	}
	if !IsNil(o.Denominator) {
		toSerialize["denominator"] = o.Denominator
	}
	if !IsNil(o.Experiments) {
		toSerialize["experiments"] = o.Experiments
	}
	if !IsNil(o.MetricGroups) {
		toSerialize["metricGroups"] = o.MetricGroups
	}
	if !IsNil(o.LastUsedInExperiment) {
		toSerialize["lastUsedInExperiment"] = o.LastUsedInExperiment
	}
	if !IsNil(o.LastUsedInGuardedRollout) {
		toSerialize["lastUsedInGuardedRollout"] = o.LastUsedInGuardedRollout
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.AttachedFeatures) {
		toSerialize["_attachedFeatures"] = o.AttachedFeatures
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetricRep) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_versionId",
		"key",
		"name",
		"kind",
		"_links",
		"tags",
		"_creationDate",
		"dataSource",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMetricRep := _MetricRep{}

	err = json.Unmarshal(data, &varMetricRep)

	if err != nil {
		return err
	}

	*o = MetricRep(varMetricRep)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "experimentCount")
		delete(additionalProperties, "metricGroupCount")
		delete(additionalProperties, "activeExperimentCount")
		delete(additionalProperties, "activeGuardedRolloutCount")
		delete(additionalProperties, "_id")
		delete(additionalProperties, "_versionId")
		delete(additionalProperties, "_version")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "_attachedFlagCount")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "_site")
		delete(additionalProperties, "_access")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "_creationDate")
		delete(additionalProperties, "lastModified")
		delete(additionalProperties, "maintainerId")
		delete(additionalProperties, "_maintainer")
		delete(additionalProperties, "description")
		delete(additionalProperties, "category")
		delete(additionalProperties, "isNumeric")
		delete(additionalProperties, "successCriteria")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "eventKey")
		delete(additionalProperties, "randomizationUnits")
		delete(additionalProperties, "analysisUnits")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "unitAggregationType")
		delete(additionalProperties, "analysisType")
		delete(additionalProperties, "percentileValue")
		delete(additionalProperties, "eventDefault")
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "lastSeen")
		delete(additionalProperties, "archived")
		delete(additionalProperties, "archivedAt")
		delete(additionalProperties, "selector")
		delete(additionalProperties, "urls")
		delete(additionalProperties, "windowStartOffset")
		delete(additionalProperties, "windowEndOffset")
		delete(additionalProperties, "winsorLowerPercentile")
		delete(additionalProperties, "winsorUpperPercentile")
		delete(additionalProperties, "winsorIncludeImputed")
		delete(additionalProperties, "traceQuery")
		delete(additionalProperties, "traceValueLocation")
		delete(additionalProperties, "denominator")
		delete(additionalProperties, "experiments")
		delete(additionalProperties, "metricGroups")
		delete(additionalProperties, "lastUsedInExperiment")
		delete(additionalProperties, "lastUsedInGuardedRollout")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "_attachedFeatures")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetricRep struct {
	value *MetricRep
	isSet bool
}

func (v NullableMetricRep) Get() *MetricRep {
	return v.value
}

func (v *NullableMetricRep) Set(val *MetricRep) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricRep) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricRep(val *MetricRep) *NullableMetricRep {
	return &NullableMetricRep{value: val, isSet: true}
}

func (v NullableMetricRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


