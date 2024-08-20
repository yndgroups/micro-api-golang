// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/order/freight_templates_undelivered.proto

package order

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
	FreightTemplatesUndeliveredService_Create_FullMethodName       = "/order.FreightTemplatesUndeliveredService/Create"
	FreightTemplatesUndeliveredService_Update_FullMethodName       = "/order.FreightTemplatesUndeliveredService/Update"
	FreightTemplatesUndeliveredService_Delete_FullMethodName       = "/order.FreightTemplatesUndeliveredService/Delete"
	FreightTemplatesUndeliveredService_FindPageList_FullMethodName = "/order.FreightTemplatesUndeliveredService/FindPageList"
	FreightTemplatesUndeliveredService_FindById_FullMethodName     = "/order.FreightTemplatesUndeliveredService/FindById"
)

// FreightTemplatesUndeliveredServiceClient is the client API for FreightTemplatesUndeliveredService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FreightTemplatesUndeliveredServiceClient interface {
	// Create 不送达地区新增信息
	Create(ctx context.Context, in *FreightTemplatesUndelivered, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error)
	// Update 更新不送达地区信息
	Update(ctx context.Context, in *FreightTemplatesUndelivered, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error)
	// Delete 删除不送达地区信息
	Delete(ctx context.Context, in *FreightTemplatesUndeliveredIds, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error)
	// FindPageList 查询不送达地区列表
	FindPageList(ctx context.Context, in *FreightTemplatesUndeliveredPageAuthQuery, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error)
	// FindById 查询不送达地区详情
	FindById(ctx context.Context, in *FreightTemplatesUndeliveredIds, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error)
}

type freightTemplatesUndeliveredServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreightTemplatesUndeliveredServiceClient(cc grpc.ClientConnInterface) FreightTemplatesUndeliveredServiceClient {
	return &freightTemplatesUndeliveredServiceClient{cc}
}

func (c *freightTemplatesUndeliveredServiceClient) Create(ctx context.Context, in *FreightTemplatesUndelivered, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error) {
	out := new(FreightTemplatesUndeliveredResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesUndeliveredService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesUndeliveredServiceClient) Update(ctx context.Context, in *FreightTemplatesUndelivered, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error) {
	out := new(FreightTemplatesUndeliveredResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesUndeliveredService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesUndeliveredServiceClient) Delete(ctx context.Context, in *FreightTemplatesUndeliveredIds, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error) {
	out := new(FreightTemplatesUndeliveredResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesUndeliveredService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesUndeliveredServiceClient) FindPageList(ctx context.Context, in *FreightTemplatesUndeliveredPageAuthQuery, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error) {
	out := new(FreightTemplatesUndeliveredResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesUndeliveredService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesUndeliveredServiceClient) FindById(ctx context.Context, in *FreightTemplatesUndeliveredIds, opts ...grpc.CallOption) (*FreightTemplatesUndeliveredResponse, error) {
	out := new(FreightTemplatesUndeliveredResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesUndeliveredService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreightTemplatesUndeliveredServiceServer is the server API for FreightTemplatesUndeliveredService service.
// All implementations must embed UnimplementedFreightTemplatesUndeliveredServiceServer
// for forward compatibility
type FreightTemplatesUndeliveredServiceServer interface {
	// Create 不送达地区新增信息
	Create(context.Context, *FreightTemplatesUndelivered) (*FreightTemplatesUndeliveredResponse, error)
	// Update 更新不送达地区信息
	Update(context.Context, *FreightTemplatesUndelivered) (*FreightTemplatesUndeliveredResponse, error)
	// Delete 删除不送达地区信息
	Delete(context.Context, *FreightTemplatesUndeliveredIds) (*FreightTemplatesUndeliveredResponse, error)
	// FindPageList 查询不送达地区列表
	FindPageList(context.Context, *FreightTemplatesUndeliveredPageAuthQuery) (*FreightTemplatesUndeliveredResponse, error)
	// FindById 查询不送达地区详情
	FindById(context.Context, *FreightTemplatesUndeliveredIds) (*FreightTemplatesUndeliveredResponse, error)
	mustEmbedUnimplementedFreightTemplatesUndeliveredServiceServer()
}

// UnimplementedFreightTemplatesUndeliveredServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFreightTemplatesUndeliveredServiceServer struct {
}

func (UnimplementedFreightTemplatesUndeliveredServiceServer) Create(context.Context, *FreightTemplatesUndelivered) (*FreightTemplatesUndeliveredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFreightTemplatesUndeliveredServiceServer) Update(context.Context, *FreightTemplatesUndelivered) (*FreightTemplatesUndeliveredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFreightTemplatesUndeliveredServiceServer) Delete(context.Context, *FreightTemplatesUndeliveredIds) (*FreightTemplatesUndeliveredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFreightTemplatesUndeliveredServiceServer) FindPageList(context.Context, *FreightTemplatesUndeliveredPageAuthQuery) (*FreightTemplatesUndeliveredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedFreightTemplatesUndeliveredServiceServer) FindById(context.Context, *FreightTemplatesUndeliveredIds) (*FreightTemplatesUndeliveredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedFreightTemplatesUndeliveredServiceServer) mustEmbedUnimplementedFreightTemplatesUndeliveredServiceServer() {
}

// UnsafeFreightTemplatesUndeliveredServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FreightTemplatesUndeliveredServiceServer will
// result in compilation errors.
type UnsafeFreightTemplatesUndeliveredServiceServer interface {
	mustEmbedUnimplementedFreightTemplatesUndeliveredServiceServer()
}

func RegisterFreightTemplatesUndeliveredServiceServer(s grpc.ServiceRegistrar, srv FreightTemplatesUndeliveredServiceServer) {
	s.RegisterService(&FreightTemplatesUndeliveredService_ServiceDesc, srv)
}

func _FreightTemplatesUndeliveredService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesUndelivered)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesUndeliveredServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesUndeliveredService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesUndeliveredServiceServer).Create(ctx, req.(*FreightTemplatesUndelivered))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesUndeliveredService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesUndelivered)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesUndeliveredServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesUndeliveredService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesUndeliveredServiceServer).Update(ctx, req.(*FreightTemplatesUndelivered))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesUndeliveredService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesUndeliveredIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesUndeliveredServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesUndeliveredService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesUndeliveredServiceServer).Delete(ctx, req.(*FreightTemplatesUndeliveredIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesUndeliveredService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesUndeliveredPageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesUndeliveredServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesUndeliveredService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesUndeliveredServiceServer).FindPageList(ctx, req.(*FreightTemplatesUndeliveredPageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesUndeliveredService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesUndeliveredIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesUndeliveredServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesUndeliveredService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesUndeliveredServiceServer).FindById(ctx, req.(*FreightTemplatesUndeliveredIds))
	}
	return interceptor(ctx, in, info, handler)
}

// FreightTemplatesUndeliveredService_ServiceDesc is the grpc.ServiceDesc for FreightTemplatesUndeliveredService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FreightTemplatesUndeliveredService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.FreightTemplatesUndeliveredService",
	HandlerType: (*FreightTemplatesUndeliveredServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FreightTemplatesUndeliveredService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FreightTemplatesUndeliveredService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FreightTemplatesUndeliveredService_Delete_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _FreightTemplatesUndeliveredService_FindPageList_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _FreightTemplatesUndeliveredService_FindById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/order/freight_templates_undelivered.proto",
}