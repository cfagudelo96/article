package gateway

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func Run(dialAddr string) error {
	ctx := context.Background()
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	conn, err := grpc.DialContext(
		ctx,
		dialAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	gwmux := runtime.NewServeMux()
	restaurantv1.RegisterRestaurantServiceHandler(ctx, gwmux, conn)
	port := os.Getenv("PORT")
	if port == "" {
		port = "11000"
	}
	gatewayAddr := "0.0.0.0:" + port
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gwmux.ServeHTTP(w, r)
		}),
	}
	return gwServer.ListenAndServe()
}
