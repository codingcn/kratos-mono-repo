package main

import (
	"flag"
	"fmt"
	logger "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"kratos-mono-repo/app/bff/interface/internal/conf"
	"kratos-mono-repo/pkg/zlog"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}
func newApp(conf *conf.Server, logger log.Logger, gs *grpc.Server, hs *http.Server, rr registry.Registrar) *kratos.App {

	return kratos.New(
		//kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(rr), // Notice 添加服务注册
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	fmt.Println("conf file", flagconf)
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	id = id + "-" + bc.Server.Name
	Name = bc.Server.Name
	Version = bc.Server.Version
	//

	zapLogger := logger.NewLogger(zlog.NewLogger(false, zap.DebugLevel, "./log/"+bc.Server.Name+"/server.log"))
	logger := log.With(zapLogger,
		//"ts", log.DefaultTimestamp,
		//"caller", log.Caller(4),
		"server.id", id,
		"server.name", Name,
		"server.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(bc.Trace.Endpoint)))
	if err != nil {
		panic(err)
	}

	tp := tracesdk.NewTracerProvider(
		// 将基于父span的采样率设置为100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// 始终确保再生成中批量处理
		tracesdk.WithBatcher(exp),
		// 在资源中记录有关此应用程序的信息
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("exporter", "jaeger"),
		)),
	)
	otel.SetTracerProvider(tp)
	app, cleanup, err := wireApp(bc.Server, bc.Registry, bc.Data, bc.Auth, logger, tp)

	log.With(log.NewStdLogger(os.Stdout),
		"server.id", app.ID(),
		"server.name", app.Name(),
		"server.version", app.Version(),
	)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
