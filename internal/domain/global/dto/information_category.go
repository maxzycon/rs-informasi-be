package dto

type InformationCategoryRow struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PayloadInformationCategory struct {
	Name string `json:"name"`
}
