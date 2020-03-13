// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/structured_snippet_placeholder_field.proto

package enums

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

// Possible values for Structured Snippet placeholder fields.
type StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField int32

const (
	// Not specified.
	StructuredSnippetPlaceholderFieldEnum_UNSPECIFIED StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	StructuredSnippetPlaceholderFieldEnum_UNKNOWN StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField = 1
	// Data Type: STRING. The category of snippet of your products/services.
	// Must match exactly one of the predefined structured snippets headers.
	// For a list, visit
	// https://developers.google.com/adwords/api/docs/appendix/structured-snippet-headers
	StructuredSnippetPlaceholderFieldEnum_HEADER StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField = 2
	// Data Type: STRING_LIST. Text values that describe your products/services.
	// All text must be family safe. Special or non-ASCII characters are not
	// permitted. A snippet can be at most 25 characters.
	StructuredSnippetPlaceholderFieldEnum_SNIPPETS StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField = 3
)

var StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "HEADER",
	3: "SNIPPETS",
}

var StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"HEADER":      2,
	"SNIPPETS":    3,
}

func (x StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField) String() string {
	return proto.EnumName(StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_name, int32(x))
}

func (StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96b948df16ee6043, []int{0, 0}
}

// Values for Structured Snippet placeholder fields.
type StructuredSnippetPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StructuredSnippetPlaceholderFieldEnum) Reset()         { *m = StructuredSnippetPlaceholderFieldEnum{} }
func (m *StructuredSnippetPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*StructuredSnippetPlaceholderFieldEnum) ProtoMessage()    {}
func (*StructuredSnippetPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b948df16ee6043, []int{0}
}

func (m *StructuredSnippetPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *StructuredSnippetPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *StructuredSnippetPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.Merge(m, src)
}
func (m *StructuredSnippetPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.Size(m)
}
func (m *StructuredSnippetPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField", StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_name, StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_value)
	proto.RegisterType((*StructuredSnippetPlaceholderFieldEnum)(nil), "google.ads.googleads.v3.enums.StructuredSnippetPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/structured_snippet_placeholder_field.proto", fileDescriptor_96b948df16ee6043)
}

var fileDescriptor_96b948df16ee6043 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xbd, 0x4e, 0xc3, 0x30,
	0x18, 0x24, 0xa9, 0x54, 0x90, 0x8b, 0x44, 0x94, 0x11, 0xd1, 0xa1, 0x95, 0x60, 0x74, 0x86, 0x6c,
	0x66, 0x4a, 0x69, 0xfa, 0x23, 0xa4, 0x10, 0x11, 0x5a, 0x24, 0x88, 0x54, 0x99, 0xda, 0x98, 0x48,
	0xa9, 0x6d, 0xc5, 0x4e, 0x9f, 0x80, 0x27, 0x61, 0xe4, 0x51, 0x78, 0x14, 0x5e, 0x81, 0x05, 0xc5,
	0xa6, 0x61, 0x82, 0x2e, 0xd6, 0xc9, 0xdf, 0x7d, 0x77, 0xdf, 0x1d, 0x98, 0x31, 0x21, 0x58, 0x49,
	0x03, 0x4c, 0x54, 0x60, 0x61, 0x83, 0xb6, 0x61, 0x40, 0x79, 0xbd, 0x51, 0x81, 0xd2, 0x55, 0xbd,
	0xd6, 0x75, 0x45, 0xc9, 0x4a, 0xf1, 0x42, 0x4a, 0xaa, 0x57, 0xb2, 0xc4, 0x6b, 0xfa, 0x22, 0x4a,
	0x42, 0xab, 0xd5, 0x73, 0x41, 0x4b, 0x02, 0x65, 0x25, 0xb4, 0xf0, 0xfb, 0x76, 0x1d, 0x62, 0xa2,
	0x60, 0xab, 0x04, 0xb7, 0x21, 0x34, 0x4a, 0xa7, 0x67, 0x3b, 0x23, 0x59, 0x04, 0x98, 0x73, 0xa1,
	0xb1, 0x2e, 0x04, 0x57, 0x76, 0x79, 0xf8, 0xea, 0x80, 0xf3, 0xac, 0xf5, 0xca, 0xac, 0x55, 0xfa,
	0xeb, 0x34, 0x69, 0x8c, 0x62, 0x5e, 0x6f, 0x86, 0x8f, 0x60, 0xb0, 0x97, 0xe8, 0x9f, 0x80, 0xde,
	0x22, 0xc9, 0xd2, 0xf8, 0x6a, 0x3e, 0x99, 0xc7, 0x63, 0xef, 0xc0, 0xef, 0x81, 0xc3, 0x45, 0x72,
	0x9d, 0xdc, 0xdc, 0x27, 0x9e, 0xe3, 0x03, 0xd0, 0x9d, 0xc5, 0xd1, 0x38, 0xbe, 0xf5, 0x5c, 0xff,
	0x18, 0x1c, 0x65, 0xc9, 0x3c, 0x4d, 0xe3, 0xbb, 0xcc, 0xeb, 0x8c, 0xbe, 0x1c, 0x30, 0x58, 0x8b,
	0x0d, 0xfc, 0x37, 0xca, 0xe8, 0x62, 0xef, 0x01, 0x69, 0x13, 0x2a, 0x75, 0x1e, 0x46, 0x3f, 0x42,
	0x4c, 0x94, 0x98, 0x33, 0x28, 0x2a, 0x16, 0x30, 0xca, 0x4d, 0xe4, 0x5d, 0xdb, 0xb2, 0x50, 0x7f,
	0x94, 0x7f, 0x69, 0xde, 0x37, 0xb7, 0x33, 0x8d, 0xa2, 0x77, 0xb7, 0x3f, 0xb5, 0x52, 0x11, 0x51,
	0xd0, 0xc2, 0x06, 0x2d, 0x43, 0xd8, 0xb4, 0xa2, 0x3e, 0x76, 0xf3, 0x3c, 0x22, 0x2a, 0x6f, 0xe7,
	0xf9, 0x32, 0xcc, 0xcd, 0xfc, 0xd3, 0x1d, 0xd8, 0x4f, 0x84, 0x22, 0xa2, 0x10, 0x6a, 0x19, 0x08,
	0x2d, 0x43, 0x84, 0x0c, 0xe7, 0xa9, 0x6b, 0x0e, 0x0b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0x0d, 0x39, 0x40, 0x14, 0x02, 0x00, 0x00,
}