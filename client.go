package agentkata

import (
	"context"
	"net/http"
	"strings"

	"github.com/agentkata/sdk-go/generated"
)

// RequestMeta describes optional solver metadata sent with action and submit requests.
type RequestMeta = generated.ExecutionMeta

// TaskActionInput is the input for a standalone task action call.
type TaskActionInput struct {
	TaskID  string
	Action  string
	Payload map[string]any
	Meta    *RequestMeta
}

// SubmitTaskInput is the input for a standalone task submission.
type SubmitTaskInput struct {
	TaskID string
	Answer any
	Meta   *RequestMeta
}

// TrackTaskActionInput is the input for a tracked task action call.
type TrackTaskActionInput struct {
	TrackID string
	TaskID  string
	Action  string
	Payload map[string]any
	Meta    *RequestMeta
}

// SubmitTrackTaskInput is the input for a tracked task submission.
type SubmitTrackTaskInput struct {
	TrackID string
	TaskID  string
	Answer  any
	Meta    *RequestMeta
}

// Client provides a small ergonomic wrapper over the generated solver client.
type Client struct {
	api      *generated.APIClient
	apiToken string
}

// NewClient constructs a solver API client.
func NewClient(baseURL, apiToken string, httpClient *http.Client) *Client {
	cfg := generated.NewConfiguration()
	cfg.Servers = generated.ServerConfigurations{{URL: normalizeBaseURL(baseURL)}}
	cfg.HTTPClient = httpClient

	return &Client{
		api:      generated.NewAPIClient(cfg),
		apiToken: apiToken,
	}
}

// Health checks solver API availability.
func (c *Client) Health(ctx context.Context) (*generated.HealthResponse, error) {
	result, _, err := c.api.SolverAPI.GetHealth(withAuth(ctx, c.apiToken)).Execute()
	return result, err
}

// TaskAction executes one standalone task action.
func (c *Client) TaskAction(ctx context.Context, input TaskActionInput) (*generated.ActionEnvelope, error) {
	request := generated.NewActionRequest()
	request.Params = input.Payload
	if input.Meta != nil {
		request.Meta = input.Meta
	}

	result, _, err := c.api.SolverAPI.
		TaskAction(withAuth(ctx, c.apiToken), input.TaskID, input.Action).
		ActionRequest(*request).
		Execute()
	return result, err
}

// SubmitTask submits one standalone task answer.
func (c *Client) SubmitTask(ctx context.Context, input SubmitTaskInput) (*generated.SubmitEnvelope, error) {
	params := generated.NewSubmitParams(input.Answer)
	request := generated.NewSubmitRequest(*params)
	if input.Meta != nil {
		request.Meta = input.Meta
	}

	result, _, err := c.api.SolverAPI.
		SubmitTask(withAuth(ctx, c.apiToken), input.TaskID).
		SubmitRequest(*request).
		Execute()
	return result, err
}

// RestartTask restarts one standalone task run.
func (c *Client) RestartTask(ctx context.Context, taskID string) (*generated.RestartEnvelope, error) {
	result, _, err := c.api.SolverAPI.RestartTask(withAuth(ctx, c.apiToken), taskID).Execute()
	return result, err
}

// TrackTaskAction executes one tracked task action.
func (c *Client) TrackTaskAction(ctx context.Context, input TrackTaskActionInput) (*generated.ActionEnvelope, error) {
	request := generated.NewActionRequest()
	request.Params = input.Payload
	if input.Meta != nil {
		request.Meta = input.Meta
	}

	result, _, err := c.api.SolverAPI.
		TrackTaskAction(withAuth(ctx, c.apiToken), input.TrackID, input.TaskID, input.Action).
		ActionRequest(*request).
		Execute()
	return result, err
}

// SubmitTrackTask submits one tracked task answer.
func (c *Client) SubmitTrackTask(ctx context.Context, input SubmitTrackTaskInput) (*generated.SubmitEnvelope, error) {
	params := generated.NewSubmitParams(input.Answer)
	request := generated.NewSubmitRequest(*params)
	if input.Meta != nil {
		request.Meta = input.Meta
	}

	result, _, err := c.api.SolverAPI.
		SubmitTrackTask(withAuth(ctx, c.apiToken), input.TrackID, input.TaskID).
		SubmitRequest(*request).
		Execute()
	return result, err
}

// RestartTrack restarts one tracked run.
func (c *Client) RestartTrack(ctx context.Context, trackID string) (*generated.RestartEnvelope, error) {
	result, _, err := c.api.SolverAPI.RestartTrack(withAuth(ctx, c.apiToken), trackID).Execute()
	return result, err
}

func normalizeBaseURL(baseURL string) string {
	trimmed := strings.TrimRight(baseURL, "/")
	if strings.HasSuffix(trimmed, "/api/agent") {
		return trimmed
	}
	return trimmed + "/api/agent"
}

func withAuth(ctx context.Context, apiToken string) context.Context {
	return context.WithValue(ctx, generated.ContextAccessToken, apiToken)
}
