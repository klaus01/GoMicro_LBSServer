// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/order/proto/order.proto

package order

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type OrderPayStatus int32

const (
	OrderPayStatus_BE_PAID       OrderPayStatus = 0
	OrderPayStatus_PAID          OrderPayStatus = 1
	OrderPayStatus_PAY_EXCEPTION OrderPayStatus = 2
)

var OrderPayStatus_name = map[int32]string{
	0: "BE_PAID",
	1: "PAID",
	2: "PAY_EXCEPTION",
}

var OrderPayStatus_value = map[string]int32{
	"BE_PAID":       0,
	"PAID":          1,
	"PAY_EXCEPTION": 2,
}

func (x OrderPayStatus) String() string {
	return proto.EnumName(OrderPayStatus_name, int32(x))
}

func (OrderPayStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{0}
}

type GetRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type SearchRequest struct {
	PageNo               uint32               `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize             uint32               `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	BeginDateTime        *timestamp.Timestamp `protobuf:"bytes,3,opt,name=beginDateTime,proto3" json:"beginDateTime,omitempty"`
	EndDateTime          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=endDateTime,proto3" json:"endDateTime,omitempty"`
	OrderId              string               `protobuf:"bytes,5,opt,name=orderId,proto3" json:"orderId,omitempty"`
	PhoneNumber          string               `protobuf:"bytes,6,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	PayStatus            OrderPayStatus       `protobuf:"varint,7,opt,name=payStatus,proto3,enum=order.OrderPayStatus" json:"payStatus,omitempty"`
	IsShipped            bool                 `protobuf:"varint,8,opt,name=isShipped,proto3" json:"isShipped,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{1}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetPageNo() uint32 {
	if m != nil {
		return m.PageNo
	}
	return 0
}

func (m *SearchRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchRequest) GetBeginDateTime() *timestamp.Timestamp {
	if m != nil {
		return m.BeginDateTime
	}
	return nil
}

func (m *SearchRequest) GetEndDateTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndDateTime
	}
	return nil
}

func (m *SearchRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *SearchRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SearchRequest) GetPayStatus() OrderPayStatus {
	if m != nil {
		return m.PayStatus
	}
	return OrderPayStatus_BE_PAID
}

func (m *SearchRequest) GetIsShipped() bool {
	if m != nil {
		return m.IsShipped
	}
	return false
}

type SearchResult struct {
	PageNo               uint32        `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageTotal            uint32        `protobuf:"varint,2,opt,name=pageTotal,proto3" json:"pageTotal,omitempty"`
	Datas                []*OrderModel `protobuf:"bytes,3,rep,name=datas,proto3" json:"datas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SearchResult) Reset()         { *m = SearchResult{} }
func (m *SearchResult) String() string { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()    {}
func (*SearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{2}
}

func (m *SearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResult.Unmarshal(m, b)
}
func (m *SearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResult.Marshal(b, m, deterministic)
}
func (m *SearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResult.Merge(m, src)
}
func (m *SearchResult) XXX_Size() int {
	return xxx_messageInfo_SearchResult.Size(m)
}
func (m *SearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResult proto.InternalMessageInfo

func (m *SearchResult) GetPageNo() uint32 {
	if m != nil {
		return m.PageNo
	}
	return 0
}

func (m *SearchResult) GetPageTotal() uint32 {
	if m != nil {
		return m.PageTotal
	}
	return 0
}

func (m *SearchResult) GetDatas() []*OrderModel {
	if m != nil {
		return m.Datas
	}
	return nil
}

type OrderModel struct {
	OrderId              string               `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty" bson:"orderId"`
	CreateAt             *timestamp.Timestamp `protobuf:"bytes,2,opt,name=createAt,proto3" json:"createAt,omitempty" bson:"createAt"`
	ProductName          string               `protobuf:"bytes,3,opt,name=productName,proto3" json:"productName,omitempty" bson:"productName"`
	ProductAmount        float32              `protobuf:"fixed32,4,opt,name=productAmount,proto3" json:"productAmount,omitempty" bson:"productAmount"`
	Name                 string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	PhoneNumber          string               `protobuf:"bytes,6,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty" bson:"phoneNumber"`
	Province             string               `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty" bson:"province"`
	City                 string               `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty" bson:"city"`
	District             string               `protobuf:"bytes,9,opt,name=district,proto3" json:"district,omitempty" bson:"district"`
	Address              string               `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty" bson:"address"`
	PayStatus            OrderPayStatus       `protobuf:"varint,11,opt,name=payStatus,proto3,enum=order.OrderPayStatus" json:"payStatus,omitempty" bson:"payStatus"`
	PayInfo              *OrderPayInfo        `protobuf:"bytes,12,opt,name=payInfo,proto3" json:"payInfo,omitempty" bson:"payInfo"`
	DeliveryInfo         *OrderDeliveryInfo   `protobuf:"bytes,13,opt,name=deliveryInfo,proto3" json:"deliveryInfo,omitempty" bson:"deliveryInfo"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" bson:"-"`
	XXX_unrecognized     []byte               `json:"-" bson:"-"`
	XXX_sizecache        int32                `json:"-" bson:"-"`
}

func (m *OrderModel) Reset()         { *m = OrderModel{} }
func (m *OrderModel) String() string { return proto.CompactTextString(m) }
func (*OrderModel) ProtoMessage()    {}
func (*OrderModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{3}
}

func (m *OrderModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderModel.Unmarshal(m, b)
}
func (m *OrderModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderModel.Marshal(b, m, deterministic)
}
func (m *OrderModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderModel.Merge(m, src)
}
func (m *OrderModel) XXX_Size() int {
	return xxx_messageInfo_OrderModel.Size(m)
}
func (m *OrderModel) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderModel.DiscardUnknown(m)
}

var xxx_messageInfo_OrderModel proto.InternalMessageInfo

func (m *OrderModel) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderModel) GetCreateAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreateAt
	}
	return nil
}

func (m *OrderModel) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *OrderModel) GetProductAmount() float32 {
	if m != nil {
		return m.ProductAmount
	}
	return 0
}

func (m *OrderModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrderModel) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *OrderModel) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *OrderModel) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *OrderModel) GetDistrict() string {
	if m != nil {
		return m.District
	}
	return ""
}

func (m *OrderModel) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *OrderModel) GetPayStatus() OrderPayStatus {
	if m != nil {
		return m.PayStatus
	}
	return OrderPayStatus_BE_PAID
}

func (m *OrderModel) GetPayInfo() *OrderPayInfo {
	if m != nil {
		return m.PayInfo
	}
	return nil
}

func (m *OrderModel) GetDeliveryInfo() *OrderDeliveryInfo {
	if m != nil {
		return m.DeliveryInfo
	}
	return nil
}

type OrderPayInfo struct {
	ModeName             string               `protobuf:"bytes,1,opt,name=modeName,proto3" json:"modeName,omitempty" bson:"modeName"`
	Money                float32              `protobuf:"fixed32,2,opt,name=money,proto3" json:"money,omitempty" bson:"money"`
	CreateAt             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createAt,proto3" json:"createAt,omitempty" bson:"createAt"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" bson:"-"`
	XXX_unrecognized     []byte               `json:"-" bson:"-"`
	XXX_sizecache        int32                `json:"-" bson:"-"`
}

func (m *OrderPayInfo) Reset()         { *m = OrderPayInfo{} }
func (m *OrderPayInfo) String() string { return proto.CompactTextString(m) }
func (*OrderPayInfo) ProtoMessage()    {}
func (*OrderPayInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{4}
}

func (m *OrderPayInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderPayInfo.Unmarshal(m, b)
}
func (m *OrderPayInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderPayInfo.Marshal(b, m, deterministic)
}
func (m *OrderPayInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPayInfo.Merge(m, src)
}
func (m *OrderPayInfo) XXX_Size() int {
	return xxx_messageInfo_OrderPayInfo.Size(m)
}
func (m *OrderPayInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPayInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPayInfo proto.InternalMessageInfo

func (m *OrderPayInfo) GetModeName() string {
	if m != nil {
		return m.ModeName
	}
	return ""
}

func (m *OrderPayInfo) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *OrderPayInfo) GetCreateAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreateAt
	}
	return nil
}

type OrderDeliveryInfo struct {
	CourierCompany       string               `protobuf:"bytes,1,opt,name=courierCompany,proto3" json:"courierCompany,omitempty" bson:"courierCompany"`
	WaybillNumber        string               `protobuf:"bytes,2,opt,name=waybillNumber,proto3" json:"waybillNumber,omitempty" bson:"waybillNumber"`
	CreateAt             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createAt,proto3" json:"createAt,omitempty" bson:"createAt"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" bson:"-"`
	XXX_unrecognized     []byte               `json:"-" bson:"-"`
	XXX_sizecache        int32                `json:"-" bson:"-"`
}

func (m *OrderDeliveryInfo) Reset()         { *m = OrderDeliveryInfo{} }
func (m *OrderDeliveryInfo) String() string { return proto.CompactTextString(m) }
func (*OrderDeliveryInfo) ProtoMessage()    {}
func (*OrderDeliveryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{5}
}

func (m *OrderDeliveryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderDeliveryInfo.Unmarshal(m, b)
}
func (m *OrderDeliveryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderDeliveryInfo.Marshal(b, m, deterministic)
}
func (m *OrderDeliveryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderDeliveryInfo.Merge(m, src)
}
func (m *OrderDeliveryInfo) XXX_Size() int {
	return xxx_messageInfo_OrderDeliveryInfo.Size(m)
}
func (m *OrderDeliveryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderDeliveryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderDeliveryInfo proto.InternalMessageInfo

func (m *OrderDeliveryInfo) GetCourierCompany() string {
	if m != nil {
		return m.CourierCompany
	}
	return ""
}

func (m *OrderDeliveryInfo) GetWaybillNumber() string {
	if m != nil {
		return m.WaybillNumber
	}
	return ""
}

func (m *OrderDeliveryInfo) GetCreateAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreateAt
	}
	return nil
}

type CreateRequest struct {
	ProductName          string   `protobuf:"bytes,1,opt,name=productName,proto3" json:"productName,omitempty"`
	ProductAmount        float32  `protobuf:"fixed32,2,opt,name=productAmount,proto3" json:"productAmount,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Province             string   `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	District             string   `protobuf:"bytes,7,opt,name=district,proto3" json:"district,omitempty"`
	Address              string   `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{6}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *CreateRequest) GetProductAmount() float32 {
	if m != nil {
		return m.ProductAmount
	}
	return 0
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *CreateRequest) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *CreateRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CreateRequest) GetDistrict() string {
	if m != nil {
		return m.District
	}
	return ""
}

func (m *CreateRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CreateResult struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResult) Reset()         { *m = CreateResult{} }
func (m *CreateResult) String() string { return proto.CompactTextString(m) }
func (*CreateResult) ProtoMessage()    {}
func (*CreateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{7}
}

func (m *CreateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResult.Unmarshal(m, b)
}
func (m *CreateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResult.Marshal(b, m, deterministic)
}
func (m *CreateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResult.Merge(m, src)
}
func (m *CreateResult) XXX_Size() int {
	return xxx_messageInfo_CreateResult.Size(m)
}
func (m *CreateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResult.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResult proto.InternalMessageInfo

func (m *CreateResult) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type SetDeliveryInfoRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	CourierCompany       string   `protobuf:"bytes,2,opt,name=courierCompany,proto3" json:"courierCompany,omitempty"`
	WaybillNumber        string   `protobuf:"bytes,3,opt,name=waybillNumber,proto3" json:"waybillNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDeliveryInfoRequest) Reset()         { *m = SetDeliveryInfoRequest{} }
func (m *SetDeliveryInfoRequest) String() string { return proto.CompactTextString(m) }
func (*SetDeliveryInfoRequest) ProtoMessage()    {}
func (*SetDeliveryInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{8}
}

func (m *SetDeliveryInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeliveryInfoRequest.Unmarshal(m, b)
}
func (m *SetDeliveryInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeliveryInfoRequest.Marshal(b, m, deterministic)
}
func (m *SetDeliveryInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeliveryInfoRequest.Merge(m, src)
}
func (m *SetDeliveryInfoRequest) XXX_Size() int {
	return xxx_messageInfo_SetDeliveryInfoRequest.Size(m)
}
func (m *SetDeliveryInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeliveryInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeliveryInfoRequest proto.InternalMessageInfo

func (m *SetDeliveryInfoRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *SetDeliveryInfoRequest) GetCourierCompany() string {
	if m != nil {
		return m.CourierCompany
	}
	return ""
}

func (m *SetDeliveryInfoRequest) GetWaybillNumber() string {
	if m != nil {
		return m.WaybillNumber
	}
	return ""
}

type SetPayInfoRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	ModeName             string   `protobuf:"bytes,2,opt,name=modeName,proto3" json:"modeName,omitempty"`
	Money                float32  `protobuf:"fixed32,3,opt,name=money,proto3" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetPayInfoRequest) Reset()         { *m = SetPayInfoRequest{} }
func (m *SetPayInfoRequest) String() string { return proto.CompactTextString(m) }
func (*SetPayInfoRequest) ProtoMessage()    {}
func (*SetPayInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc4dee17e320b1ee, []int{9}
}

func (m *SetPayInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPayInfoRequest.Unmarshal(m, b)
}
func (m *SetPayInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPayInfoRequest.Marshal(b, m, deterministic)
}
func (m *SetPayInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPayInfoRequest.Merge(m, src)
}
func (m *SetPayInfoRequest) XXX_Size() int {
	return xxx_messageInfo_SetPayInfoRequest.Size(m)
}
func (m *SetPayInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPayInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetPayInfoRequest proto.InternalMessageInfo

func (m *SetPayInfoRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *SetPayInfoRequest) GetModeName() string {
	if m != nil {
		return m.ModeName
	}
	return ""
}

func (m *SetPayInfoRequest) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func init() {
	proto.RegisterEnum("order.OrderPayStatus", OrderPayStatus_name, OrderPayStatus_value)
	proto.RegisterType((*GetRequest)(nil), "order.GetRequest")
	proto.RegisterType((*SearchRequest)(nil), "order.SearchRequest")
	proto.RegisterType((*SearchResult)(nil), "order.SearchResult")
	proto.RegisterType((*OrderModel)(nil), "order.OrderModel")
	proto.RegisterType((*OrderPayInfo)(nil), "order.OrderPayInfo")
	proto.RegisterType((*OrderDeliveryInfo)(nil), "order.OrderDeliveryInfo")
	proto.RegisterType((*CreateRequest)(nil), "order.CreateRequest")
	proto.RegisterType((*CreateResult)(nil), "order.CreateResult")
	proto.RegisterType((*SetDeliveryInfoRequest)(nil), "order.SetDeliveryInfoRequest")
	proto.RegisterType((*SetPayInfoRequest)(nil), "order.SetPayInfoRequest")
}

func init() {
	proto.RegisterFile("srv/order/proto/order.proto", fileDescriptor_dc4dee17e320b1ee)
}

var fileDescriptor_dc4dee17e320b1ee = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x5f, 0x6e, 0xfa, 0x46,
	0x10, 0xc6, 0x36, 0x7f, 0x07, 0x48, 0xc3, 0x26, 0x8d, 0x2c, 0x92, 0xaa, 0xc8, 0xaa, 0x52, 0x54,
	0x29, 0x44, 0x22, 0x6a, 0x1f, 0xaa, 0x3c, 0x84, 0x26, 0x28, 0xa2, 0x52, 0x09, 0x32, 0x3c, 0xb4,
	0x4f, 0x91, 0xc1, 0x13, 0x62, 0x09, 0x7b, 0x5d, 0x7b, 0x9d, 0x96, 0x3e, 0xf5, 0x1a, 0xbd, 0x46,
	0x8f, 0xd2, 0x33, 0xf4, 0x06, 0xbd, 0x40, 0xe5, 0x5d, 0xff, 0x25, 0x71, 0x82, 0x7e, 0x2f, 0x68,
	0x67, 0xe6, 0x9b, 0xdd, 0xd9, 0xf9, 0x3e, 0xcf, 0x02, 0xa7, 0xbe, 0xf7, 0x72, 0x49, 0x3d, 0x13,
	0xbd, 0x4b, 0xd7, 0xa3, 0x8c, 0x8a, 0xf5, 0x80, 0xaf, 0x49, 0x85, 0x1b, 0xdd, 0xd3, 0x35, 0xa5,
	0xeb, 0x0d, 0x0a, 0xc0, 0x32, 0x78, 0xba, 0x44, 0xdb, 0x65, 0x5b, 0x81, 0xe9, 0x7e, 0xb9, 0x1b,
	0x64, 0x96, 0x8d, 0x3e, 0x33, 0x6c, 0x57, 0x00, 0xb4, 0x73, 0x80, 0x7b, 0x64, 0x3a, 0xfe, 0x1a,
	0xa0, 0xcf, 0x88, 0x0a, 0x35, 0xbe, 0xe9, 0xc4, 0x54, 0xa5, 0x9e, 0xd4, 0x6f, 0xe8, 0xb1, 0xa9,
	0xfd, 0x23, 0x43, 0x7b, 0x8e, 0x86, 0xb7, 0x7a, 0x8e, 0xb1, 0x27, 0x50, 0x75, 0x8d, 0x35, 0x4e,
	0x29, 0x87, 0xb6, 0xf5, 0xc8, 0x22, 0x5d, 0xa8, 0x87, 0xab, 0xb9, 0xf5, 0x07, 0xaa, 0x32, 0x8f,
	0x24, 0x36, 0xb9, 0x81, 0xf6, 0x12, 0xd7, 0x96, 0x73, 0x67, 0x30, 0x5c, 0x58, 0x36, 0xaa, 0x4a,
	0x4f, 0xea, 0x37, 0x87, 0xdd, 0x81, 0x28, 0x73, 0x10, 0x97, 0x39, 0x58, 0xc4, 0x65, 0xea, 0xf9,
	0x04, 0x72, 0x0d, 0x4d, 0x74, 0xcc, 0x24, 0xbf, 0xfc, 0x61, 0x7e, 0x16, 0x9e, 0xbd, 0x5f, 0x25,
	0x77, 0x3f, 0xd2, 0x83, 0xa6, 0xfb, 0x4c, 0x1d, 0x9c, 0x06, 0xf6, 0x12, 0x3d, 0xb5, 0xca, 0xa3,
	0x59, 0x17, 0xb9, 0x82, 0x86, 0x6b, 0x6c, 0xe7, 0xcc, 0x60, 0x81, 0xaf, 0xd6, 0x7a, 0x52, 0xff,
	0x60, 0xf8, 0xf9, 0x40, 0xf0, 0xf1, 0x10, 0xfe, 0xce, 0xe2, 0xa0, 0x9e, 0xe2, 0xc8, 0x19, 0x34,
	0x2c, 0x7f, 0xfe, 0x6c, 0xb9, 0x2e, 0x9a, 0x6a, 0xbd, 0x27, 0xf5, 0xeb, 0x7a, 0xea, 0xd0, 0x6c,
	0x68, 0xc5, 0x3d, 0xf5, 0x83, 0x4d, 0x71, 0x4b, 0xcf, 0xc2, 0xa3, 0xd7, 0xb8, 0xa0, 0xcc, 0xd8,
	0x44, 0x3d, 0x4d, 0x1d, 0xe4, 0x6b, 0xa8, 0x98, 0x06, 0x33, 0x7c, 0x55, 0xe9, 0x29, 0xfd, 0xe6,
	0xb0, 0x93, 0x2d, 0xea, 0x27, 0x6a, 0xe2, 0x46, 0x17, 0x71, 0xed, 0x5f, 0x05, 0x20, 0xf5, 0x16,
	0x93, 0x4d, 0xbe, 0x83, 0xfa, 0xca, 0x43, 0x83, 0xe1, 0x88, 0xf1, 0xe3, 0xde, 0xef, 0x70, 0x82,
	0xe5, 0x4d, 0xf4, 0xa8, 0x19, 0xac, 0xd8, 0xd4, 0x88, 0xc8, 0x0d, 0x9b, 0x98, 0xba, 0xc8, 0x57,
	0xd0, 0x8e, 0xcc, 0x91, 0x4d, 0x03, 0x87, 0x71, 0x02, 0x65, 0x3d, 0xef, 0x24, 0x04, 0xca, 0x4e,
	0xb8, 0x81, 0xe0, 0x88, 0xaf, 0xf7, 0x20, 0x28, 0x14, 0x9e, 0x47, 0x5f, 0x2c, 0x67, 0x85, 0x9c,
	0x9f, 0x86, 0x9e, 0xd8, 0xe1, 0x8e, 0x2b, 0x8b, 0x6d, 0x39, 0x05, 0x0d, 0x9d, 0xaf, 0x43, 0xbc,
	0x69, 0xf9, 0xcc, 0xb3, 0x56, 0x4c, 0x6d, 0x08, 0x7c, 0x6c, 0x87, 0xbd, 0x31, 0x4c, 0xd3, 0x43,
	0xdf, 0x57, 0x41, 0xf4, 0x26, 0x32, 0xf3, 0x32, 0x68, 0xee, 0x29, 0x83, 0x0b, 0xa8, 0xb9, 0xc6,
	0x76, 0xe2, 0x3c, 0x51, 0xb5, 0xc5, 0xfb, 0x79, 0xb4, 0x93, 0x12, 0x86, 0xf4, 0x18, 0x43, 0xae,
	0xa1, 0x65, 0xe2, 0xc6, 0x7a, 0x41, 0x4f, 0xe4, 0xb4, 0x79, 0x8e, 0x9a, 0xcd, 0xb9, 0xcb, 0xc4,
	0xf5, 0x1c, 0x5a, 0xfb, 0x1d, 0x5a, 0xd9, 0x6d, 0xc3, 0x7b, 0xda, 0xd4, 0x44, 0x4e, 0x89, 0x20,
	0x3a, 0xb1, 0xc9, 0x31, 0x54, 0x6c, 0xea, 0xe0, 0x96, 0xd3, 0x2c, 0xeb, 0xc2, 0xc8, 0xf1, 0xaf,
	0xec, 0xcf, 0xbf, 0xf6, 0x97, 0x04, 0x9d, 0x57, 0xd5, 0x91, 0x73, 0x38, 0x58, 0xd1, 0xc0, 0xb3,
	0xd0, 0xbb, 0xa5, 0xb6, 0x6b, 0x38, 0xdb, 0xa8, 0x8a, 0x1d, 0x6f, 0xa8, 0x8d, 0xdf, 0x8c, 0xed,
	0xd2, 0xda, 0x6c, 0x22, 0x8e, 0x65, 0x0e, 0xcb, 0x3b, 0x3f, 0xb9, 0xb6, 0xff, 0x24, 0x68, 0xdf,
	0x72, 0x23, 0x1e, 0x60, 0x3b, 0x6a, 0x95, 0xf6, 0x50, 0xab, 0xfc, 0x9e, 0x5a, 0x95, 0x62, 0xb5,
	0x96, 0xdf, 0x57, 0x6b, 0xa5, 0x40, 0xad, 0xd5, 0x02, 0xb5, 0xd6, 0x8a, 0xd5, 0x5a, 0xcf, 0xa9,
	0x55, 0xeb, 0x43, 0x2b, 0xbe, 0x34, 0x9f, 0x30, 0xc5, 0x03, 0xfe, 0x4f, 0x09, 0x4e, 0xe6, 0xc8,
	0x72, 0xba, 0xfa, 0xe8, 0x55, 0x78, 0x83, 0x5a, 0x79, 0x3f, 0x6a, 0x95, 0x37, 0xa8, 0xd5, 0x1e,
	0xa1, 0x33, 0x47, 0x16, 0x7f, 0x0d, 0x1f, 0x1e, 0x9e, 0xd5, 0xb5, 0x5c, 0xa4, 0x6b, 0x25, 0xa3,
	0xeb, 0x6f, 0xbe, 0x87, 0x83, 0xfc, 0x37, 0x4a, 0x9a, 0x50, 0xfb, 0x61, 0xfc, 0x38, 0x1b, 0x4d,
	0xee, 0x0e, 0x4b, 0xa4, 0x0e, 0x65, 0xbe, 0x92, 0x48, 0x07, 0xda, 0xb3, 0xd1, 0x2f, 0x8f, 0xe3,
	0x9f, 0x6f, 0xc7, 0xb3, 0xc5, 0xe4, 0x61, 0x7a, 0x28, 0x0f, 0xff, 0x96, 0xa1, 0xc2, 0x93, 0xc9,
	0x05, 0x28, 0xf7, 0xc8, 0x48, 0x3c, 0x67, 0xd3, 0xe7, 0xb3, 0xfb, 0x7a, 0xf4, 0x6a, 0x25, 0xf2,
	0x2d, 0x54, 0xc5, 0x90, 0x27, 0xc7, 0x51, 0x38, 0xf7, 0x8e, 0x76, 0x8f, 0x76, 0xbc, 0x21, 0x4f,
	0x22, 0x4d, 0x30, 0x97, 0xa4, 0xe5, 0xd4, 0x9b, 0xa4, 0x65, 0xe9, 0xd5, 0x4a, 0xe4, 0x47, 0xf8,
	0x6c, 0x87, 0x45, 0xf2, 0x45, 0x72, 0xc0, 0x5b, 0xec, 0x76, 0x4f, 0x5e, 0x7d, 0x3e, 0xe3, 0xf0,
	0x0f, 0x84, 0x56, 0x22, 0x37, 0x00, 0x29, 0x1f, 0x44, 0x4d, 0xb7, 0xc9, 0x53, 0x54, 0xbc, 0xc3,
	0xb2, 0xca, 0x3d, 0x57, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xed, 0x39, 0x46, 0xc8, 0x08,
	0x00, 0x00,
}
