package main

import (
	pb "Go-project/proto"
	"Go-project/server"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitTagServer() *grpc.Server {
	wire.Build(grpc.NewServer)
	return &grpc.Server{}
}

func main() {
	s := InitTagServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
