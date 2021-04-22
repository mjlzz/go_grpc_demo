package main

import (
	"context"
	"log"
	//"os"
	"time"

	"google.golang.org/grpc"
	pb "grpcdemo/proto"
)

const (
	address = "localhost:8001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	play(conn)

	upload(conn)
}

func play(conn *grpc.ClientConn) {
	c := pb.NewPlaySvcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Play(ctx, &pb.PlayReq{Address: "http://pic.baidu.com", Nums: 3})
	if err != nil {
		log.Fatalf("could not play: %v", err)
	}
	log.Printf("Upload: %d", r.GetCode())
}

func upload(conn *grpc.ClientConn) {
	c := pb.NewUploadSvcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Upload(ctx, &pb.UploadReq{Name: "Super Mario"})
	if err != nil {
		log.Fatalf("could not upload: %v", err)
	}
	log.Printf("Upload: %s", r.GetRes())
}
