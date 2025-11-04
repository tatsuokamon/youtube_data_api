package youtubedataapi

import (
	"fmt"
	"net/url"
)

func NewPlaylistRequest() *PlaylistRequest {
	return &PlaylistRequest{
		Part: PartTypeSnippet,
		PlaylistID: "",
	}
}

type PlaylistRequest struct {
	Key string `json:"key"`

	PlaylistID string `json:"playlistId"`
	Part       string `json:"part"`
}

func (r *PlaylistRequest) ToURL() (string, error) {
	urlPrefix := "https://www.googleapis.com/youtube/v3/playlistItems"
	v := url.Values{}

	if r.Key == "" {
		return "", ErrNoKeyError
	}
	v.Add("key", r.Key)

	if r.PlaylistID == "" {
		return "", ErrUndefinedPlaylistID
	}
	v.Add("playlistId", r.PlaylistID)

	if r.Part != PartTypeSnippet {
		v.Add("part", r.Part)
	}

	return fmt.Sprintf("%s?%s", urlPrefix, v.Encode()), nil
}
