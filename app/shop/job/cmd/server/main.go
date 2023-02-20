package main

import (
	"flag"

	"github.com/lalifeier/vvgo-mall/app/shop/job/internal/conf"
	"github.com/lalifeier/vvgo-mall/pkg/bootstrap"
	"github.com/tx7do/kratos-transport/transport/kafka"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Service = bootstrap.NewServiceInfo(
		"vvgo-mall.shop.job",
		"1.0.0",
		"",
	)

	Flags = bootstrap.NewCommandFlags()
)

func init() {
	Flags.Init()
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar, ks *kafka.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			ks,
		),
		kratos.Registrar(rr),
	)
}

func loadConfig() (*conf.Bootstrap, *conf.Registry) {
	c := bootstrap.NewConfigProvider(Flags.ConfigType, Flags.ConfigHost, Flags.Conf, Service.Name)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	return &bc, &rc
}

func main() {
	flag.Parse()

	logger := bootstrap.NewLoggerProvider(&Service)

	bc, rc := loadConfig()
	if bc == nil || rc == nil {
		panic("load config failed")
	}

	err := bootstrap.NewTracerProvider(bc.Trace.Batcher, bc.Trace.Endpoint, Flags.Env, &Service)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, rc, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}