/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListParticipantResponse struct for ListParticipantResponse
type ListParticipantResponse struct {
	End             int32                                  `json:"End,omitempty"`
	FirstPageUri    string                                 `json:"FirstPageUri,omitempty"`
	NextPageUri     string                                 `json:"NextPageUri,omitempty"`
	Page            int32                                  `json:"Page,omitempty"`
	PageSize        int32                                  `json:"PageSize,omitempty"`
	Participants    []ApiV2010AccountConferenceParticipant `json:"Participants,omitempty"`
	PreviousPageUri string                                 `json:"PreviousPageUri,omitempty"`
	Start           int32                                  `json:"Start,omitempty"`
	Uri             string                                 `json:"Uri,omitempty"`
}
