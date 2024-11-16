package server

import (
	"Bangseungjae/cockroach/cockroach/handlers"
	"Bangseungjae/cockroach/cockroach/repositories"
	"Bangseungjae/cockroach/cockroach/usecases"
	"Bangseungjae/cockroach/config"
	"Bangseungjae/cockroach/database"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover()) // Recover: 패닉 발생 시 서버가 크래시되지 않고 복구할 수 있도록 도와줍니다.
	s.app.Use(middleware.Logger())  // Logger: HTTP 요청과 응답을 로깅합니다.

	s.app.GET("/v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	s.initializeCockroachHttpHandler()

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeCockroachHttpHandler() {
	// Initialize all layers
	cockroachPostgersRepository := repositories.NewCockroachPostgresRepository(s.db)
	cockroachFCMMessaging := repositories.NewCockroachFCMMessaging()

	cockroachUsecase := usecases.NewCockroachUsecaseImpl(
		cockroachPostgersRepository,
		cockroachFCMMessaging,
	)

	cockroachHttpHandler := handlers.NewCockroachHttpHandler(cockroachUsecase)

	// Routers
	cockroachRouters := s.app.Group("/v1/cockroach")
	cockroachRouters.POST("", cockroachHttpHandler.DetectCockroach)
}
