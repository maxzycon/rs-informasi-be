package service

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"

	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
)

type GlobalService interface {
	GetAllUserPluck(ctx context.Context, user *authutil.UserClaims) (resp []*dto.UserRowPluck, err error)

	// ---- Floor
	GetFloorPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetFloorPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetFloorById(ctx context.Context, id int) (resp *dto.FloorRow, err error)
	GetAllFloorByUser(ctx context.Context) (resp []*dto.FloorUserRow, err error)
	CreateFloor(ctx context.Context, payload *dto.PayloadFloor) (resp *int64, err error)
	UpdateFloorById(ctx context.Context, id int, payload *dto.PayloadFloor) (resp *int64, err error)
	DeleteFloorById(ctx context.Context, id int) (resp *int64, err error)

	// ---- MerchantCategory
	GetMerchantCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetMerchantCategoryPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetMerchantCategoryById(ctx context.Context, id int) (resp *dto.MerchantCategoryRow, err error)
	CreateMerchantCategory(ctx context.Context, payload *dto.PayloadMerchantCategory) (resp *int64, err error)
	UpdateMerchantCategoryById(ctx context.Context, id int, payload *dto.PayloadMerchantCategory) (resp *int64, err error)
	DeleteMerchantCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- MerchantSpecialization
	GetMerchantSpecializationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetMerchantSpecializationPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetMerchantSpecializationById(ctx context.Context, id int) (resp *dto.MerchantSpecializationRow, err error)
	CreateMerchantSpecialization(ctx context.Context, payload *dto.PayloadMerchantSpecialization) (resp *int64, err error)
	UpdateMerchantSpecializationById(ctx context.Context, id int, payload *dto.PayloadMerchantSpecialization) (resp *int64, err error)
	DeleteMerchantSpecializationById(ctx context.Context, id int) (resp *int64, err error)

	// ---- MerchantInformationCategory
	GetInformationCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetInformationCategoryPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetInformationCategoryById(ctx context.Context, id int) (resp *dto.InformationCategoryRow, err error)
	CreateInformationCategory(ctx context.Context, payload *dto.PayloadInformationCategory) (resp *int64, err error)
	UpdateInformationCategoryById(ctx context.Context, id int, payload *dto.PayloadInformationCategory) (resp *int64, err error)
	DeleteInformationCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- ProductCategory
	GetProductCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetProductCategoryPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetProductCategoryById(ctx context.Context, id int) (resp *dto.ProductCategoryRow, err error)
	CreateProductCategory(ctx context.Context, payload *dto.PayloadProductCategory) (resp *int64, err error)
	UpdateProductCategoryById(ctx context.Context, id int, payload *dto.PayloadProductCategory) (resp *int64, err error)
	DeleteProductCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Facility
	GetFacilityPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetFacilityPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetFacilityById(ctx context.Context, id int) (resp *dto.FacilityRow, err error)
	CreateFacility(ctx context.Context, payload *dto.PayloadFacility) (resp *int64, err error)
	UpdateFacilityById(ctx context.Context, id int, payload *dto.PayloadFacility) (resp *int64, err error)
	DeleteFacilityById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Organ
	GetOrganPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetOrganPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetOrganById(ctx context.Context, id int) (resp *dto.OrganRow, err error)
	CreateOrgan(ctx context.Context, payload *dto.PayloadOrgan) (resp *int64, err error)
	UpdateOrganById(ctx context.Context, id int, payload *dto.PayloadOrgan) (resp *int64, err error)
	DeleteOrganById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Information
	GetInformationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetInformationPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetInformationById(ctx context.Context, id int) (resp *dto.InformationRow, err error)
	CreateInformation(ctx context.Context, payload *dto.PayloadInformation) (resp *int64, err error)
	UpdateInformationById(ctx context.Context, id int, payload *dto.PayloadInformation) (resp *int64, err error)
	DeleteInformationById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Product
	GetProductPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetProductPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetProductById(ctx context.Context, id int) (resp *dto.ProductRow, err error)
	CreateProduct(ctx context.Context, payload *dto.PayloadProduct) (resp *int64, err error)
	UpdateProductById(ctx context.Context, id int, payload *dto.PayloadProduct) (resp *int64, err error)
	DeleteProductById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Service
	GetServicePaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetServicePluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetServiceById(ctx context.Context, id int) (resp *dto.ServiceRow, err error)
	CreateService(ctx context.Context, payload *dto.PayloadService) (resp *int64, err error)
	UpdateServiceById(ctx context.Context, id int, payload *dto.PayloadService) (resp *int64, err error)
	DeleteServiceById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Doctor
	GetDoctorPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	GetDoctorPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error)
	GetDoctorById(ctx context.Context, id int) (resp *dto.DoctorByIdRow, err error)
	CreateDoctor(ctx context.Context, payload *dto.PayloadDoctor) (resp *int64, err error)
	UpdateDoctorById(ctx context.Context, id int, payload *dto.PayloadDoctor) (resp *int64, err error)
	DeleteDoctorById(ctx context.Context, id int) (resp *int64, err error)

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

	// ----- Analytic
	GetDashboardAnalytic(ctx context.Context) (resp *dto.SummaryDashboardWrapper, err error)

	// ----- Dashboard queue
	GetMerchantDetailAdvertisement(ctx context.Context, merchantIdStr string) (resp *dto.AdvertisementMerchant, err error)
}
