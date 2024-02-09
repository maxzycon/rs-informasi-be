package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/cmd"
	"github.com/maxzycon/rs-farmasi-be/internal/config"
)

func main() {
	config.Init()
	conf := config.Get()
	log.Info("[InitialEnv] env set successfully")
	cmd.InitWebservice(&cmd.InitWebserviceParam{Conf: conf})
}
