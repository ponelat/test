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

type NewComment struct {

	Position int32 `json:"position"`

	Body string `json:"body"`

	Replies []NewReply `json:"replies,omitempty"`
}