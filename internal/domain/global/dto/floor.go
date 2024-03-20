package dto

type FloorRow struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FloorUserRow struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	IsUsed bool   `json:"is_used"`
}

type PayloadFloor struct {
	Name string `json:"name"`
}

type WrapperUpdateFloorUser struct {
	Data []*PayloadFloorUser `json:"data"`
}

type PayloadFloorUser struct {
	ID uint `json:"id"`
}
