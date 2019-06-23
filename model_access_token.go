/*
 * Leap Play
 *
 * Learn more at https://www.leap-play.com
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package leapclient

type AccessToken struct {
	Token string `json:"token,omitempty"`
	ExpiresIn int32 `json:"expiresIn,omitempty"`
}
