// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceServiceClient is the client API for ServiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceServiceClient interface {
	// 对象存储账号管理
	Oss(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*OssResponse, error)
	OssList(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*OssListResponse, error)
	OssPagination(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*OssPaginationResponse, error)
	CreateOss(ctx context.Context, in *CreateOssPost, opts ...grpc.CallOption) (*OssResponse, error)
	UpdateOss(ctx context.Context, in *UpdateOssPost, opts ...grpc.CallOption) (*OssResponse, error)
	ChangeOssStatus(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*OssResponse, error)
	// 桶管理
	OssBucket(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*OssBucketResponse, error)
	OssBuckets(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*OssBucketsResponse, error)
	OssBucketPagination(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*OssBucketPaginationResponse, error)
	CreateOssBucket(ctx context.Context, in *CreateOssBucketPost, opts ...grpc.CallOption) (*OssBucketResponse, error)
	UpdateOssBucket(ctx context.Context, in *UpdateOssBucketPost, opts ...grpc.CallOption) (*OssBucketResponse, error)
	ChangeOssBucketStatus(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*OssBucketResponse, error)
}

type serviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceServiceClient(cc grpc.ClientConnInterface) ServiceServiceClient {
	return &serviceServiceClient{cc}
}

func (c *serviceServiceClient) Oss(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*OssResponse, error) {
	out := new(OssResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/Oss", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) OssList(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*OssListResponse, error) {
	out := new(OssListResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/OssList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) OssPagination(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*OssPaginationResponse, error) {
	out := new(OssPaginationResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/OssPagination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) CreateOss(ctx context.Context, in *CreateOssPost, opts ...grpc.CallOption) (*OssResponse, error) {
	out := new(OssResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/CreateOss", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) UpdateOss(ctx context.Context, in *UpdateOssPost, opts ...grpc.CallOption) (*OssResponse, error) {
	out := new(OssResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/UpdateOss", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) ChangeOssStatus(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*OssResponse, error) {
	out := new(OssResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/ChangeOssStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) OssBucket(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*OssBucketResponse, error) {
	out := new(OssBucketResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/OssBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) OssBuckets(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*OssBucketsResponse, error) {
	out := new(OssBucketsResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/OssBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) OssBucketPagination(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*OssBucketPaginationResponse, error) {
	out := new(OssBucketPaginationResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/OssBucketPagination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) CreateOssBucket(ctx context.Context, in *CreateOssBucketPost, opts ...grpc.CallOption) (*OssBucketResponse, error) {
	out := new(OssBucketResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/CreateOssBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) UpdateOssBucket(ctx context.Context, in *UpdateOssBucketPost, opts ...grpc.CallOption) (*OssBucketResponse, error) {
	out := new(OssBucketResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/UpdateOssBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) ChangeOssBucketStatus(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*OssBucketResponse, error) {
	out := new(OssBucketResponse)
	err := c.cc.Invoke(ctx, "/service.ServiceService/ChangeOssBucketStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServiceServer is the server API for ServiceService service.
// All implementations must embed UnimplementedServiceServiceServer
// for forward compatibility
type ServiceServiceServer interface {
	// 对象存储账号管理
	Oss(context.Context, *InfoPost) (*OssResponse, error)
	OssList(context.Context, *EmptyPost) (*OssListResponse, error)
	OssPagination(context.Context, *PaginationPost) (*OssPaginationResponse, error)
	CreateOss(context.Context, *CreateOssPost) (*OssResponse, error)
	UpdateOss(context.Context, *UpdateOssPost) (*OssResponse, error)
	ChangeOssStatus(context.Context, *ChangeStatusPost) (*OssResponse, error)
	// 桶管理
	OssBucket(context.Context, *InfoPost) (*OssBucketResponse, error)
	OssBuckets(context.Context, *EmptyPost) (*OssBucketsResponse, error)
	OssBucketPagination(context.Context, *PaginationPost) (*OssBucketPaginationResponse, error)
	CreateOssBucket(context.Context, *CreateOssBucketPost) (*OssBucketResponse, error)
	UpdateOssBucket(context.Context, *UpdateOssBucketPost) (*OssBucketResponse, error)
	ChangeOssBucketStatus(context.Context, *ChangeStatusPost) (*OssBucketResponse, error)
	mustEmbedUnimplementedServiceServiceServer()
}

// UnimplementedServiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServiceServer struct {
}

func (UnimplementedServiceServiceServer) Oss(context.Context, *InfoPost) (*OssResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Oss not implemented")
}
func (UnimplementedServiceServiceServer) OssList(context.Context, *EmptyPost) (*OssListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OssList not implemented")
}
func (UnimplementedServiceServiceServer) OssPagination(context.Context, *PaginationPost) (*OssPaginationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OssPagination not implemented")
}
func (UnimplementedServiceServiceServer) CreateOss(context.Context, *CreateOssPost) (*OssResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOss not implemented")
}
func (UnimplementedServiceServiceServer) UpdateOss(context.Context, *UpdateOssPost) (*OssResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOss not implemented")
}
func (UnimplementedServiceServiceServer) ChangeOssStatus(context.Context, *ChangeStatusPost) (*OssResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeOssStatus not implemented")
}
func (UnimplementedServiceServiceServer) OssBucket(context.Context, *InfoPost) (*OssBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OssBucket not implemented")
}
func (UnimplementedServiceServiceServer) OssBuckets(context.Context, *EmptyPost) (*OssBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OssBuckets not implemented")
}
func (UnimplementedServiceServiceServer) OssBucketPagination(context.Context, *PaginationPost) (*OssBucketPaginationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OssBucketPagination not implemented")
}
func (UnimplementedServiceServiceServer) CreateOssBucket(context.Context, *CreateOssBucketPost) (*OssBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOssBucket not implemented")
}
func (UnimplementedServiceServiceServer) UpdateOssBucket(context.Context, *UpdateOssBucketPost) (*OssBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOssBucket not implemented")
}
func (UnimplementedServiceServiceServer) ChangeOssBucketStatus(context.Context, *ChangeStatusPost) (*OssBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeOssBucketStatus not implemented")
}
func (UnimplementedServiceServiceServer) mustEmbedUnimplementedServiceServiceServer() {}

// UnsafeServiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServiceServer will
// result in compilation errors.
type UnsafeServiceServiceServer interface {
	mustEmbedUnimplementedServiceServiceServer()
}

func RegisterServiceServiceServer(s grpc.ServiceRegistrar, srv ServiceServiceServer) {
	s.RegisterService(&ServiceService_ServiceDesc, srv)
}

func _ServiceService_Oss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Oss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/Oss",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Oss(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_OssList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).OssList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/OssList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).OssList(ctx, req.(*EmptyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_OssPagination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).OssPagination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/OssPagination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).OssPagination(ctx, req.(*PaginationPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_CreateOss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOssPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).CreateOss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/CreateOss",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).CreateOss(ctx, req.(*CreateOssPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_UpdateOss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOssPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).UpdateOss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/UpdateOss",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).UpdateOss(ctx, req.(*UpdateOssPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_ChangeOssStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).ChangeOssStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/ChangeOssStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).ChangeOssStatus(ctx, req.(*ChangeStatusPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_OssBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).OssBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/OssBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).OssBucket(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_OssBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).OssBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/OssBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).OssBuckets(ctx, req.(*EmptyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_OssBucketPagination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).OssBucketPagination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/OssBucketPagination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).OssBucketPagination(ctx, req.(*PaginationPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_CreateOssBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOssBucketPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).CreateOssBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/CreateOssBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).CreateOssBucket(ctx, req.(*CreateOssBucketPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_UpdateOssBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOssBucketPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).UpdateOssBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/UpdateOssBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).UpdateOssBucket(ctx, req.(*UpdateOssBucketPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_ChangeOssBucketStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).ChangeOssBucketStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ServiceService/ChangeOssBucketStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).ChangeOssBucketStatus(ctx, req.(*ChangeStatusPost))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceService_ServiceDesc is the grpc.ServiceDesc for ServiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ServiceService",
	HandlerType: (*ServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Oss",
			Handler:    _ServiceService_Oss_Handler,
		},
		{
			MethodName: "OssList",
			Handler:    _ServiceService_OssList_Handler,
		},
		{
			MethodName: "OssPagination",
			Handler:    _ServiceService_OssPagination_Handler,
		},
		{
			MethodName: "CreateOss",
			Handler:    _ServiceService_CreateOss_Handler,
		},
		{
			MethodName: "UpdateOss",
			Handler:    _ServiceService_UpdateOss_Handler,
		},
		{
			MethodName: "ChangeOssStatus",
			Handler:    _ServiceService_ChangeOssStatus_Handler,
		},
		{
			MethodName: "OssBucket",
			Handler:    _ServiceService_OssBucket_Handler,
		},
		{
			MethodName: "OssBuckets",
			Handler:    _ServiceService_OssBuckets_Handler,
		},
		{
			MethodName: "OssBucketPagination",
			Handler:    _ServiceService_OssBucketPagination_Handler,
		},
		{
			MethodName: "CreateOssBucket",
			Handler:    _ServiceService_CreateOssBucket_Handler,
		},
		{
			MethodName: "UpdateOssBucket",
			Handler:    _ServiceService_UpdateOssBucket_Handler,
		},
		{
			MethodName: "ChangeOssBucketStatus",
			Handler:    _ServiceService_ChangeOssBucketStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}
