package signals

import (
	"context"
	"time"

	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/server"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Shutdown struct {
	// pool                  *redis.Pool
	// tracerProvider        *sdktrace.TracerProvider
	serverShutdownTimeout time.Duration
}

func NewShutdown(serverShutdownTimeout time.Duration) (*Shutdown, error) {
	srv := &Shutdown{
		// logger:                logger,
		serverShutdownTimeout: serverShutdownTimeout,
	}

	return srv, nil
}

func (s *Shutdown) Graceful(stopCh <-chan struct{}, svr *server.Server, engine *gorm.DB) {
	ctx := context.Background()

	// wait for the server to gracefully terminate
	<-stopCh
	ctx, cancel := context.WithTimeout(ctx, s.serverShutdownTimeout)
	defer cancel()

	// all calls to /healthz and /readyz will fail from now on
	// atomic.StoreInt32(healthy, 0)
	// atomic.StoreInt32(ready, 0)

	// close cache pool
	// if s.pool != nil {
	// 	_ = s.pool.Close()
	// }

	log.Info("Shutting down HTTP/HTTPS server. ", s.serverShutdownTimeout)

	// There could be a period where a terminating pod may still receive requests. Implementing a brief wait can mitigate this.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination
	// the readiness check interval must be lower than the timeout
	// if viper.GetString("level") != "debug" {
	// 	time.Sleep(3 * time.Second)
	// }

	// // stop OpenTelemetry tracer provider
	// if s.tracerProvider != nil {
	// 	if err := s.tracerProvider.Shutdown(ctx); err != nil {
	// 		s.logger.Warn("stopping tracer provider", zap.Error(err))
	// 	}
	// }

	// determine if the GRPC was started
	// if grpcServer != nil {
	// 	s.logger.Info("Shutting down GRPC server", zap.Duration("timeout", s.serverShutdownTimeout))
	// 	grpcServer.GracefulStop()
	// }

	// determine if the http server was started
	if err := svr.Shutdown(ctx); err != nil {
		log.Warn("HTTP server graceful shutdown failed", err)
	}

	if engine != nil {
		if err := dao.Close(engine); err != nil {
			log.Warn("gorm close failed", err)
		}
	}

	log.Error("Shutdown complete.")
}
