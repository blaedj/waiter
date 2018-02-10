# generate the protos for this project
gen:
	protoc --proto_path=$(GOPATH)/src:. --twirp_out=. --go_out=. ./rpc/checkin/service.proto
