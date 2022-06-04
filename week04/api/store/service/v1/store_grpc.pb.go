// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: api/store/service/v1/store.proto

package v1

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

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreClient interface {
	CreateGoods(ctx context.Context, in *CreateGoodsReq, opts ...grpc.CallOption) (*CreateGoodsRsp, error)
	IncGoodsNum(ctx context.Context, in *IncGoodsNumReq, opts ...grpc.CallOption) (*IncGoodsNumRsp, error)
	ListGoods(ctx context.Context, in *ListGoodsReq, opts ...grpc.CallOption) (*ListGoodsRsp, error)
	DecGoodsNum(ctx context.Context, in *DecGoodsNumReq, opts ...grpc.CallOption) (*DecGoodsNumRsp, error)
}

type storeClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreClient(cc grpc.ClientConnInterface) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) CreateGoods(ctx context.Context, in *CreateGoodsReq, opts ...grpc.CallOption) (*CreateGoodsRsp, error) {
	out := new(CreateGoodsRsp)
	err := c.cc.Invoke(ctx, "/store.service.v1.Store/CreateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) IncGoodsNum(ctx context.Context, in *IncGoodsNumReq, opts ...grpc.CallOption) (*IncGoodsNumRsp, error) {
	out := new(IncGoodsNumRsp)
	err := c.cc.Invoke(ctx, "/store.service.v1.Store/IncGoodsNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) ListGoods(ctx context.Context, in *ListGoodsReq, opts ...grpc.CallOption) (*ListGoodsRsp, error) {
	out := new(ListGoodsRsp)
	err := c.cc.Invoke(ctx, "/store.service.v1.Store/ListGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) DecGoodsNum(ctx context.Context, in *DecGoodsNumReq, opts ...grpc.CallOption) (*DecGoodsNumRsp, error) {
	out := new(DecGoodsNumRsp)
	err := c.cc.Invoke(ctx, "/store.service.v1.Store/DecGoodsNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
// All implementations must embed UnimplementedStoreServer
// for forward compatibility
type StoreServer interface {
	CreateGoods(context.Context, *CreateGoodsReq) (*CreateGoodsRsp, error)
	IncGoodsNum(context.Context, *IncGoodsNumReq) (*IncGoodsNumRsp, error)
	ListGoods(context.Context, *ListGoodsReq) (*ListGoodsRsp, error)
	DecGoodsNum(context.Context, *DecGoodsNumReq) (*DecGoodsNumRsp, error)
	mustEmbedUnimplementedStoreServer()
}

// UnimplementedStoreServer must be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (UnimplementedStoreServer) CreateGoods(context.Context, *CreateGoodsReq) (*CreateGoodsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoods not implemented")
}
func (UnimplementedStoreServer) IncGoodsNum(context.Context, *IncGoodsNumReq) (*IncGoodsNumRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncGoodsNum not implemented")
}
func (UnimplementedStoreServer) ListGoods(context.Context, *ListGoodsReq) (*ListGoodsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoods not implemented")
}
func (UnimplementedStoreServer) DecGoodsNum(context.Context, *DecGoodsNumReq) (*DecGoodsNumRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecGoodsNum not implemented")
}
func (UnimplementedStoreServer) mustEmbedUnimplementedStoreServer() {}

// UnsafeStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServer will
// result in compilation errors.
type UnsafeStoreServer interface {
	mustEmbedUnimplementedStoreServer()
}

func RegisterStoreServer(s grpc.ServiceRegistrar, srv StoreServer) {
	s.RegisterService(&Store_ServiceDesc, srv)
}

func _Store_CreateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).CreateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.service.v1.Store/CreateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).CreateGoods(ctx, req.(*CreateGoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_IncGoodsNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncGoodsNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).IncGoodsNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.service.v1.Store/IncGoodsNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).IncGoodsNum(ctx, req.(*IncGoodsNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_ListGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).ListGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.service.v1.Store/ListGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).ListGoods(ctx, req.(*ListGoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_DecGoodsNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecGoodsNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).DecGoodsNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.service.v1.Store/DecGoodsNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).DecGoodsNum(ctx, req.(*DecGoodsNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Store_ServiceDesc is the grpc.ServiceDesc for Store service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Store_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "store.service.v1.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGoods",
			Handler:    _Store_CreateGoods_Handler,
		},
		{
			MethodName: "IncGoodsNum",
			Handler:    _Store_IncGoodsNum_Handler,
		},
		{
			MethodName: "ListGoods",
			Handler:    _Store_ListGoods_Handler,
		},
		{
			MethodName: "DecGoodsNum",
			Handler:    _Store_DecGoodsNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/store/service/v1/store.proto",
}
