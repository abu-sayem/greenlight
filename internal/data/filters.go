package data

import (
	"greenlight.abusayem.net/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page >= 0, "page", "page must be greater than or equal to 0")
	v.Check(f.Page <= 10_000_000, "page", "page must be less than or equal to 10,000,000")
	v.Check(f.PageSize >= 0, "page_size", "page_size must be greater than or equal to 0")
	v.Check(f.PageSize <= 100, "page_size", "page_size must be less than or equal to 100")

	v.Check(validator.In(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}
