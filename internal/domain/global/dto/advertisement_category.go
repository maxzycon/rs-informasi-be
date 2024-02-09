package dto

type AdvertisementCategoryRow struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type PayloadAdvertisementCategory struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
