package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	//assetfs "github.com/philips/go-bindata-assetfs"
	//"github.com/philips/grpc-gateway-example/pkg/ui/data/swagger"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	v1armytypespb "github.com/happilymarrieddad/old-world-api3/pb/proto/armytypes"
	authpb "github.com/happilymarrieddad/old-world-api3/pb/proto/auth"
	v1compositiontypespb "github.com/happilymarrieddad/old-world-api3/pb/proto/compositionTypes"
	v1gamespb "github.com/happilymarrieddad/old-world-api3/pb/proto/games"
	v1itemtypespb "github.com/happilymarrieddad/old-world-api3/pb/proto/itemtypes"
	v1optiontypespb "github.com/happilymarrieddad/old-world-api3/pb/proto/optiontypes"
	v1statisticspb "github.com/happilymarrieddad/old-world-api3/pb/proto/statistics"
	v1trooptypespb "github.com/happilymarrieddad/old-world-api3/pb/proto/trooptypes"
	v1unittypespb "github.com/happilymarrieddad/old-world-api3/pb/proto/unittypes"
	v1userarmiespb "github.com/happilymarrieddad/old-world-api3/pb/proto/userarmies"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

// TODO: figure out how to do this
// func serveSwagger(mux *runtime.ServeMux) {
// 	mime.AddExtensionType(".svg", "image/svg+xml")

// 	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
// 	fileServer := http.FileServer(&assetfs.AssetFS{
// 		Asset:    swagger.Asset,
// 		AssetDir: swagger.AssetDir,
// 		Prefix:   "third_party/swagger-ui",
// 	})
// 	prefix := "/swagger-ui/"
// 	mux.Handle("GET", prefix, runtime.HandlerFunc(fileServer),

// 	http.StripPrefix(prefix, fileServer))
// }

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := authpb.RegisterAuthHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	// V1
	if err := v1gamespb.RegisterV1GamesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := v1armytypespb.RegisterV1ArmyTypesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := v1compositiontypespb.RegisterV1CompositionTypesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := v1itemtypespb.RegisterV1ItemTypesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := v1optiontypespb.RegisterV1OptionTypesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := v1statisticspb.RegisterV1StatisticsHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := v1trooptypespb.RegisterV1TroopTypesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := v1unittypespb.RegisterV1UnitTypesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := v1userarmiespb.RegisterV1UserArmiesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := v1userarmiespb.RegisterV1UserArmyUnitsHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	port := int(8081)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      cors.AllowAll().Handler(mux),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	log.Printf("Server running on port %d\n", port)
	return s.ListenAndServe()
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
