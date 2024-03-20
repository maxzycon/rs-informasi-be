package database

import (
	"fmt"
	"time"

	"github.com/maxzycon/rs-informasi-be/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2/log"
)

type InitMariaDBParams struct {
	Conf *config.MariaDBConfig
}

func InitMariaDB(params *InitMariaDBParams) (db *gorm.DB, err error) {

	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true&loc=%s&charset=utf8mb4",
		params.Conf.Username, params.Conf.Password,
		params.Conf.Address, params.Conf.DBName, "Local",
	)

	for i := 10; i > 0; i-- {
		db, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})
		if err == nil {
			break
		}
		log.Errorf("[InitMariaDB] error init opening db for %s: %+v, retrying in %d second", dataSource, err, i)
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return
	}

	for i := 10; i > 0; i-- {
		err = db.Error

		if err == nil {
			break
		}
		log.Errorf("[InitMariaDB] error ping db for %s: %+v, retrying in %d second", dataSource, err, i)
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return
	}

	log.Info("[InitMariaDB] db init successfully")
	return
}
