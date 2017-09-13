package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"example.com/api"
	"github.com/tnarg/rules_gogo_proto_examples/cmd/myserviced/insecure"
)

//
// Route matcher for gRPC traffic
//
func grpcRouteMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	return r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc")
}

//
// Instantiate a new GRPC Server
//
func newGRPCHandler() http.Handler {
	grpcServer := grpc.NewServer(
		grpc.Creds(credentials.NewClientTLSFromCert(insecure.CertPool, insecure.Addr)),
	)
	api.RegisterMyServiceServer(grpcServer, newMyServiceImpl())
	return grpcServer
}

//
// Instantiate a new REST Gateway HTTP Handler
//
func newRESTGatewayHandler() http.Handler {
	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: insecure.Addr,
		RootCAs:    insecure.CertPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	ctx := context.Background()
	restHandler := runtime.NewServeMux()
	err := api.RegisterMyServiceHandlerFromEndpoint(ctx, restHandler, insecure.Addr, dopts)
	if err != nil {
		panic(fmt.Errorf("serve: %v\n", err))
	}

	return restHandler
}

//
// Instantiate a new HTTP Handler to serve OpenAPI docs
//
func newDocsHandler() http.Handler {
	docsFS := &assetfs.AssetFS{
		Asset:     api.Asset,
		AssetDir:  api.AssetDir,
		AssetInfo: api.AssetInfo,
	}

	return http.StripPrefix("/doc/", http.FileServer(docsFS))
}

func main() {
	router := mux.NewRouter()
	router.MatcherFunc(grpcRouteMatcher).Handler(newGRPCHandler())
	router.PathPrefix("/api/").Handler(newRESTGatewayHandler())
	router.PathPrefix("/doc/").Handler(newDocsHandler())

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", insecure.Port))
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    insecure.Addr,
		Handler: router,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*insecure.KeyPair},
			NextProtos:   []string{"h2"},
		},
	}

	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
