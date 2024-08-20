// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/common/website_stated.proto

package common

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

const (
	WebsiteStatedService_Create_FullMethodName       = "/common.WebsiteStatedService/Create"
	WebsiteStatedService_Update_FullMethodName       = "/common.WebsiteStatedService/Update"
	WebsiteStatedService_Delete_FullMethodName       = "/common.WebsiteStatedService/Delete"
	WebsiteStatedService_FindById_FullMethodName     = "/common.WebsiteStatedService/FindById"
	WebsiteStatedService_FindPageList_FullMethodName = "/common.WebsiteStatedService/FindPageList"
)

// WebsiteStatedServiceClient is the client API for WebsiteStatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebsiteStatedServiceClient interface {
	// 新增
	Create(ctx context.Context, in *WebsiteStated, opts ...grpc.CallOption) (*WebsiteStatedResponse, error)
	// 修改
	Update(ctx context.Context, in *WebsiteStated, opts ...grpc.CallOption) (*WebsiteStatedResponse, error)
	// 删除
	Delete(ctx context.Context, in *WebsiteStatedIds, opts ...grpc.CallOption) (*WebsiteStatedResponse, error)
	// 查询 详情
	FindById(ctx context.Context, in *WebsiteStatedIds, opts ...grpc.CallOption) (*WebsiteStatedResponse, error)
	// 查询 分页
	FindPageList(ctx context.Context, in *WebsiteStatedPageAuthQuery, opts ...grpc.CallOption) (*WebsiteStatedResponse, error)
}

type websiteStatedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebsiteStatedServiceClient(cc grpc.ClientConnInterface) WebsiteStatedServiceClient {
	return &websiteStatedServiceClient{cc}
}

func (c *websiteStatedServiceClient) Create(ctx context.Context, in *WebsiteStated, opts ...grpc.CallOption) (*WebsiteStatedResponse, error) {
	out := new(WebsiteStatedResponse)
	err := c.cc.Invoke(ctx, WebsiteStatedService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteStatedServiceClient) Update(ctx context.Context, in *WebsiteStated, opts ...grpc.CallOption) (*WebsiteStatedResponse, error) {
	out := new(WebsiteStatedResponse)
	err := c.cc.Invoke(ctx, WebsiteStatedService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteStatedServiceClient) Delete(ctx context.Context, in *WebsiteStatedIds, opts ...grpc.CallOption) (*WebsiteStatedResponse, error) {
	out := new(WebsiteStatedResponse)
	err := c.cc.Invoke(ctx, WebsiteStatedService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteStatedServiceClient) FindById(ctx context.Context, in *WebsiteStatedIds, opts ...grpc.CallOption) (*WebsiteStatedResponse, error) {
	out := new(WebsiteStatedResponse)
	err := c.cc.Invoke(ctx, WebsiteStatedService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteStatedServiceClient) FindPageList(ctx context.Context, in *WebsiteStatedPageAuthQuery, opts ...grpc.CallOption) (*WebsiteStatedResponse, error) {
	out := new(WebsiteStatedResponse)
	err := c.cc.Invoke(ctx, WebsiteStatedService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebsiteStatedServiceServer is the server API for WebsiteStatedService service.
// All implementations must embed UnimplementedWebsiteStatedServiceServer
// for forward compatibility
type WebsiteStatedServiceServer interface {
	// 新增
	Create(context.Context, *WebsiteStated) (*WebsiteStatedResponse, error)
	// 修改
	Update(context.Context, *WebsiteStated) (*WebsiteStatedResponse, error)
	// 删除
	Delete(context.Context, *WebsiteStatedIds) (*WebsiteStatedResponse, error)
	// 查询 详情
	FindById(context.Context, *WebsiteStatedIds) (*WebsiteStatedResponse, error)
	// 查询 分页
	FindPageList(context.Context, *WebsiteStatedPageAuthQuery) (*WebsiteStatedResponse, error)
	mustEmbedUnimplementedWebsiteStatedServiceServer()
}

// UnimplementedWebsiteStatedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWebsiteStatedServiceServer struct {
}

func (UnimplementedWebsiteStatedServiceServer) Create(context.Context, *WebsiteStated) (*WebsiteStatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWebsiteStatedServiceServer) Update(context.Context, *WebsiteStated) (*WebsiteStatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWebsiteStatedServiceServer) Delete(context.Context, *WebsiteStatedIds) (*WebsiteStatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWebsiteStatedServiceServer) FindById(context.Context, *WebsiteStatedIds) (*WebsiteStatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedWebsiteStatedServiceServer) FindPageList(context.Context, *WebsiteStatedPageAuthQuery) (*WebsiteStatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedWebsiteStatedServiceServer) mustEmbedUnimplementedWebsiteStatedServiceServer() {}

// UnsafeWebsiteStatedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebsiteStatedServiceServer will
// result in compilation errors.
type UnsafeWebsiteStatedServiceServer interface {
	mustEmbedUnimplementedWebsiteStatedServiceServer()
}

func RegisterWebsiteStatedServiceServer(s grpc.ServiceRegistrar, srv WebsiteStatedServiceServer) {
	s.RegisterService(&WebsiteStatedService_ServiceDesc, srv)
}

func _WebsiteStatedService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebsiteStated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteStatedServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteStatedService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteStatedServiceServer).Create(ctx, req.(*WebsiteStated))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteStatedService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebsiteStated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteStatedServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteStatedService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteStatedServiceServer).Update(ctx, req.(*WebsiteStated))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteStatedService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebsiteStatedIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteStatedServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteStatedService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteStatedServiceServer).Delete(ctx, req.(*WebsiteStatedIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteStatedService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebsiteStatedIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteStatedServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteStatedService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteStatedServiceServer).FindById(ctx, req.(*WebsiteStatedIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebsiteStatedService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebsiteStatedPageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsiteStatedServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebsiteStatedService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsiteStatedServiceServer).FindPageList(ctx, req.(*WebsiteStatedPageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// WebsiteStatedService_ServiceDesc is the grpc.ServiceDesc for WebsiteStatedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebsiteStatedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.WebsiteStatedService",
	HandlerType: (*WebsiteStatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WebsiteStatedService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _WebsiteStatedService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WebsiteStatedService_Delete_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _WebsiteStatedService_FindById_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _WebsiteStatedService_FindPageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/common/website_stated.proto",
}
