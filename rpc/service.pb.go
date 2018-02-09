// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/service.proto

/*
Package waiter is a generated protocol buffer package.

It is generated from these files:
	rpc/service.proto

It has these top-level messages:
	CheckInAck
	CheckInLog
*/
package waiter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// a simple OK or ERROR status
type DinerStatus int32

const (
	DinerStatus_OK    DinerStatus = 0
	DinerStatus_ERROR DinerStatus = 1
)

var DinerStatus_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}
var DinerStatus_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x DinerStatus) String() string {
	return proto.EnumName(DinerStatus_name, int32(x))
}
func (DinerStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CheckInAck struct {
}

func (m *CheckInAck) Reset()                    { *m = CheckInAck{} }
func (m *CheckInAck) String() string            { return proto.CompactTextString(m) }
func (*CheckInAck) ProtoMessage()               {}
func (*CheckInAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// a checking log contains basic information for a checkin.
type CheckInLog struct {
	Ulid    string      `protobuf:"bytes,1,opt,name=ulid" json:"ulid,omitempty"`
	Status  DinerStatus `protobuf:"varint,2,opt,name=status,enum=blaedj.waiter.DinerStatus" json:"status,omitempty"`
	Message string      `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *CheckInLog) Reset()                    { *m = CheckInLog{} }
func (m *CheckInLog) String() string            { return proto.CompactTextString(m) }
func (*CheckInLog) ProtoMessage()               {}
func (*CheckInLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CheckInLog) GetUlid() string {
	if m != nil {
		return m.Ulid
	}
	return ""
}

func (m *CheckInLog) GetStatus() DinerStatus {
	if m != nil {
		return m.Status
	}
	return DinerStatus_OK
}

func (m *CheckInLog) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckInAck)(nil), "blaedj.waiter.CheckInAck")
	proto.RegisterType((*CheckInLog)(nil), "blaedj.waiter.CheckInLog")
	proto.RegisterEnum("blaedj.waiter.DinerStatus", DinerStatus_name, DinerStatus_value)
}

func init() { proto.RegisterFile("rpc/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2a, 0x48, 0xd6,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d,
	0xca, 0x49, 0x4c, 0x4d, 0xc9, 0xd2, 0x2b, 0x4f, 0xcc, 0x2c, 0x49, 0x2d, 0x52, 0xe2, 0xe1, 0xe2,
	0x72, 0xce, 0x48, 0x4d, 0xce, 0xf6, 0xcc, 0x73, 0x4c, 0xce, 0x56, 0xca, 0x83, 0xf3, 0x7c, 0xf2,
	0xd3, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x73, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x21, 0x23, 0x2e, 0xb6, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0x3e, 0x23, 0x29, 0x3d, 0x14, 0xf3, 0xf4, 0x5c, 0x32, 0xf3, 0x52, 0x8b, 0x82, 0xc1, 0x2a,
	0x82, 0xa0, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x98,
	0xc1, 0x46, 0xc1, 0xb8, 0x5a, 0x0a, 0x5c, 0xdc, 0x48, 0x1a, 0x84, 0xd8, 0xb8, 0x98, 0xfc, 0xbd,
	0x05, 0x18, 0x84, 0x38, 0xb9, 0x58, 0x5d, 0x83, 0x82, 0xfc, 0x83, 0x04, 0x18, 0x8d, 0x3c, 0xb9,
	0xd8, 0xc2, 0xc1, 0x26, 0x0b, 0xd9, 0x73, 0xb1, 0x43, 0xdd, 0x26, 0x24, 0x89, 0x66, 0x29, 0xc2,
	0xcd, 0x52, 0x38, 0xa4, 0x1c, 0x93, 0xb3, 0x9d, 0x38, 0xa2, 0xd8, 0x20, 0x82, 0x49, 0x6c, 0xe0,
	0xa0, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x35, 0xde, 0x8c, 0xb8, 0x1f, 0x01, 0x00, 0x00,
}
