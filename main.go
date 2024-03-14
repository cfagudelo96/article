package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"

	"github.com/cfagudelo96/article/core/restaurant"
	restaurantstore "github.com/cfagudelo96/article/core/restaurant/store"
	"github.com/cfagudelo96/article/gateway"
	restauranthdlr "github.com/cfagudelo96/article/handlers/restaurant"
	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type config struct {
	Address string `default:"127.0.0.1:10000"`
	DB      struct {
		User         string `required:"true" split_words:"true"`
		Password     string `required:"true" split_words:"true"`
		Address      string `required:"true" split_words:"true"`
		Name         string `required:"true" split_words:"true"`
		MaxIdleConns int    `default:"2" split_words:"true"`
		MaxOpenConns int    `split_words:"true"`
		DisableTLS   bool   `default:"true" split_words:"true"`
	}
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err.Error())
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	var conf config
	if err := envconfig.Process("", &conf); err != nil {
		return fmt.Errorf("processing environment: %w", err)
	}

	dbConn, err := buildDBConnection(ctx, conf)
	if err != nil {
		return fmt.Errorf("building DB connection: %w", err)
	}
	defer dbConn.Close(ctx)

	lis, err := net.Listen("tcp", conf.Address)
	if err != nil {
		return fmt.Errorf("listening to %s: %w", conf.Address, err)
	}

	s := grpc.NewServer()

	rs := restaurantstore.NewStore(dbConn)
	rc := restaurant.NewCore(rs)
	rh, err := restauranthdlr.NewHandler(rc)
	if err != nil {
		return fmt.Errorf("initializing handler: %w", err)
	}

	restaurantv1.RegisterRestaurantServiceServer(s, rh)
	log.Info("Serving gRPC on https://", conf.Address)

	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalf("serving server: %w", err)
		}
	}()

	if err = gateway.Run("dns:///" + conf.Address); err != nil {
		return fmt.Errorf("running gateway: %w", err)
	}
	return nil
}

func interceptError(
	ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (any, error) {
	resp, err := handler(ctx, req)
	if err == nil {
		return resp, nil
	}
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Internal, "Unexpected error")
		d := &errdetails.DebugInfo{
			Detail: err.Error(),
		}
		s, err = s.WithDetails(d)
		if err != nil {
			return nil, status.New(codes.Internal, "Unexpected error").Err()
		}
		return nil, s.Err()
	}
	return nil, s.Err()
}

func buildDBConnection(ctx context.Context, conf config) (*pgx.Conn, error) {
	dbConf := conf.DB
	sslMode := "require"
	if dbConf.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("search_path", "public")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(dbConf.User, dbConf.Password),
		Host:     dbConf.Address,
		Path:     dbConf.Name,
		RawQuery: q.Encode(),
	}

	dbURL := u.String()
	m, err := migrate.New("file://core/data/migrations", dbURL)
	if err != nil {
		return nil, fmt.Errorf("initializing migrator: %w", err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("applying migrations: %w", err)
	}
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("connecting to DB: %w", err)
	}
	return conn, nil
}
