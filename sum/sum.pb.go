// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum.proto

package sum

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SumRequest struct {
	FirstOperand         int32    `protobuf:"varint,1,opt,name=firstOperand,proto3" json:"firstOperand,omitempty"`
	SecondOperand        int32    `protobuf:"varint,2,opt,name=secondOperand,proto3" json:"secondOperand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetFirstOperand() int32 {
	if m != nil {
		return m.FirstOperand
	}
	return 0
}

func (m *SumRequest) GetSecondOperand() int32 {
	if m != nil {
		return m.SecondOperand
	}
	return 0
}

type ResultReply struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultReply) Reset()         { *m = ResultReply{} }
func (m *ResultReply) String() string { return proto.CompactTextString(m) }
func (*ResultReply) ProtoMessage()    {}
func (*ResultReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{1}
}

func (m *ResultReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultReply.Unmarshal(m, b)
}
func (m *ResultReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultReply.Marshal(b, m, deterministic)
}
func (m *ResultReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultReply.Merge(m, src)
}
func (m *ResultReply) XXX_Size() int {
	return xxx_messageInfo_ResultReply.Size(m)
}
func (m *ResultReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultReply.DiscardUnknown(m)
}

var xxx_messageInfo_ResultReply proto.InternalMessageInfo

func (m *ResultReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "sum.SumRequest")
	proto.RegisterType((*ResultReply)(nil), "sum.ResultReply")
}

func init() { proto.RegisterFile("sum.proto", fileDescriptor_62743f9cdc99b9fd) }

var fileDescriptor_62743f9cdc99b9fd = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2e, 0xcd, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2e, 0xcd, 0x55, 0x0a, 0xe3, 0xe2, 0x0a, 0x2e,
	0xcd, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe2, 0xe2, 0x49, 0xcb, 0x2c, 0x2a,
	0x2e, 0xf1, 0x2f, 0x48, 0x2d, 0x4a, 0xcc, 0x4b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x42,
	0x11, 0x13, 0x52, 0xe1, 0xe2, 0x2d, 0x4e, 0x4d, 0xce, 0xcf, 0x4b, 0x81, 0x29, 0x62, 0x02, 0x2b,
	0x42, 0x15, 0x54, 0x52, 0xe5, 0xe2, 0x0e, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x09, 0x4a, 0x2d, 0xc8,
	0xa9, 0x14, 0x12, 0xe3, 0x62, 0x2b, 0x02, 0x73, 0xa1, 0x46, 0x42, 0x79, 0x46, 0x0e, 0x5c, 0xdc,
	0xc1, 0xa5, 0xb9, 0xce, 0xf9, 0xb9, 0x05, 0xa5, 0x25, 0xa9, 0x45, 0x42, 0x86, 0x5c, 0x5c, 0x50,
	0x76, 0x70, 0x69, 0xae, 0x10, 0xbf, 0x1e, 0xc8, 0xb1, 0x08, 0xe7, 0x49, 0x09, 0x80, 0x05, 0x90,
	0xcc, 0x55, 0x62, 0x48, 0x62, 0x03, 0x7b, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x61, 0xa4,
	0x12, 0xad, 0xd9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SumComputerClient is the client API for SumComputer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumComputerClient interface {
	ComputeSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*ResultReply, error)
}

type sumComputerClient struct {
	cc grpc.ClientConnInterface
}

func NewSumComputerClient(cc grpc.ClientConnInterface) SumComputerClient {
	return &sumComputerClient{cc}
}

func (c *sumComputerClient) ComputeSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*ResultReply, error) {
	out := new(ResultReply)
	err := c.cc.Invoke(ctx, "/sum.SumComputer/ComputeSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumComputerServer is the server API for SumComputer service.
type SumComputerServer interface {
	ComputeSum(context.Context, *SumRequest) (*ResultReply, error)
}

// UnimplementedSumComputerServer can be embedded to have forward compatible implementations.
type UnimplementedSumComputerServer struct {
}

func (*UnimplementedSumComputerServer) ComputeSum(ctx context.Context, req *SumRequest) (*ResultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeSum not implemented")
}

func RegisterSumComputerServer(s *grpc.Server, srv SumComputerServer) {
	s.RegisterService(&_SumComputer_serviceDesc, srv)
}

func _SumComputer_ComputeSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumComputerServer).ComputeSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumComputer/ComputeSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumComputerServer).ComputeSum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumComputer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sum.SumComputer",
	HandlerType: (*SumComputerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeSum",
			Handler:    _SumComputer_ComputeSum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sum.proto",
}
