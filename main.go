package main

import (
	"fmt"
	"github.com/dedidot/grpc-book-server-streaming/proto"
	"github.com/dedidot/grpc-book-server-streaming/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Service starting ...... ")

	lis, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		fmt.Println("Failed to listen.... ", err)
	}

	s := grpc.NewServer()
	proto.RegisterBookServiceServer(s, &server.Server{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting server ..... ")
		if err := s.Serve(lis); err != nil {
			fmt.Println("Failed to starting server ... ", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stop the server .....")
	s.Stop()
	fmt.Println("Stopping listener ....")
	lis.Close()
	fmt.Println("End program ......")
}
