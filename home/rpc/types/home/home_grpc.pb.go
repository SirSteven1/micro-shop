// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: home.proto

package home

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

// HomeClient is the client API for Home service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeClient interface {
	Banner(ctx context.Context, in *BannerReq, opts ...grpc.CallOption) (*CommonResply, error)
	Recommend(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*CommonResply, error)
	RankingList(ctx context.Context, in *RankingListReq, opts ...grpc.CallOption) (*CommonResply, error)
}

type homeClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeClient(cc grpc.ClientConnInterface) HomeClient {
	return &homeClient{cc}
}

func (c *homeClient) Banner(ctx context.Context, in *BannerReq, opts ...grpc.CallOption) (*CommonResply, error) {
	out := new(CommonResply)
	err := c.cc.Invoke(ctx, "/home.home/Banner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeClient) Recommend(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*CommonResply, error) {
	out := new(CommonResply)
	err := c.cc.Invoke(ctx, "/home.home/Recommend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeClient) RankingList(ctx context.Context, in *RankingListReq, opts ...grpc.CallOption) (*CommonResply, error) {
	out := new(CommonResply)
	err := c.cc.Invoke(ctx, "/home.home/RankingList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeServer is the server API for Home service.
// All implementations must embed UnimplementedHomeServer
// for forward compatibility
type HomeServer interface {
	Banner(context.Context, *BannerReq) (*CommonResply, error)
	Recommend(context.Context, *RecommendReq) (*CommonResply, error)
	RankingList(context.Context, *RankingListReq) (*CommonResply, error)
	mustEmbedUnimplementedHomeServer()
}

// UnimplementedHomeServer must be embedded to have forward compatible implementations.
type UnimplementedHomeServer struct {
}

func (UnimplementedHomeServer) Banner(context.Context, *BannerReq) (*CommonResply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Banner not implemented")
}
func (UnimplementedHomeServer) Recommend(context.Context, *RecommendReq) (*CommonResply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recommend not implemented")
}
func (UnimplementedHomeServer) RankingList(context.Context, *RankingListReq) (*CommonResply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RankingList not implemented")
}
func (UnimplementedHomeServer) mustEmbedUnimplementedHomeServer() {}

// UnsafeHomeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeServer will
// result in compilation errors.
type UnsafeHomeServer interface {
	mustEmbedUnimplementedHomeServer()
}

func RegisterHomeServer(s grpc.ServiceRegistrar, srv HomeServer) {
	s.RegisterService(&Home_ServiceDesc, srv)
}

func _Home_Banner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeServer).Banner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/home.home/Banner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeServer).Banner(ctx, req.(*BannerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Home_Recommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeServer).Recommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/home.home/Recommend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeServer).Recommend(ctx, req.(*RecommendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Home_RankingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankingListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeServer).RankingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/home.home/RankingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeServer).RankingList(ctx, req.(*RankingListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Home_ServiceDesc is the grpc.ServiceDesc for Home service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Home_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "home.home",
	HandlerType: (*HomeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Banner",
			Handler:    _Home_Banner_Handler,
		},
		{
			MethodName: "Recommend",
			Handler:    _Home_Recommend_Handler,
		},
		{
			MethodName: "RankingList",
			Handler:    _Home_RankingList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "home.proto",
}
