// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/common/area.proto

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
	AreaService_Create_FullMethodName       = "/common.AreaService/Create"
	AreaService_Update_FullMethodName       = "/common.AreaService/Update"
	AreaService_Delete_FullMethodName       = "/common.AreaService/Delete"
	AreaService_FindById_FullMethodName     = "/common.AreaService/FindById"
	AreaService_FindPageList_FullMethodName = "/common.AreaService/FindPageList"
	AreaService_FindTree_FullMethodName     = "/common.AreaService/FindTree"
)

// AreaServiceClient is the client API for AreaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AreaServiceClient interface {
	// 新增
	Create(ctx context.Context, in *Area, opts ...grpc.CallOption) (*AreaResponse, error)
	// 修改
	Update(ctx context.Context, in *Area, opts ...grpc.CallOption) (*AreaResponse, error)
	// 删除
	Delete(ctx context.Context, in *AreaIds, opts ...grpc.CallOption) (*AreaResponse, error)
	// 查询 详情
	FindById(ctx context.Context, in *AreaIds, opts ...grpc.CallOption) (*AreaResponse, error)
	// 查询 分页
	FindPageList(ctx context.Context, in *AreaPageAuthQuery, opts ...grpc.CallOption) (*AreaResponse, error)
	// 查询三级地区
	FindTree(ctx context.Context, in *AreaTags, opts ...grpc.CallOption) (*AreaResponse, error)
}

type areaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAreaServiceClient(cc grpc.ClientConnInterface) AreaServiceClient {
	return &areaServiceClient{cc}
}

func (c *areaServiceClient) Create(ctx context.Context, in *Area, opts ...grpc.CallOption) (*AreaResponse, error) {
	out := new(AreaResponse)
	err := c.cc.Invoke(ctx, AreaService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) Update(ctx context.Context, in *Area, opts ...grpc.CallOption) (*AreaResponse, error) {
	out := new(AreaResponse)
	err := c.cc.Invoke(ctx, AreaService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) Delete(ctx context.Context, in *AreaIds, opts ...grpc.CallOption) (*AreaResponse, error) {
	out := new(AreaResponse)
	err := c.cc.Invoke(ctx, AreaService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) FindById(ctx context.Context, in *AreaIds, opts ...grpc.CallOption) (*AreaResponse, error) {
	out := new(AreaResponse)
	err := c.cc.Invoke(ctx, AreaService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) FindPageList(ctx context.Context, in *AreaPageAuthQuery, opts ...grpc.CallOption) (*AreaResponse, error) {
	out := new(AreaResponse)
	err := c.cc.Invoke(ctx, AreaService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) FindTree(ctx context.Context, in *AreaTags, opts ...grpc.CallOption) (*AreaResponse, error) {
	out := new(AreaResponse)
	err := c.cc.Invoke(ctx, AreaService_FindTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AreaServiceServer is the server API for AreaService service.
// All implementations must embed UnimplementedAreaServiceServer
// for forward compatibility
type AreaServiceServer interface {
	// 新增
	Create(context.Context, *Area) (*AreaResponse, error)
	// 修改
	Update(context.Context, *Area) (*AreaResponse, error)
	// 删除
	Delete(context.Context, *AreaIds) (*AreaResponse, error)
	// 查询 详情
	FindById(context.Context, *AreaIds) (*AreaResponse, error)
	// 查询 分页
	FindPageList(context.Context, *AreaPageAuthQuery) (*AreaResponse, error)
	// 查询三级地区
	FindTree(context.Context, *AreaTags) (*AreaResponse, error)
	mustEmbedUnimplementedAreaServiceServer()
}

// UnimplementedAreaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAreaServiceServer struct {
}

func (UnimplementedAreaServiceServer) Create(context.Context, *Area) (*AreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAreaServiceServer) Update(context.Context, *Area) (*AreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAreaServiceServer) Delete(context.Context, *AreaIds) (*AreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAreaServiceServer) FindById(context.Context, *AreaIds) (*AreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedAreaServiceServer) FindPageList(context.Context, *AreaPageAuthQuery) (*AreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedAreaServiceServer) FindTree(context.Context, *AreaTags) (*AreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTree not implemented")
}
func (UnimplementedAreaServiceServer) mustEmbedUnimplementedAreaServiceServer() {}

// UnsafeAreaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AreaServiceServer will
// result in compilation errors.
type UnsafeAreaServiceServer interface {
	mustEmbedUnimplementedAreaServiceServer()
}

func RegisterAreaServiceServer(s grpc.ServiceRegistrar, srv AreaServiceServer) {
	s.RegisterService(&AreaService_ServiceDesc, srv)
}

func _AreaService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Area)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AreaService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).Create(ctx, req.(*Area))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Area)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AreaService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).Update(ctx, req.(*Area))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AreaService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).Delete(ctx, req.(*AreaIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AreaService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).FindById(ctx, req.(*AreaIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaPageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AreaService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).FindPageList(ctx, req.(*AreaPageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_FindTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaTags)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).FindTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AreaService_FindTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).FindTree(ctx, req.(*AreaTags))
	}
	return interceptor(ctx, in, info, handler)
}

// AreaService_ServiceDesc is the grpc.ServiceDesc for AreaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AreaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.AreaService",
	HandlerType: (*AreaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AreaService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AreaService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AreaService_Delete_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _AreaService_FindById_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _AreaService_FindPageList_Handler,
		},
		{
			MethodName: "FindTree",
			Handler:    _AreaService_FindTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/common/area.proto",
}