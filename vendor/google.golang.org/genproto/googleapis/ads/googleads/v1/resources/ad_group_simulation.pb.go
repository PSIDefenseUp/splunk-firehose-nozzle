// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_simulation.proto

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

// An ad group simulation. Supported combinations of advertising
// channel type, simulation type and simulation modification method is
// detailed below respectively.
//
// 1. SEARCH - CPC_BID - DEFAULT
// 2. SEARCH - CPC_BID - UNIFORM
// 3. SEARCH - TARGET_CPA - UNIFORM
// 4. DISPLAY - CPC_BID - DEFAULT
// 5. DISPLAY - CPC_BID - UNIFORM
// 6. DISPLAY - TARGET_CPA - UNIFORM
// 7. VIDEO - CPV_BID - DEFAULT
// 8. VIDEO - CPV_BID - UNIFORM
type AdGroupSimulation struct {
	// Immutable. The resource name of the ad group simulation.
	// Ad group simulation resource names have the form:
	//
	// `customers/{customer_id}/adGroupSimulations/{ad_group_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Ad group id of the simulation.
	AdGroupId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=ad_group_id,json=adGroupId,proto3" json:"ad_group_id,omitempty"`
	// Output only. The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// Output only. How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,4,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v1.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// Output only. First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate *wrappers.StringValue `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Output only. Last day on which the simulation is based, in YYYY-MM-DD format
	EndDate *wrappers.StringValue `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are valid to be assigned to PointList:
	//	*AdGroupSimulation_CpcBidPointList
	//	*AdGroupSimulation_CpvBidPointList
	//	*AdGroupSimulation_TargetCpaPointList
	PointList            isAdGroupSimulation_PointList `protobuf_oneof:"point_list"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AdGroupSimulation) Reset()         { *m = AdGroupSimulation{} }
func (m *AdGroupSimulation) String() string { return proto.CompactTextString(m) }
func (*AdGroupSimulation) ProtoMessage()    {}
func (*AdGroupSimulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c8573d5f5d3a023, []int{0}
}

func (m *AdGroupSimulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupSimulation.Unmarshal(m, b)
}
func (m *AdGroupSimulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupSimulation.Marshal(b, m, deterministic)
}
func (m *AdGroupSimulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupSimulation.Merge(m, src)
}
func (m *AdGroupSimulation) XXX_Size() int {
	return xxx_messageInfo_AdGroupSimulation.Size(m)
}
func (m *AdGroupSimulation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupSimulation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupSimulation proto.InternalMessageInfo

func (m *AdGroupSimulation) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupSimulation) GetAdGroupId() *wrappers.Int64Value {
	if m != nil {
		return m.AdGroupId
	}
	return nil
}

func (m *AdGroupSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if m != nil {
		return m.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (m *AdGroupSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if m != nil {
		return m.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (m *AdGroupSimulation) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *AdGroupSimulation) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type isAdGroupSimulation_PointList interface {
	isAdGroupSimulation_PointList()
}

type AdGroupSimulation_CpcBidPointList struct {
	CpcBidPointList *common.CpcBidSimulationPointList `protobuf:"bytes,8,opt,name=cpc_bid_point_list,json=cpcBidPointList,proto3,oneof"`
}

type AdGroupSimulation_CpvBidPointList struct {
	CpvBidPointList *common.CpvBidSimulationPointList `protobuf:"bytes,10,opt,name=cpv_bid_point_list,json=cpvBidPointList,proto3,oneof"`
}

type AdGroupSimulation_TargetCpaPointList struct {
	TargetCpaPointList *common.TargetCpaSimulationPointList `protobuf:"bytes,9,opt,name=target_cpa_point_list,json=targetCpaPointList,proto3,oneof"`
}

func (*AdGroupSimulation_CpcBidPointList) isAdGroupSimulation_PointList() {}

func (*AdGroupSimulation_CpvBidPointList) isAdGroupSimulation_PointList() {}

func (*AdGroupSimulation_TargetCpaPointList) isAdGroupSimulation_PointList() {}

func (m *AdGroupSimulation) GetPointList() isAdGroupSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *AdGroupSimulation) GetCpcBidPointList() *common.CpcBidSimulationPointList {
	if x, ok := m.GetPointList().(*AdGroupSimulation_CpcBidPointList); ok {
		return x.CpcBidPointList
	}
	return nil
}

func (m *AdGroupSimulation) GetCpvBidPointList() *common.CpvBidSimulationPointList {
	if x, ok := m.GetPointList().(*AdGroupSimulation_CpvBidPointList); ok {
		return x.CpvBidPointList
	}
	return nil
}

func (m *AdGroupSimulation) GetTargetCpaPointList() *common.TargetCpaSimulationPointList {
	if x, ok := m.GetPointList().(*AdGroupSimulation_TargetCpaPointList); ok {
		return x.TargetCpaPointList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupSimulation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupSimulation_CpcBidPointList)(nil),
		(*AdGroupSimulation_CpvBidPointList)(nil),
		(*AdGroupSimulation_TargetCpaPointList)(nil),
	}
}

func init() {
	proto.RegisterType((*AdGroupSimulation)(nil), "google.ads.googleads.v1.resources.AdGroupSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_simulation.proto", fileDescriptor_8c8573d5f5d3a023)
}

var fileDescriptor_8c8573d5f5d3a023 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6a, 0x13, 0x41,
	0x18, 0x36, 0x49, 0x8f, 0xd3, 0xaa, 0xb8, 0xa2, 0xc4, 0x5a, 0xb4, 0x15, 0x0a, 0x45, 0x64, 0x86,
	0xb4, 0x22, 0xb8, 0xad, 0xd0, 0x4d, 0x95, 0x5a, 0xb1, 0x52, 0xd2, 0x12, 0x51, 0x02, 0xcb, 0x64,
	0x67, 0xba, 0x1d, 0xcc, 0x1c, 0xd8, 0x99, 0x5d, 0x29, 0xa5, 0x57, 0xbe, 0x81, 0x8f, 0xe0, 0xa5,
	0x8f, 0xe2, 0x53, 0xf4, 0xba, 0x97, 0x5e, 0x7a, 0x25, 0x99, 0x3d, 0x24, 0x26, 0x4d, 0x1b, 0xef,
	0xfe, 0xdd, 0xff, 0x3b, 0xfc, 0xff, 0xb7, 0xb3, 0x03, 0x36, 0x42, 0x29, 0xc3, 0x0e, 0x45, 0x98,
	0x68, 0x94, 0x96, 0xdd, 0x2a, 0xa9, 0xa1, 0x88, 0x6a, 0x19, 0x47, 0x01, 0xd5, 0x08, 0x13, 0x3f,
	0x8c, 0x64, 0xac, 0x7c, 0xcd, 0x78, 0xdc, 0xc1, 0x86, 0x49, 0x01, 0x55, 0x24, 0x8d, 0x74, 0x96,
	0x53, 0x06, 0xc4, 0x44, 0xc3, 0x82, 0x0c, 0x93, 0x1a, 0x2c, 0xc8, 0x0b, 0x68, 0x94, 0x7e, 0x20,
	0x39, 0x97, 0x02, 0x0d, 0x6a, 0x2e, 0xd4, 0x47, 0x11, 0xa8, 0x88, 0xb9, 0xee, 0xc3, 0xfb, 0x5c,
	0x12, 0x76, 0xc4, 0x82, 0xec, 0x81, 0x9a, 0x63, 0x49, 0x32, 0x8d, 0xf5, 0xb1, 0x35, 0xcc, 0x89,
	0xa2, 0x19, 0xe9, 0x71, 0x4e, 0x52, 0x0c, 0x1d, 0x31, 0xda, 0x21, 0x7e, 0x9b, 0x1e, 0xe3, 0x84,
	0xc9, 0x28, 0x03, 0x3c, 0xe8, 0x03, 0xe4, 0x0b, 0x66, 0xad, 0x47, 0x59, 0xcb, 0x3e, 0xb5, 0xe3,
	0x23, 0xf4, 0x35, 0xc2, 0x4a, 0xd1, 0x48, 0x67, 0xfd, 0xc5, 0x3e, 0x2a, 0x16, 0x42, 0x1a, 0xeb,
	0x9e, 0x75, 0x9f, 0xfc, 0x9e, 0x06, 0x77, 0x3c, 0xb2, 0xd3, 0xcd, 0xf8, 0xa0, 0x18, 0xcd, 0xf9,
	0x08, 0x6e, 0xe6, 0x2e, 0xbe, 0xc0, 0x9c, 0x56, 0x4b, 0x4b, 0xa5, 0xd5, 0xd9, 0xfa, 0xda, 0xb9,
	0x37, 0xf9, 0xc7, 0x7b, 0x06, 0x9e, 0xf6, 0x02, 0xcf, 0x2a, 0xc5, 0x34, 0x0c, 0x24, 0x47, 0x43,
	0x52, 0x8d, 0xf9, 0x5c, 0xe8, 0x03, 0xe6, 0xd4, 0xd9, 0x02, 0x73, 0xc5, 0x27, 0x65, 0xa4, 0x5a,
	0x5e, 0x2a, 0xad, 0xce, 0xad, 0x3d, 0xcc, 0x54, 0x60, 0xbe, 0x02, 0xdc, 0x15, 0xe6, 0xc5, 0xf3,
	0x26, 0xee, 0xc4, 0xb4, 0x5e, 0x39, 0xf7, 0x2a, 0x8d, 0x59, 0x9c, 0xea, 0xee, 0x12, 0xe7, 0x13,
	0x98, 0xe8, 0x06, 0x57, 0xad, 0x2c, 0x95, 0x56, 0x6f, 0xad, 0x6d, 0xc1, 0x51, 0xc7, 0xc0, 0xc6,
	0x0d, 0x7b, 0x83, 0x1c, 0x9e, 0x28, 0xfa, 0x46, 0xc4, 0x7c, 0xe0, 0x55, 0xaa, 0x6f, 0x25, 0x9d,
	0xef, 0x25, 0x70, 0xf7, 0x92, 0x0f, 0x5b, 0x9d, 0xb0, 0x56, 0xad, 0xb1, 0xad, 0xf6, 0xfa, 0x34,
	0xf6, 0xac, 0xc4, 0x80, 0xf1, 0x30, 0x20, 0x1d, 0xc3, 0xe1, 0x43, 0x0d, 0x67, 0x0b, 0x00, 0x6d,
	0x70, 0x64, 0x7c, 0x82, 0x0d, 0xad, 0x4e, 0xda, 0xc0, 0x16, 0x87, 0x02, 0x3b, 0x30, 0x11, 0x13,
	0x61, 0x7f, 0x62, 0x96, 0xf4, 0x1a, 0x1b, 0xea, 0x6c, 0x82, 0x19, 0x2a, 0x48, 0xca, 0x9f, 0x1a,
	0x97, 0x3f, 0x4d, 0x05, 0xb1, 0x6c, 0x0e, 0x9c, 0x40, 0x05, 0x7e, 0x9b, 0x11, 0x5f, 0x49, 0x26,
	0x8c, 0xdf, 0x61, 0xda, 0x54, 0x67, 0xac, 0xce, 0xcb, 0x91, 0x91, 0xa4, 0x7f, 0x18, 0xdc, 0x56,
	0x41, 0x9d, 0x91, 0xde, 0xe2, 0xfb, 0x5d, 0x85, 0xf7, 0x4c, 0x1b, 0x6b, 0xf2, 0xf6, 0x46, 0xe3,
	0x76, 0x60, 0x11, 0xc5, 0xfb, 0xd4, 0x2e, 0x19, 0xb4, 0x03, 0xe3, 0xda, 0x25, 0xd7, 0xda, 0x25,
	0xff, 0xd8, 0x25, 0xe0, 0x9e, 0xc1, 0x51, 0x48, 0x8d, 0x1f, 0x28, 0xdc, 0xef, 0x38, 0x6b, 0x1d,
	0x37, 0xaf, 0x73, 0x3c, 0xb4, 0xe4, 0x6d, 0x85, 0xaf, 0x30, 0x75, 0x4c, 0x0e, 0x2a, 0x5a, 0xae,
	0xb8, 0xf0, 0xbe, 0xfc, 0xcf, 0x6f, 0xe4, 0xbc, 0x0a, 0x62, 0x6d, 0x24, 0xa7, 0x91, 0x46, 0xa7,
	0x79, 0x79, 0x86, 0xf0, 0x20, 0x4e, 0xa3, 0xd3, 0x4b, 0xae, 0xcc, 0xb3, 0xfa, 0x3c, 0x00, 0xbd,
	0xe5, 0xea, 0xdf, 0xca, 0x60, 0x25, 0x90, 0x1c, 0x5e, 0x7b, 0x85, 0xd6, 0xef, 0x0f, 0x4d, 0xb2,
	0xdf, 0x3d, 0x32, 0xfb, 0xa5, 0xcf, 0xef, 0x32, 0x72, 0x28, 0x3b, 0x58, 0x84, 0x50, 0x46, 0x21,
	0x0a, 0xa9, 0xb0, 0x07, 0x0a, 0xf5, 0xb6, 0xb9, 0xe2, 0x6e, 0xdf, 0x28, 0xaa, 0x1f, 0xe5, 0xca,
	0x8e, 0xe7, 0xfd, 0x2c, 0x2f, 0xef, 0xa4, 0x92, 0x1e, 0xd1, 0x30, 0x2d, 0xbb, 0x55, 0xb3, 0x06,
	0x1b, 0x39, 0xf2, 0x57, 0x8e, 0x69, 0x79, 0x44, 0xb7, 0x0a, 0x4c, 0xab, 0x59, 0x6b, 0x15, 0x98,
	0x8b, 0xf2, 0x4a, 0xda, 0x70, 0x5d, 0x8f, 0x68, 0xd7, 0x2d, 0x50, 0xae, 0xdb, 0xac, 0xb9, 0x6e,
	0x81, 0x6b, 0x4f, 0xd9, 0x61, 0xd7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x02, 0xd7, 0xe9, 0x1f,
	0x87, 0x06, 0x00, 0x00,
}