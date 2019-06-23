/*
 * Leap Play
 *
 * Learn more at https://www.leap-play.com
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package leapclient

type StationCurrentState struct {
	StationId string `json:"stationId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ControlMode string `json:"controlMode,omitempty"`
	NetworkState string `json:"networkState,omitempty"`
	Session *Session `json:"session,omitempty"`
}