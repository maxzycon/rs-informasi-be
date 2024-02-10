package impl

import (
	"context"

	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	typequeue "github.com/maxzycon/rs-farmasi-be/pkg/constant/type_queue"
)

var bulanMap = map[string]string{
	"1":  "Januari",
	"2":  "Februari",
	"3":  "Maret",
	"4":  "April",
	"5":  "Mei",
	"6":  "Juni",
	"7":  "July",
	"8":  "Agustus",
	"9":  "September",
	"10": "Oktober",
	"11": "November",
	"12": "Desember",
}

func (s *GlobalService) GetDashboardAnalytic(ctx context.Context) (resp *dto.SummaryDashboardWrapper, err error) {
	resp = &dto.SummaryDashboardWrapper{
		QueuesChart: make([]*struct {
			Label string "json:\"label\""
			Value uint64 "json:\"value\""
		}, 0),
		QueuesLocationChart: make([]*struct {
			Label string "json:\"label\""
			Value uint64 "json:\"value\""
		}, 0),
		QueuesType: struct {
			Racikan    uint64 "json:\"racikan\""
			NonRacikan uint64 "json:\"non_racikan\""
		}{
			Racikan:    0,
			NonRacikan: 0,
		},
		QueuesAvgType: struct {
			Racikan    float64 "json:\"racikan\""
			NonRacikan float64 "json:\"non_racikan\""
		}{
			Racikan:    0,
			NonRacikan: 0,
		},
		QueuesAvgStatus: struct {
			ValidationToProcess         float64 "json:\"validation_to_process\""
			ProcessToReadyToBeSubmitted float64 "json:\"process_to_ready_to_be_submitted\""
			ReadyToBeSubmitted          float64 "json:\"ready_to_be_submitted_to_submitted\""
		}{
			ValidationToProcess:         0,
			ProcessToReadyToBeSubmitted: 0,
			ReadyToBeSubmitted:          0,
		},
	}

	user, _ := authutil.GetCredential(ctx)

	if user.MerchantID == nil {
		err = s.Summary(resp)
		if err != nil {
			return
		}
		return
	}

	err = s.SummaryByMerchantId(resp, *user.MerchantID)
	if err != nil {
		return
	}
	return
}

func (s *GlobalService) SummaryByMerchantId(dto *dto.SummaryDashboardWrapper, id uint) (err error) {
	yearlyChart, err := s.db.
		Raw(`SELECT COUNT(id) as total, MONTH(created_at) as bulan FROM queues
		WHERE YEAR(created_at) = YEAR(now())
		AND deleted_at IS NULL AND merchant_id = ?
		GROUP BY MONTH(created_at)`, id).
		Rows()

	if err != nil {
		return
	}

	for yearlyChart.Next() {
		temp := &struct {
			Label string "json:\"label\""
			Value uint64 "json:\"value\""
		}{}
		err = yearlyChart.Scan(&temp.Value, &temp.Label)
		if err != nil {
			return
		}
		temp.Label = bulanMap[temp.Label]
		dto.QueuesChart = append(dto.QueuesChart, temp)
	}
	// ----- Location
	locationChart, err := s.db.
		Raw(`SELECT COUNT(q.id) as total, l.name 
		FROM queues q
		LEFT JOIN locations l ON l.id = q.location_id
		WHERE YEAR(q.created_at) = YEAR(now()) AND MONTH(q.created_at) = MONTH(now()) AND merchant_id = ? AND q.deleted_at IS NULL 
		GROUP BY location_id`, id).
		Rows()
	if err != nil {
		return
	}
	for locationChart.Next() {
		temp := &struct {
			Label string "json:\"label\""
			Value uint64 "json:\"value\""
		}{}
		err = locationChart.Scan(&temp.Value, &temp.Label)
		if err != nil {
			return
		}
		dto.QueuesLocationChart = append(dto.QueuesLocationChart, temp)
	}
	// ----- Type
	typeChart, err := s.db.Raw(`SELECT COUNT(id) as total, type FROM queues
		WHERE YEAR(created_at) = YEAR(now()) AND MONTH(created_at) = MONTH(now()) AND merchant_id = ?
		AND deleted_at IS NULL
		GROUP BY type`, id).
		Rows()
	if err != nil {
		return
	}
	for typeChart.Next() {
		temp := &struct {
			Type  uint
			Value uint64
		}{}
		err = typeChart.Scan(&temp.Value, &temp.Type)

		if err != nil {
			return
		}

		if temp.Type == typequeue.RACIKAN {
			dto.QueuesType.Racikan = temp.Value
		}

		if temp.Type == typequeue.NON_RACIKAN {
			dto.QueuesType.NonRacikan = temp.Value
		}
	}
	// ----- avg by type (difference in minutes)
	avgType, err := s.db.Raw(`SELECT AVG(COALESCE(TIMESTAMPDIFF(MINUTE, p.created_at, d.created_at), 0)) AS avg_difference, q.type
		FROM queues q
		LEFT JOIN (SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id))  as l ON l.queue_id = q.id
		JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 2))  as p ON p.queue_id = q.id 
		JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 3))  as d ON d.queue_id = q.id 
		WHERE q.merchant_id = ? AND q.deleted_at IS NULL 
		GROUP BY q.type`, id).Rows()

	if err != nil {
		return
	}

	for avgType.Next() {
		temp := &struct {
			Type  uint
			Value float64
		}{}
		err = avgType.Scan(&temp.Value, &temp.Type)

		if err != nil {
			return
		}

		if temp.Type == typequeue.RACIKAN {
			dto.QueuesAvgType.Racikan = temp.Value
		}

		if temp.Type == typequeue.NON_RACIKAN {
			dto.QueuesAvgType.NonRacikan = temp.Value
		}
	}

	// ----- avg by status
	err = s.db.Raw(`SELECT 
	COALESCE(AVG(TIMESTAMPDIFF(MINUTE, f.created_at, t.created_at)), 0) as validation_to_process,
	COALESCE(AVG(TIMESTAMPDIFF(MINUTE, t.created_at, th.created_at)), 0) as process_to_ready_to_be_submitted,
	COALESCE(AVG(TIMESTAMPDIFF(MINUTE, th.created_at, fr.created_at)), 0) as ready_to_be_submitted_to_submitted
	FROM queues q
	LEFT JOIN (SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id))  as l ON l.queue_id = q.id
	JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 1))  as f ON f.queue_id = q.id 
	JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 2))  as t ON t.queue_id = q.id 
	JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 3))  as th ON th.queue_id = q.id 
	JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 4))  as fr ON fr.queue_id = q.id 
	WHERE q.merchant_id = ? AND q.deleted_at IS NULL`, id).
		Row().
		Scan(&dto.QueuesAvgStatus.ValidationToProcess, &dto.QueuesAvgStatus.ProcessToReadyToBeSubmitted, &dto.QueuesAvgStatus.ReadyToBeSubmitted)

	if err != nil {
		return
	}

	return
}

func (s *GlobalService) Summary(dto *dto.SummaryDashboardWrapper) (err error) {
	yearlyChart, err := s.db.
		Raw(`SELECT COUNT(id) as total, MONTH(created_at) as bulan FROM queues
		WHERE YEAR(created_at) = YEAR(now())
		AND deleted_at IS NULL
		GROUP BY MONTH(created_at)`).
		Rows()

	if err != nil {
		return
	}

	for yearlyChart.Next() {
		temp := &struct {
			Label string "json:\"label\""
			Value uint64 "json:\"value\""
		}{}
		err = yearlyChart.Scan(&temp.Value, &temp.Label)
		if err != nil {
			return
		}
		temp.Label = bulanMap[temp.Label]
		dto.QueuesChart = append(dto.QueuesChart, temp)
	}
	// ----- Location
	locationChart, err := s.db.
		Raw(`SELECT COUNT(q.id) as total, l.name 
		FROM queues q
		LEFT JOIN locations l ON l.id = q.location_id
		WHERE YEAR(q.created_at) = YEAR(now()) AND MONTH(q.created_at) = MONTH(now()) AND q.deleted_at IS NULL 
		GROUP BY location_id`).
		Rows()
	if err != nil {
		return
	}
	for locationChart.Next() {
		temp := &struct {
			Label string "json:\"label\""
			Value uint64 "json:\"value\""
		}{}
		err = locationChart.Scan(&temp.Value, &temp.Label)
		if err != nil {
			return
		}
		dto.QueuesLocationChart = append(dto.QueuesLocationChart, temp)
	}
	// ----- Type
	typeChart, err := s.db.Raw(`SELECT COUNT(id) as total, type FROM queues
		WHERE YEAR(created_at) = YEAR(now()) AND MONTH(created_at) = MONTH(now())
		AND deleted_at IS NULL
		GROUP BY type`).
		Rows()
	if err != nil {
		return
	}
	for typeChart.Next() {
		temp := &struct {
			Type  uint
			Value uint64
		}{}
		err = typeChart.Scan(&temp.Value, &temp.Type)

		if err != nil {
			return
		}

		if temp.Type == typequeue.RACIKAN {
			dto.QueuesType.Racikan = temp.Value
		}

		if temp.Type == typequeue.NON_RACIKAN {
			dto.QueuesType.NonRacikan = temp.Value
		}
	}
	// ----- avg by type (difference in minutes)
	avgType, err := s.db.Raw(`SELECT AVG(COALESCE(TIMESTAMPDIFF(MINUTE, p.created_at, d.created_at), 0)) AS avg_difference, q.type
		FROM queues q
		LEFT JOIN (SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id))  as l ON l.queue_id = q.id
		JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 2))  as p ON p.queue_id = q.id 
		JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 3))  as d ON d.queue_id = q.id 
		WHERE q.deleted_at IS NULL 
		GROUP BY q.type`).Rows()

	if err != nil {
		return
	}

	for avgType.Next() {
		temp := &struct {
			Type  uint
			Value float64
		}{}
		err = avgType.Scan(&temp.Value, &temp.Type)

		if err != nil {
			return
		}

		if temp.Type == typequeue.RACIKAN {
			dto.QueuesAvgType.Racikan = temp.Value
		}

		if temp.Type == typequeue.NON_RACIKAN {
			dto.QueuesAvgType.NonRacikan = temp.Value
		}
	}

	// ----- avg by status
	err = s.db.Raw(`SELECT 
	COALESCE(AVG(TIMESTAMPDIFF(MINUTE, f.created_at, t.created_at)), 0) as validation_to_process,
	COALESCE(AVG(TIMESTAMPDIFF(MINUTE, t.created_at, th.created_at)), 0) as process_to_ready_to_be_submitted,
	COALESCE(AVG(TIMESTAMPDIFF(MINUTE, th.created_at, fr.created_at)), 0) as ready_to_be_submitted_to_submitted
	FROM queues q
	LEFT JOIN (SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id))  as l ON l.queue_id = q.id
	JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 1))  as f ON f.queue_id = q.id 
	JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 2))  as t ON t.queue_id = q.id 
	JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 3))  as th ON th.queue_id = q.id 
	JOIN (SELECT queue_id, status, id, created_at FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id and qh.status = 4))  as fr ON fr.queue_id = q.id
	WHERE q.deleted_at IS NULL `).
		Row().
		Scan(&dto.QueuesAvgStatus.ValidationToProcess, &dto.QueuesAvgStatus.ProcessToReadyToBeSubmitted, &dto.QueuesAvgStatus.ReadyToBeSubmitted)

	if err != nil {
		return
	}

	return
}
