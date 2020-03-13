// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/shopping_performance_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [ShoppingPerformanceViewService.GetShoppingPerformanceView][google.ads.googleads.v1.services.ShoppingPerformanceViewService.GetShoppingPerformanceView].
type GetShoppingPerformanceViewRequest struct {
	// Required. The resource name of the Shopping performance view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShoppingPerformanceViewRequest) Reset()         { *m = GetShoppingPerformanceViewRequest{} }
func (m *GetShoppingPerformanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetShoppingPerformanceViewRequest) ProtoMessage()    {}
func (*GetShoppingPerformanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6c4cf22050d6d3, []int{0}
}

func (m *GetShoppingPerformanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Unmarshal(m, b)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.Merge(m, src)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Size(m)
}
func (m *GetShoppingPerformanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShoppingPerformanceViewRequest proto.InternalMessageInfo

func (m *GetShoppingPerformanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetShoppingPerformanceViewRequest)(nil), "google.ads.googleads.v1.services.GetShoppingPerformanceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/shopping_performance_view_service.proto", fileDescriptor_4d6c4cf22050d6d3)
}

var fileDescriptor_4d6c4cf22050d6d3 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0x55, 0x72, 0x12, 0x12, 0x11, 0x2c, 0x59, 0x38, 0x05, 0x04, 0xe5, 0xb8, 0xe1, 0xc4, 0x60,
	0x2b, 0x30, 0x20, 0x8c, 0x18, 0x5c, 0x86, 0xb2, 0x00, 0xd5, 0x9d, 0x94, 0x01, 0x45, 0x8a, 0xdc,
	0xe4, 0x6b, 0x6a, 0x29, 0xb1, 0x83, 0x9d, 0xa6, 0x03, 0x62, 0x61, 0x64, 0xe5, 0x1f, 0x30, 0x32,
	0xf0, 0x43, 0xba, 0xb2, 0x31, 0x31, 0x30, 0x31, 0xf0, 0x1b, 0x50, 0xea, 0x38, 0x6d, 0x4f, 0x4a,
	0xbb, 0x3d, 0xe5, 0xbd, 0xbc, 0xf7, 0xf9, 0x7b, 0x9f, 0xf7, 0x3a, 0x97, 0x32, 0x2f, 0x00, 0xb3,
	0x4c, 0x63, 0x03, 0x5b, 0xd4, 0x84, 0x58, 0x83, 0x6a, 0x78, 0x0a, 0x1a, 0xeb, 0x85, 0xac, 0x2a,
	0x2e, 0xf2, 0xa4, 0x02, 0x35, 0x97, 0xaa, 0x64, 0x22, 0x85, 0xa4, 0xe1, 0xb0, 0x4a, 0x3a, 0x09,
	0xaa, 0x94, 0xac, 0xa5, 0x3f, 0x32, 0xbf, 0x23, 0x96, 0x69, 0xd4, 0x3b, 0xa1, 0x26, 0x44, 0xd6,
	0x29, 0xa0, 0x43, 0x59, 0x0a, 0xb4, 0x5c, 0xaa, 0x83, 0x61, 0x26, 0x24, 0xb8, 0x67, 0x2d, 0x2a,
	0x8e, 0x99, 0x10, 0xb2, 0x66, 0x35, 0x97, 0x42, 0x77, 0xec, 0x9d, 0x1d, 0x36, 0x2d, 0x38, 0x88,
	0xba, 0x23, 0x1e, 0xec, 0x10, 0x73, 0x0e, 0x45, 0x96, 0xcc, 0x60, 0xc1, 0x1a, 0x2e, 0x95, 0x11,
	0x9c, 0xbd, 0xf1, 0x1e, 0x4e, 0xa0, 0xbe, 0xea, 0xd2, 0xa7, 0xdb, 0xf0, 0x88, 0xc3, 0xea, 0x12,
	0x3e, 0x2c, 0x41, 0xd7, 0xfe, 0x85, 0x77, 0xdb, 0x4e, 0x9a, 0x08, 0x56, 0xc2, 0xa9, 0x33, 0x72,
	0x2e, 0x6e, 0x8e, 0x4f, 0x7e, 0x53, 0xf7, 0xf2, 0x96, 0x65, 0xde, 0xb2, 0x12, 0x9e, 0xfc, 0x70,
	0xbd, 0xfb, 0x03, 0x66, 0x57, 0x66, 0x1b, 0xfe, 0x3f, 0xc7, 0x0b, 0x86, 0x23, 0xfd, 0x57, 0xe8,
	0xd8, 0x3a, 0xd1, 0xd1, 0x81, 0x03, 0x32, 0x68, 0xd2, 0x6f, 0x1c, 0x0d, 0x58, 0x9c, 0xbd, 0xfb,
	0x45, 0xf7, 0x5f, 0xfb, 0xf9, 0xe7, 0x9f, 0xaf, 0xee, 0x73, 0xff, 0x59, 0x5b, 0xd8, 0xc7, 0x3d,
	0xe6, 0x65, 0xba, 0xd4, 0xb5, 0x2c, 0x41, 0x69, 0xfc, 0xb8, 0x6f, 0xf0, 0x9a, 0xdf, 0xa7, 0xe0,
	0xee, 0x9a, 0x9e, 0x6e, 0x67, 0xe8, 0x50, 0xc5, 0x35, 0x4a, 0x65, 0x39, 0xfe, 0xe2, 0x7a, 0xe7,
	0xa9, 0x2c, 0x8f, 0x3e, 0x7a, 0xfc, 0xe8, 0xf0, 0x5a, 0xa7, 0x6d, 0x9b, 0x53, 0xe7, 0x7d, 0x77,
	0xd6, 0x28, 0x97, 0x05, 0x13, 0x39, 0x92, 0x2a, 0xc7, 0x39, 0x88, 0x4d, 0xd7, 0x78, 0x1b, 0x3d,
	0x7c, 0xf5, 0x2f, 0x2c, 0xf8, 0xe6, 0x9e, 0x4c, 0x28, 0xfd, 0xee, 0x8e, 0x26, 0xc6, 0x90, 0x66,
	0x1a, 0x19, 0xd8, 0xa2, 0x28, 0x44, 0x5d, 0xb0, 0x5e, 0x5b, 0x49, 0x4c, 0x33, 0x1d, 0xf7, 0x92,
	0x38, 0x0a, 0x63, 0x2b, 0xf9, 0xeb, 0x9e, 0x9b, 0xef, 0x84, 0xd0, 0x4c, 0x13, 0xd2, 0x8b, 0x08,
	0x89, 0x42, 0x42, 0xac, 0x6c, 0x76, 0x63, 0x33, 0xe7, 0xd3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x83, 0x0f, 0xde, 0x53, 0x9c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShoppingPerformanceViewServiceClient is the client API for ShoppingPerformanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShoppingPerformanceViewServiceClient interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error)
}

type shoppingPerformanceViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShoppingPerformanceViewServiceClient(cc grpc.ClientConnInterface) ShoppingPerformanceViewServiceClient {
	return &shoppingPerformanceViewServiceClient{cc}
}

func (c *shoppingPerformanceViewServiceClient) GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error) {
	out := new(resources.ShoppingPerformanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ShoppingPerformanceViewService/GetShoppingPerformanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingPerformanceViewServiceServer is the server API for ShoppingPerformanceViewService service.
type ShoppingPerformanceViewServiceServer interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(context.Context, *GetShoppingPerformanceViewRequest) (*resources.ShoppingPerformanceView, error)
}

// UnimplementedShoppingPerformanceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShoppingPerformanceViewServiceServer struct {
}

func (*UnimplementedShoppingPerformanceViewServiceServer) GetShoppingPerformanceView(ctx context.Context, req *GetShoppingPerformanceViewRequest) (*resources.ShoppingPerformanceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShoppingPerformanceView not implemented")
}

func RegisterShoppingPerformanceViewServiceServer(s *grpc.Server, srv ShoppingPerformanceViewServiceServer) {
	s.RegisterService(&_ShoppingPerformanceViewService_serviceDesc, srv)
}

func _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShoppingPerformanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ShoppingPerformanceViewService/GetShoppingPerformanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, req.(*GetShoppingPerformanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingPerformanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ShoppingPerformanceViewService",
	HandlerType: (*ShoppingPerformanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShoppingPerformanceView",
			Handler:    _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/shopping_performance_view_service.proto",
}