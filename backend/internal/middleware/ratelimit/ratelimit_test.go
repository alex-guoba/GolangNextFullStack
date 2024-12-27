package ratelimit

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSentinelMiddleware(t *testing.T) {
	type args struct {
		thresh   float64
		strategy int
		method   string
		path     string
		reqPath  string
		handler  func(ctx *gin.Context)
		body     io.Reader
	}
	type want struct {
		code int
	}
	var (
		tests = []struct {
			name string
			args args
			want want
		}{
			{
				name: "pass strategy non-adaptive",
				args: args{
					thresh:   1000,
					strategy: int(system.NoAdaptive),
					method:   http.MethodGet,
					path:     "/ping",
					reqPath:  "/ping",
					handler: func(ctx *gin.Context) {
						ctx.String(http.StatusOK, "ping")
					},
					body: nil,
				},
				want: want{
					code: http.StatusOK,
				},
			},
			{
				name: "pass strategy BBR",
				args: args{
					thresh:   1000,
					strategy: int(system.BBR),
					method:   http.MethodGet,
					path:     "/ping",
					reqPath:  "/ping",
					handler: func(ctx *gin.Context) {
						ctx.String(http.StatusOK, "ping")
					},
					body: nil,
				},
				want: want{
					code: http.StatusOK,
				},
			},
			{
				name: "failed strategy non-adaptivet",
				args: args{
					thresh:   0.0001, // smaller enough
					strategy: int(system.NoAdaptive),
					method:   http.MethodPost,
					path:     "/api/users/:id",
					reqPath:  "/api/users/123",
					handler: func(ctx *gin.Context) {
						ctx.String(http.StatusOK, "ping")
					},
					body: nil,
				},
				want: want{
					code: http.StatusTooManyRequests,
				},
			},
			{
				name: "pass strategy non-adaptivet",
				args: args{
					thresh:   0.0001, // smaller enough
					strategy: int(system.BBR),
					method:   http.MethodPost,
					path:     "/api/users/:id",
					reqPath:  "/api/users/123",
					handler: func(ctx *gin.Context) {
						ctx.String(http.StatusOK, "ping")
					},
					body: nil,
				},
				want: want{
					code: http.StatusOK,
				},
			},
		}
	)

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.Use(New("", tt.args.thresh, tt.args.strategy))
			router.Handle(tt.args.method, tt.args.path, tt.args.handler)
			r := httptest.NewRequest(tt.args.method, tt.args.reqPath, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
		})
	}
}
