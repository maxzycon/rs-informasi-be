package dto

import "gorm.io/datatypes"

type ProductRow struct {
	ID                uint                `json:"id"`
	Name              string              `json:"name"`
	CategoryProductID uint                `json:"category_product_id"`
	CategoryName      string              `json:"category_name"`
	Price             float64             `json:"price"`
	IsDiscount        bool                `json:"is_discount"`
	AmountDiscount    *float64            `json:"amount_discount"`
	StartDiscount     *datatypes.Date     `json:"start_discount"`
	EndDiscount       *datatypes.Date     `json:"end_discount"`
	Photo             *string             `json:"photo"`
	DetailProduct     []*DetailProductRow `json:"detail"`
}

type PayloadProduct struct {
	Name              string                  `json:"name"`
	CategoryProductID uint                    `json:"category_product_id"`
	Price             float64                 `json:"price"`
	IsDiscount        bool                    `json:"is_discount"`
	AmountDiscount    *float64                `json:"amount_discount"`
	StartDiscount     *string                 `json:"start_discount"`
	EndDiscount       *string                 `json:"end_discount"`
	Photo             *string                 `json:"photo"`
	DetailProduct     []*PayloadDetailProduct `json:"detail"`
}

type PayloadDetailProduct struct {
	Description string `json:"description"`
}

type DetailProductRow struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
}
