// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/account_budget.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// An account-level budget. It contains information about the budget itself,
// as well as the most recently approved changes to the budget and proposed
// changes that are pending approval. The proposed changes that are pending
// approval, if any, are found in 'pending_proposal'.  Effective details about
// the budget are found in fields prefixed 'approved_', 'adjusted_' and those
// without a prefix.  Since some effective details may differ from what the user
// had originally requested (e.g. spending limit), these differences are
// juxtaposed via 'proposed_', 'approved_', and possibly 'adjusted_' fields.
//
// This resource is mutated using AccountBudgetProposal and cannot be mutated
// directly. A budget may have at most one pending proposal at any given time.
// It is read through pending_proposal.
//
// Once approved, a budget may be subject to adjustments, such as credit
// adjustments.  Adjustments create differences between the 'approved' and
// 'adjusted' fields, which would otherwise be identical.
type AccountBudget struct {
	// Immutable. The resource name of the account-level budget.
	// AccountBudget resource names have the form:
	//
	// `customers/{customer_id}/accountBudgets/{account_budget_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the account-level budget.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The resource name of the billing setup associated with this account-level
	// budget.  BillingSetup resource names have the form:
	//
	// `customers/{customer_id}/billingSetups/{billing_setup_id}`
	BillingSetup *wrappers.StringValue `protobuf:"bytes,3,opt,name=billing_setup,json=billingSetup,proto3" json:"billing_setup,omitempty"`
	// Output only. The status of this account-level budget.
	Status enums.AccountBudgetStatusEnum_AccountBudgetStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.AccountBudgetStatusEnum_AccountBudgetStatus" json:"status,omitempty"`
	// Output only. The name of the account-level budget.
	Name *wrappers.StringValue `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The proposed start time of the account-level budget in
	// yyyy-MM-dd HH:mm:ss format.  If a start time type of NOW was proposed,
	// this is the time of request.
	ProposedStartDateTime *wrappers.StringValue `protobuf:"bytes,6,opt,name=proposed_start_date_time,json=proposedStartDateTime,proto3" json:"proposed_start_date_time,omitempty"`
	// Output only. The approved start time of the account-level budget in yyyy-MM-dd HH:mm:ss
	// format.
	//
	// For example, if a new budget is approved after the proposed start time,
	// the approved start time is the time of approval.
	ApprovedStartDateTime *wrappers.StringValue `protobuf:"bytes,7,opt,name=approved_start_date_time,json=approvedStartDateTime,proto3" json:"approved_start_date_time,omitempty"`
	// Output only. The total adjustments amount.
	//
	// An example of an adjustment is courtesy credits.
	TotalAdjustmentsMicros *wrappers.Int64Value `protobuf:"bytes,18,opt,name=total_adjustments_micros,json=totalAdjustmentsMicros,proto3" json:"total_adjustments_micros,omitempty"`
	// Output only. The value of Ads that have been served, in micros.
	//
	// This includes overdelivery costs, in which case a credit might be
	// automatically applied to the budget (see total_adjustments_micros).
	AmountServedMicros *wrappers.Int64Value `protobuf:"bytes,19,opt,name=amount_served_micros,json=amountServedMicros,proto3" json:"amount_served_micros,omitempty"`
	// Output only. A purchase order number is a value that helps users reference this budget
	// in their monthly invoices.
	PurchaseOrderNumber *wrappers.StringValue `protobuf:"bytes,20,opt,name=purchase_order_number,json=purchaseOrderNumber,proto3" json:"purchase_order_number,omitempty"`
	// Output only. Notes associated with the budget.
	Notes *wrappers.StringValue `protobuf:"bytes,21,opt,name=notes,proto3" json:"notes,omitempty"`
	// Output only. The pending proposal to modify this budget, if applicable.
	PendingProposal *AccountBudget_PendingAccountBudgetProposal `protobuf:"bytes,22,opt,name=pending_proposal,json=pendingProposal,proto3" json:"pending_proposal,omitempty"`
	// The proposed end time of the account-level budget.
	//
	// Types that are valid to be assigned to ProposedEndTime:
	//	*AccountBudget_ProposedEndDateTime
	//	*AccountBudget_ProposedEndTimeType
	ProposedEndTime isAccountBudget_ProposedEndTime `protobuf_oneof:"proposed_end_time"`
	// The approved end time of the account-level budget.
	//
	// For example, if a budget's end time is updated and the proposal is approved
	// after the proposed end time, the approved end time is the time of approval.
	//
	// Types that are valid to be assigned to ApprovedEndTime:
	//	*AccountBudget_ApprovedEndDateTime
	//	*AccountBudget_ApprovedEndTimeType
	ApprovedEndTime isAccountBudget_ApprovedEndTime `protobuf_oneof:"approved_end_time"`
	// The proposed spending limit.
	//
	// Types that are valid to be assigned to ProposedSpendingLimit:
	//	*AccountBudget_ProposedSpendingLimitMicros
	//	*AccountBudget_ProposedSpendingLimitType
	ProposedSpendingLimit isAccountBudget_ProposedSpendingLimit `protobuf_oneof:"proposed_spending_limit"`
	// The approved spending limit.
	//
	// For example, if the amount already spent by the account exceeds the
	// proposed spending limit at the time the proposal is approved, the approved
	// spending limit is set to the amount already spent.
	//
	// Types that are valid to be assigned to ApprovedSpendingLimit:
	//	*AccountBudget_ApprovedSpendingLimitMicros
	//	*AccountBudget_ApprovedSpendingLimitType
	ApprovedSpendingLimit isAccountBudget_ApprovedSpendingLimit `protobuf_oneof:"approved_spending_limit"`
	// The spending limit after adjustments have been applied.  Adjustments are
	// stored in total_adjustments_micros.
	//
	// This value has the final say on how much the account is allowed to spend.
	//
	// Types that are valid to be assigned to AdjustedSpendingLimit:
	//	*AccountBudget_AdjustedSpendingLimitMicros
	//	*AccountBudget_AdjustedSpendingLimitType
	AdjustedSpendingLimit isAccountBudget_AdjustedSpendingLimit `protobuf_oneof:"adjusted_spending_limit"`
	XXX_NoUnkeyedLiteral  struct{}                              `json:"-"`
	XXX_unrecognized      []byte                                `json:"-"`
	XXX_sizecache         int32                                 `json:"-"`
}

func (m *AccountBudget) Reset()         { *m = AccountBudget{} }
func (m *AccountBudget) String() string { return proto.CompactTextString(m) }
func (*AccountBudget) ProtoMessage()    {}
func (*AccountBudget) Descriptor() ([]byte, []int) {
	return fileDescriptor_4320b2be24569afc, []int{0}
}

func (m *AccountBudget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudget.Unmarshal(m, b)
}
func (m *AccountBudget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudget.Marshal(b, m, deterministic)
}
func (m *AccountBudget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudget.Merge(m, src)
}
func (m *AccountBudget) XXX_Size() int {
	return xxx_messageInfo_AccountBudget.Size(m)
}
func (m *AccountBudget) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudget.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudget proto.InternalMessageInfo

func (m *AccountBudget) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AccountBudget) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AccountBudget) GetBillingSetup() *wrappers.StringValue {
	if m != nil {
		return m.BillingSetup
	}
	return nil
}

func (m *AccountBudget) GetStatus() enums.AccountBudgetStatusEnum_AccountBudgetStatus {
	if m != nil {
		return m.Status
	}
	return enums.AccountBudgetStatusEnum_UNSPECIFIED
}

func (m *AccountBudget) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AccountBudget) GetProposedStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ProposedStartDateTime
	}
	return nil
}

func (m *AccountBudget) GetApprovedStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ApprovedStartDateTime
	}
	return nil
}

func (m *AccountBudget) GetTotalAdjustmentsMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TotalAdjustmentsMicros
	}
	return nil
}

func (m *AccountBudget) GetAmountServedMicros() *wrappers.Int64Value {
	if m != nil {
		return m.AmountServedMicros
	}
	return nil
}

func (m *AccountBudget) GetPurchaseOrderNumber() *wrappers.StringValue {
	if m != nil {
		return m.PurchaseOrderNumber
	}
	return nil
}

func (m *AccountBudget) GetNotes() *wrappers.StringValue {
	if m != nil {
		return m.Notes
	}
	return nil
}

func (m *AccountBudget) GetPendingProposal() *AccountBudget_PendingAccountBudgetProposal {
	if m != nil {
		return m.PendingProposal
	}
	return nil
}

type isAccountBudget_ProposedEndTime interface {
	isAccountBudget_ProposedEndTime()
}

type AccountBudget_ProposedEndDateTime struct {
	ProposedEndDateTime *wrappers.StringValue `protobuf:"bytes,8,opt,name=proposed_end_date_time,json=proposedEndDateTime,proto3,oneof"`
}

type AccountBudget_ProposedEndTimeType struct {
	ProposedEndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,9,opt,name=proposed_end_time_type,json=proposedEndTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*AccountBudget_ProposedEndDateTime) isAccountBudget_ProposedEndTime() {}

func (*AccountBudget_ProposedEndTimeType) isAccountBudget_ProposedEndTime() {}

func (m *AccountBudget) GetProposedEndTime() isAccountBudget_ProposedEndTime {
	if m != nil {
		return m.ProposedEndTime
	}
	return nil
}

func (m *AccountBudget) GetProposedEndDateTime() *wrappers.StringValue {
	if x, ok := m.GetProposedEndTime().(*AccountBudget_ProposedEndDateTime); ok {
		return x.ProposedEndDateTime
	}
	return nil
}

func (m *AccountBudget) GetProposedEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetProposedEndTime().(*AccountBudget_ProposedEndTimeType); ok {
		return x.ProposedEndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isAccountBudget_ApprovedEndTime interface {
	isAccountBudget_ApprovedEndTime()
}

type AccountBudget_ApprovedEndDateTime struct {
	ApprovedEndDateTime *wrappers.StringValue `protobuf:"bytes,10,opt,name=approved_end_date_time,json=approvedEndDateTime,proto3,oneof"`
}

type AccountBudget_ApprovedEndTimeType struct {
	ApprovedEndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,11,opt,name=approved_end_time_type,json=approvedEndTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*AccountBudget_ApprovedEndDateTime) isAccountBudget_ApprovedEndTime() {}

func (*AccountBudget_ApprovedEndTimeType) isAccountBudget_ApprovedEndTime() {}

func (m *AccountBudget) GetApprovedEndTime() isAccountBudget_ApprovedEndTime {
	if m != nil {
		return m.ApprovedEndTime
	}
	return nil
}

func (m *AccountBudget) GetApprovedEndDateTime() *wrappers.StringValue {
	if x, ok := m.GetApprovedEndTime().(*AccountBudget_ApprovedEndDateTime); ok {
		return x.ApprovedEndDateTime
	}
	return nil
}

func (m *AccountBudget) GetApprovedEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetApprovedEndTime().(*AccountBudget_ApprovedEndTimeType); ok {
		return x.ApprovedEndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isAccountBudget_ProposedSpendingLimit interface {
	isAccountBudget_ProposedSpendingLimit()
}

type AccountBudget_ProposedSpendingLimitMicros struct {
	ProposedSpendingLimitMicros *wrappers.Int64Value `protobuf:"bytes,12,opt,name=proposed_spending_limit_micros,json=proposedSpendingLimitMicros,proto3,oneof"`
}

type AccountBudget_ProposedSpendingLimitType struct {
	ProposedSpendingLimitType enums.SpendingLimitTypeEnum_SpendingLimitType `protobuf:"varint,13,opt,name=proposed_spending_limit_type,json=proposedSpendingLimitType,proto3,enum=google.ads.googleads.v1.enums.SpendingLimitTypeEnum_SpendingLimitType,oneof"`
}

func (*AccountBudget_ProposedSpendingLimitMicros) isAccountBudget_ProposedSpendingLimit() {}

func (*AccountBudget_ProposedSpendingLimitType) isAccountBudget_ProposedSpendingLimit() {}

func (m *AccountBudget) GetProposedSpendingLimit() isAccountBudget_ProposedSpendingLimit {
	if m != nil {
		return m.ProposedSpendingLimit
	}
	return nil
}

func (m *AccountBudget) GetProposedSpendingLimitMicros() *wrappers.Int64Value {
	if x, ok := m.GetProposedSpendingLimit().(*AccountBudget_ProposedSpendingLimitMicros); ok {
		return x.ProposedSpendingLimitMicros
	}
	return nil
}

func (m *AccountBudget) GetProposedSpendingLimitType() enums.SpendingLimitTypeEnum_SpendingLimitType {
	if x, ok := m.GetProposedSpendingLimit().(*AccountBudget_ProposedSpendingLimitType); ok {
		return x.ProposedSpendingLimitType
	}
	return enums.SpendingLimitTypeEnum_UNSPECIFIED
}

type isAccountBudget_ApprovedSpendingLimit interface {
	isAccountBudget_ApprovedSpendingLimit()
}

type AccountBudget_ApprovedSpendingLimitMicros struct {
	ApprovedSpendingLimitMicros *wrappers.Int64Value `protobuf:"bytes,14,opt,name=approved_spending_limit_micros,json=approvedSpendingLimitMicros,proto3,oneof"`
}

type AccountBudget_ApprovedSpendingLimitType struct {
	ApprovedSpendingLimitType enums.SpendingLimitTypeEnum_SpendingLimitType `protobuf:"varint,15,opt,name=approved_spending_limit_type,json=approvedSpendingLimitType,proto3,enum=google.ads.googleads.v1.enums.SpendingLimitTypeEnum_SpendingLimitType,oneof"`
}

func (*AccountBudget_ApprovedSpendingLimitMicros) isAccountBudget_ApprovedSpendingLimit() {}

func (*AccountBudget_ApprovedSpendingLimitType) isAccountBudget_ApprovedSpendingLimit() {}

func (m *AccountBudget) GetApprovedSpendingLimit() isAccountBudget_ApprovedSpendingLimit {
	if m != nil {
		return m.ApprovedSpendingLimit
	}
	return nil
}

func (m *AccountBudget) GetApprovedSpendingLimitMicros() *wrappers.Int64Value {
	if x, ok := m.GetApprovedSpendingLimit().(*AccountBudget_ApprovedSpendingLimitMicros); ok {
		return x.ApprovedSpendingLimitMicros
	}
	return nil
}

func (m *AccountBudget) GetApprovedSpendingLimitType() enums.SpendingLimitTypeEnum_SpendingLimitType {
	if x, ok := m.GetApprovedSpendingLimit().(*AccountBudget_ApprovedSpendingLimitType); ok {
		return x.ApprovedSpendingLimitType
	}
	return enums.SpendingLimitTypeEnum_UNSPECIFIED
}

type isAccountBudget_AdjustedSpendingLimit interface {
	isAccountBudget_AdjustedSpendingLimit()
}

type AccountBudget_AdjustedSpendingLimitMicros struct {
	AdjustedSpendingLimitMicros *wrappers.Int64Value `protobuf:"bytes,16,opt,name=adjusted_spending_limit_micros,json=adjustedSpendingLimitMicros,proto3,oneof"`
}

type AccountBudget_AdjustedSpendingLimitType struct {
	AdjustedSpendingLimitType enums.SpendingLimitTypeEnum_SpendingLimitType `protobuf:"varint,17,opt,name=adjusted_spending_limit_type,json=adjustedSpendingLimitType,proto3,enum=google.ads.googleads.v1.enums.SpendingLimitTypeEnum_SpendingLimitType,oneof"`
}

func (*AccountBudget_AdjustedSpendingLimitMicros) isAccountBudget_AdjustedSpendingLimit() {}

func (*AccountBudget_AdjustedSpendingLimitType) isAccountBudget_AdjustedSpendingLimit() {}

func (m *AccountBudget) GetAdjustedSpendingLimit() isAccountBudget_AdjustedSpendingLimit {
	if m != nil {
		return m.AdjustedSpendingLimit
	}
	return nil
}

func (m *AccountBudget) GetAdjustedSpendingLimitMicros() *wrappers.Int64Value {
	if x, ok := m.GetAdjustedSpendingLimit().(*AccountBudget_AdjustedSpendingLimitMicros); ok {
		return x.AdjustedSpendingLimitMicros
	}
	return nil
}

func (m *AccountBudget) GetAdjustedSpendingLimitType() enums.SpendingLimitTypeEnum_SpendingLimitType {
	if x, ok := m.GetAdjustedSpendingLimit().(*AccountBudget_AdjustedSpendingLimitType); ok {
		return x.AdjustedSpendingLimitType
	}
	return enums.SpendingLimitTypeEnum_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AccountBudget) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AccountBudget_ProposedEndDateTime)(nil),
		(*AccountBudget_ProposedEndTimeType)(nil),
		(*AccountBudget_ApprovedEndDateTime)(nil),
		(*AccountBudget_ApprovedEndTimeType)(nil),
		(*AccountBudget_ProposedSpendingLimitMicros)(nil),
		(*AccountBudget_ProposedSpendingLimitType)(nil),
		(*AccountBudget_ApprovedSpendingLimitMicros)(nil),
		(*AccountBudget_ApprovedSpendingLimitType)(nil),
		(*AccountBudget_AdjustedSpendingLimitMicros)(nil),
		(*AccountBudget_AdjustedSpendingLimitType)(nil),
	}
}

// A pending proposal associated with the enclosing account-level budget,
// if applicable.
type AccountBudget_PendingAccountBudgetProposal struct {
	// Output only. The resource name of the proposal.
	// AccountBudgetProposal resource names have the form:
	//
	// `customers/{customer_id}/accountBudgetProposals/{account_budget_proposal_id}`
	AccountBudgetProposal *wrappers.StringValue `protobuf:"bytes,1,opt,name=account_budget_proposal,json=accountBudgetProposal,proto3" json:"account_budget_proposal,omitempty"`
	// Output only. The type of this proposal, e.g. END to end the budget associated
	// with this proposal.
	ProposalType enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType `protobuf:"varint,2,opt,name=proposal_type,json=proposalType,proto3,enum=google.ads.googleads.v1.enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType" json:"proposal_type,omitempty"`
	// Output only. The name to assign to the account-level budget.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The start time in yyyy-MM-dd HH:mm:ss format.
	StartDateTime *wrappers.StringValue `protobuf:"bytes,4,opt,name=start_date_time,json=startDateTime,proto3" json:"start_date_time,omitempty"`
	// Output only. A purchase order number is a value that helps users reference this budget
	// in their monthly invoices.
	PurchaseOrderNumber *wrappers.StringValue `protobuf:"bytes,9,opt,name=purchase_order_number,json=purchaseOrderNumber,proto3" json:"purchase_order_number,omitempty"`
	// Output only. Notes associated with this budget.
	Notes *wrappers.StringValue `protobuf:"bytes,10,opt,name=notes,proto3" json:"notes,omitempty"`
	// Output only. The time when this account-level budget proposal was created.
	// Formatted as yyyy-MM-dd HH:mm:ss.
	CreationDateTime *wrappers.StringValue `protobuf:"bytes,11,opt,name=creation_date_time,json=creationDateTime,proto3" json:"creation_date_time,omitempty"`
	// The end time of the account-level budget.
	//
	// Types that are valid to be assigned to EndTime:
	//	*AccountBudget_PendingAccountBudgetProposal_EndDateTime
	//	*AccountBudget_PendingAccountBudgetProposal_EndTimeType
	EndTime isAccountBudget_PendingAccountBudgetProposal_EndTime `protobuf_oneof:"end_time"`
	// The spending limit.
	//
	// Types that are valid to be assigned to SpendingLimit:
	//	*AccountBudget_PendingAccountBudgetProposal_SpendingLimitMicros
	//	*AccountBudget_PendingAccountBudgetProposal_SpendingLimitType
	SpendingLimit        isAccountBudget_PendingAccountBudgetProposal_SpendingLimit `protobuf_oneof:"spending_limit"`
	XXX_NoUnkeyedLiteral struct{}                                                   `json:"-"`
	XXX_unrecognized     []byte                                                     `json:"-"`
	XXX_sizecache        int32                                                      `json:"-"`
}

func (m *AccountBudget_PendingAccountBudgetProposal) Reset() {
	*m = AccountBudget_PendingAccountBudgetProposal{}
}
func (m *AccountBudget_PendingAccountBudgetProposal) String() string {
	return proto.CompactTextString(m)
}
func (*AccountBudget_PendingAccountBudgetProposal) ProtoMessage() {}
func (*AccountBudget_PendingAccountBudgetProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4320b2be24569afc, []int{0, 0}
}

func (m *AccountBudget_PendingAccountBudgetProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudget_PendingAccountBudgetProposal.Unmarshal(m, b)
}
func (m *AccountBudget_PendingAccountBudgetProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudget_PendingAccountBudgetProposal.Marshal(b, m, deterministic)
}
func (m *AccountBudget_PendingAccountBudgetProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudget_PendingAccountBudgetProposal.Merge(m, src)
}
func (m *AccountBudget_PendingAccountBudgetProposal) XXX_Size() int {
	return xxx_messageInfo_AccountBudget_PendingAccountBudgetProposal.Size(m)
}
func (m *AccountBudget_PendingAccountBudgetProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudget_PendingAccountBudgetProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudget_PendingAccountBudgetProposal proto.InternalMessageInfo

func (m *AccountBudget_PendingAccountBudgetProposal) GetAccountBudgetProposal() *wrappers.StringValue {
	if m != nil {
		return m.AccountBudgetProposal
	}
	return nil
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetProposalType() enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType {
	if m != nil {
		return m.ProposalType
	}
	return enums.AccountBudgetProposalTypeEnum_UNSPECIFIED
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.StartDateTime
	}
	return nil
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetPurchaseOrderNumber() *wrappers.StringValue {
	if m != nil {
		return m.PurchaseOrderNumber
	}
	return nil
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetNotes() *wrappers.StringValue {
	if m != nil {
		return m.Notes
	}
	return nil
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetCreationDateTime() *wrappers.StringValue {
	if m != nil {
		return m.CreationDateTime
	}
	return nil
}

type isAccountBudget_PendingAccountBudgetProposal_EndTime interface {
	isAccountBudget_PendingAccountBudgetProposal_EndTime()
}

type AccountBudget_PendingAccountBudgetProposal_EndDateTime struct {
	EndDateTime *wrappers.StringValue `protobuf:"bytes,5,opt,name=end_date_time,json=endDateTime,proto3,oneof"`
}

type AccountBudget_PendingAccountBudgetProposal_EndTimeType struct {
	EndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,6,opt,name=end_time_type,json=endTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*AccountBudget_PendingAccountBudgetProposal_EndDateTime) isAccountBudget_PendingAccountBudgetProposal_EndTime() {
}

func (*AccountBudget_PendingAccountBudgetProposal_EndTimeType) isAccountBudget_PendingAccountBudgetProposal_EndTime() {
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetEndTime() isAccountBudget_PendingAccountBudgetProposal_EndTime {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetEndDateTime() *wrappers.StringValue {
	if x, ok := m.GetEndTime().(*AccountBudget_PendingAccountBudgetProposal_EndDateTime); ok {
		return x.EndDateTime
	}
	return nil
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetEndTime().(*AccountBudget_PendingAccountBudgetProposal_EndTimeType); ok {
		return x.EndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isAccountBudget_PendingAccountBudgetProposal_SpendingLimit interface {
	isAccountBudget_PendingAccountBudgetProposal_SpendingLimit()
}

type AccountBudget_PendingAccountBudgetProposal_SpendingLimitMicros struct {
	SpendingLimitMicros *wrappers.Int64Value `protobuf:"bytes,7,opt,name=spending_limit_micros,json=spendingLimitMicros,proto3,oneof"`
}

type AccountBudget_PendingAccountBudgetProposal_SpendingLimitType struct {
	SpendingLimitType enums.SpendingLimitTypeEnum_SpendingLimitType `protobuf:"varint,8,opt,name=spending_limit_type,json=spendingLimitType,proto3,enum=google.ads.googleads.v1.enums.SpendingLimitTypeEnum_SpendingLimitType,oneof"`
}

func (*AccountBudget_PendingAccountBudgetProposal_SpendingLimitMicros) isAccountBudget_PendingAccountBudgetProposal_SpendingLimit() {
}

func (*AccountBudget_PendingAccountBudgetProposal_SpendingLimitType) isAccountBudget_PendingAccountBudgetProposal_SpendingLimit() {
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetSpendingLimit() isAccountBudget_PendingAccountBudgetProposal_SpendingLimit {
	if m != nil {
		return m.SpendingLimit
	}
	return nil
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetSpendingLimitMicros() *wrappers.Int64Value {
	if x, ok := m.GetSpendingLimit().(*AccountBudget_PendingAccountBudgetProposal_SpendingLimitMicros); ok {
		return x.SpendingLimitMicros
	}
	return nil
}

func (m *AccountBudget_PendingAccountBudgetProposal) GetSpendingLimitType() enums.SpendingLimitTypeEnum_SpendingLimitType {
	if x, ok := m.GetSpendingLimit().(*AccountBudget_PendingAccountBudgetProposal_SpendingLimitType); ok {
		return x.SpendingLimitType
	}
	return enums.SpendingLimitTypeEnum_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AccountBudget_PendingAccountBudgetProposal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AccountBudget_PendingAccountBudgetProposal_EndDateTime)(nil),
		(*AccountBudget_PendingAccountBudgetProposal_EndTimeType)(nil),
		(*AccountBudget_PendingAccountBudgetProposal_SpendingLimitMicros)(nil),
		(*AccountBudget_PendingAccountBudgetProposal_SpendingLimitType)(nil),
	}
}

func init() {
	proto.RegisterType((*AccountBudget)(nil), "google.ads.googleads.v1.resources.AccountBudget")
	proto.RegisterType((*AccountBudget_PendingAccountBudgetProposal)(nil), "google.ads.googleads.v1.resources.AccountBudget.PendingAccountBudgetProposal")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/account_budget.proto", fileDescriptor_4320b2be24569afc)
}

var fileDescriptor_4320b2be24569afc = []byte{
	// 1132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0xec, 0x24, 0x4d, 0x98, 0x38, 0x3f, 0x72, 0x9d, 0xca, 0x71, 0xd0, 0xa5, 0x03, 0xb2,
	0xe5, 0xa6, 0x52, 0x9d, 0x65, 0x29, 0xe6, 0xdd, 0x4c, 0xc6, 0xba, 0xb4, 0xc5, 0xda, 0x06, 0x72,
	0x67, 0x0c, 0x45, 0x30, 0x81, 0xb6, 0x58, 0x55, 0x83, 0x25, 0x0a, 0x22, 0x95, 0xa1, 0x28, 0x02,
	0xec, 0x01, 0x86, 0x3d, 0xc4, 0x2e, 0x77, 0xb9, 0xab, 0x3d, 0xc3, 0x9e, 0x22, 0xd7, 0x7d, 0x84,
	0xde, 0x6c, 0x10, 0x29, 0xd2, 0x96, 0x2d, 0xc5, 0x0a, 0x9a, 0x3b, 0x9a, 0x3c, 0xdf, 0x77, 0xce,
	0x47, 0x1d, 0x9e, 0x73, 0x0c, 0x8e, 0x5d, 0x8c, 0xdd, 0x11, 0x32, 0xa0, 0x43, 0x0c, 0xbe, 0x4c,
	0x56, 0xe7, 0x6d, 0x23, 0x42, 0x04, 0xc7, 0xd1, 0x10, 0x11, 0x03, 0x0e, 0x87, 0x38, 0x0e, 0xa8,
	0x3d, 0x88, 0x1d, 0x17, 0x51, 0x3d, 0x8c, 0x30, 0xc5, 0xea, 0x3d, 0x6e, 0xac, 0x43, 0x87, 0xe8,
	0x12, 0xa7, 0x9f, 0xb7, 0x75, 0x89, 0xdb, 0xf9, 0xb6, 0x88, 0x1a, 0x05, 0xb1, 0x3f, 0x4d, 0x6b,
	0x87, 0x11, 0x0e, 0x31, 0x81, 0x23, 0x9b, 0xbe, 0x0d, 0x11, 0x77, 0xb2, 0xf3, 0xf5, 0xb5, 0x18,
	0x08, 0x85, 0x34, 0x26, 0x29, 0xf4, 0xe1, 0xd5, 0x50, 0x12, 0xa2, 0xc0, 0xf1, 0x02, 0xd7, 0x1e,
	0x79, 0xbe, 0x47, 0x27, 0x7d, 0xde, 0xbf, 0x1a, 0x48, 0x3d, 0x1f, 0x4d, 0x9a, 0x7f, 0x2a, 0xcc,
	0x43, 0xcf, 0x78, 0xed, 0xa1, 0x91, 0x63, 0x0f, 0xd0, 0x1b, 0x78, 0xee, 0xe1, 0x28, 0x35, 0x68,
	0x4e, 0x18, 0x88, 0xbb, 0x49, 0x8f, 0xee, 0xa6, 0x47, 0xec, 0xd7, 0x20, 0x7e, 0x6d, 0xfc, 0x1a,
	0xc1, 0x30, 0x44, 0x91, 0xd0, 0xb0, 0x3b, 0x01, 0x85, 0x41, 0x80, 0x29, 0xa4, 0x1e, 0x0e, 0xd2,
	0xd3, 0xcf, 0xfe, 0x69, 0x81, 0x9a, 0xc9, 0x6f, 0xa0, 0xcb, 0x2e, 0x40, 0xed, 0x81, 0x9a, 0xf0,
	0x60, 0x07, 0xd0, 0x47, 0x9a, 0xb2, 0xa7, 0x1c, 0xac, 0x74, 0xf5, 0x4b, 0x73, 0xf1, 0x83, 0x79,
	0x00, 0x3e, 0x1f, 0x7f, 0xa7, 0x74, 0x15, 0x7a, 0x44, 0x1f, 0x62, 0xdf, 0xc8, 0xd0, 0x58, 0x6b,
	0x82, 0xe4, 0x39, 0xf4, 0x91, 0xfa, 0x00, 0x54, 0x3c, 0x47, 0xab, 0xec, 0x29, 0x07, 0xab, 0x87,
	0xad, 0x14, 0xa8, 0x8b, 0x88, 0xf5, 0x27, 0x01, 0x3d, 0x3e, 0xea, 0xc3, 0x51, 0x8c, 0xba, 0xd5,
	0x4b, 0xb3, 0x6a, 0x55, 0x3c, 0x47, 0x0d, 0x41, 0x6d, 0xe0, 0x8d, 0x46, 0xc9, 0xed, 0x12, 0x44,
	0xe3, 0x50, 0xab, 0x32, 0xf0, 0xee, 0x0c, 0xb8, 0x47, 0x23, 0x2f, 0x70, 0x39, 0xfa, 0xfe, 0xa5,
	0x59, 0xfd, 0x60, 0x7e, 0x01, 0xf6, 0x0b, 0x83, 0xec, 0x72, 0xce, 0x5e, 0x42, 0x69, 0xad, 0x0d,
	0x26, 0x7e, 0xa9, 0x2e, 0x58, 0xe2, 0x1f, 0x5f, 0x5b, 0xd8, 0x53, 0x0e, 0xd6, 0x0f, 0x9f, 0xea,
	0x45, 0xd9, 0xc9, 0x3e, 0xa2, 0x9e, 0xd1, 0xdb, 0x63, 0xc8, 0x47, 0x41, 0xec, 0xe7, 0xed, 0x73,
	0x59, 0x29, 0xbd, 0xfa, 0x15, 0x58, 0x60, 0x17, 0xbb, 0x58, 0x42, 0x11, 0x03, 0x32, 0x73, 0xf5,
	0x15, 0xd0, 0x78, 0x7a, 0x23, 0x27, 0xc9, 0xd2, 0x88, 0xda, 0x0e, 0xa4, 0xc8, 0x4e, 0x72, 0x49,
	0x5b, 0x2a, 0x4b, 0xd5, 0x10, 0x14, 0xbd, 0x84, 0xe1, 0x3b, 0x48, 0xd1, 0x4b, 0x8f, 0x73, 0xc3,
	0x30, 0x8c, 0xf0, 0x79, 0x0e, 0xf7, 0xad, 0xd2, 0xdc, 0x82, 0x62, 0x86, 0x9b, 0x62, 0x0a, 0x47,
	0x36, 0x74, 0x7e, 0x89, 0x09, 0xf5, 0x51, 0x40, 0x89, 0xed, 0x7b, 0xc3, 0x08, 0x13, 0x4d, 0x2d,
	0x99, 0x11, 0xdb, 0x8c, 0xc1, 0x1c, 0x13, 0x3c, 0x63, 0x78, 0xd5, 0x02, 0xb7, 0xa1, 0xcf, 0x9e,
	0x2f, 0x41, 0x51, 0x12, 0x7c, 0xca, 0x5b, 0x2f, 0xc9, 0xab, 0x72, 0x74, 0x8f, 0x81, 0x53, 0xce,
	0x1f, 0x41, 0x23, 0x8c, 0xa3, 0xe1, 0x1b, 0x48, 0x90, 0x8d, 0x23, 0x07, 0x45, 0x76, 0x10, 0xfb,
	0x03, 0x14, 0x69, 0xb7, 0xcb, 0x5e, 0x44, 0x5d, 0xe0, 0x5f, 0x24, 0xf0, 0xe7, 0x0c, 0xad, 0x3e,
	0x04, 0x8b, 0x01, 0xa6, 0x88, 0x68, 0x8d, 0xb2, 0x34, 0xdc, 0x5e, 0xbd, 0x00, 0x9b, 0xa2, 0xce,
	0x88, 0xf2, 0xa6, 0x6d, 0x33, 0x8e, 0x67, 0xfa, 0xdc, 0xfa, 0x99, 0xcd, 0x46, 0xfd, 0x94, 0x13,
	0x65, 0x36, 0x4f, 0x53, 0x52, 0xee, 0x74, 0x23, 0xf5, 0x25, 0x76, 0xd5, 0x9f, 0xc0, 0xb6, 0x4c,
	0x3b, 0x14, 0x38, 0x13, 0x89, 0xb1, 0x5c, 0x52, 0xc8, 0xe3, 0x4f, 0xac, 0xba, 0xa0, 0x78, 0x14,
	0x38, 0x32, 0x31, 0xc2, 0x29, 0x66, 0x59, 0x15, 0xb5, 0x15, 0xf6, 0x00, 0x8f, 0xe6, 0x3c, 0xc0,
	0x84, 0xe4, 0xe5, 0xdb, 0x10, 0xb1, 0x57, 0x27, 0x7e, 0xe4, 0x79, 0x14, 0x67, 0x89, 0x16, 0x99,
	0xe6, 0x59, 0x2d, 0xa0, 0xac, 0x16, 0xc5, 0xaa, 0x0b, 0x8a, 0x29, 0x2d, 0x19, 0xe6, 0xb1, 0x96,
	0xd5, 0x8f, 0xd5, 0x92, 0xf5, 0x28, 0xb5, 0xb8, 0xe0, 0xee, 0xb8, 0x1c, 0x64, 0x1b, 0x51, 0xfa,
	0x08, 0xd6, 0xca, 0x3d, 0x82, 0xc7, 0x15, 0xab, 0x25, 0xab, 0x42, 0x4a, 0xf4, 0x43, 0xc2, 0x93,
	0xbe, 0x87, 0x3f, 0x14, 0xb0, 0x5b, 0xe4, 0x89, 0x29, 0xac, 0x31, 0x85, 0xdf, 0xcf, 0x51, 0x98,
	0xa1, 0x96, 0x52, 0x67, 0x76, 0x45, 0x48, 0xcd, 0xdc, 0x90, 0x84, 0xf2, 0x71, 0xb1, 0xca, 0x55,
	0xbe, 0x5e, 0x52, 0x79, 0xd5, 0x6a, 0xc9, 0x9a, 0x55, 0xa0, 0xbc, 0xc8, 0x13, 0x53, 0xbe, 0x71,
	0xf3, 0xca, 0xab, 0x56, 0x33, 0x37, 0x24, 0xa9, 0x9c, 0xd5, 0xc0, 0x42, 0xe5, 0x9b, 0x25, 0x95,
	0x2f, 0x58, 0x2d, 0xc1, 0x54, 0xa8, 0xbc, 0xc0, 0x13, 0x53, 0xbe, 0x75, 0xf3, 0xca, 0x17, 0xac,
	0x66, 0x6e, 0x48, 0x89, 0xc5, 0xce, 0xdf, 0xcb, 0x60, 0xf7, 0xaa, 0xe2, 0xa5, 0xfe, 0xae, 0x80,
	0x3b, 0x05, 0xc3, 0x20, 0x9b, 0x60, 0xe6, 0x3d, 0xee, 0x63, 0x36, 0x3a, 0x3c, 0x00, 0x7a, 0xb9,
	0xf9, 0x46, 0xf8, 0xb5, 0x1a, 0x30, 0x37, 0x9c, 0xdf, 0x14, 0x50, 0xcb, 0x0c, 0xa3, 0x6c, 0xf8,
	0x59, 0x3f, 0xec, 0x5f, 0x67, 0xa8, 0x10, 0x6c, 0xf2, 0xe6, 0x0a, 0x4f, 0x79, 0xed, 0x5e, 0x0b,
	0x27, 0xb6, 0xe4, 0x98, 0x51, 0xbd, 0xde, 0x98, 0xf1, 0x04, 0x6c, 0x4c, 0x4f, 0x00, 0x0b, 0x65,
	0x19, 0x6a, 0x24, 0xd3, 0xf9, 0x0b, 0x3b, 0xe9, 0xca, 0xcd, 0x74, 0x52, 0x70, 0xcd, 0x4e, 0xfa,
	0x02, 0xa8, 0xc3, 0x08, 0xb1, 0xf9, 0x77, 0x42, 0xdd, 0x6a, 0x59, 0x96, 0x4d, 0x01, 0x96, 0x02,
	0x4f, 0x40, 0x2d, 0xdb, 0x46, 0x16, 0xcb, 0xb7, 0xc4, 0x55, 0x34, 0xd1, 0x3e, 0x6c, 0x4e, 0x34,
	0xee, 0x1a, 0x4b, 0x1f, 0xdf, 0x01, 0x13, 0x07, 0xb2, 0x5b, 0xf4, 0x41, 0x23, 0xbf, 0x60, 0xdc,
	0x2a, 0x59, 0x30, 0x14, 0xab, 0x4e, 0x72, 0x0a, 0xc5, 0x05, 0xa8, 0xe7, 0x95, 0x87, 0xe5, 0x9b,
	0x2f, 0x0f, 0x8a, 0xb5, 0x45, 0x66, 0x4e, 0x00, 0x58, 0x16, 0xf7, 0xd6, 0xdd, 0x04, 0xeb, 0xd9,
	0x50, 0x3a, 0xf0, 0xbd, 0xf9, 0x73, 0xd9, 0x3f, 0x2c, 0xea, 0xd1, 0x30, 0x26, 0x14, 0xfb, 0x28,
	0x22, 0xc6, 0x3b, 0xb1, 0xbc, 0x30, 0x32, 0xaf, 0x9b, 0x18, 0xef, 0xb2, 0x25, 0xe6, 0xa2, 0x5b,
	0x07, 0x5b, 0x33, 0x33, 0x4c, 0xb2, 0x39, 0x33, 0x0c, 0x74, 0x9b, 0xe0, 0x4e, 0x41, 0x17, 0x4d,
	0x8e, 0x0a, 0xda, 0x0c, 0x3b, 0xca, 0xaf, 0xc3, 0xdd, 0xff, 0x14, 0xb0, 0x3f, 0xc4, 0xfe, 0xfc,
	0x19, 0xb0, 0xab, 0x4e, 0xd7, 0x0d, 0x8a, 0x4f, 0x95, 0x57, 0x4f, 0x53, 0xa0, 0x8b, 0x47, 0x30,
	0x70, 0x75, 0x1c, 0xb9, 0x86, 0x8b, 0x02, 0x96, 0x00, 0xc6, 0xf8, 0xa6, 0xae, 0xf8, 0x4f, 0xff,
	0x8d, 0x5c, 0xfd, 0x59, 0xa9, 0x9e, 0x98, 0xe6, 0x5f, 0x95, 0x7b, 0x27, 0x9c, 0xd2, 0x74, 0x88,
	0xce, 0x97, 0xc9, 0xaa, 0xdf, 0xd6, 0x2d, 0x61, 0xf9, 0xaf, 0xb0, 0x39, 0x33, 0x1d, 0x72, 0x26,
	0x6d, 0xce, 0xfa, 0xed, 0x33, 0x69, 0xf3, 0xbe, 0xb2, 0xcf, 0x0f, 0x3a, 0x1d, 0xd3, 0x21, 0x9d,
	0x8e, 0xb4, 0xea, 0x74, 0xfa, 0xed, 0x4e, 0x47, 0xda, 0x0d, 0x96, 0x58, 0xb0, 0x5f, 0xfe, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0xc8, 0xcb, 0x3f, 0x95, 0x7f, 0x10, 0x00, 0x00,
}