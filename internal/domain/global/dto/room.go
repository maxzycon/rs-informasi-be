package dto

type RoomRow struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	FloorId   uint   `json:"floor_id"`
	FloorName string `json:"floor_name"`
}

type PayloadRoom struct {
	Name    string `json:"name"`
	FloorId uint   `json:"floor_id"`
}
