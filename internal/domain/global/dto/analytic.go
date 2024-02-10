package dto

type SummaryDashboardWrapper struct {
	QueuesChart []*struct {
		Label string `json:"label"`
		Value uint64 `json:"value"`
	} `json:"queues_chart"`

	QueuesLocationChart []*struct {
		Label string `json:"label"`
		Value uint64 `json:"value"`
	} `json:"queues_location_chart"`

	QueuesType struct {
		Racikan    uint64 `json:"racikan"`
		NonRacikan uint64 `json:"non_racikan"`
	} `json:"queues_type"`

	QueuesAvgType struct {
		Racikan    float64 `json:"racikan"`
		NonRacikan float64 `json:"non_racikan"`
	} `json:"queues_avg_type"`

	QueuesAvgStatus struct {
		ValidationToProcess         float64 `json:"validation_to_process"`
		ProcessToReadyToBeSubmitted float64 `json:"process_to_ready_to_be_submitted"`
		ReadyToBeSubmitted          float64 `json:"ready_to_be_submitted_to_submitted"`
	} `json:"queues_avg_status"`
}
