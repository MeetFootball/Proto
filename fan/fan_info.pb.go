// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fan/fan_info.proto

package fan

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

// FanInfo 粉丝信息
type FanInfo struct {
	FanID                uint64   `protobuf:"varint,1,opt,name=FanID,proto3" json:"fan_id"`
	Sex                  uint64   `protobuf:"varint,2,opt,name=Sex,proto3" json:"sex"`
	SecretSex            bool     `protobuf:"varint,3,opt,name=SecretSex,proto3" json:"secret_sex"`
	Age                  uint64   `protobuf:"varint,4,opt,name=Age,proto3" json:"age"`
	Phone                string   `protobuf:"bytes,5,opt,name=Phone,proto3" json:"phone"`
	CountryCode          uint64   `protobuf:"varint,6,opt,name=CountryCode,proto3" json:"country_code"`
	ProvinceCode         uint64   `protobuf:"varint,7,opt,name=ProvinceCode,proto3" json:"province_code"`
	CityCode             uint64   `protobuf:"varint,8,opt,name=CityCode,proto3" json:"city_code"`
	Country              string   `protobuf:"bytes,9,opt,name=Country,proto3" json:"country"`
	Province             string   `protobuf:"bytes,10,opt,name=Province,proto3" json:"province"`
	City                 string   `protobuf:"bytes,11,opt,name=City,proto3" json:"city"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FanInfo) Reset()         { *m = FanInfo{} }
func (m *FanInfo) String() string { return proto.CompactTextString(m) }
func (*FanInfo) ProtoMessage()    {}
func (*FanInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d472f4e7e4272c5b, []int{0}
}
func (m *FanInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FanInfo.Unmarshal(m, b)
}
func (m *FanInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FanInfo.Marshal(b, m, deterministic)
}
func (m *FanInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FanInfo.Merge(m, src)
}
func (m *FanInfo) XXX_Size() int {
	return xxx_messageInfo_FanInfo.Size(m)
}
func (m *FanInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FanInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FanInfo proto.InternalMessageInfo

func (m *FanInfo) GetFanID() uint64 {
	if m != nil {
		return m.FanID
	}
	return 0
}

func (m *FanInfo) GetSex() uint64 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *FanInfo) GetSecretSex() bool {
	if m != nil {
		return m.SecretSex
	}
	return false
}

func (m *FanInfo) GetAge() uint64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *FanInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *FanInfo) GetCountryCode() uint64 {
	if m != nil {
		return m.CountryCode
	}
	return 0
}

func (m *FanInfo) GetProvinceCode() uint64 {
	if m != nil {
		return m.ProvinceCode
	}
	return 0
}

func (m *FanInfo) GetCityCode() uint64 {
	if m != nil {
		return m.CityCode
	}
	return 0
}

func (m *FanInfo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *FanInfo) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *FanInfo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

// ChangeSexPost 修改性别
type ChangeSexPost struct {
	Sex                  uint64   `protobuf:"varint,1,opt,name=Sex,proto3" json:"sex"`
	SecretSex            bool     `protobuf:"varint,2,opt,name=SecretSex,proto3" json:"secret_sex"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeSexPost) Reset()         { *m = ChangeSexPost{} }
func (m *ChangeSexPost) String() string { return proto.CompactTextString(m) }
func (*ChangeSexPost) ProtoMessage()    {}
func (*ChangeSexPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_d472f4e7e4272c5b, []int{1}
}
func (m *ChangeSexPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeSexPost.Unmarshal(m, b)
}
func (m *ChangeSexPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeSexPost.Marshal(b, m, deterministic)
}
func (m *ChangeSexPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeSexPost.Merge(m, src)
}
func (m *ChangeSexPost) XXX_Size() int {
	return xxx_messageInfo_ChangeSexPost.Size(m)
}
func (m *ChangeSexPost) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeSexPost.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeSexPost proto.InternalMessageInfo

func (m *ChangeSexPost) GetSex() uint64 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *ChangeSexPost) GetSecretSex() bool {
	if m != nil {
		return m.SecretSex
	}
	return false
}

// FanInfoResponse 返回粉丝信息
type FanInfoResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"message"`
	Data                 *FanInfo `protobuf:"bytes,2,opt,name=Data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FanInfoResponse) Reset()         { *m = FanInfoResponse{} }
func (m *FanInfoResponse) String() string { return proto.CompactTextString(m) }
func (*FanInfoResponse) ProtoMessage()    {}
func (*FanInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d472f4e7e4272c5b, []int{2}
}
func (m *FanInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FanInfoResponse.Unmarshal(m, b)
}
func (m *FanInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FanInfoResponse.Marshal(b, m, deterministic)
}
func (m *FanInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FanInfoResponse.Merge(m, src)
}
func (m *FanInfoResponse) XXX_Size() int {
	return xxx_messageInfo_FanInfoResponse.Size(m)
}
func (m *FanInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FanInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FanInfoResponse proto.InternalMessageInfo

func (m *FanInfoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FanInfoResponse) GetData() *FanInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*FanInfo)(nil), "fan.FanInfo")
	proto.RegisterType((*ChangeSexPost)(nil), "fan.ChangeSexPost")
	proto.RegisterType((*FanInfoResponse)(nil), "fan.FanInfoResponse")
}

func init() { proto.RegisterFile("fan/fan_info.proto", fileDescriptor_d472f4e7e4272c5b) }

var fileDescriptor_d472f4e7e4272c5b = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x8a, 0xdb, 0x30,
	0x14, 0x85, 0xf1, 0xd8, 0x89, 0xed, 0xeb, 0xa4, 0x3f, 0x5a, 0xa9, 0xa5, 0xa0, 0x10, 0x28, 0xa4,
	0xa5, 0x64, 0x60, 0x4a, 0x57, 0x5d, 0x35, 0x19, 0x0a, 0x5d, 0x14, 0x82, 0xb2, 0x29, 0xdd, 0x04,
	0xd5, 0x96, 0x3d, 0x5e, 0x54, 0x32, 0xb6, 0x5a, 0xd2, 0x07, 0xeb, 0xeb, 0xe8, 0x01, 0xf4, 0x14,
	0x45, 0xd7, 0xf2, 0x4c, 0x67, 0x95, 0x9d, 0xf9, 0xce, 0xb9, 0xe7, 0x5c, 0xa3, 0x0b, 0xa4, 0x16,
	0xea, 0xba, 0x16, 0xea, 0xd4, 0xaa, 0x5a, 0x6f, 0xbb, 0x5e, 0x1b, 0x4d, 0xe2, 0x5a, 0xa8, 0x97,
	0x45, 0xa3, 0xaf, 0x9b, 0x40, 0xd6, 0x7f, 0x63, 0x48, 0x3f, 0x0b, 0xf5, 0x45, 0xd5, 0x9a, 0xac,
	0x60, 0xe6, 0x3f, 0x6f, 0x69, 0xb4, 0x8a, 0x36, 0xc9, 0x0e, 0x9c, 0x65, 0x73, 0x0c, 0xa8, 0xf8,
	0x28, 0x90, 0x17, 0x10, 0x1f, 0xe5, 0x99, 0x5e, 0xa1, 0x9e, 0x3a, 0xcb, 0xe2, 0x41, 0x9e, 0xb9,
	0x67, 0xe4, 0x1d, 0xe4, 0x47, 0x59, 0xf6, 0xd2, 0x78, 0x43, 0xbc, 0x8a, 0x36, 0xd9, 0xee, 0x89,
	0xb3, 0x0c, 0x06, 0x84, 0x27, 0xef, 0x7b, 0x30, 0xf8, 0xa0, 0x4f, 0x8d, 0xa4, 0xc9, 0x43, 0x90,
	0x68, 0x24, 0xf7, 0x8c, 0x30, 0x98, 0x1d, 0xee, 0xb4, 0x92, 0x74, 0xb6, 0x8a, 0x36, 0xf9, 0x2e,
	0x77, 0x96, 0xcd, 0x3a, 0x0f, 0xf8, 0xc8, 0xc9, 0x0d, 0x14, 0x7b, 0xfd, 0x4b, 0x99, 0xfe, 0xcf,
	0x5e, 0x57, 0x92, 0xce, 0x31, 0xe3, 0x99, 0xb3, 0x6c, 0x51, 0x8e, 0xf8, 0x54, 0xea, 0x4a, 0xf2,
	0xff, 0x4d, 0xe4, 0x03, 0x2c, 0x0e, 0xbd, 0xfe, 0xdd, 0xaa, 0x52, 0xe2, 0x50, 0x8a, 0x43, 0xcf,
	0x9d, 0x65, 0xcb, 0x2e, 0xf0, 0x71, 0xea, 0x91, 0x8d, 0xbc, 0x81, 0x6c, 0xdf, 0x9a, 0xb1, 0x27,
	0xc3, 0x91, 0xa5, 0xb3, 0x2c, 0x2f, 0x5b, 0x13, 0x4a, 0xee, 0x65, 0xf2, 0x1a, 0xd2, 0x50, 0x48,
	0x73, 0x5c, 0xbc, 0x70, 0x96, 0xa5, 0x61, 0x23, 0x3e, 0x69, 0x64, 0x03, 0xd9, 0xd4, 0x40, 0x01,
	0x7d, 0x0b, 0x67, 0x59, 0x36, 0x2d, 0xc1, 0xef, 0x55, 0xf2, 0x0a, 0x12, 0x1f, 0x4e, 0x0b, 0x74,
	0x65, 0xce, 0xb2, 0xc4, 0xf7, 0x72, 0xa4, 0xeb, 0x6f, 0xb0, 0xdc, 0xdf, 0x09, 0xd5, 0xc8, 0xa3,
	0x3c, 0x1f, 0xf4, 0x60, 0xa6, 0xa7, 0x89, 0x2e, 0x3d, 0xcd, 0xd5, 0x85, 0xa7, 0x59, 0x57, 0xf0,
	0x34, 0x1c, 0x04, 0x97, 0x43, 0xa7, 0xd5, 0x80, 0xff, 0xf6, 0x55, 0x0e, 0x83, 0x68, 0x24, 0xe6,
	0x87, 0x7f, 0xfb, 0x39, 0x22, 0x3e, 0x69, 0xe4, 0x2d, 0x24, 0xb7, 0xc2, 0x08, 0xac, 0x28, 0x6e,
	0x16, 0xdb, 0x5a, 0xa8, 0x6d, 0x88, 0x1a, 0xf7, 0xaf, 0x84, 0x11, 0x1c, 0x3d, 0xbb, 0xe2, 0x7b,
	0xbe, 0xf5, 0xd7, 0xf9, 0xb1, 0x16, 0xea, 0xc7, 0x1c, 0x6f, 0xf1, 0xfd, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0x96, 0x35, 0xc3, 0xb3, 0x02, 0x00, 0x00,
}
