package flip

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// RefreshArgs are used as filters to determine what to refresh
type RefreshArgs struct {
	EntityID  string   `mapstructure:"entity_id" json:"entity_id"`
	MetricID  string   `mapstructure:"metric_id" json:"metric_id"`
	DatasetID string   `mapstructure:"dataset_id" json:"dataset_id"`
	Tags      []string `mapstructure:"tags" json:"tags"`
	Stage     string   `mapstructure:"stage" json:"stage"`
}

// RefreshResponse returns the RPC response
type RefreshResponse struct {
	Message     string      `mapstructure:"message"`
	Status      string      `mapstructure:"status"`
	JobID       string      `mapstructure:"job_id"`
	RefreshArgs RefreshArgs `mapstructure:"refresh_args"`
}

// Refresh creates a refresh job
func (c Client) Refresh(refreshArgs RefreshArgs) (*RefreshResponse, error) {
	var input = make(map[string]interface{})
	if refreshArgs.EntityID != "" {
		input["entity_id"] = refreshArgs.EntityID
	}
	if refreshArgs.MetricID != "" {
		input["metric_id"] = refreshArgs.MetricID
	}
	if refreshArgs.DatasetID != "" {
		input["dataset_id"] = refreshArgs.DatasetID
	}
	if refreshArgs.Stage != "" {
		input["stage"] = refreshArgs.Stage
	}
	if len(refreshArgs.Tags) > 0 {
		input["tags"] = refreshArgs.Tags
	}

	var response RefreshResponse

	rpc, err := c.CallRPC("RPCService.Refresh", input)
	if err != nil {
		return &response, err
	}

	err = mapstructure.Decode(rpc.Result, &response)
	if err != nil {
		return &response, errors.Wrap(err, "error decoding into `RefreshResponse`")
	}

	return &response, nil
}

// GetRefreshJob returns a refresh job
func (c Client) GetRefreshJob(jobID string) (*RefreshResponse, error) {
	var input = make(map[string]interface{})
	input["job_id"] = jobID

	var response RefreshResponse

	rpc, err := c.CallRPC("RPCService.GetRefreshJob", input)
	if err != nil {
		return &response, err
	}

	err = mapstructure.Decode(rpc.Result, &response)
	if err != nil {
		return &response, errors.Wrap(err, "error decoding into `GetRefreshJob`")
	}

	return &response, nil
}
