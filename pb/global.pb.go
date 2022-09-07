// 全局 Protobuf

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: global/global.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmptyPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyPost) Reset() {
	*x = EmptyPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyPost) ProtoMessage() {}

func (x *EmptyPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyPost.ProtoReflect.Descriptor instead.
func (*EmptyPost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{0}
}

type InfoPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID"`
}

func (x *InfoPost) Reset() {
	*x = InfoPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoPost) ProtoMessage() {}

func (x *InfoPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoPost.ProtoReflect.Descriptor instead.
func (*InfoPost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{1}
}

func (x *InfoPost) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type KeyPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key"`
}

func (x *KeyPost) Reset() {
	*x = KeyPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyPost) ProtoMessage() {}

func (x *KeyPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyPost.ProtoReflect.Descriptor instead.
func (*KeyPost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{2}
}

func (x *KeyPost) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type KeywordPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=Keyword,proto3" json:"Keyword"`
}

func (x *KeywordPost) Reset() {
	*x = KeywordPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPost) ProtoMessage() {}

func (x *KeywordPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPost.ProtoReflect.Descriptor instead.
func (*KeywordPost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{3}
}

func (x *KeywordPost) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type DatePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string `protobuf:"bytes,1,opt,name=Start,proto3" json:"Start"`
	End   string `protobuf:"bytes,2,opt,name=End,proto3" json:"End"`
}

func (x *DatePost) Reset() {
	*x = DatePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatePost) ProtoMessage() {}

func (x *DatePost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatePost.ProtoReflect.Descriptor instead.
func (*DatePost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{4}
}

func (x *DatePost) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *DatePost) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

// PaginationPost 普通分页 传参
type PaginationPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size uint64 `protobuf:"varint,1,opt,name=Size,proto3" json:"Size"`
	Page uint64 `protobuf:"varint,2,opt,name=Page,proto3" json:"Page"`
}

func (x *PaginationPost) Reset() {
	*x = PaginationPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationPost) ProtoMessage() {}

func (x *PaginationPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationPost.ProtoReflect.Descriptor instead.
func (*PaginationPost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{5}
}

func (x *PaginationPost) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PaginationPost) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

// SearchPaginationPost 搜索分页 传参
type SearchPaginationPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size    uint64 `protobuf:"varint,1,opt,name=Size,proto3" json:"Size"`
	Page    uint64 `protobuf:"varint,2,opt,name=Page,proto3" json:"Page"`
	Keyword string `protobuf:"bytes,3,opt,name=Keyword,proto3" json:"Keyword"`
	Status  uint64 `protobuf:"varint,4,opt,name=Status,proto3" json:"Status"`
	Start   string `protobuf:"bytes,5,opt,name=Start,proto3" json:"Start"`
	End     string `protobuf:"bytes,6,opt,name=End,proto3" json:"End"`
}

func (x *SearchPaginationPost) Reset() {
	*x = SearchPaginationPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPaginationPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPaginationPost) ProtoMessage() {}

func (x *SearchPaginationPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPaginationPost.ProtoReflect.Descriptor instead.
func (*SearchPaginationPost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{6}
}

func (x *SearchPaginationPost) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchPaginationPost) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchPaginationPost) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchPaginationPost) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SearchPaginationPost) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *SearchPaginationPost) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

// ChangeStatusPost 通用修改状态 传参
type ChangeStatusPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID"`
	Status uint64 `protobuf:"varint,2,opt,name=Status,proto3" json:"Status"`
}

func (x *ChangeStatusPost) Reset() {
	*x = ChangeStatusPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatusPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusPost) ProtoMessage() {}

func (x *ChangeStatusPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusPost.ProtoReflect.Descriptor instead.
func (*ChangeStatusPost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{7}
}

func (x *ChangeStatusPost) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ChangeStatusPost) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ChangeStatusByKeyPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key"`
	Status uint64 `protobuf:"varint,2,opt,name=Status,proto3" json:"Status"`
}

func (x *ChangeStatusByKeyPost) Reset() {
	*x = ChangeStatusByKeyPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatusByKeyPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusByKeyPost) ProtoMessage() {}

func (x *ChangeStatusByKeyPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusByKeyPost.ProtoReflect.Descriptor instead.
func (*ChangeStatusByKeyPost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{8}
}

func (x *ChangeStatusByKeyPost) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ChangeStatusByKeyPost) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// UpdateAreaPost 更新所在区域
type UpdateAreaPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode  uint64 `protobuf:"varint,1,opt,name=country_code,json=countryCode,proto3" json:"country_code"`
	ProvinceCode uint64 `protobuf:"varint,2,opt,name=province_code,json=provinceCode,proto3" json:"province_code"`
	CityCode     uint64 `protobuf:"varint,3,opt,name=city_code,json=cityCode,proto3" json:"city_code"`
	Country      string `protobuf:"bytes,4,opt,name=country,proto3" json:"country"`
	Province     string `protobuf:"bytes,5,opt,name=province,proto3" json:"province"`
	City         string `protobuf:"bytes,6,opt,name=city,proto3" json:"city"`
}

func (x *UpdateAreaPost) Reset() {
	*x = UpdateAreaPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAreaPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAreaPost) ProtoMessage() {}

func (x *UpdateAreaPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAreaPost.ProtoReflect.Descriptor instead.
func (*UpdateAreaPost) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateAreaPost) GetCountryCode() uint64 {
	if x != nil {
		return x.CountryCode
	}
	return 0
}

func (x *UpdateAreaPost) GetProvinceCode() uint64 {
	if x != nil {
		return x.ProvinceCode
	}
	return 0
}

func (x *UpdateAreaPost) GetCityCode() uint64 {
	if x != nil {
		return x.CityCode
	}
	return 0
}

func (x *UpdateAreaPost) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UpdateAreaPost) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *UpdateAreaPost) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string            `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message"`
	Data    map[string]string `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{10}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type StringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    string `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data"`
}

func (x *StringResponse) Reset() {
	*x = StringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringResponse) ProtoMessage() {}

func (x *StringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringResponse.ProtoReflect.Descriptor instead.
func (*StringResponse) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{11}
}

func (x *StringResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StringResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// BoolResponse 返回 Bool
type BoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    bool   `protobuf:"varint,2,opt,name=data,proto3" json:"data"`
}

func (x *BoolResponse) Reset() {
	*x = BoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResponse) ProtoMessage() {}

func (x *BoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResponse.ProtoReflect.Descriptor instead.
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{12}
}

func (x *BoolResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BoolResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    bool   `protobuf:"varint,2,opt,name=data,proto3" json:"data"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{13}
}

func (x *CheckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CheckResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    bool   `protobuf:"varint,2,opt,name=data,proto3" json:"data"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_global_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_global_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_global_global_proto_rawDescGZIP(), []int{14}
}

func (x *ResultResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResultResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

var File_global_global_proto protoreflect.FileDescriptor

var file_global_global_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x22, 0x0b, 0x0a,
	0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x1a, 0x0a, 0x08, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x1b, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4b, 0x65, 0x79, 0x22, 0x27, 0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x32, 0x0a, 0x08,
	0x44, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x45, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x6e, 0x64,
	0x22, 0x38, 0x0a, 0x0e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x14, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x45, 0x6e, 0x64, 0x22, 0x3a, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x41, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x65, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x8d, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x13, 0x42, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x01, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_global_global_proto_rawDescOnce sync.Once
	file_global_global_proto_rawDescData = file_global_global_proto_rawDesc
)

func file_global_global_proto_rawDescGZIP() []byte {
	file_global_global_proto_rawDescOnce.Do(func() {
		file_global_global_proto_rawDescData = protoimpl.X.CompressGZIP(file_global_global_proto_rawDescData)
	})
	return file_global_global_proto_rawDescData
}

var file_global_global_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_global_global_proto_goTypes = []interface{}{
	(*EmptyPost)(nil),             // 0: global.EmptyPost
	(*InfoPost)(nil),              // 1: global.InfoPost
	(*KeyPost)(nil),               // 2: global.KeyPost
	(*KeywordPost)(nil),           // 3: global.KeywordPost
	(*DatePost)(nil),              // 4: global.DatePost
	(*PaginationPost)(nil),        // 5: global.PaginationPost
	(*SearchPaginationPost)(nil),  // 6: global.SearchPaginationPost
	(*ChangeStatusPost)(nil),      // 7: global.ChangeStatusPost
	(*ChangeStatusByKeyPost)(nil), // 8: global.ChangeStatusByKeyPost
	(*UpdateAreaPost)(nil),        // 9: global.UpdateAreaPost
	(*Response)(nil),              // 10: global.Response
	(*StringResponse)(nil),        // 11: global.StringResponse
	(*BoolResponse)(nil),          // 12: global.BoolResponse
	(*CheckResponse)(nil),         // 13: global.CheckResponse
	(*ResultResponse)(nil),        // 14: global.ResultResponse
	nil,                           // 15: global.Response.DataEntry
}
var file_global_global_proto_depIdxs = []int32{
	15, // 0: global.Response.Data:type_name -> global.Response.DataEntry
	1,  // [1:1] is the sub-list for method output_type
	1,  // [1:1] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_global_global_proto_init() }
func file_global_global_proto_init() {
	if File_global_global_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_global_global_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatePost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPaginationPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStatusPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStatusByKeyPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAreaPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_global_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_global_global_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_global_global_proto_goTypes,
		DependencyIndexes: file_global_global_proto_depIdxs,
		MessageInfos:      file_global_global_proto_msgTypes,
	}.Build()
	File_global_global_proto = out.File
	file_global_global_proto_rawDesc = nil
	file_global_global_proto_goTypes = nil
	file_global_global_proto_depIdxs = nil
}
