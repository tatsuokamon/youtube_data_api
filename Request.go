package youtubedataapi

import (
	"fmt"
	"net/url"
)

const (
	NoSpecify = iota
	YoutubeTypeVideo
	YoutubeTypePlaylist
	YoutubeTypeChannel
)

type Request struct {
	Query string `json:"query"`
	Key   string `json:"key"`
	Type  int    `json:"type"`
}

func (r *Request) SetKey(key string) {
	r.Key = key
}

func (r *Request) ToURL() (string, error) {
	urlPrefix := "https://www.googleapis.com/youtube/v3/search"
	v := url.Values{}

	if r.Key == "" {
		return "", ErrNoKeyError
	}

	v.Add("key", r.Key)
	if r.Query != "" {
		v.Add("query", r.Query)
	}
	if r.Type != NoSpecify {
		value, err := fromIntToTypeKeyword(r.Type)
		if err != nil {
			return "", err
		}
		v.Add("type", value)
	}

	return fmt.Sprintf("%s?%s", urlPrefix, v.Encode()), nil
}

func fromIntToTypeKeyword(i int) (string, error) {
	switch i {
	case YoutubeTypeVideo:
		{
			return "video", nil
		}
	case YoutubeTypePlaylist:
		{
			return "playlist", nil
		}
	case YoutubeTypeChannel:
		{
			return "channel", nil
		}
	default:
		{
			return "", ErrUndefinedYoutubeType
		}
	}
}
