package app

import (
	"TestErajaya/config"
	"TestErajaya/setup"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func RunApp() {
	var err error
	config.DBConnection, err = config.StartConnectionDb(config.DBConnectionStruct{
		Username: "root",
		Password: "secret",
		Port:     "5432",
		Hostname: "localhost",
		DBName:   "erajaya",
	})
	if err != nil {
		log.Fatal(err)
	}

	config.StartConnectingRedis()

	err = config.MigrateModel(config.DBConnection, setup.GetListModel()...)
	if err != nil {
		log.Fatal(err)
	}

	repository := setup.RegisterRepository(config.DBConnection)

	useCase := setup.RegisterUseCase(&repository, config.DBConnection)

	router := setup.RegisterRouter(&useCase)

	// start server
	fmt.Println("start server ...")
	app := fiber.New()

	// setup routing
	setup.SetupRouting(router, app)

	err = app.Listen(":3400")
	if err != nil {
		log.Fatal(err)
	}

}
