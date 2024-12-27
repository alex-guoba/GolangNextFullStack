package middleware

import (
	"github.com/alex-guoba/gin-clean-template/internal/middleware/ratelimit"
	"github.com/alex-guoba/gin-clean-template/pkg/setting"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	gswag "github.com/swaggo/gin-swagger"  // gin-swagger middleware
)

// UseDefault set the default middleware
func UseDefault(r *gin.Engine, cfg *setting.Configuration) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// rate limit middleware
	if cfg.Ratelimit.Enable {
		limiter := ratelimit.New(cfg.Ratelimit.ConfigFile,
			cfg.Ratelimit.CPULoadThresh, cfg.Ratelimit.CPULoadStrategy)
		if limiter == nil {
			log.Error("init rate limit middleware failed, ignored")
		} else {
			r.Use(limiter)
		}
	}
	// r.Static("/static", "./web/dist")

	// swagger middleware
	r.GET("/swagger/*any", gswag.WrapHandler(swaggerfiles.Handler))
}
