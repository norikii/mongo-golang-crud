package main

import (
	"fmt"
	"github.com/tatrasoft/mongo-golang-crud/database"
	"github.com/tatrasoft/mongo-golang-crud/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

var MongoClient *mongo.Client

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50051...")

	// Start our listener, 50051 is the default gRPC port
	listener, err := net.Listen("tcp", ":50051")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)
	// Create BlogService type
	srv := handlers.BlogServiceServer{}
	// Register the service with the server
	//blogpb.RegisterBlogServiceServer(s, srv)

	MongoClient, err = database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	// start the server in child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	fmt.Println("Server successfully started on port :50051")

	// stop server using shutdown hook
	// crate a channel to receive OS signals
	c := make(chan os.Signal)

	// relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// block main routine until a signal is received
	// as long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// server stops
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	err = database.CloseConnection(MongoClient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}

func GetServer() {

}

