package dto

type MerchantRow struct {
	ID                 uint    `json:"id"`
	Name               string  `json:"name"`
	Email              string  `json:"email"`
	PICName            string  `json:"pic_name"`
	Phone              string  `json:"phone"`
	Photo              *string `json:"photo"`
	Address            string  `json:"address"`
	MerchantCategoryID uint    `json:"merchant_category_id"`
}

type MerchantRowPaginated struct {
	ID           uint    `json:"id"`
	IDStr        string  `json:"id_str"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	PICName      string  `json:"pic_name"`
	Phone        string  `json:"phone"`
	Photo        *string `json:"photo"`
	Address      string  `json:"address"`
	CategoryID   uint    `json:"category_merchant_id"`
	CategoryName string  `json:"category_merchant_name"`
}

type MerchantWrapper struct {
	SummaryMerchant *SummaryMerchant     `json:"summary"`
	Items           *MerchantDataWrapper `json:"items"`
}

type MerchantDataWrapper struct {
	Merchants []*MerchantRowPaginated `json:"merchants"`
	Paginator DefaultPaginationDtoRow `json:"paginator"`
}

type SummaryMerchant struct {
	TotalMerchant uint64 `json:"total"`
}

type ParamsPaginationMerchant struct {
	Search   *string `query:"search"`
	Order    string  `query:"order"`
	SortBy   string  `query:"sort_by"`
	Limit    uint64  `query:"limit"`
	Page     uint64  `query:"page"`
	Category uint64  `query:"category"`
}

type PayloadMerchant struct {
	Name               string  `json:"name"`
	Email              string  `json:"email"`
	PICName            string  `json:"pic_name"`
	Phone              string  `json:"phone"`
	Photo              *string `json:"photo"`
	Address            string  `json:"address"`
	MerchantCategoryID uint    `json:"merchant_category_id"`
}

type PayloadUpdateConfig struct {
	RunningText *string `json:"running_text"`
}
