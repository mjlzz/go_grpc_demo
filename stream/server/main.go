/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	pb "streamgrpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloC(stream pb.Greeter_SayHelloCServer) error {
	var result string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "Hello " + result})
		}
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("Received stream request: %v", req.GetName())
		result += req.GetName()
	}
	return nil
}

func (s *server) SayHelloS(req *pb.HelloRequest, stream pb.Greeter_SayHelloSServer) error {
	msg := req.GetName()
	for _, byt := range msg {
		rsp := pb.HelloReply{Message: "Hello " + string(byt)}
		err := stream.Send(&rsp)
		if err != nil {
			log.Fatalf("failed to send stream rsp: %v", err)
			return err
		}
		log.Printf("Send stream rsp: %v", rsp.GetMessage())
	}
	return nil
}

func (s *server) SayHelloB(stream pb.Greeter_SayHelloBServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("Received stream request: %v", req.GetName())

		if err2 := stream.Send(&pb.HelloReply{Message: "Hello " + req.GetName() + req.GetName()}); err2 != nil {
			log.Fatalf("failed to reply: %v", err2)
		}
	}
	return nil
}

func main() {
	fmt.Println("start server...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
