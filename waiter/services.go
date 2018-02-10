package waiterservice

// to re-gen the proto gen'd files:
// protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/service.proto

import (
	"context"
	pb "github.com/blaedj/waiter/rpc"
	twirp "github.com/twitchtv/twirp"
)

type Server struct{}

func (s *Server) CheckIn(ctx context.Context, log *pb.CheckInLog) (ack *pb.CheckInAck, err error) {
	if len(log.Message) < 1 {
		return nil, twirp.InvalidArgumentError("message", "Give me a message")
	}
	return &pb.CheckInAck{}, nil
}
