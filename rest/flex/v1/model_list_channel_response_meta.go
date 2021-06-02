/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListChannelResponseMeta struct for ListChannelResponseMeta
type ListChannelResponseMeta struct {
	FirstPageUrl    string `json:"first_page_url,omitempty"`
	Key             string `json:"key,omitempty"`
	NextPageUrl     string `json:"next_page_url,omitempty"`
	Page            int32  `json:"page,omitempty"`
	PageSize        int32  `json:"page_size,omitempty"`
	PreviousPageUrl string `json:"previous_page_url,omitempty"`
	Url             string `json:"url,omitempty"`
}
