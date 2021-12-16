package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "streamgrpc/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	streamclient, err := c.SayHelloC(context.Background())
	for i := 0; i < 5; i++ {
		streamclient.Send(&pb.HelloRequest{Name: fmt.Sprint(i)})
	}
	r2, err := streamclient.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r2.GetMessage())

	streamRes, err := c.SayHelloS(ctx, &pb.HelloRequest{Name: "56789"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Print("==>start get stream rsp")
	for {
		rsp, err := streamRes.Recv()
		if err == io.EOF {
			log.Print("<==end get stream rsp")
			break
		}
		if err != nil {
			log.Fatalf("greet err: %v", err)
		}
		log.Printf("get stream rsp: %s", rsp.GetMessage())
	}

	streamB, err := c.SayHelloB(context.Background())
	go func() {
		log.Print("==>start get stream rsp")
		for {
			rsp, err := streamB.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("greet err: %v", err)
			}
			log.Printf("get stream rsp: %s", rsp.GetMessage())
		}
		log.Print("<==end get stream rsp")
	}()
	for i := 0; i < 5; i++ {
		streamB.Send(&pb.HelloRequest{Name: fmt.Sprint(i)})
	}
	streamB.CloseSend()
	time.Sleep(time.Second)
}
