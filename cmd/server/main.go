package main

import (
	"github.com/blaedj/waiter"
	pb "github.com/blaedj/waiter/proto"
	"net/http"
)

func main() {
	server := waiter.WaiterServer{} // implements waiter interface
	twirpHandler := pb.NewWaiterServer(server, nil)
	http.ListenAndServe(":8082", twirpHandler)
}
