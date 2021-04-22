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

func (s *server) Upload(ctx context.Context, in *pb.UploadReq) (*pb.UploadRes, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.UploadRes{Res: "upload " + in.GetName() + " success"}, nil
}

type server2 struct {
}

func (s *server2) Play(ctx context.Context, in *pb.PlayReq) (*pb.PlayRes, error) {
	log.Printf("Received: %v", in.GetAddress())
	return &pb.PlayRes{Code: 200}, nil
}

var gServer = grpc.NewServer()

func StartServer(addr string) {
	log.Printf("start grpc servers...")

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("net listen err: %v", err)
	}

	uploadSvcServer := &server{}
	pb.RegisterUploadSvcServer(gServer, uploadSvcServer)

	playSvcServer := &server2{}
	pb.RegisterPlaySvcServer(gServer, playSvcServer)

	if err = gServer.Serve(conn); err != nil {
		log.Fatalf("grpc server err:%v", err)
	}
}

func StopServer() {
	gServer.GracefulStop()
	log.Println("grpc server stop")
}
