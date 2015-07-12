package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"os"
	"time"
)

var (
	logger *log.Logger
)

func init() {
	log.SetLevel(log.InfoLevel)
	logger = log.New()
	logger.Out = os.Stderr
}

func getLogger() *log.Logger {
	return logger
}

func getMiddleWareLogger() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			start := time.Now()

			entry := logger.WithFields(log.Fields{
				"request": c.Request().RequestURI,
				"method":  c.Request().Method,
				"remote":  c.Request().RemoteAddr,
			})

			if reqID := c.Request().Header.Get("X-Request-Id"); reqID != "" {
				entry = entry.WithField("request_id", reqID)
			}

			if err := next(c); err != nil {
				c.Error(err)
			}

			latency := time.Since(start)
			entry.WithFields(log.Fields{
				"status": c.Response().Status(),
				"took":   latency,
			}).Info("Handled request")

			return nil
		}
	}
}
