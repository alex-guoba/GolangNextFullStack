package server

import (
	"context"
	"net/http"

	"github.com/alex-guoba/gin-clean-template/internal/middleware"
	"github.com/alex-guoba/gin-clean-template/pkg/setting"
	"github.com/alex-guoba/gin-clean-template/server/api"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//	@title			Gin-Clean-Template
//	@version		1.0
//	@description	Clean Architecture template for Golang Gin services
//	@contact.name	gelco
//	@contact.url	https://github.com/alex-guoba/gin-clean-template
//	@license.name	MIT License
//	@license.url	https://github.com/alex-guoba/gin-clean-template/blob/main/LICENSE
//	@host			localhost:8080
//	@BasePath		/
//	@schemes		http https

type Server struct {
	Router *gin.Engine
	Svr    *http.Server
	Config *setting.Configuration
	DB     *gorm.DB
}

func NewServer(cfg *setting.Configuration, db *gorm.DB) *Server {
	r := gin.New()

	middleware.UseDefault(r, cfg)

	api.SetRouters(r, cfg, db)

	srv := &http.Server{
		Addr:           ":" + cfg.Server.HTTPPort,
		Handler:        r,
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		Config: cfg,
		Svr:    srv,
		Router: r,
		DB:     db,
	}
}

func (s *Server) Start() error {
	// Timeout: https://adam-p.ca/blog/2022/01/golang-http-server-timeouts/
	go func() {
		log.Info("Starting HTTP Server at :", s.Config.Server.HTTPPort)
		if err := s.Svr.ListenAndServe(); err != http.ErrServerClosed {
			log.Error("HTTP server expcetpion. ", err)
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	// TODO: add code
	return s.Svr.Shutdown(ctx)
}
