/*
LaunchDarkly REST API

This documentation describes LaunchDarkly's REST API. To access the complete OpenAPI spec directly, use [Get OpenAPI spec](https://launchdarkly.com/docs/api/other/get-openapi-spec).  To learn how to use LaunchDarkly using the user interface (UI) instead, read our [product documentation](https://launchdarkly.com/docs/home).  ## Authentication  LaunchDarkly's REST API uses the HTTPS protocol with a minimum TLS version of 1.2.  All REST API resources are authenticated with either [personal or service access tokens](https://launchdarkly.com/docs/home/account/api), or session cookies. Other authentication mechanisms are not supported. You can manage personal access tokens on your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page in the LaunchDarkly UI.  LaunchDarkly also has SDK keys, mobile keys, and client-side IDs that are used by our server-side SDKs, mobile SDKs, and JavaScript-based SDKs, respectively. **These keys cannot be used to access our REST API**. These keys are environment-specific, and can only perform read-only operations such as fetching feature flag settings.  | Auth mechanism                                                                                  | Allowed resources                                                                                     | Use cases                                          | | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | -------------------------------------------------- | | [Personal or service access tokens](https://launchdarkly.com/docs/home/account/api) | Can be customized on a per-token basis                                                                | Building scripts, custom integrations, data export. | | SDK keys                                                                                        | Can only access read-only resources specific to server-side SDKs. Restricted to a single environment. | Server-side SDKs                     | | Mobile keys                                                                                     | Can only access read-only resources specific to mobile SDKs, and only for flags marked available to mobile keys. Restricted to a single environment.           | Mobile SDKs                                        | | Client-side ID                                                                                  | Can only access read-only resources specific to JavaScript-based client-side SDKs, and only for flags marked available to client-side. Restricted to a single environment.           | Client-side JavaScript                             |  > #### Keep your access tokens and SDK keys private > > Access tokens should _never_ be exposed in untrusted contexts. Never put an access token in client-side JavaScript, or embed it in a mobile application. LaunchDarkly has special mobile keys that you can embed in mobile apps. If you accidentally expose an access token or SDK key, you can reset it from your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page. > > The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.  ### Authentication using request header  The preferred way to authenticate with the API is by adding an `Authorization` header containing your access token to your requests. The value of the `Authorization` header must be your access token.  Manage personal access tokens from the [**Authorization**](https://app.launchdarkly.com/settings/authorization) page.  ### Authentication using session cookie  For testing purposes, you can make API calls directly from your web browser. If you are logged in to the LaunchDarkly application, the API will use your existing session to authenticate calls.  Depending on the permissions granted as part of your [role](https://launchdarkly.com/docs/home/account/roles), you may not have permission to perform some API calls. You will receive a `401` response code in that case.  > ### Modifying the Origin header causes an error > > LaunchDarkly validates that the Origin header for any API request authenticated by a session cookie matches the expected Origin header. The expected Origin header is `https://app.launchdarkly.com`. > > If the Origin header does not match what's expected, LaunchDarkly returns an error. This error can prevent the LaunchDarkly app from working correctly. > > Any browser extension that intentionally changes the Origin header can cause this problem. For example, Cross-Origin Resource Sharing (CORS) extensions used during development can modify the Origin header and cause the app to fail. > > To prevent this error, do not modify your Origin header. > > LaunchDarkly does not require origin matching when authenticating with an access token, so this issue does not affect normal API usage.  ## Representations  All resources expect and return JSON response bodies. Error responses also send a JSON body. To learn more about the error format of the API, read [Errors](https://launchdarkly.com/docs/api#errors).  In practice this means that you always get a response with a `Content-Type` header set to `application/json`.  In addition, request bodies for `PATCH`, `POST`, and `PUT` requests must be encoded as JSON with a `Content-Type` header set to `application/json`.  Unless an endpoint's documentation says otherwise, the API ignores request body fields that are not part of that endpoint's request schema, and does not return an error for them. A successful response does not mean that every field you sent was applied. To confirm the result, check the response body.  ### Summary and detailed representations  When you fetch a list of resources, the response includes only the most important attributes of each resource. This is a _summary representation_ of the resource. When you fetch an individual resource, such as a single feature flag, you receive a _detailed representation_ of the resource.  The best way to find a detailed representation is to follow links. Every summary representation includes a link to its detailed representation.  ### Expanding responses  Sometimes the detailed representation of a resource does not include all of the attributes of the resource by default. If this is the case, the request method will clearly document this and describe which attributes you can include in an expanded response.  To include the additional attributes, append the `expand` request parameter to your request and add a comma-separated list of the attributes to include. For example, when you append `?expand=members,maintainers` to the [Get team](https://launchdarkly.com/docs/api/teams/get-team) endpoint, the expanded response includes both of these attributes.  ### Links and addressability  The best way to navigate the API is by following links. These are attributes in representations that link to other resources. The API always uses the same format for links:  - Links to other resources within the API are encapsulated in a `_links` object - If the resource has a corresponding link to HTML content on the site, it is stored in a special `_site` link  Each link has two attributes:  - An `href`, which contains the URL - A `type`, which describes the content type  For example, a feature resource might return the following:  ```json {   \"_links\": {     \"parent\": {       \"href\": \"/api/features\",       \"type\": \"application/json\"     },     \"self\": {       \"href\": \"/api/features/sort.order\",       \"type\": \"application/json\"     }   },   \"_site\": {     \"href\": \"/features/sort.order\",     \"type\": \"text/html\"   } } ```  From this, you can navigate to the parent collection of features by following the `parent` link, or navigate to the site page for the feature by following the `_site` link.  Collections are always represented as a JSON object with an `items` attribute containing an array of representations. Like all other representations, collections have `_links` defined at the top level.  Paginated collections include `first`, `last`, `next`, and `prev` links containing a URL with the respective set of elements in the collection.  ## Updates  Resources that accept partial updates use the `PATCH` verb. Most resources support the [JSON patch](https://launchdarkly.com/docs/api#updates-using-json-patch) format. Some resources also support the [JSON merge patch](https://launchdarkly.com/docs/api#updates-using-json-merge-patch) format, and some resources support the [semantic patch](https://launchdarkly.com/docs/api#updates-using-semantic-patch) format, which is a way to specify the modifications to perform as a set of executable instructions. Each resource supports optional [comments](https://launchdarkly.com/docs/api#updates-with-comments) that you can submit with updates. Comments appear in outgoing webhooks, the audit log, and other integrations.  When a resource supports both JSON patch and semantic patch, we document both in the request method. However, the specific request body fields and descriptions included in our documentation only match one type of patch or the other.  ### Updates using JSON patch  [JSON patch](https://datatracker.ietf.org/doc/html/rfc6902) is a way to specify the modifications to perform on a resource. JSON patch uses paths and a limited set of operations to describe how to transform the current state of the resource into a new state. JSON patch documents are always arrays, where each element contains an operation, a path to the field to update, and the new value.  For example, in this feature flag representation:  ```json {     \"name\": \"New recommendations engine\",     \"key\": \"engine.enable\",     \"description\": \"This is the description\",     ... } ``` You can change the feature flag's description with the following patch document:  ```json [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"This is the new description\" }] ```  You can specify multiple modifications to perform in a single request. You can also test that certain preconditions are met before applying the patch:  ```json [   { \"op\": \"test\", \"path\": \"/version\", \"value\": 10 },   { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" } ] ```  The above patch request tests whether the feature flag's `version` is `10`, and if so, changes the feature flag's description.  Attributes that are not editable, such as a resource's `_links`, have names that start with an underscore.  ### Updates using JSON merge patch  [JSON merge patch](https://datatracker.ietf.org/doc/html/rfc7386) is another format for specifying the modifications to perform on a resource. JSON merge patch is less expressive than JSON patch. However, in many cases it is simpler to construct a merge patch document. For example, you can change a feature flag's description with the following merge patch document:  ```json {   \"description\": \"New flag description\" } ```  ### Updates using semantic patch  Some resources support the semantic patch format. A semantic patch is a way to specify the modifications to perform on a resource as a set of executable instructions.  Semantic patch allows you to be explicit about intent using precise, custom instructions. In many cases, you can define semantic patch instructions independently of the current state of the resource. This can be useful when defining a change that may be applied at a future date.  To make a semantic patch request, you must append `domain-model=launchdarkly.semanticpatch` to your `Content-Type` header.  Here's how:  ``` Content-Type: application/json; domain-model=launchdarkly.semanticpatch ```  If you call a semantic patch resource without this header, you will receive a `400` response because your semantic patch will be interpreted as a JSON patch.  The body of a semantic patch request takes the following properties:  * `comment` (string): (Optional) A description of the update. * `environmentKey` (string): (Required for some resources only) The environment key. * `instructions` (array): (Required) A list of actions the update should perform. Each action in the list must be an object with a `kind` property that indicates the instruction. If the instruction requires parameters, you must include those parameters as additional fields in the object. The documentation for each resource that supports semantic patch includes the available instructions and any additional parameters.  For example:  ```json {   \"comment\": \"optional comment\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  Semantic patches are not applied partially; either all of the instructions are applied or none of them are. If **any** instruction is invalid, the endpoint returns an error and will not change the resource. If all instructions are valid, the request succeeds and the resources are updated if necessary, or left unchanged if they are already in the state you request.  ### Updates with comments  You can submit optional comments with `PATCH` changes.  To submit a comment along with a JSON patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"patch\": [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" }] } ```  To submit a comment along with a JSON merge patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"merge\": { \"description\": \"New flag description\" } } ```  To submit a comment along with a semantic patch, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  ## Errors  The API always returns errors in a common format. Here's an example:  ```json {   \"code\": \"invalid_request\",   \"message\": \"A feature with that key already exists\",   \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\" } ```  The `code` indicates the general class of error. The `message` is a human-readable explanation of what went wrong. The `id` is a unique identifier. Use it when you're working with LaunchDarkly Support to debug a problem with a specific API call.  ### HTTP status error response codes  | Code | Definition        | Description                                                                                       | Possible Solution                                                | | ---- | ----------------- | ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- | | 400  | Invalid request       | The request cannot be understood.                                    | Ensure JSON syntax in request body is correct.                   | | 401  | Invalid access token      | Requestor is unauthorized or does not have permission for this API call.                                                | Ensure your API access token is valid and has the appropriate permissions.                                     | | 403  | Forbidden         | Requestor does not have access to this resource.                                                | Ensure that the account member or access token has proper permissions set. | | 404  | Invalid resource identifier | The requested resource is not valid. | Ensure that the resource is correctly identified by ID or key. | | 405  | Method not allowed | The request method is not allowed on this resource. | Ensure that the HTTP verb is correct. | | 409  | Conflict          | The API request can not be completed because it conflicts with a concurrent API request. | Retry your request.                                              | | 422  | Unprocessable entity | The API request can not be completed because the update description can not be understood. | Ensure that the request body is correct for the type of patch you are using, either JSON patch or semantic patch. | 429  | Too many requests | Read [Rate limiting](https://launchdarkly.com/docs/api#rate-limiting).                                               | Wait and try again later.                                        |  ## CORS  The LaunchDarkly API supports Cross Origin Resource Sharing (CORS) for AJAX requests from any origin. If an `Origin` header is given in a request, it will be echoed as an explicitly allowed origin. Otherwise the request returns a wildcard, `Access-Control-Allow-Origin: *`. For more information on CORS, read the [CORS W3C Recommendation](http://www.w3.org/TR/cors). Example CORS headers might look like:  ```http Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, Authorization Access-Control-Allow-Methods: OPTIONS, GET, DELETE, PATCH Access-Control-Allow-Origin: * Access-Control-Max-Age: 300 ```  You can make authenticated CORS calls just as you would make same-origin calls, using either [token or session-based authentication](https://launchdarkly.com/docs/api#authentication). If you are using session authentication, you should set the `withCredentials` property for your `xhr` request to `true`. You should never expose your access tokens to untrusted entities.  ## Rate limiting  We use several rate-limiting strategies to ensure the availability of our APIs. Rate-limited calls to our APIs return a `429` status code and include headers to indicate the current rate limit status. The specific headers returned depend on the API route that was called. Limits differ based on the route, authentication mechanism, and other factors.  Each set of headers below appears only when the corresponding limit is being enforced for your call. A given route may be subject to any combination of these limits, so a response can include one, several, or none of these headers. A missing header indicates that the limit was not applied to this specific call; it does not necessarily indicate that the limit does not exist. To reduce usage before hitting a `429` status, program against whichever rate limit headers are present rather than expecting a specific header.  We do not publicly document the specific number of calls permitted by any of these limits, and these limits may change. We encourage clients to program against the specification and rely on the headers described below, rather than hardcoding the current limits.  > ### Rate limiting and SDKs > > LaunchDarkly SDKs are never rate limited and do not use the API endpoints defined here. LaunchDarkly uses a different set of approaches, including streaming/server-sent events and a global CDN, to ensure availability to the routes used by LaunchDarkly SDKs.  ### Global rate limits  Authenticated requests are subject to a global limit. This is the maximum number of calls that your account can make to the API per ten seconds. All service and personal access tokens on the account share this limit, so exceeding the limit with one access token will impact other tokens. Calls that are subject to global rate limits may return the headers below:  | Header name                    | Description                                                                      | | ------------------------------ | -------------------------------------------------------------------------------- | | `X-Ratelimit-Global-Limit`     | The maximum number of requests the account is permitted to make per ten seconds. | | `X-Ratelimit-Global-Remaining` | The number of requests remaining in the current global rate limit window.        | | `X-Ratelimit-Reset`            | The time at which the current rate limit window resets in epoch milliseconds.    |  ### Route-level rate limits  Some authenticated routes have custom rate limits. These also reset every ten seconds. Any service or personal access tokens hitting the same route share this limit, so exceeding the limit with one access token may impact other tokens. Calls that are subject to route-level rate limits return the headers below:  | Header name                   | Description                                                                                           | | ----------------------------- | ----------------------------------------------------------------------------------------------------- | | `X-Ratelimit-Route-Limit`     | The maximum number of requests to the current route permitted per ten seconds.           | | `X-Ratelimit-Route-Remaining` | The number of requests remaining for the current route in the current rate limit window. | | `X-Ratelimit-Reset`           | The time at which the current rate limit window resets in epoch milliseconds.            |  A _route_ represents a specific URL pattern and verb. For example, the [Delete environment](https://launchdarkly.com/docs/api/environments/delete-environment) endpoint is considered a single route, and each call to delete an environment counts against your route-level rate limit for that route.  ### Access token rate limits  Some calls are rate limited per access token. Unlike the global and route-level limits, this limit applies to a single service or personal access token on its own. Exceeding a limit with one access token does not affect other tokens on the account. Calls that are subject to access token rate limits return these headers:  | Header name                        | Description                                                                             | | ---------------------------------- | --------------------------------------------------------------------------------------- | | `X-Ratelimit-Auth-Token-Limit`     | The maximum number of requests the access token can make per ten seconds.               | | `X-Ratelimit-Auth-Token-Remaining` | The number of requests remaining for the access token in the current rate limit window. | | `X-Ratelimit-Auth-Token-Reset`     | The time at which the current rate limit window resets in epoch milliseconds.           |  Unlike the other rate limits, access token rate limits report their own reset time in the `X-Ratelimit-Auth-Token-Reset` header instead of in `X-Ratelimit-Reset`.  ### IP-based rate limiting  We also employ IP-based rate limiting on some API routes. If you hit an IP-based rate limit, your API response will include a `Retry-After` header indicating how long to wait before re-trying the call. Clients must wait at least `Retry-After` seconds before making additional calls to our API, and should employ jitter and backoff strategies to avoid triggering rate limits again.  ## OpenAPI (Swagger) and client libraries  We have a [complete OpenAPI (Swagger) specification](https://app.launchdarkly.com/api/v2/openapi.json) for our API.  We auto-generate multiple client libraries based on our OpenAPI specification. To learn more, visit the [collection of client libraries on GitHub](https://github.com/search?q=topic%3Alaunchdarkly-api+org%3Alaunchdarkly&type=Repositories). Alternatively, you can use the specification to generate client libraries to interact with our REST API in your language of choice. Or, you can refer to our API endpoints' documentation for guidance on how to make requests with a common HTTP library in your language of choice.  Our OpenAPI specification is supported by several API-based tools such as Postman and Insomnia. In many cases, you can directly import our specification to explore our APIs.  ## Method overriding  Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those environments, our API endpoints that use `DELETE`, `PATCH`, and `PUT` verbs are inaccessible.  To avoid this issue, our API supports the `X-HTTP-Method-Override` header, allowing clients to \"tunnel\" `DELETE`, `PATCH`, and `PUT` requests using a `POST` request.  For example, to call a `PATCH` endpoint using a `POST` request, you can include `X-HTTP-Method-Override:PATCH` as a header.  ## Beta resources  We sometimes release new API resources in **beta** status before we release them with general availability.  Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  We try to promote resources into general availability as quickly as possible. This happens after sufficient testing and when we're satisfied that we no longer need to make backwards-incompatible changes.  We mark beta resources with a \"Beta\" callout in our documentation, pictured below:  > ### This feature is in beta > > To use this feature, pass in a header including the `LD-API-Version` key with value set to `beta`. Use this header with each call. To learn more, read [Beta resources](https://launchdarkly.com/docs/api#beta-resources). > > Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  ### Using beta resources  To use a beta resource, you must include a header in the request. If you call a beta resource without this header, you receive a `403` response.  Use this header:  ``` LD-API-Version: beta ```  ## Federal and EU environments  In addition to the commercial versions, LaunchDarkly offers instances for federal agencies and those based in the European Union (EU).  ### Federal environments  The version of LaunchDarkly that is available on domains controlled by the United States government is different from the version of LaunchDarkly available to the general public. If you are an employee or contractor for a United States federal agency and use LaunchDarkly in your work, you likely use the federal instance of LaunchDarkly.  If you are working in the federal instance of LaunchDarkly, the base URI for each request is `https://app.launchdarkly.us`.  To learn more, read [LaunchDarkly in federal environments](https://launchdarkly.com/docs/home/infrastructure/federal).  ### EU environments  The version of LaunchDarkly that is available in the EU is different from the version of LaunchDarkly available to other regions. If you are based in the EU, you likely use the EU instance of LaunchDarkly. The LaunchDarkly EU instance complies with EU data residency principles, including the protection and confidentiality of EU customer information.  If you are working in the EU instance of LaunchDarkly, the base URI for each request is `https://app.eu.launchdarkly.com`.  To learn more, read [LaunchDarkly in the European Union (EU)](https://launchdarkly.com/docs/home/infrastructure/eu).  ## Versioning  We try hard to keep our REST API backwards compatible, but we occasionally have to make backwards-incompatible changes in the process of shipping new features. These breaking changes can cause unexpected behavior if you don't prepare for them accordingly.  Updates to our REST API include support for the latest features in LaunchDarkly. We also release a new version of our REST API every time we make a breaking change. We provide simultaneous support for multiple API versions so you can migrate from your current API version to a new version at your own pace.  ### Setting the API version per request  You can set the API version on a specific request by sending an `LD-API-Version` header, as shown in the example below:  ``` LD-API-Version: 20240415 ```  The header value is the version number of the API version you would like to request. The number for each version corresponds to the date the version was released in `yyyymmdd` format. In the example above the version `20240415` corresponds to April 15, 2024.  ### Setting the API version per access token  When you create an access token, you must specify a specific version of the API to use. This ensures that integrations using this token cannot be broken by version changes.  Tokens created before versioning was released have their version set to `20160426`, which is the version of the API that existed before the current versioning scheme, so that they continue working the same way they did before versioning.  If you would like to upgrade your integration to use a new API version, you can explicitly set the header described above.  > ### Best practice: Set the header for every client or integration > > We recommend that you set the API version header explicitly in any client or integration you build. > > Only rely on the access token API version during manual testing.  ### API version changelog  <table>   <tr>     <th>Version</th>     <th>Changes</th>     <th>End of life (EOL)</th>   </tr>   <tr>     <td>`20240415`</td>     <td>       <ul><li>Changed several endpoints from unpaginated to paginated. Use the `limit` and `offset` query parameters to page through the results.</li> <li>Changed the [list access tokens](https://launchdarkly.com/docs/api/access-tokens/get-tokens) endpoint: <ul><li>Response is now paginated with a default limit of `25`</li></ul></li> <li>Changed the [list account members](https://launchdarkly.com/docs/api/account-members/get-members) endpoint: <ul><li>The `accessCheck` filter is no longer available</li></ul></li> <li>Changed the [list custom roles](https://launchdarkly.com/docs/api/custom-roles/get-custom-roles) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `environments` field is now only returned if the request is filtered by environment, using the `filterEnv` query parameter</li><li>The `followerId`, `hasDataExport`, `status`, `contextKindTargeted`, and `segmentTargeted` filters are no longer available</li><li>The `compare` query parameter is no longer available</li></ul></li> <li>Changed the [list segments](https://launchdarkly.com/docs/api/segments/get-segments) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list teams](https://launchdarkly.com/docs/api/teams/get-teams) endpoint: <ul><li>The `expand` parameter no longer supports including `projects` or `roles`</li><li>In paginated results, the maximum page size is now 100</li></ul></li> <li>Changed the [get workflows](https://launchdarkly.com/docs/api/workflows/get-workflows) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `_conflicts` field in the response is no longer available</li></ul></li> </ul>     </td>     <td>Current</td>   </tr>   <tr>     <td>`20220603`</td>     <td>       <ul><li>Changed the [list projects](https://launchdarkly.com/docs/api/projects/get-projects) return value:<ul><li>Response is now paginated with a default limit of `20`.</li><li>Added support for filter and sort.</li><li>The project `environments` field is now expandable. This field is omitted by default.</li></ul></li><li>Changed the [get project](https://launchdarkly.com/docs/api/projects/get-project) return value:<ul><li>The `environments` field is now expandable. This field is omitted by default.</li></ul></li></ul>     </td>     <td>2025-04-15</td>   </tr>   <tr>     <td>`20210729`</td>     <td>       <ul><li>Changed the [create approval request](https://launchdarkly.com/docs/api/approvals/post-approval-request) return value. It now returns HTTP Status Code `201` instead of `200`.</li><li> Changed the [get user](https://launchdarkly.com/docs/api/users/get-user) return value. It now returns a user record, not a user. </li><li>Added additional optional fields to environment, segments, flags, members, and segments, including the ability to create big segments. </li><li> Added default values for flag variations when new environments are created. </li><li>Added filtering and pagination for getting flags and members, including `limit`, `number`, `filter`, and `sort` query parameters. </li><li>Added endpoints for expiring user targets for flags and segments, scheduled changes, access tokens, Relay Proxy configuration, integrations and subscriptions, and approvals. </li></ul>     </td>     <td>2023-06-03</td>   </tr>   <tr>     <td>`20191212`</td>     <td>       <ul><li>[List feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) now defaults to sending summaries of feature flag configurations, equivalent to setting the query parameter `summary=true`. Summaries omit flag targeting rules and individual user targets from the payload. </li><li> Added endpoints for flags, flag status, projects, environments, audit logs, members, users, custom roles, segments, usage, streams, events, and data export. </li></ul>     </td>     <td>2022-07-29</td>   </tr>   <tr>     <td>`20160426`</td>     <td>       <ul><li>Initial versioning of API. Tokens created before versioning have their version set to this.</li></ul>     </td>     <td>2020-12-12</td>   </tr> </table>  To learn more about how EOL is determined, read LaunchDarkly's [End of Life (EOL) Policy](https://launchdarkly.com/policies/end-of-life-policy/). 

API version: 2.0
Contact: support@launchdarkly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ldapi

import (
	"encoding/json"
	"fmt"
)

// checks if the EnvironmentRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentRelease{}

// EnvironmentRelease A flag release in one environment, made up of one or more ordered phases. The top-level status is derived from the phases: it reflects the active phase, reports `waiting` between successfully completed phases, and reports the terminal outcome of the release. 
type EnvironmentRelease struct {
	Id string `json:"id"`
	Status ReleaseStatus `json:"status"`
	StatusReason *ReleaseStatusReason `json:"statusReason,omitempty"`
	ProjectKey string `json:"projectKey"`
	EnvironmentKey string `json:"environmentKey"`
	FlagKey string `json:"flagKey"`
	// ID of the variation the release started from.
	OriginalVariationId string `json:"originalVariationId"`
	// The value of the variation. This may be a string, boolean, number, JSON object, or array.
	OriginalVariationValue interface{} `json:"originalVariationValue"`
	// ID of the variation being released.
	TargetVariationId string `json:"targetVariationId"`
	// The value of the variation. This may be a string, boolean, number, JSON object, or array.
	TargetVariationValue interface{} `json:"targetVariationValue"`
	// Key of the context kind the release randomizes on.
	RandomizationUnit string `json:"randomizationUnit"`
	CleanupBehavior CleanupBehavior `json:"cleanupBehavior"`
	Kind ReleaseType `json:"kind"`
	Metrics []ReleaseMetric `json:"metrics"`
	Phases []EnvironmentReleasePhase `json:"phases"`
	// Zero-based index of the active phase. Null before the first phase, between phases, and after the release reaches a terminal status.
	CurrentPhaseIndex NullableInt32 `json:"currentPhaseIndex"`
	AdvancementMode EnvironmentReleaseAdvancementMode `json:"advancementMode"`
	// Creation time in Unix milliseconds.
	CreatedAt int64 `json:"createdAt"`
	// Start of the first phase in Unix milliseconds. Absent until a phase has started.
	StartedAt *int64 `json:"startedAt,omitempty"`
	// End of the last started phase in Unix milliseconds. Present only after the release reaches a terminal status of `completed` or `reverted`.
	EndedAt *int64 `json:"endedAt,omitempty"`
	Links *map[string]EnvironmentReleaseLink `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentRelease EnvironmentRelease

// NewEnvironmentRelease instantiates a new EnvironmentRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentRelease(id string, status ReleaseStatus, projectKey string, environmentKey string, flagKey string, originalVariationId string, originalVariationValue interface{}, targetVariationId string, targetVariationValue interface{}, randomizationUnit string, cleanupBehavior CleanupBehavior, kind ReleaseType, metrics []ReleaseMetric, phases []EnvironmentReleasePhase, currentPhaseIndex NullableInt32, advancementMode EnvironmentReleaseAdvancementMode, createdAt int64) *EnvironmentRelease {
	this := EnvironmentRelease{}
	this.Id = id
	this.Status = status
	this.ProjectKey = projectKey
	this.EnvironmentKey = environmentKey
	this.FlagKey = flagKey
	this.OriginalVariationId = originalVariationId
	this.OriginalVariationValue = originalVariationValue
	this.TargetVariationId = targetVariationId
	this.TargetVariationValue = targetVariationValue
	this.RandomizationUnit = randomizationUnit
	this.CleanupBehavior = cleanupBehavior
	this.Kind = kind
	this.Metrics = metrics
	this.Phases = phases
	this.CurrentPhaseIndex = currentPhaseIndex
	this.AdvancementMode = advancementMode
	this.CreatedAt = createdAt
	return &this
}

// NewEnvironmentReleaseWithDefaults instantiates a new EnvironmentRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentReleaseWithDefaults() *EnvironmentRelease {
	this := EnvironmentRelease{}
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentRelease) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentRelease) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *EnvironmentRelease) GetStatus() ReleaseStatus {
	if o == nil {
		var ret ReleaseStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetStatusOk() (*ReleaseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EnvironmentRelease) SetStatus(v ReleaseStatus) {
	o.Status = v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *EnvironmentRelease) GetStatusReason() ReleaseStatusReason {
	if o == nil || IsNil(o.StatusReason) {
		var ret ReleaseStatusReason
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetStatusReasonOk() (*ReleaseStatusReason, bool) {
	if o == nil || IsNil(o.StatusReason) {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *EnvironmentRelease) HasStatusReason() bool {
	if o != nil && !IsNil(o.StatusReason) {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given ReleaseStatusReason and assigns it to the StatusReason field.
func (o *EnvironmentRelease) SetStatusReason(v ReleaseStatusReason) {
	o.StatusReason = &v
}

// GetProjectKey returns the ProjectKey field value
func (o *EnvironmentRelease) GetProjectKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetProjectKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectKey, true
}

// SetProjectKey sets field value
func (o *EnvironmentRelease) SetProjectKey(v string) {
	o.ProjectKey = v
}

// GetEnvironmentKey returns the EnvironmentKey field value
func (o *EnvironmentRelease) GetEnvironmentKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentKey
}

// GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetEnvironmentKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentKey, true
}

// SetEnvironmentKey sets field value
func (o *EnvironmentRelease) SetEnvironmentKey(v string) {
	o.EnvironmentKey = v
}

// GetFlagKey returns the FlagKey field value
func (o *EnvironmentRelease) GetFlagKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlagKey
}

// GetFlagKeyOk returns a tuple with the FlagKey field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetFlagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagKey, true
}

// SetFlagKey sets field value
func (o *EnvironmentRelease) SetFlagKey(v string) {
	o.FlagKey = v
}

// GetOriginalVariationId returns the OriginalVariationId field value
func (o *EnvironmentRelease) GetOriginalVariationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalVariationId
}

// GetOriginalVariationIdOk returns a tuple with the OriginalVariationId field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetOriginalVariationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalVariationId, true
}

// SetOriginalVariationId sets field value
func (o *EnvironmentRelease) SetOriginalVariationId(v string) {
	o.OriginalVariationId = v
}

// GetOriginalVariationValue returns the OriginalVariationValue field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *EnvironmentRelease) GetOriginalVariationValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.OriginalVariationValue
}

// GetOriginalVariationValueOk returns a tuple with the OriginalVariationValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentRelease) GetOriginalVariationValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OriginalVariationValue) {
		return nil, false
	}
	return &o.OriginalVariationValue, true
}

// SetOriginalVariationValue sets field value
func (o *EnvironmentRelease) SetOriginalVariationValue(v interface{}) {
	o.OriginalVariationValue = v
}

// GetTargetVariationId returns the TargetVariationId field value
func (o *EnvironmentRelease) GetTargetVariationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetVariationId
}

// GetTargetVariationIdOk returns a tuple with the TargetVariationId field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetTargetVariationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetVariationId, true
}

// SetTargetVariationId sets field value
func (o *EnvironmentRelease) SetTargetVariationId(v string) {
	o.TargetVariationId = v
}

// GetTargetVariationValue returns the TargetVariationValue field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *EnvironmentRelease) GetTargetVariationValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.TargetVariationValue
}

// GetTargetVariationValueOk returns a tuple with the TargetVariationValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentRelease) GetTargetVariationValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TargetVariationValue) {
		return nil, false
	}
	return &o.TargetVariationValue, true
}

// SetTargetVariationValue sets field value
func (o *EnvironmentRelease) SetTargetVariationValue(v interface{}) {
	o.TargetVariationValue = v
}

// GetRandomizationUnit returns the RandomizationUnit field value
func (o *EnvironmentRelease) GetRandomizationUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RandomizationUnit
}

// GetRandomizationUnitOk returns a tuple with the RandomizationUnit field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetRandomizationUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RandomizationUnit, true
}

// SetRandomizationUnit sets field value
func (o *EnvironmentRelease) SetRandomizationUnit(v string) {
	o.RandomizationUnit = v
}

// GetCleanupBehavior returns the CleanupBehavior field value
func (o *EnvironmentRelease) GetCleanupBehavior() CleanupBehavior {
	if o == nil {
		var ret CleanupBehavior
		return ret
	}

	return o.CleanupBehavior
}

// GetCleanupBehaviorOk returns a tuple with the CleanupBehavior field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetCleanupBehaviorOk() (*CleanupBehavior, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CleanupBehavior, true
}

// SetCleanupBehavior sets field value
func (o *EnvironmentRelease) SetCleanupBehavior(v CleanupBehavior) {
	o.CleanupBehavior = v
}

// GetKind returns the Kind field value
func (o *EnvironmentRelease) GetKind() ReleaseType {
	if o == nil {
		var ret ReleaseType
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetKindOk() (*ReleaseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *EnvironmentRelease) SetKind(v ReleaseType) {
	o.Kind = v
}

// GetMetrics returns the Metrics field value
func (o *EnvironmentRelease) GetMetrics() []ReleaseMetric {
	if o == nil {
		var ret []ReleaseMetric
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetMetricsOk() ([]ReleaseMetric, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *EnvironmentRelease) SetMetrics(v []ReleaseMetric) {
	o.Metrics = v
}

// GetPhases returns the Phases field value
func (o *EnvironmentRelease) GetPhases() []EnvironmentReleasePhase {
	if o == nil {
		var ret []EnvironmentReleasePhase
		return ret
	}

	return o.Phases
}

// GetPhasesOk returns a tuple with the Phases field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetPhasesOk() ([]EnvironmentReleasePhase, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phases, true
}

// SetPhases sets field value
func (o *EnvironmentRelease) SetPhases(v []EnvironmentReleasePhase) {
	o.Phases = v
}

// GetCurrentPhaseIndex returns the CurrentPhaseIndex field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *EnvironmentRelease) GetCurrentPhaseIndex() int32 {
	if o == nil || o.CurrentPhaseIndex.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CurrentPhaseIndex.Get()
}

// GetCurrentPhaseIndexOk returns a tuple with the CurrentPhaseIndex field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentRelease) GetCurrentPhaseIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPhaseIndex.Get(), o.CurrentPhaseIndex.IsSet()
}

// SetCurrentPhaseIndex sets field value
func (o *EnvironmentRelease) SetCurrentPhaseIndex(v int32) {
	o.CurrentPhaseIndex.Set(&v)
}

// GetAdvancementMode returns the AdvancementMode field value
func (o *EnvironmentRelease) GetAdvancementMode() EnvironmentReleaseAdvancementMode {
	if o == nil {
		var ret EnvironmentReleaseAdvancementMode
		return ret
	}

	return o.AdvancementMode
}

// GetAdvancementModeOk returns a tuple with the AdvancementMode field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetAdvancementModeOk() (*EnvironmentReleaseAdvancementMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdvancementMode, true
}

// SetAdvancementMode sets field value
func (o *EnvironmentRelease) SetAdvancementMode(v EnvironmentReleaseAdvancementMode) {
	o.AdvancementMode = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentRelease) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentRelease) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *EnvironmentRelease) GetStartedAt() int64 {
	if o == nil || IsNil(o.StartedAt) {
		var ret int64
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetStartedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *EnvironmentRelease) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given int64 and assigns it to the StartedAt field.
func (o *EnvironmentRelease) SetStartedAt(v int64) {
	o.StartedAt = &v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *EnvironmentRelease) GetEndedAt() int64 {
	if o == nil || IsNil(o.EndedAt) {
		var ret int64
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetEndedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.EndedAt) {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *EnvironmentRelease) HasEndedAt() bool {
	if o != nil && !IsNil(o.EndedAt) {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given int64 and assigns it to the EndedAt field.
func (o *EnvironmentRelease) SetEndedAt(v int64) {
	o.EndedAt = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnvironmentRelease) GetLinks() map[string]EnvironmentReleaseLink {
	if o == nil || IsNil(o.Links) {
		var ret map[string]EnvironmentReleaseLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRelease) GetLinksOk() (*map[string]EnvironmentReleaseLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnvironmentRelease) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]EnvironmentReleaseLink and assigns it to the Links field.
func (o *EnvironmentRelease) SetLinks(v map[string]EnvironmentReleaseLink) {
	o.Links = &v
}

func (o EnvironmentRelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentRelease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	if !IsNil(o.StatusReason) {
		toSerialize["statusReason"] = o.StatusReason
	}
	toSerialize["projectKey"] = o.ProjectKey
	toSerialize["environmentKey"] = o.EnvironmentKey
	toSerialize["flagKey"] = o.FlagKey
	toSerialize["originalVariationId"] = o.OriginalVariationId
	if o.OriginalVariationValue != nil {
		toSerialize["originalVariationValue"] = o.OriginalVariationValue
	}
	toSerialize["targetVariationId"] = o.TargetVariationId
	if o.TargetVariationValue != nil {
		toSerialize["targetVariationValue"] = o.TargetVariationValue
	}
	toSerialize["randomizationUnit"] = o.RandomizationUnit
	toSerialize["cleanupBehavior"] = o.CleanupBehavior
	toSerialize["kind"] = o.Kind
	toSerialize["metrics"] = o.Metrics
	toSerialize["phases"] = o.Phases
	toSerialize["currentPhaseIndex"] = o.CurrentPhaseIndex.Get()
	toSerialize["advancementMode"] = o.AdvancementMode
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !IsNil(o.EndedAt) {
		toSerialize["endedAt"] = o.EndedAt
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentRelease) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"projectKey",
		"environmentKey",
		"flagKey",
		"originalVariationId",
		"originalVariationValue",
		"targetVariationId",
		"targetVariationValue",
		"randomizationUnit",
		"cleanupBehavior",
		"kind",
		"metrics",
		"phases",
		"currentPhaseIndex",
		"advancementMode",
		"createdAt",
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

	varEnvironmentRelease := _EnvironmentRelease{}

	err = json.Unmarshal(data, &varEnvironmentRelease)

	if err != nil {
		return err
	}

	*o = EnvironmentRelease(varEnvironmentRelease)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusReason")
		delete(additionalProperties, "projectKey")
		delete(additionalProperties, "environmentKey")
		delete(additionalProperties, "flagKey")
		delete(additionalProperties, "originalVariationId")
		delete(additionalProperties, "originalVariationValue")
		delete(additionalProperties, "targetVariationId")
		delete(additionalProperties, "targetVariationValue")
		delete(additionalProperties, "randomizationUnit")
		delete(additionalProperties, "cleanupBehavior")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "metrics")
		delete(additionalProperties, "phases")
		delete(additionalProperties, "currentPhaseIndex")
		delete(additionalProperties, "advancementMode")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "startedAt")
		delete(additionalProperties, "endedAt")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvironmentRelease struct {
	value *EnvironmentRelease
	isSet bool
}

func (v NullableEnvironmentRelease) Get() *EnvironmentRelease {
	return v.value
}

func (v *NullableEnvironmentRelease) Set(val *EnvironmentRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentRelease(val *EnvironmentRelease) *NullableEnvironmentRelease {
	return &NullableEnvironmentRelease{value: val, isSet: true}
}

func (v NullableEnvironmentRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


