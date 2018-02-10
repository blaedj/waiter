package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/blaedj/waiter/proto"
	"net/http"
	"os"
)

var flURI = flag.String("uri", "http://localhost:8765",
	`Supply a string to use as the server uri for the client to connect to
    Example:
       --uri http://localhost:8765`)

var flULID = flag.String("ulid", "abc123", `supply a valid ULID identifier`)

func main() {
	flag.Parse()
	client := pb.NewWaiterProtobufClient(*flURI, &http.Client{})
	ack, err := client.CheckIn(context.Background(), &pb.CheckInLog{
		Ulid:    *flULID,
		Status:  pb.DinerStatus_OK,
		Message: "this is a message"})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Waiter acked %v\n", ack)
	//client := pb.NewWaiterProtobufClient("http://localhost:8990", &http.Client{})
}
