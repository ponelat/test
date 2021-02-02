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

type PluginConfigurationUpdateOptions struct {

	// Available options PLAN = disable plugins based on the current owner's plan. See plan details for more info on specific plugins allowed per plan 
	Disable string `json:"disable,omitempty"`
}