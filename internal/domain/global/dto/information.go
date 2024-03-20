package dto

type InformationRow struct {
	ID                    uint    `json:"id"`
	Name                  string  `json:"name"`
	InformationCategoryID uint    `json:"information_category_id"`
	Photo                 *string `json:"photo"`
	Desc                  *string `json:"description"`
}

type PayloadInformation struct {
	Name                  string  `json:"name"`
	InformationCategoryID uint    `json:"information_category_id"`
	Photo                 *string `json:"photo"`
	Desc                  *string `json:"description"`
}
