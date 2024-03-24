package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/maxzycon/rs-informasi-be/internal/config"
	S3Controller "github.com/maxzycon/rs-informasi-be/internal/domain/s3/controller"
	S3Repository "github.com/maxzycon/rs-informasi-be/internal/domain/s3/repository/impl"
	S3Service "github.com/maxzycon/rs-informasi-be/internal/domain/s3/service/impl"
	UserController "github.com/maxzycon/rs-informasi-be/internal/domain/user/controller"
	UserRepository "github.com/maxzycon/rs-informasi-be/internal/domain/user/repository/impl"
	UserService "github.com/maxzycon/rs-informasi-be/internal/domain/user/service/impl"
	"github.com/maxzycon/rs-informasi-be/pkg/database"
	middleware2 "github.com/maxzycon/rs-informasi-be/pkg/middleware"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/sirupsen/logrus"

	GlobalController "github.com/maxzycon/rs-informasi-be/internal/domain/global/controller"
	GlobalRepository "github.com/maxzycon/rs-informasi-be/internal/domain/global/repository/impl"
	GlobalService "github.com/maxzycon/rs-informasi-be/internal/domain/global/service/impl"
	"github.com/mikhail-bigun/fiberlogrus"
)

type InitWebserviceParam struct {
	Conf *config.Config
}

func InitWebservice(params *InitWebserviceParam) {
	app := fiber.New(fiber.Config{
		BodyLimit: 1000 * 1024 * 1024, // set 1000MB
	})

	db, err := database.InitMariaDB(&database.InitMariaDBParams{
		Conf: &params.Conf.MariaDBConfig,
	})

	log := logrus.New()
	log.SetReportCaller(true)

	db.AutoMigrate(
		// --- global
		&model.User{},
		&model.Merchant{},
		&model.MerchantCategory{},
		&model.Floor{},
		&model.Facility{},
		&model.Services{},
		&model.ProductCategory{},
		&model.Product{},
		&model.DetailProduct{},
		&model.InformationCategory{},
		&model.Information{},
		&model.Specialization{},
		&model.Doctor{},
		&model.DoctorEducation{},
		&model.DoctorSkill{},
		&model.DoctorSlot{},
		&model.AdvertisementCategory{},
		&model.Advertisement{},
		&model.Organ{},
		&model.Room{},
		&model.LogsPage{},
	)

	if err != nil {
		log.Errorf("[webservice.go][InitWebservice] err init mysql :%+v", err)
		return
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*, https://information.devku.xyz",
		AllowOriginsFunc: func(origin string) bool {
			return params.Conf.ENVIRONMENT == "dev"
		},
	}))

	// app.Use(logger.New())  // Logger middleware
	app.Use(recover.New()) // --- recover panic
	app.Use(fiberlogrus.New(fiberlogrus.Config{
		Tags: []string{
			fiberlogrus.TagPath,
			fiberlogrus.TagHost,
			fiberlogrus.TagQueryStringParams,
			// add method field
			fiberlogrus.TagMethod,
			// add status field
			fiberlogrus.TagStatus,
			// add value from locals
			fiberlogrus.AttachKeyTag(fiberlogrus.TagLocals, "requestid"),
			// add certain header
			fiberlogrus.AttachKeyTag(fiberlogrus.TagReqHeader, "custom-header"),
		},
	})) // ----- logger

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")
	api.Get("/health", monitor.New())

	// ------- User
	userRepository := UserRepository.New(&UserRepository.NewUserRepository{Conf: params.Conf, Db: db})
	userService := UserService.New(&UserService.NewUserServiceParams{Conf: params.Conf, UserRepository: userRepository})

	// ------  middleware
	middleware := middleware2.GlobalMiddleware{
		UserService: userService,
		Conf:        params.Conf,
	}

	userController := UserController.New(&UserController.UsersControllerParams{
		V1:          v1,
		Conf:        params.Conf,
		UserService: userService,
		Middleware:  middleware,
	})

	// -------- AWS S3
	s3Repository := S3Repository.New(&S3Repository.NewS3RepositoryParams{
		Conf: params.Conf,
		Db:   db,
	})
	s3Service := S3Service.New(&S3Service.NewS3ServiceParams{
		Conf:         params.Conf,
		S3Repository: s3Repository,
	})

	s3Controller := S3Controller.New(&S3Controller.S3ControllerParams{
		V1:        v1,
		Conf:      params.Conf,
		S3Service: s3Service,
	})

	// --------- Main
	globalRepository := GlobalRepository.New(&GlobalRepository.NewGlobalRepository{
		Conf: params.Conf,
		Db:   db,
		Log:  log,
	})

	globalService := GlobalService.New(&GlobalService.NewGlobalServiceParams{
		Conf:             params.Conf,
		GlobalRepository: globalRepository,
		S3Service:        s3Service,
		Db:               db,
		Log:              log,
	})

	globalController := GlobalController.New(&GlobalController.GlobalControllerParams{
		V1:            v1,
		Conf:          params.Conf,
		GlobalService: globalService,
		Middleware:    middleware,
		Log:           log,
	})

	userController.Init()
	s3Controller.Init()
	globalController.Init()

	app.Listen(fmt.Sprintf(":%s", params.Conf.AppAddress))
}
