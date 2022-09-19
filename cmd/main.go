package main

import (
	"os"

	"github.com/Gym-Apps/gym-backend/internal/config/db"
	"github.com/Gym-Apps/gym-backend/router"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("gymapp"),
		newrelic.ConfigLicense("eu01xx250142fac95cf54b8f52f65ea9f8d4NRAL"),
		newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(config *newrelic.Config) {
			logrus.SetLevel(logrus.DebugLevel)
			config.Logger = nrlogrus.StandardLogger()
		},
	)
	if err != nil {
		panic(err)
	}

	db := db.Connect(app)

	e := echo.New()
	router.Init(e, db, app)
	e.Logger.Fatal(e.Start(":8080"))
}
