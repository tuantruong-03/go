package dto

type QueryResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
