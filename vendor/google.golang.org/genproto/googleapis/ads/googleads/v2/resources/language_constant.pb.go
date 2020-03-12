// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/language_constant.proto

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
	return fileDescriptor_7db6cb90de59ea00, []int{0}
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
	proto.RegisterType((*LanguageConstant)(nil), "google.ads.googleads.v2.resources.LanguageConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/language_constant.proto", fileDescriptor_7db6cb90de59ea00)
}

var fileDescriptor_7db6cb90de59ea00 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x99, 0x76, 0x57, 0x30, 0x2a, 0x48, 0x41, 0x18, 0xc7, 0x45, 0x77, 0x85, 0x81, 0x15,
	0x21, 0xd1, 0xfa, 0x07, 0x8c, 0x07, 0x49, 0x3d, 0x2c, 0x8a, 0xc8, 0x32, 0x62, 0x0f, 0x32, 0x30,
	0x64, 0x9a, 0x6c, 0x2c, 0x74, 0xf2, 0x96, 0x24, 0x1d, 0x0f, 0xe2, 0xc5, 0x8f, 0xb2, 0x47, 0x3f,
	0x8a, 0x9f, 0x62, 0xcf, 0xfb, 0x11, 0x3c, 0x49, 0xdb, 0xb4, 0x53, 0x66, 0xc0, 0x3f, 0xb7, 0xa7,
	0x7d, 0x9f, 0xdf, 0x93, 0x27, 0xe1, 0x45, 0x2f, 0x14, 0x80, 0x2a, 0x24, 0xe1, 0xc2, 0x92, 0x56,
	0xd6, 0x6a, 0x1d, 0x13, 0x23, 0x2d, 0x54, 0x26, 0x93, 0x96, 0x14, 0x5c, 0xab, 0x8a, 0x2b, 0xb9,
	0xc8, 0x40, 0x5b, 0xc7, 0xb5, 0xc3, 0xa5, 0x01, 0x07, 0xd1, 0x51, 0xeb, 0xc7, 0x5c, 0x58, 0xdc,
	0xa3, 0x78, 0x1d, 0xe3, 0x1e, 0x9d, 0xdc, 0xeb, 0xd2, 0xcb, 0x9c, 0x9c, 0xe5, 0xb2, 0x10, 0x8b,
	0xa5, 0xfc, 0xcc, 0xd7, 0x39, 0x98, 0x36, 0x63, 0x72, 0x7b, 0x60, 0xe8, 0x30, 0x3f, 0xba, 0xeb,
	0x47, 0xcd, 0xd7, 0xb2, 0x3a, 0x23, 0x5f, 0x0c, 0x2f, 0x4b, 0x69, 0xac, 0x9f, 0x1f, 0x0c, 0x50,
	0xae, 0x35, 0x38, 0xee, 0x72, 0xd0, 0x7e, 0x7a, 0xff, 0x3c, 0x44, 0x37, 0xdf, 0xf9, 0xe2, 0xaf,
	0x7d, 0xef, 0x28, 0x45, 0x37, 0xba, 0x43, 0x16, 0x9a, 0xaf, 0xe4, 0x78, 0x74, 0x38, 0x3a, 0xbe,
	0x9a, 0x3c, 0xbe, 0x60, 0xfb, 0xbf, 0xd8, 0x43, 0xf4, 0x60, 0x73, 0x0b, 0xaf, 0xca, 0xdc, 0xe2,
	0x0c, 0x56, 0x64, 0x3b, 0x69, 0x76, 0xbd, 0xcb, 0x79, 0xcf, 0x57, 0x32, 0x7a, 0x84, 0x82, 0x5c,
	0x8c, 0x83, 0xc3, 0xd1, 0xf1, 0xb5, 0xf8, 0x8e, 0x67, 0x71, 0xd7, 0x1b, 0xbf, 0xd1, 0xee, 0xf9,
	0xd3, 0x94, 0x17, 0x95, 0x4c, 0xc2, 0x0b, 0x16, 0xce, 0x82, 0x5c, 0x44, 0xcf, 0xd0, 0x5e, 0x06,
	0x42, 0x8e, 0xc3, 0x86, 0x39, 0xd8, 0x61, 0x3e, 0x38, 0x93, 0x6b, 0x35, 0x80, 0x1a, 0x7b, 0x8d,
	0x35, 0xbd, 0xf7, 0xfe, 0x19, 0xab, 0xed, 0xd1, 0x2b, 0x84, 0x1c, 0x37, 0x4a, 0x3a, 0xbe, 0x2c,
	0xe4, 0x78, 0xbf, 0x81, 0x27, 0x3b, 0x70, 0x02, 0x50, 0x0c, 0xd0, 0x01, 0x42, 0x3f, 0x5e, 0xb2,
	0xd9, 0x7f, 0x3c, 0x4f, 0x34, 0x2d, 0xb6, 0xfe, 0x58, 0xf2, 0x75, 0x67, 0x8d, 0xbe, 0x25, 0xdf,
	0x03, 0x34, 0xcd, 0x60, 0x85, 0xff, 0xba, 0x48, 0xc9, 0xad, 0xed, 0x23, 0x4e, 0xeb, 0xd6, 0xa7,
	0xa3, 0x4f, 0x6f, 0x3d, 0xab, 0xa0, 0xce, 0xc7, 0x60, 0x14, 0x51, 0x52, 0x37, 0x77, 0x22, 0x9b,
	0x96, 0x7f, 0x58, 0xef, 0x97, 0xbd, 0x3a, 0x0f, 0xc2, 0x13, 0xc6, 0x7e, 0x04, 0x47, 0x27, 0x6d,
	0x24, 0x13, 0x16, 0xb7, 0xb2, 0x56, 0x69, 0x8c, 0x67, 0x9d, 0xf3, 0x67, 0xe7, 0x99, 0x33, 0x61,
	0xe7, 0xbd, 0x67, 0x9e, 0xc6, 0xf3, 0xde, 0x73, 0x19, 0x4c, 0xdb, 0x01, 0xa5, 0x4c, 0x58, 0x4a,
	0x7b, 0x17, 0xa5, 0x69, 0x4c, 0x69, 0xef, 0x5b, 0x5e, 0x69, 0xca, 0x3e, 0xf9, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x4c, 0xee, 0x09, 0x79, 0x8a, 0x03, 0x00, 0x00,
}
