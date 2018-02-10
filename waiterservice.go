package waiter

import (
	"context"
	pb "github.com/blaedj/waiter/proto"
)

type WaiterServer struct{}

// Server implements the waiter interface
func (s *WaiterServer) CheckIn(ctx context.Context, log *pb.CheckInLog) (ack *pb.CheckInAck, err error) {
	if len(log.Message) < 1 {
		return nil, twirp.InvalidArgumentError("message", "Give me a message")
	}
	return &pb.CheckInAck{}, nil
}
