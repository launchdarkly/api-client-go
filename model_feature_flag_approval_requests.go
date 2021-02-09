/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 4.0.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlagApprovalRequests struct {
	Links *Links `json:"_links,omitempty"`
	Items []FeatureFlagApprovalRequest `json:"items,omitempty"`
}