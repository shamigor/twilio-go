/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSyncListPermissionResponse struct for ListSyncListPermissionResponse
type ListSyncListPermissionResponse struct {
	Meta        ListServiceResponseMeta                   `json:"meta,omitempty"`
	Permissions []SyncV1ServiceSyncListSyncListPermission `json:"permissions,omitempty"`
}
