package repository

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

type GlobalRepository interface {
	// ---- users
	FindAllUser(ctx context.Context, claims *authutil.UserClaims) (resp []*model.User, err error)

	// ---- Floors
	FindFloorById(ctx context.Context, id int) (resp *model.Floor, err error)
	FindAllFloor(ctx context.Context) (resp []*model.Floor, err error)
	FindFloorPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateFloor(ctx context.Context, entity *model.Floor) (resp *int64, err error)
	UpdateFloorById(ctx context.Context, id int, entity *model.Floor) (resp *int64, err error)
	DeleteFloorById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Rooms
	FindRoomById(ctx context.Context, id int) (resp *model.Room, err error)
	FindAllRoom(ctx context.Context) (resp []*model.Room, err error)
	FindRoomPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateRoom(ctx context.Context, entity *model.Room) (resp *int64, err error)
	UpdateRoomById(ctx context.Context, id int, entity *model.Room) (resp *int64, err error)
	DeleteRoomById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Merchant Category
	FindMerchantCategoryById(ctx context.Context, id int) (resp *model.MerchantCategory, err error)
	FindAllMerchantCategory(ctx context.Context) (resp []*model.MerchantCategory, err error)
	FindMerchantCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateMerchantCategory(ctx context.Context, entity *model.MerchantCategory) (resp *int64, err error)
	UpdateMerchantCategoryById(ctx context.Context, id int, entity *model.MerchantCategory) (resp *int64, err error)
	DeleteMerchantCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Merchant Specialization
	FindMerchantSpecializationById(ctx context.Context, id int) (resp *model.Specialization, err error)
	FindAllMerchantSpecialization(ctx context.Context) (resp []*model.Specialization, err error)
	FindMerchantSpecializationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateMerchantSpecialization(ctx context.Context, entity *model.Specialization) (resp *int64, err error)
	UpdateMerchantSpecializationById(ctx context.Context, id int, entity *model.Specialization) (resp *int64, err error)
	DeleteMerchantSpecializationById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Merchant Information category
	FindInformationCategoryById(ctx context.Context, id int) (resp *model.InformationCategory, err error)
	FindAllInformationCategory(ctx context.Context) (resp []*model.InformationCategory, err error)
	FindInformationCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateInformationCategory(ctx context.Context, entity *model.InformationCategory) (resp *int64, err error)
	UpdateInformationCategoryById(ctx context.Context, id int, entity *model.InformationCategory) (resp *int64, err error)
	DeleteInformationCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Product Category
	FindProductCategoryById(ctx context.Context, id int) (resp *model.ProductCategory, err error)
	FindAllProductCategory(ctx context.Context) (resp []*model.ProductCategory, err error)
	FindProductCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateProductCategory(ctx context.Context, entity *model.ProductCategory) (resp *int64, err error)
	UpdateProductCategoryById(ctx context.Context, id int, entity *model.ProductCategory) (resp *int64, err error)
	DeleteProductCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Merchant Facility
	FindFacilityById(ctx context.Context, id int) (resp *model.Facility, err error)
	FindAllFacility(ctx context.Context) (resp []*model.Facility, err error)
	FindFacilityPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateFacility(ctx context.Context, entity *model.Facility) (resp *int64, err error)
	UpdateFacilityById(ctx context.Context, id int, entity *model.Facility) (resp *int64, err error)
	DeleteFacilityById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Organs
	FindOrganById(ctx context.Context, id int) (resp *model.Organ, err error)
	FindAllOrgan(ctx context.Context) (resp []*model.Organ, err error)
	FindOrganPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateOrgan(ctx context.Context, entity *model.Organ) (resp *int64, err error)
	UpdateOrganById(ctx context.Context, id int, entity *model.Organ) (resp *int64, err error)
	DeleteOrganById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Services
	FindServicesById(ctx context.Context, id int) (resp *model.Services, err error)
	FindAllServices(ctx context.Context) (resp []*model.Services, err error)
	FindServicesPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateServices(ctx context.Context, entity *model.Services) (resp *int64, err error)
	UpdateServicesById(ctx context.Context, id int, entity *model.Services) (resp *int64, err error)
	DeleteServicesById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Information
	FindInformationById(ctx context.Context, id int) (resp *model.Information, err error)
	FindAllInformation(ctx context.Context) (resp []*model.Information, err error)
	FindInformationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateInformation(ctx context.Context, entity *model.Information) (resp *int64, err error)
	UpdateInformationById(ctx context.Context, id int, entity *model.Information) (resp *int64, err error)
	DeleteInformationById(ctx context.Context, id int) (resp *int64, err error)

	// ---- Product
	FindProductById(ctx context.Context, id int) (resp *model.Product, err error)
	FindAllProduct(ctx context.Context) (resp []*model.Product, err error)
	FindProductPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateProduct(ctx context.Context, entity *model.Product) (resp *int64, err error)
	UpdateProductById(ctx context.Context, id int, entity *model.Product) (resp *int64, err error)
	DeleteProductById(ctx context.Context, id int) (resp *int64, err error)

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

	// ---- Doctors
	FindDoctorById(ctx context.Context, id int) (resp *model.Doctor, err error)
	FindAllDoctor(ctx context.Context) (resp []*model.Doctor, err error)
	FindDoctorPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateDoctor(ctx context.Context, entity *model.Doctor) (resp *int64, err error)
	UpdateDoctorById(ctx context.Context, id int, entity *model.Doctor) (resp *int64, err error)
	DeleteDoctorById(ctx context.Context, id int) (resp *int64, err error)
}
