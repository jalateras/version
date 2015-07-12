package main

import (
	"os"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Version = APP_VERSION
	app.Author = APP_AUTHOR
	app.Usage = APP_USAGE
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Value:  "0.0.0.0",
			Usage:  "The host interface that this service should bind too",
			EnvVar: "HOST",
		},
		cli.IntFlag{
			Name:   "port",
			Value:  3000,
			Usage:  "The port that this service will listen on",
			EnvVar: "PORT",
		},
		cli.BoolFlag{
			Name:   "debug, d",
			Usage:  "Run in debug mode",
			EnvVar: "DEBUG",
		},
	}
	app.Action = createService()
	app.Run(os.Args)
}

func createService() func(*cli.Context) {

	return func(ctx *cli.Context) {
		logger := getLogger()

		server := echo.New()
		server.SetDebug(ctx.Bool("debug"))
		server.Use(getMiddleWareLogger())
		server.Use(mw.Recover())

		// Mount all the routes
		mountVersionRoutes(server)

		// Start server
		endpoint := ctx.String("host") + ":" + strconv.Itoa(ctx.Int("port"))
		logger.WithFields(logrus.Fields{
			"endpoint": endpoint,
		}).Info("Starting server")
		server.Run(endpoint)
	}
}
