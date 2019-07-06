// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

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

// sum request and response
type SumRequest struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
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

func (m *SumRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *SumRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

// div request and response
type DivRequest struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DivRequest) Reset()         { *m = DivRequest{} }
func (m *DivRequest) String() string { return proto.CompactTextString(m) }
func (*DivRequest) ProtoMessage()    {}
func (*DivRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *DivRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DivRequest.Unmarshal(m, b)
}
func (m *DivRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DivRequest.Marshal(b, m, deterministic)
}
func (m *DivRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DivRequest.Merge(m, src)
}
func (m *DivRequest) XXX_Size() int {
	return xxx_messageInfo_DivRequest.Size(m)
}
func (m *DivRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DivRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DivRequest proto.InternalMessageInfo

func (m *DivRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *DivRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type DivResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DivResponse) Reset()         { *m = DivResponse{} }
func (m *DivResponse) String() string { return proto.CompactTextString(m) }
func (*DivResponse) ProtoMessage()    {}
func (*DivResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *DivResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DivResponse.Unmarshal(m, b)
}
func (m *DivResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DivResponse.Marshal(b, m, deterministic)
}
func (m *DivResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DivResponse.Merge(m, src)
}
func (m *DivResponse) XXX_Size() int {
	return xxx_messageInfo_DivResponse.Size(m)
}
func (m *DivResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DivResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DivResponse proto.InternalMessageInfo

func (m *DivResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

// multiply request and response
type MultiplyRequest struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyRequest) Reset()         { *m = MultiplyRequest{} }
func (m *MultiplyRequest) String() string { return proto.CompactTextString(m) }
func (*MultiplyRequest) ProtoMessage()    {}
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *MultiplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyRequest.Unmarshal(m, b)
}
func (m *MultiplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyRequest.Marshal(b, m, deterministic)
}
func (m *MultiplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyRequest.Merge(m, src)
}
func (m *MultiplyRequest) XXX_Size() int {
	return xxx_messageInfo_MultiplyRequest.Size(m)
}
func (m *MultiplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyRequest proto.InternalMessageInfo

func (m *MultiplyRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *MultiplyRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type MultiplyResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyResponse) Reset()         { *m = MultiplyResponse{} }
func (m *MultiplyResponse) String() string { return proto.CompactTextString(m) }
func (*MultiplyResponse) ProtoMessage()    {}
func (*MultiplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *MultiplyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyResponse.Unmarshal(m, b)
}
func (m *MultiplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyResponse.Marshal(b, m, deterministic)
}
func (m *MultiplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyResponse.Merge(m, src)
}
func (m *MultiplyResponse) XXX_Size() int {
	return xxx_messageInfo_MultiplyResponse.Size(m)
}
func (m *MultiplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyResponse proto.InternalMessageInfo

func (m *MultiplyResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

// sub request and response
type SubRequest struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubRequest) Reset()         { *m = SubRequest{} }
func (m *SubRequest) String() string { return proto.CompactTextString(m) }
func (*SubRequest) ProtoMessage()    {}
func (*SubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *SubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubRequest.Unmarshal(m, b)
}
func (m *SubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubRequest.Marshal(b, m, deterministic)
}
func (m *SubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubRequest.Merge(m, src)
}
func (m *SubRequest) XXX_Size() int {
	return xxx_messageInfo_SubRequest.Size(m)
}
func (m *SubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubRequest proto.InternalMessageInfo

func (m *SubRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *SubRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SubResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubResponse) Reset()         { *m = SubResponse{} }
func (m *SubResponse) String() string { return proto.CompactTextString(m) }
func (*SubResponse) ProtoMessage()    {}
func (*SubResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *SubResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubResponse.Unmarshal(m, b)
}
func (m *SubResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubResponse.Marshal(b, m, deterministic)
}
func (m *SubResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubResponse.Merge(m, src)
}
func (m *SubResponse) XXX_Size() int {
	return xxx_messageInfo_SubResponse.Size(m)
}
func (m *SubResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubResponse proto.InternalMessageInfo

func (m *SubResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*DivRequest)(nil), "calculator.DivRequest")
	proto.RegisterType((*DivResponse)(nil), "calculator.DivResponse")
	proto.RegisterType((*MultiplyRequest)(nil), "calculator.MultiplyRequest")
	proto.RegisterType((*MultiplyResponse)(nil), "calculator.MultiplyResponse")
	proto.RegisterType((*SubRequest)(nil), "calculator.SubRequest")
	proto.RegisterType((*SubResponse)(nil), "calculator.SubResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x47, 0x30, 0x0b, 0x92, 0x90, 0x38, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x08, 0x11, 0xa5, 0x10, 0x2e, 0xae, 0xe0, 0xd2, 0xdc, 0xa0, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x45, 0x2e, 0x9e, 0xb4, 0xcc, 0xa2, 0xe2, 0x92, 0xf8, 0xbc,
	0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x6e, 0xb0, 0x98, 0x1f,
	0x58, 0x48, 0x48, 0x99, 0x8b, 0xb7, 0x38, 0x35, 0x39, 0x3f, 0x2f, 0x05, 0xa6, 0x86, 0x09, 0xac,
	0x86, 0x07, 0x22, 0x08, 0x51, 0xa4, 0xa4, 0xca, 0xc5, 0x0d, 0x36, 0xb5, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x55, 0x48, 0x8c, 0x8b, 0xad, 0x28, 0xb5, 0xb8, 0x34, 0xa7, 0x04, 0x6a, 0x20, 0x94, 0x07,
	0xb2, 0xdc, 0x25, 0xb3, 0x8c, 0x06, 0x96, 0x83, 0x4d, 0x25, 0x60, 0x79, 0x24, 0x17, 0xbf, 0x6f,
	0x69, 0x4e, 0x49, 0x66, 0x41, 0x4e, 0x25, 0xb5, 0x5d, 0xa0, 0xc5, 0x25, 0x80, 0x30, 0x9a, 0x70,
	0x18, 0x04, 0x97, 0x26, 0xd1, 0x24, 0x02, 0x92, 0x08, 0x59, 0x6e, 0xd4, 0xc9, 0xc4, 0x25, 0xe0,
	0x0c, 0x4d, 0x0c, 0xa9, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x16, 0x5c, 0xcc, 0xc1,
	0xa5, 0xb9, 0x42, 0x62, 0x7a, 0x48, 0x09, 0x07, 0x91, 0x46, 0xa4, 0xc4, 0x31, 0xc4, 0x21, 0x96,
	0x28, 0x31, 0x80, 0x74, 0xba, 0x64, 0x96, 0xa1, 0xea, 0x44, 0x44, 0x30, 0xaa, 0x4e, 0xa4, 0x28,
	0x52, 0x62, 0x10, 0x72, 0xe7, 0xe2, 0x80, 0x85, 0x98, 0x90, 0x34, 0xb2, 0x32, 0xb4, 0x28, 0x92,
	0x92, 0xc1, 0x2e, 0x89, 0xec, 0x84, 0xe0, 0xd2, 0x24, 0x74, 0xc7, 0x27, 0xe1, 0x70, 0x7c, 0x12,
	0x42, 0xa7, 0x13, 0x5f, 0x14, 0x0f, 0x72, 0xb6, 0x49, 0x62, 0x03, 0x67, 0x16, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xed, 0x8a, 0xa3, 0xe0, 0x58, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Div(ctx context.Context, in *DivRequest, opts ...grpc.CallOption) (*DivResponse, error)
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
	Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error)
}

type calculateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculateServiceClient(cc *grpc.ClientConn) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) Div(ctx context.Context, in *DivRequest, opts ...grpc.CallOption) (*DivResponse, error) {
	out := new(DivResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error) {
	out := new(SubResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServiceServer is the server API for CalculateService service.
type CalculateServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Div(context.Context, *DivRequest) (*DivResponse, error)
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
	Sub(context.Context, *SubRequest) (*SubResponse, error)
}

// UnimplementedCalculateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (*UnimplementedCalculateServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculateServiceServer) Div(ctx context.Context, req *DivRequest) (*DivResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}
func (*UnimplementedCalculateServiceServer) Multiply(ctx context.Context, req *MultiplyRequest) (*MultiplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (*UnimplementedCalculateServiceServer) Sub(ctx context.Context, req *SubRequest) (*SubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}

func RegisterCalculateServiceServer(s *grpc.Server, srv CalculateServiceServer) {
	s.RegisterService(&_CalculateService_serviceDesc, srv)
}

func _CalculateService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Div(ctx, req.(*DivRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Sub(ctx, req.(*SubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculateService_Sum_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _CalculateService_Div_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CalculateService_Multiply_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _CalculateService_Sub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
