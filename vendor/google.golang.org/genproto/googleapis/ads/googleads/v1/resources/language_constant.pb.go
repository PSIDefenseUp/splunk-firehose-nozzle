// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/language_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A language.
type LanguageConstant struct {
	// Immutable. The resource name of the language constant.
	// Language constant resource names have the form:
	//
	// `languageConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the language constant.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The language code, e.g. "en_US", "en_AU", "es", "fr", etc.
	Code *wrappers.StringValue `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// Output only. The full name of the language in English, e.g., "English (US)", "Spanish",
	// etc.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Whether the language is targetable.
	Targetable           *wrappers.BoolValue `protobuf:"bytes,5,opt,name=targetable,proto3" json:"targetable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LanguageConstant) Reset()         { *m = LanguageConstant{} }
func (m *LanguageConstant) String() string { return proto.CompactTextString(m) }
func (*LanguageConstant) ProtoMessage()    {}
func (*LanguageConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_74e5fb6fb55865cb, []int{0}
}

func (m *LanguageConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LanguageConstant.Unmarshal(m, b)
}
func (m *LanguageConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LanguageConstant.Marshal(b, m, deterministic)
}
func (m *LanguageConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LanguageConstant.Merge(m, src)
}
func (m *LanguageConstant) XXX_Size() int {
	return xxx_messageInfo_LanguageConstant.Size(m)
}
func (m *LanguageConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_LanguageConstant.DiscardUnknown(m)
}

var xxx_messageInfo_LanguageConstant proto.InternalMessageInfo

func (m *LanguageConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *LanguageConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *LanguageConstant) GetCode() *wrappers.StringValue {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *LanguageConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *LanguageConstant) GetTargetable() *wrappers.BoolValue {
	if m != nil {
		return m.Targetable
	}
	return nil
}

func init() {
	proto.RegisterType((*LanguageConstant)(nil), "google.ads.googleads.v1.resources.LanguageConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/language_constant.proto", fileDescriptor_74e5fb6fb55865cb)
}

var fileDescriptor_74e5fb6fb55865cb = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x99, 0x76, 0x57, 0x30, 0x2a, 0x48, 0x41, 0x18, 0xc7, 0x45, 0x77, 0x85, 0x81, 0x15,
	0x21, 0xb1, 0xfe, 0x03, 0xe3, 0x41, 0x52, 0x0f, 0x8b, 0x22, 0xb2, 0x8c, 0xd8, 0x83, 0x0c, 0x0c,
	0x99, 0x26, 0x1b, 0x0b, 0x9d, 0xbc, 0x25, 0x49, 0xc7, 0x83, 0x78, 0xf1, 0xa3, 0xec, 0xd1, 0x8f,
	0xe2, 0xa7, 0xd8, 0xf3, 0x7e, 0x04, 0x4f, 0xd2, 0x36, 0xed, 0x94, 0x19, 0xf0, 0xcf, 0xed, 0x69,
	0xdf, 0xe7, 0xf7, 0xe4, 0x49, 0x78, 0xd1, 0x0b, 0x05, 0xa0, 0x0a, 0x49, 0xb8, 0xb0, 0xa4, 0x95,
	0xb5, 0x5a, 0xc7, 0xc4, 0x48, 0x0b, 0x95, 0xc9, 0xa4, 0x25, 0x05, 0xd7, 0xaa, 0xe2, 0x4a, 0x2e,
	0x32, 0xd0, 0xd6, 0x71, 0xed, 0x70, 0x69, 0xc0, 0x41, 0x74, 0xd4, 0xfa, 0x31, 0x17, 0x16, 0xf7,
	0x28, 0x5e, 0xc7, 0xb8, 0x47, 0x27, 0xf7, 0xba, 0xf4, 0x32, 0x27, 0x67, 0xb9, 0x2c, 0xc4, 0x62,
	0x29, 0x3f, 0xf3, 0x75, 0x0e, 0xa6, 0xcd, 0x98, 0xdc, 0x1e, 0x18, 0x3a, 0xcc, 0x8f, 0xee, 0xfa,
	0x51, 0xf3, 0xb5, 0xac, 0xce, 0xc8, 0x17, 0xc3, 0xcb, 0x52, 0x1a, 0xeb, 0xe7, 0x07, 0x03, 0x94,
	0x6b, 0x0d, 0x8e, 0xbb, 0x1c, 0xb4, 0x9f, 0xde, 0x3f, 0x0f, 0xd1, 0xcd, 0x77, 0xbe, 0xf8, 0x6b,
	0xdf, 0x3b, 0x4a, 0xd1, 0x8d, 0xee, 0x90, 0x85, 0xe6, 0x2b, 0x39, 0x1e, 0x1d, 0x8e, 0x8e, 0xaf,
	0x26, 0xf1, 0x05, 0xdb, 0xff, 0xc5, 0x1e, 0xa2, 0x07, 0x9b, 0x5b, 0x78, 0x55, 0xe6, 0x16, 0x67,
	0xb0, 0x22, 0xdb, 0x49, 0xb3, 0xeb, 0x5d, 0xce, 0x7b, 0xbe, 0x92, 0xd1, 0x23, 0x14, 0xe4, 0x62,
	0x1c, 0x1c, 0x8e, 0x8e, 0xaf, 0x3d, 0xbe, 0xe3, 0x59, 0xdc, 0xf5, 0xc6, 0x6f, 0xb4, 0x7b, 0xfe,
	0x34, 0xe5, 0x45, 0x25, 0x93, 0xf0, 0x82, 0x85, 0xb3, 0x20, 0x17, 0xd1, 0x33, 0xb4, 0x97, 0x81,
	0x90, 0xe3, 0xb0, 0x61, 0x0e, 0x76, 0x98, 0x0f, 0xce, 0xe4, 0x5a, 0x0d, 0xa0, 0xc6, 0x5e, 0x63,
	0x4d, 0xef, 0xbd, 0x7f, 0xc6, 0x6a, 0x7b, 0xf4, 0x0a, 0x21, 0xc7, 0x8d, 0x92, 0x8e, 0x2f, 0x0b,
	0x39, 0xde, 0x6f, 0xe0, 0xc9, 0x0e, 0x9c, 0x00, 0x14, 0x03, 0x74, 0x80, 0xd0, 0x8f, 0x97, 0x6c,
	0xf6, 0x1f, 0xcf, 0x13, 0x4d, 0x8b, 0xad, 0x3f, 0x96, 0x7c, 0xdd, 0x59, 0xa3, 0x6f, 0xc9, 0xf7,
	0x00, 0x4d, 0x33, 0x58, 0xe1, 0xbf, 0x2e, 0x52, 0x72, 0x6b, 0xfb, 0x88, 0xd3, 0xba, 0xf5, 0xe9,
	0xe8, 0xd3, 0x5b, 0xcf, 0x2a, 0xa8, 0xf3, 0x31, 0x18, 0x45, 0x94, 0xd4, 0xcd, 0x9d, 0xc8, 0xa6,
	0xe5, 0x1f, 0xd6, 0xfb, 0x65, 0xaf, 0xce, 0x83, 0xf0, 0x84, 0xb1, 0x1f, 0xc1, 0xd1, 0x49, 0x1b,
	0xc9, 0x84, 0xc5, 0xad, 0xac, 0x55, 0x1a, 0xe3, 0x59, 0xe7, 0xfc, 0xd9, 0x79, 0xe6, 0x4c, 0xd8,
	0x79, 0xef, 0x99, 0xa7, 0xf1, 0xbc, 0xf7, 0x5c, 0x06, 0xd3, 0x76, 0x40, 0x29, 0x13, 0x96, 0xd2,
	0xde, 0x45, 0x69, 0x1a, 0x53, 0xda, 0xfb, 0x96, 0x57, 0x9a, 0xb2, 0x4f, 0x7e, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x52, 0x74, 0x28, 0x76, 0x8a, 0x03, 0x00, 0x00,
}