// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/user_interest.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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
	TaxonomyType enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType `protobuf:"varint,2,opt,name=taxonomy_type,json=taxonomyType,proto3,enum=google.ads.googleads.v2.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType" json:"taxonomy_type,omitempty"`
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
	return fileDescriptor_3f0e313311c98402, []int{0}
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
	proto.RegisterType((*UserInterest)(nil), "google.ads.googleads.v2.resources.UserInterest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/user_interest.proto", fileDescriptor_3f0e313311c98402)
}

var fileDescriptor_3f0e313311c98402 = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x55, 0xe2, 0xb6, 0x9f, 0x3e, 0xf7, 0x07, 0xb0, 0x58, 0x98, 0x52, 0x41, 0x8a, 0x54, 0x91,
	0x4d, 0x67, 0x24, 0x43, 0x59, 0x18, 0x21, 0x34, 0x29, 0xa8, 0x6a, 0x17, 0xa8, 0x0a, 0x21, 0x0b,
	0x88, 0xb0, 0x26, 0xf6, 0xd4, 0x1d, 0x34, 0x9e, 0xb1, 0x66, 0xc6, 0x81, 0xa8, 0x2a, 0x0f, 0xc3,
	0x92, 0x47, 0xe1, 0x29, 0xba, 0xee, 0x1b, 0xc0, 0x02, 0xa1, 0xda, 0x63, 0x77, 0x42, 0x15, 0x28,
	0xbb, 0x9b, 0xdc, 0x73, 0xce, 0x3d, 0x3e, 0x57, 0x77, 0xdc, 0x9d, 0x54, 0x88, 0x94, 0x11, 0x88,
	0x13, 0x05, 0xab, 0xf2, 0xa2, 0x9a, 0x04, 0x50, 0x12, 0x25, 0x0a, 0x19, 0x13, 0x05, 0x0b, 0x45,
	0x64, 0x44, 0xb9, 0x26, 0x92, 0x28, 0x0d, 0x72, 0x29, 0xb4, 0xf0, 0x36, 0x2b, 0x2c, 0xc0, 0x89,
	0x02, 0x0d, 0x0d, 0x4c, 0x02, 0xd0, 0xd0, 0xd6, 0x5f, 0xcc, 0x53, 0x8e, 0x45, 0x96, 0x09, 0x0e,
	0x63, 0x49, 0x35, 0x91, 0x54, 0xf0, 0x28, 0xc6, 0x9a, 0xa4, 0x42, 0x4e, 0x23, 0x3c, 0xc1, 0x94,
	0xe1, 0x31, 0x65, 0x54, 0x4f, 0xab, 0x41, 0xeb, 0xcf, 0xe7, 0xa9, 0x10, 0x5e, 0x64, 0xbf, 0x79,
	0x8b, 0x34, 0xfe, 0x24, 0xb8, 0xc8, 0xa6, 0x91, 0x9e, 0xe6, 0xc4, 0x08, 0xdc, 0xaf, 0x05, 0x72,
	0x0a, 0x8f, 0x28, 0x61, 0x49, 0x34, 0x26, 0xc7, 0x78, 0x42, 0x85, 0x34, 0x80, 0x3b, 0x16, 0xa0,
	0x76, 0x6f, 0x5a, 0xf7, 0x4c, 0xab, 0xfc, 0x35, 0x2e, 0x8e, 0xe0, 0x47, 0x89, 0xf3, 0x9c, 0x48,
	0x65, 0xfa, 0x1b, 0x16, 0x15, 0x73, 0x2e, 0x34, 0xd6, 0x54, 0x70, 0xd3, 0x7d, 0xf0, 0x7d, 0xd1,
	0x5d, 0x79, 0xa3, 0x88, 0xdc, 0x37, 0xf6, 0xbc, 0xbe, 0xbb, 0x5a, 0x0f, 0x88, 0x38, 0xce, 0x88,
	0xdf, 0xea, 0xb4, 0xba, 0xff, 0xf7, 0xb6, 0xcf, 0xd0, 0xe2, 0x0f, 0xf4, 0xd0, 0xdd, 0xba, 0x0c,
	0xd2, 0x54, 0x39, 0x55, 0x20, 0x16, 0x19, 0xb4, 0x55, 0xfa, 0x2b, 0xb5, 0xc6, 0x2b, 0x9c, 0x11,
	0xef, 0xb3, 0xbb, 0x3a, 0xf3, 0xd5, 0x7e, 0xbb, 0xd3, 0xea, 0xae, 0x05, 0x03, 0x30, 0x6f, 0x41,
	0x65, 0x6e, 0xc0, 0x56, 0x1c, 0x18, 0xfe, 0x60, 0x9a, 0x93, 0x97, 0xbc, 0xc8, 0xe6, 0x36, 0x7b,
	0xce, 0x19, 0x72, 0xfa, 0x2b, 0xda, 0xfa, 0xcb, 0x3b, 0x70, 0x6f, 0xce, 0xee, 0x80, 0x26, 0xbe,
	0xd3, 0x69, 0x75, 0x97, 0x83, 0xbb, 0xb5, 0x85, 0x3a, 0x3d, 0xb0, 0xcf, 0xf5, 0x93, 0xc7, 0x43,
	0xcc, 0x0a, 0xa3, 0xb4, 0x56, 0x58, 0x83, 0xf6, 0x13, 0x6f, 0xc7, 0x5d, 0x28, 0x63, 0x59, 0x28,
	0xf9, 0x1b, 0x57, 0xf8, 0xaf, 0xb5, 0xa4, 0x3c, 0xb5, 0x04, 0x4a, 0xb8, 0x77, 0xea, 0xde, 0x9e,
	0xb5, 0x90, 0x63, 0x49, 0xb8, 0xf6, 0x17, 0xaf, 0x21, 0xb3, 0x7d, 0x86, 0x9c, 0x7f, 0xc8, 0xde,
	0xb3, 0x1d, 0x1f, 0x96, 0x63, 0xbc, 0x3d, 0xf7, 0x06, 0xc3, 0x05, 0x8f, 0x8f, 0x49, 0x12, 0x69,
	0x11, 0x61, 0xc6, 0xfc, 0xa5, 0x72, 0xf2, 0xfa, 0x95, 0xc9, 0x3d, 0x21, 0x98, 0x65, 0x7f, 0xb5,
	0xe6, 0x0d, 0x04, 0x62, 0xcc, 0xfb, 0xe0, 0xae, 0x59, 0x07, 0x40, 0x89, 0xf2, 0xff, 0xeb, 0x38,
	0xdd, 0xe5, 0xe0, 0xd9, 0xdc, 0x5d, 0x56, 0x97, 0x04, 0x76, 0xeb, 0x4b, 0xda, 0x35, 0x87, 0x84,
	0xac, 0x3b, 0x32, 0x51, 0xcf, 0x2a, 0x87, 0xef, 0xcf, 0xd1, 0xbb, 0x6b, 0x7e, 0xb4, 0x17, 0xc4,
	0x85, 0xd2, 0x22, 0x23, 0x52, 0xc1, 0x93, 0xba, 0x3c, 0x85, 0x76, 0x12, 0x0a, 0x9e, 0xcc, 0x6c,
	0xe0, 0xb4, 0xf7, 0xb3, 0xe5, 0x6e, 0xc5, 0x22, 0x03, 0x7f, 0x7d, 0x26, 0x7a, 0xb7, 0xec, 0x59,
	0x87, 0x17, 0x49, 0x1d, 0xb6, 0xde, 0x1e, 0x18, 0x5e, 0x2a, 0x18, 0xe6, 0x29, 0x10, 0x32, 0x85,
	0x29, 0xe1, 0x65, 0x8e, 0xf0, 0xd2, 0xea, 0x1f, 0x1e, 0xad, 0xa7, 0x4d, 0xf5, 0xa5, 0xed, 0xec,
	0x21, 0xf4, 0xb5, 0xbd, 0xb9, 0x57, 0x49, 0xa2, 0x44, 0x81, 0xaa, 0xbc, 0xa8, 0x86, 0x01, 0xe8,
	0xd7, 0xc8, 0x6f, 0x35, 0x66, 0x84, 0x12, 0x35, 0x6a, 0x30, 0xa3, 0x61, 0x30, 0x6a, 0x30, 0xe7,
	0xed, 0xad, 0xaa, 0x11, 0x86, 0x28, 0x51, 0x61, 0xd8, 0xa0, 0xc2, 0x70, 0x18, 0x84, 0x61, 0x83,
	0x1b, 0x2f, 0x95, 0x66, 0x1f, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x77, 0x49, 0xd1, 0x8f, 0x60,
	0x05, 0x00, 0x00,
}