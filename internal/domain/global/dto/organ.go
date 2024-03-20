package dto

type OrganRow struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PayloadOrgan struct {
	Name string `json:"name"`
}
