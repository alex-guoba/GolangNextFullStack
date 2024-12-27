package ratelimit

import (
	"log"
	"net/http"
	"runtime"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/system"

	"github.com/gin-gonic/gin"
)

const (
	DftCPULoadThreshPerCore = 2
)

func initConfig(cfg string, loadThresh float64, loadStrategy int) error {
	// config file format: https://github.com/alibaba/sentinel-golang/wiki/%E5%90%AF%E5%8A%A8%E9%85%8D%E7%BD%AE
	if len(cfg) == 0 {
		if err := sentinel.InitDefault(); err != nil {
			return err
		}
	} else {
		if err := sentinel.InitWithConfigFile(cfg); err != nil {
			return err
		}
	}

	// init rules. see https://github.com/alibaba/sentinel-golang/wiki/%E6%B5%81%E9%87%8F%E6%8E%A7%E5%88%B6
	// we just used system rules here by default.
	// how system adaptive rules works: https://sentinelguard.io/zh-cn/docs/system-adaptive-protection.html
	var loadTrigger = loadThresh
	if loadTrigger <= 0 {
		loadTrigger = float64(DftCPULoadThreshPerCore * runtime.NumCPU())
	}
	if loadStrategy < int(system.NoAdaptive) || loadStrategy > int(system.BBR) {
		loadStrategy = int(system.NoAdaptive)
	}
	_, err := system.LoadRules([]*system.Rule{
		{
			ID:           "DefaultSystemRule",
			MetricType:   system.Load,
			TriggerCount: loadTrigger,
			Strategy:     system.AdaptiveStrategy(loadStrategy),
		},
	})
	if err != nil {
		return err
	}

	log.Printf("Init sentinel success, load trigger %v, load strategy %v", loadTrigger, loadStrategy)
	return nil
}

// See: https://pkg.go.dev/github.com/alibaba/sentinel-golang/pkg/adapters/gin
func New(cfg string, loadThresh float64, loadStrategy int) gin.HandlerFunc {
	// init default
	if err := initConfig(cfg, loadThresh, loadStrategy); err != nil {
		return nil
	}

	return func(c *gin.Context) {
		resourceName := c.Request.Method + ":" + c.FullPath()
		// resourceName := c.GetHeader("X-Real-IP")

		entry, err := sentinel.Entry(
			resourceName,
			sentinel.WithResourceType(base.ResTypeAPIGateway),
			sentinel.WithTrafficType(base.Inbound),
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, map[string]any{
				"err":  "too many request; the quota used up",
				"code": 10222,
			})
			return
		}

		defer entry.Exit()
		c.Next()
	}
}
