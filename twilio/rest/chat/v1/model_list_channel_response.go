/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListChannelResponse struct for ListChannelResponse
type ListChannelResponse struct {
	Channels []ChatV1ServiceChannel     `json:"Channels,omitempty"`
	Meta     ListCredentialResponseMeta `json:"Meta,omitempty"`
}
