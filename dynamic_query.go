package flip

import (
	"github.com/FlipsideCrypto/flip-rpc-client-go/dynamicquery"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// ExecuteDynamicQueryResponse returns the RPC response
type ExecuteDynamicQueryResponse struct {
	Results     []interface{} `mapstructure:"results"`
	ResultCount int           `mapstructure:"result_count"`
	CompiledSQL string        `mapstructure:"compiled_sql"`
}

// ExecuteDynamicQuery returns the query results
func (c Client) ExecuteDynamicQuery(query dynamicquery.Query, debug bool) (*ExecuteDynamicQueryResponse, error) {
	var input = make(map[string]interface{})
	input["query"] = query
	input["debug"] = debug

	var response ExecuteDynamicQueryResponse

	result, err := c.CallRPC("RPCService.ExecuteDynamicQuery", input)
	if err != nil {
		return &response, err
	}

	err = mapstructure.Decode(result, &response)
	if err != nil {
		return &response, errors.Wrap(err, "error decoding into `ExecuteDynamicQueryResponse`")
	}

	return &response, nil
}
