package flip

import (
	"github.com/FlipsideCrypto/flip-rpc-client-go/segment"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// GetSegmentMembersResponse returns the RPC response
type GetSegmentMembersResponse struct {
	Members     []string `mapstructure:"members"`
	MemberCount int      `mapstructure:"member_count"`
}

// GetSegmentMembers returns the members belonging to the result set of a condition.
func (c Client) GetSegmentMembers(condition segment.Condition) (*GetSegmentMembersResponse, error) {
	var input = make(map[string]segment.Condition)
	input["segment"] = condition

	var segmentMembers GetSegmentMembersResponse

	result, err := c.CallRPC("RPCService.GetSegmentMembers", input)
	if err != nil {
		return &segmentMembers, err
	}

	err = mapstructure.Decode(result, &segmentMembers)
	if err != nil {
		return &segmentMembers, errors.Wrap(err, "error decoding into `GetSegmentMembersResponse`")
	}

	return &segmentMembers, nil
}

// IntersectMembersToSegmentResponse returns the RPC response
type IntersectMembersToSegmentResponse struct {
	SegmentMemberCount int      `mapstructure:"segment_member_count"`
	Matches            []string `mapstructure:"matches"`
	MatchCount         int      `mapstructure:"match_count"`
}

// IntersectMembersToSegment returns the intersection between a set of inputs against a segment formed by conditions
func (c Client) IntersectMembersToSegment(members []string, condition segment.Condition) (*IntersectMembersToSegmentResponse, error) {
	var input = make(map[string]interface{})
	input["segment"] = condition
	input["members"] = members

	var intersectResponse IntersectMembersToSegmentResponse

	result, err := c.CallRPC("RPCService.IntersectMembersToSegment", input)
	if err != nil {
		return &intersectResponse, err
	}

	err = mapstructure.Decode(result, &intersectResponse)
	if err != nil {
		return &intersectResponse, errors.Wrap(err, "error decoding into `IntersectMembersToSegmentResponse`")
	}

	return &intersectResponse, nil
}
