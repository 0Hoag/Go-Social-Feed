package paginator

import "math"

const (
	defaultPage  = 1
	defaultLimit = 15
)

// PaginatorQuery is a struct that contains the page and limit of a request.
type PaginatorQuery struct {
	Page        int   `json:"page" form:"page"`
	Limit       int64 `json:"limit" form:"limit"`
	ShiftOffset int64
}

// Adjust adjusts the paginator's page and limit to the default values if they are invalid.
func (p *PaginatorQuery) Adjust() {
	if p.Page < 1 {
		p.Page = defaultPage
	}

	if p.Limit < 1 {
		p.Limit = defaultLimit
	}
}

// Offset returns the offset of the paginator.
func (p *PaginatorQuery) Offset() int64 {
	offset := int64(p.Page-1)*p.Limit - p.ShiftOffset
	if offset < 0 {
		return 0
	}
	return offset
}

type Paginator struct {
	Total       int64
	Count       int64
	PerPage     int64
	CurrentPage int
}

// TotalPages returns the total pages of the paginator.
func (p Paginator) TotalPages() int {
	if p.Total == 0 {
		return 0
	}

	return int(math.Ceil(float64(p.Total) / float64(p.PerPage)))
}

// ToResponse converts the paginator to a response.
func (p Paginator) ToResponse() PaginatorResponse {
	return PaginatorResponse{
		Total:       p.Total,
		Count:       p.Count,
		PerPage:     p.PerPage,
		CurrentPage: p.CurrentPage,
		TotalPages:  p.TotalPages(),
	}
}

// PaginatorResponse is a struct that contains the response of a paginator.
type PaginatorResponse struct {
	Total       int64 `json:"total"`
	Count       int64 `json:"count"`
	PerPage     int64 `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
}
