package waiter

// to re-gen the proto gen'd files:
// protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./proto/service.proto

import (
	"flag"
	pb "github.com/blaedj/waiter/proto"
	"net/http"
)

var flURI = flag.String("uri", "localhost:8765",
	`Supply a string to use as the server uri.
    Example:
       ./server --uri localhost:3456
`)

func main() {
	//flag.Parse()
	//log.Printf("using uri: %v", *flURI)
	//http.HandleFunc("/checkin", checkinHandler)
	//	http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	//log.Fatal(http.ListenAndServe(*flURI, nil))

	server := WaiterServer{} // implements waiter interface
	twirpHandler := pb.NewWaiterServer(server, nil)
	http.ListenAndServe(":8082", twirpHandler)
}
