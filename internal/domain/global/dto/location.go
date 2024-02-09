package dto

type LocationRow struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PayloadLocation struct {
	Name string `json:"name"`
}
