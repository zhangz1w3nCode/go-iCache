// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.3
// source: cache-monitor.proto

package cache_monitor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CacheMonitorService_GetCacheUserAddressList_FullMethodName = "/user.CacheMonitorService/GetCacheUserAddressList"
	CacheMonitorService_GetCacheUserAppNameList_FullMethodName = "/user.CacheMonitorService/GetCacheUserAppNameList"
	CacheMonitorService_GetCacheNameList_FullMethodName        = "/user.CacheMonitorService/GetCacheNameList"
)

// CacheMonitorServiceClient is the client API for CacheMonitorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheMonitorServiceClient interface {
	// 获取使用了iCache的机器的ip:哪个appName对应的机器用了iCache
	GetCacheUserAddressList(ctx context.Context, in *GetCacheUserAddressListRequest, opts ...grpc.CallOption) (*GetCacheUserAddressListResponse, error)
	// 获取使用了iCache的服务名称列表：biz-app、query-app
	GetCacheUserAppNameList(ctx context.Context, in *GetCacheUserAppNameListRequest, opts ...grpc.CallOption) (*GetCacheUserAppNameListResponse, error)
	// 该AppName下有哪些缓存：如productCache、userCache
	GetCacheNameList(ctx context.Context, in *GetCacheNameListRequest, opts ...grpc.CallOption) (*GetCacheNameListResponse, error)
}

type cacheMonitorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheMonitorServiceClient(cc grpc.ClientConnInterface) CacheMonitorServiceClient {
	return &cacheMonitorServiceClient{cc}
}

func (c *cacheMonitorServiceClient) GetCacheUserAddressList(ctx context.Context, in *GetCacheUserAddressListRequest, opts ...grpc.CallOption) (*GetCacheUserAddressListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCacheUserAddressListResponse)
	err := c.cc.Invoke(ctx, CacheMonitorService_GetCacheUserAddressList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheMonitorServiceClient) GetCacheUserAppNameList(ctx context.Context, in *GetCacheUserAppNameListRequest, opts ...grpc.CallOption) (*GetCacheUserAppNameListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCacheUserAppNameListResponse)
	err := c.cc.Invoke(ctx, CacheMonitorService_GetCacheUserAppNameList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheMonitorServiceClient) GetCacheNameList(ctx context.Context, in *GetCacheNameListRequest, opts ...grpc.CallOption) (*GetCacheNameListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCacheNameListResponse)
	err := c.cc.Invoke(ctx, CacheMonitorService_GetCacheNameList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheMonitorServiceServer is the server API for CacheMonitorService service.
// All implementations must embed UnimplementedCacheMonitorServiceServer
// for forward compatibility.
type CacheMonitorServiceServer interface {
	// 获取使用了iCache的机器的ip:哪个appName对应的机器用了iCache
	GetCacheUserAddressList(context.Context, *GetCacheUserAddressListRequest) (*GetCacheUserAddressListResponse, error)
	// 获取使用了iCache的服务名称列表：biz-app、query-app
	GetCacheUserAppNameList(context.Context, *GetCacheUserAppNameListRequest) (*GetCacheUserAppNameListResponse, error)
	// 该AppName下有哪些缓存：如productCache、userCache
	GetCacheNameList(context.Context, *GetCacheNameListRequest) (*GetCacheNameListResponse, error)
	mustEmbedUnimplementedCacheMonitorServiceServer()
}

// UnimplementedCacheMonitorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCacheMonitorServiceServer struct{}

func (UnimplementedCacheMonitorServiceServer) GetCacheUserAddressList(context.Context, *GetCacheUserAddressListRequest) (*GetCacheUserAddressListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheUserAddressList not implemented")
}
func (UnimplementedCacheMonitorServiceServer) GetCacheUserAppNameList(context.Context, *GetCacheUserAppNameListRequest) (*GetCacheUserAppNameListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheUserAppNameList not implemented")
}
func (UnimplementedCacheMonitorServiceServer) GetCacheNameList(context.Context, *GetCacheNameListRequest) (*GetCacheNameListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheNameList not implemented")
}
func (UnimplementedCacheMonitorServiceServer) mustEmbedUnimplementedCacheMonitorServiceServer() {}
func (UnimplementedCacheMonitorServiceServer) testEmbeddedByValue()                             {}

// UnsafeCacheMonitorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheMonitorServiceServer will
// result in compilation errors.
type UnsafeCacheMonitorServiceServer interface {
	mustEmbedUnimplementedCacheMonitorServiceServer()
}

func RegisterCacheMonitorServiceServer(s grpc.ServiceRegistrar, srv CacheMonitorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCacheMonitorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CacheMonitorService_ServiceDesc, srv)
}

func _CacheMonitorService_GetCacheUserAddressList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheUserAddressListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheMonitorServiceServer).GetCacheUserAddressList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheMonitorService_GetCacheUserAddressList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheMonitorServiceServer).GetCacheUserAddressList(ctx, req.(*GetCacheUserAddressListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheMonitorService_GetCacheUserAppNameList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheUserAppNameListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheMonitorServiceServer).GetCacheUserAppNameList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheMonitorService_GetCacheUserAppNameList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheMonitorServiceServer).GetCacheUserAppNameList(ctx, req.(*GetCacheUserAppNameListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheMonitorService_GetCacheNameList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheNameListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheMonitorServiceServer).GetCacheNameList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheMonitorService_GetCacheNameList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheMonitorServiceServer).GetCacheNameList(ctx, req.(*GetCacheNameListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheMonitorService_ServiceDesc is the grpc.ServiceDesc for CacheMonitorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheMonitorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.CacheMonitorService",
	HandlerType: (*CacheMonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCacheUserAddressList",
			Handler:    _CacheMonitorService_GetCacheUserAddressList_Handler,
		},
		{
			MethodName: "GetCacheUserAppNameList",
			Handler:    _CacheMonitorService_GetCacheUserAppNameList_Handler,
		},
		{
			MethodName: "GetCacheNameList",
			Handler:    _CacheMonitorService_GetCacheNameList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache-monitor.proto",
}
