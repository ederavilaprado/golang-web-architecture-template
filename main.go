package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/ederavilaprado/golang-web-architecture-template/apis"
	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/ederavilaprado/golang-web-architecture-template/daos"
	"github.com/ederavilaprado/golang-web-architecture-template/errors"
	"github.com/ederavilaprado/golang-web-architecture-template/services"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// loading app flags before start the server
	app.LoadAppFlags()

	// load application configurations
	if err := app.LoadConfig(); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	// Starting DB...
	db, err := sqlx.Connect("postgres", app.Config.DSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	// // TODO: find a good number for the connection pool
	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(10)

	// wire up API routing
	http.Handle("/", buildRouter(logger, db))

	// start the server
	address := fmt.Sprintf("%v:%v", app.Config.Host, app.Config.Port)
	logger.Infof("Server (%v) is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *sqlx.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort() // skip all other middlewares/handlers
		return c.Write("OK " + app.Version)
	})

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		// TODO: create an middleware helper for transaction...
		// app.Transactional(db),
	)

	rg := router.Group("/v1")

	// TODO: JWT + Session here...

	rg.Post("/auth", apis.Auth(app.Config.JWTSigningKey))
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  apis.JWTHandler,
	}))

	artistDAO := daos.NewArtistDAO(db)
	apis.ServeArtistResource(rg, services.NewArtistService(artistDAO))

	// wire up more resource APIs here

	return router
}
