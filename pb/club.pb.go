// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: club/club.proto

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

var File_club_club_proto protoreflect.FileDescriptor

var file_club_club_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x6c, 0x75, 0x62, 0x1a, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x6c,
	0x75, 0x62, 0x2f, 0x66, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6c,
	0x75, 0x62, 0x2f, 0x66, 0x61, 0x6e, 0x5f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x94, 0x0a, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x46, 0x61, 0x6e, 0x12, 0x10, 0x2e, 0x63, 0x6c,
	0x75, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x46, 0x61, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x02, 0x4d, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2b, 0x0a, 0x03, 0x46, 0x61, 0x6e, 0x12, 0x0f, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x6c, 0x75,
	0x62, 0x2e, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46,
	0x61, 0x6e, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x61, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x12, 0x16, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61,
	0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04,
	0x43, 0x6c, 0x75, 0x62, 0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x46, 0x61,
	0x6e, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x09, 0x4d, 0x79, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x11, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x46, 0x61,
	0x6e, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x46, 0x61, 0x6e, 0x43, 0x6c,
	0x75, 0x62, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61,
	0x6e, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6c, 0x75, 0x62,
	0x2e, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x11,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x41, 0x72, 0x65, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x15,
	0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75,
	0x62, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x41, 0x67, 0x72, 0x65, 0x65, 0x4a, 0x6f,
	0x69, 0x6e, 0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x63, 0x6c, 0x75,
	0x62, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x6f, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x05, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x46,
	0x61, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x12, 0x11, 0x2e,
	0x63, 0x6c, 0x75, 0x62, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_club_club_proto_goTypes = []interface{}{
	(*NewFanPost)(nil),                // 0: club.NewFanPost
	(*EmptyPost)(nil),                 // 1: global.EmptyPost
	(*KeyPost)(nil),                   // 2: global.KeyPost
	(*ChangeFanNicknamePost)(nil),     // 3: club.ChangeFanNicknamePost
	(*ChangeFanAccountPost)(nil),      // 4: club.ChangeFanAccountPost
	(*UpdateAreaPost)(nil),            // 5: global.UpdateAreaPost
	(*InfoPost)(nil),                  // 6: global.InfoPost
	(*FanClubPaginationPost)(nil),     // 7: club.FanClubPaginationPost
	(*CreateFanClubPost)(nil),         // 8: club.CreateFanClubPost
	(*InviteFanPost)(nil),             // 9: club.InviteFanPost
	(*RejectInvitePost)(nil),          // 10: club.RejectInvitePost
	(*JoinFanClubPost)(nil),           // 11: club.JoinFanClubPost
	(*LeaveFanClubPost)(nil),          // 12: club.LeaveFanClubPost
	(*KickOutPost)(nil),               // 13: club.KickOutPost
	(*BoolResponse)(nil),              // 14: global.BoolResponse
	(*MeResponse)(nil),                // 15: club.MeResponse
	(*FanResponse)(nil),               // 16: club.FanResponse
	(*FanClubResponse)(nil),           // 17: club.FanClubResponse
	(*FanClubPaginationResponse)(nil), // 18: club.FanClubPaginationResponse
}
var file_club_club_proto_depIdxs = []int32{
	0,  // 0: club.ClubService.NewFan:input_type -> club.NewFanPost
	1,  // 1: club.ClubService.Me:input_type -> global.EmptyPost
	2,  // 2: club.ClubService.Fan:input_type -> global.KeyPost
	3,  // 3: club.ClubService.ChangeNickname:input_type -> club.ChangeFanNicknamePost
	4,  // 4: club.ClubService.ChangeAccount:input_type -> club.ChangeFanAccountPost
	5,  // 5: club.ClubService.UpdateFanArea:input_type -> global.UpdateAreaPost
	6,  // 6: club.ClubService.Club:input_type -> global.InfoPost
	1,  // 7: club.ClubService.MyFanClub:input_type -> global.EmptyPost
	7,  // 8: club.ClubService.Pagination:input_type -> club.FanClubPaginationPost
	8,  // 9: club.ClubService.Create:input_type -> club.CreateFanClubPost
	1,  // 10: club.ClubService.Transfer:input_type -> global.EmptyPost
	5,  // 11: club.ClubService.UpdateClubArea:input_type -> global.UpdateAreaPost
	9,  // 12: club.ClubService.Invite:input_type -> club.InviteFanPost
	6,  // 13: club.ClubService.CancelInvite:input_type -> global.InfoPost
	6,  // 14: club.ClubService.AcceptInvite:input_type -> global.InfoPost
	10, // 15: club.ClubService.RejectInvite:input_type -> club.RejectInvitePost
	11, // 16: club.ClubService.Join:input_type -> club.JoinFanClubPost
	6,  // 17: club.ClubService.CancelJoin:input_type -> global.InfoPost
	6,  // 18: club.ClubService.AgreeJoin:input_type -> global.InfoPost
	11, // 19: club.ClubService.RejectJoin:input_type -> club.JoinFanClubPost
	12, // 20: club.ClubService.Leave:input_type -> club.LeaveFanClubPost
	13, // 21: club.ClubService.KickOut:input_type -> club.KickOutPost
	14, // 22: club.ClubService.NewFan:output_type -> global.BoolResponse
	15, // 23: club.ClubService.Me:output_type -> club.MeResponse
	16, // 24: club.ClubService.Fan:output_type -> club.FanResponse
	14, // 25: club.ClubService.ChangeNickname:output_type -> global.BoolResponse
	14, // 26: club.ClubService.ChangeAccount:output_type -> global.BoolResponse
	14, // 27: club.ClubService.UpdateFanArea:output_type -> global.BoolResponse
	17, // 28: club.ClubService.Club:output_type -> club.FanClubResponse
	17, // 29: club.ClubService.MyFanClub:output_type -> club.FanClubResponse
	18, // 30: club.ClubService.Pagination:output_type -> club.FanClubPaginationResponse
	17, // 31: club.ClubService.Create:output_type -> club.FanClubResponse
	14, // 32: club.ClubService.Transfer:output_type -> global.BoolResponse
	14, // 33: club.ClubService.UpdateClubArea:output_type -> global.BoolResponse
	14, // 34: club.ClubService.Invite:output_type -> global.BoolResponse
	14, // 35: club.ClubService.CancelInvite:output_type -> global.BoolResponse
	14, // 36: club.ClubService.AcceptInvite:output_type -> global.BoolResponse
	14, // 37: club.ClubService.RejectInvite:output_type -> global.BoolResponse
	14, // 38: club.ClubService.Join:output_type -> global.BoolResponse
	14, // 39: club.ClubService.CancelJoin:output_type -> global.BoolResponse
	14, // 40: club.ClubService.AgreeJoin:output_type -> global.BoolResponse
	14, // 41: club.ClubService.RejectJoin:output_type -> global.BoolResponse
	14, // 42: club.ClubService.Leave:output_type -> global.BoolResponse
	14, // 43: club.ClubService.KickOut:output_type -> global.BoolResponse
	22, // [22:44] is the sub-list for method output_type
	0,  // [0:22] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_club_club_proto_init() }
func file_club_club_proto_init() {
	if File_club_club_proto != nil {
		return
	}
	file_global_global_proto_init()
	file_club_fan_proto_init()
	file_club_fan_club_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_club_club_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_club_club_proto_goTypes,
		DependencyIndexes: file_club_club_proto_depIdxs,
	}.Build()
	File_club_club_proto = out.File
	file_club_club_proto_rawDesc = nil
	file_club_club_proto_goTypes = nil
	file_club_club_proto_depIdxs = nil
}
