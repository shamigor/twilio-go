/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV1VoiceVoiceNumber struct for PricingV1VoiceVoiceNumber
type PricingV1VoiceVoiceNumber struct {
	// The name of the country
	Country          *string                                    `json:"country,omitempty"`
	InboundCallPrice *PricingV1VoiceVoiceNumberInboundCallPrice `json:"inbound_call_price,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The phone number
	Number            *string                                     `json:"number,omitempty"`
	OutboundCallPrice *PricingV1VoiceVoiceNumberOutboundCallPrice `json:"outbound_call_price,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"price_unit,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
