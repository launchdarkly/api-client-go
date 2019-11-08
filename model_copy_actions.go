/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.21
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type CopyActions string

// List of CopyActions
const (
	UPDATE_ON CopyActions = "updateOn"
	UPDATE_PREREQUISITES CopyActions = "updatePrerequisites"
	UPDATE_TARGETS CopyActions = "updateTargets"
	UPDATE_RULES CopyActions = "updateRules"
	UPDATE_FALLTHROUGH CopyActions = "updateFallthrough"
	UPDATE_OFF_VARIATION CopyActions = "updateOffVariation"
)
