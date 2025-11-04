package youtubedataapi

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestPlaylist(t *testing.T) {
	key := getKey()

	pr := NewPlaylistRequest()
	pr.SetKey(key)
	pr.PlaylistID = "PLloSdCaKLbA9X-RsTt3MFA4jDxyB0AGAD"
	u, err := pr.ToURL()
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
