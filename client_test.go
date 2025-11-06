package youtubedataapi

import "testing"

func TestClient(t *testing.T) {
	c := NewClient()
	for r := range readyRequest {
		res, err := c.Get(r)
		if err != nil {
			t.Log(err)
		}

		t.Log(res)
	}
}
