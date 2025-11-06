package youtubedataapi

type SearchResponse struct {
	*ResponseBase
	Items []SearchItem `json:"items"`
}

type SearchItem struct {
	ItemBase
	ID ID `json:"id"`
}

func (si SearchItem) GetID() ID {
	return  si.ID
}

func (sr SearchResponse)GetItemsIter(yield func(int, Item)bool) {
	for index, item := range sr.Items {
		if !yield(index, item) {
			return
		}
	}
}
