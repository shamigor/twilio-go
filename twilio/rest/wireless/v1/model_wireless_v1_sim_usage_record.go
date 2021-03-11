/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// WirelessV1SimUsageRecord struct for WirelessV1SimUsageRecord
type WirelessV1SimUsageRecord struct {
	AccountSid *string                 `json:"AccountSid,omitempty"`
	Commands   *map[string]interface{} `json:"Commands,omitempty"`
	Data       *map[string]interface{} `json:"Data,omitempty"`
	Period     *map[string]interface{} `json:"Period,omitempty"`
	SimSid     *string                 `json:"SimSid,omitempty"`
}
