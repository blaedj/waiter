
to re-gen the proto gen'd files:
```
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. \ 
  ./rpc/<service_name>/service.proto
```
