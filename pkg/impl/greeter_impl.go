package impl

import (
	context "context"
	"fmt"
	"os"

	"github.com/kameshsampath/istio-grpc-example/pkg/greeter"
)

type SimpleGreeterServer struct {
	greeter.UnimplementedGreeterServer
}

// Greet implements greeter.GreeterServer
func (*SimpleGreeterServer) Greet(ctx context.Context, req *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	if req.Name == "guruji" {
		return &greeter.GreetResponse{
			Message: "Jai Guru",
			Version: os.Getenv("VERSION"),
		}, nil
	}
	return &greeter.GreetResponse{
		Message: fmt.Sprintf("Hello, %s", req.Name),
		Version: os.Getenv("VERSION"),
	}, nil
}

var _ greeter.GreeterServer = (*SimpleGreeterServer)(nil)
