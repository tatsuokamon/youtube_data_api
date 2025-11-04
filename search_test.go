package youtubedataapi

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
)

func getKey() string {
	return os.Getenv("YUMKEY")
}

func TestSearch(t *testing.T) {
	key := getKey()

	// SearchRequest
	sr := NewSearchRequest()
	sr.SetKey(key)
	sr.Query = "もこう"
	sr.Part = PartTypeSnippet

	u, err := sr.ToURL()
	if err != nil {
		t.Log(err)
	}
	t.Log(u)

	resp, err := http.Get(u)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(content))

	var res Response
	err = json.Unmarshal(content, &res)
	if err != nil {
		t.Log(err)
	}

	t.Log(res)
}
