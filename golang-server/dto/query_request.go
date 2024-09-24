package dto

import (
	"fmt"
	"strings"
)

type QueryRequest struct {
	Keyword  string            `json:"keyword"`
	PageSize int               `json:"pageSize"`
	Page     int               `json:"page"`
	Sort     []SortCriteria    `json:"sort"`   // Multiple sorting criteria
	Filter   map[string]string `json:"filter"` // Key-value pairs for filtering
}

type SortCriteria struct {
	Field string `json:"field"`
	Order string `json:"order"` // "asc" or "desc"
}

func (qr QueryRequest) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Keyword: %d\n", qr.Keyword))
	sb.WriteString(fmt.Sprintf("PageSize: %d\n", qr.PageSize))
	sb.WriteString(fmt.Sprintf("Page: %d\n", qr.Page))

	if len(qr.Sort) > 0 {
		sb.WriteString("Sort:\n")
		for _, sort := range qr.Sort {
			sb.WriteString(fmt.Sprintf("  - Field: %s, Order: %s\n", sort.Field, sort.Order))
		}
	} else {
		sb.WriteString("Sort: None\n")
	}

	if len(qr.Filter) > 0 {
		sb.WriteString("Filter:\n")
		for key, value := range qr.Filter {
			sb.WriteString(fmt.Sprintf("  - %s: %s\n", key, value))
		}
	} else {
		sb.WriteString("Filter: None\n")
	}

	return sb.String()
}
