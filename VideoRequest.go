package youtubedataapi

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

func NewVideoRequest() *VideoRequest {
	return &VideoRequest{
		Part: PartTypeSnippet,
	}
}

type VideoRequest struct {
	Key    string `json:"key"`
	ID     string `json:"id"`
	Part   string `json:"part"`
	idList []string
}

func (r *VideoRequest) SetKey(key string) {
	r.Key = key
}

func (r *VideoRequest) AddID(id string) {
	r.idList = append(r.idList, id)
	r.ID = strings.Join(r.idList, ",")
}

func (r *VideoRequest) DeleteID(id string) {
	if index := slices.Index(r.idList, id); index != -1 {
		r.idList = slices.Delete(r.idList, index, index + 1)
	}
	fmt.Println(r.idList)
	r.ID = strings.Join(r.idList, ",")
	fmt.Println(r.ID)
}

func (r *VideoRequest) ToURL() (string, error) {
	urlPrefix := "https://www.googleapis.com/youtube/v3/videos"
	v := url.Values{}

	if r.Key == "" {
		return "", ErrNoKeyError
	}
	v.Add("key", r.Key)

	if r.ID == "" {
		return "", ErrUndefinedVideoID
	}
	v.Add("id", r.ID)

	if r.Part == "" {
		return "", ErrUndefinedYoutubePart
	}
	v.Add("part", r.Part)

	return fmt.Sprintf("%s?%s", urlPrefix, v.Encode()), nil
}
