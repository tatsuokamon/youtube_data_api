package youtubedataapi

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

func NewVideoRequest() *VideoRequest {
	return &VideoRequest{
		RequestBase: RequestBase{Part: PartTypeSnippet},
	}
}

type VideoRequest struct {
	RequestBase
	idList []string
}

func (r *VideoRequest) ResponseType() VideoResponse{
	return VideoResponse{}
}

func (r *VideoRequest) ID() string {
	return strings.Join(r.idList, ",")
}

func (r *VideoRequest) AddID(id string) {
	r.idList = append(r.idList, id)
}

func (r *VideoRequest) DeleteID(id string) {
	if index := slices.Index(r.idList, id); index != -1 {
		r.idList = slices.Delete(r.idList, index, index+1)
	}
	fmt.Println(r.idList)
}

func (r *VideoRequest) EncodedQuery() (string, error){
	v := url.Values{}

	if r.Key == "" {
		return "", ErrNoKeyError
	}
	v.Add("key", r.Key)

	if r.ID() == "" {
		return "", ErrUndefinedVideoID
	}
	v.Add("id", r.ID())

	if r.Part == "" {
		return "", ErrUndefinedYoutubePart
	}
	v.Add("part", r.Part)

	return v.Encode(), nil
}

func (r *VideoRequest) ToURL() (string, error) {
	return ToURLBase("https://www.googleapis.com/youtube/v3/videos", r)
}
