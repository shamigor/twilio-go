/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListPhoneNumberCountryResponse struct for ListPhoneNumberCountryResponse
type ListPhoneNumberCountryResponse struct {
	Countries []PricingV1PhoneNumberPhoneNumberCountry `json:"Countries,omitempty"`
	Meta      ListMessagingCountryResponseMeta         `json:"Meta,omitempty"`
}
