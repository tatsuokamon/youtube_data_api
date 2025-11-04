package youtubedataapi

import "fmt"

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

func (s Snippet) ThumbnailsIter(yield func(Thumbnail) bool) {
	if s.Thumbnails.Default.URL != "" {
		if !(yield(s.Thumbnails.Default)) {
			return
		}
	}
	if s.Thumbnails.High.URL != "" {
		if !(yield(s.Thumbnails.Default)) {
			return
		}
	}
	if s.Thumbnails.Medium.URL != "" {
		if !(yield(s.Thumbnails.Default)) {
			return
		}
	}
}

type Thumbnail struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

func (s Snippet) String() string {
	result := fmt.Sprintf(`
		-- Title 		%s
		-- Description: 	%s

		-- ChanneldID:		%s
		-- ChanneldTitle:	%s
	`, s.Title, s.Description, s.ChannelID, s.ChannelTitle)
	for t := range s.ThumbnailsIter {
		result += t.String()
	}

	return result
}

func (t Thumbnail) String() string {
	return fmt.Sprintf(`
		Thumbnail:
			-- Height: %d
			-- Width:  %d
			-- URL:    %s
	`, t.Height, t.Width, t.URL)
}
