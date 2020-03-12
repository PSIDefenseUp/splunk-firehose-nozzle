// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/mobile_device_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
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

// Request message for [MobileDeviceConstantService.GetMobileDeviceConstant][google.ads.googleads.v3.services.MobileDeviceConstantService.GetMobileDeviceConstant].
type GetMobileDeviceConstantRequest struct {
	// Required. Resource name of the mobile device to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMobileDeviceConstantRequest) Reset()         { *m = GetMobileDeviceConstantRequest{} }
func (m *GetMobileDeviceConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetMobileDeviceConstantRequest) ProtoMessage()    {}
func (*GetMobileDeviceConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3657ef6f30658246, []int{0}
}

func (m *GetMobileDeviceConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Unmarshal(m, b)
}
func (m *GetMobileDeviceConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetMobileDeviceConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMobileDeviceConstantRequest.Merge(m, src)
}
func (m *GetMobileDeviceConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Size(m)
}
func (m *GetMobileDeviceConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMobileDeviceConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMobileDeviceConstantRequest proto.InternalMessageInfo

func (m *GetMobileDeviceConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMobileDeviceConstantRequest)(nil), "google.ads.googleads.v3.services.GetMobileDeviceConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/mobile_device_constant_service.proto", fileDescriptor_3657ef6f30658246)
}

var fileDescriptor_3657ef6f30658246 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4f, 0x8b, 0xd3, 0x40,
	0x1c, 0x25, 0x29, 0x08, 0x06, 0xbd, 0xe4, 0xd2, 0x92, 0x8a, 0x86, 0xd2, 0x43, 0x51, 0x9c, 0x01,
	0x73, 0x10, 0x46, 0x14, 0xa7, 0x55, 0x2a, 0x82, 0x52, 0x2a, 0xf4, 0x20, 0x81, 0x30, 0x4d, 0xc6,
	0x38, 0x90, 0xcc, 0xd4, 0x4c, 0x9a, 0x8b, 0x78, 0xf1, 0xe2, 0x07, 0x10, 0xf6, 0x03, 0xec, 0x71,
	0x3f, 0x4a, 0xaf, 0x7b, 0xda, 0x3d, 0xed, 0x61, 0x4f, 0xfb, 0x29, 0x96, 0x64, 0x32, 0x69, 0x0b,
	0x4d, 0x7b, 0x7b, 0xc9, 0x7b, 0xbf, 0xf7, 0x7e, 0x7f, 0xc6, 0xfa, 0x18, 0x0b, 0x11, 0x27, 0x14,
	0x92, 0x48, 0x42, 0x05, 0x4b, 0x54, 0x78, 0x50, 0xd2, 0xac, 0x60, 0x21, 0x95, 0x30, 0x15, 0x4b,
	0x96, 0xd0, 0x20, 0xa2, 0xe5, 0x67, 0x10, 0x0a, 0x2e, 0x73, 0xc2, 0xf3, 0xa0, 0xe6, 0xc1, 0x2a,
	0x13, 0xb9, 0xb0, 0x5d, 0x55, 0x0b, 0x48, 0x24, 0x41, 0x63, 0x03, 0x0a, 0x0f, 0x68, 0x1b, 0xe7,
	0x5d, 0x5b, 0x50, 0x46, 0xa5, 0x58, 0x67, 0xed, 0x49, 0x2a, 0xc1, 0x79, 0xa2, 0xeb, 0x57, 0x0c,
	0x12, 0xce, 0x45, 0x4e, 0x72, 0x26, 0xb8, 0xac, 0xd9, 0xee, 0x0e, 0x1b, 0x26, 0x8c, 0x36, 0x65,
	0xcf, 0x76, 0x88, 0x1f, 0x8c, 0x26, 0x51, 0xb0, 0xa4, 0x3f, 0x49, 0xc1, 0x44, 0xa6, 0x04, 0x83,
	0xcf, 0xd6, 0xd3, 0x29, 0xcd, 0xbf, 0x54, 0xd1, 0x1f, 0xaa, 0xe4, 0x49, 0x1d, 0x3c, 0xa7, 0xbf,
	0xd6, 0x54, 0xe6, 0xf6, 0xc8, 0x7a, 0xac, 0x7b, 0x0c, 0x38, 0x49, 0x69, 0xcf, 0x70, 0x8d, 0xd1,
	0xc3, 0x71, 0xe7, 0x06, 0x9b, 0xf3, 0x47, 0x9a, 0xf9, 0x4a, 0x52, 0xfa, 0xea, 0xcc, 0xb4, 0xfa,
	0x87, 0x9c, 0xbe, 0xa9, 0x25, 0xd8, 0x57, 0x86, 0xd5, 0x6d, 0x09, 0xb3, 0xdf, 0x83, 0x53, 0x2b,
	0x04, 0xc7, 0xfb, 0x74, 0x5e, 0xb7, 0x3a, 0x34, 0x2b, 0x06, 0x87, 0xea, 0x07, 0x93, 0x6b, 0xbc,
	0x3f, 0xe1, 0xdf, 0xcb, 0xdb, 0xff, 0xe6, 0x4b, 0xfb, 0x45, 0x79, 0x9e, 0xdf, 0x7b, 0xcc, 0xdb,
	0xf4, 0x80, 0x81, 0x84, 0xcf, 0xff, 0x38, 0xfd, 0x0d, 0xee, 0x6d, 0x43, 0x6b, 0xb4, 0x62, 0x12,
	0x84, 0x22, 0x1d, 0xff, 0x33, 0xad, 0x61, 0x28, 0xd2, 0x93, 0x23, 0x8e, 0xdd, 0x23, 0xeb, 0x9b,
	0x95, 0xf7, 0x9a, 0x19, 0xdf, 0x3f, 0xd5, 0x2e, 0xb1, 0x48, 0x08, 0x8f, 0x81, 0xc8, 0x62, 0x18,
	0x53, 0x5e, 0x5d, 0x13, 0x6e, 0x73, 0xdb, 0x5f, 0xf4, 0x1b, 0x0d, 0xce, 0xcd, 0xce, 0x14, 0xe3,
	0x0b, 0xd3, 0x9d, 0x2a, 0x43, 0x1c, 0x49, 0xa0, 0x60, 0x89, 0x16, 0x1e, 0xa8, 0x83, 0xe5, 0x46,
	0x4b, 0x7c, 0x1c, 0x49, 0xbf, 0x91, 0xf8, 0x0b, 0xcf, 0xd7, 0x92, 0x3b, 0x73, 0xa8, 0xfe, 0x23,
	0x84, 0x23, 0x89, 0x50, 0x23, 0x42, 0x68, 0xe1, 0x21, 0xa4, 0x65, 0xcb, 0x07, 0x55, 0x9f, 0xde,
	0x7d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x70, 0x2b, 0x8c, 0x78, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MobileDeviceConstantServiceClient is the client API for MobileDeviceConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MobileDeviceConstantServiceClient interface {
	// Returns the requested mobile device constant in full detail.
	GetMobileDeviceConstant(ctx context.Context, in *GetMobileDeviceConstantRequest, opts ...grpc.CallOption) (*resources.MobileDeviceConstant, error)
}

type mobileDeviceConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMobileDeviceConstantServiceClient(cc grpc.ClientConnInterface) MobileDeviceConstantServiceClient {
	return &mobileDeviceConstantServiceClient{cc}
}

func (c *mobileDeviceConstantServiceClient) GetMobileDeviceConstant(ctx context.Context, in *GetMobileDeviceConstantRequest, opts ...grpc.CallOption) (*resources.MobileDeviceConstant, error) {
	out := new(resources.MobileDeviceConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.MobileDeviceConstantService/GetMobileDeviceConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileDeviceConstantServiceServer is the server API for MobileDeviceConstantService service.
type MobileDeviceConstantServiceServer interface {
	// Returns the requested mobile device constant in full detail.
	GetMobileDeviceConstant(context.Context, *GetMobileDeviceConstantRequest) (*resources.MobileDeviceConstant, error)
}

// UnimplementedMobileDeviceConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMobileDeviceConstantServiceServer struct {
}

func (*UnimplementedMobileDeviceConstantServiceServer) GetMobileDeviceConstant(ctx context.Context, req *GetMobileDeviceConstantRequest) (*resources.MobileDeviceConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMobileDeviceConstant not implemented")
}

func RegisterMobileDeviceConstantServiceServer(s *grpc.Server, srv MobileDeviceConstantServiceServer) {
	s.RegisterService(&_MobileDeviceConstantService_serviceDesc, srv)
}

func _MobileDeviceConstantService_GetMobileDeviceConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMobileDeviceConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileDeviceConstantServiceServer).GetMobileDeviceConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.MobileDeviceConstantService/GetMobileDeviceConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileDeviceConstantServiceServer).GetMobileDeviceConstant(ctx, req.(*GetMobileDeviceConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MobileDeviceConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.MobileDeviceConstantService",
	HandlerType: (*MobileDeviceConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMobileDeviceConstant",
			Handler:    _MobileDeviceConstantService_GetMobileDeviceConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/mobile_device_constant_service.proto",
}
