package flip

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// Dataset is a set of rules that define attributes over a distinct set of members
type Dataset struct {
	Identifier   string      `mapstructure:"identifier" json:"identifier"`
	Name         string      `mapstructure:"name" json:"name"`
	OutputEngine string      `mapstructure:"output_engine" json:"output_engine"`
	Description  string      `mapstructure:"description" json:"description"`
	Tags         []string    `mapstructure:"tags" json:"tags"`
	Partitions   []Partition `mapstructure:"partitions" json:"partitions"`
}

// GetDatasetsResponse returns the RPC response
type GetDatasetsResponse struct {
	Datasets []Dataset `mapstructure:"datasets"`
}

// GetDatasets returns the partitions belonging to a member
func (c Client) GetDatasets(entityID string, ownerID string) (*GetDatasetsResponse, error) {
	var input = make(map[string]interface{})
	if entityID != "" {
		input["entity_id"] = entityID
	}

	if ownerID != "" {
		input["owner_id"] = ownerID
	}

	var response GetDatasetsResponse

	rpc, err := c.CallRPC("RPCService.GetDatasets", input)
	if err != nil {
		return &response, err
	}

	err = mapstructure.Decode(rpc.Result, &response)
	if err != nil {
		return &response, errors.Wrap(err, "error decoding into `GetDatasets`")
	}

	return &response, nil
}
