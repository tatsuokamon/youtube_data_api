package youtubedataapi

type PlaylistResponse struct {
	*ResponseBase
	Items []PlaylistItem `json:"items"`
}

type PlaylistItem struct {
	ItemBase
}

func (pi PlaylistItem) GetID() ID {
	return pi.Snippet.ResourceID
}

func (pr PlaylistResponse) GetItemsIter(yield func(int, Item) bool) {
	for index, item := range pr.Items {
		if !yield(index, item) {
			return
		}
	}
}
