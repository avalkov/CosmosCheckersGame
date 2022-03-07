// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmoscheckersgame/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// this line is used by starport scaffolding # 3
type QueryGetNextGameRequest struct {
}

func (m *QueryGetNextGameRequest) Reset()         { *m = QueryGetNextGameRequest{} }
func (m *QueryGetNextGameRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetNextGameRequest) ProtoMessage()    {}
func (*QueryGetNextGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0acbb2f8ad96b12c, []int{0}
}
func (m *QueryGetNextGameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetNextGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetNextGameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetNextGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetNextGameRequest.Merge(m, src)
}
func (m *QueryGetNextGameRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetNextGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetNextGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetNextGameRequest proto.InternalMessageInfo

type QueryGetNextGameResponse struct {
	NextGame *NextGame `protobuf:"bytes,1,opt,name=NextGame,proto3" json:"NextGame,omitempty"`
}

func (m *QueryGetNextGameResponse) Reset()         { *m = QueryGetNextGameResponse{} }
func (m *QueryGetNextGameResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetNextGameResponse) ProtoMessage()    {}
func (*QueryGetNextGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0acbb2f8ad96b12c, []int{1}
}
func (m *QueryGetNextGameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetNextGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetNextGameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetNextGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetNextGameResponse.Merge(m, src)
}
func (m *QueryGetNextGameResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetNextGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetNextGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetNextGameResponse proto.InternalMessageInfo

func (m *QueryGetNextGameResponse) GetNextGame() *NextGame {
	if m != nil {
		return m.NextGame
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetNextGameRequest)(nil), "avalkov.cosmoscheckersgame.cosmoscheckersgame.QueryGetNextGameRequest")
	proto.RegisterType((*QueryGetNextGameResponse)(nil), "avalkov.cosmoscheckersgame.cosmoscheckersgame.QueryGetNextGameResponse")
}

func init() { proto.RegisterFile("cosmoscheckersgame/query.proto", fileDescriptor_0acbb2f8ad96b12c) }

var fileDescriptor_0acbb2f8ad96b12c = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0x2a, 0x4e, 0x4f, 0xcc, 0x4d, 0xd5, 0x2f,
	0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x4d, 0x2c, 0x4b, 0xcc,
	0xc9, 0xce, 0x2f, 0xd3, 0xc3, 0x54, 0x87, 0x45, 0x48, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27,
	0x55, 0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf,
	0x18, 0x62, 0x98, 0x94, 0x16, 0x44, 0x87, 0x7e, 0x52, 0x62, 0x31, 0xd4, 0x16, 0xfd, 0x32, 0xc3,
	0xa4, 0xd4, 0x92, 0x44, 0x43, 0xfd, 0x82, 0xc4, 0xf4, 0xcc, 0x3c, 0xb0, 0x62, 0xa8, 0x5a, 0x25,
	0x2c, 0x0e, 0xcb, 0x4b, 0xad, 0x28, 0x89, 0x07, 0x5b, 0x0d, 0x56, 0xa3, 0x24, 0xc9, 0x25, 0x1e,
	0x08, 0x32, 0xc5, 0x3d, 0xb5, 0xc4, 0x2f, 0xb5, 0xa2, 0xc4, 0x3d, 0x31, 0x37, 0x35, 0x28, 0xb5,
	0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x29, 0x9f, 0x4b, 0x02, 0x53, 0xaa, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x55, 0x28, 0x98, 0x8b, 0x03, 0x26, 0x26, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x64, 0xae, 0x47,
	0x92, 0x37, 0xf5, 0xe0, 0x46, 0xc2, 0x0d, 0x32, 0x7a, 0xc6, 0xc8, 0xc5, 0x0a, 0xb6, 0x51, 0xe8,
	0x0e, 0x23, 0xc2, 0x7c, 0x21, 0x37, 0x12, 0x4d, 0xc6, 0xe1, 0x1f, 0x29, 0x77, 0x8a, 0xcd, 0x81,
	0x78, 0x5e, 0xc9, 0xbe, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x96, 0x42, 0xe6, 0xfa, 0x50, 0x03, 0xf5,
	0xb1, 0x04, 0x34, 0x8e, 0xb0, 0x07, 0x19, 0xe4, 0x14, 0x75, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x0e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0x70,
	0xc3, 0x9d, 0xc1, 0x26, 0x39, 0x43, 0x4d, 0x02, 0x19, 0xa0, 0x5f, 0x81, 0xcd, 0xf8, 0x92, 0xca,
	0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0xbc, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x49,
	0x18, 0xa0, 0x96, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a nextGame by index.
	NextGame(ctx context.Context, in *QueryGetNextGameRequest, opts ...grpc.CallOption) (*QueryGetNextGameResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) NextGame(ctx context.Context, in *QueryGetNextGameRequest, opts ...grpc.CallOption) (*QueryGetNextGameResponse, error) {
	out := new(QueryGetNextGameResponse)
	err := c.cc.Invoke(ctx, "/avalkov.cosmoscheckersgame.cosmoscheckersgame.Query/NextGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a nextGame by index.
	NextGame(context.Context, *QueryGetNextGameRequest) (*QueryGetNextGameResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) NextGame(ctx context.Context, req *QueryGetNextGameRequest) (*QueryGetNextGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextGame not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_NextGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetNextGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NextGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avalkov.cosmoscheckersgame.cosmoscheckersgame.Query/NextGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NextGame(ctx, req.(*QueryGetNextGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "avalkov.cosmoscheckersgame.cosmoscheckersgame.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NextGame",
			Handler:    _Query_NextGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmoscheckersgame/query.proto",
}

func (m *QueryGetNextGameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetNextGameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetNextGameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetNextGameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetNextGameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetNextGameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextGame != nil {
		{
			size, err := m.NextGame.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetNextGameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetNextGameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextGame != nil {
		l = m.NextGame.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetNextGameRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetNextGameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetNextGameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetNextGameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetNextGameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetNextGameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextGame", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NextGame == nil {
				m.NextGame = &NextGame{}
			}
			if err := m.NextGame.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
