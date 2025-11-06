package youtubedataapi

import "fmt"

type ID struct {
	Kind       string `json:"kind"`
	VideoID    string `json:"videoId"`
	ChannelID  string `json:"channelId"`
	PlaylistID string `json:"playlistId"`
}

func (id ID) String() string {
	result := fmt.Sprintf(`\tKind: %s\n`, id.Kind)

	for field, i := range map[string]string{
		"VideoID":    id.VideoID,
		"ChannelID":  id.ChannelID,
		"PlaylistID": id.PlaylistID,
	} {
		result += fmt.Sprintf(`\t%s: %s`, field, i)
	}

	return result
}

type Item interface {
	GetEtag() string
	GetSnippet() Snippet

	GetID() ID // 各Itemで追加
}

type Items []Item

type ItemBase struct {
	Etag    string  `json:"etag"`
	Snippet Snippet `json:"snippet"`
}

func (i ItemBase) GetEtag() string {
	return i.Etag
}

func (i ItemBase) GetSnippet() Snippet {
	return i.Snippet
}

func ItemStringBase(i Item) string {
	return fmt.Sprintf(
		`
	--------------------
%s
\tSnippet: %s
	--------------------
		 `, i.GetID().String(), i.GetSnippet().String())
}
