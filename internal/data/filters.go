package data

import (
	"math"
	"strings"

	"greenlight.abusayem.net/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page >= 0, "page", "page must be greater than or equal to 0")
	v.Check(f.Page <= 10_000_000, "page", "page must be less than or equal to 10,000,000")
	v.Check(f.PageSize >= 0, "page_size", "page_size must be greater than or equal to 0")
	v.Check(f.PageSize <= 100, "page_size", "page_size must be less than or equal to 100")

	v.Check(validator.In(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}

// Check if the given sort value is in in SortSafelist
// and extract the column name by stripping the leading '-'
func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if safeValue == f.Sort {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("Unsafe sort parameter: " + f.Sort)
}

// return sort direction asc desc depending on prefix
func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (f Filters) limit() int {
	return f.PageSize
}

func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}
