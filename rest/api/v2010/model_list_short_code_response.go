/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListShortCodeResponse struct for ListShortCodeResponse
type ListShortCodeResponse struct {
	End             int32                      `json:"end,omitempty"`
	FirstPageUri    string                     `json:"first_page_uri,omitempty"`
	NextPageUri     string                     `json:"next_page_uri,omitempty"`
	Page            int32                      `json:"page,omitempty"`
	PageSize        int32                      `json:"page_size,omitempty"`
	PreviousPageUri string                     `json:"previous_page_uri,omitempty"`
	ShortCodes      []ApiV2010AccountShortCode `json:"short_codes,omitempty"`
	Start           int32                      `json:"start,omitempty"`
	Uri             string                     `json:"uri,omitempty"`
}
