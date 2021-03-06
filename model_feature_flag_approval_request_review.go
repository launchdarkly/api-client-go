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

type FeatureFlagApprovalRequestReview struct {
	// A unix epoch time in milliseconds specifying the date the approval request was reviewed
	CreationDate int32 `json:"creationDate,omitempty"`
	Kind *FeatureFlagApprovalRequestReviewStatus `json:"kind,omitempty"`
	// The unique resource id.
	MemberId string `json:"memberId,omitempty"`
	// The unique resource id.
	Id string `json:"_id,omitempty"`
}
