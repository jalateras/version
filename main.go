package main

import (
	"net/http"
	"os"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
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
	app.Action = launchService

	app.Run(os.Args)
}

func launchService(ctx *cli.Context) {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", HomeHandler)

	server := negroni.Classic();
	server.UseHandler(router)
	server.Run(ctx.String("host") + ":" + strconv.Itoa(ctx.Int("port")))
	fmt.Println("Server running on :" + strconv.Itoa(ctx.Int("port")))
}

func HomeHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Home")
}
