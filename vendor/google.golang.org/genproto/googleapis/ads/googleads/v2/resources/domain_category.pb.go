// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/domain_category.proto

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

// A category generated automatically by crawling a domain. If a campaign uses
// the DynamicSearchAdsSetting, then domain categories will be generated for
// the domain. The categories can be targeted using WebpageConditionInfo.
// See: https://support.google.com/google-ads/answer/2471185
type DomainCategory struct {
	// Immutable. The resource name of the domain category.
	// Domain category resource names have the form:
	//
	// `customers/{customer_id}/domainCategories/{campaign_id}~{category_base64}~{language_code}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The campaign this category is recommended for.
	Campaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Output only. Recommended category for the website domain. e.g. if you have a website
	// about electronics, the categories could be "cameras", "televisions", etc.
	Category *wrappers.StringValue `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	// Output only. The language code specifying the language of the website. e.g. "en" for
	// English. The language can be specified in the DynamicSearchAdsSetting
	// required for dynamic search ads. This is the language of the pages from
	// your website that you want Google Ads to find, create ads for,
	// and match searches with.
	LanguageCode *wrappers.StringValue `protobuf:"bytes,4,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Output only. The domain for the website. The domain can be specified in the
	// DynamicSearchAdsSetting required for dynamic search ads.
	Domain *wrappers.StringValue `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
	// Output only. Fraction of pages on your site that this category matches.
	CoverageFraction *wrappers.DoubleValue `protobuf:"bytes,6,opt,name=coverage_fraction,json=coverageFraction,proto3" json:"coverage_fraction,omitempty"`
	// Output only. The position of this category in the set of categories. Lower numbers
	// indicate a better match for the domain. null indicates not recommended.
	CategoryRank *wrappers.Int64Value `protobuf:"bytes,7,opt,name=category_rank,json=categoryRank,proto3" json:"category_rank,omitempty"`
	// Output only. Indicates whether this category has sub-categories.
	HasChildren *wrappers.BoolValue `protobuf:"bytes,8,opt,name=has_children,json=hasChildren,proto3" json:"has_children,omitempty"`
	// Output only. The recommended cost per click for the category.
	RecommendedCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,9,opt,name=recommended_cpc_bid_micros,json=recommendedCpcBidMicros,proto3" json:"recommended_cpc_bid_micros,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *DomainCategory) Reset()         { *m = DomainCategory{} }
func (m *DomainCategory) String() string { return proto.CompactTextString(m) }
func (*DomainCategory) ProtoMessage()    {}
func (*DomainCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b5d3cffbd1a298, []int{0}
}

func (m *DomainCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainCategory.Unmarshal(m, b)
}
func (m *DomainCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainCategory.Marshal(b, m, deterministic)
}
func (m *DomainCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainCategory.Merge(m, src)
}
func (m *DomainCategory) XXX_Size() int {
	return xxx_messageInfo_DomainCategory.Size(m)
}
func (m *DomainCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainCategory.DiscardUnknown(m)
}

var xxx_messageInfo_DomainCategory proto.InternalMessageInfo

func (m *DomainCategory) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DomainCategory) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *DomainCategory) GetCategory() *wrappers.StringValue {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *DomainCategory) GetLanguageCode() *wrappers.StringValue {
	if m != nil {
		return m.LanguageCode
	}
	return nil
}

func (m *DomainCategory) GetDomain() *wrappers.StringValue {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *DomainCategory) GetCoverageFraction() *wrappers.DoubleValue {
	if m != nil {
		return m.CoverageFraction
	}
	return nil
}

func (m *DomainCategory) GetCategoryRank() *wrappers.Int64Value {
	if m != nil {
		return m.CategoryRank
	}
	return nil
}

func (m *DomainCategory) GetHasChildren() *wrappers.BoolValue {
	if m != nil {
		return m.HasChildren
	}
	return nil
}

func (m *DomainCategory) GetRecommendedCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedCpcBidMicros
	}
	return nil
}

func init() {
	proto.RegisterType((*DomainCategory)(nil), "google.ads.googleads.v2.resources.DomainCategory")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/domain_category.proto", fileDescriptor_68b5d3cffbd1a298)
}

var fileDescriptor_68b5d3cffbd1a298 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0xd5, 0xf6, 0xbf, 0xfe, 0x37, 0x6f, 0x43, 0x10, 0x0e, 0x84, 0x32, 0xc1, 0x06, 0x9a,
	0xd8, 0x2e, 0xb6, 0x54, 0x10, 0x13, 0x41, 0x1c, 0x92, 0x8e, 0x4d, 0x20, 0x31, 0x4d, 0x05, 0xf5,
	0x80, 0x8a, 0x22, 0xd7, 0xf6, 0x52, 0x6b, 0x89, 0x1d, 0xd9, 0x49, 0x11, 0x9a, 0x26, 0xf1, 0x59,
	0x38, 0x72, 0xe0, 0x83, 0xf0, 0x29, 0x76, 0xde, 0x47, 0xd8, 0x09, 0x25, 0xb1, 0xb3, 0x8c, 0x69,
	0xd0, 0xdb, 0x5b, 0xbd, 0xcf, 0xf3, 0x7b, 0x9f, 0x57, 0xf5, 0x1b, 0xb0, 0x13, 0x49, 0x19, 0xc5,
	0x0c, 0x61, 0xaa, 0x51, 0x55, 0x16, 0xd5, 0xac, 0x8f, 0x14, 0xd3, 0x32, 0x57, 0x84, 0x69, 0x44,
	0x65, 0x82, 0xb9, 0x08, 0x09, 0xce, 0x58, 0x24, 0xd5, 0x57, 0x98, 0x2a, 0x99, 0x49, 0x67, 0xa3,
	0x52, 0x43, 0x4c, 0x35, 0xac, 0x8d, 0x70, 0xd6, 0x87, 0xb5, 0xb1, 0xf7, 0xc8, 0xb2, 0x53, 0x8e,
	0x8e, 0x38, 0x8b, 0x69, 0x38, 0x61, 0x53, 0x3c, 0xe3, 0x52, 0x55, 0x8c, 0xde, 0xfd, 0x86, 0xc0,
	0xda, 0x4c, 0xeb, 0xa1, 0x69, 0x95, 0xbf, 0x26, 0xf9, 0x11, 0xfa, 0xa2, 0x70, 0x9a, 0x32, 0xa5,
	0x4d, 0x7f, 0xad, 0x61, 0xc5, 0x42, 0xc8, 0x0c, 0x67, 0x5c, 0x0a, 0xd3, 0x7d, 0xfc, 0xb3, 0x0b,
	0x6e, 0xed, 0x96, 0xb1, 0x07, 0x26, 0xb5, 0xf3, 0x11, 0xac, 0xda, 0x11, 0xa1, 0xc0, 0x09, 0x73,
	0x5b, 0xeb, 0xad, 0xad, 0xa5, 0x00, 0x9d, 0xf9, 0x0b, 0x17, 0xfe, 0x36, 0x78, 0x7a, 0xb9, 0x83,
	0xa9, 0x52, 0xae, 0x21, 0x91, 0x09, 0xba, 0xca, 0x19, 0xae, 0x58, 0xca, 0x01, 0x4e, 0x98, 0x43,
	0xc0, 0x22, 0xc1, 0x49, 0x8a, 0x79, 0x24, 0xdc, 0xf6, 0x7a, 0x6b, 0x6b, 0xb9, 0xbf, 0x66, 0xfc,
	0xd0, 0x26, 0x87, 0x1f, 0x32, 0xc5, 0x45, 0x34, 0xc2, 0x71, 0xce, 0x82, 0xed, 0x33, 0xbf, 0x73,
	0xe1, 0x3f, 0x01, 0x1b, 0x37, 0x8e, 0x1b, 0x18, 0xdc, 0xb0, 0x06, 0x3b, 0xaf, 0x8b, 0x21, 0xd5,
	0x78, 0xb7, 0x33, 0xc7, 0x90, 0xce, 0x99, 0xdf, 0x19, 0xd6, 0x16, 0x67, 0x0f, 0xac, 0xc6, 0x58,
	0x44, 0x39, 0x8e, 0x58, 0x48, 0x24, 0x65, 0xee, 0x7f, 0xf3, 0x32, 0x56, 0xac, 0x6f, 0x20, 0x29,
	0x73, 0x5e, 0x82, 0x6e, 0xf5, 0x14, 0xdc, 0x85, 0x79, 0x01, 0xc6, 0xe0, 0x1c, 0x80, 0x3b, 0x44,
	0xce, 0x98, 0x2a, 0x22, 0x1c, 0x29, 0x4c, 0x8a, 0xff, 0xca, 0xed, 0xde, 0x40, 0xd9, 0x95, 0xf9,
	0x24, 0x66, 0x0d, 0xca, 0x6d, 0xeb, 0xdd, 0x33, 0x56, 0xe7, 0x0d, 0x58, 0xb5, 0xeb, 0x85, 0x0a,
	0x8b, 0x63, 0xf7, 0xff, 0x92, 0xf5, 0xe0, 0x1a, 0xeb, 0xad, 0xc8, 0x5e, 0x3c, 0x6f, 0x6e, 0x64,
	0x6d, 0x43, 0x2c, 0x8e, 0x9d, 0x00, 0xac, 0x4c, 0xb1, 0x0e, 0xc9, 0x94, 0xc7, 0x54, 0x31, 0xe1,
	0x2e, 0x96, 0x94, 0xde, 0x35, 0x4a, 0x20, 0x65, 0xdc, 0x80, 0x2c, 0x4f, 0xb1, 0x1e, 0x18, 0x8f,
	0xf3, 0x19, 0xf4, 0x14, 0x23, 0x32, 0x49, 0x98, 0xa0, 0x8c, 0x86, 0x24, 0x25, 0xe1, 0x84, 0xd3,
	0x30, 0xe1, 0x44, 0x49, 0xed, 0x2e, 0xcd, 0x99, 0xeb, 0x5e, 0x83, 0x31, 0x48, 0x49, 0xc0, 0xe9,
	0xfb, 0x12, 0xe0, 0xb1, 0x73, 0x7f, 0x32, 0xf7, 0xe3, 0x74, 0x76, 0x48, 0xae, 0x33, 0x99, 0x30,
	0xa5, 0xd1, 0x89, 0x2d, 0x4f, 0xcd, 0x01, 0x1b, 0x11, 0x67, 0x1a, 0x9d, 0xfc, 0x71, 0xd2, 0xa7,
	0xc1, 0xb7, 0x36, 0xd8, 0x24, 0x32, 0x81, 0xff, 0x3c, 0xea, 0xe0, 0xee, 0xd5, 0x91, 0x87, 0xc5,
	0x46, 0x87, 0xad, 0x4f, 0xef, 0x8c, 0x33, 0x92, 0xc5, 0x9b, 0x81, 0x52, 0x45, 0x28, 0x62, 0xa2,
	0xdc, 0x17, 0x5d, 0x66, 0xfe, 0xcb, 0x67, 0xe6, 0x55, 0x5d, 0x7d, 0x6f, 0x77, 0xf6, 0x7d, 0xff,
	0x47, 0x7b, 0x63, 0xbf, 0x42, 0xfa, 0x54, 0xc3, 0xaa, 0x2c, 0xaa, 0x51, 0x1f, 0x0e, 0xad, 0xf2,
	0x97, 0xd5, 0x8c, 0x7d, 0xaa, 0xc7, 0xb5, 0x66, 0x3c, 0xea, 0x8f, 0x6b, 0xcd, 0x79, 0x7b, 0xb3,
	0x6a, 0x78, 0x9e, 0x4f, 0xb5, 0xe7, 0xd5, 0x2a, 0xcf, 0x1b, 0xf5, 0x3d, 0xaf, 0xd6, 0x4d, 0xba,
	0x65, 0xd8, 0x67, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xbf, 0xc8, 0x0b, 0x12, 0x05, 0x00,
	0x00,
}