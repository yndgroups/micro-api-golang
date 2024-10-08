// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/order/freight_templates_region.proto

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
	FreightTemplatesRegionService_Create_FullMethodName       = "/order.FreightTemplatesRegionService/Create"
	FreightTemplatesRegionService_Update_FullMethodName       = "/order.FreightTemplatesRegionService/Update"
	FreightTemplatesRegionService_Delete_FullMethodName       = "/order.FreightTemplatesRegionService/Delete"
	FreightTemplatesRegionService_FindPageList_FullMethodName = "/order.FreightTemplatesRegionService/FindPageList"
	FreightTemplatesRegionService_FindById_FullMethodName     = "/order.FreightTemplatesRegionService/FindById"
)

// FreightTemplatesRegionServiceClient is the client API for FreightTemplatesRegionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FreightTemplatesRegionServiceClient interface {
	// Create 区域运费新增信息
	Create(ctx context.Context, in *FreightTemplatesRegion, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error)
	// Update 更新区域运费信息
	Update(ctx context.Context, in *FreightTemplatesRegion, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error)
	// Delete 删除区域运费信息
	Delete(ctx context.Context, in *FreightTemplatesRegionIds, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error)
	// FindPageList 查询区域运费列表
	FindPageList(ctx context.Context, in *FreightTemplatesRegionPageAuthQuery, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error)
	// FindById 查询区域运费详情
	FindById(ctx context.Context, in *FreightTemplatesRegionIds, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error)
}

type freightTemplatesRegionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreightTemplatesRegionServiceClient(cc grpc.ClientConnInterface) FreightTemplatesRegionServiceClient {
	return &freightTemplatesRegionServiceClient{cc}
}

func (c *freightTemplatesRegionServiceClient) Create(ctx context.Context, in *FreightTemplatesRegion, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error) {
	out := new(FreightTemplatesRegionResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesRegionService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesRegionServiceClient) Update(ctx context.Context, in *FreightTemplatesRegion, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error) {
	out := new(FreightTemplatesRegionResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesRegionService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesRegionServiceClient) Delete(ctx context.Context, in *FreightTemplatesRegionIds, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error) {
	out := new(FreightTemplatesRegionResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesRegionService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesRegionServiceClient) FindPageList(ctx context.Context, in *FreightTemplatesRegionPageAuthQuery, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error) {
	out := new(FreightTemplatesRegionResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesRegionService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightTemplatesRegionServiceClient) FindById(ctx context.Context, in *FreightTemplatesRegionIds, opts ...grpc.CallOption) (*FreightTemplatesRegionResponse, error) {
	out := new(FreightTemplatesRegionResponse)
	err := c.cc.Invoke(ctx, FreightTemplatesRegionService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreightTemplatesRegionServiceServer is the server API for FreightTemplatesRegionService service.
// All implementations must embed UnimplementedFreightTemplatesRegionServiceServer
// for forward compatibility
type FreightTemplatesRegionServiceServer interface {
	// Create 区域运费新增信息
	Create(context.Context, *FreightTemplatesRegion) (*FreightTemplatesRegionResponse, error)
	// Update 更新区域运费信息
	Update(context.Context, *FreightTemplatesRegion) (*FreightTemplatesRegionResponse, error)
	// Delete 删除区域运费信息
	Delete(context.Context, *FreightTemplatesRegionIds) (*FreightTemplatesRegionResponse, error)
	// FindPageList 查询区域运费列表
	FindPageList(context.Context, *FreightTemplatesRegionPageAuthQuery) (*FreightTemplatesRegionResponse, error)
	// FindById 查询区域运费详情
	FindById(context.Context, *FreightTemplatesRegionIds) (*FreightTemplatesRegionResponse, error)
	mustEmbedUnimplementedFreightTemplatesRegionServiceServer()
}

// UnimplementedFreightTemplatesRegionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFreightTemplatesRegionServiceServer struct {
}

func (UnimplementedFreightTemplatesRegionServiceServer) Create(context.Context, *FreightTemplatesRegion) (*FreightTemplatesRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFreightTemplatesRegionServiceServer) Update(context.Context, *FreightTemplatesRegion) (*FreightTemplatesRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFreightTemplatesRegionServiceServer) Delete(context.Context, *FreightTemplatesRegionIds) (*FreightTemplatesRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFreightTemplatesRegionServiceServer) FindPageList(context.Context, *FreightTemplatesRegionPageAuthQuery) (*FreightTemplatesRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedFreightTemplatesRegionServiceServer) FindById(context.Context, *FreightTemplatesRegionIds) (*FreightTemplatesRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedFreightTemplatesRegionServiceServer) mustEmbedUnimplementedFreightTemplatesRegionServiceServer() {
}

// UnsafeFreightTemplatesRegionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FreightTemplatesRegionServiceServer will
// result in compilation errors.
type UnsafeFreightTemplatesRegionServiceServer interface {
	mustEmbedUnimplementedFreightTemplatesRegionServiceServer()
}

func RegisterFreightTemplatesRegionServiceServer(s grpc.ServiceRegistrar, srv FreightTemplatesRegionServiceServer) {
	s.RegisterService(&FreightTemplatesRegionService_ServiceDesc, srv)
}

func _FreightTemplatesRegionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesRegion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesRegionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesRegionService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesRegionServiceServer).Create(ctx, req.(*FreightTemplatesRegion))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesRegionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesRegion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesRegionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesRegionService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesRegionServiceServer).Update(ctx, req.(*FreightTemplatesRegion))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesRegionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesRegionIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesRegionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesRegionService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesRegionServiceServer).Delete(ctx, req.(*FreightTemplatesRegionIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesRegionService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesRegionPageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesRegionServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesRegionService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesRegionServiceServer).FindPageList(ctx, req.(*FreightTemplatesRegionPageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightTemplatesRegionService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreightTemplatesRegionIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightTemplatesRegionServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreightTemplatesRegionService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightTemplatesRegionServiceServer).FindById(ctx, req.(*FreightTemplatesRegionIds))
	}
	return interceptor(ctx, in, info, handler)
}

// FreightTemplatesRegionService_ServiceDesc is the grpc.ServiceDesc for FreightTemplatesRegionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FreightTemplatesRegionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.FreightTemplatesRegionService",
	HandlerType: (*FreightTemplatesRegionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FreightTemplatesRegionService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FreightTemplatesRegionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FreightTemplatesRegionService_Delete_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _FreightTemplatesRegionService_FindPageList_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _FreightTemplatesRegionService_FindById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/order/freight_templates_region.proto",
}
