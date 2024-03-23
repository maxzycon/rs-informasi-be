package dto

type FacilityRow struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Photo       *string `json:"photo"`
	Description *string `json:"description"`
}

type PayloadFacility struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Photo       *string `json:"photo"`
}
