package youtubedataapi

import "fmt"

type Request interface {
	ToURL() (string, error)
	SetKey(string)
	SetNextPageToken(string)
	EncodedQuery() (string, error)
	responseType()Response
}

type RequestBase struct {
	Key string `json:"key"`

	MaxResults    int    `json:"maxResults"`
	NextPageToken string `json:"nextPageToken"`

	Part string `json:"part"`
}

func (r *RequestBase) SetKey(key string) {
	r.Key = key
}

func (r *RequestBase) SetNextPageToken(token string) {
	r.NextPageToken = token
}

func ToURLBase(prefix string, r Request) (string, error) {
	encodedQuery, err := r.EncodedQuery()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s?%s", prefix, encodedQuery), nil
}
