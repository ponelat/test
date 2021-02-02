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

// details for one side of a ComparisonPart
type ComparisonDetail struct {

	StartPosition *Position `json:"startPosition"`

	EndPosition *Position `json:"endPosition"`

	// the content of the API definition between `startPosition` and `endPosition`
	Content string `json:"content"`

	// the path to the object (notation similar to JSONPath)
	Path string `json:"path,omitempty"`
}