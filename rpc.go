package flip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// RPCPayload the input structure for the RPC request
type RPCPayload struct {
	ID     int           `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

// RPCResponse the response structure of the RPC interface
type RPCResponse struct {
	ID     int         `json:"id"`
	Error  string      `json:"error"`
	Result interface{} `json:"result"`
}

// CallRPC returns a response from the RPC interface
func (c Client) CallRPC(method string, param interface{}) (interface{}, error) {
	params := make([]interface{}, 0)
	params = append(params, param)

	payload := RPCPayload{
		ID:     0,
		Method: method,
		Params: params,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling json for condition %v", payload))
	}

	req, _ := http.NewRequest("POST", c.RPCURL, bytes.NewBuffer(body))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	responseBody, errBodyRead := ioutil.ReadAll(res.Body)
	if errBodyRead != nil {
		return nil, errBodyRead
	}

	var rpcResponse RPCResponse

	err = json.Unmarshal([]byte(responseBody), &rpcResponse)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("flip rpc responded with non-200 for %s, err: %s", c.RPCURL, rpcResponse.Error))
	}

	return rpcResponse.Result, nil
}
