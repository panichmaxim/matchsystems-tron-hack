package models

type Cursor struct {
	Cursor  int64 `json:"cursor"`
	HasMore bool  `json:"has_more"`
}

//func NewBlogPostCursor(entries []*BlogPost, hasMore bool) *Cursor {
//	if len(entries) == 0 {
//		return &Cursor{}
//	}
//
//	return &Cursor{
//		Cursor:  entries[len(entries)-1].ID,
//		HasMore: hasMore,
//	}
//}
