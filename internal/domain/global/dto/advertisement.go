package dto

import uuid "github.com/satori/go.uuid"

type AdvertisementRow struct {
	ID                        uint    `json:"id"`
	Name                      string  `json:"name"`
	Company                   string  `json:"company"`
	DateStart                 string  `json:"date_start"`
	DateEnd                   string  `json:"date_end"`
	MerchantID                *uint   `json:"merchant_id"`
	MerchantName              *string `json:"merchant_name"`
	CategoryAdvertisementID   *uint   `json:"category_advertisement_id"`
	CategoryAdvertisementName *string `json:"category_advertisement_name"`
	Status                    string  `json:"status"`
}

type AdvertisementDetailRow struct {
	ID                        uint    `json:"id"`
	Name                      string  `json:"name"`
	Company                   string  `json:"company"`
	DateStart                 string  `json:"date_start"`
	DateEnd                   string  `json:"date_end"`
	MerchantID                *uint   `json:"merchant_id"`
	MerchantName              *string `json:"merchant_name"`
	CategoryAdvertisementID   *uint   `json:"category_advertisement_id"`
	CategoryAdvertisementName *string `json:"category_advertisement_name"`
	Path                      string  `json:"document_path"`
	Description               *string `json:"description"`
}

type AdvertisementData struct {
	Advertisements []*AdvertisementRow     `json:"advertisement"`
	Paginator      DefaultPaginationDtoRow `json:"paginator"`
}

type AdvertisementWrapper struct {
	Summary           *SummaryAdvertisement `json:"summary"`
	AdvertisementData AdvertisementData     `json:"items"`
}

type SummaryAdvertisement struct {
	OnGoing    uint64 `json:"ongoing"`
	Finished   uint64 `json:"finished"`
	OnSchedule uint64 `json:"on_schedule"`
}

type PayloadAdvertisement struct {
	Name                    string  `json:"name"`
	Company                 string  `json:"company"`
	DateStart               string  `json:"date_start"`
	DateEnd                 string  `json:"date_end"`
	MerchantID              *uint   `json:"merchant_id"`
	CategoryAdvertisementID *uint   `json:"category_advertisement_id"`
	DocumentPath            string  `json:"document_path"`
	Description             *string `json:"description"`
}

type ParamsPaginationAdvertisement struct {
	Search   *string `query:"search"`
	Order    string  `query:"order"`
	SortBy   string  `query:"sort_by"`
	Limit    uint64  `query:"limit"`
	Page     uint64  `query:"page"`
	Category uint64  `query:"category"`
}

type AdvertisementContentWrapper struct {
	Daily    []*AdvertisementListData `json:"daily_content"`
	Contents []*AdvertisementListData `json:"content"`
}

type AdvertisementListData struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type AdvertisementMerchant struct {
	IDStr uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Photo *string   `json:"photo"`
}
