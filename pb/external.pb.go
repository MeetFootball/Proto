// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: external/external.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_external_external_proto protoreflect.FileDescriptor

var file_external_external_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x1a, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xe5, 0x0b, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x08, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x13,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x57, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x57, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x11, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a,
	0x14, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x1e, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x12, 0x1e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x10, 0x2e,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0e, 0x4e, 0x65, 0x77,
	0x73, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x65, 0x77,
	0x73, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x47, 0x72, 0x61,
	0x62, 0x4e, 0x65, 0x77, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x62, 0x4e,
	0x65, 0x77, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x08, 0x47, 0x72, 0x61, 0x62, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x16, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x47, 0x72, 0x61, 0x62, 0x4e, 0x65, 0x77, 0x73, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77,
	0x73, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x54, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x65, 0x77, 0x73, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_external_external_proto_goTypes = []interface{}{
	(*InfoPost)(nil),                       // 0: global.InfoPost
	(*EmptyPost)(nil),                      // 1: global.EmptyPost
	(*SearchPaginationPost)(nil),           // 2: global.SearchPaginationPost
	(*CreateWebsitePost)(nil),              // 3: external.CreateWebsitePost
	(*UpdateWebsitePost)(nil),              // 4: external.UpdateWebsitePost
	(*ChangeStatusPost)(nil),               // 5: global.ChangeStatusPost
	(*SearchWebsiteJobPaginationPost)(nil), // 6: external.SearchWebsiteJobPaginationPost
	(*CreateWebsiteJobPost)(nil),           // 7: external.CreateWebsiteJobPost
	(*UpdateWebsiteJobPost)(nil),           // 8: external.UpdateWebsiteJobPost
	(*SearchNewsPaginationPost)(nil),       // 9: external.SearchNewsPaginationPost
	(*ManualGrabNewsListPost)(nil),         // 10: external.ManualGrabNewsListPost
	(*GrabNewsPost)(nil),                   // 11: external.GrabNewsPost
	(*TranslateNewsPost)(nil),              // 12: external.TranslateNewsPost
	(*TranslateNewsContentPost)(nil),       // 13: external.TranslateNewsContentPost
	(*UpdateNewsPost)(nil),                 // 14: external.UpdateNewsPost
	(*RemoveNewsPost)(nil),                 // 15: external.RemoveNewsPost
	(*WebsiteResponse)(nil),                // 16: external.WebsiteResponse
	(*WebsitesResponse)(nil),               // 17: external.WebsitesResponse
	(*WebsitePaginationResponse)(nil),      // 18: external.WebsitePaginationResponse
	(*WebsiteJobResponse)(nil),             // 19: external.WebsiteJobResponse
	(*WebsiteJobsResponse)(nil),            // 20: external.WebsiteJobsResponse
	(*WebsiteJobPaginationResponse)(nil),   // 21: external.WebsiteJobPaginationResponse
	(*NewsResponse)(nil),                   // 22: external.NewsResponse
	(*NewsPaginationResponse)(nil),         // 23: external.NewsPaginationResponse
	(*Response)(nil),                       // 24: global.Response
}
var file_external_external_proto_depIdxs = []int32{
	0,  // 0: external.ExternalService.Website:input_type -> global.InfoPost
	1,  // 1: external.ExternalService.Websites:input_type -> global.EmptyPost
	2,  // 2: external.ExternalService.WebsitePagination:input_type -> global.SearchPaginationPost
	3,  // 3: external.ExternalService.CreateWebsite:input_type -> external.CreateWebsitePost
	4,  // 4: external.ExternalService.UpdateWebsite:input_type -> external.UpdateWebsitePost
	5,  // 5: external.ExternalService.ChangeWebsiteStatus:input_type -> global.ChangeStatusPost
	0,  // 6: external.ExternalService.WebsiteJob:input_type -> global.InfoPost
	1,  // 7: external.ExternalService.WebsiteJobs:input_type -> global.EmptyPost
	6,  // 8: external.ExternalService.WebsiteJobPagination:input_type -> external.SearchWebsiteJobPaginationPost
	7,  // 9: external.ExternalService.CreateWebsiteJob:input_type -> external.CreateWebsiteJobPost
	8,  // 10: external.ExternalService.UpdateWebsiteJob:input_type -> external.UpdateWebsiteJobPost
	5,  // 11: external.ExternalService.ChangeWebsiteJobStatus:input_type -> global.ChangeStatusPost
	0,  // 12: external.ExternalService.News:input_type -> global.InfoPost
	9,  // 13: external.ExternalService.NewsPagination:input_type -> external.SearchNewsPaginationPost
	10, // 14: external.ExternalService.ManualGrabNewsList:input_type -> external.ManualGrabNewsListPost
	11, // 15: external.ExternalService.GrabNews:input_type -> external.GrabNewsPost
	12, // 16: external.ExternalService.TranslateNews:input_type -> external.TranslateNewsPost
	13, // 17: external.ExternalService.TranslateNewsContent:input_type -> external.TranslateNewsContentPost
	14, // 18: external.ExternalService.UpdateNews:input_type -> external.UpdateNewsPost
	15, // 19: external.ExternalService.RemoveNews:input_type -> external.RemoveNewsPost
	16, // 20: external.ExternalService.Website:output_type -> external.WebsiteResponse
	17, // 21: external.ExternalService.Websites:output_type -> external.WebsitesResponse
	18, // 22: external.ExternalService.WebsitePagination:output_type -> external.WebsitePaginationResponse
	16, // 23: external.ExternalService.CreateWebsite:output_type -> external.WebsiteResponse
	16, // 24: external.ExternalService.UpdateWebsite:output_type -> external.WebsiteResponse
	16, // 25: external.ExternalService.ChangeWebsiteStatus:output_type -> external.WebsiteResponse
	19, // 26: external.ExternalService.WebsiteJob:output_type -> external.WebsiteJobResponse
	20, // 27: external.ExternalService.WebsiteJobs:output_type -> external.WebsiteJobsResponse
	21, // 28: external.ExternalService.WebsiteJobPagination:output_type -> external.WebsiteJobPaginationResponse
	19, // 29: external.ExternalService.CreateWebsiteJob:output_type -> external.WebsiteJobResponse
	19, // 30: external.ExternalService.UpdateWebsiteJob:output_type -> external.WebsiteJobResponse
	19, // 31: external.ExternalService.ChangeWebsiteJobStatus:output_type -> external.WebsiteJobResponse
	22, // 32: external.ExternalService.News:output_type -> external.NewsResponse
	23, // 33: external.ExternalService.NewsPagination:output_type -> external.NewsPaginationResponse
	24, // 34: external.ExternalService.ManualGrabNewsList:output_type -> global.Response
	24, // 35: external.ExternalService.GrabNews:output_type -> global.Response
	22, // 36: external.ExternalService.TranslateNews:output_type -> external.NewsResponse
	22, // 37: external.ExternalService.TranslateNewsContent:output_type -> external.NewsResponse
	22, // 38: external.ExternalService.UpdateNews:output_type -> external.NewsResponse
	22, // 39: external.ExternalService.RemoveNews:output_type -> external.NewsResponse
	20, // [20:40] is the sub-list for method output_type
	0,  // [0:20] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_external_external_proto_init() }
func file_external_external_proto_init() {
	if File_external_external_proto != nil {
		return
	}
	file_global_global_proto_init()
	file_external_website_proto_init()
	file_external_website_job_proto_init()
	file_external_news_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_external_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_external_proto_goTypes,
		DependencyIndexes: file_external_external_proto_depIdxs,
	}.Build()
	File_external_external_proto = out.File
	file_external_external_proto_rawDesc = nil
	file_external_external_proto_goTypes = nil
	file_external_external_proto_depIdxs = nil
}
