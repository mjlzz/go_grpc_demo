package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpcdemo/proto"
)

const (
	addr = "127.0.0.1:8001"
)

func main() {
	StartServer(addr)
}

type server struct {
}

func (s *server) Play(ctx context.Context, in *pb.PlayReq) (*pb.PlayRes, error) {
	log.Print("Received play: ", in.GetAddress(), "  ", in.GetNums())
	log.Print("Received play: ", in)
	return &pb.PlayRes{Code: 200}, nil
}

func (s *server) Stop(ctx context.Context, in *pb.StopReq) (*pb.StopRes, error) {
	return &pb.StopRes{Status: "on"}, nil
}

var gServer = grpc.NewServer()

func StartServer(addr string) {
	log.Printf("start grpc servers...")

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("net listen err: %v", err)
	}

	playSvcServer := &server{}
	pb.RegisterPlaySvcServer(gServer, playSvcServer)

	if err = gServer.Serve(conn); err != nil {
		log.Fatalf("grpc server err:%v", err)
	}
}

func StopServer() {
	gServer.GracefulStop()
	log.Println("grpc server stop")
}
