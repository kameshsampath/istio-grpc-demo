package main

import (
	"log"
	"net"

	"github.com/kameshsampath/istio-grpc-example/pkg/greeter"
	"github.com/kameshsampath/istio-grpc-example/pkg/impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	creds "google.golang.org/grpc/credentials/xds"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/xds"
)

func main() {

	// using xDS for security configuration via discovery
	serverCreds, err := creds.NewServerCredentials(creds.ServerOptions{
		FallbackCreds: insecure.NewCredentials(),
	})

	if err != nil {
		log.Fatal(err)
	}

	s := xds.NewGRPCServer(grpc.Creds(serverCreds))
	reflection.Register(s)
	greeter.RegisterGreeterServer(s, &impl.SimpleGreeterServer{})

	// Listener where there service listens
	lis, err := net.Listen("tcp", ":9090")
	log.Printf("Server Listening, %s", lis.Addr())
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
