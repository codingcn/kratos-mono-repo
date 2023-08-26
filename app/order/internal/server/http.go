package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-mono-repo/app/order/internal/conf"
	"kratos-mono-repo/app/order/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, order *service.OrderService, logger log.Logger) *http.Server {
	return nil
}
