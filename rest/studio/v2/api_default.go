/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL        string
	requestHandler *twilio.RequestHandler
}

func NewDefaultApiService(requestHandler *twilio.RequestHandler) *DefaultApiService {
	return &DefaultApiService{
		requestHandler: requestHandler,
		baseURL:        "https://studio.twilio.com",
	}
}

func NewDefaultApiServiceWithClient(client twilio.BaseClient) *DefaultApiService {
	return NewDefaultApiService(twilio.NewRequestHandler(client))
}

// Optional parameters for the method 'CreateExecution'
type CreateExecutionParams struct {
	// The Twilio phone number to send messages or initiate calls from during the Flow's Execution. Available as variable `{{flow.channel.address}}`. For SMS, this can also be a Messaging Service SID.
	From *string `json:"From,omitempty"`
	// JSON data that will be added to the Flow's context and that can be accessed as variables inside your Flow. For example, if you pass in `Parameters={\\\"name\\\":\\\"Zeke\\\"}`, a widget in your Flow can reference the variable `{{flow.data.name}}`, which returns \\\"Zeke\\\". Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
	Parameters *map[string]interface{} `json:"Parameters,omitempty"`
	// The Contact phone number to start a Studio Flow Execution, available as variable `{{contact.channel.address}}`.
	To *string `json:"To,omitempty"`
}

func (params *CreateExecutionParams) SetFrom(From string) *CreateExecutionParams {
	params.From = &From
	return params
}
func (params *CreateExecutionParams) SetParameters(Parameters map[string]interface{}) *CreateExecutionParams {
	params.Parameters = &Parameters
	return params
}
func (params *CreateExecutionParams) SetTo(To string) *CreateExecutionParams {
	params.To = &To
	return params
}

// Triggers a new Execution for the Flow
func (c *DefaultApiService) CreateExecution(FlowSid string, params *CreateExecutionParams) (*StudioV2FlowExecution, error) {
	path := "/v2/Flows/{FlowSid}/Executions"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Parameters != nil {
		v, err := json.Marshal(params.Parameters)

		if err != nil {
			return nil, err
		}

		data.Set("Parameters", string(v))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowExecution{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'CreateFlow'
type CreateFlowParams struct {
	// Description of change made in the revision.
	CommitMessage *string `json:"CommitMessage,omitempty"`
	// JSON representation of flow definition.
	Definition *map[string]interface{} `json:"Definition,omitempty"`
	// The string that you assigned to describe the Flow.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The status of the Flow. Can be: `draft` or `published`.
	Status *string `json:"Status,omitempty"`
}

func (params *CreateFlowParams) SetCommitMessage(CommitMessage string) *CreateFlowParams {
	params.CommitMessage = &CommitMessage
	return params
}
func (params *CreateFlowParams) SetDefinition(Definition map[string]interface{}) *CreateFlowParams {
	params.Definition = &Definition
	return params
}
func (params *CreateFlowParams) SetFriendlyName(FriendlyName string) *CreateFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateFlowParams) SetStatus(Status string) *CreateFlowParams {
	params.Status = &Status
	return params
}

// Create a Flow.
func (c *DefaultApiService) CreateFlow(params *CreateFlowParams) (*StudioV2Flow, error) {
	path := "/v2/Flows"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CommitMessage != nil {
		data.Set("CommitMessage", *params.CommitMessage)
	}
	if params != nil && params.Definition != nil {
		v, err := json.Marshal(params.Definition)

		if err != nil {
			return nil, err
		}

		data.Set("Definition", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete the Execution and all Steps relating to it.
func (c *DefaultApiService) DeleteExecution(FlowSid string, Sid string) error {
	path := "/v2/Flows/{FlowSid}/Executions/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Delete a specific Flow.
func (c *DefaultApiService) DeleteFlow(Sid string) error {
	path := "/v2/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Retrieve an Execution
func (c *DefaultApiService) FetchExecution(FlowSid string, Sid string) (*StudioV2FlowExecution, error) {
	path := "/v2/Flows/{FlowSid}/Executions/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowExecution{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Retrieve the most recent context for an Execution.
func (c *DefaultApiService) FetchExecutionContext(FlowSid string, ExecutionSid string) (*StudioV2FlowExecutionExecutionContext, error) {
	path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Context"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowExecutionExecutionContext{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Retrieve a Step.
func (c *DefaultApiService) FetchExecutionStep(FlowSid string, ExecutionSid string, Sid string) (*StudioV2FlowExecutionExecutionStep, error) {
	path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowExecutionExecutionStep{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Retrieve the context for an Execution Step.
func (c *DefaultApiService) FetchExecutionStepContext(FlowSid string, ExecutionSid string, StepSid string) (*StudioV2FlowExecutionExecutionStepExecutionStepContext, error) {
	path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)
	path = strings.Replace(path, "{"+"StepSid"+"}", StepSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowExecutionExecutionStepExecutionStepContext{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Retrieve a specific Flow.
func (c *DefaultApiService) FetchFlow(Sid string) (*StudioV2Flow, error) {
	path := "/v2/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Retrieve a specific Flow revision.
func (c *DefaultApiService) FetchFlowRevision(Sid string, Revision string) (*StudioV2FlowFlowRevision, error) {
	path := "/v2/Flows/{Sid}/Revisions/{Revision}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)
	path = strings.Replace(path, "{"+"Revision"+"}", Revision, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowFlowRevision{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch flow test users
func (c *DefaultApiService) FetchTestUser(Sid string) (*StudioV2FlowTestUser, error) {
	path := "/v2/Flows/{Sid}/TestUsers"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowTestUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListExecution'
type ListExecutionParams struct {
	// Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
	DateCreatedFrom *time.Time `json:"DateCreatedFrom,omitempty"`
	// Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
	DateCreatedTo *time.Time `json:"DateCreatedTo,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListExecutionParams) SetDateCreatedFrom(DateCreatedFrom time.Time) *ListExecutionParams {
	params.DateCreatedFrom = &DateCreatedFrom
	return params
}
func (params *ListExecutionParams) SetDateCreatedTo(DateCreatedTo time.Time) *ListExecutionParams {
	params.DateCreatedTo = &DateCreatedTo
	return params
}
func (params *ListExecutionParams) SetPageSize(PageSize int32) *ListExecutionParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Executions for the Flow.
func (c *DefaultApiService) ListExecution(FlowSid string, params *ListExecutionParams) (*ListExecutionResponse, error) {
	path := "/v2/Flows/{FlowSid}/Executions"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreatedFrom != nil {
		data.Set("DateCreatedFrom", fmt.Sprint((*params.DateCreatedFrom).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedTo != nil {
		data.Set("DateCreatedTo", fmt.Sprint((*params.DateCreatedTo).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExecutionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListExecutionStep'
type ListExecutionStepParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListExecutionStepParams) SetPageSize(PageSize int32) *ListExecutionStepParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Steps for an Execution.
func (c *DefaultApiService) ListExecutionStep(FlowSid string, ExecutionSid string, params *ListExecutionStepParams) (*ListExecutionStepResponse, error) {
	path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExecutionStepResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlow'
type ListFlowParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListFlowParams) SetPageSize(PageSize int32) *ListFlowParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Flows.
func (c *DefaultApiService) ListFlow(params *ListFlowParams) (*ListFlowResponse, error) {
	path := "/v2/Flows"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlowRevision'
type ListFlowRevisionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListFlowRevisionParams) SetPageSize(PageSize int32) *ListFlowRevisionParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Flows revisions.
func (c *DefaultApiService) ListFlowRevision(Sid string, params *ListFlowRevisionParams) (*ListFlowRevisionResponse, error) {
	path := "/v2/Flows/{Sid}/Revisions"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowRevisionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateExecution'
type UpdateExecutionParams struct {
	// The status of the Execution. Can only be `ended`.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateExecutionParams) SetStatus(Status string) *UpdateExecutionParams {
	params.Status = &Status
	return params
}

// Update the status of an Execution to &#x60;ended&#x60;.
func (c *DefaultApiService) UpdateExecution(FlowSid string, Sid string, params *UpdateExecutionParams) (*StudioV2FlowExecution, error) {
	path := "/v2/Flows/{FlowSid}/Executions/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowExecution{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateFlow'
type UpdateFlowParams struct {
	// Description of change made in the revision.
	CommitMessage *string `json:"CommitMessage,omitempty"`
	// JSON representation of flow definition.
	Definition *map[string]interface{} `json:"Definition,omitempty"`
	// The string that you assigned to describe the Flow.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The status of the Flow. Can be: `draft` or `published`.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateFlowParams) SetCommitMessage(CommitMessage string) *UpdateFlowParams {
	params.CommitMessage = &CommitMessage
	return params
}
func (params *UpdateFlowParams) SetDefinition(Definition map[string]interface{}) *UpdateFlowParams {
	params.Definition = &Definition
	return params
}
func (params *UpdateFlowParams) SetFriendlyName(FriendlyName string) *UpdateFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateFlowParams) SetStatus(Status string) *UpdateFlowParams {
	params.Status = &Status
	return params
}

// Update a Flow.
func (c *DefaultApiService) UpdateFlow(Sid string, params *UpdateFlowParams) (*StudioV2Flow, error) {
	path := "/v2/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CommitMessage != nil {
		data.Set("CommitMessage", *params.CommitMessage)
	}
	if params != nil && params.Definition != nil {
		v, err := json.Marshal(params.Definition)

		if err != nil {
			return nil, err
		}

		data.Set("Definition", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateFlowValidate'
type UpdateFlowValidateParams struct {
	// Description of change made in the revision.
	CommitMessage *string `json:"CommitMessage,omitempty"`
	// JSON representation of flow definition.
	Definition *map[string]interface{} `json:"Definition,omitempty"`
	// The string that you assigned to describe the Flow.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The status of the Flow. Can be: `draft` or `published`.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateFlowValidateParams) SetCommitMessage(CommitMessage string) *UpdateFlowValidateParams {
	params.CommitMessage = &CommitMessage
	return params
}
func (params *UpdateFlowValidateParams) SetDefinition(Definition map[string]interface{}) *UpdateFlowValidateParams {
	params.Definition = &Definition
	return params
}
func (params *UpdateFlowValidateParams) SetFriendlyName(FriendlyName string) *UpdateFlowValidateParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateFlowValidateParams) SetStatus(Status string) *UpdateFlowValidateParams {
	params.Status = &Status
	return params
}

// Validate flow JSON definition
func (c *DefaultApiService) UpdateFlowValidate(params *UpdateFlowValidateParams) (*StudioV2FlowValidate, error) {
	path := "/v2/Flows/Validate"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CommitMessage != nil {
		data.Set("CommitMessage", *params.CommitMessage)
	}
	if params != nil && params.Definition != nil {
		v, err := json.Marshal(params.Definition)

		if err != nil {
			return nil, err
		}

		data.Set("Definition", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowValidate{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateTestUser'
type UpdateTestUserParams struct {
	// List of test user identities that can test draft versions of the flow.
	TestUsers *[]string `json:"TestUsers,omitempty"`
}

func (params *UpdateTestUserParams) SetTestUsers(TestUsers []string) *UpdateTestUserParams {
	params.TestUsers = &TestUsers
	return params
}

// Update flow test users
func (c *DefaultApiService) UpdateTestUser(Sid string, params *UpdateTestUserParams) (*StudioV2FlowTestUser, error) {
	path := "/v2/Flows/{Sid}/TestUsers"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.TestUsers != nil {
		data.Set("TestUsers", strings.Join(*params.TestUsers, ","))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowTestUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
