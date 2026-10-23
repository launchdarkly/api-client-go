/*
LaunchDarkly REST API

This documentation describes LaunchDarkly's REST API. To access the complete OpenAPI spec directly, use [Get OpenAPI spec](https://launchdarkly.com/docs/api/other/get-openapi-spec).  To learn how to use LaunchDarkly using the user interface (UI) instead, read our [product documentation](https://launchdarkly.com/docs/home).  ## Authentication  LaunchDarkly's REST API uses the HTTPS protocol with a minimum TLS version of 1.2.  All REST API resources are authenticated with either [personal or service access tokens](https://launchdarkly.com/docs/home/account/api), or session cookies. Other authentication mechanisms are not supported. You can manage personal access tokens on your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page in the LaunchDarkly UI.  LaunchDarkly also has SDK keys, mobile keys, and client-side IDs that are used by our server-side SDKs, mobile SDKs, and JavaScript-based SDKs, respectively. **These keys cannot be used to access our REST API**. These keys are environment-specific, and can only perform read-only operations such as fetching feature flag settings.  | Auth mechanism                                                                                  | Allowed resources                                                                                     | Use cases                                          | | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | -------------------------------------------------- | | [Personal or service access tokens](https://launchdarkly.com/docs/home/account/api) | Can be customized on a per-token basis                                                                | Building scripts, custom integrations, data export. | | SDK keys                                                                                        | Can only access read-only resources specific to server-side SDKs. Restricted to a single environment. | Server-side SDKs                     | | Mobile keys                                                                                     | Can only access read-only resources specific to mobile SDKs, and only for flags marked available to mobile keys. Restricted to a single environment.           | Mobile SDKs                                        | | Client-side ID                                                                                  | Can only access read-only resources specific to JavaScript-based client-side SDKs, and only for flags marked available to client-side. Restricted to a single environment.           | Client-side JavaScript                             |  > #### Keep your access tokens and SDK keys private > > Access tokens should _never_ be exposed in untrusted contexts. Never put an access token in client-side JavaScript, or embed it in a mobile application. LaunchDarkly has special mobile keys that you can embed in mobile apps. If you accidentally expose an access token or SDK key, you can reset it from your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page. > > The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.  ### Authentication using request header  The preferred way to authenticate with the API is by adding an `Authorization` header containing your access token to your requests. The value of the `Authorization` header must be your access token.  Manage personal access tokens from the [**Authorization**](https://app.launchdarkly.com/settings/authorization) page.  ### Authentication using session cookie  For testing purposes, you can make API calls directly from your web browser. If you are logged in to the LaunchDarkly application, the API will use your existing session to authenticate calls.  Depending on the permissions granted as part of your [role](https://launchdarkly.com/docs/home/account/roles), you may not have permission to perform some API calls. You will receive a `401` response code in that case.  > ### Modifying the Origin header causes an error > > LaunchDarkly validates that the Origin header for any API request authenticated by a session cookie matches the expected Origin header. The expected Origin header is `https://app.launchdarkly.com`. > > If the Origin header does not match what's expected, LaunchDarkly returns an error. This error can prevent the LaunchDarkly app from working correctly. > > Any browser extension that intentionally changes the Origin header can cause this problem. For example, Cross-Origin Resource Sharing (CORS) extensions used during development can modify the Origin header and cause the app to fail. > > To prevent this error, do not modify your Origin header. > > LaunchDarkly does not require origin matching when authenticating with an access token, so this issue does not affect normal API usage.  ## Representations  All resources expect and return JSON response bodies. Error responses also send a JSON body. To learn more about the error format of the API, read [Errors](https://launchdarkly.com/docs/api#errors).  In practice this means that you always get a response with a `Content-Type` header set to `application/json`.  In addition, request bodies for `PATCH`, `POST`, and `PUT` requests must be encoded as JSON with a `Content-Type` header set to `application/json`.  Unless an endpoint's documentation says otherwise, the API ignores request body fields that are not part of that endpoint's request schema, and does not return an error for them. A successful response does not mean that every field you sent was applied. To confirm the result, check the response body.  ### Summary and detailed representations  When you fetch a list of resources, the response includes only the most important attributes of each resource. This is a _summary representation_ of the resource. When you fetch an individual resource, such as a single feature flag, you receive a _detailed representation_ of the resource.  The best way to find a detailed representation is to follow links. Every summary representation includes a link to its detailed representation.  ### Expanding responses  Sometimes the detailed representation of a resource does not include all of the attributes of the resource by default. If this is the case, the request method will clearly document this and describe which attributes you can include in an expanded response.  To include the additional attributes, append the `expand` request parameter to your request and add a comma-separated list of the attributes to include. For example, when you append `?expand=members,maintainers` to the [Get team](https://launchdarkly.com/docs/api/teams/get-team) endpoint, the expanded response includes both of these attributes.  ### Links and addressability  The best way to navigate the API is by following links. These are attributes in representations that link to other resources. The API always uses the same format for links:  - Links to other resources within the API are encapsulated in a `_links` object - If the resource has a corresponding link to HTML content on the site, it is stored in a special `_site` link  Each link has two attributes:  - An `href`, which contains the URL - A `type`, which describes the content type  For example, a feature resource might return the following:  ```json {   \"_links\": {     \"parent\": {       \"href\": \"/api/features\",       \"type\": \"application/json\"     },     \"self\": {       \"href\": \"/api/features/sort.order\",       \"type\": \"application/json\"     }   },   \"_site\": {     \"href\": \"/features/sort.order\",     \"type\": \"text/html\"   } } ```  From this, you can navigate to the parent collection of features by following the `parent` link, or navigate to the site page for the feature by following the `_site` link.  Collections are always represented as a JSON object with an `items` attribute containing an array of representations. Like all other representations, collections have `_links` defined at the top level.  Paginated collections include `first`, `last`, `next`, and `prev` links containing a URL with the respective set of elements in the collection.  ## Updates  Resources that accept partial updates use the `PATCH` verb. Most resources support the [JSON patch](https://launchdarkly.com/docs/api#updates-using-json-patch) format. Some resources also support the [JSON merge patch](https://launchdarkly.com/docs/api#updates-using-json-merge-patch) format, and some resources support the [semantic patch](https://launchdarkly.com/docs/api#updates-using-semantic-patch) format, which is a way to specify the modifications to perform as a set of executable instructions. Each resource supports optional [comments](https://launchdarkly.com/docs/api#updates-with-comments) that you can submit with updates. Comments appear in outgoing webhooks, the audit log, and other integrations.  When a resource supports both JSON patch and semantic patch, we document both in the request method. However, the specific request body fields and descriptions included in our documentation only match one type of patch or the other.  ### Updates using JSON patch  [JSON patch](https://datatracker.ietf.org/doc/html/rfc6902) is a way to specify the modifications to perform on a resource. JSON patch uses paths and a limited set of operations to describe how to transform the current state of the resource into a new state. JSON patch documents are always arrays, where each element contains an operation, a path to the field to update, and the new value.  For example, in this feature flag representation:  ```json {     \"name\": \"New recommendations engine\",     \"key\": \"engine.enable\",     \"description\": \"This is the description\",     ... } ``` You can change the feature flag's description with the following patch document:  ```json [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"This is the new description\" }] ```  You can specify multiple modifications to perform in a single request. You can also test that certain preconditions are met before applying the patch:  ```json [   { \"op\": \"test\", \"path\": \"/version\", \"value\": 10 },   { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" } ] ```  The above patch request tests whether the feature flag's `version` is `10`, and if so, changes the feature flag's description.  Attributes that are not editable, such as a resource's `_links`, have names that start with an underscore.  ### Updates using JSON merge patch  [JSON merge patch](https://datatracker.ietf.org/doc/html/rfc7386) is another format for specifying the modifications to perform on a resource. JSON merge patch is less expressive than JSON patch. However, in many cases it is simpler to construct a merge patch document. For example, you can change a feature flag's description with the following merge patch document:  ```json {   \"description\": \"New flag description\" } ```  ### Updates using semantic patch  Some resources support the semantic patch format. A semantic patch is a way to specify the modifications to perform on a resource as a set of executable instructions.  Semantic patch allows you to be explicit about intent using precise, custom instructions. In many cases, you can define semantic patch instructions independently of the current state of the resource. This can be useful when defining a change that may be applied at a future date.  To make a semantic patch request, you must append `domain-model=launchdarkly.semanticpatch` to your `Content-Type` header.  Here's how:  ``` Content-Type: application/json; domain-model=launchdarkly.semanticpatch ```  If you call a semantic patch resource without this header, you will receive a `400` response because your semantic patch will be interpreted as a JSON patch.  The body of a semantic patch request takes the following properties:  * `comment` (string): (Optional) A description of the update. * `environmentKey` (string): (Required for some resources only) The environment key. * `instructions` (array): (Required) A list of actions the update should perform. Each action in the list must be an object with a `kind` property that indicates the instruction. If the instruction requires parameters, you must include those parameters as additional fields in the object. The documentation for each resource that supports semantic patch includes the available instructions and any additional parameters.  For example:  ```json {   \"comment\": \"optional comment\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  Semantic patches are not applied partially; either all of the instructions are applied or none of them are. If **any** instruction is invalid, the endpoint returns an error and will not change the resource. If all instructions are valid, the request succeeds and the resources are updated if necessary, or left unchanged if they are already in the state you request.  ### Updates with comments  You can submit optional comments with `PATCH` changes.  To submit a comment along with a JSON patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"patch\": [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" }] } ```  To submit a comment along with a JSON merge patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"merge\": { \"description\": \"New flag description\" } } ```  To submit a comment along with a semantic patch, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  ## Errors  The API always returns errors in a common format. Here's an example:  ```json {   \"code\": \"invalid_request\",   \"message\": \"A feature with that key already exists\",   \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\" } ```  The `code` indicates the general class of error. The `message` is a human-readable explanation of what went wrong. The `id` is a unique identifier. Use it when you're working with LaunchDarkly Support to debug a problem with a specific API call.  ### HTTP status error response codes  | Code | Definition        | Description                                                                                       | Possible Solution                                                | | ---- | ----------------- | ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- | | 400  | Invalid request       | The request cannot be understood.                                    | Ensure JSON syntax in request body is correct.                   | | 401  | Invalid access token      | Requestor is unauthorized or does not have permission for this API call.                                                | Ensure your API access token is valid and has the appropriate permissions.                                     | | 403  | Forbidden         | Requestor does not have access to this resource.                                                | Ensure that the account member or access token has proper permissions set. | | 404  | Invalid resource identifier | The requested resource is not valid. | Ensure that the resource is correctly identified by ID or key. | | 405  | Method not allowed | The request method is not allowed on this resource. | Ensure that the HTTP verb is correct. | | 409  | Conflict          | The API request can not be completed because it conflicts with a concurrent API request. | Retry your request.                                              | | 422  | Unprocessable entity | The API request can not be completed because the update description can not be understood. | Ensure that the request body is correct for the type of patch you are using, either JSON patch or semantic patch. | 429  | Too many requests | Read [Rate limiting](https://launchdarkly.com/docs/api#rate-limiting).                                               | Wait and try again later.                                        |  ## CORS  The LaunchDarkly API supports Cross Origin Resource Sharing (CORS) for AJAX requests from any origin. If an `Origin` header is given in a request, it will be echoed as an explicitly allowed origin. Otherwise the request returns a wildcard, `Access-Control-Allow-Origin: *`. For more information on CORS, read the [CORS W3C Recommendation](http://www.w3.org/TR/cors). Example CORS headers might look like:  ```http Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, Authorization Access-Control-Allow-Methods: OPTIONS, GET, DELETE, PATCH Access-Control-Allow-Origin: * Access-Control-Max-Age: 300 ```  You can make authenticated CORS calls just as you would make same-origin calls, using either [token or session-based authentication](https://launchdarkly.com/docs/api#authentication). If you are using session authentication, you should set the `withCredentials` property for your `xhr` request to `true`. You should never expose your access tokens to untrusted entities.  ## Rate limiting  We use several rate-limiting strategies to ensure the availability of our APIs. Rate-limited calls to our APIs return a `429` status code and include headers to indicate the current rate limit status. The specific headers returned depend on the API route that was called. Limits differ based on the route, authentication mechanism, and other factors.  Each set of headers below appears only when the corresponding limit is being enforced for your call. A given route may be subject to any combination of these limits, so a response can include one, several, or none of these headers. A missing header indicates that the limit was not applied to this specific call; it does not necessarily indicate that the limit does not exist. To reduce usage before hitting a `429` status, program against whichever rate limit headers are present rather than expecting a specific header.  We do not publicly document the specific number of calls permitted by any of these limits, and these limits may change. We encourage clients to program against the specification and rely on the headers described below, rather than hardcoding the current limits.  > ### Rate limiting and SDKs > > LaunchDarkly SDKs are never rate limited and do not use the API endpoints defined here. LaunchDarkly uses a different set of approaches, including streaming/server-sent events and a global CDN, to ensure availability to the routes used by LaunchDarkly SDKs.  ### Global rate limits  Authenticated requests are subject to a global limit. This is the maximum number of calls that your account can make to the API per ten seconds. All service and personal access tokens on the account share this limit, so exceeding the limit with one access token will impact other tokens. Calls that are subject to global rate limits may return the headers below:  | Header name                    | Description                                                                      | | ------------------------------ | -------------------------------------------------------------------------------- | | `X-Ratelimit-Global-Limit`     | The maximum number of requests the account is permitted to make per ten seconds. | | `X-Ratelimit-Global-Remaining` | The number of requests remaining in the current global rate limit window.        | | `X-Ratelimit-Reset`            | The time at which the current rate limit window resets in epoch milliseconds.    |  ### Route-level rate limits  Some authenticated routes have custom rate limits. These also reset every ten seconds. Any service or personal access tokens hitting the same route share this limit, so exceeding the limit with one access token may impact other tokens. Calls that are subject to route-level rate limits return the headers below:  | Header name                   | Description                                                                                           | | ----------------------------- | ----------------------------------------------------------------------------------------------------- | | `X-Ratelimit-Route-Limit`     | The maximum number of requests to the current route permitted per ten seconds.           | | `X-Ratelimit-Route-Remaining` | The number of requests remaining for the current route in the current rate limit window. | | `X-Ratelimit-Reset`           | The time at which the current rate limit window resets in epoch milliseconds.            |  A _route_ represents a specific URL pattern and verb. For example, the [Delete environment](https://launchdarkly.com/docs/api/environments/delete-environment) endpoint is considered a single route, and each call to delete an environment counts against your route-level rate limit for that route.  ### Access token rate limits  Some calls are rate limited per access token. Unlike the global and route-level limits, this limit applies to a single service or personal access token on its own. Exceeding a limit with one access token does not affect other tokens on the account. Calls that are subject to access token rate limits return these headers:  | Header name                        | Description                                                                             | | ---------------------------------- | --------------------------------------------------------------------------------------- | | `X-Ratelimit-Auth-Token-Limit`     | The maximum number of requests the access token can make per ten seconds.               | | `X-Ratelimit-Auth-Token-Remaining` | The number of requests remaining for the access token in the current rate limit window. | | `X-Ratelimit-Auth-Token-Reset`     | The time at which the current rate limit window resets in epoch milliseconds.           |  Unlike the other rate limits, access token rate limits report their own reset time in the `X-Ratelimit-Auth-Token-Reset` header instead of in `X-Ratelimit-Reset`.  ### IP-based rate limiting  We also employ IP-based rate limiting on some API routes. If you hit an IP-based rate limit, your API response will include a `Retry-After` header indicating how long to wait before re-trying the call. Clients must wait at least `Retry-After` seconds before making additional calls to our API, and should employ jitter and backoff strategies to avoid triggering rate limits again.  ## OpenAPI (Swagger) and client libraries  We have a [complete OpenAPI (Swagger) specification](https://app.launchdarkly.com/api/v2/openapi.json) for our API.  We auto-generate multiple client libraries based on our OpenAPI specification. To learn more, visit the [collection of client libraries on GitHub](https://github.com/search?q=topic%3Alaunchdarkly-api+org%3Alaunchdarkly&type=Repositories). Alternatively, you can use the specification to generate client libraries to interact with our REST API in your language of choice. Or, you can refer to our API endpoints' documentation for guidance on how to make requests with a common HTTP library in your language of choice.  Our OpenAPI specification is supported by several API-based tools such as Postman and Insomnia. In many cases, you can directly import our specification to explore our APIs.  ## Method overriding  Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those environments, our API endpoints that use `DELETE`, `PATCH`, and `PUT` verbs are inaccessible.  To avoid this issue, our API supports the `X-HTTP-Method-Override` header, allowing clients to \"tunnel\" `DELETE`, `PATCH`, and `PUT` requests using a `POST` request.  For example, to call a `PATCH` endpoint using a `POST` request, you can include `X-HTTP-Method-Override:PATCH` as a header.  ## Beta resources  We sometimes release new API resources in **beta** status before we release them with general availability.  Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  We try to promote resources into general availability as quickly as possible. This happens after sufficient testing and when we're satisfied that we no longer need to make backwards-incompatible changes.  We mark beta resources with a \"Beta\" callout in our documentation, pictured below:  > ### This feature is in beta > > To use this feature, pass in a header including the `LD-API-Version` key with value set to `beta`. Use this header with each call. To learn more, read [Beta resources](https://launchdarkly.com/docs/api#beta-resources). > > Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  ### Using beta resources  To use a beta resource, you must include a header in the request. If you call a beta resource without this header, you receive a `403` response.  Use this header:  ``` LD-API-Version: beta ```  ## Federal and EU environments  In addition to the commercial versions, LaunchDarkly offers instances for federal agencies and those based in the European Union (EU).  ### Federal environments  The version of LaunchDarkly that is available on domains controlled by the United States government is different from the version of LaunchDarkly available to the general public. If you are an employee or contractor for a United States federal agency and use LaunchDarkly in your work, you likely use the federal instance of LaunchDarkly.  If you are working in the federal instance of LaunchDarkly, the base URI for each request is `https://app.launchdarkly.us`.  To learn more, read [LaunchDarkly in federal environments](https://launchdarkly.com/docs/home/infrastructure/federal).  ### EU environments  The version of LaunchDarkly that is available in the EU is different from the version of LaunchDarkly available to other regions. If you are based in the EU, you likely use the EU instance of LaunchDarkly. The LaunchDarkly EU instance complies with EU data residency principles, including the protection and confidentiality of EU customer information.  If you are working in the EU instance of LaunchDarkly, the base URI for each request is `https://app.eu.launchdarkly.com`.  To learn more, read [LaunchDarkly in the European Union (EU)](https://launchdarkly.com/docs/home/infrastructure/eu).  ## Versioning  We try hard to keep our REST API backwards compatible, but we occasionally have to make backwards-incompatible changes in the process of shipping new features. These breaking changes can cause unexpected behavior if you don't prepare for them accordingly.  Updates to our REST API include support for the latest features in LaunchDarkly. We also release a new version of our REST API every time we make a breaking change. We provide simultaneous support for multiple API versions so you can migrate from your current API version to a new version at your own pace.  ### Setting the API version per request  You can set the API version on a specific request by sending an `LD-API-Version` header, as shown in the example below:  ``` LD-API-Version: 20240415 ```  The header value is the version number of the API version you would like to request. The number for each version corresponds to the date the version was released in `yyyymmdd` format. In the example above the version `20240415` corresponds to April 15, 2024.  ### Setting the API version per access token  When you create an access token, you must specify a specific version of the API to use. This ensures that integrations using this token cannot be broken by version changes.  Tokens created before versioning was released have their version set to `20160426`, which is the version of the API that existed before the current versioning scheme, so that they continue working the same way they did before versioning.  If you would like to upgrade your integration to use a new API version, you can explicitly set the header described above.  > ### Best practice: Set the header for every client or integration > > We recommend that you set the API version header explicitly in any client or integration you build. > > Only rely on the access token API version during manual testing.  ### API version changelog  <table>   <tr>     <th>Version</th>     <th>Changes</th>     <th>End of life (EOL)</th>   </tr>   <tr>     <td>`20240415`</td>     <td>       <ul><li>Changed several endpoints from unpaginated to paginated. Use the `limit` and `offset` query parameters to page through the results.</li> <li>Changed the [list access tokens](https://launchdarkly.com/docs/api/access-tokens/get-tokens) endpoint: <ul><li>Response is now paginated with a default limit of `25`</li></ul></li> <li>Changed the [list account members](https://launchdarkly.com/docs/api/account-members/get-members) endpoint: <ul><li>The `accessCheck` filter is no longer available</li></ul></li> <li>Changed the [list custom roles](https://launchdarkly.com/docs/api/custom-roles/get-custom-roles) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `environments` field is now only returned if the request is filtered by environment, using the `filterEnv` query parameter</li><li>The `followerId`, `hasDataExport`, `status`, `contextKindTargeted`, and `segmentTargeted` filters are no longer available</li><li>The `compare` query parameter is no longer available</li></ul></li> <li>Changed the [list segments](https://launchdarkly.com/docs/api/segments/get-segments) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list teams](https://launchdarkly.com/docs/api/teams/get-teams) endpoint: <ul><li>The `expand` parameter no longer supports including `projects` or `roles`</li><li>In paginated results, the maximum page size is now 100</li></ul></li> <li>Changed the [get workflows](https://launchdarkly.com/docs/api/workflows/get-workflows) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `_conflicts` field in the response is no longer available</li></ul></li> </ul>     </td>     <td>Current</td>   </tr>   <tr>     <td>`20220603`</td>     <td>       <ul><li>Changed the [list projects](https://launchdarkly.com/docs/api/projects/get-projects) return value:<ul><li>Response is now paginated with a default limit of `20`.</li><li>Added support for filter and sort.</li><li>The project `environments` field is now expandable. This field is omitted by default.</li></ul></li><li>Changed the [get project](https://launchdarkly.com/docs/api/projects/get-project) return value:<ul><li>The `environments` field is now expandable. This field is omitted by default.</li></ul></li></ul>     </td>     <td>2025-04-15</td>   </tr>   <tr>     <td>`20210729`</td>     <td>       <ul><li>Changed the [create approval request](https://launchdarkly.com/docs/api/approvals/post-approval-request) return value. It now returns HTTP Status Code `201` instead of `200`.</li><li> Changed the [get user](https://launchdarkly.com/docs/api/users/get-user) return value. It now returns a user record, not a user. </li><li>Added additional optional fields to environment, segments, flags, members, and segments, including the ability to create big segments. </li><li> Added default values for flag variations when new environments are created. </li><li>Added filtering and pagination for getting flags and members, including `limit`, `number`, `filter`, and `sort` query parameters. </li><li>Added endpoints for expiring user targets for flags and segments, scheduled changes, access tokens, Relay Proxy configuration, integrations and subscriptions, and approvals. </li></ul>     </td>     <td>2023-06-03</td>   </tr>   <tr>     <td>`20191212`</td>     <td>       <ul><li>[List feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) now defaults to sending summaries of feature flag configurations, equivalent to setting the query parameter `summary=true`. Summaries omit flag targeting rules and individual user targets from the payload. </li><li> Added endpoints for flags, flag status, projects, environments, audit logs, members, users, custom roles, segments, usage, streams, events, and data export. </li></ul>     </td>     <td>2022-07-29</td>   </tr>   <tr>     <td>`20160426`</td>     <td>       <ul><li>Initial versioning of API. Tokens created before versioning have their version set to this.</li></ul>     </td>     <td>2020-12-12</td>   </tr> </table>  To learn more about how EOL is determined, read LaunchDarkly's [End of Life (EOL) Policy](https://launchdarkly.com/policies/end-of-life-policy/). 

API version: 2.0
Contact: support@launchdarkly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ldapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type EnvironmentReleasesBetaApi interface {

	/*
	GetEnvironmentRelease Get an environment release

	Returns an environment release and all its phases in phase-index order.
The release ID must belong to the flag and environment in the path.
Returns `404` for a release that ended more than 90 days ago.
Pass `expand=metrics` to attach each metric's definition to the top-level
`metrics` array. Analysis settings come from the version pinned when the
release started (`metricVersionId`); name, description, tags, maintainer,
and archive state reflect the current metric and may change afterwards.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param flagKey
	@param environmentKey
	@param releaseId
	@return ApiGetEnvironmentReleaseRequest
	*/
	GetEnvironmentRelease(ctx context.Context, projectKey string, flagKey string, environmentKey string, releaseId string) ApiGetEnvironmentReleaseRequest

	// GetEnvironmentReleaseExecute executes the request
	//  @return EnvironmentRelease
	GetEnvironmentReleaseExecute(r ApiGetEnvironmentReleaseRequest) (*EnvironmentRelease, *http.Response, error)

	/*
	ListEnvironmentReleases List environment releases

	Returns the environment releases for the flag in the environment, newest first.
Each release includes its phases in phase-index order.
Metric definitions are not included. To get them, call
`getEnvironmentRelease` with `expand=metrics`.
Returns an empty `items` array when the flag has no releases in the environment,
or when no release matches the `filter`.
Releases that ended more than 90 days ago are not returned.

Results are paginated. By default, the first 25 releases are returned. Use
`limit` and `offset` to page through the rest, and `totalCount` to learn how
many releases match in total. The `_links` object includes `first`, `prev`,
`next`, and `last` links where they apply.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param flagKey
	@param environmentKey
	@return ApiListEnvironmentReleasesRequest
	*/
	ListEnvironmentReleases(ctx context.Context, projectKey string, flagKey string, environmentKey string) ApiListEnvironmentReleasesRequest

	// ListEnvironmentReleasesExecute executes the request
	//  @return EnvironmentReleaseCollection
	ListEnvironmentReleasesExecute(r ApiListEnvironmentReleasesRequest) (*EnvironmentReleaseCollection, *http.Response, error)

	/*
	StartEnvironmentRelease Start an environment release

	Creates and starts a release on a flag in an environment, either from
inline configuration or from an environment release plan. A release
has one or more phases. Each phase is an audience, targeted by one
flag rule, and has its own steps: the schedule of allocations at
which that rule serves the target variation. After all of a phase's
steps complete, the phase serves the target variation to 100% of its
contexts.

**Inline configuration.** Each entry in `phases` must specify exactly
one targeting mode:

- **Existing rule**: set `ruleId`. The rule must be enabled. Only a
  single-phase release can use an existing rule.
- **New rule**: set `clauses`, with an optional `description`. New
  rules are appended to the end of the flag's rule list. To insert
  them before an existing rule instead, set the top-level
  `beforeRuleId`, which applies to every phase.

Every phase's rule, including an existing rule passed by `ruleId`,
must include
`{"contextKind": "<randomizationUnit>", "attribute": "kind", "op": "in", "values": ["<randomizationUnit>"]}`,
not negated and with no other values, so the release only reaches
contexts of the `randomizationUnit` kind, or multi-contexts
containing the `randomizationUnit` kind. An existing rule needs the
clause even if its other clauses only match that context kind. A
phase without it returns `400`.

For multiple phases, set `phases` with at least two entries. Each phase
creates a rule. The first rule is enabled and starts immediately.
Subsequent rules stay disabled until the release advances to them:
automatically when `advancementMode` is `automatic`, or by an explicit
advance when it is `manual`. `advancementMode` is required for a
multi-phase release.

Releases on the flag's default rule are not supported. A phase that
omits both targeting modes is rejected.

The flag must be on in the environment, or the request returns
`409`. The release only reaches contexts that match a phase's rule.
Other contexts keep getting the flag's default rule, which may
already serve the target variation.

The release type is inferred from `metrics`: an omitted or empty array
creates a progressive release; one or more metrics creates a guarded
release. Metrics must be listed individually; metric groups are not
accepted.

Every start runs release health checks, described in
`ReleaseHealthCheckKind`. A flag that hasn't recently been
evaluated for contexts of the `randomizationUnit` kind fails
`flag_evaluations_with_context`, and one with no recent evaluations
at all also fails `flag_evaluations_any`. Guarded releases also
check that each metric receives events. A failing
check returns `400` with code `health_checks_failed`. To start the
release anyway, list the failing kinds in `bypassHealthChecks`.

**From a release plan.** Set `releasePlanKey` to start from a plan in
this environment instead of configuring the release inline; a request
that also sets fields the plan supplies is rejected. The plan's current
version runs as stored, with metric groups expanded into their members.
Health checks run against the plan's randomization unit and metrics.
Returns `400` if the plan is missing or incomplete, and `422` until
environment release plans are available for the account.

Requires the `updateRules` action on the flag in this environment, plus
`updateFallthrough` when `cleanupBehavior` is
`remove_rules_and_update_default_rule`. A release started from a plan
requires both. The start is recorded in the flag's audit log as a
targeting change that carries `comment`.

In an environment that requires approvals, this endpoint fails with
`405`. Tokens granted the `bypassRequiredApproval` action can still
call it to start the release directly.

This operation is not idempotent. Retrying a request that already
started a release returns `409` with code `release_already_active`.
The earlier request may have succeeded; call `listEnvironmentReleases`
to confirm the running release is yours.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param flagKey
	@param environmentKey
	@return ApiStartEnvironmentReleaseRequest
	*/
	StartEnvironmentRelease(ctx context.Context, projectKey string, flagKey string, environmentKey string) ApiStartEnvironmentReleaseRequest

	// StartEnvironmentReleaseExecute executes the request
	//  @return EnvironmentRelease
	StartEnvironmentReleaseExecute(r ApiStartEnvironmentReleaseRequest) (*EnvironmentRelease, *http.Response, error)

	/*
	StopEnvironmentRelease Stop an environment release

	Stops a release and resolves its phases according to the `resolution`
you specify. To learn how each resolution treats completed and pending
phases, read the `resolution` schema.

Requires the `updateRules` action on the flag in this environment, or
`updateFallthrough` for a historical release that ran on the default
rule. The stop is recorded in the flag's audit log as a targeting
change that names the resolution and carries `comment`.

In an environment that requires approvals, this endpoint fails with
`405`. Tokens granted the `bypassRequiredApproval` action can still
call it to stop the release directly.

Stopping a release that has already completed or been stopped returns
`200` and makes no changes. Returns `404` for a release that ended more
than 90 days ago.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param flagKey
	@param environmentKey
	@param releaseId
	@return ApiStopEnvironmentReleaseRequest
	*/
	StopEnvironmentRelease(ctx context.Context, projectKey string, flagKey string, environmentKey string, releaseId string) ApiStopEnvironmentReleaseRequest

	// StopEnvironmentReleaseExecute executes the request
	//  @return EnvironmentRelease
	StopEnvironmentReleaseExecute(r ApiStopEnvironmentReleaseRequest) (*EnvironmentRelease, *http.Response, error)
}

// EnvironmentReleasesBetaApiService EnvironmentReleasesBetaApi service
type EnvironmentReleasesBetaApiService service

type ApiGetEnvironmentReleaseRequest struct {
	ctx context.Context
	ApiService EnvironmentReleasesBetaApi
	lDAPIVersion *string
	projectKey string
	flagKey string
	environmentKey string
	releaseId string
	expand *[]string
}

// Version of the endpoint.
func (r ApiGetEnvironmentReleaseRequest) LDAPIVersion(lDAPIVersion string) ApiGetEnvironmentReleaseRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// Optional fields to expand.
func (r ApiGetEnvironmentReleaseRequest) Expand(expand []string) ApiGetEnvironmentReleaseRequest {
	r.expand = &expand
	return r
}

func (r ApiGetEnvironmentReleaseRequest) Execute() (*EnvironmentRelease, *http.Response, error) {
	return r.ApiService.GetEnvironmentReleaseExecute(r)
}

/*
GetEnvironmentRelease Get an environment release

Returns an environment release and all its phases in phase-index order.
The release ID must belong to the flag and environment in the path.
Returns `404` for a release that ended more than 90 days ago.
Pass `expand=metrics` to attach each metric's definition to the top-level
`metrics` array. Analysis settings come from the version pinned when the
release started (`metricVersionId`); name, description, tags, maintainer,
and archive state reflect the current metric and may change afterwards.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param flagKey
 @param environmentKey
 @param releaseId
 @return ApiGetEnvironmentReleaseRequest
*/
func (a *EnvironmentReleasesBetaApiService) GetEnvironmentRelease(ctx context.Context, projectKey string, flagKey string, environmentKey string, releaseId string) ApiGetEnvironmentReleaseRequest {
	return ApiGetEnvironmentReleaseRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		flagKey: flagKey,
		environmentKey: environmentKey,
		releaseId: releaseId,
	}
}

// Execute executes the request
//  @return EnvironmentRelease
func (a *EnvironmentReleasesBetaApiService) GetEnvironmentReleaseExecute(r ApiGetEnvironmentReleaseRequest) (*EnvironmentRelease, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentRelease
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentReleasesBetaApiService.GetEnvironmentRelease")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/releases/{releaseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"flagKey"+"}", url.PathEscape(parameterValueToString(r.flagKey, "flagKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", url.PathEscape(parameterValueToString(r.environmentKey, "environmentKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"releaseId"+"}", url.PathEscape(parameterValueToString(r.releaseId, "releaseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	if r.expand != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expand", r.expand, "form", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListEnvironmentReleasesRequest struct {
	ctx context.Context
	ApiService EnvironmentReleasesBetaApi
	lDAPIVersion *string
	projectKey string
	flagKey string
	environmentKey string
	filter *string
	limit *int32
	offset *int32
	snapshot *string
}

// Version of the endpoint.
func (r ApiListEnvironmentReleasesRequest) LDAPIVersion(lDAPIVersion string) ApiListEnvironmentReleasesRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// A queryfilter expression that narrows the list. Supported fields and operators: &#x60;kind&#x60; (&#x60;equals&#x60;) with values &#x60;progressive&#x60; or &#x60;guarded&#x60;. 
func (r ApiListEnvironmentReleasesRequest) Filter(filter string) ApiListEnvironmentReleasesRequest {
	r.filter = &filter
	return r
}

// The number of releases to return. Defaults to 25. Maximum is 100.
func (r ApiListEnvironmentReleasesRequest) Limit(limit int32) ApiListEnvironmentReleasesRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten releases and then returns the next releases in the list, up to the query &#x60;limit&#x60;.
func (r ApiListEnvironmentReleasesRequest) Offset(offset int32) ApiListEnvironmentReleasesRequest {
	r.offset = &offset
	return r
}

// Pins the read to a moment in time (nanoseconds since the Unix epoch) so pages walked in sequence reflect one consistent snapshot. Follow the pagination links rather than constructing this value. Out-of-range values are ignored.
func (r ApiListEnvironmentReleasesRequest) Snapshot(snapshot string) ApiListEnvironmentReleasesRequest {
	r.snapshot = &snapshot
	return r
}

func (r ApiListEnvironmentReleasesRequest) Execute() (*EnvironmentReleaseCollection, *http.Response, error) {
	return r.ApiService.ListEnvironmentReleasesExecute(r)
}

/*
ListEnvironmentReleases List environment releases

Returns the environment releases for the flag in the environment, newest first.
Each release includes its phases in phase-index order.
Metric definitions are not included. To get them, call
`getEnvironmentRelease` with `expand=metrics`.
Returns an empty `items` array when the flag has no releases in the environment,
or when no release matches the `filter`.
Releases that ended more than 90 days ago are not returned.

Results are paginated. By default, the first 25 releases are returned. Use
`limit` and `offset` to page through the rest, and `totalCount` to learn how
many releases match in total. The `_links` object includes `first`, `prev`,
`next`, and `last` links where they apply.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param flagKey
 @param environmentKey
 @return ApiListEnvironmentReleasesRequest
*/
func (a *EnvironmentReleasesBetaApiService) ListEnvironmentReleases(ctx context.Context, projectKey string, flagKey string, environmentKey string) ApiListEnvironmentReleasesRequest {
	return ApiListEnvironmentReleasesRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		flagKey: flagKey,
		environmentKey: environmentKey,
	}
}

// Execute executes the request
//  @return EnvironmentReleaseCollection
func (a *EnvironmentReleasesBetaApiService) ListEnvironmentReleasesExecute(r ApiListEnvironmentReleasesRequest) (*EnvironmentReleaseCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentReleaseCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentReleasesBetaApiService.ListEnvironmentReleases")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/releases"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"flagKey"+"}", url.PathEscape(parameterValueToString(r.flagKey, "flagKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", url.PathEscape(parameterValueToString(r.environmentKey, "environmentKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
        var defaultValue int32 = 25
        parameterAddToHeaderOrQuery(localVarQueryParams, "limit", defaultValue, "form", "")
        r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	} else {
        var defaultValue int32 = 0
        parameterAddToHeaderOrQuery(localVarQueryParams, "offset", defaultValue, "form", "")
        r.offset = &defaultValue
	}
	if r.snapshot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "_snapshot", r.snapshot, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStartEnvironmentReleaseRequest struct {
	ctx context.Context
	ApiService EnvironmentReleasesBetaApi
	lDAPIVersion *string
	projectKey string
	flagKey string
	environmentKey string
	startEnvironmentReleaseRequest *StartEnvironmentReleaseRequest
}

// Version of the endpoint.
func (r ApiStartEnvironmentReleaseRequest) LDAPIVersion(lDAPIVersion string) ApiStartEnvironmentReleaseRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiStartEnvironmentReleaseRequest) StartEnvironmentReleaseRequest(startEnvironmentReleaseRequest StartEnvironmentReleaseRequest) ApiStartEnvironmentReleaseRequest {
	r.startEnvironmentReleaseRequest = &startEnvironmentReleaseRequest
	return r
}

func (r ApiStartEnvironmentReleaseRequest) Execute() (*EnvironmentRelease, *http.Response, error) {
	return r.ApiService.StartEnvironmentReleaseExecute(r)
}

/*
StartEnvironmentRelease Start an environment release

Creates and starts a release on a flag in an environment, either from
inline configuration or from an environment release plan. A release
has one or more phases. Each phase is an audience, targeted by one
flag rule, and has its own steps: the schedule of allocations at
which that rule serves the target variation. After all of a phase's
steps complete, the phase serves the target variation to 100% of its
contexts.

**Inline configuration.** Each entry in `phases` must specify exactly
one targeting mode:

- **Existing rule**: set `ruleId`. The rule must be enabled. Only a
  single-phase release can use an existing rule.
- **New rule**: set `clauses`, with an optional `description`. New
  rules are appended to the end of the flag's rule list. To insert
  them before an existing rule instead, set the top-level
  `beforeRuleId`, which applies to every phase.

Every phase's rule, including an existing rule passed by `ruleId`,
must include
`{"contextKind": "<randomizationUnit>", "attribute": "kind", "op": "in", "values": ["<randomizationUnit>"]}`,
not negated and with no other values, so the release only reaches
contexts of the `randomizationUnit` kind, or multi-contexts
containing the `randomizationUnit` kind. An existing rule needs the
clause even if its other clauses only match that context kind. A
phase without it returns `400`.

For multiple phases, set `phases` with at least two entries. Each phase
creates a rule. The first rule is enabled and starts immediately.
Subsequent rules stay disabled until the release advances to them:
automatically when `advancementMode` is `automatic`, or by an explicit
advance when it is `manual`. `advancementMode` is required for a
multi-phase release.

Releases on the flag's default rule are not supported. A phase that
omits both targeting modes is rejected.

The flag must be on in the environment, or the request returns
`409`. The release only reaches contexts that match a phase's rule.
Other contexts keep getting the flag's default rule, which may
already serve the target variation.

The release type is inferred from `metrics`: an omitted or empty array
creates a progressive release; one or more metrics creates a guarded
release. Metrics must be listed individually; metric groups are not
accepted.

Every start runs release health checks, described in
`ReleaseHealthCheckKind`. A flag that hasn't recently been
evaluated for contexts of the `randomizationUnit` kind fails
`flag_evaluations_with_context`, and one with no recent evaluations
at all also fails `flag_evaluations_any`. Guarded releases also
check that each metric receives events. A failing
check returns `400` with code `health_checks_failed`. To start the
release anyway, list the failing kinds in `bypassHealthChecks`.

**From a release plan.** Set `releasePlanKey` to start from a plan in
this environment instead of configuring the release inline; a request
that also sets fields the plan supplies is rejected. The plan's current
version runs as stored, with metric groups expanded into their members.
Health checks run against the plan's randomization unit and metrics.
Returns `400` if the plan is missing or incomplete, and `422` until
environment release plans are available for the account.

Requires the `updateRules` action on the flag in this environment, plus
`updateFallthrough` when `cleanupBehavior` is
`remove_rules_and_update_default_rule`. A release started from a plan
requires both. The start is recorded in the flag's audit log as a
targeting change that carries `comment`.

In an environment that requires approvals, this endpoint fails with
`405`. Tokens granted the `bypassRequiredApproval` action can still
call it to start the release directly.

This operation is not idempotent. Retrying a request that already
started a release returns `409` with code `release_already_active`.
The earlier request may have succeeded; call `listEnvironmentReleases`
to confirm the running release is yours.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param flagKey
 @param environmentKey
 @return ApiStartEnvironmentReleaseRequest
*/
func (a *EnvironmentReleasesBetaApiService) StartEnvironmentRelease(ctx context.Context, projectKey string, flagKey string, environmentKey string) ApiStartEnvironmentReleaseRequest {
	return ApiStartEnvironmentReleaseRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		flagKey: flagKey,
		environmentKey: environmentKey,
	}
}

// Execute executes the request
//  @return EnvironmentRelease
func (a *EnvironmentReleasesBetaApiService) StartEnvironmentReleaseExecute(r ApiStartEnvironmentReleaseRequest) (*EnvironmentRelease, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentRelease
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentReleasesBetaApiService.StartEnvironmentRelease")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/releases"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"flagKey"+"}", url.PathEscape(parameterValueToString(r.flagKey, "flagKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", url.PathEscape(parameterValueToString(r.environmentKey, "environmentKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.startEnvironmentReleaseRequest == nil {
		return localVarReturnValue, nil, reportError("startEnvironmentReleaseRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.startEnvironmentReleaseRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v StartEnvironmentReleaseBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStopEnvironmentReleaseRequest struct {
	ctx context.Context
	ApiService EnvironmentReleasesBetaApi
	lDAPIVersion *string
	projectKey string
	flagKey string
	environmentKey string
	releaseId string
	stopEnvironmentReleaseRequest *StopEnvironmentReleaseRequest
}

// Version of the endpoint.
func (r ApiStopEnvironmentReleaseRequest) LDAPIVersion(lDAPIVersion string) ApiStopEnvironmentReleaseRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiStopEnvironmentReleaseRequest) StopEnvironmentReleaseRequest(stopEnvironmentReleaseRequest StopEnvironmentReleaseRequest) ApiStopEnvironmentReleaseRequest {
	r.stopEnvironmentReleaseRequest = &stopEnvironmentReleaseRequest
	return r
}

func (r ApiStopEnvironmentReleaseRequest) Execute() (*EnvironmentRelease, *http.Response, error) {
	return r.ApiService.StopEnvironmentReleaseExecute(r)
}

/*
StopEnvironmentRelease Stop an environment release

Stops a release and resolves its phases according to the `resolution`
you specify. To learn how each resolution treats completed and pending
phases, read the `resolution` schema.

Requires the `updateRules` action on the flag in this environment, or
`updateFallthrough` for a historical release that ran on the default
rule. The stop is recorded in the flag's audit log as a targeting
change that names the resolution and carries `comment`.

In an environment that requires approvals, this endpoint fails with
`405`. Tokens granted the `bypassRequiredApproval` action can still
call it to stop the release directly.

Stopping a release that has already completed or been stopped returns
`200` and makes no changes. Returns `404` for a release that ended more
than 90 days ago.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param flagKey
 @param environmentKey
 @param releaseId
 @return ApiStopEnvironmentReleaseRequest
*/
func (a *EnvironmentReleasesBetaApiService) StopEnvironmentRelease(ctx context.Context, projectKey string, flagKey string, environmentKey string, releaseId string) ApiStopEnvironmentReleaseRequest {
	return ApiStopEnvironmentReleaseRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		flagKey: flagKey,
		environmentKey: environmentKey,
		releaseId: releaseId,
	}
}

// Execute executes the request
//  @return EnvironmentRelease
func (a *EnvironmentReleasesBetaApiService) StopEnvironmentReleaseExecute(r ApiStopEnvironmentReleaseRequest) (*EnvironmentRelease, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentRelease
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentReleasesBetaApiService.StopEnvironmentRelease")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/releases/{releaseId}/stop"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"flagKey"+"}", url.PathEscape(parameterValueToString(r.flagKey, "flagKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", url.PathEscape(parameterValueToString(r.environmentKey, "environmentKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"releaseId"+"}", url.PathEscape(parameterValueToString(r.releaseId, "releaseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.stopEnvironmentReleaseRequest == nil {
		return localVarReturnValue, nil, reportError("stopEnvironmentReleaseRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.stopEnvironmentReleaseRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
