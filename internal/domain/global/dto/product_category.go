package dto

type ProductCategoryRow struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PayloadProductCategory struct {
	Name string `json:"name"`
}
