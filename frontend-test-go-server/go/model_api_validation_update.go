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

// Contains information required to update an APIs validation status
type ApiValidationUpdate struct {

	HasValidationErrors bool `json:"hasValidationErrors,omitempty"`

	ValidationErrorsSeverity []string `json:"validationErrorsSeverity,omitempty"`
}