package client

import (
	"io"
	"log"
	"os"
	"time"

	_ "github.com/asim/go-micro/plugins/client/grpc/v4"
	_ "github.com/asim/go-micro/plugins/registry/consul/v4"
	_ "github.com/asim/go-micro/plugins/registry/kubernetes/v4"
	kb "github.com/asim/go-micro/plugins/registry/kubernetes/v4"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go-micro.dev/v4/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func isUseKubernetes() bool {
	return os.Getenv("DISCOVERY_DRIVER") == "kubernetes"
}

func isUseConsul() bool {
	return os.Getenv("DISCOVERY_DRIVER") == "consul"
}

func GetClientOptions() []client.Option {
	options := []client.Option{
		client.RequestTimeout(300 * time.Second),
	}

	if isUseKubernetes() {
		options = append(options, client.Registry(kb.NewRegistry()))
	}

	return options
}

func grpcClientOptions() []grpc.DialOption {
	tracer, closer, err := openTracingConfig()
	if err != nil {
		log.Fatalf("an error accoured when create tracer: ", err)
	}
	defer closer.Close()

	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(tracer),
			),
		),
	}
}

func openTracingConfig() (opentracing.Tracer, io.Closer, error) {
	// tracing using opentracing and jaeger
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, nil, err
	}

	return cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
}

func NewServiceAuthClient() (*grpc.ClientConn, error) {
	return grpc.Dial(os.Getenv("AUTH_SERVICE_HOST"), GetClientOptions())
}

func NewServiceUserClient() (*grpc.ClientConn, error) {
	return grpc.Dial(os.Getenv("USER_SERVICE_HOST"), GetClientOptions())
}

func NewServiceRoleClient() (*grpc.ClientConn, error) {
	return grpc.Dial(os.Getenv("ROLE_SERVICE_HOST"), GetClientOptions())
}

func NewServiceMasterClient() (*grpc.ClientConn, error) {
	return grpc.Dial(os.Getenv("MASTER_SERVICE_HOST"), GetClientOptions())
}
