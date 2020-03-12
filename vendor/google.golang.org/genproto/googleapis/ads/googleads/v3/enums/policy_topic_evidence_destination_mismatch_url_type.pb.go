// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/policy_topic_evidence_destination_mismatch_url_type.proto

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

// The possible policy topic evidence destination mismatch url types.
type PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType int32

const (
	// No value has been specified.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_UNSPECIFIED PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_UNKNOWN PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 1
	// The display url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_DISPLAY_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 2
	// The final url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_FINAL_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 3
	// The final mobile url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_FINAL_MOBILE_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 4
	// The tracking url template, with substituted desktop url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_TRACKING_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 5
	// The tracking url template, with substituted mobile url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_MOBILE_TRACKING_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 6
)

var PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DISPLAY_URL",
	3: "FINAL_URL",
	4: "FINAL_MOBILE_URL",
	5: "TRACKING_URL",
	6: "MOBILE_TRACKING_URL",
}

var PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"DISPLAY_URL":         2,
	"FINAL_URL":           3,
	"FINAL_MOBILE_URL":    4,
	"TRACKING_URL":        5,
	"MOBILE_TRACKING_URL": 6,
}

func (x PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType) String() string {
	return proto.EnumName(PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_name, int32(x))
}

func (PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1fe4244cca06d810, []int{0, 0}
}

// Container for enum describing possible policy topic evidence destination
// mismatch url types.
type PolicyTopicEvidenceDestinationMismatchUrlTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) Reset() {
	*m = PolicyTopicEvidenceDestinationMismatchUrlTypeEnum{}
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) String() string {
	return proto.CompactTextString(m)
}
func (*PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) ProtoMessage() {}
func (*PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fe4244cca06d810, []int{0}
}

func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.Unmarshal(m, b)
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.Marshal(b, m, deterministic)
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.Merge(m, src)
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.Size(m)
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType", PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_name, PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_value)
	proto.RegisterType((*PolicyTopicEvidenceDestinationMismatchUrlTypeEnum)(nil), "google.ads.googleads.v3.enums.PolicyTopicEvidenceDestinationMismatchUrlTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/policy_topic_evidence_destination_mismatch_url_type.proto", fileDescriptor_1fe4244cca06d810)
}

var fileDescriptor_1fe4244cca06d810 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xc1, 0xaa, 0x9b, 0x40,
	0x14, 0xad, 0xa6, 0x4d, 0xe9, 0xa4, 0xa5, 0x62, 0x0b, 0x85, 0xd2, 0x2c, 0x92, 0x7d, 0x47, 0x5a,
	0x77, 0xd3, 0xd5, 0x98, 0x98, 0x20, 0x31, 0x46, 0x92, 0x98, 0xd0, 0x22, 0x88, 0xd5, 0xc1, 0x0e,
	0xe8, 0x8c, 0x38, 0x26, 0x90, 0x2f, 0xe8, 0x7f, 0x74, 0xd9, 0x45, 0x3f, 0xa4, 0xdf, 0xf1, 0x56,
	0xef, 0x2b, 0x1e, 0x8e, 0x26, 0x8f, 0xb7, 0x78, 0x0f, 0xb2, 0x91, 0x73, 0xef, 0xb9, 0xf7, 0x1c,
	0xe7, 0x5c, 0xb0, 0xcf, 0x38, 0xcf, 0x72, 0x62, 0xc4, 0xa9, 0x30, 0x5a, 0xd8, 0xa0, 0xa3, 0x69,
	0x10, 0x76, 0x28, 0x84, 0x51, 0xf2, 0x9c, 0x26, 0xa7, 0xa8, 0xe6, 0x25, 0x4d, 0x22, 0x72, 0xa4,
	0x29, 0x61, 0x09, 0x89, 0x52, 0x22, 0x6a, 0xca, 0xe2, 0x9a, 0x72, 0x16, 0x15, 0x54, 0x14, 0x71,
	0x9d, 0xfc, 0x8a, 0x0e, 0x55, 0x1e, 0xd5, 0xa7, 0x92, 0xc0, 0xb2, 0xe2, 0x35, 0xd7, 0x87, 0xad,
	0x1a, 0x8c, 0x53, 0x01, 0x2f, 0xc2, 0xf0, 0x68, 0x42, 0x29, 0xfc, 0xf1, 0xd3, 0xd9, 0xb7, 0xa4,
	0x46, 0xcc, 0x18, 0xaf, 0xa5, 0x9a, 0x68, 0x97, 0xc7, 0x37, 0x0a, 0xf8, 0xe2, 0x4b, 0xeb, 0x6d,
	0xe3, 0x6c, 0x77, 0xc6, 0xd3, 0x7b, 0xdf, 0x65, 0x67, 0x1b, 0x54, 0xf9, 0xf6, 0x54, 0x12, 0x9b,
	0x1d, 0x8a, 0xf1, 0x3f, 0x05, 0x7c, 0xbe, 0x6a, 0x4b, 0x7f, 0x0b, 0x06, 0x81, 0xb7, 0xf1, 0xed,
	0x89, 0x33, 0x73, 0xec, 0xa9, 0xf6, 0x4c, 0x1f, 0x80, 0x97, 0x81, 0xb7, 0xf0, 0x56, 0x7b, 0x4f,
	0x53, 0x1a, 0x76, 0xea, 0x6c, 0x7c, 0x17, 0x7f, 0x8f, 0x82, 0xb5, 0xab, 0xa9, 0xfa, 0x1b, 0xf0,
	0x6a, 0xe6, 0x78, 0xd8, 0x95, 0x65, 0x4f, 0x7f, 0x0f, 0xb4, 0xb6, 0x5c, 0xae, 0x2c, 0xc7, 0xb5,
	0x65, 0xf7, 0xb9, 0xae, 0x81, 0xd7, 0xdb, 0x35, 0x9e, 0x2c, 0x1c, 0x6f, 0x2e, 0x3b, 0x2f, 0xf4,
	0x0f, 0xe0, 0x5d, 0x37, 0xf1, 0x80, 0xe8, 0x5b, 0xbf, 0x55, 0x30, 0x4a, 0x78, 0x01, 0x9f, 0x8c,
	0xca, 0xfa, 0x7a, 0xd5, 0x9b, 0xfc, 0x26, 0x40, 0x5f, 0xf9, 0x61, 0x75, 0xa2, 0x19, 0xcf, 0x63,
	0x96, 0x41, 0x5e, 0x65, 0x46, 0x46, 0x98, 0x8c, 0xf7, 0x7c, 0xe8, 0x92, 0x8a, 0x47, 0xee, 0xfe,
	0x4d, 0x7e, 0xff, 0xa8, 0xbd, 0x39, 0xc6, 0x7f, 0xd5, 0xe1, 0xbc, 0x95, 0xc2, 0xa9, 0x80, 0x2d,
	0x6c, 0xd0, 0xce, 0x84, 0x4d, 0xea, 0xe2, 0xff, 0x99, 0x0f, 0x71, 0x2a, 0xc2, 0x0b, 0x1f, 0xee,
	0xcc, 0x50, 0xf2, 0xb7, 0xea, 0xa8, 0x6d, 0x22, 0x84, 0x53, 0x81, 0xd0, 0x65, 0x02, 0xa1, 0x9d,
	0x89, 0x90, 0x9c, 0xf9, 0xd9, 0x97, 0x3f, 0x66, 0xde, 0x05, 0x00, 0x00, 0xff, 0xff, 0x73, 0xb6,
	0x6d, 0xf9, 0x8f, 0x02, 0x00, 0x00,
}
