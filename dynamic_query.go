package flip

import (
	"github.com/FlipsideCrypto/flip-rpc-client-go/dynamicquery"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// ExecuteDynamicQueryResponse returns the RPC response
type ExecuteDynamicQueryResponse struct {
	Results     []interface{} `mapstructure:"results" json:"results"`
	ResultCount int           `mapstructure:"result_count" json:"result_count"`
	CompiledSQL string        `mapstructure:"compiled_sql" json:"compiled_sql,omitempty"`
	TTLHash     string        `mapstructure:"ttl_hash" json:"ttl_hash,omitempty"`
	TTLStart    string        `mapstructure:"ttl_start" json:"ttl_start,omitempty"`
	TTLExpiry   string        `mapstructure:"ttl_expiry" json:"ttl_expiry,omitempty"`
	Error       string        `mapstructure:"error,omitempty" json:"error,omitempty"`
}

// ExecuteDynamicQuery returns the query results
func (c Client) ExecuteDynamicQuery(query dynamicquery.Query, debug bool, ttlSeconds int) (*ExecuteDynamicQueryResponse, error) {
	var input = make(map[string]interface{})
	input["query"] = query
	input["debug"] = debug
	input["ttl_seconds"] = ttlSeconds

	var response ExecuteDynamicQueryResponse

	rpc, err := c.CallRPC("RPCService.ExecuteDynamicQuery", input)

	if err != nil {
		return &response, err
	}

	err = mapstructure.Decode(rpc.Result, &response)
	if err != nil {
		return &response, errors.Wrap(err, "error decoding into `ExecuteDynamicQueryResponse`")
	}
	response.Error = rpc.Error

	return &response, nil
}
