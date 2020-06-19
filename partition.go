package flip

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// Partition is a slice of a dataset corresponding to a particular member attribute
type Partition struct {
	ID          string `mapstructure:"id" json:"id"`
	EntityID    string `mapstructure:"entity_id" json:"entity_id"`
	OwnerID     string `mapstructure:"owner_id" json:"owner_id"`
	Name        string `mapstructure:"name" json:"name"`
	IsRanked    bool   `mapstructure:"is_ranked" json:"is_ranked"`
	Cardinality int    `mapstructure:"cardinality" json:"cardinality"`
	MaxRank     int    `mapstructure:"max_rank" json:"max_rank"`
	MinRank     int    `mapstructure:"min_rank" json:"min_rank"`
}

// GetMemberPartitionsResponse returns the RPC response
type GetMemberPartitionsResponse struct {
	PartitionCount int         `mapstructure:"partition_count"`
	Partitions     []Partition `mapstructure:"partitions"`
}

// GetMemberPartitions returns the partitions belonging to a member
func (c Client) GetMemberPartitions(entityID string, memberID string) (*GetMemberPartitionsResponse, error) {
	var input = make(map[string]interface{})
	input["entity_id"] = entityID
	input["member_id"] = memberID

	var response GetMemberPartitionsResponse

	rpc, err := c.CallRPC("RPCService.GetMemberPartitions", input)
	if err != nil {
		return &response, err
	}

	err = mapstructure.Decode(rpc.Result, &response)
	if err != nil {
		return &response, errors.Wrap(err, "error decoding into `GetMemberPartitions`")
	}

	return &response, nil
}
