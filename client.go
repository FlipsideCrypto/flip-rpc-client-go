package flip

import "fmt"

// Config allows a consuming app to set up API Key
type Config struct {
	APIKey  string
	BaseURL string
}

// Client allows access to the Flip RPC Interface
type Client struct {
	APIKey  string
	BaseURL string
	RPCURL  string
}

// NewClient returns a new Databridge Client
func NewClient(config Config) (Client, error) {
	c := Client{}
	c.APIKey = config.APIKey
	c.BaseURL = config.BaseURL
	c.RPCURL = fmt.Sprintf("%s/rpc?api_key=%s", c.BaseURL, c.APIKey)

	return c, nil
}
