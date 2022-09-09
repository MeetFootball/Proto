// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package match

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

// MatchServiceClient is the client API for MatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchServiceClient interface {
	// 联赛球队
	LeaguesTree(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*LeaguesTreeResponse, error)
	// 俱乐部
	Club(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*ClubResponse, error)
	LeaguesClub(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error)
}

type matchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchServiceClient(cc grpc.ClientConnInterface) MatchServiceClient {
	return &matchServiceClient{cc}
}

func (c *matchServiceClient) LeaguesTree(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*LeaguesTreeResponse, error) {
	out := new(LeaguesTreeResponse)
	err := c.cc.Invoke(ctx, "/match.MatchService/LeaguesTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) Club(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*ClubResponse, error) {
	out := new(ClubResponse)
	err := c.cc.Invoke(ctx, "/match.MatchService/Club", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) LeaguesClub(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/match.MatchService/LeaguesClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchServiceServer is the server API for MatchService service.
// All implementations must embed UnimplementedMatchServiceServer
// for forward compatibility
type MatchServiceServer interface {
	// 联赛球队
	LeaguesTree(context.Context, *EmptyPost) (*LeaguesTreeResponse, error)
	// 俱乐部
	Club(context.Context, *InfoPost) (*ClubResponse, error)
	LeaguesClub(context.Context, *InfoPost) (*BoolResponse, error)
	mustEmbedUnimplementedMatchServiceServer()
}

// UnimplementedMatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchServiceServer struct {
}

func (UnimplementedMatchServiceServer) LeaguesTree(context.Context, *EmptyPost) (*LeaguesTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaguesTree not implemented")
}
func (UnimplementedMatchServiceServer) Club(context.Context, *InfoPost) (*ClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Club not implemented")
}
func (UnimplementedMatchServiceServer) LeaguesClub(context.Context, *InfoPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaguesClub not implemented")
}
func (UnimplementedMatchServiceServer) mustEmbedUnimplementedMatchServiceServer() {}

// UnsafeMatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchServiceServer will
// result in compilation errors.
type UnsafeMatchServiceServer interface {
	mustEmbedUnimplementedMatchServiceServer()
}

func RegisterMatchServiceServer(s grpc.ServiceRegistrar, srv MatchServiceServer) {
	s.RegisterService(&MatchService_ServiceDesc, srv)
}

func _MatchService_LeaguesTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).LeaguesTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/match.MatchService/LeaguesTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).LeaguesTree(ctx, req.(*EmptyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_Club_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).Club(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/match.MatchService/Club",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).Club(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_LeaguesClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).LeaguesClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/match.MatchService/LeaguesClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).LeaguesClub(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchService_ServiceDesc is the grpc.ServiceDesc for MatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "match.MatchService",
	HandlerType: (*MatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LeaguesTree",
			Handler:    _MatchService_LeaguesTree_Handler,
		},
		{
			MethodName: "Club",
			Handler:    _MatchService_Club_Handler,
		},
		{
			MethodName: "LeaguesClub",
			Handler:    _MatchService_LeaguesClub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "match/match.proto",
}