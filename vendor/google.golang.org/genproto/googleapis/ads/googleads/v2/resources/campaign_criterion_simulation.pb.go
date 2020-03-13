// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_criterion_simulation.proto

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

// A campaign criterion simulation. Supported combinations of advertising
// channel type, criterion ids, simulation type and simulation modification
// method is detailed below respectively.
//
// 1. SEARCH - 30000,30001,30002 - BID_MODIFIER - UNIFORM
// 2. SHOPPING - 30000,30001,30002 - BID_MODIFIER - UNIFORM
// 3. DISPLAY - 30001 - BID_MODIFIER - UNIFORM
type CampaignCriterionSimulation struct {
	// Immutable. The resource name of the campaign criterion simulation.
	// Campaign criterion simulation resource names have the form:
	//
	// `customers/{customer_id}/campaignCriterionSimulations/{campaign_id}~{criterion_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Campaign ID of the simulation.
	CampaignId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	// Output only. Criterion ID of the simulation.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// Output only. The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// Output only. How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,5,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v2.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// Output only. First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate *wrappers.StringValue `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Output only. Last day on which the simulation is based, in YYYY-MM-DD format.
	EndDate *wrappers.StringValue `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are valid to be assigned to PointList:
	//	*CampaignCriterionSimulation_BidModifierPointList
	PointList            isCampaignCriterionSimulation_PointList `protobuf_oneof:"point_list"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CampaignCriterionSimulation) Reset()         { *m = CampaignCriterionSimulation{} }
func (m *CampaignCriterionSimulation) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterionSimulation) ProtoMessage()    {}
func (*CampaignCriterionSimulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5832e91cef223b5, []int{0}
}

func (m *CampaignCriterionSimulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterionSimulation.Unmarshal(m, b)
}
func (m *CampaignCriterionSimulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterionSimulation.Marshal(b, m, deterministic)
}
func (m *CampaignCriterionSimulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterionSimulation.Merge(m, src)
}
func (m *CampaignCriterionSimulation) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterionSimulation.Size(m)
}
func (m *CampaignCriterionSimulation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterionSimulation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterionSimulation proto.InternalMessageInfo

func (m *CampaignCriterionSimulation) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterionSimulation) GetCampaignId() *wrappers.Int64Value {
	if m != nil {
		return m.CampaignId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if m != nil {
		return m.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if m != nil {
		return m.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type isCampaignCriterionSimulation_PointList interface {
	isCampaignCriterionSimulation_PointList()
}

type CampaignCriterionSimulation_BidModifierPointList struct {
	BidModifierPointList *common.BidModifierSimulationPointList `protobuf:"bytes,8,opt,name=bid_modifier_point_list,json=bidModifierPointList,proto3,oneof"`
}

func (*CampaignCriterionSimulation_BidModifierPointList) isCampaignCriterionSimulation_PointList() {}

func (m *CampaignCriterionSimulation) GetPointList() isCampaignCriterionSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetBidModifierPointList() *common.BidModifierSimulationPointList {
	if x, ok := m.GetPointList().(*CampaignCriterionSimulation_BidModifierPointList); ok {
		return x.BidModifierPointList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterionSimulation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterionSimulation_BidModifierPointList)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignCriterionSimulation)(nil), "google.ads.googleads.v2.resources.CampaignCriterionSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_criterion_simulation.proto", fileDescriptor_a5832e91cef223b5)
}

var fileDescriptor_a5832e91cef223b5 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x4e, 0xd4, 0x40,
	0x18, 0xb7, 0xbb, 0xfc, 0x1d, 0xd0, 0x43, 0x35, 0xb1, 0x02, 0xd1, 0xc5, 0x84, 0x84, 0xd3, 0x34,
	0x29, 0x84, 0x43, 0x25, 0x86, 0x16, 0x09, 0x62, 0xc4, 0xe0, 0x62, 0x36, 0xd1, 0x6c, 0xd2, 0xcc,
	0x76, 0x86, 0x32, 0xc9, 0x76, 0xa6, 0x99, 0x99, 0x62, 0x88, 0xf2, 0x00, 0x1e, 0xb8, 0x18, 0x9f,
	0xc0, 0xa3, 0x8f, 0xe2, 0x53, 0x70, 0xe6, 0x11, 0x3c, 0x99, 0x9d, 0x4e, 0xdb, 0x0d, 0xb0, 0xeb,
	0xc6, 0xdb, 0xd7, 0x7e, 0xbf, 0x3f, 0x33, 0xbf, 0x6f, 0x66, 0xc0, 0x5e, 0xc2, 0x79, 0xd2, 0x27,
	0x2e, 0xc2, 0xd2, 0x2d, 0xca, 0x41, 0x75, 0xe6, 0xb9, 0x82, 0x48, 0x9e, 0x8b, 0x98, 0x48, 0x37,
	0x46, 0x69, 0x86, 0x68, 0xc2, 0xa2, 0x58, 0x50, 0x45, 0x04, 0xe5, 0x2c, 0x92, 0x34, 0xcd, 0xfb,
	0x48, 0x51, 0xce, 0x60, 0x26, 0xb8, 0xe2, 0xf6, 0x6a, 0xc1, 0x85, 0x08, 0x4b, 0x58, 0xc9, 0xc0,
	0x33, 0x0f, 0x56, 0x32, 0x4b, 0xee, 0x28, 0xa7, 0x98, 0xa7, 0x29, 0x67, 0xee, 0x4d, 0xcd, 0xa5,
	0x70, 0x14, 0x81, 0xb0, 0x3c, 0x95, 0x43, 0xf8, 0x28, 0xe5, 0x98, 0x9e, 0xd0, 0xd8, 0x7c, 0x10,
	0x75, 0xca, 0xb1, 0xd1, 0xd8, 0x98, 0x58, 0x43, 0x9d, 0x67, 0xc4, 0x90, 0x9e, 0x95, 0xa4, 0x8c,
	0xba, 0x27, 0x94, 0xf4, 0x71, 0xd4, 0x23, 0xa7, 0xe8, 0x8c, 0x72, 0x61, 0x00, 0x4f, 0x86, 0x00,
	0xe5, 0x06, 0x4d, 0xeb, 0xa9, 0x69, 0xe9, 0xaf, 0x5e, 0x7e, 0xe2, 0x7e, 0x16, 0x28, 0xcb, 0x88,
	0x90, 0xa6, 0xbf, 0x32, 0x44, 0x45, 0x8c, 0x71, 0xa5, 0xdd, 0x4d, 0xf7, 0xf9, 0x8f, 0x59, 0xb0,
	0xbc, 0x6b, 0xe2, 0xde, 0x2d, 0xd3, 0x3e, 0xae, 0x16, 0x69, 0x23, 0x70, 0xbf, 0xf4, 0x8b, 0x18,
	0x4a, 0x89, 0x63, 0xb5, 0xac, 0xf5, 0xf9, 0x70, 0xfb, 0x2a, 0x98, 0xfe, 0x13, 0x6c, 0x81, 0xcd,
	0x3a, 0x7a, 0x53, 0x65, 0x54, 0xc2, 0x98, 0xa7, 0xee, 0x18, 0xd1, 0xf6, 0x62, 0x29, 0xf9, 0x0e,
	0xa5, 0xc4, 0x0e, 0xc0, 0x42, 0x35, 0x70, 0x8a, 0x9d, 0x46, 0xcb, 0x5a, 0x5f, 0xf0, 0x96, 0x8d,
	0x1e, 0x2c, 0xb7, 0x05, 0x0f, 0x98, 0xda, 0xda, 0xec, 0xa0, 0x7e, 0x4e, 0xc2, 0xe6, 0x55, 0xd0,
	0x6c, 0x83, 0x92, 0x74, 0x80, 0xed, 0x5d, 0xb0, 0x58, 0x1f, 0x15, 0x8a, 0x9d, 0xe6, 0x84, 0x1a,
	0x0b, 0x15, 0xeb, 0x00, 0xdb, 0x1f, 0xc1, 0xd4, 0x60, 0x24, 0xce, 0x54, 0xcb, 0x5a, 0x7f, 0xe0,
	0xed, 0xc0, 0x51, 0x07, 0x4c, 0x0f, 0x12, 0xd6, 0xdb, 0xf9, 0x70, 0x9e, 0x91, 0x3d, 0x96, 0xa7,
	0x37, 0x7e, 0x15, 0x0e, 0x5a, 0xd2, 0xfe, 0x6e, 0x81, 0x87, 0x77, 0x1c, 0x19, 0x67, 0x5a, 0x5b,
	0x75, 0x27, 0xb6, 0x3a, 0x1c, 0xd2, 0x38, 0xd4, 0x12, 0x37, 0x8c, 0x6f, 0x03, 0x8a, 0x65, 0xd8,
	0xe9, 0xad, 0x86, 0xbd, 0x03, 0x80, 0x54, 0x48, 0xa8, 0x08, 0x23, 0x45, 0x9c, 0x19, 0x1d, 0xd9,
	0xca, 0xad, 0xc8, 0x8e, 0x95, 0xa0, 0x2c, 0x19, 0xca, 0x6c, 0x5e, 0x93, 0x5e, 0x21, 0x45, 0xec,
	0x6d, 0x30, 0x47, 0x18, 0x2e, 0xf8, 0xb3, 0x93, 0xf2, 0x67, 0x09, 0xc3, 0x9a, 0xfd, 0x15, 0x3c,
	0xee, 0x51, 0x6c, 0xae, 0x12, 0x11, 0x51, 0xc6, 0x29, 0x53, 0x51, 0x9f, 0x4a, 0xe5, 0xcc, 0x69,
	0xb1, 0x97, 0x23, 0x73, 0x29, 0x2e, 0x30, 0x0c, 0x29, 0x3e, 0x34, 0xec, 0x3a, 0x82, 0xa3, 0x81,
	0xcc, 0x5b, 0x2a, 0x95, 0xb6, 0x7b, 0x7d, 0xaf, 0xfd, 0xa8, 0x57, 0xc3, 0xaa, 0xa6, 0x7f, 0x69,
	0x5d, 0x07, 0xdf, 0xac, 0xff, 0x3b, 0xc0, 0xf6, 0xfb, 0x38, 0x97, 0x8a, 0xa7, 0x44, 0x48, 0xf7,
	0x4b, 0x59, 0x5e, 0x54, 0xcf, 0xd6, 0x1d, 0x8c, 0x01, 0x6e, 0xdc, 0xa3, 0x76, 0x11, 0x2e, 0x02,
	0x50, 0x07, 0x10, 0x5e, 0x36, 0xc0, 0x5a, 0xcc, 0x53, 0xf8, 0xcf, 0x47, 0x2e, 0x6c, 0x8d, 0x59,
	0xe7, 0xd1, 0x60, 0x08, 0x47, 0xd6, 0xa7, 0x37, 0x46, 0x26, 0xe1, 0x7d, 0xc4, 0x12, 0xc8, 0x45,
	0xe2, 0x26, 0x84, 0xe9, 0x11, 0xb9, 0xf5, 0xae, 0xc7, 0xbc, 0xc8, 0x2f, 0xaa, 0xea, 0x67, 0xa3,
	0xb9, 0x1f, 0x04, 0xbf, 0x1a, 0xab, 0xfb, 0x85, 0x64, 0x80, 0x25, 0x2c, 0xca, 0x41, 0xd5, 0xf1,
	0x60, 0xbb, 0x44, 0xfe, 0x2e, 0x31, 0xdd, 0x00, 0xcb, 0x6e, 0x85, 0xe9, 0x76, 0xbc, 0x6e, 0x85,
	0xb9, 0x6e, 0xac, 0x15, 0x0d, 0xdf, 0x0f, 0xb0, 0xf4, 0xfd, 0x0a, 0xe5, 0xfb, 0x1d, 0xcf, 0xf7,
	0x2b, 0x5c, 0x6f, 0x46, 0x2f, 0x76, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0xa8, 0x39,
	0xfc, 0x3d, 0x06, 0x00, 0x00,
}