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

type Webhooks struct {

	Links Links `json:"_links,omitempty"`

	Items []Webhook `json:"items,omitempty"`
}
