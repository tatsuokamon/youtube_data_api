package youtubedataapi

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestVideo(t *testing.T) {
	key := getKey()
	vr := NewVideoRequest()
	vr.SetKey(key)
	vr.AddID("XlRyUkvF8RM") // マーモットの動画
	vr.AddID("Iy1MdbgZWrs") // スカタンクの動画
	vr.AddID("QI76ErPwO_I") // ゼルダの動画

	vr.DeleteID("Iy1MdbgZWrs") // IDの削除(表示されるのはまーもっととゼルダなはず)

	u, err := vr.ToURL()
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
