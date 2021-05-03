// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v2alpha/tapds.proto

package envoy_service_tap_v2alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TapResource struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config               *TapConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TapResource) Reset()         { *m = TapResource{} }
func (m *TapResource) String() string { return proto.CompactTextString(m) }
func (*TapResource) ProtoMessage()    {}
func (*TapResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eef591ebf2a5317, []int{0}
}

func (m *TapResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapResource.Unmarshal(m, b)
}
func (m *TapResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapResource.Marshal(b, m, deterministic)
}
func (m *TapResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapResource.Merge(m, src)
}
func (m *TapResource) XXX_Size() int {
	return xxx_messageInfo_TapResource.Size(m)
}
func (m *TapResource) XXX_DiscardUnknown() {
	xxx_messageInfo_TapResource.DiscardUnknown(m)
}

var xxx_messageInfo_TapResource proto.InternalMessageInfo

func (m *TapResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TapResource) GetConfig() *TapConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*TapResource)(nil), "envoy.service.tap.v2alpha.TapResource")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v2alpha/tapds.proto", fileDescriptor_4eef591ebf2a5317)
}

var fileDescriptor_4eef591ebf2a5317 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0xbb, 0x2e, 0xa2, 0xea, 0x72, 0x00, 0xb9, 0x87, 0x82, 0xa1, 0x2d, 0xa2, 0xb4, 0x45,
	0x1c, 0xd6, 0x95, 0x7b, 0x43, 0x3d, 0xb9, 0xa8, 0x67, 0x64, 0x7c, 0xe9, 0xa9, 0x1a, 0xec, 0x2d,
	0xac, 0x64, 0xef, 0x6e, 0xbd, 0x6b, 0xab, 0x5c, 0xaa, 0x28, 0xa7, 0xdc, 0xf3, 0x36, 0x79, 0x82,
	0x5c, 0xa3, 0xe4, 0x11, 0xf2, 0x14, 0x39, 0x45, 0xfe, 0x03, 0x82, 0x44, 0xce, 0x29, 0xb7, 0xd5,
	0x7c, 0xdf, 0xf7, 0x1b, 0x8f, 0x67, 0xf0, 0x27, 0xca, 0x33, 0xb1, 0xb5, 0x15, 0x4d, 0x32, 0x16,
	0x50, 0x5b, 0x83, 0xb4, 0x33, 0x07, 0x22, 0xb9, 0x81, 0xfc, 0x1d, 0x2a, 0x22, 0x13, 0xa1, 0x85,
	0xd9, 0x2b, 0x6c, 0xa4, 0xb2, 0x11, 0x0d, 0x92, 0x54, 0x36, 0x6b, 0x50, 0x12, 0x40, 0x32, 0x3b,
	0x73, 0xec, 0x90, 0xa9, 0x40, 0x64, 0x34, 0xd9, 0x96, 0x41, 0xeb, 0x73, 0x3d, 0x3f, 0x10, 0x71,
	0x2c, 0x78, 0xe5, 0x1b, 0xac, 0x85, 0x58, 0x47, 0xb4, 0xc0, 0x00, 0xe7, 0x42, 0x83, 0x66, 0x82,
	0x57, 0xed, 0xad, 0x77, 0x69, 0x28, 0xe1, 0xb0, 0x6e, 0x2b, 0x0d, 0x3a, 0xdd, 0xc9, 0x6f, 0x33,
	0x88, 0x58, 0x08, 0x9a, 0xda, 0xbb, 0x47, 0x29, 0x8c, 0x36, 0xb8, 0xe5, 0x83, 0xf4, 0xa8, 0x12,
	0x69, 0x12, 0x50, 0xb3, 0x8f, 0x1b, 0x1c, 0x62, 0xda, 0x45, 0x43, 0x34, 0x79, 0xed, 0xbe, 0xba,
	0x73, 0x1b, 0x89, 0x31, 0x44, 0x5e, 0x51, 0x34, 0xbf, 0xe3, 0x66, 0x20, 0xf8, 0x1f, 0xb6, 0xee,
	0x1a, 0x43, 0x34, 0x69, 0x39, 0x63, 0x52, 0x3b, 0x33, 0xf1, 0x41, 0xfe, 0x28, 0xbc, 0x5e, 0x95,
	0x71, 0x6e, 0x0c, 0xfc, 0xc6, 0x07, 0x39, 0xdf, 0x8d, 0xbf, 0x2c, 0x53, 0xe6, 0x2f, 0xdc, 0x59,
	0xea, 0x84, 0x42, 0xbc, 0x8f, 0x28, 0xf3, 0x7d, 0x45, 0x06, 0xc9, 0x48, 0xe6, 0x90, 0x7d, 0xc6,
	0xa3, 0x7f, 0x53, 0xaa, 0xb4, 0xf5, 0xa1, 0x56, 0x57, 0x52, 0x70, 0x45, 0x47, 0x2f, 0x26, 0xe8,
	0x2b, 0x32, 0x57, 0xb8, 0x3d, 0xa7, 0x91, 0x86, 0x03, 0xf2, 0xc7, 0x07, 0xc9, 0x5c, 0x7e, 0x84,
	0x1f, 0x3f, 0x6d, 0x3a, 0xea, 0xf1, 0x1f, 0xb7, 0x7f, 0x52, 0x1d, 0x6c, 0x9e, 0xf3, 0xeb, 0xa7,
	0xa7, 0xd7, 0xb7, 0xe7, 0x46, 0x7f, 0xd4, 0x3b, 0xba, 0x97, 0x99, 0x06, 0xf9, 0xbb, 0xfc, 0x99,
	0xaa, 0x30, 0xbc, 0x9c, 0xa1, 0xa9, 0xeb, 0x5e, 0x9c, 0x5c, 0x5e, 0x35, 0x8d, 0x0e, 0xc2, 0x5f,
	0x98, 0x28, 0xc1, 0x32, 0x11, 0xff, 0xb6, 0xf5, 0xbb, 0x71, 0xb1, 0x9f, 0xdf, 0xed, 0x22, 0xdf,
	0xff, 0x02, 0x9d, 0x21, 0xb4, 0x6a, 0x16, 0xb7, 0xf0, 0xed, 0x3e, 0x00, 0x00, 0xff, 0xff, 0xf3,
	0xba, 0xec, 0x43, 0xeb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapDiscoveryServiceClient is the client API for TapDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapDiscoveryServiceClient interface {
	StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error)
	DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error)
	FetchTapConfigs(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type tapDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapDiscoveryServiceClient(cc *grpc.ClientConn) TapDiscoveryServiceClient {
	return &tapDiscoveryServiceClient{cc}
}

func (c *tapDiscoveryServiceClient) StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[0], "/envoy.service.tap.v2alpha.TapDiscoveryService/StreamTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceStreamTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_StreamTapConfigsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceStreamTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[1], "/envoy.service.tap.v2alpha.TapDiscoveryService/DeltaTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceDeltaTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_DeltaTapConfigsClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceDeltaTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) FetchTapConfigs(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.tap.v2alpha.TapDiscoveryService/FetchTapConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TapDiscoveryServiceServer is the server API for TapDiscoveryService service.
type TapDiscoveryServiceServer interface {
	StreamTapConfigs(TapDiscoveryService_StreamTapConfigsServer) error
	DeltaTapConfigs(TapDiscoveryService_DeltaTapConfigsServer) error
	FetchTapConfigs(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedTapDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTapDiscoveryServiceServer struct {
}

func (*UnimplementedTapDiscoveryServiceServer) StreamTapConfigs(srv TapDiscoveryService_StreamTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) DeltaTapConfigs(srv TapDiscoveryService_DeltaTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) FetchTapConfigs(ctx context.Context, req *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTapConfigs not implemented")
}

func RegisterTapDiscoveryServiceServer(s *grpc.Server, srv TapDiscoveryServiceServer) {
	s.RegisterService(&_TapDiscoveryService_serviceDesc, srv)
}

func _TapDiscoveryService_StreamTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).StreamTapConfigs(&tapDiscoveryServiceStreamTapConfigsServer{stream})
}

type TapDiscoveryService_StreamTapConfigsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceStreamTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_DeltaTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).DeltaTapConfigs(&tapDiscoveryServiceDeltaTapConfigsServer{stream})
}

type TapDiscoveryService_DeltaTapConfigsServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceDeltaTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_FetchTapConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.tap.v2alpha.TapDiscoveryService/FetchTapConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TapDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v2alpha.TapDiscoveryService",
	HandlerType: (*TapDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTapConfigs",
			Handler:    _TapDiscoveryService_FetchTapConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTapConfigs",
			Handler:       _TapDiscoveryService_StreamTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaTapConfigs",
			Handler:       _TapDiscoveryService_DeltaTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v2alpha/tapds.proto",
}