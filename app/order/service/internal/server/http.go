package server

import (
	"kratos-mono-repo/app/order/service/internal/conf"
	"kratos-mono-repo/app/order/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, order *service.OrderService, logger log.Logger) *http.Server {
	return nil
}
