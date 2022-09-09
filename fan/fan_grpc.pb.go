// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fan

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

// FanServiceClient is the client API for FanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FanServiceClient interface {
	// Fan
	NewFan(ctx context.Context, in *NewFanPost, opts ...grpc.CallOption) (*BoolResponse, error)
	Me(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*MeResponse, error)
	Fan(ctx context.Context, in *KeyPost, opts ...grpc.CallOption) (*FanResponse, error)
	SelfInfo(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*FanInfoResponse, error)
	FanInfo(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*FanInfoResponse, error)
	ChangeNickname(ctx context.Context, in *ChangeFanNicknamePost, opts ...grpc.CallOption) (*BoolResponse, error)
	ChangeAccount(ctx context.Context, in *ChangeFanAccountPost, opts ...grpc.CallOption) (*BoolResponse, error)
	ChangeSex(ctx context.Context, in *ChangeSexPost, opts ...grpc.CallOption) (*BoolResponse, error)
	UpdateFanArea(ctx context.Context, in *UpdateAreaPost, opts ...grpc.CallOption) (*BoolResponse, error)
	// Club
	MyClub(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*ClubResponse, error)
	ChooseClub(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error)
	// FanClub
	FanClub(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*FanClubResponse, error)
	MyFanClub(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*FanClubResponse, error)
	Pagination(ctx context.Context, in *FanClubPaginationPost, opts ...grpc.CallOption) (*FanClubPaginationResponse, error)
	Create(ctx context.Context, in *CreateFanClubPost, opts ...grpc.CallOption) (*FanClubResponse, error)
	Transfer(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*BoolResponse, error)
	UpdateClubArea(ctx context.Context, in *UpdateAreaPost, opts ...grpc.CallOption) (*BoolResponse, error)
	// Invite
	Invite(ctx context.Context, in *InviteFanPost, opts ...grpc.CallOption) (*BoolResponse, error)
	CancelInvite(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error)
	AcceptInvite(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error)
	RejectInvite(ctx context.Context, in *RejectInvitePost, opts ...grpc.CallOption) (*BoolResponse, error)
	// Join
	Join(ctx context.Context, in *JoinFanClubPost, opts ...grpc.CallOption) (*BoolResponse, error)
	CancelJoin(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error)
	AgreeJoin(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error)
	RejectJoin(ctx context.Context, in *JoinFanClubPost, opts ...grpc.CallOption) (*BoolResponse, error)
	// Leave
	Leave(ctx context.Context, in *LeaveFanClubPost, opts ...grpc.CallOption) (*BoolResponse, error)
	KickOut(ctx context.Context, in *KickOutPost, opts ...grpc.CallOption) (*BoolResponse, error)
}

type fanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFanServiceClient(cc grpc.ClientConnInterface) FanServiceClient {
	return &fanServiceClient{cc}
}

func (c *fanServiceClient) NewFan(ctx context.Context, in *NewFanPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/NewFan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) Me(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*MeResponse, error) {
	out := new(MeResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/Me", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) Fan(ctx context.Context, in *KeyPost, opts ...grpc.CallOption) (*FanResponse, error) {
	out := new(FanResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/Fan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) SelfInfo(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*FanInfoResponse, error) {
	out := new(FanInfoResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/SelfInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) FanInfo(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*FanInfoResponse, error) {
	out := new(FanInfoResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/FanInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) ChangeNickname(ctx context.Context, in *ChangeFanNicknamePost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/ChangeNickname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) ChangeAccount(ctx context.Context, in *ChangeFanAccountPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/ChangeAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) ChangeSex(ctx context.Context, in *ChangeSexPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/ChangeSex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) UpdateFanArea(ctx context.Context, in *UpdateAreaPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/UpdateFanArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) MyClub(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*ClubResponse, error) {
	out := new(ClubResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/MyClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) ChooseClub(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/ChooseClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) FanClub(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*FanClubResponse, error) {
	out := new(FanClubResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/FanClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) MyFanClub(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*FanClubResponse, error) {
	out := new(FanClubResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/MyFanClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) Pagination(ctx context.Context, in *FanClubPaginationPost, opts ...grpc.CallOption) (*FanClubPaginationResponse, error) {
	out := new(FanClubPaginationResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/Pagination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) Create(ctx context.Context, in *CreateFanClubPost, opts ...grpc.CallOption) (*FanClubResponse, error) {
	out := new(FanClubResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) Transfer(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) UpdateClubArea(ctx context.Context, in *UpdateAreaPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/UpdateClubArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) Invite(ctx context.Context, in *InviteFanPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/Invite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) CancelInvite(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/CancelInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) AcceptInvite(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/AcceptInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) RejectInvite(ctx context.Context, in *RejectInvitePost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/RejectInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) Join(ctx context.Context, in *JoinFanClubPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) CancelJoin(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/CancelJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) AgreeJoin(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/AgreeJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) RejectJoin(ctx context.Context, in *JoinFanClubPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/RejectJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) Leave(ctx context.Context, in *LeaveFanClubPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanServiceClient) KickOut(ctx context.Context, in *KickOutPost, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/fan.FanService/KickOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FanServiceServer is the server API for FanService service.
// All implementations must embed UnimplementedFanServiceServer
// for forward compatibility
type FanServiceServer interface {
	// Fan
	NewFan(context.Context, *NewFanPost) (*BoolResponse, error)
	Me(context.Context, *EmptyPost) (*MeResponse, error)
	Fan(context.Context, *KeyPost) (*FanResponse, error)
	SelfInfo(context.Context, *EmptyPost) (*FanInfoResponse, error)
	FanInfo(context.Context, *InfoPost) (*FanInfoResponse, error)
	ChangeNickname(context.Context, *ChangeFanNicknamePost) (*BoolResponse, error)
	ChangeAccount(context.Context, *ChangeFanAccountPost) (*BoolResponse, error)
	ChangeSex(context.Context, *ChangeSexPost) (*BoolResponse, error)
	UpdateFanArea(context.Context, *UpdateAreaPost) (*BoolResponse, error)
	// Club
	MyClub(context.Context, *EmptyPost) (*ClubResponse, error)
	ChooseClub(context.Context, *InfoPost) (*BoolResponse, error)
	// FanClub
	FanClub(context.Context, *InfoPost) (*FanClubResponse, error)
	MyFanClub(context.Context, *EmptyPost) (*FanClubResponse, error)
	Pagination(context.Context, *FanClubPaginationPost) (*FanClubPaginationResponse, error)
	Create(context.Context, *CreateFanClubPost) (*FanClubResponse, error)
	Transfer(context.Context, *EmptyPost) (*BoolResponse, error)
	UpdateClubArea(context.Context, *UpdateAreaPost) (*BoolResponse, error)
	// Invite
	Invite(context.Context, *InviteFanPost) (*BoolResponse, error)
	CancelInvite(context.Context, *InfoPost) (*BoolResponse, error)
	AcceptInvite(context.Context, *InfoPost) (*BoolResponse, error)
	RejectInvite(context.Context, *RejectInvitePost) (*BoolResponse, error)
	// Join
	Join(context.Context, *JoinFanClubPost) (*BoolResponse, error)
	CancelJoin(context.Context, *InfoPost) (*BoolResponse, error)
	AgreeJoin(context.Context, *InfoPost) (*BoolResponse, error)
	RejectJoin(context.Context, *JoinFanClubPost) (*BoolResponse, error)
	// Leave
	Leave(context.Context, *LeaveFanClubPost) (*BoolResponse, error)
	KickOut(context.Context, *KickOutPost) (*BoolResponse, error)
	mustEmbedUnimplementedFanServiceServer()
}

// UnimplementedFanServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFanServiceServer struct {
}

func (UnimplementedFanServiceServer) NewFan(context.Context, *NewFanPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewFan not implemented")
}
func (UnimplementedFanServiceServer) Me(context.Context, *EmptyPost) (*MeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Me not implemented")
}
func (UnimplementedFanServiceServer) Fan(context.Context, *KeyPost) (*FanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fan not implemented")
}
func (UnimplementedFanServiceServer) SelfInfo(context.Context, *EmptyPost) (*FanInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelfInfo not implemented")
}
func (UnimplementedFanServiceServer) FanInfo(context.Context, *InfoPost) (*FanInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FanInfo not implemented")
}
func (UnimplementedFanServiceServer) ChangeNickname(context.Context, *ChangeFanNicknamePost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNickname not implemented")
}
func (UnimplementedFanServiceServer) ChangeAccount(context.Context, *ChangeFanAccountPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAccount not implemented")
}
func (UnimplementedFanServiceServer) ChangeSex(context.Context, *ChangeSexPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSex not implemented")
}
func (UnimplementedFanServiceServer) UpdateFanArea(context.Context, *UpdateAreaPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFanArea not implemented")
}
func (UnimplementedFanServiceServer) MyClub(context.Context, *EmptyPost) (*ClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyClub not implemented")
}
func (UnimplementedFanServiceServer) ChooseClub(context.Context, *InfoPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChooseClub not implemented")
}
func (UnimplementedFanServiceServer) FanClub(context.Context, *InfoPost) (*FanClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FanClub not implemented")
}
func (UnimplementedFanServiceServer) MyFanClub(context.Context, *EmptyPost) (*FanClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyFanClub not implemented")
}
func (UnimplementedFanServiceServer) Pagination(context.Context, *FanClubPaginationPost) (*FanClubPaginationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pagination not implemented")
}
func (UnimplementedFanServiceServer) Create(context.Context, *CreateFanClubPost) (*FanClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFanServiceServer) Transfer(context.Context, *EmptyPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedFanServiceServer) UpdateClubArea(context.Context, *UpdateAreaPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClubArea not implemented")
}
func (UnimplementedFanServiceServer) Invite(context.Context, *InviteFanPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invite not implemented")
}
func (UnimplementedFanServiceServer) CancelInvite(context.Context, *InfoPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelInvite not implemented")
}
func (UnimplementedFanServiceServer) AcceptInvite(context.Context, *InfoPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptInvite not implemented")
}
func (UnimplementedFanServiceServer) RejectInvite(context.Context, *RejectInvitePost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectInvite not implemented")
}
func (UnimplementedFanServiceServer) Join(context.Context, *JoinFanClubPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedFanServiceServer) CancelJoin(context.Context, *InfoPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelJoin not implemented")
}
func (UnimplementedFanServiceServer) AgreeJoin(context.Context, *InfoPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgreeJoin not implemented")
}
func (UnimplementedFanServiceServer) RejectJoin(context.Context, *JoinFanClubPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectJoin not implemented")
}
func (UnimplementedFanServiceServer) Leave(context.Context, *LeaveFanClubPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (UnimplementedFanServiceServer) KickOut(context.Context, *KickOutPost) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickOut not implemented")
}
func (UnimplementedFanServiceServer) mustEmbedUnimplementedFanServiceServer() {}

// UnsafeFanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FanServiceServer will
// result in compilation errors.
type UnsafeFanServiceServer interface {
	mustEmbedUnimplementedFanServiceServer()
}

func RegisterFanServiceServer(s grpc.ServiceRegistrar, srv FanServiceServer) {
	s.RegisterService(&FanService_ServiceDesc, srv)
}

func _FanService_NewFan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewFanPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).NewFan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/NewFan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).NewFan(ctx, req.(*NewFanPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_Me_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).Me(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/Me",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).Me(ctx, req.(*EmptyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_Fan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).Fan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/Fan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).Fan(ctx, req.(*KeyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_SelfInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).SelfInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/SelfInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).SelfInfo(ctx, req.(*EmptyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_FanInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).FanInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/FanInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).FanInfo(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_ChangeNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeFanNicknamePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).ChangeNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/ChangeNickname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).ChangeNickname(ctx, req.(*ChangeFanNicknamePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_ChangeAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeFanAccountPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).ChangeAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/ChangeAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).ChangeAccount(ctx, req.(*ChangeFanAccountPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_ChangeSex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSexPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).ChangeSex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/ChangeSex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).ChangeSex(ctx, req.(*ChangeSexPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_UpdateFanArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAreaPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).UpdateFanArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/UpdateFanArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).UpdateFanArea(ctx, req.(*UpdateAreaPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_MyClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).MyClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/MyClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).MyClub(ctx, req.(*EmptyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_ChooseClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).ChooseClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/ChooseClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).ChooseClub(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_FanClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).FanClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/FanClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).FanClub(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_MyFanClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).MyFanClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/MyFanClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).MyFanClub(ctx, req.(*EmptyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_Pagination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FanClubPaginationPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).Pagination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/Pagination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).Pagination(ctx, req.(*FanClubPaginationPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFanClubPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).Create(ctx, req.(*CreateFanClubPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).Transfer(ctx, req.(*EmptyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_UpdateClubArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAreaPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).UpdateClubArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/UpdateClubArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).UpdateClubArea(ctx, req.(*UpdateAreaPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_Invite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteFanPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).Invite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/Invite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).Invite(ctx, req.(*InviteFanPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_CancelInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).CancelInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/CancelInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).CancelInvite(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_AcceptInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).AcceptInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/AcceptInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).AcceptInvite(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_RejectInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectInvitePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).RejectInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/RejectInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).RejectInvite(ctx, req.(*RejectInvitePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinFanClubPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).Join(ctx, req.(*JoinFanClubPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_CancelJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).CancelJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/CancelJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).CancelJoin(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_AgreeJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).AgreeJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/AgreeJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).AgreeJoin(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_RejectJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinFanClubPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).RejectJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/RejectJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).RejectJoin(ctx, req.(*JoinFanClubPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveFanClubPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).Leave(ctx, req.(*LeaveFanClubPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanService_KickOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickOutPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanServiceServer).KickOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fan.FanService/KickOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanServiceServer).KickOut(ctx, req.(*KickOutPost))
	}
	return interceptor(ctx, in, info, handler)
}

// FanService_ServiceDesc is the grpc.ServiceDesc for FanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fan.FanService",
	HandlerType: (*FanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewFan",
			Handler:    _FanService_NewFan_Handler,
		},
		{
			MethodName: "Me",
			Handler:    _FanService_Me_Handler,
		},
		{
			MethodName: "Fan",
			Handler:    _FanService_Fan_Handler,
		},
		{
			MethodName: "SelfInfo",
			Handler:    _FanService_SelfInfo_Handler,
		},
		{
			MethodName: "FanInfo",
			Handler:    _FanService_FanInfo_Handler,
		},
		{
			MethodName: "ChangeNickname",
			Handler:    _FanService_ChangeNickname_Handler,
		},
		{
			MethodName: "ChangeAccount",
			Handler:    _FanService_ChangeAccount_Handler,
		},
		{
			MethodName: "ChangeSex",
			Handler:    _FanService_ChangeSex_Handler,
		},
		{
			MethodName: "UpdateFanArea",
			Handler:    _FanService_UpdateFanArea_Handler,
		},
		{
			MethodName: "MyClub",
			Handler:    _FanService_MyClub_Handler,
		},
		{
			MethodName: "ChooseClub",
			Handler:    _FanService_ChooseClub_Handler,
		},
		{
			MethodName: "FanClub",
			Handler:    _FanService_FanClub_Handler,
		},
		{
			MethodName: "MyFanClub",
			Handler:    _FanService_MyFanClub_Handler,
		},
		{
			MethodName: "Pagination",
			Handler:    _FanService_Pagination_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FanService_Create_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _FanService_Transfer_Handler,
		},
		{
			MethodName: "UpdateClubArea",
			Handler:    _FanService_UpdateClubArea_Handler,
		},
		{
			MethodName: "Invite",
			Handler:    _FanService_Invite_Handler,
		},
		{
			MethodName: "CancelInvite",
			Handler:    _FanService_CancelInvite_Handler,
		},
		{
			MethodName: "AcceptInvite",
			Handler:    _FanService_AcceptInvite_Handler,
		},
		{
			MethodName: "RejectInvite",
			Handler:    _FanService_RejectInvite_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _FanService_Join_Handler,
		},
		{
			MethodName: "CancelJoin",
			Handler:    _FanService_CancelJoin_Handler,
		},
		{
			MethodName: "AgreeJoin",
			Handler:    _FanService_AgreeJoin_Handler,
		},
		{
			MethodName: "RejectJoin",
			Handler:    _FanService_RejectJoin_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _FanService_Leave_Handler,
		},
		{
			MethodName: "KickOut",
			Handler:    _FanService_KickOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fan/fan.proto",
}