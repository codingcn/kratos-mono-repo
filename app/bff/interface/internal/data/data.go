package data

import (
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	userv1 "kratos-mono-repo/api/user/service/v1"
	"kratos-mono-repo/app/bff/interface/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewGormClient,
	NewRedisCmd,
	NewUserUsecase,
	NewUserServiceClient,
)

// Data .
type Data struct {
	// TODO wrapped database client
	db       *gorm.DB
	redisCli redis.Cmdable
	uc       userv1.UserClient
}

// NewData .
func NewData(c *conf.Data,
	db *gorm.DB,
	redisCli redis.Cmdable,
	uc userv1.UserClient,
	logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:       db,
		redisCli: redisCli,
		uc:       uc,
	}, cleanup, nil
}
func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Addr
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}

func NewRegistrar(conf *conf.Registry) (registry.Registrar, func(), error) {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Addr
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(true), consul.WithHealthCheckInterval(3))
	cleanFunc := func() {
	}
	return r, cleanFunc, nil
}
func NewGormClient(conf *conf.Data, logger log.Logger) (*gorm.DB, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "server-server/data/gorm"))
	client, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		ConnPool: nil,
	})
	client = client.Debug()
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	cleanFunc := func() {}
	return client, cleanFunc, nil
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "server-server/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		Password:     conf.Redis.Password,
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}
func NewUserServiceClient(r registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///kratos.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}
