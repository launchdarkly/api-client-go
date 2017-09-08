/* 
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@launchdarkly.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type Webhook struct {

	Links Links `json:"_links,omitempty"`

	Id string `json:"_id,omitempty"`

	Url string `json:"url,omitempty"`

	Secret string `json:"secret,omitempty"`

	On bool `json:"on,omitempty"`

	Tags []string `json:"tags,omitempty"`
}
