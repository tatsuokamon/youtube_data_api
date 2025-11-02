package youtubedataapi

type Snippet struct {
	ChannelID            string `json:"channelId"`
	ChannelTitle         string `json:"channelTitle"`
	Description          string `json:"description"`
	LiveBroadcastContent string `json:"liveBroadcastContent"`
	PublishTime          string `json:"publishTime"`
	PublishedAt          string `json:"publishedAt"`
	Thumbnails           struct {
		Default Thumbnail `json:"default"`
		High    Thumbnail `json:"high"`
		Medium  Thumbnail `json:"medium"`
	} `json:"thumbnails"`

	Title string `json:"title"`
}

type Thumbnail struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}
