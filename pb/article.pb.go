// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: article/article.proto

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

var File_article_article_proto protoreflect.FileDescriptor

var file_article_article_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x1a, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfd, 0x02, 0x0a, 0x0e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x54,
	0x61, 0x67, 0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x54, 0x61,
	0x67, 0x73, 0x12, 0x11, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x0d, 0x54, 0x61, 0x67, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x54, 0x61, 0x67, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x16, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x67, 0x12, 0x16, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x61, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x6f, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_article_article_proto_goTypes = []interface{}{
	(*InfoPost)(nil),              // 0: global.InfoPost
	(*EmptyPost)(nil),             // 1: global.EmptyPost
	(*PaginationPost)(nil),        // 2: global.PaginationPost
	(*CreateTagPost)(nil),         // 3: article.CreateTagPost
	(*UpdateTagPost)(nil),         // 4: article.UpdateTagPost
	(*ChangeStatusPost)(nil),      // 5: global.ChangeStatusPost
	(*TagResponse)(nil),           // 6: article.TagResponse
	(*TagsResponse)(nil),          // 7: article.TagsResponse
	(*TagPaginationResponse)(nil), // 8: article.TagPaginationResponse
}
var file_article_article_proto_depIdxs = []int32{
	0, // 0: article.ArticleService.Tag:input_type -> global.InfoPost
	1, // 1: article.ArticleService.Tags:input_type -> global.EmptyPost
	2, // 2: article.ArticleService.TagPagination:input_type -> global.PaginationPost
	3, // 3: article.ArticleService.CreateTag:input_type -> article.CreateTagPost
	4, // 4: article.ArticleService.UpdateTag:input_type -> article.UpdateTagPost
	5, // 5: article.ArticleService.ChangeTagStatus:input_type -> global.ChangeStatusPost
	6, // 6: article.ArticleService.Tag:output_type -> article.TagResponse
	7, // 7: article.ArticleService.Tags:output_type -> article.TagsResponse
	8, // 8: article.ArticleService.TagPagination:output_type -> article.TagPaginationResponse
	6, // 9: article.ArticleService.CreateTag:output_type -> article.TagResponse
	6, // 10: article.ArticleService.UpdateTag:output_type -> article.TagResponse
	6, // 11: article.ArticleService.ChangeTagStatus:output_type -> article.TagResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_article_article_proto_init() }
func file_article_article_proto_init() {
	if File_article_article_proto != nil {
		return
	}
	file_global_global_proto_init()
	file_article_tag_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_article_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_article_article_proto_goTypes,
		DependencyIndexes: file_article_article_proto_depIdxs,
	}.Build()
	File_article_article_proto = out.File
	file_article_article_proto_rawDesc = nil
	file_article_article_proto_goTypes = nil
	file_article_article_proto_depIdxs = nil
}