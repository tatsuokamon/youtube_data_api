package youtubedataapi

import (
	"encoding/json"
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

func (c *Client) Get(req Request, st any) (error) {
	u, err := req.ToURL()
	if err != nil {
		return err
	}

	resp, err := c.client.Get(u)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, st)
	if err != nil {
		return err
	}
	return nil

}
