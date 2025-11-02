package youtubedataapi

type Response struct {
	Items         []Item `json:"items"`
	NextPageToken string `json:"nextPageToken"`
}

type PageInfo struct {
	ResultsPerPage int `json:"resultsPerPage"`
	TotalResults   int `json:"toralResults"`
}
