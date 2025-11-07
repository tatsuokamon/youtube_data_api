package youtubedataapi

import "testing"

func TestClient(t *testing.T) {
	c := NewClient()

	key := getKey()
	sr := NewSearchRequest()
	sr.SetKey(key)
	sr.Query = "mokou"

	res := sr.responseType()
	err := c.Get(sr, &res)
	if err != nil {
		t.Log(err)
	}

	t.Log(res)
}
