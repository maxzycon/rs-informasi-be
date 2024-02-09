package dto

type DefaultPluck struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DefaultPaginationDtoRow struct {
	CurrentPage   uint64 `json:"current_page"`
	RecordPerPage uint64 `json:"per_page"`
	LastPage      uint64 `json:"last_page"`
	TotalItem     uint64 `json:"total_item"`
}
