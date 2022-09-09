// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: fan/fan.proto

package fan

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

// Me 我的信息
type Me struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account        string `protobuf:"bytes,1,opt,name=account,proto3" json:"account"`
	Nickname       string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname"`
	Avatar         string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar"`
	Club           string `protobuf:"bytes,4,opt,name=club,proto3" json:"club"`
	Authentication bool   `protobuf:"varint,5,opt,name=authentication,proto3" json:"authentication"`
}

func (x *Me) Reset() {
	*x = Me{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fan_fan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Me) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Me) ProtoMessage() {}

func (x *Me) ProtoReflect() protoreflect.Message {
	mi := &file_fan_fan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Me.ProtoReflect.Descriptor instead.
func (*Me) Descriptor() ([]byte, []int) {
	return file_fan_fan_proto_rawDescGZIP(), []int{0}
}

func (x *Me) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Me) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Me) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Me) GetClub() string {
	if x != nil {
		return x.Club
	}
	return ""
}

func (x *Me) GetAuthentication() bool {
	if x != nil {
		return x.Authentication
	}
	return false
}

// Fan 粉丝信息
type Fan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account        string `protobuf:"bytes,1,opt,name=account,proto3" json:"account"`
	Nickname       string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname"`
	Avatar         string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar"`
	Club           string `protobuf:"bytes,4,opt,name=club,proto3" json:"club"`
	Authentication bool   `protobuf:"varint,5,opt,name=authentication,proto3" json:"authentication"`
}

func (x *Fan) Reset() {
	*x = Fan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fan_fan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fan) ProtoMessage() {}

func (x *Fan) ProtoReflect() protoreflect.Message {
	mi := &file_fan_fan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fan.ProtoReflect.Descriptor instead.
func (*Fan) Descriptor() ([]byte, []int) {
	return file_fan_fan_proto_rawDescGZIP(), []int{1}
}

func (x *Fan) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Fan) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Fan) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Fan) GetClub() string {
	if x != nil {
		return x.Club
	}
	return ""
}

func (x *Fan) GetAuthentication() bool {
	if x != nil {
		return x.Authentication
	}
	return false
}

// NewFanPost 新用户
type NewFanPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WeChatUserID uint64 `protobuf:"varint,1,opt,name=WeChatUserID,proto3" json:"WeChatUserID"`
	Account      string `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account"`
	Avatar       string `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar"`
	Nickname     string `protobuf:"bytes,4,opt,name=Nickname,proto3" json:"Nickname"`
}

func (x *NewFanPost) Reset() {
	*x = NewFanPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fan_fan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewFanPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewFanPost) ProtoMessage() {}

func (x *NewFanPost) ProtoReflect() protoreflect.Message {
	mi := &file_fan_fan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewFanPost.ProtoReflect.Descriptor instead.
func (*NewFanPost) Descriptor() ([]byte, []int) {
	return file_fan_fan_proto_rawDescGZIP(), []int{2}
}

func (x *NewFanPost) GetWeChatUserID() uint64 {
	if x != nil {
		return x.WeChatUserID
	}
	return 0
}

func (x *NewFanPost) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *NewFanPost) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *NewFanPost) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

// ChangeFanNicknamePost 修改昵称
type ChangeFanNicknamePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname string `protobuf:"bytes,1,opt,name=Nickname,proto3" json:"Nickname"`
}

func (x *ChangeFanNicknamePost) Reset() {
	*x = ChangeFanNicknamePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fan_fan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeFanNicknamePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeFanNicknamePost) ProtoMessage() {}

func (x *ChangeFanNicknamePost) ProtoReflect() protoreflect.Message {
	mi := &file_fan_fan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeFanNicknamePost.ProtoReflect.Descriptor instead.
func (*ChangeFanNicknamePost) Descriptor() ([]byte, []int) {
	return file_fan_fan_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeFanNicknamePost) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

// ChangeFanAccountPost 修改账号
type ChangeFanAccountPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account"`
}

func (x *ChangeFanAccountPost) Reset() {
	*x = ChangeFanAccountPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fan_fan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeFanAccountPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeFanAccountPost) ProtoMessage() {}

func (x *ChangeFanAccountPost) ProtoReflect() protoreflect.Message {
	mi := &file_fan_fan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeFanAccountPost.ProtoReflect.Descriptor instead.
func (*ChangeFanAccountPost) Descriptor() ([]byte, []int) {
	return file_fan_fan_proto_rawDescGZIP(), []int{4}
}

func (x *ChangeFanAccountPost) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type MeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *Me    `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *MeResponse) Reset() {
	*x = MeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fan_fan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeResponse) ProtoMessage() {}

func (x *MeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fan_fan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeResponse.ProtoReflect.Descriptor instead.
func (*MeResponse) Descriptor() ([]byte, []int) {
	return file_fan_fan_proto_rawDescGZIP(), []int{5}
}

func (x *MeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MeResponse) GetData() *Me {
	if x != nil {
		return x.Data
	}
	return nil
}

// FanResponse 返回 粉丝信息
type FanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *Fan   `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *FanResponse) Reset() {
	*x = FanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fan_fan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanResponse) ProtoMessage() {}

func (x *FanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fan_fan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanResponse.ProtoReflect.Descriptor instead.
func (*FanResponse) Descriptor() ([]byte, []int) {
	return file_fan_fan_proto_rawDescGZIP(), []int{6}
}

func (x *FanResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FanResponse) GetData() *Fan {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_fan_fan_proto protoreflect.FileDescriptor

var file_fan_fan_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x61, 0x6e, 0x2f, 0x66, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x66, 0x61, 0x6e, 0x1a, 0x10, 0x66, 0x61, 0x6e, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x66, 0x61, 0x6e, 0x2f, 0x66, 0x61, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x66, 0x61, 0x6e, 0x2f,
	0x66, 0x61, 0x6e, 0x5f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e,
	0x01, 0x0a, 0x02, 0x4d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6c, 0x75, 0x62, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x8f, 0x01, 0x0a, 0x03, 0x46, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x75, 0x62, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6c, 0x75, 0x62, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x7e, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x46, 0x61, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x33, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x61, 0x6e, 0x4e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x46, 0x61, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x0a, 0x4d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a,
	0x0b, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x46, 0x61, 0x6e, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xb1, 0x0b, 0x0a, 0x0a, 0x46, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x46, 0x61, 0x6e, 0x12, 0x0f, 0x2e,
	0x66, 0x61, 0x6e, 0x2e, 0x4e, 0x65, 0x77, 0x46, 0x61, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x02, 0x4d, 0x65, 0x12, 0x0e, 0x2e, 0x66, 0x61, 0x6e, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x66, 0x61, 0x6e, 0x2e,
	0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x03,
	0x46, 0x61, 0x6e, 0x12, 0x0c, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x6f, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x66, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x46, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x07, 0x46, 0x61, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0d, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x46, 0x61, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x2e,
	0x66, 0x61, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x61, 0x6e, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x19, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x61, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x78, 0x12, 0x12, 0x2e, 0x66,
	0x61, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x78, 0x50, 0x6f, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x12, 0x13, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61,
	0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x06, 0x4d, 0x79, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x0e, 0x2e, 0x66, 0x61, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e,
	0x2e, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x0a, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x0d, 0x2e,
	0x66, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66,
	0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x07, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x0d, 0x2e, 0x66,
	0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x61,
	0x6e, 0x2e, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x4d, 0x79, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62,
	0x12, 0x0e, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x46, 0x61, 0x6e,
	0x43, 0x6c, 0x75, 0x62, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x43, 0x6c,
	0x75, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x46, 0x61, 0x6e,
	0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f,
	0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x66, 0x61, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x41, 0x72, 0x65,
	0x61, 0x12, 0x13, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x65, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x46, 0x61, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x0d,
	0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x12, 0x0d, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x66, 0x61, 0x6e,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x6f, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a,
	0x6f, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x41, 0x67, 0x72, 0x65, 0x65,
	0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66,
	0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x66, 0x61, 0x6e,
	0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x6f, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x07, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75,
	0x74, 0x12, 0x10, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x66, 0x61,
	0x6e, 0x3b, 0x66, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fan_fan_proto_rawDescOnce sync.Once
	file_fan_fan_proto_rawDescData = file_fan_fan_proto_rawDesc
)

func file_fan_fan_proto_rawDescGZIP() []byte {
	file_fan_fan_proto_rawDescOnce.Do(func() {
		file_fan_fan_proto_rawDescData = protoimpl.X.CompressGZIP(file_fan_fan_proto_rawDescData)
	})
	return file_fan_fan_proto_rawDescData
}

var file_fan_fan_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_fan_fan_proto_goTypes = []interface{}{
	(*Me)(nil),                        // 0: fan.Me
	(*Fan)(nil),                       // 1: fan.Fan
	(*NewFanPost)(nil),                // 2: fan.NewFanPost
	(*ChangeFanNicknamePost)(nil),     // 3: fan.ChangeFanNicknamePost
	(*ChangeFanAccountPost)(nil),      // 4: fan.ChangeFanAccountPost
	(*MeResponse)(nil),                // 5: fan.MeResponse
	(*FanResponse)(nil),               // 6: fan.FanResponse
	(*EmptyPost)(nil),                 // 7: fan.EmptyPost
	(*KeyPost)(nil),                   // 8: fan.KeyPost
	(*InfoPost)(nil),                  // 9: fan.InfoPost
	(*ChangeSexPost)(nil),             // 10: fan.ChangeSexPost
	(*UpdateAreaPost)(nil),            // 11: fan.UpdateAreaPost
	(*FanClubPaginationPost)(nil),     // 12: fan.FanClubPaginationPost
	(*CreateFanClubPost)(nil),         // 13: fan.CreateFanClubPost
	(*InviteFanPost)(nil),             // 14: fan.InviteFanPost
	(*RejectInvitePost)(nil),          // 15: fan.RejectInvitePost
	(*JoinFanClubPost)(nil),           // 16: fan.JoinFanClubPost
	(*LeaveFanClubPost)(nil),          // 17: fan.LeaveFanClubPost
	(*KickOutPost)(nil),               // 18: fan.KickOutPost
	(*BoolResponse)(nil),              // 19: fan.BoolResponse
	(*FanInfoResponse)(nil),           // 20: fan.FanInfoResponse
	(*ClubResponse)(nil),              // 21: fan.ClubResponse
	(*FanClubResponse)(nil),           // 22: fan.FanClubResponse
	(*FanClubPaginationResponse)(nil), // 23: fan.FanClubPaginationResponse
}
var file_fan_fan_proto_depIdxs = []int32{
	0,  // 0: fan.MeResponse.data:type_name -> fan.Me
	1,  // 1: fan.FanResponse.data:type_name -> fan.Fan
	2,  // 2: fan.FanService.NewFan:input_type -> fan.NewFanPost
	7,  // 3: fan.FanService.Me:input_type -> fan.EmptyPost
	8,  // 4: fan.FanService.Fan:input_type -> fan.KeyPost
	7,  // 5: fan.FanService.SelfInfo:input_type -> fan.EmptyPost
	9,  // 6: fan.FanService.FanInfo:input_type -> fan.InfoPost
	3,  // 7: fan.FanService.ChangeNickname:input_type -> fan.ChangeFanNicknamePost
	4,  // 8: fan.FanService.ChangeAccount:input_type -> fan.ChangeFanAccountPost
	10, // 9: fan.FanService.ChangeSex:input_type -> fan.ChangeSexPost
	11, // 10: fan.FanService.UpdateFanArea:input_type -> fan.UpdateAreaPost
	7,  // 11: fan.FanService.MyClub:input_type -> fan.EmptyPost
	9,  // 12: fan.FanService.ChooseClub:input_type -> fan.InfoPost
	9,  // 13: fan.FanService.FanClub:input_type -> fan.InfoPost
	7,  // 14: fan.FanService.MyFanClub:input_type -> fan.EmptyPost
	12, // 15: fan.FanService.Pagination:input_type -> fan.FanClubPaginationPost
	13, // 16: fan.FanService.Create:input_type -> fan.CreateFanClubPost
	7,  // 17: fan.FanService.Transfer:input_type -> fan.EmptyPost
	11, // 18: fan.FanService.UpdateClubArea:input_type -> fan.UpdateAreaPost
	14, // 19: fan.FanService.Invite:input_type -> fan.InviteFanPost
	9,  // 20: fan.FanService.CancelInvite:input_type -> fan.InfoPost
	9,  // 21: fan.FanService.AcceptInvite:input_type -> fan.InfoPost
	15, // 22: fan.FanService.RejectInvite:input_type -> fan.RejectInvitePost
	16, // 23: fan.FanService.Join:input_type -> fan.JoinFanClubPost
	9,  // 24: fan.FanService.CancelJoin:input_type -> fan.InfoPost
	9,  // 25: fan.FanService.AgreeJoin:input_type -> fan.InfoPost
	16, // 26: fan.FanService.RejectJoin:input_type -> fan.JoinFanClubPost
	17, // 27: fan.FanService.Leave:input_type -> fan.LeaveFanClubPost
	18, // 28: fan.FanService.KickOut:input_type -> fan.KickOutPost
	19, // 29: fan.FanService.NewFan:output_type -> fan.BoolResponse
	5,  // 30: fan.FanService.Me:output_type -> fan.MeResponse
	6,  // 31: fan.FanService.Fan:output_type -> fan.FanResponse
	20, // 32: fan.FanService.SelfInfo:output_type -> fan.FanInfoResponse
	20, // 33: fan.FanService.FanInfo:output_type -> fan.FanInfoResponse
	19, // 34: fan.FanService.ChangeNickname:output_type -> fan.BoolResponse
	19, // 35: fan.FanService.ChangeAccount:output_type -> fan.BoolResponse
	19, // 36: fan.FanService.ChangeSex:output_type -> fan.BoolResponse
	19, // 37: fan.FanService.UpdateFanArea:output_type -> fan.BoolResponse
	21, // 38: fan.FanService.MyClub:output_type -> fan.ClubResponse
	19, // 39: fan.FanService.ChooseClub:output_type -> fan.BoolResponse
	22, // 40: fan.FanService.FanClub:output_type -> fan.FanClubResponse
	22, // 41: fan.FanService.MyFanClub:output_type -> fan.FanClubResponse
	23, // 42: fan.FanService.Pagination:output_type -> fan.FanClubPaginationResponse
	22, // 43: fan.FanService.Create:output_type -> fan.FanClubResponse
	19, // 44: fan.FanService.Transfer:output_type -> fan.BoolResponse
	19, // 45: fan.FanService.UpdateClubArea:output_type -> fan.BoolResponse
	19, // 46: fan.FanService.Invite:output_type -> fan.BoolResponse
	19, // 47: fan.FanService.CancelInvite:output_type -> fan.BoolResponse
	19, // 48: fan.FanService.AcceptInvite:output_type -> fan.BoolResponse
	19, // 49: fan.FanService.RejectInvite:output_type -> fan.BoolResponse
	19, // 50: fan.FanService.Join:output_type -> fan.BoolResponse
	19, // 51: fan.FanService.CancelJoin:output_type -> fan.BoolResponse
	19, // 52: fan.FanService.AgreeJoin:output_type -> fan.BoolResponse
	19, // 53: fan.FanService.RejectJoin:output_type -> fan.BoolResponse
	19, // 54: fan.FanService.Leave:output_type -> fan.BoolResponse
	19, // 55: fan.FanService.KickOut:output_type -> fan.BoolResponse
	29, // [29:56] is the sub-list for method output_type
	2,  // [2:29] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_fan_fan_proto_init() }
func file_fan_fan_proto_init() {
	if File_fan_fan_proto != nil {
		return
	}
	file_fan_global_proto_init()
	file_fan_fan_info_proto_init()
	file_fan_fan_club_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fan_fan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Me); i {
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
		file_fan_fan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fan); i {
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
		file_fan_fan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewFanPost); i {
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
		file_fan_fan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeFanNicknamePost); i {
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
		file_fan_fan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeFanAccountPost); i {
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
		file_fan_fan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeResponse); i {
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
		file_fan_fan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanResponse); i {
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
			RawDescriptor: file_fan_fan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fan_fan_proto_goTypes,
		DependencyIndexes: file_fan_fan_proto_depIdxs,
		MessageInfos:      file_fan_fan_proto_msgTypes,
	}.Build()
	File_fan_fan_proto = out.File
	file_fan_fan_proto_rawDesc = nil
	file_fan_fan_proto_goTypes = nil
	file_fan_fan_proto_depIdxs = nil
}