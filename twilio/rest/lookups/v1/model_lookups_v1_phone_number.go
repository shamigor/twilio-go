/*
 * Twilio - Lookups
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// LookupsV1PhoneNumber struct for LookupsV1PhoneNumber
type LookupsV1PhoneNumber struct {
	AddOns         *map[string]interface{} `json:"AddOns,omitempty"`
	CallerName     *map[string]interface{} `json:"CallerName,omitempty"`
	Carrier        *map[string]interface{} `json:"Carrier,omitempty"`
	CountryCode    *string                 `json:"CountryCode,omitempty"`
	NationalFormat *string                 `json:"NationalFormat,omitempty"`
	PhoneNumber    *string                 `json:"PhoneNumber,omitempty"`
	Url            *string                 `json:"Url,omitempty"`
}
