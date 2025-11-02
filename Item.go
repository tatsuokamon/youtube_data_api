package youtubedataapi

type Item struct {
	Etag string `json:"eta"`
	ID struct {
		Kind string `json:"kind"`
		VideoID string `json:"videoId"`
		ChannelID string `json:"channelId"`
		PlaylistID string `json:"playlistId"`
	} `json:"id"`
	Snippet Snippet `json:"snippet"`
}
