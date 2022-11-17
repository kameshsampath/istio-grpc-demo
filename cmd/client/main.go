package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/kameshsampath/istio-grpc-example/pkg/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	creds "google.golang.org/grpc/credentials/xds"
	_ "google.golang.org/grpc/xds"
)

func main() {
	var serviceName, serviceNamespace, servicePort string

	flag.StringVar(&serviceName, "name", "", "name of the service e.g. greeter")
	flag.StringVar(&serviceNamespace, "namespace", "", "namespace of the service e.g. demos")
	flag.StringVar(&servicePort, "port", "", "service port e.g. 9090")

	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// using xDS for security configuration via discovery
	creds, err := creds.NewClientCredentials(
		creds.ClientOptions{
			FallbackCreds: insecure.NewCredentials(),
		})

	if err != nil {
		log.Fatal(err)
	}

	con, err := grpc.DialContext(ctx,
		fmt.Sprintf("xds:///%s.%s.svc.cluster.local:%s", serviceName, serviceNamespace, servicePort),
		grpc.WithTransportCredentials(creds),
	)

	if err != nil {
		log.Fatal(err)
	}

	client := greeter.NewGreeterClient(con)

	for {
		message, err := client.Greet(context.Background(), &greeter.GreetRequest{
			Name: "guruji",
		})
		if err != nil {
			log.Printf("Call failed with error: %v", err)
		}
		if err == nil {
			log.Printf("%s-%s", message.Message, message.Version)
		}
		time.Sleep(5 * time.Second)
	}
}
