package youtubedataapi

import (
	"net/url"
	"strconv"
)

const (
	NoSpecify           string = ""
	YoutubeTypeVideo    string = "video"
	YoutubeTypePlaylist string = "playlist"
	YoutubeTypeChannel  string = "channel"
)

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{
		RequestBase: RequestBase{
			Part: PartTypeSnippet,
		},
		Type: NoSpecify,
	}
}

type SearchRequest struct {
	RequestBase
	Type  string `json:"type"`
	Query string `json:"query"`
}

func (r *SearchRequest)responseType() Response{
	return SearchResponse{}
}

func (r *SearchRequest) EncodedQuery() (string, error) {
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

	return v.Encode(), nil
}

func (r *SearchRequest) ToURL() (string, error) {
	return ToURLBase("https://www.googleapis.com/youtube/v3/search", r)
}
