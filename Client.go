package youtubedataapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewClient() Client {
	return Client{
		client: &http.Client{},
	}
}

type Client struct {
	client *http.Client
}

func (c *Client) Get(req Request) (Response, error) {
	u, err := req.ToURL()
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Get(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := req.responseType()
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
