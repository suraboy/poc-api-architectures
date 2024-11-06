package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/suraboy/poc-api-architectures/internal"
	"github.com/suraboy/poc-api-architectures/internal/rest"
	"github.com/suraboy/poc-api-architectures/internal/soap"
)

func main() {
	conf, err := internal.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	e := echo.New()

	soapSvc := soap.NewService(conf)
	soapHdl := soap.NewHandler(soapSvc)

	restSvc := rest.NewService(conf)
	restHdl := rest.NewHandler(restSvc)

	e.GET("/soap/user", soapHdl.GetUser)
	e.GET("/rest/user", restHdl.GetUser)

	err = e.Start(fmt.Sprintf(":%s", conf.Server.Port))
	if err != nil {
		return
	}
	defer e.Shutdown(context.Background())
}
