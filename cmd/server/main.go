package main

import (
	waiter "github.com/blaedj/waiter/internal"
	pb "github.com/blaedj/waiter/rpc/checkin"
	"net/http"
)

func main() {
	server := waiter.WaiterServer{} // implements waiter interface
	twirpHandler := pb.NewWaiterServer(server, nil)
	http.ListenAndServe(":8082", twirpHandler)
}
