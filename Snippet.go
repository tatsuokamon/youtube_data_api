package youtubedataapi

import "fmt"

type Snippet struct {
	ChannelID            string     `json:"channelId"`
	ChannelTitle         string     `json:"channelTitle"`

	Description          string     `json:"description"`
	LiveBroadcastContent string     `json:"liveBroadcastContent"`
	PublishTime          string     `json:"publishTime"`
	PublishedAt          string     `json:"publishedAt"`

	Thumbnails           Thumbnails `json:"thumbnails"`
	ResourceID ID `json:"resourceId"`
	Title string `json:"title"`
}

type Thumbnails struct {
	Default  Thumbnail `json:"default"`
	High     Thumbnail `json:"high"`
	Medium   Thumbnail `json:"medium"`
	Standard Thumbnail `json:"standard"`
	Maxres   Thumbnail `json:"maxres"`
}

func (ts Thumbnails) Iter(yield func(Thumbnail) bool) {
	for _, t := range []Thumbnail{
		ts.Default,
		ts.High,
		ts.Medium,
		ts.Standard,
		ts.Maxres,
	} {
		if t.URL != "" {
			if !yield(t) {
				return
			}
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
	for t := range s.Thumbnails.Iter {
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
