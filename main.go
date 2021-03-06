package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/ederavilaprado/golang-web-architecture-template/apis"
	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/ederavilaprado/golang-web-architecture-template/daos"
	"github.com/ederavilaprado/golang-web-architecture-template/errors"
	"github.com/ederavilaprado/golang-web-architecture-template/services"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	// creating the logger
	logger := logrus.New()
	// starting DB...
	db, err := sqlx.Connect("postgres", app.Config.DSN)
	if err != nil {
		logger.Fatalln(err)
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

func buildRouter(logger *logrus.Logger, db *sqlx.DB) *echo.Echo {
	router := echo.New()

	// When running in debug mode,the returned JSON is always "pretty printed"
	router.Debug = app.Config.Debug

	// Removing unnecessary trailing slash at the end of the path
	router.Pre(middleware.RemoveTrailingSlash())

	// TODO: implement 2 variables to healthcheck... one "completely" open,
	// another closed and with the power to check the entire status of the app
	router.Match([]string{"GET", "HEAD"}, "/health", func(c echo.Context) error {
		return c.String(http.StatusOK, app.Version)
	})

	router.Use(
		app.Init(logger),
		app.RequestContextMiddleware(logger),
		middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize: 1 << 10, // 1 KB
		}),
	)

	// Handling all the panic errors for the api
	// Needs the "recover" middleware
	router.HTTPErrorHandler = func(err error, c echo.Context) {
		c.HTML(http.StatusInternalServerError, "Internal Server Error... from panic")
	}

	// 	cors.Handler(cors.Options{
	// 		AllowOrigins: "*",
	// 		AllowHeaders: "*",
	// 		AllowMethods: "*",
	// 	}),
	// 	// TODO: create an middleware helper for transaction...
	// 	// app.Transactional(db),
	// )

	rg := router.Group("/v1")

	// TODO: JWT + Session here...
	// rg.Post("/auth", apis.Auth(app.Config.JWTSigningKey))
	// rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
	// 	SigningMethod: app.Config.JWTSigningMethod,
	// 	TokenHandler:  apis.JWTHandler,
	// }))

	artistDAO := daos.NewArtistDAO(db)
	apis.ServeArtistResource(rg, services.NewArtistService(artistDAO))

	// wire up more resource APIs here

	return router
}
