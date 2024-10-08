// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/order/freight_templates_free.proto

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
	FreightTemplatesFreeService_Create_FullMethodName       = "/order.FreightTemplatesFreeService/Create"
	FreightTemplatesFreeService_Update_FullMethodName       = "/order.FreightTemplatesFreeService/Update"
	FreightTemplatesFreeService_Delete_FullMethodName       = "/order.FreightTemplatesFreeService/Delete"
	FreightTemplatesFreeService_FindPageList_FullMethodName = "/order.FreightTemplatesFreeService/FindPageList"
	FreightTemplatesFreeService_FindById_FullMethodName     = "/order.FreightTemplatesFreeService/FindById"
)

// FreightTemplatesFreeServiceClient is the client API for FreightTemplatesFreeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FreightTemplatesFreeServiceClient interface {
	// Create 地区新增信息
	Create(ctx context.Context, in *FreightTemplatesFree, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error)
	// Update 更新地区信息
	Update(ctx context.Context, in *FreightTemplatesFree, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error)
	// Delete 删除地区信息
	Delete(ctx context.Context, in *FreightTemplatesFreeIds, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error)
	// FindPageList 查询地区列表
	FindPageList(ctx context.Context, in *FreightTemplatesFreePageAuthQuery, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error)
	// FindById 查询地区详情
	FindById(ctx context.Context, in *FreightTemplatesFreeIds, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error)
}

type freightTemplatesFreeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreightTemplatesFreeServiceClient(cc grpc.ClientConnInterface) FreightTemplatesFreeServiceClient {
	return &freightTemplatesFreeServiceClient{cc}
}

func (c *freightTemplatesFreeServiceClient) Create(ctx context.Context, in *FreightTemplatesFree, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error) {
	out := new(FreightTemplatesFreeResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesFreeService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesFreeServiceClient) Update(ctx context.Context, in *FreightTemplatesFree, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error) {
	out := new(FreightTemplatesFreeResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesFreeService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesFreeServiceClient) Delete(ctx context.Context, in *FreightTemplatesFreeIds, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error) {
	out := new(FreightTemplatesFreeResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesFreeService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesFreeServiceClient) FindPageList(ctx context.Context, in *FreightTemplatesFreePageAuthQuery, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error) {
	out := new(FreightTemplatesFreeResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesFreeService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesFreeServiceClient) FindById(ctx context.Context, in *FreightTemplatesFreeIds, opts ...grpc.CallOption) (*FreightTemplatesFreeResponse, error) {
	out := new(FreightTemplatesFreeResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesFreeService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreightTemplatesFreeServiceServer is the server API for FreightTemplatesFreeService service.
// All implementations must embed UnimplementedFreightTemplatesFreeServiceServer
// for forward compatibility
type FreightTemplatesFreeServiceServer interface {
	// Create 地区新增信息
	Create(context.Context, *FreightTemplatesFree) (*FreightTemplatesFreeResponse, error)
	// Update 更新地区信息
	Update(context.Context, *FreightTemplatesFree) (*FreightTemplatesFreeResponse, error)
	// Delete 删除地区信息
	Delete(context.Context, *FreightTemplatesFreeIds) (*FreightTemplatesFreeResponse, error)
	// FindPageList 查询地区列表
	FindPageList(context.Context, *FreightTemplatesFreePageAuthQuery) (*FreightTemplatesFreeResponse, error)
	// FindById 查询地区详情
	FindById(context.Context, *FreightTemplatesFreeIds) (*FreightTemplatesFreeResponse, error)
	mustEmbedUnimplementedFreightTemplatesFreeServiceServer()
}

// UnimplementedFreightTemplatesFreeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFreightTemplatesFreeServiceServer struct {
}

func (UnimplementedFreightTemplatesFreeServiceServer) Create(context.Context, *FreightTemplatesFree) (*FreightTemplatesFreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFreightTemplatesFreeServiceServer) Update(context.Context, *FreightTemplatesFree) (*FreightTemplatesFreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFreightTemplatesFreeServiceServer) Delete(context.Context, *FreightTemplatesFreeIds) (*FreightTemplatesFreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFreightTemplatesFreeServiceServer) FindPageList(context.Context, *FreightTemplatesFreePageAuthQuery) (*FreightTemplatesFreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedFreightTemplatesFreeServiceServer) FindById(context.Context, *FreightTemplatesFreeIds) (*FreightTemplatesFreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedFreightTemplatesFreeServiceServer) mustEmbedUnimplementedFreightTemplatesFreeServiceServer() {
}

// UnsafeFreightTemplatesFreeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FreightTemplatesFreeServiceServer will
// result in compilation errors.
type UnsafeFreightTemplatesFreeServiceServer interface {
	mustEmbedUnimplementedFreightTemplatesFreeServiceServer()
}

func RegisterFreightTemplatesFreeServiceServer(s grpc.ServiceRegistrar, srv FreightTemplatesFreeServiceServer) {
	s.RegisterService(&FreightTemplatesFreeService_ServiceDesc, srv)
}

func _FreightTemplatesFreeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesFree)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesFreeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesFreeService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesFreeServiceServer).Create(ctx, req.(*FreightTemplatesFree))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesFreeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesFree)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesFreeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesFreeService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesFreeServiceServer).Update(ctx, req.(*FreightTemplatesFree))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesFreeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesFreeIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesFreeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesFreeService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesFreeServiceServer).Delete(ctx, req.(*FreightTemplatesFreeIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesFreeService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesFreePageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesFreeServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesFreeService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesFreeServiceServer).FindPageList(ctx, req.(*FreightTemplatesFreePageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesFreeService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesFreeIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesFreeServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesFreeService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesFreeServiceServer).FindById(ctx, req.(*FreightTemplatesFreeIds))
	}
	return interceptor(ctx, in, info, handler)
}

// FreightTemplatesFreeService_ServiceDesc is the grpc.ServiceDesc for FreightTemplatesFreeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FreightTemplatesFreeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.FreightTemplatesFreeService",
	HandlerType: (*FreightTemplatesFreeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FreightTemplatesFreeService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FreightTemplatesFreeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FreightTemplatesFreeService_Delete_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _FreightTemplatesFreeService_FindPageList_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _FreightTemplatesFreeService_FindById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/order/freight_templates_free.proto",
}
