// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/user_interest.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A user interest: a particular interest-based vertical to be targeted.
type UserInterest struct {
	// Immutable. The resource name of the user interest.
	// User interest resource names have the form:
	//
	// `customers/{customer_id}/userInterests/{user_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Taxonomy type of the user interest.
	TaxonomyType enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType `protobuf:"varint,2,opt,name=taxonomy_type,json=taxonomyType,proto3,enum=google.ads.googleads.v1.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType" json:"taxonomy_type,omitempty"`
	// Output only. The ID of the user interest.
	UserInterestId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=user_interest_id,json=userInterestId,proto3" json:"user_interest_id,omitempty"`
	// Output only. The name of the user interest.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The parent of the user interest.
	UserInterestParent *wrappers.StringValue `protobuf:"bytes,5,opt,name=user_interest_parent,json=userInterestParent,proto3" json:"user_interest_parent,omitempty"`
	// Output only. True if the user interest is launched to all channels and locales.
	LaunchedToAll *wrappers.BoolValue `protobuf:"bytes,6,opt,name=launched_to_all,json=launchedToAll,proto3" json:"launched_to_all,omitempty"`
	// Output only. Availability information of the user interest.
	Availabilities       []*common.CriterionCategoryAvailability `protobuf:"bytes,7,rep,name=availabilities,proto3" json:"availabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *UserInterest) Reset()         { *m = UserInterest{} }
func (m *UserInterest) String() string { return proto.CompactTextString(m) }
func (*UserInterest) ProtoMessage()    {}
func (*UserInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01746905ff80465b, []int{0}
}

func (m *UserInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInterest.Unmarshal(m, b)
}
func (m *UserInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInterest.Marshal(b, m, deterministic)
}
func (m *UserInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInterest.Merge(m, src)
}
func (m *UserInterest) XXX_Size() int {
	return xxx_messageInfo_UserInterest.Size(m)
}
func (m *UserInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInterest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInterest proto.InternalMessageInfo

func (m *UserInterest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserInterest) GetTaxonomyType() enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType {
	if m != nil {
		return m.TaxonomyType
	}
	return enums.UserInterestTaxonomyTypeEnum_UNSPECIFIED
}

func (m *UserInterest) GetUserInterestId() *wrappers.Int64Value {
	if m != nil {
		return m.UserInterestId
	}
	return nil
}

func (m *UserInterest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserInterest) GetUserInterestParent() *wrappers.StringValue {
	if m != nil {
		return m.UserInterestParent
	}
	return nil
}

func (m *UserInterest) GetLaunchedToAll() *wrappers.BoolValue {
	if m != nil {
		return m.LaunchedToAll
	}
	return nil
}

func (m *UserInterest) GetAvailabilities() []*common.CriterionCategoryAvailability {
	if m != nil {
		return m.Availabilities
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInterest)(nil), "google.ads.googleads.v1.resources.UserInterest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/user_interest.proto", fileDescriptor_01746905ff80465b)
}

var fileDescriptor_01746905ff80465b = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x55, 0xe2, 0xb6, 0x9f, 0x3e, 0xf7, 0x07, 0xb0, 0x58, 0x98, 0x52, 0x41, 0x8a, 0x54, 0x91,
	0x4d, 0x67, 0x94, 0x40, 0x59, 0x18, 0x21, 0x34, 0x29, 0xa8, 0x6a, 0x17, 0xa8, 0x0a, 0x21, 0x0b,
	0x88, 0xb0, 0x26, 0xf6, 0xd4, 0x1d, 0x34, 0x9e, 0xb1, 0x66, 0xc6, 0x81, 0xa8, 0x2a, 0x0f, 0xc3,
	0x92, 0x47, 0xe1, 0x29, 0xba, 0xee, 0x1b, 0xc0, 0x02, 0xa1, 0xda, 0x63, 0x77, 0x42, 0x15, 0x28,
	0xbb, 0x9b, 0xdc, 0x73, 0xce, 0x3d, 0x3e, 0x57, 0x77, 0xdc, 0x9d, 0x44, 0x88, 0x84, 0x11, 0x88,
	0x63, 0x05, 0xcb, 0xf2, 0xa2, 0x9a, 0x74, 0xa0, 0x24, 0x4a, 0xe4, 0x32, 0x22, 0x0a, 0xe6, 0x8a,
	0xc8, 0x90, 0x72, 0x4d, 0x24, 0x51, 0x1a, 0x64, 0x52, 0x68, 0xe1, 0x6d, 0x96, 0x58, 0x80, 0x63,
	0x05, 0x6a, 0x1a, 0x98, 0x74, 0x40, 0x4d, 0x5b, 0x7f, 0x31, 0x4f, 0x39, 0x12, 0x69, 0x2a, 0x38,
	0x8c, 0x24, 0xd5, 0x44, 0x52, 0xc1, 0xc3, 0x08, 0x6b, 0x92, 0x08, 0x39, 0x0d, 0xf1, 0x04, 0x53,
	0x86, 0xc7, 0x94, 0x51, 0x3d, 0x2d, 0x07, 0xad, 0x3f, 0x9f, 0xa7, 0x42, 0x78, 0x9e, 0xfe, 0xe6,
	0x2d, 0xd4, 0xf8, 0x93, 0xe0, 0x22, 0x9d, 0x86, 0x7a, 0x9a, 0x11, 0x23, 0x70, 0xbf, 0x12, 0xc8,
	0x28, 0x3c, 0xa2, 0x84, 0xc5, 0xe1, 0x98, 0x1c, 0xe3, 0x09, 0x15, 0xd2, 0x00, 0xee, 0x58, 0x80,
	0xca, 0xbd, 0x69, 0xdd, 0x33, 0xad, 0xe2, 0xd7, 0x38, 0x3f, 0x82, 0x1f, 0x25, 0xce, 0x32, 0x22,
	0x95, 0xe9, 0x6f, 0x58, 0x54, 0xcc, 0xb9, 0xd0, 0x58, 0x53, 0xc1, 0x4d, 0xf7, 0xc1, 0xf7, 0x45,
	0x77, 0xe5, 0x8d, 0x22, 0x72, 0xdf, 0xd8, 0xf3, 0xfa, 0xee, 0x6a, 0x35, 0x20, 0xe4, 0x38, 0x25,
	0x7e, 0xa3, 0xd5, 0x68, 0xff, 0xdf, 0xdb, 0x3e, 0x43, 0x8b, 0x3f, 0xd0, 0x43, 0x77, 0xeb, 0x32,
	0x48, 0x53, 0x65, 0x54, 0x81, 0x48, 0xa4, 0xd0, 0x56, 0xe9, 0xaf, 0x54, 0x1a, 0xaf, 0x70, 0x4a,
	0xbc, 0xcf, 0xee, 0xea, 0xcc, 0x57, 0xfb, 0xcd, 0x56, 0xa3, 0xbd, 0xd6, 0x1d, 0x80, 0x79, 0x0b,
	0x2a, 0x72, 0x03, 0xb6, 0xe2, 0xc0, 0xf0, 0x07, 0xd3, 0x8c, 0xbc, 0xe4, 0x79, 0x3a, 0xb7, 0xd9,
	0x73, 0xce, 0x90, 0xd3, 0x5f, 0xd1, 0xd6, 0x5f, 0xde, 0x81, 0x7b, 0x73, 0x76, 0x07, 0x34, 0xf6,
	0x9d, 0x56, 0xa3, 0xbd, 0xdc, 0xbd, 0x5b, 0x59, 0xa8, 0xd2, 0x03, 0xfb, 0x5c, 0x3f, 0x79, 0x3c,
	0xc4, 0x2c, 0x37, 0x4a, 0x6b, 0xb9, 0x35, 0x68, 0x3f, 0xf6, 0x76, 0xdc, 0x85, 0x22, 0x96, 0x85,
	0x82, 0xbf, 0x71, 0x85, 0xff, 0x5a, 0x4b, 0xca, 0x13, 0x4b, 0xa0, 0x80, 0x7b, 0xa7, 0xee, 0xed,
	0x59, 0x0b, 0x19, 0x96, 0x84, 0x6b, 0x7f, 0xf1, 0x1a, 0x32, 0xdb, 0x67, 0xc8, 0xf9, 0x87, 0xec,
	0x3d, 0xdb, 0xf1, 0x61, 0x31, 0xc6, 0xdb, 0x73, 0x6f, 0x30, 0x9c, 0xf3, 0xe8, 0x98, 0xc4, 0xa1,
	0x16, 0x21, 0x66, 0xcc, 0x5f, 0x2a, 0x26, 0xaf, 0x5f, 0x99, 0xdc, 0x13, 0x82, 0x59, 0xf6, 0x57,
	0x2b, 0xde, 0x40, 0x20, 0xc6, 0xbc, 0x0f, 0xee, 0x9a, 0x75, 0x00, 0x94, 0x28, 0xff, 0xbf, 0x96,
	0xd3, 0x5e, 0xee, 0x3e, 0x9b, 0xbb, 0xcb, 0xf2, 0x92, 0xc0, 0x6e, 0x75, 0x49, 0xbb, 0xe6, 0x90,
	0x90, 0x75, 0x47, 0x26, 0xea, 0x59, 0xe5, 0xe0, 0xfd, 0x39, 0x7a, 0x77, 0xcd, 0x8f, 0xf6, 0xba,
	0x51, 0xae, 0xb4, 0x48, 0x89, 0x54, 0xf0, 0xa4, 0x2a, 0x4f, 0xa1, 0x9d, 0x84, 0x82, 0x27, 0x33,
	0x1b, 0x38, 0xed, 0xfd, 0x6c, 0xb8, 0x5b, 0x91, 0x48, 0xc1, 0x5f, 0x9f, 0x89, 0xde, 0x2d, 0x7b,
	0xd6, 0xe1, 0x45, 0x52, 0x87, 0x8d, 0xb7, 0x07, 0x86, 0x97, 0x08, 0x86, 0x79, 0x02, 0x84, 0x4c,
	0x60, 0x42, 0x78, 0x91, 0x23, 0xbc, 0xb4, 0xfa, 0x87, 0x47, 0xeb, 0x69, 0x5d, 0x7d, 0x69, 0x3a,
	0x7b, 0x08, 0x7d, 0x6d, 0x6e, 0xee, 0x95, 0x92, 0x28, 0x56, 0xa0, 0x2c, 0x2f, 0xaa, 0x61, 0x07,
	0xf4, 0x2b, 0xe4, 0xb7, 0x0a, 0x33, 0x42, 0xb1, 0x1a, 0xd5, 0x98, 0xd1, 0xb0, 0x33, 0xaa, 0x31,
	0xe7, 0xcd, 0xad, 0xb2, 0x11, 0x04, 0x28, 0x56, 0x41, 0x50, 0xa3, 0x82, 0x60, 0xd8, 0x09, 0x82,
	0x1a, 0x37, 0x5e, 0x2a, 0xcc, 0x3e, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x20, 0xc9, 0xfb, 0xae,
	0x60, 0x05, 0x00, 0x00,
}