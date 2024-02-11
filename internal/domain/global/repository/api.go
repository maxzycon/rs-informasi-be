package repository

import (
	"context"

	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

type GlobalRepository interface {
	// ---- users
	FindAllUser(ctx context.Context, claims *authutil.UserClaims) (resp []*model.User, err error)

	// ---- Locations
	FindLocationById(ctx context.Context, id int) (resp *model.Location, err error)
	FindAllLocation(ctx context.Context) (resp []*model.Location, err error)
	FindLocationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateLocation(ctx context.Context, entity *model.Location) (resp *int64, err error)
	UpdateLocationById(ctx context.Context, id int, entity *model.Location) (resp *int64, err error)
	DeleteLocationById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Merchant Category
	FindMerchantCategoryById(ctx context.Context, id int) (resp *model.MerchantCategory, err error)
	FindAllMerchantCategory(ctx context.Context) (resp []*model.MerchantCategory, err error)
	FindMerchantCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateMerchantCategory(ctx context.Context, entity *model.MerchantCategory) (resp *int64, err error)
	UpdateMerchantCategoryById(ctx context.Context, id int, entity *model.MerchantCategory) (resp *int64, err error)
	DeleteMerchantCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Merchants
	FindMerchantById(ctx context.Context, id int) (resp *model.Merchant, err error)
	FindAllMerchant(ctx context.Context) (resp []*model.Merchant, err error)
	FindMerchantPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateMerchant(ctx context.Context, entity *model.Merchant) (resp *int64, err error)
	UpdateMerchantById(ctx context.Context, id int, entity *model.Merchant) (resp *int64, err error)
	DeleteMerchantById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Advertisement Category
	FindAdvertisementCategoryById(ctx context.Context, id int) (resp *model.AdvertisementCategory, err error)
	FindAllAdvertisementCategory(ctx context.Context) (resp []*model.AdvertisementCategory, err error)
	FindAdvertisementCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateAdvertisementCategory(ctx context.Context, entity *model.AdvertisementCategory) (resp *int64, err error)
	UpdateAdvertisementCategoryById(ctx context.Context, id int, entity *model.AdvertisementCategory) (resp *int64, err error)
	DeleteAdvertisementCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Advertisement
	FindAdvertisementById(ctx context.Context, id int) (resp *model.Advertisement, err error)
	FindAllAdvertisement(ctx context.Context) (resp []*model.Advertisement, err error)
	FindAdvertisementPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateAdvertisement(ctx context.Context, entity *model.Advertisement) (resp *int64, err error)
	UpdateAdvertisementById(ctx context.Context, id int, entity *model.Advertisement) (resp *int64, err error)
	DeleteAdvertisementById(ctx context.Context, id int) (resp *int64, err error)

	// ----- Queue
	DeleteQueueById(ctx context.Context, id int) (resp *int64, err error)
	UpdateQueueById(ctx context.Context, id int, entity *model.Queue) (resp *int64, err error)
}
