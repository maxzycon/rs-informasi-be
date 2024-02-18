package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-farmasi-be/internal/config"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/service"
	"github.com/maxzycon/rs-farmasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-farmasi-be/pkg/middleware"
)

const (
	GetUserPluck = "/users_pluck"

	GetLocationPluck     = "locations/list"
	GetLocationPaginated = "locations/paginated"
	Location             = "locations"
	LocationById         = "locations/:id"

	GetMerchantCategoryPluck     = "merchant_categories/list"
	GetMerchantCategoryPaginated = "merchant_categories/paginated"
	MerchantCategory             = "merchant_categories"
	MerchantCategoryById         = "merchant_categories/:id"

	GetMerchantPluck     = "merchants/list"
	GetMerchantPaginated = "merchants/paginated"
	Merchant             = "merchants"
	MerchantById         = "merchants/:id"
	MerchantConfigById   = Merchant + "/detail/config"

	GetAdvertisementCategoryPluck     = "advertisement_categories/list"
	GetAdvertisementCategoryPaginated = "advertisement_categories/paginated"
	AdvertisementCategory             = "advertisement_categories"
	AdvertisementCategoryById         = "advertisement_categories/:id"

	GetAdvertisementPluck     = "advertisements/list"
	GetAdvertisementPaginated = "advertisements/paginated"
	Advertisement             = "advertisements"
	AdvertisementContent      = "advertisements/content/:id"
	AdvertisementById         = "advertisements/:id"

	GetQueuePaginated = "queues/paginated"
	Queue             = "queues"
	QueueById         = "queues/:id"
	QueueBySearch     = "queues/detail"
	UpdateQueueFu     = "queues/fu/:id"
	QueueStatusById   = "queues/status/:id"

	AnalyticDashboard = "analytic/dashboard"
	DisplayDashboard  = "dashboard"
)

type GlobalControllerParams struct {
	V1            fiber.Router
	Conf          *config.Config
	GlobalService service.GlobalService
	Middleware    middleware.GlobalMiddleware
}
type GlobalController struct {
	v1            fiber.Router
	conf          *config.Config
	globalService service.GlobalService
	middleware    middleware.GlobalMiddleware
}

func New(params *GlobalControllerParams) *GlobalController {
	return &GlobalController{
		v1:            params.V1,
		conf:          params.Conf,
		globalService: params.GlobalService,
		middleware:    params.Middleware,
	}
}

func (pc *GlobalController) Init() {
	// ---- User
	pc.v1.Get(GetUserPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllUserPluck)

	// ---- Location
	pc.v1.Get(GetLocationPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllLocationPluck)
	pc.v1.Get(GetLocationPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerGetLocationPaginated)
	pc.v1.Get(LocationById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerGetLocationById)
	pc.v1.Post(Location, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerCreateLocation)
	pc.v1.Put(LocationById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerUpdateLocation)
	pc.v1.Delete(LocationById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerDeleteLocation)

	// ---- MerchantCategory
	pc.v1.Get(GetMerchantCategoryPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllMerchantCategoryPluck)
	pc.v1.Get(GetMerchantCategoryPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING}), pc.handlerGetMerchantCategoryPaginated)
	pc.v1.Get(MerchantCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerGetMerchantCategoryById)
	pc.v1.Post(MerchantCategory, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerCreateMerchantCategory)
	pc.v1.Put(MerchantCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerUpdateMerchantCategory)
	pc.v1.Delete(MerchantCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerDeleteMerchantCategory)

	// ---- Merchant
	pc.v1.Get(GetMerchantPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllMerchantPluck)
	pc.v1.Get(GetMerchantPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING}), pc.handlerGetMerchantPaginated)
	pc.v1.Get(MerchantById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerGetMerchantById)
	pc.v1.Post(Merchant, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerCreateMerchant)
	pc.v1.Put(MerchantById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerUpdateMerchant)
	pc.v1.Put(MerchantConfigById, pc.middleware.Protected([]uint{role.ROLE_SUPER_ADMIN}), pc.handlerUpdateMerchantConfig)
	pc.v1.Delete(MerchantById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerDeleteMerchant)

	// ---- AdvertisementCategory
	pc.v1.Get(GetAdvertisementCategoryPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllAdvertisementCategoryPluck)
	pc.v1.Get(GetAdvertisementCategoryPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING}), pc.handlerGetAdvertisementCategoryPaginated)
	pc.v1.Get(AdvertisementCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerGetAdvertisementCategoryById)
	pc.v1.Post(AdvertisementCategory, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerCreateAdvertisementCategory)
	pc.v1.Put(AdvertisementCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerUpdateAdvertisementCategory)
	pc.v1.Delete(AdvertisementCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerDeleteAdvertisementCategory)

	// ---- Advertisement
	pc.v1.Get(GetAdvertisementPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllAdvertisementPluck)
	pc.v1.Get(GetAdvertisementPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA}), pc.handlerGetAdvertisementPaginated)
	pc.v1.Get(AdvertisementById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA}), pc.handlerGetAdvertisementById)
	pc.v1.Post(Advertisement, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA}), pc.handlerCreateAdvertisement)
	pc.v1.Put(AdvertisementById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA}), pc.handlerUpdateAdvertisement)
	pc.v1.Delete(AdvertisementById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA}), pc.handlerDeleteAdvertisement)

	// ----- Queue
	pc.v1.Get(GetQueuePaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN, role.ROLE_KASIR, role.ROLE_FARMASI}), pc.handlerGetQueuePaginated)
	pc.v1.Get(QueueById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN, role.ROLE_KASIR, role.ROLE_FARMASI}), pc.handlerGetQueueById)
	pc.v1.Post(Queue, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN, role.ROLE_KASIR}), pc.handlerCreateQueue)
	pc.v1.Put(QueueStatusById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN, role.ROLE_KASIR, role.ROLE_FARMASI}), pc.handlerUpdateStatusQueue)
	pc.v1.Put(QueueById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN, role.ROLE_KASIR, role.ROLE_FARMASI}), pc.handlerUpdateQueueById)
	pc.v1.Delete(QueueById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN, role.ROLE_KASIR, role.ROLE_FARMASI}), pc.handlerDeleteQueue)

	// ----- Analytic
	pc.v1.Get(AnalyticDashboard, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA, role.ROLE_SUPER_ADMIN}), pc.handlerGetDashboardAnalytic)

	// ----- Display dashboard
	pc.v1.Get(DisplayDashboard, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA, role.ROLE_SUPER_ADMIN}), pc.handlerDisplayQueue)
	pc.v1.Get(QueueBySearch+"/search", pc.handlerQueueBySearch)
	pc.v1.Put(UpdateQueueFu, pc.handlerUpdateFollowUpPhone)

	// ----- Content
	pc.v1.Get(AdvertisementContent, pc.handlerContentAdvertisement)
}
