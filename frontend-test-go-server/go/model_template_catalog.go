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

type TemplateCatalog struct {

	Templates []Template `json:"templates,omitempty"`

	Default_ *Template `json:"default,omitempty"`
}
