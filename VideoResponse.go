package youtubedataapi

type VideoResponse struct {
	*ResponseBase
	Items []VideoItem `json:"items"`
}

type VideoItem struct {
	ItemBase
	ID ID `json:"id"`
}

func (vi VideoItem) GetID() ID {
	return vi.ID
}

func (vr VideoResponse) GetItemsIter(yield func(int, Item)bool) {
	for index, item := range vr.Items {
		if !yield(index, item) {
			return
		}
	}
}
