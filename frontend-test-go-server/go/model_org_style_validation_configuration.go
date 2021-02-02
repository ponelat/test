/*
 * SwaggerHub Registry API
 *
 * The registry API for SwaggerHub
 *
 * API version: 1.0.97
 * Contact: info@swaggerhub.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OrgStyleValidationConfiguration struct {

	Title string `json:"title,omitempty"`

	Properties map[string]OrgStyleValidationRuleGroup `json:"properties,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	ScanRunning bool `json:"scanRunning,omitempty"`

	Description string `json:"description,omitempty"`

	CustomRules string `json:"customRules,omitempty"`

	RulesInSync bool `json:"rulesInSync,omitempty"`
}