package main

const (
	APP_NAME    = "version"
	APP_VERSION = "0.0.2"
	APP_AUTHOR  = "Jim Alateras"
	APP_USAGE   = `Version Microservice

This service support the PageLoad Version API
`
)

type VersionInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Author  string `json:"author"`
}
