package pagination

import (
	"fmt"

	"gorm.io/gorm"
)

type DefaultPaginationPayload struct {
	Search *string `query:"search"`
	Order  string  `query:"order"`
	SortBy string  `query:"sort_by"`
	Limit  int     `query:"limit"`
	Page   int     `query:"page"`
}

type DefaultPaginationRow struct {
	CurrentPage   int `json:"current_page"`
	RecordPerPage int `json:"per_page"`
	LastPage      int `json:"last_page"`
	TotalItem     int `json:"total_item"`
}

type DefaultPagination struct {
	Items     interface{}          `json:"items"`
	Paginator DefaultPaginationRow `json:"paginator"`
}

func (pgs *DefaultPaginationPayload) ToPaginationPayloadManual() *DefaultPaginationPayload {
	if pgs.Limit < 1 {
		pgs.Limit = 1
	}

	pgs.Page = pgs.Page - 1

	return pgs
}

func (pgs *DefaultPaginationPayload) ToPaginationManual(pagination *DefaultPaginationRow, totalRows int64) *DefaultPaginationRow {
	total := totalRows / int64(pgs.Limit)
	remainder := totalRows % int64(pgs.Limit)

	resp := pagination

	if remainder == 0 {
		resp.LastPage = int(total)
	} else {
		resp.LastPage = int(total + 1)
	}
	// Set current/record per page meta data
	resp.TotalItem = int(totalRows)
	resp.RecordPerPage = pgs.Limit
	resp.CurrentPage = pgs.Page + 1
	return resp
}

func (pgs *DefaultPaginationPayload) Pagination(value interface{}, pagination *DefaultPaginationRow, db *gorm.DB) func(*gorm.DB) *gorm.DB {
	if pgs.Limit < 1 {
		pgs.Limit = 1
	}

	pgs.Page = pgs.Page - 1

	var totalRows int64
	//sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	//	return tx.Model(value).Count(&totalRows)
	//})
	//fmt.Println(sql)
	//fmt.Println("=====")
	db.Model(value).Session(&gorm.Session{}).Count(&totalRows)

	total := totalRows / int64(pgs.Limit)
	remainder := totalRows % int64(pgs.Limit)

	resp := pagination

	if remainder == 0 {
		resp.LastPage = int(total)
	} else {
		resp.LastPage = int(total + 1)
	}

	// Set current/record per page meta data
	resp.TotalItem = int(totalRows)
	resp.RecordPerPage = pgs.Limit
	resp.CurrentPage = pgs.Page + 1

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pgs.Page * pgs.Limit).Limit(pgs.Limit).Order(fmt.Sprintf("%s %s", pgs.SortBy, pgs.Order))
	}
}

func (pgs *DefaultPaginationPayload) PaginationV3(pagination *DefaultPaginationRow, sql *gorm.DB) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pgs.Limit < 1 {
			pgs.Limit = 1
		}

		pgs.Page = pgs.Page - 1
		var totalRows int64

		sql.Session(&gorm.Session{}).Count(&totalRows)

		total := totalRows / int64(pgs.Limit)
		remainder := totalRows % int64(pgs.Limit)

		resp := pagination

		if remainder == 0 {
			resp.LastPage = int(total)
		} else {
			resp.LastPage = int(total + 1)
		}

		// Set current/record per page meta data
		resp.TotalItem = int(totalRows)
		resp.RecordPerPage = pgs.Limit
		resp.CurrentPage = pgs.Page + 1

		return db.Offset(pgs.Page * pgs.Limit).Limit(pgs.Limit).Order(fmt.Sprintf("%s %s", pgs.SortBy, pgs.Order))
	}
}

func (pgs *DefaultPaginationPayload) PaginationV2(pagination *DefaultPaginationRow) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pgs.Limit < 1 {
			pgs.Limit = 1
		}

		pgs.Page = pgs.Page - 1
		var totalRows int64
		db.Session(&gorm.Session{}).Count(&totalRows)

		total := totalRows / int64(pgs.Limit)
		remainder := totalRows % int64(pgs.Limit)

		resp := pagination

		if remainder == 0 {
			resp.LastPage = int(total)
		} else {
			resp.LastPage = int(total + 1)
		}

		// Set current/record per page meta data
		resp.TotalItem = int(totalRows)
		resp.RecordPerPage = pgs.Limit
		resp.CurrentPage = pgs.Page + 1

		return db.Offset(pgs.Page * pgs.Limit).Limit(pgs.Limit).Order(fmt.Sprintf("%s %s", pgs.SortBy, pgs.Order))
	}
}
