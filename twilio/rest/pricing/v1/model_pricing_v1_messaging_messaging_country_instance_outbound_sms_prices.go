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

// PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices struct for PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices
type PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices struct {
	Carrier string                                                       `json:"Carrier,omitempty"`
	Mcc     string                                                       `json:"Mcc,omitempty"`
	Mnc     string                                                       `json:"Mnc,omitempty"`
	Prices  []PricingV1MessagingMessagingCountryInstanceInboundSmsPrices `json:"Prices,omitempty"`
}