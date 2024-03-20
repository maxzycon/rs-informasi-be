package dto

type ServiceRow struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Photo *string `json:"photo"`
	Desc  *string `json:"description"`
}

type PayloadService struct {
	Name  string  `json:"name"`
	Photo *string `json:"photo"`
	Desc  *string `json:"description"`
}
