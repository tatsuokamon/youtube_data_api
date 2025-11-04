package youtubedataapi

import "fmt"

type Item struct {
	Etag string `json:"etag"`
	ID struct {
		Kind string `json:"kind"`
		VideoID string `json:"videoId"`
		ChannelID string `json:"channelId"`
		PlaylistID string `json:"playlistId"`
	} `json:"id"`
	Snippet Snippet `json:"snippet"`
}

func (i Item) String() string{
	return fmt.Sprintf(
	`
	--------------------
	Kind: %s
	VideoID: %s
	ChannelID: %s
	PlaylistID: %s
	Snippet: %s
	--------------------
		 `,i.ID.Kind, i.ID.VideoID, i.ID.ChannelID, i.ID.PlaylistID, i.Snippet.String())
}
