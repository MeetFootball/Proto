// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: we_chat/global.proto

package we_chat

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EmptyPost struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyPost) Reset()         { *m = EmptyPost{} }
func (m *EmptyPost) String() string { return proto.CompactTextString(m) }
func (*EmptyPost) ProtoMessage()    {}
func (*EmptyPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{0}
}
func (m *EmptyPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyPost.Unmarshal(m, b)
}
func (m *EmptyPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyPost.Marshal(b, m, deterministic)
}
func (m *EmptyPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyPost.Merge(m, src)
}
func (m *EmptyPost) XXX_Size() int {
	return xxx_messageInfo_EmptyPost.Size(m)
}
func (m *EmptyPost) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyPost.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyPost proto.InternalMessageInfo

type InfoPost struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoPost) Reset()         { *m = InfoPost{} }
func (m *InfoPost) String() string { return proto.CompactTextString(m) }
func (*InfoPost) ProtoMessage()    {}
func (*InfoPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{1}
}
func (m *InfoPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoPost.Unmarshal(m, b)
}
func (m *InfoPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoPost.Marshal(b, m, deterministic)
}
func (m *InfoPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoPost.Merge(m, src)
}
func (m *InfoPost) XXX_Size() int {
	return xxx_messageInfo_InfoPost.Size(m)
}
func (m *InfoPost) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoPost.DiscardUnknown(m)
}

var xxx_messageInfo_InfoPost proto.InternalMessageInfo

func (m *InfoPost) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type KeyPost struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyPost) Reset()         { *m = KeyPost{} }
func (m *KeyPost) String() string { return proto.CompactTextString(m) }
func (*KeyPost) ProtoMessage()    {}
func (*KeyPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{2}
}
func (m *KeyPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyPost.Unmarshal(m, b)
}
func (m *KeyPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyPost.Marshal(b, m, deterministic)
}
func (m *KeyPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPost.Merge(m, src)
}
func (m *KeyPost) XXX_Size() int {
	return xxx_messageInfo_KeyPost.Size(m)
}
func (m *KeyPost) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPost.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPost proto.InternalMessageInfo

func (m *KeyPost) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type KeywordPost struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=Keyword,proto3" json:"keyword"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPost) Reset()         { *m = KeywordPost{} }
func (m *KeywordPost) String() string { return proto.CompactTextString(m) }
func (*KeywordPost) ProtoMessage()    {}
func (*KeywordPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{3}
}
func (m *KeywordPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPost.Unmarshal(m, b)
}
func (m *KeywordPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPost.Marshal(b, m, deterministic)
}
func (m *KeywordPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPost.Merge(m, src)
}
func (m *KeywordPost) XXX_Size() int {
	return xxx_messageInfo_KeywordPost.Size(m)
}
func (m *KeywordPost) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPost.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPost proto.InternalMessageInfo

func (m *KeywordPost) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type DatePost struct {
	Start                string   `protobuf:"bytes,1,opt,name=Start,proto3" json:"start"`
	End                  string   `protobuf:"bytes,2,opt,name=End,proto3" json:"end"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatePost) Reset()         { *m = DatePost{} }
func (m *DatePost) String() string { return proto.CompactTextString(m) }
func (*DatePost) ProtoMessage()    {}
func (*DatePost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{4}
}
func (m *DatePost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatePost.Unmarshal(m, b)
}
func (m *DatePost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatePost.Marshal(b, m, deterministic)
}
func (m *DatePost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatePost.Merge(m, src)
}
func (m *DatePost) XXX_Size() int {
	return xxx_messageInfo_DatePost.Size(m)
}
func (m *DatePost) XXX_DiscardUnknown() {
	xxx_messageInfo_DatePost.DiscardUnknown(m)
}

var xxx_messageInfo_DatePost proto.InternalMessageInfo

func (m *DatePost) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *DatePost) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

// PaginationPost ???????????? ??????
type PaginationPost struct {
	Size_                uint64   `protobuf:"varint,1,opt,name=Size,proto3" json:"size"`
	Page                 uint64   `protobuf:"varint,2,opt,name=Page,proto3" json:"page"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaginationPost) Reset()         { *m = PaginationPost{} }
func (m *PaginationPost) String() string { return proto.CompactTextString(m) }
func (*PaginationPost) ProtoMessage()    {}
func (*PaginationPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{5}
}
func (m *PaginationPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaginationPost.Unmarshal(m, b)
}
func (m *PaginationPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaginationPost.Marshal(b, m, deterministic)
}
func (m *PaginationPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaginationPost.Merge(m, src)
}
func (m *PaginationPost) XXX_Size() int {
	return xxx_messageInfo_PaginationPost.Size(m)
}
func (m *PaginationPost) XXX_DiscardUnknown() {
	xxx_messageInfo_PaginationPost.DiscardUnknown(m)
}

var xxx_messageInfo_PaginationPost proto.InternalMessageInfo

func (m *PaginationPost) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *PaginationPost) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

// SearchPaginationPost ???????????? ??????
type SearchPaginationPost struct {
	Size_                uint64   `protobuf:"varint,1,opt,name=Size,proto3" json:"size"`
	Page                 uint64   `protobuf:"varint,2,opt,name=Page,proto3" json:"page"`
	Keyword              string   `protobuf:"bytes,3,opt,name=Keyword,proto3" json:"keyword"`
	Status               uint64   `protobuf:"varint,4,opt,name=Status,proto3" json:"status"`
	Start                string   `protobuf:"bytes,5,opt,name=Start,proto3" json:"start"`
	End                  string   `protobuf:"bytes,6,opt,name=End,proto3" json:"end"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPaginationPost) Reset()         { *m = SearchPaginationPost{} }
func (m *SearchPaginationPost) String() string { return proto.CompactTextString(m) }
func (*SearchPaginationPost) ProtoMessage()    {}
func (*SearchPaginationPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{6}
}
func (m *SearchPaginationPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPaginationPost.Unmarshal(m, b)
}
func (m *SearchPaginationPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPaginationPost.Marshal(b, m, deterministic)
}
func (m *SearchPaginationPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPaginationPost.Merge(m, src)
}
func (m *SearchPaginationPost) XXX_Size() int {
	return xxx_messageInfo_SearchPaginationPost.Size(m)
}
func (m *SearchPaginationPost) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPaginationPost.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPaginationPost proto.InternalMessageInfo

func (m *SearchPaginationPost) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *SearchPaginationPost) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchPaginationPost) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *SearchPaginationPost) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SearchPaginationPost) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *SearchPaginationPost) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

// ChangeStatusPost ?????????????????? ??????
type ChangeStatusPost struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	Status               uint64   `protobuf:"varint,2,opt,name=Status,proto3" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusPost) Reset()         { *m = ChangeStatusPost{} }
func (m *ChangeStatusPost) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusPost) ProtoMessage()    {}
func (*ChangeStatusPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{7}
}
func (m *ChangeStatusPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusPost.Unmarshal(m, b)
}
func (m *ChangeStatusPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusPost.Marshal(b, m, deterministic)
}
func (m *ChangeStatusPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusPost.Merge(m, src)
}
func (m *ChangeStatusPost) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusPost.Size(m)
}
func (m *ChangeStatusPost) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusPost.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusPost proto.InternalMessageInfo

func (m *ChangeStatusPost) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ChangeStatusPost) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ChangeStatusByKeyPost struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"key"`
	Status               uint64   `protobuf:"varint,2,opt,name=Status,proto3" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusByKeyPost) Reset()         { *m = ChangeStatusByKeyPost{} }
func (m *ChangeStatusByKeyPost) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusByKeyPost) ProtoMessage()    {}
func (*ChangeStatusByKeyPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{8}
}
func (m *ChangeStatusByKeyPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusByKeyPost.Unmarshal(m, b)
}
func (m *ChangeStatusByKeyPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusByKeyPost.Marshal(b, m, deterministic)
}
func (m *ChangeStatusByKeyPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusByKeyPost.Merge(m, src)
}
func (m *ChangeStatusByKeyPost) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusByKeyPost.Size(m)
}
func (m *ChangeStatusByKeyPost) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusByKeyPost.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusByKeyPost proto.InternalMessageInfo

func (m *ChangeStatusByKeyPost) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ChangeStatusByKeyPost) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

// UpdateAreaPost ??????????????????
type UpdateAreaPost struct {
	CountryCode          uint64   `protobuf:"varint,1,opt,name=CountryCode,proto3" json:"country_code"`
	ProvinceCode         uint64   `protobuf:"varint,2,opt,name=ProvinceCode,proto3" json:"province_code"`
	CityCode             uint64   `protobuf:"varint,3,opt,name=CityCode,proto3" json:"city_code"`
	Country              string   `protobuf:"bytes,4,opt,name=Country,proto3" json:"country"`
	Province             string   `protobuf:"bytes,5,opt,name=Province,proto3" json:"province"`
	City                 string   `protobuf:"bytes,6,opt,name=City,proto3" json:"city"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAreaPost) Reset()         { *m = UpdateAreaPost{} }
func (m *UpdateAreaPost) String() string { return proto.CompactTextString(m) }
func (*UpdateAreaPost) ProtoMessage()    {}
func (*UpdateAreaPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{9}
}
func (m *UpdateAreaPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAreaPost.Unmarshal(m, b)
}
func (m *UpdateAreaPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAreaPost.Marshal(b, m, deterministic)
}
func (m *UpdateAreaPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAreaPost.Merge(m, src)
}
func (m *UpdateAreaPost) XXX_Size() int {
	return xxx_messageInfo_UpdateAreaPost.Size(m)
}
func (m *UpdateAreaPost) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAreaPost.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAreaPost proto.InternalMessageInfo

func (m *UpdateAreaPost) GetCountryCode() uint64 {
	if m != nil {
		return m.CountryCode
	}
	return 0
}

func (m *UpdateAreaPost) GetProvinceCode() uint64 {
	if m != nil {
		return m.ProvinceCode
	}
	return 0
}

func (m *UpdateAreaPost) GetCityCode() uint64 {
	if m != nil {
		return m.CityCode
	}
	return 0
}

func (m *UpdateAreaPost) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *UpdateAreaPost) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *UpdateAreaPost) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type StringResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"message"`
	Data                 string   `protobuf:"bytes,2,opt,name=Data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringResponse) Reset()         { *m = StringResponse{} }
func (m *StringResponse) String() string { return proto.CompactTextString(m) }
func (*StringResponse) ProtoMessage()    {}
func (*StringResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{10}
}
func (m *StringResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringResponse.Unmarshal(m, b)
}
func (m *StringResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringResponse.Marshal(b, m, deterministic)
}
func (m *StringResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringResponse.Merge(m, src)
}
func (m *StringResponse) XXX_Size() int {
	return xxx_messageInfo_StringResponse.Size(m)
}
func (m *StringResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StringResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StringResponse proto.InternalMessageInfo

func (m *StringResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StringResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// BoolResponse ?????? Bool
type BoolResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"message"`
	Data                 bool     `protobuf:"varint,2,opt,name=Data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolResponse) Reset()         { *m = BoolResponse{} }
func (m *BoolResponse) String() string { return proto.CompactTextString(m) }
func (*BoolResponse) ProtoMessage()    {}
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{11}
}
func (m *BoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolResponse.Unmarshal(m, b)
}
func (m *BoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolResponse.Marshal(b, m, deterministic)
}
func (m *BoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolResponse.Merge(m, src)
}
func (m *BoolResponse) XXX_Size() int {
	return xxx_messageInfo_BoolResponse.Size(m)
}
func (m *BoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BoolResponse proto.InternalMessageInfo

func (m *BoolResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BoolResponse) GetData() bool {
	if m != nil {
		return m.Data
	}
	return false
}

type CheckResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"message"`
	Data                 bool     `protobuf:"varint,2,opt,name=Data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{12}
}
func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CheckResponse) GetData() bool {
	if m != nil {
		return m.Data
	}
	return false
}

type ResultResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"message"`
	Data                 bool     `protobuf:"varint,2,opt,name=Data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultResponse) Reset()         { *m = ResultResponse{} }
func (m *ResultResponse) String() string { return proto.CompactTextString(m) }
func (*ResultResponse) ProtoMessage()    {}
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7150585e4c50cc, []int{13}
}
func (m *ResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultResponse.Unmarshal(m, b)
}
func (m *ResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultResponse.Marshal(b, m, deterministic)
}
func (m *ResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultResponse.Merge(m, src)
}
func (m *ResultResponse) XXX_Size() int {
	return xxx_messageInfo_ResultResponse.Size(m)
}
func (m *ResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResultResponse proto.InternalMessageInfo

func (m *ResultResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResultResponse) GetData() bool {
	if m != nil {
		return m.Data
	}
	return false
}

func init() {
	proto.RegisterType((*EmptyPost)(nil), "we_chat.EmptyPost")
	proto.RegisterType((*InfoPost)(nil), "we_chat.InfoPost")
	proto.RegisterType((*KeyPost)(nil), "we_chat.KeyPost")
	proto.RegisterType((*KeywordPost)(nil), "we_chat.KeywordPost")
	proto.RegisterType((*DatePost)(nil), "we_chat.DatePost")
	proto.RegisterType((*PaginationPost)(nil), "we_chat.PaginationPost")
	proto.RegisterType((*SearchPaginationPost)(nil), "we_chat.SearchPaginationPost")
	proto.RegisterType((*ChangeStatusPost)(nil), "we_chat.ChangeStatusPost")
	proto.RegisterType((*ChangeStatusByKeyPost)(nil), "we_chat.ChangeStatusByKeyPost")
	proto.RegisterType((*UpdateAreaPost)(nil), "we_chat.UpdateAreaPost")
	proto.RegisterType((*StringResponse)(nil), "we_chat.StringResponse")
	proto.RegisterType((*BoolResponse)(nil), "we_chat.BoolResponse")
	proto.RegisterType((*CheckResponse)(nil), "we_chat.CheckResponse")
	proto.RegisterType((*ResultResponse)(nil), "we_chat.ResultResponse")
}

func init() { proto.RegisterFile("we_chat/global.proto", fileDescriptor_cd7150585e4c50cc) }

var fileDescriptor_cd7150585e4c50cc = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x9c, 0xc4, 0x71, 0x26, 0x3f, 0x6a, 0xad, 0x82, 0x02, 0x02, 0x19, 0xad, 0x40, 0x2a,
	0x97, 0x46, 0xe2, 0xe7, 0xc4, 0x09, 0x27, 0x45, 0x8a, 0x02, 0x28, 0x5a, 0x13, 0x0e, 0x5c, 0xaa,
	0xad, 0xbd, 0x38, 0x56, 0x52, 0xaf, 0x65, 0x6f, 0xa8, 0xdc, 0xf7, 0xe0, 0xe1, 0xb8, 0xf8, 0x01,
	0xfc, 0x14, 0x68, 0xc7, 0xeb, 0x36, 0x08, 0x5a, 0x90, 0xc8, 0xc9, 0xf6, 0xf7, 0xcd, 0x7c, 0xdf,
	0x78, 0x66, 0x76, 0xe1, 0xe8, 0x92, 0x9f, 0xf9, 0x2b, 0x26, 0xc7, 0xe1, 0x46, 0x9c, 0xb3, 0xcd,
	0x49, 0x92, 0x0a, 0x29, 0xec, 0x8e, 0x46, 0x1f, 0xf6, 0x42, 0x31, 0x0e, 0x45, 0x85, 0x92, 0x1e,
	0x74, 0x4f, 0x2f, 0x12, 0x99, 0x2f, 0x44, 0x26, 0x09, 0x01, 0x6b, 0x16, 0x7f, 0x15, 0xea, 0xdd,
	0xbe, 0x0f, 0xc6, 0x6c, 0x3a, 0x6a, 0x3c, 0x69, 0x1c, 0xb7, 0x5c, 0xb3, 0x2c, 0x1c, 0x23, 0x0a,
	0xa8, 0x31, 0x9b, 0x92, 0xa7, 0xd0, 0x99, 0x73, 0x0c, 0xb7, 0x1f, 0x40, 0x73, 0xce, 0x73, 0x8c,
	0xe9, 0xba, 0x9d, 0xb2, 0x70, 0x9a, 0x6b, 0x9e, 0x53, 0x85, 0x91, 0x57, 0xd0, 0x9b, 0xf3, 0xfc,
	0x52, 0xa4, 0x01, 0x46, 0x3e, 0xc3, 0x24, 0xf5, 0xa9, 0xa3, 0x7b, 0x65, 0xe1, 0x74, 0xd6, 0x15,
	0x44, 0x6b, 0x8e, 0xbc, 0x03, 0x6b, 0xca, 0x24, 0xc7, 0x14, 0x07, 0xda, 0x9e, 0x64, 0xa9, 0xd4,
	0x09, 0xdd, 0xb2, 0x70, 0xda, 0x99, 0x02, 0x68, 0x85, 0x2b, 0xf7, 0xd3, 0x38, 0x18, 0x19, 0x37,
	0xee, 0x3c, 0x0e, 0xa8, 0xc2, 0xc8, 0x7b, 0x18, 0x2e, 0x58, 0x18, 0xc5, 0x4c, 0x46, 0x22, 0x46,
	0xb5, 0x47, 0xd0, 0xf2, 0xa2, 0x2b, 0xae, 0xff, 0xc7, 0x2a, 0x0b, 0xa7, 0x95, 0x45, 0x57, 0x9c,
	0x22, 0xaa, 0xd8, 0x05, 0x0b, 0x39, 0x6a, 0x69, 0x36, 0x61, 0x21, 0xa7, 0x88, 0x92, 0x1f, 0x0d,
	0x38, 0xf2, 0x38, 0x4b, 0xfd, 0xd5, 0xfe, 0x44, 0x77, 0x3b, 0xd2, 0xbc, 0xbd, 0x23, 0x36, 0x01,
	0xd3, 0x93, 0x4c, 0x6e, 0xb3, 0x51, 0x0b, 0x65, 0xa0, 0x2c, 0x1c, 0x33, 0x43, 0x84, 0x6a, 0xe6,
	0xa6, 0x53, 0xed, 0xbb, 0x3b, 0x65, 0xfe, 0xa1, 0x53, 0x1f, 0xe1, 0x60, 0xb2, 0x62, 0x71, 0xc8,
	0x2b, 0xad, 0xbb, 0x26, 0xbf, 0x53, 0x8b, 0x71, 0x5b, 0x2d, 0xe4, 0x33, 0xdc, 0xdb, 0xd5, 0x73,
	0xf3, 0xbf, 0xef, 0xca, 0x3f, 0xe9, 0x7e, 0x37, 0x60, 0xb8, 0x4c, 0x02, 0x26, 0xf9, 0xdb, 0x94,
	0x33, 0x54, 0x7c, 0x01, 0xbd, 0x89, 0xd8, 0xc6, 0x32, 0xcd, 0x27, 0x22, 0xa8, 0x87, 0x70, 0x50,
	0x16, 0x4e, 0xdf, 0xaf, 0xe0, 0x33, 0x5f, 0x04, 0x9c, 0xee, 0x06, 0xd9, 0xaf, 0xa1, 0xbf, 0x48,
	0xc5, 0xb7, 0x28, 0xf6, 0x39, 0x26, 0x55, 0x86, 0x87, 0x65, 0xe1, 0x0c, 0x12, 0x8d, 0x57, 0x59,
	0xbf, 0x84, 0xd9, 0xcf, 0xc1, 0x9a, 0x44, 0xb2, 0xf2, 0x69, 0x62, 0xca, 0xa0, 0x2c, 0x9c, 0xae,
	0x1f, 0x49, 0x6d, 0x72, 0x4d, 0xab, 0xb9, 0x6a, 0x43, 0x9c, 0x98, 0x9e, 0xab, 0xae, 0x88, 0xd6,
	0x9c, 0x7d, 0x0c, 0x56, 0xed, 0xa0, 0xc7, 0xd6, 0x2f, 0x0b, 0xc7, 0xaa, 0x8b, 0xa0, 0xd7, 0xac,
	0x5a, 0x23, 0x25, 0xae, 0xa7, 0x87, 0x6b, 0xa4, 0x7c, 0x29, 0xa2, 0x64, 0x09, 0x43, 0x4f, 0xa6,
	0x51, 0x1c, 0x52, 0x9e, 0x25, 0x22, 0xce, 0xb0, 0x80, 0x0f, 0x3c, 0xcb, 0xd4, 0xe6, 0xed, 0x1c,
	0xb5, 0x8b, 0x0a, 0xa2, 0x35, 0xa7, 0x64, 0xa7, 0x4c, 0x32, 0x7d, 0x7c, 0x50, 0x36, 0x60, 0x92,
	0x51, 0x44, 0x89, 0x07, 0x7d, 0x57, 0x88, 0xcd, 0xff, 0x88, 0x5a, 0xbf, 0x89, 0x7e, 0x82, 0xc1,
	0x64, 0xc5, 0xfd, 0xf5, 0x7e, 0x55, 0x97, 0x30, 0xa4, 0x3c, 0xdb, 0x6e, 0xe4, 0x5e, 0x65, 0xdd,
	0xc7, 0xae, 0x59, 0xdd, 0x9e, 0x8b, 0xc6, 0x97, 0xc3, 0x93, 0xb1, 0xbe, 0x3b, 0xdf, 0xe8, 0xe7,
	0xb9, 0x89, 0xb7, 0xe7, 0xcb, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x22, 0xa7, 0x3d, 0x6b,
	0x05, 0x00, 0x00,
}
