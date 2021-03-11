/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics struct for TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics
type TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics struct {
	AccountSid   *string                 `json:"AccountSid,omitempty"`
	Cumulative   *map[string]interface{} `json:"Cumulative,omitempty"`
	Realtime     *map[string]interface{} `json:"Realtime,omitempty"`
	TaskQueueSid *string                 `json:"TaskQueueSid,omitempty"`
	Url          *string                 `json:"Url,omitempty"`
	WorkspaceSid *string                 `json:"WorkspaceSid,omitempty"`
}
