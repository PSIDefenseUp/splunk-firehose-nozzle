// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/product_group_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A product group view.
type ProductGroupView struct {
	// Immutable. The resource name of the product group view.
	// Product group view resource names have the form:
	//
	// `customers/{customer_id}/productGroupViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductGroupView) Reset()         { *m = ProductGroupView{} }
func (m *ProductGroupView) String() string { return proto.CompactTextString(m) }
func (*ProductGroupView) ProtoMessage()    {}
func (*ProductGroupView) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3a6978f2d357ed9, []int{0}
}

func (m *ProductGroupView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductGroupView.Unmarshal(m, b)
}
func (m *ProductGroupView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductGroupView.Marshal(b, m, deterministic)
}
func (m *ProductGroupView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductGroupView.Merge(m, src)
}
func (m *ProductGroupView) XXX_Size() int {
	return xxx_messageInfo_ProductGroupView.Size(m)
}
func (m *ProductGroupView) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductGroupView.DiscardUnknown(m)
}

var xxx_messageInfo_ProductGroupView proto.InternalMessageInfo

func (m *ProductGroupView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductGroupView)(nil), "google.ads.googleads.v1.resources.ProductGroupView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/product_group_view.proto", fileDescriptor_e3a6978f2d357ed9)
}

var fileDescriptor_e3a6978f2d357ed9 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0x69, 0x5f, 0x5e, 0xc1, 0xa2, 0x20, 0x03, 0x41, 0x87, 0xa0, 0x13, 0x06, 0x8a, 0x90,
	0x50, 0xbc, 0x65, 0xa7, 0xec, 0x32, 0xf0, 0x20, 0x63, 0x87, 0x1e, 0xa4, 0x50, 0xb2, 0x36, 0x76,
	0xc1, 0xb5, 0xff, 0x92, 0xb4, 0xdd, 0x61, 0xec, 0xe2, 0x47, 0xf1, 0xe8, 0x47, 0xf1, 0x0b, 0x78,
	0xf5, 0xbc, 0x8f, 0xe0, 0x49, 0xba, 0x2c, 0xd9, 0x98, 0xa0, 0x78, 0x7b, 0xe0, 0xff, 0x7b, 0x9e,
	0x3c, 0x3c, 0xf1, 0x48, 0x0a, 0x90, 0x4e, 0x39, 0x66, 0x89, 0xc2, 0x5a, 0x36, 0xaa, 0xf6, 0xb1,
	0xe4, 0x0a, 0x2a, 0x19, 0x73, 0x85, 0x0b, 0x09, 0x49, 0x15, 0x97, 0x51, 0x2a, 0xa1, 0x2a, 0xa2,
	0x5a, 0xf0, 0x19, 0x2a, 0x24, 0x94, 0xd0, 0xea, 0x68, 0x03, 0x62, 0x89, 0x42, 0xd6, 0x8b, 0x6a,
	0x1f, 0x59, 0x6f, 0xfb, 0xdc, 0xc4, 0x17, 0x02, 0x3f, 0x0a, 0x3e, 0x4d, 0xa2, 0x31, 0x9f, 0xb0,
	0x5a, 0x80, 0xd4, 0x19, 0xed, 0xd3, 0x2d, 0xc0, 0xd8, 0xd6, 0xa7, 0xb3, 0xad, 0x13, 0xcb, 0x73,
	0x28, 0x59, 0x29, 0x20, 0x57, 0xfa, 0x7a, 0xf9, 0xee, 0x78, 0x47, 0x43, 0xdd, 0x6c, 0xd0, 0x14,
	0x0b, 0x04, 0x9f, 0xb5, 0x02, 0xef, 0xd0, 0x84, 0x44, 0x39, 0xcb, 0xf8, 0x89, 0x73, 0xe1, 0x5c,
	0xed, 0xf7, 0xfd, 0x0f, 0xfa, 0xff, 0x93, 0xde, 0x78, 0xd7, 0x9b, 0x96, 0x6b, 0x55, 0x08, 0x85,
	0x62, 0xc8, 0xf0, 0x6e, 0xd2, 0xe8, 0xc0, 0xe4, 0xdc, 0xb3, 0x8c, 0x93, 0xa7, 0x25, 0x9d, 0xfc,
	0xc1, 0xdd, 0xea, 0xc5, 0x95, 0x2a, 0x21, 0xe3, 0x52, 0xe1, 0xb9, 0x91, 0x0b, 0x33, 0xa4, 0xc5,
	0x14, 0x9e, 0x7f, 0xdf, 0x76, 0xd1, 0x7f, 0x76, 0xbd, 0x6e, 0x0c, 0x19, 0xfa, 0x75, 0xdd, 0xfe,
	0xf1, 0xee, 0xc3, 0xc3, 0x66, 0x9a, 0xa1, 0xf3, 0x70, 0xb7, 0xf6, 0xa6, 0x30, 0x65, 0x79, 0x8a,
	0x40, 0xa6, 0x38, 0xe5, 0xf9, 0x6a, 0x38, 0xbc, 0xe9, 0xfe, 0xc3, 0xa7, 0xf7, 0xac, 0x7a, 0x71,
	0xff, 0x0d, 0x28, 0x7d, 0x75, 0x3b, 0x03, 0x1d, 0x49, 0x13, 0x85, 0xb4, 0x6c, 0x54, 0xe0, 0xa3,
	0x91, 0x21, 0xdf, 0x0c, 0x13, 0xd2, 0x44, 0x85, 0x96, 0x09, 0x03, 0x3f, 0xb4, 0xcc, 0xd2, 0xed,
	0xea, 0x03, 0x21, 0x34, 0x51, 0x84, 0x58, 0x8a, 0x90, 0xc0, 0x27, 0xc4, 0x72, 0xe3, 0xbd, 0x55,
	0xd9, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x41, 0x87, 0x17, 0xa0, 0x02, 0x00, 0x00,
}