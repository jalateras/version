package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/codegangsta/cli"
)

const (
	APP_NAME = "version"
	APP_VERSION = "0.0.1"
	APP_AUTHOR = "PageLoad"
	APP_USAGE = `Version Microservice

This service support the PageLoad Version API
`
)

type VersionInfo struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Author string `json:"author"`
}

func main()  {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Version = APP_VERSION
	app.Author = APP_AUTHOR
	app.Usage = APP_USAGE
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "0.0.0.0",
			Usage: "The host interface that this service should bind too",
			EnvVar: "HOST",
		},
		cli.IntFlag{
			Name: "port",
			Value: 3000,
			Usage: "The port that this service will listen on",
			EnvVar: "PORT",
		},
	}
	app.Action = createService()

	app.Run(os.Args)
}

func createService() func(*cli.Context) {
	return func(ctx *cli.Context) {
		server := echo.New()
		server.SetDebug(false)
		server.Use(mw.Logger())
		server.Use(mw.Recover())

		server.Get("/", HomeHandler)

		server.Run(ctx.String("host") + ":" + strconv.Itoa(ctx.Int("port")))
	}
}

func HomeHandler(ctx *echo.Context) error {
	versionInfo := VersionInfo{
		Name: APP_NAME,
		Version: APP_VERSION,
		Author: APP_AUTHOR,
	}

	return ctx.JSON(http.StatusOK, &versionInfo)
}
