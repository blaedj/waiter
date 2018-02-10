package waiter

import (
	"context"
	pb "github.com/blaedj/waiter/rpc/checkin"
	"github.com/twitchtv/twirp"
	logger "log"
)

type WaiterServer struct{}

// Server implements the waiter interface
func (s WaiterServer) CheckIn(ctx context.Context, log *pb.CheckInLog) (ack *pb.CheckInAck, err error) {
	if len(log.Message) < 1 {
		return nil, twirp.InvalidArgumentError("message", "Give me a message")
	}
	logger.Printf("recieved log from diner: %v, with message: '%v'. Status: %v",
		log.Ulid, log.Message, log.Status)
	return &pb.CheckInAck{Status: 0}, nil
}
