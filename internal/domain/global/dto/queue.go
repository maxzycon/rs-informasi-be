package dto

import "time"

type QueueRow struct {
	ID            uint   `json:"id"`
	QueueNo       string `json:"queue_no"`
	MedicalRecord string `json:"medical_record"`
	Type          uint   `json:"type"`
	LocationID    uint   `json:"location_id"`
	Status        uint   `json:"status"`
	EstEnd        string `json:"est_end"`
}

type QueueRowPaginated struct {
	ID            uint       `json:"id"`
	IDStr         string     `json:"id_str"`
	QueueNo       string     `json:"queue_no"`
	MedicalRecord string     `json:"medical_record"`
	LocationID    uint       `json:"location_id"`
	LocationName  string     `json:"location_name"`
	Type          uint       `json:"type"`
	MerchantID    uint       `json:"merchant_id"`
	MerchantName  string     `json:"merchant_name"`
	CreatedByID   uint       `json:"created_by_id"`
	CreatedBy     string     `json:"created_by"`
	Status        uint       `json:"last_status"`
	IsFollowUp    bool       `json:"is_follow_up"`
	FollowUpPhone *string    `json:"follow_up_number"`
	EstEnd        *time.Time `json:"est_end"`
	CreatedAt     time.Time  `json:"created_at"`
}

type QueueWrapper struct {
	Summary *SummaryQueue     `json:"summary"`
	Items   *QueueDataWrapper `json:"items"`
}

type QueueDataWrapper struct {
	Queues    []*QueueRowPaginated    `json:"queues"`
	Paginator DefaultPaginationDtoRow `json:"paginator"`
}

type SummaryQueue struct {
	Total           uint64 `json:"total"`
	TotalDiserahkan uint64 `json:"total_diserahkan"`
	TotalRacikan    uint64 `json:"total_racikan"`
	TotalNonRacikan uint64 `json:"total_non_racikan"`
}

type PayloadQueue struct {
	QueueNo       string `json:"queue_no"`
	MedicalRecord string `json:"medical_record"`
	Type          uint   `json:"type"`
	LocationID    uint   `json:"location_id"`
	MerchantID    *uint  `json:"merchant_id"`
}

type PayloadUpdateQueue struct {
	Type      uint    `json:"type"`
	Duration  float64 `json:"duration"`
	NewStatus uint    `json:"new_status"`
}

type ParamsQueueQueries struct {
	Search *string `query:"search"`
	Order  string  `query:"order"`
	SortBy string  `query:"sort_by"`
	Limit  uint64  `query:"limit"`
	Page   uint64  `query:"page"`

	Status      []uint `query:"status"`
	Type        []uint `query:"type"`
	LocationIDS []uint `query:"locations"`
}

type ParamsQueueDisplay struct {
	Search        *string `query:"search"`
	Limit         uint64  `query:"limit"`
	Page          uint64  `query:"page"`
	MerchantIdStr string  `query:"merchant_id_str"`
}

type QueueDataDisplayWrapper struct {
	Queues    []*QueueDataDisplay     `json:"queues"`
	Paginator DefaultPaginationDtoRow `json:"paginator"`
}

type QueueDataDisplay struct {
	ID            uint   `json:"id"`
	MedicalRecord string `json:"medical_record"`
	QueueNo       string `json:"queue_no"`
	Type          uint   `json:"type"`
	Status        uint   `json:"status"`
	EstEnd        string `json:"est_end"`
}

type QueueUserSearch struct {
	ID            string  `json:"id_str"`
	MedicalRecord string  `json:"medical_record"`
	QueueNo       string  `json:"queue_no"`
	Type          uint    `json:"type"`
	Status        uint    `json:"status"`
	EstEnd        string  `json:"est_end"`
	IsFollowUp    bool    `json:"is_follow_up"`
	FollowUpPhone *string `json:"last_follow_up_number"`
}

type PayloadUpdateFollowUpQueue struct {
	NewFollowUpPhoneNo string `json:"new_follow_up_no"`
}
