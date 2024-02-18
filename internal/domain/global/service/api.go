package service

import (
	"context"

	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"

	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
)

type GlobalService interface {
	GetAllUserPluck(ctx context.Context, user *authutil.UserClaims) (resp []*dto.UserRowPluck, err error)

	// ---- Location
	GetLocationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetLocationPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetLocationById(ctx context.Context, id int) (resp *dto.LocationRow, err error)
	CreateLocation(ctx context.Context, payload *dto.PayloadLocation) (resp *int64, err error)
	UpdateLocationById(ctx context.Context, id int, payload *dto.PayloadLocation) (resp *int64, err error)
	DeleteLocationById(ctx context.Context, id int) (resp *int64, err error)

	// ---- MerchantCategory
	GetMerchantCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetMerchantCategoryPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetMerchantCategoryById(ctx context.Context, id int) (resp *dto.MerchantCategoryRow, err error)
	CreateMerchantCategory(ctx context.Context, payload *dto.PayloadMerchantCategory) (resp *int64, err error)
	UpdateMerchantCategoryById(ctx context.Context, id int, payload *dto.PayloadMerchantCategory) (resp *int64, err error)
	DeleteMerchantCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Merchant
	GetMerchantPaginated(ctx context.Context, payload *dto.ParamsPaginationMerchant) (resp dto.MerchantWrapper, err error)
	GetMerchantPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetMerchantById(ctx context.Context, id int) (resp *dto.MerchantRow, err error)
	GetRunningTextByMerchantIdStr(ctx context.Context, id string) (resp *dto.RunningText, err error)
	CreateMerchant(ctx context.Context, payload *dto.PayloadMerchant) (resp *int64, err error)
	UpdateMerchantById(ctx context.Context, id int, payload *dto.PayloadMerchant) (resp *int64, err error)
	UpdateMerchantConfigById(ctx context.Context, id int, payload *dto.PayloadUpdateConfig) (resp *int64, err error)
	DeleteMerchantById(ctx context.Context, id int) (resp *int64, err error)

	// ---- AdvertisementCategory
	GetAdvertisementCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetAdvertisementCategoryPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetAdvertisementCategoryById(ctx context.Context, id int) (resp *dto.AdvertisementCategoryRow, err error)
	CreateAdvertisementCategory(ctx context.Context, payload *dto.PayloadAdvertisementCategory) (resp *int64, err error)
	UpdateAdvertisementCategoryById(ctx context.Context, id int, payload *dto.PayloadAdvertisementCategory) (resp *int64, err error)
	DeleteAdvertisementCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Advertisement
	GetAdvertisementPaginated(ctx context.Context, payload *dto.ParamsPaginationAdvertisement) (resp *dto.AdvertisementWrapper, err error)
	GetAdvertisementPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetAdvertisementById(ctx context.Context, id int) (resp *dto.AdvertisementDetailRow, err error)
	CreateAdvertisement(ctx context.Context, payload *dto.PayloadAdvertisement) (resp *int64, err error)
	UpdateAdvertisementById(ctx context.Context, id int, payload *dto.PayloadAdvertisement) (resp *int64, err error)
	DeleteAdvertisementById(ctx context.Context, id int) (resp *int64, err error)
	GetListContent(ctx context.Context, merchantIDstr string) (resp *dto.AdvertisementContentWrapper, err error)

	// ----- Queues
	GetQueuePaginated(ctx context.Context, payload *dto.ParamsQueueQueries) (resp dto.QueueWrapper, err error)
	GetQueueById(ctx context.Context, id int) (resp *dto.QueueRowDetail, err error)
	CreateQueue(ctx context.Context, payload *dto.PayloadQueue) (resp *int64, err error)
	UpdateStatusQueueById(ctx context.Context, id int, payload *dto.PayloadUpdateQueue) (resp *int64, err error)
	UpdateQueueById(ctx context.Context, id int, payload *dto.PayloadQueue) (resp *int64, err error)
	DeleteQueueById(ctx context.Context, id int) (resp *int64, err error)

	// ----- Analytic
	GetDashboardAnalytic(ctx context.Context) (resp *dto.SummaryDashboardWrapper, err error)

	// ----- Dashboard queue
	GetDashboardDisplay(ctx context.Context, payload *dto.ParamsQueueDisplay, merchantIdStr string) (resp *dto.QueueDataDisplayWrapper, err error)

	// ----- Get Queue detail by search
	GetQueueBySearch(ctx context.Context, merchantId string, search string) (resp *dto.QueueUserSearch, err error)

	// ----- Update queue
	UpdateFuQueueNo(ctx context.Context, id string, newPhone string) (err error)
}
