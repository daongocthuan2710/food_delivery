package common

type Paging struct {
	Page  int   `json:"page" format:"page"`
	Limit int   `json:"limit" format:"limit"`
	Total int64 `json:"total"`
}

func (p *Paging) Process() error {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}

	return nil
}
