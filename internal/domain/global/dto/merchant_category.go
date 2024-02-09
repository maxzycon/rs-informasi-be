package dto

type MerchantCategoryRow struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PayloadMerchantCategory struct {
	Name string `json:"name"`
}
