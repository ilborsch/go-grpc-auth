package suite

import (
	"context"
	ssov1 "github.com/GolangLessons/protos/gen/go/sso"
	"go-grpc-auth/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"strconv"
	"testing"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath("../config/local_test.yaml")

	ctx, timeoutCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	defer t.Cleanup(func() {
		t.Helper()
		timeoutCtx()
	})

	cc, err := grpc.DialContext(context.Background(), grpcAddress(cfg), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc connection failed: %s", err)
	}
	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: ssov1.NewAuthClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort("127.0.0.1", strconv.Itoa(cfg.GRPC.Port))
}
