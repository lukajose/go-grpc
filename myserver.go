
package main

import (
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
)
type server struct{}
func main() {
	fmt.Println("Hello golang server")
	lis, err := net.Listen("tcp","0.0.0.0:50051") 
	if err != nil {
		log.Fatalf("Error happened: %v",err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGuestServiceServer(s * grpc.Server,)

}
