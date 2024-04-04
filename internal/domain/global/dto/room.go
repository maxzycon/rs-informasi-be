package dto

type RoomRow struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Photo       *string `json:"photo"`
	Description string  `json:"description"`
	FloorId     uint    `json:"floor_id"`
	FloorName   string  `json:"floor_name"`
}

type PayloadRoom struct {
	Name        string  `json:"name"`
	FloorId     uint    `json:"floor_id"`
	Photo       *string `json:"photo"`
	Description string  `json:"description"`
}
