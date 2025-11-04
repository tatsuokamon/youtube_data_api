package youtubedataapi

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	NoSpecify string = ""
	YoutubeTypeVideo string = "video"
	YoutubeTypePlaylist string = "playlist"
	YoutubeTypeChannel string = "channel"
)

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{
		Part: PartTypeSnippet,
		MaxResults: 50,
	}
}

type SearchRequest struct {
	Key string `json:"key"`
	Part   string `json:"part"`
	Query  string `json:"query"`
	Type    string `json:"type"`
	MaxResults int `json:"maxResults"`
}

func (r *SearchRequest) SetKey(key string) {
	r.Key = key
}

func (r *SearchRequest) ToURL() (string, error) {
	urlPrefix := "https://www.googleapis.com/youtube/v3/search"
	v := url.Values{}

	if r.Key == "" {
		return "", ErrNoKeyError
	}
	v.Add("key", r.Key)

	if r.Query != "" {
		v.Add("q", r.Query)
	}

	if r.Part != "" {
		v.Add("part", r.Part)
	}

	if r.Type != NoSpecify {
		v.Add("type", r.Type)
	}

	if r.MaxResults != 0 {
		v.Add("maxResults", strconv.Itoa(r.MaxResults))
	}

	return fmt.Sprintf("%s?%s", urlPrefix, v.Encode()), nil
}
