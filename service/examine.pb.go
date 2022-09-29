// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/examine.proto

package service

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

// ExamineText 文字审核结果
type ExamineText struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result"`
	Words                []string `protobuf:"bytes,2,rep,name=words,proto3" json:"words"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExamineText) Reset()         { *m = ExamineText{} }
func (m *ExamineText) String() string { return proto.CompactTextString(m) }
func (*ExamineText) ProtoMessage()    {}
func (*ExamineText) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa288c9b7c6699e, []int{0}
}
func (m *ExamineText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamineText.Unmarshal(m, b)
}
func (m *ExamineText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamineText.Marshal(b, m, deterministic)
}
func (m *ExamineText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamineText.Merge(m, src)
}
func (m *ExamineText) XXX_Size() int {
	return xxx_messageInfo_ExamineText.Size(m)
}
func (m *ExamineText) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamineText.DiscardUnknown(m)
}

var xxx_messageInfo_ExamineText proto.InternalMessageInfo

func (m *ExamineText) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *ExamineText) GetWords() []string {
	if m != nil {
		return m.Words
	}
	return nil
}

// ExamineImage 图片审核结果
type ExamineImage struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExamineImage) Reset()         { *m = ExamineImage{} }
func (m *ExamineImage) String() string { return proto.CompactTextString(m) }
func (*ExamineImage) ProtoMessage()    {}
func (*ExamineImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa288c9b7c6699e, []int{1}
}
func (m *ExamineImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamineImage.Unmarshal(m, b)
}
func (m *ExamineImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamineImage.Marshal(b, m, deterministic)
}
func (m *ExamineImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamineImage.Merge(m, src)
}
func (m *ExamineImage) XXX_Size() int {
	return xxx_messageInfo_ExamineImage.Size(m)
}
func (m *ExamineImage) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamineImage.DiscardUnknown(m)
}

var xxx_messageInfo_ExamineImage proto.InternalMessageInfo

func (m *ExamineImage) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *ExamineImage) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// ExamineTextPost 文字审核
type ExamineTextPost struct {
	UUID                 uint64   `protobuf:"varint,1,opt,name=UUID,proto3" json:"uuid"`
	Client               uint64   `protobuf:"varint,2,opt,name=Client,proto3" json:"client"`
	Text                 string   `protobuf:"bytes,3,opt,name=Text,proto3" json:"text"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExamineTextPost) Reset()         { *m = ExamineTextPost{} }
func (m *ExamineTextPost) String() string { return proto.CompactTextString(m) }
func (*ExamineTextPost) ProtoMessage()    {}
func (*ExamineTextPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa288c9b7c6699e, []int{2}
}
func (m *ExamineTextPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamineTextPost.Unmarshal(m, b)
}
func (m *ExamineTextPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamineTextPost.Marshal(b, m, deterministic)
}
func (m *ExamineTextPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamineTextPost.Merge(m, src)
}
func (m *ExamineTextPost) XXX_Size() int {
	return xxx_messageInfo_ExamineTextPost.Size(m)
}
func (m *ExamineTextPost) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamineTextPost.DiscardUnknown(m)
}

var xxx_messageInfo_ExamineTextPost proto.InternalMessageInfo

func (m *ExamineTextPost) GetUUID() uint64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

func (m *ExamineTextPost) GetClient() uint64 {
	if m != nil {
		return m.Client
	}
	return 0
}

func (m *ExamineTextPost) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// ExamineImagePost 图片审核
type ExamineImagePost struct {
	UUID                 uint64   `protobuf:"varint,1,opt,name=UUID,proto3" json:"uuid"`
	Client               uint64   `protobuf:"varint,2,opt,name=Client,proto3" json:"client"`
	Type                 uint64   `protobuf:"varint,3,opt,name=Type,proto3" json:"type"`
	Image                string   `protobuf:"bytes,4,opt,name=Image,proto3" json:"image"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExamineImagePost) Reset()         { *m = ExamineImagePost{} }
func (m *ExamineImagePost) String() string { return proto.CompactTextString(m) }
func (*ExamineImagePost) ProtoMessage()    {}
func (*ExamineImagePost) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa288c9b7c6699e, []int{3}
}
func (m *ExamineImagePost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamineImagePost.Unmarshal(m, b)
}
func (m *ExamineImagePost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamineImagePost.Marshal(b, m, deterministic)
}
func (m *ExamineImagePost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamineImagePost.Merge(m, src)
}
func (m *ExamineImagePost) XXX_Size() int {
	return xxx_messageInfo_ExamineImagePost.Size(m)
}
func (m *ExamineImagePost) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamineImagePost.DiscardUnknown(m)
}

var xxx_messageInfo_ExamineImagePost proto.InternalMessageInfo

func (m *ExamineImagePost) GetUUID() uint64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

func (m *ExamineImagePost) GetClient() uint64 {
	if m != nil {
		return m.Client
	}
	return 0
}

func (m *ExamineImagePost) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ExamineImagePost) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

// ExamineTextResponse 返回 文字审核结果
type ExamineTextResponse struct {
	Message              string       `protobuf:"bytes,1,opt,name=Message,proto3" json:"message"`
	Data                 *ExamineText `protobuf:"bytes,2,opt,name=Data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ExamineTextResponse) Reset()         { *m = ExamineTextResponse{} }
func (m *ExamineTextResponse) String() string { return proto.CompactTextString(m) }
func (*ExamineTextResponse) ProtoMessage()    {}
func (*ExamineTextResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa288c9b7c6699e, []int{4}
}
func (m *ExamineTextResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamineTextResponse.Unmarshal(m, b)
}
func (m *ExamineTextResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamineTextResponse.Marshal(b, m, deterministic)
}
func (m *ExamineTextResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamineTextResponse.Merge(m, src)
}
func (m *ExamineTextResponse) XXX_Size() int {
	return xxx_messageInfo_ExamineTextResponse.Size(m)
}
func (m *ExamineTextResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamineTextResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExamineTextResponse proto.InternalMessageInfo

func (m *ExamineTextResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ExamineTextResponse) GetData() *ExamineText {
	if m != nil {
		return m.Data
	}
	return nil
}

// ExamineImageResponse 返回 审核结果
type ExamineImageResponse struct {
	Message              string        `protobuf:"bytes,1,opt,name=Message,proto3" json:"message"`
	Data                 *ExamineImage `protobuf:"bytes,2,opt,name=Data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExamineImageResponse) Reset()         { *m = ExamineImageResponse{} }
func (m *ExamineImageResponse) String() string { return proto.CompactTextString(m) }
func (*ExamineImageResponse) ProtoMessage()    {}
func (*ExamineImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa288c9b7c6699e, []int{5}
}
func (m *ExamineImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamineImageResponse.Unmarshal(m, b)
}
func (m *ExamineImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamineImageResponse.Marshal(b, m, deterministic)
}
func (m *ExamineImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamineImageResponse.Merge(m, src)
}
func (m *ExamineImageResponse) XXX_Size() int {
	return xxx_messageInfo_ExamineImageResponse.Size(m)
}
func (m *ExamineImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamineImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExamineImageResponse proto.InternalMessageInfo

func (m *ExamineImageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ExamineImageResponse) GetData() *ExamineImage {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ExamineText)(nil), "service.ExamineText")
	proto.RegisterType((*ExamineImage)(nil), "service.ExamineImage")
	proto.RegisterType((*ExamineTextPost)(nil), "service.ExamineTextPost")
	proto.RegisterType((*ExamineImagePost)(nil), "service.ExamineImagePost")
	proto.RegisterType((*ExamineTextResponse)(nil), "service.ExamineTextResponse")
	proto.RegisterType((*ExamineImageResponse)(nil), "service.ExamineImageResponse")
}

func init() { proto.RegisterFile("service/examine.proto", fileDescriptor_5fa288c9b7c6699e) }

var fileDescriptor_5fa288c9b7c6699e = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xc7, 0x49, 0x9b, 0x5f, 0xff, 0x4c, 0x7e, 0xa0, 0xa6, 0x2d, 0x04, 0x11, 0x52, 0x16, 0x84,
	0x9e, 0x52, 0x68, 0x8f, 0xde, 0x62, 0x3d, 0xf4, 0x20, 0xc8, 0x6a, 0x2f, 0xde, 0x62, 0x3b, 0x94,
	0x40, 0xdb, 0x8d, 0xbb, 0x5b, 0x4d, 0x9f, 0xc3, 0xf7, 0xcb, 0x03, 0xe4, 0x29, 0x24, 0xb3, 0x5b,
	0x8c, 0x3d, 0x29, 0x78, 0x1a, 0xf8, 0xec, 0xec, 0xe7, 0x3b, 0x3b, 0x2c, 0x0c, 0x14, 0xca, 0xb7,
	0x74, 0x89, 0x63, 0xcc, 0x93, 0x6d, 0xba, 0xc3, 0x28, 0x93, 0x42, 0x0b, 0xbf, 0x6d, 0xf1, 0xa5,
	0xb7, 0x16, 0xe3, 0xb5, 0x30, 0x94, 0x71, 0xf0, 0xee, 0x4c, 0xdb, 0x13, 0xe6, 0xda, 0x67, 0xd0,
	0x92, 0xa8, 0xf6, 0x1b, 0x1d, 0x38, 0x43, 0x67, 0xd4, 0x89, 0xa1, 0x2c, 0x42, 0x4b, 0xb8, 0xad,
	0x7e, 0x08, 0xff, 0xde, 0x85, 0x5c, 0xa9, 0xa0, 0x31, 0x6c, 0x8e, 0xba, 0x71, 0xb7, 0x2c, 0x42,
	0x03, 0xb8, 0x29, 0xec, 0x11, 0xfe, 0x5b, 0xe7, 0x7c, 0x9b, 0xac, 0xf1, 0xa7, 0x52, 0x94, 0x52,
	0xc8, 0xa0, 0x31, 0x74, 0x8e, 0x52, 0x02, 0xdc, 0x14, 0xf6, 0x0a, 0x67, 0xb5, 0x41, 0x1f, 0x84,
	0xd2, 0xfe, 0x15, 0xb8, 0x8b, 0xc5, 0x7c, 0x46, 0x56, 0x37, 0xee, 0x94, 0x45, 0xe8, 0xee, 0xf7,
	0xe9, 0x8a, 0x13, 0xad, 0x52, 0x6f, 0x37, 0x29, 0xee, 0x34, 0x29, 0x5d, 0x93, 0xba, 0x24, 0xc2,
	0xed, 0x49, 0x65, 0xa8, 0x6c, 0x41, 0x93, 0x42, 0xc9, 0xa0, 0x31, 0xd7, 0x9c, 0x28, 0xfb, 0x70,
	0xe0, 0xbc, 0xfe, 0x90, 0x3f, 0x0c, 0x3d, 0x64, 0x48, 0xa1, 0xd6, 0xa0, 0x0f, 0x19, 0x72, 0xa2,
	0xd5, 0x22, 0x28, 0x2c, 0x70, 0xbf, 0x16, 0x91, 0x56, 0x80, 0x1b, 0xce, 0x32, 0xe8, 0xd5, 0x16,
	0xc1, 0x51, 0x65, 0x62, 0xa7, 0xd0, 0xbf, 0x86, 0xf6, 0x3d, 0x2a, 0x55, 0xdd, 0x74, 0xe8, 0xa6,
	0x57, 0x16, 0x61, 0x7b, 0x6b, 0x10, 0x3f, 0x9e, 0xf9, 0x13, 0x70, 0x67, 0x89, 0x4e, 0x68, 0x3c,
	0x6f, 0xd2, 0x8f, 0xec, 0xa7, 0x88, 0x6a, 0x4a, 0x33, 0xd2, 0x2a, 0xd1, 0x09, 0xa7, 0x5e, 0x26,
	0xa1, 0x5f, 0x5f, 0xc3, 0x6f, 0x23, 0xa7, 0xdf, 0x22, 0x07, 0xa7, 0x91, 0xe4, 0x3c, 0xcd, 0x8c,
	0x7b, 0xcf, 0x17, 0xd1, 0xd8, 0x76, 0xde, 0xd8, 0xfa, 0xd2, 0xa2, 0x3f, 0x3b, 0xfd, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0x8d, 0x73, 0x32, 0xe2, 0x02, 0x00, 0x00,
}
