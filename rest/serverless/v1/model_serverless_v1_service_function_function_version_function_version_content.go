/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent struct for ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent
type ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent struct {
	// The SID of the Account that created the Function Version resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The content of the Function Version resource
	Content *string `json:"content,omitempty"`
	// The SID of the Function that is the parent of the Function Version
	FunctionSid *string `json:"function_sid,omitempty"`
	// The SID of the Service that the Function Version resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Function Version resource
	Sid *string `json:"sid,omitempty"`
	Url *string `json:"url,omitempty"`
}
