// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_audience_view.proto

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

// A campaign audience view.
// Includes performance data from interests and remarketing lists for Display
// Network and YouTube Network ads, and remarketing lists for search ads (RLSA),
// aggregated by campaign and audience criterion. This view only includes
// audiences attached at the campaign level.
type CampaignAudienceView struct {
	// Immutable. The resource name of the campaign audience view.
	// Campaign audience view resource names have the form:
	//
	// `customers/{customer_id}/campaignAudienceViews/{campaign_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignAudienceView) Reset()         { *m = CampaignAudienceView{} }
func (m *CampaignAudienceView) String() string { return proto.CompactTextString(m) }
func (*CampaignAudienceView) ProtoMessage()    {}
func (*CampaignAudienceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_466ee4be4842c8d4, []int{0}
}

func (m *CampaignAudienceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignAudienceView.Unmarshal(m, b)
}
func (m *CampaignAudienceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignAudienceView.Marshal(b, m, deterministic)
}
func (m *CampaignAudienceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignAudienceView.Merge(m, src)
}
func (m *CampaignAudienceView) XXX_Size() int {
	return xxx_messageInfo_CampaignAudienceView.Size(m)
}
func (m *CampaignAudienceView) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignAudienceView.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignAudienceView proto.InternalMessageInfo

func (m *CampaignAudienceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*CampaignAudienceView)(nil), "google.ads.googleads.v1.resources.CampaignAudienceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_audience_view.proto", fileDescriptor_466ee4be4842c8d4)
}

var fileDescriptor_466ee4be4842c8d4 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x69, 0x45, 0xc1, 0xa2, 0x97, 0xe1, 0xc1, 0x0d, 0x41, 0x27, 0x0c, 0xbc, 0x98, 0x50,
	0xc4, 0x4b, 0x04, 0x21, 0xdb, 0x61, 0xe0, 0x41, 0xc6, 0x0e, 0x3d, 0x8c, 0x42, 0xc9, 0x9a, 0x18,
	0x03, 0x6b, 0x52, 0x9a, 0xae, 0x3b, 0xc8, 0xae, 0x7e, 0x10, 0x8f, 0x7e, 0x14, 0x3f, 0x85, 0xe7,
	0xf9, 0x0d, 0x3c, 0x49, 0x97, 0x26, 0xdb, 0x61, 0x28, 0xde, 0x1e, 0xfc, 0x7f, 0xef, 0xf5, 0xf5,
	0x91, 0xe0, 0x9e, 0x2b, 0xc5, 0x67, 0x0c, 0x12, 0xaa, 0xa1, 0x91, 0xb5, 0xaa, 0x42, 0x58, 0x30,
	0xad, 0xe6, 0x45, 0xca, 0x34, 0x4c, 0x49, 0x96, 0x13, 0xc1, 0x65, 0x42, 0xe6, 0x54, 0x30, 0x99,
	0xb2, 0xa4, 0x12, 0x6c, 0x01, 0xf2, 0x42, 0x95, 0xaa, 0xd5, 0x35, 0x26, 0x40, 0xa8, 0x06, 0xce,
	0x0f, 0xaa, 0x10, 0x38, 0x7f, 0xe7, 0xdc, 0x7e, 0x22, 0x17, 0xf0, 0x49, 0xb0, 0x19, 0x4d, 0xa6,
	0xec, 0x99, 0x54, 0x42, 0x15, 0x26, 0xa3, 0xd3, 0xde, 0x02, 0xac, 0xad, 0x39, 0x9d, 0x6d, 0x9d,
	0x88, 0x94, 0xaa, 0x24, 0xa5, 0x50, 0x52, 0x9b, 0xeb, 0xe5, 0x97, 0x17, 0x9c, 0x0c, 0x9a, 0x76,
	0xb8, 0x29, 0x17, 0x09, 0xb6, 0x68, 0x4d, 0x82, 0x63, 0x1b, 0x94, 0x48, 0x92, 0xb1, 0x53, 0xef,
	0xc2, 0xbb, 0x3a, 0xec, 0xdf, 0x7e, 0xe2, 0xfd, 0x6f, 0x0c, 0x83, 0xeb, 0x4d, 0xd3, 0x46, 0xe5,
	0x42, 0x83, 0x54, 0x65, 0x70, 0x57, 0xda, 0xf8, 0xc8, 0x66, 0x3d, 0x92, 0x8c, 0xa1, 0xc5, 0x0a,
	0x97, 0xff, 0x4c, 0x68, 0x0d, 0xd2, 0xb9, 0x2e, 0x55, 0xc6, 0x0a, 0x0d, 0x5f, 0xac, 0x5c, 0xba,
	0x61, 0xb7, 0xd1, 0x1a, 0xd8, 0xb9, 0xf7, 0xb2, 0xff, 0xea, 0x07, 0xbd, 0x54, 0x65, 0xe0, 0xcf,
	0xc5, 0xfb, 0xed, 0x5d, 0x25, 0x46, 0xf5, 0x64, 0x23, 0x6f, 0xf2, 0xd0, 0xf8, 0xb9, 0x9a, 0x11,
	0xc9, 0x81, 0x2a, 0x38, 0xe4, 0x4c, 0xae, 0x07, 0x85, 0x9b, 0x7f, 0xf9, 0xe5, 0x41, 0xdc, 0x39,
	0xf5, 0xe6, 0xef, 0x0d, 0x31, 0x7e, 0xf7, 0xbb, 0x43, 0x13, 0x89, 0xa9, 0x06, 0x46, 0xd6, 0x2a,
	0x0a, 0xc1, 0xd8, 0x92, 0x1f, 0x96, 0x89, 0x31, 0xd5, 0xb1, 0x63, 0xe2, 0x28, 0x8c, 0x1d, 0xb3,
	0xf2, 0x7b, 0xe6, 0x80, 0x10, 0xa6, 0x1a, 0x21, 0x47, 0x21, 0x14, 0x85, 0x08, 0x39, 0x6e, 0x7a,
	0xb0, 0x2e, 0x7b, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x92, 0x80, 0x2e, 0xe4, 0xbc, 0x02, 0x00,
	0x00,
}