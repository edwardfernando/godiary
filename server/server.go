package server

import (
	"fmt"
	"net/http"

	"github.com/edwardfernando/godiary/config"
	"github.com/edwardfernando/godiary/entry/handler"
	"github.com/edwardfernando/godiary/entry/repository"
	"github.com/edwardfernando/godiary/entry/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq" // load postgres
)

const (
	serverLabel = "[godiary]"
)

// Server struct type
type Server struct {
	Config      *config.ApplicationConfig
	ServerReady chan bool
}

// StartEchoServer To start echo HTTP server
func (s *Server) StartEchoServer() {
	logrus.Infof("Start Echo Server")

	dbString := s.Config.AppDBConnectionURL()
	appDB, err := gorm.Open("postgres", dbString)
	if err != nil {
		logrus.Errorf("failed to connect to postgres db: %s", err.Error())
	}

	appDB.DB().SetMaxIdleConns(s.Config.AppDBMaxIdleConn)
	appDB.DB().SetMaxOpenConns(s.Config.AppDBMaxOpenConn)
	appDB.DB().SetConnMaxLifetime(s.Config.AppDBConnMaxLifetimeDuration)

	defer appDB.Close()

	echoServer := echo.New()

	echoServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	go func() {
		if err := echoServer.Start(fmt.Sprintf(":%d", s.Config.AppPort)); err != nil {
			logrus.Errorf("%s %s", serverLabel, err.Error())
			logrus.Infof("%s shutting down the server", serverLabel)
		}
	}()

	entryRepository := repository.NewEntryRepository(appDB)
	entryUsecase := usecase.NewEntryUsecase(entryRepository)
	entryHandler := handler.NewEntryHTTPHandler(entryUsecase)

	echoServer.POST("/entries", entryHandler.PostEntry)

	logrus.Infof("%s Server started", serverLabel)

	if s.ServerReady != nil {
		s.ServerReady <- true
	}
}
