// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmoscheckersgame/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("cosmoscheckersgame/tx.proto", fileDescriptor_95cac711d7033459) }

var fileDescriptor_95cac711d7033459 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0x2a, 0x4e, 0x4f, 0xcc, 0x4d, 0xd5, 0x2f,
	0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x4d, 0x2c, 0x4b, 0xcc, 0xc9, 0xce, 0x2f,
	0xd3, 0xc3, 0x54, 0x84, 0x45, 0xc8, 0x88, 0x95, 0x8b, 0xd9, 0xb7, 0x38, 0xdd, 0x29, 0xea, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1,
	0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x1c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xa1, 0x46, 0xeb, 0x3b, 0x83, 0xcd, 0x71, 0x86, 0x9a, 0xe3, 0x0e,
	0xb2, 0xbf, 0x42, 0x1f, 0x9b, 0xa3, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x0e, 0x33, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0x87, 0x4c, 0xbd, 0xb7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "avalkov.cosmoscheckersgame.cosmoscheckersgame.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cosmoscheckersgame/tx.proto",
}
