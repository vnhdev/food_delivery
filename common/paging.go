package common

type Paging struct {
	Page       int    `json:"page" form:"page"`
	Limit      int    `json:"Limit" form:"limit"`
	Total      int64  `json:"total" form:"total"`
	FakeCursor string `json:"cursor" form:"cursor"`
	NExtCursor string `json:"next_cursor"`
}

func (p *Paging) Fullfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}
}
