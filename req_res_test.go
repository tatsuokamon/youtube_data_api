package youtubedataapi

import (
	"io"
	"net/http"
	"testing"
)

func readyRequest(yield func(Request) bool) {
	key := getKey()

	sr := NewSearchRequest()
	sr.Query = "もこう"
	sr.SetKey(key)
	if !yield(sr) {
		return
	}

	vr := NewVideoRequest()
	vr.SetKey(key)
	vr.AddID("XlRyUkvF8RM") // マーモットの動画
	vr.AddID("Iy1MdbgZWrs") // スカタンクの動画
	vr.AddID("QI76ErPwO_I") // ゼルダの動画

	vr.DeleteID("Iy1MdbgZWrs") // IDの削除(表示されるのはまーもっととゼルダなはず)
	if !yield(vr) {
		return
	}

	pr := NewPlaylistRequest()
	pr.SetKey(key)
	pr.PlaylistID = "PLloSdCaKLbA9X-RsTt3MFA4jDxyB0AGAD"
	if !yield(pr) {
		return
	}
}

func TestReq(t *testing.T) {
	client := &http.Client{}

	for r := range readyRequest {
		u, err := r.ToURL()
		if err != nil {
			t.Log(err)
		}
		resp, err := client.Get(u)
		if err != nil {
			t.Log(err)
		}
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Log(err)
		}

		t.Log(string(content))
	}
}
