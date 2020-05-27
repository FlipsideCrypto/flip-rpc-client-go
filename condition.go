package flip

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// Gte = greater than or equal to
type Gte struct {
	PartitionID string  `json:"partition_id"`
	Value       float64 `json:"value"`
}

// Lte = less than or equal to
type Lte struct {
	PartitionID string  `json:"partition_id"`
	Value       float64 `json:"value"`
}

// Lt = less than
type Lt struct {
	PartitionID string  `json:"partition_id"`
	Value       float64 `json:"value"`
}

// Gt = greater than
type Gt struct {
	PartitionID string  `json:"partition_id"`
	Value       float64 `json:"value"`
}

// Condition is set of logic
type Condition struct {
	PartitionID string  `json:"partition_id"`
	Value       float64 `json:"value"`

	Or  []*Condition `json:"or"`
	And []*Condition `json:"and"`
	Gt  Gt           `json:"gt"`
	Gte Gte          `json:"gte"`
	Lt  Lt           `json:"lt"`
	Lte Lte          `json:"lte"`
}

// GetConditionMembersResponse returns the RPC response
type GetConditionMembersResponse struct {
	Members     []string `mapstructure:"members"`
	MemberCount int      `mapstructure:"member_count"`
}

// GetConditionMembers returns the members belonging to the result set of a condition.
func (c Client) GetConditionMembers(condition Condition) (*GetConditionMembersResponse, error) {
	var input = make(map[string]Condition)
	input["condition"] = condition

	var conditionMembers GetConditionMembersResponse

	result, err := c.CallRPC("RPCService.GetConditionMembers", input)
	if err != nil {
		return &conditionMembers, err
	}

	err = mapstructure.Decode(result, &conditionMembers)
	if err != nil {
		return &conditionMembers, errors.Wrap(err, "error decoding into `GetConditionMembersResponse`")
	}

	return &conditionMembers, nil
}

// IntersectMembersToConditionResponse returns the RPC response
type IntersectMembersToConditionResponse struct {
	ConditionSetCount int      `mapstructure:"condition_set_count"`
	Matches           []string `mapstructure:"matches"`
	MatchCount        int      `mapstructure:"match_count"`
}

// IntersectMembersToCondition returns the intersection between a set of inputs against conditions
func (c Client) IntersectMembersToCondition(members []string, condition Condition) (*IntersectMembersToConditionResponse, error) {
	var input = make(map[string]interface{})
	input["condition"] = condition
	input["members"] = members

	var intersectResponse IntersectMembersToConditionResponse

	result, err := c.CallRPC("RPCService.IntersectMembersToCondition", input)
	if err != nil {
		return &intersectResponse, err
	}

	err = mapstructure.Decode(result, &intersectResponse)
	if err != nil {
		return &intersectResponse, errors.Wrap(err, "error decoding into `IntersectMembersToConditionResponse`")
	}

	return &intersectResponse, nil
}
