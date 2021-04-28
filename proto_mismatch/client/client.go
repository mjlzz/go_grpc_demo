package main

import (
	"context"
	"log"
	//"os"
	"time"

	"google.golang.org/grpc"
	pb "grpcdemo/proto2"
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

	// rpc call pass, but parameters mismatch
	play(conn)

	// could not play: rpc error: code = Unimplemented desc = unknown method Stop2 for service api.PlaySvc
	// stop(conn)
}

func play(conn *grpc.ClientConn) {
	c := pb.NewPlaySvcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Play(ctx, &pb.PlayReq2{A: "http://pic.baidu.com", Nums: 3, Cat: "mm"})
	if err != nil {
		log.Fatalf("could not play: %v", err)
	}
	log.Printf("Upload: %d", r.GetCode())

}

func stop(conn *grpc.ClientConn) {
	c := pb.NewPlaySvcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Stop2(ctx, &pb.StopReq{Id: 123})
	if err != nil {
		log.Fatalf("could not play: %v", err)
	}
	log.Printf("Upload: %v", r)
}
