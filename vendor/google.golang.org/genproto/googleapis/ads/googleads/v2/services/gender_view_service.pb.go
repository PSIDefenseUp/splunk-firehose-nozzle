// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/gender_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for [GenderViewService.GetGenderView][google.ads.googleads.v2.services.GenderViewService.GetGenderView].
type GetGenderViewRequest struct {
	// Required. The resource name of the gender view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGenderViewRequest) Reset()         { *m = GetGenderViewRequest{} }
func (m *GetGenderViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetGenderViewRequest) ProtoMessage()    {}
func (*GetGenderViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9903af7aa827f0d3, []int{0}
}

func (m *GetGenderViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGenderViewRequest.Unmarshal(m, b)
}
func (m *GetGenderViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGenderViewRequest.Marshal(b, m, deterministic)
}
func (m *GetGenderViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGenderViewRequest.Merge(m, src)
}
func (m *GetGenderViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetGenderViewRequest.Size(m)
}
func (m *GetGenderViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGenderViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGenderViewRequest proto.InternalMessageInfo

func (m *GetGenderViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGenderViewRequest)(nil), "google.ads.googleads.v2.services.GetGenderViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/gender_view_service.proto", fileDescriptor_9903af7aa827f0d3)
}

var fileDescriptor_9903af7aa827f0d3 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0xeb, 0xd3, 0x40,
	0x18, 0x26, 0xf9, 0x81, 0x60, 0xf0, 0x37, 0x18, 0x44, 0x4b, 0x14, 0x2c, 0xa5, 0x43, 0x29, 0xf4,
	0x0e, 0x52, 0x70, 0x38, 0x11, 0xbc, 0x82, 0xc4, 0x49, 0x4a, 0x85, 0x0c, 0x12, 0x08, 0xd7, 0xe4,
	0x35, 0x1e, 0x24, 0xb9, 0x7a, 0x97, 0xa6, 0x83, 0xb8, 0xf8, 0x15, 0xfc, 0x06, 0x8e, 0x7e, 0x07,
	0xbf, 0x40, 0x57, 0x37, 0x27, 0x07, 0x27, 0x27, 0xbf, 0x80, 0x20, 0xe9, 0xe5, 0x92, 0x56, 0x2d,
	0xdd, 0x1e, 0xee, 0xf9, 0xf3, 0xbe, 0xef, 0x93, 0x38, 0x24, 0x13, 0x22, 0xcb, 0x01, 0xb3, 0x54,
	0x61, 0x0d, 0x1b, 0x54, 0xfb, 0x58, 0x81, 0xac, 0x79, 0x02, 0x0a, 0x67, 0x50, 0xa6, 0x20, 0xe3,
	0x9a, 0xc3, 0x2e, 0x6e, 0x1f, 0xd1, 0x46, 0x8a, 0x4a, 0xb8, 0x43, 0x6d, 0x40, 0x2c, 0x55, 0xa8,
	0xf3, 0xa2, 0xda, 0x47, 0xc6, 0xeb, 0xcd, 0xcf, 0xa5, 0x4b, 0x50, 0x62, 0x2b, 0xff, 0x8a, 0xd7,
	0xb1, 0xde, 0x03, 0x63, 0xda, 0x70, 0xcc, 0xca, 0x52, 0x54, 0xac, 0xe2, 0xa2, 0x54, 0x2d, 0x7b,
	0xef, 0x88, 0x4d, 0x72, 0x0e, 0x65, 0xd5, 0x12, 0x0f, 0x8f, 0x88, 0xd7, 0x1c, 0xf2, 0x34, 0x5e,
	0xc3, 0x1b, 0x56, 0x73, 0x21, 0xb5, 0x60, 0xf4, 0xd4, 0xb9, 0x13, 0x40, 0x15, 0x1c, 0xe6, 0x85,
	0x1c, 0x76, 0x2b, 0x78, 0xbb, 0x05, 0x55, 0xb9, 0x13, 0xe7, 0xda, 0xac, 0x13, 0x97, 0xac, 0x80,
	0x81, 0x35, 0xb4, 0x26, 0x37, 0x17, 0x57, 0xdf, 0xa9, 0xbd, 0xba, 0x65, 0x98, 0x17, 0xac, 0x00,
	0xff, 0x97, 0xe5, 0xdc, 0xee, 0xfd, 0x2f, 0xf5, 0x95, 0xee, 0x17, 0xcb, 0xb9, 0x3e, 0x09, 0x76,
	0x1f, 0xa1, 0x4b, 0xcd, 0xa0, 0xff, 0x6d, 0xe2, 0xcd, 0xce, 0xfa, 0xba, 0xbe, 0x50, 0xef, 0x1a,
	0x3d, 0xfb, 0x46, 0x4f, 0x37, 0xff, 0xf0, 0xf5, 0xc7, 0x47, 0x1b, 0xbb, 0xb3, 0xa6, 0xe1, 0x77,
	0x27, 0xcc, 0x93, 0x64, 0xab, 0x2a, 0x51, 0x80, 0x54, 0x78, 0xda, 0x56, 0xde, 0x44, 0x28, 0x3c,
	0x7d, 0xef, 0xdd, 0xdf, 0xd3, 0x41, 0x3f, 0xac, 0x45, 0x1b, 0xae, 0x50, 0x22, 0x8a, 0xc5, 0x6f,
	0xcb, 0x19, 0x27, 0xa2, 0xb8, 0x78, 0xd0, 0xe2, 0xee, 0x3f, 0xc5, 0x2c, 0x9b, 0xd6, 0x97, 0xd6,
	0xab, 0xe7, 0xad, 0x37, 0x13, 0x39, 0x2b, 0x33, 0x24, 0x64, 0xd6, 0x6c, 0x70, 0xf8, 0x26, 0xb8,
	0x9f, 0x76, 0xfe, 0x0f, 0x7c, 0x6c, 0xc0, 0x27, 0xfb, 0x2a, 0xa0, 0xf4, 0xb3, 0x3d, 0x0c, 0x74,
	0x20, 0x4d, 0x15, 0xd2, 0xb0, 0x41, 0xa1, 0x8f, 0xda, 0xc1, 0x6a, 0x6f, 0x24, 0x11, 0x4d, 0x55,
	0xd4, 0x49, 0xa2, 0xd0, 0x8f, 0x8c, 0xe4, 0xa7, 0x3d, 0xd6, 0xef, 0x84, 0xd0, 0x54, 0x11, 0xd2,
	0x89, 0x08, 0x09, 0x7d, 0x42, 0x8c, 0x6c, 0x7d, 0xe3, 0xb0, 0xe7, 0xfc, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0e, 0xff, 0x5b, 0x69, 0x28, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GenderViewServiceClient is the client API for GenderViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GenderViewServiceClient interface {
	// Returns the requested gender view in full detail.
	GetGenderView(ctx context.Context, in *GetGenderViewRequest, opts ...grpc.CallOption) (*resources.GenderView, error)
}

type genderViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenderViewServiceClient(cc grpc.ClientConnInterface) GenderViewServiceClient {
	return &genderViewServiceClient{cc}
}

func (c *genderViewServiceClient) GetGenderView(ctx context.Context, in *GetGenderViewRequest, opts ...grpc.CallOption) (*resources.GenderView, error) {
	out := new(resources.GenderView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.GenderViewService/GetGenderView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenderViewServiceServer is the server API for GenderViewService service.
type GenderViewServiceServer interface {
	// Returns the requested gender view in full detail.
	GetGenderView(context.Context, *GetGenderViewRequest) (*resources.GenderView, error)
}

// UnimplementedGenderViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGenderViewServiceServer struct {
}

func (*UnimplementedGenderViewServiceServer) GetGenderView(ctx context.Context, req *GetGenderViewRequest) (*resources.GenderView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenderView not implemented")
}

func RegisterGenderViewServiceServer(s *grpc.Server, srv GenderViewServiceServer) {
	s.RegisterService(&_GenderViewService_serviceDesc, srv)
}

func _GenderViewService_GetGenderView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGenderViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenderViewServiceServer).GetGenderView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.GenderViewService/GetGenderView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenderViewServiceServer).GetGenderView(ctx, req.(*GetGenderViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GenderViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.GenderViewService",
	HandlerType: (*GenderViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGenderView",
			Handler:    _GenderViewService_GetGenderView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/gender_view_service.proto",
}
