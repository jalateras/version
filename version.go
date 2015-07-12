package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func mountVersionRoutes(server *echo.Echo) {
	server.Get("/", getVersion)
}

func getVersion(ctx *echo.Context) error {
	ctx.Response().Header().Set(echo.ContentType, echo.ApplicationJSON)
	return ctx.JSON(http.StatusOK, VersionInfo{
		Name:    APP_NAME,
		Version: APP_VERSION,
		Author:  APP_AUTHOR,
	})
}
