package youtubedataapi

import (
	"net/url"
)

func NewPlaylistRequest() *PlaylistRequest {
	return &PlaylistRequest{
		RequestBase: RequestBase{Part: PartTypeSnippet},
		PlaylistID:  "",
	}
}

type PlaylistRequest struct {
	RequestBase
	PlaylistID string `json:"playlistId"`
}

func (r *PlaylistRequest) responseType() Response{
	return PlaylistResponse{}
}

func (r *PlaylistRequest) EncodedQuery() (string, error) {
	v := url.Values{}

	if r.Key == "" {
		return "", ErrNoKeyError
	}
	v.Add("key", r.Key)

	if r.PlaylistID == "" {
		return "", ErrUndefinedPlaylistID
	}
	v.Add("playlistId", r.PlaylistID)

	v.Add("part", r.Part)
	return v.Encode(), nil
}

func (r *PlaylistRequest) ToURL() (string, error) {
	return ToURLBase("https://www.googleapis.com/youtube/v3/playlistItems", r)
}
