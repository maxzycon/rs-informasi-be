package dto

type LocationRow struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type LocationUserRow struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	IsUsed bool   `json:"is_used"`
}

type PayloadLocation struct {
	Name string `json:"name"`
}

type WrapperUpdateLocationUser struct {
	Data []*PayloadLocationUser `json:"data"`
}

type PayloadLocationUser struct {
	ID uint `json:"id"`
}
