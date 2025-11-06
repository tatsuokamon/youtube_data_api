package youtubedataapi

type Response interface {
	GetNextPageToken() (string, error)
	GetItemsIter(func(int, Item) bool) 
}

type ResponseBase struct {
	Kind         string   `json:"kind"`
	NextPageToken string   `json:"nextPageToken"`
	PageInfo     PageInfo `json:"pageInfo"`
}

type PageInfo struct {
	ResultsPerPage int `json:"resultsPerPage"`
	TotalResults   int `json:"toralResults"`
}

func (r *ResponseBase) GetNextPageToken() (string, error) {
	if r.NextPageToken == "" {
		return "", ErrNotFoundNextPageToken
	}

	return r.NextPageToken, nil
}
