package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/config"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/service"
	"github.com/maxzycon/rs-informasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-informasi-be/pkg/middleware"
	"github.com/sirupsen/logrus"
)

const (
	GetUserPluck = "/users_pluck"

	GetFloorPluck     = "floors/list"
	GetFloorUser      = "floors/user"
	GetFloorPaginated = "floors/paginated"
	Floor             = "floors"
	FloorById         = "floors/:id"

	GetRoomPluck     = "rooms/list"
	GetRoomPaginated = "rooms/paginated"
	Room             = "rooms"
	RoomById         = "rooms/:id"

	GetFacilityPluck     = "facilities/list"
	GetFacilityPaginated = "facilities/paginated"
	Facility             = "facilities"
	FacilityById         = "facilities/:id"

	GetInformationCategoryPluck     = "information_categories/list"
	GetInformationCategoryPaginated = "information_categories/paginated"
	InformationCategory             = "information_categories"
	InformationCategoryById         = "information_categories/:id"

	GetProductCategoryPluck     = "product_categories/list"
	GetProductCategoryPaginated = "product_categories/paginated"
	ProductCategory             = "product_categories"
	ProductCategoryById         = "product_categories/:id"

	GetOrganPluck     = "organs/list"
	GetOrganPaginated = "organs/paginated"
	Organ             = "organs"
	OrganById         = "organs/:id"

	GetServicePluck     = "services/list"
	GetServicePaginated = "services/paginated"
	Service             = "services"
	ServiceById         = "services/:id"

	GetInformationPluck     = "informations/list"
	GetInformationPaginated = "informations/paginated"
	Information             = "informations"
	InformationById         = "informations/:id"

	GetProductPluck     = "products/list"
	GetProductPaginated = "products/paginated"
	Product             = "products"
	ProductById         = "products/:id"

	GetDoctorPluck     = "doctors/list"
	GetDoctorPaginated = "doctors/paginated"
	Doctor             = "doctors"
	DoctorById         = "doctors/:id"

	GetMerchantCategoryPluck     = "merchant_categories/list"
	GetMerchantCategoryPaginated = "merchant_categories/paginated"
	MerchantCategory             = "merchant_categories"
	MerchantCategoryById         = "merchant_categories/:id"

	GetMerchantSpecializationPluck     = "merchant_specializations/list"
	GetMerchantSpecializationPaginated = "merchant_specializations/paginated"
	MerchantSpecialization             = "merchant_specializations"
	MerchantSpecializationById         = "merchant_specializations/:id"

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

	AnalyticDashboard = "analytic/dashboard"
	DisplayDashboard  = "dashboard"
	RunningText       = "running_text/:id"
)

type GlobalControllerParams struct {
	V1            fiber.Router
	Conf          *config.Config
	GlobalService service.GlobalService
	Middleware    middleware.GlobalMiddleware
	Log           *logrus.Logger
}
type GlobalController struct {
	v1            fiber.Router
	conf          *config.Config
	globalService service.GlobalService
	middleware    middleware.GlobalMiddleware
	log           *logrus.Logger
}

func New(params *GlobalControllerParams) *GlobalController {
	return &GlobalController{
		v1:            params.V1,
		conf:          params.Conf,
		globalService: params.GlobalService,
		middleware:    params.Middleware,
		log:           params.Log,
	}
}

func (pc *GlobalController) Init() {
	// ---- User
	pc.v1.Get(GetUserPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllUserPluck)

	// ---- Floor
	pc.v1.Get(GetFloorPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllFloorPluck)
	pc.v1.Get(GetFloorUser, pc.middleware.Protected([]uint{role.ROLE_SUPER_ADMIN}), pc.handlerGetAllFloorUser)
	pc.v1.Get(GetFloorPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetFloorPaginated)
	pc.v1.Get(FloorById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetFloorById)
	pc.v1.Post(Floor, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateFloor)
	pc.v1.Put(FloorById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateFloor)
	pc.v1.Delete(FloorById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteFloor)

	// ---- Rooms
	pc.v1.Get(GetRoomPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllRoomPluck)
	pc.v1.Get(GetRoomPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetRoomPaginated)
	pc.v1.Get(RoomById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetRoomById)
	pc.v1.Post(Room, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateRoom)
	pc.v1.Put(RoomById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateRoom)
	pc.v1.Delete(RoomById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteRoom)

	// ---- Facility
	pc.v1.Get(GetFacilityPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllFacilityPluck)
	pc.v1.Get(GetFacilityPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetFacilityPaginated)
	pc.v1.Get(FacilityById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetFacilityById)
	pc.v1.Post(Facility, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateFacility)
	pc.v1.Put(FacilityById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateFacility)
	pc.v1.Delete(FacilityById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteFacility)

	// ---- Information categories
	pc.v1.Get(GetInformationCategoryPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllInformationCategoryPluck)
	pc.v1.Get(GetInformationCategoryPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetInformationCategoryPaginated)
	pc.v1.Get(InformationCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetInformationCategoryById)
	pc.v1.Post(InformationCategory, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateInformationCategory)
	pc.v1.Put(InformationCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateInformationCategory)
	pc.v1.Delete(InformationCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteInformationCategory)

	// ---- Product categories
	pc.v1.Get(GetProductCategoryPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllProductCategoryPluck)
	pc.v1.Get(GetProductCategoryPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetProductCategoryPaginated)
	pc.v1.Get(ProductCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetProductCategoryById)
	pc.v1.Post(ProductCategory, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateProductCategory)
	pc.v1.Put(ProductCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateProductCategory)
	pc.v1.Delete(ProductCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteProductCategory)

	// ---- Organs
	pc.v1.Get(GetOrganPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllOrganPluck)
	pc.v1.Get(GetOrganPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetOrganPaginated)
	pc.v1.Get(OrganById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetOrganById)
	pc.v1.Post(Organ, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateOrgan)
	pc.v1.Put(OrganById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateOrgan)
	pc.v1.Delete(OrganById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteOrgan)

	// ---- Services
	pc.v1.Get(GetServicePluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllServicePluck)
	pc.v1.Get(GetServicePaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetServicePaginated)
	pc.v1.Get(ServiceById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetServiceById)
	pc.v1.Post(Service, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateService)
	pc.v1.Put(ServiceById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateService)
	pc.v1.Delete(ServiceById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteService)

	// ---- Information
	pc.v1.Get(GetInformationPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllInformationPluck)
	pc.v1.Get(GetInformationPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetInformationPaginated)
	pc.v1.Get(InformationById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetInformationById)
	pc.v1.Post(Information, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateInformation)
	pc.v1.Put(InformationById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateInformation)
	pc.v1.Delete(InformationById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteInformation)

	// ---- Product
	pc.v1.Get(GetProductPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllProductPluck)
	pc.v1.Get(GetProductPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetProductPaginated)
	pc.v1.Get(ProductById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetProductById)
	pc.v1.Post(Product, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateProduct)
	pc.v1.Put(ProductById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateProduct)
	pc.v1.Delete(ProductById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteProduct)

	// ---- Doctors
	pc.v1.Get(GetDoctorPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllDoctorPluck)
	pc.v1.Get(GetDoctorPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetDoctorPaginated)
	pc.v1.Get(DoctorById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetDoctorById)
	pc.v1.Post(Doctor, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateDoctor)
	pc.v1.Put(DoctorById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateDoctor)
	pc.v1.Delete(DoctorById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteDoctor)

	// ---- MerchantCategory
	pc.v1.Get(GetMerchantCategoryPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllMerchantCategoryPluck)
	pc.v1.Get(GetMerchantCategoryPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING}), pc.handlerGetMerchantCategoryPaginated)
	pc.v1.Get(MerchantCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerGetMerchantCategoryById)
	pc.v1.Post(MerchantCategory, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerCreateMerchantCategory)
	pc.v1.Put(MerchantCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerUpdateMerchantCategory)
	pc.v1.Delete(MerchantCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerDeleteMerchantCategory)

	// ---- Specialization
	pc.v1.Get(GetMerchantSpecializationPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllMerchantSpecializationPluck)
	pc.v1.Get(GetMerchantSpecializationPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetMerchantSpecializationPaginated)
	pc.v1.Get(MerchantSpecializationById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetMerchantSpecializationById)
	pc.v1.Post(MerchantSpecialization, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateMerchantSpecialization)
	pc.v1.Put(MerchantSpecializationById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateMerchantSpecialization)
	pc.v1.Delete(MerchantSpecializationById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteMerchantSpecialization)

	// ---- Merchant
	pc.v1.Get(GetMerchantPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllMerchantPluck)
	pc.v1.Get(GetMerchantPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_MARKETING}), pc.handlerGetMerchantPaginated)
	pc.v1.Get(MerchantById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerGetMerchantById)
	pc.v1.Post(Merchant, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerCreateMerchant)
	pc.v1.Put(MerchantById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerUpdateMerchant)
	pc.v1.Put(MerchantConfigById, pc.middleware.Protected([]uint{role.ROLE_SUPER_ADMIN}), pc.handlerUpdateMerchantConfig)
	pc.v1.Delete(MerchantById, pc.middleware.Protected([]uint{role.ROLE_OWNER}), pc.handlerDeleteMerchant)

	// ---- AdvertisementCategory
	pc.v1.Get(GetAdvertisementCategoryPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllAdvertisementCategoryPluck)
	pc.v1.Get(GetAdvertisementCategoryPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAdvertisementCategoryPaginated)
	pc.v1.Get(AdvertisementCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAdvertisementCategoryById)
	pc.v1.Post(AdvertisementCategory, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateAdvertisementCategory)
	pc.v1.Put(AdvertisementCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateAdvertisementCategory)
	pc.v1.Delete(AdvertisementCategoryById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteAdvertisementCategory)

	// ---- Advertisement
	pc.v1.Get(GetAdvertisementPluck, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAllAdvertisementPluck)
	pc.v1.Get(GetAdvertisementPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAdvertisementPaginated)
	pc.v1.Get(AdvertisementById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetAdvertisementById)
	pc.v1.Post(Advertisement, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateAdvertisement)
	pc.v1.Put(AdvertisementById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateAdvertisement)
	pc.v1.Delete(AdvertisementById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteAdvertisement)

	// ----- Analytic
	pc.v1.Get(AnalyticDashboard, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetDashboardAnalytic)

	// ----- Content
	pc.v1.Get(AdvertisementContent, pc.handlerContentAdvertisement)
	pc.v1.Get("advertisements/merchant/:id", pc.handlerMerchantAdvertisement)

	pc.v1.Get(RunningText, pc.handlerGetRunningTextByMerchantIdStr)
}
