package youtubedataapi

type Response struct {
	Kind string `json:"kind"`
	Items         []Item `json:"items"`
	NextPageToken string `json:"nextPageToken"`
	PageInfo PageInfo `json:"pageInfo"`
}

type PageInfo struct {
	ResultsPerPage int `json:"resultsPerPage"`
	TotalResults   int `json:"toralResults"`
}
