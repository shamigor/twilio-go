/*
 * Twilio - Accounts
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCredentialAwsResponse struct for ListCredentialAwsResponse
type ListCredentialAwsResponse struct {
	Credentials []AccountsV1CredentialCredentialAws `json:"Credentials,omitempty"`
	Meta        ListCredentialAwsResponseMeta       `json:"Meta,omitempty"`
}
