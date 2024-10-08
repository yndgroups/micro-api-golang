// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/common/department.proto

package common

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	global "protobuf/build/global"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DepartmentService_Create_FullMethodName   = "/common.DepartmentService/Create"
	DepartmentService_Update_FullMethodName   = "/common.DepartmentService/Update"
	DepartmentService_Delete_FullMethodName   = "/common.DepartmentService/Delete"
	DepartmentService_FindById_FullMethodName = "/common.DepartmentService/FindById"
	DepartmentService_FindTree_FullMethodName = "/common.DepartmentService/FindTree"
)

// DepartmentServiceClient is the client API for DepartmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepartmentServiceClient interface {
	// 新增部门
	Create(ctx context.Context, in *Department, opts ...grpc.CallOption) (*DepartmentResponse, error)
	// 修改部门
	Update(ctx context.Context, in *Department, opts ...grpc.CallOption) (*DepartmentResponse, error)
	// 查询部门
	Delete(ctx context.Context, in *DepartmentIds, opts ...grpc.CallOption) (*DepartmentResponse, error)
	// 查询部门详情
	FindById(ctx context.Context, in *DepartmentIds, opts ...grpc.CallOption) (*DepartmentResponse, error)
	// 查询部门数据
	FindTree(ctx context.Context, in *global.Auth, opts ...grpc.CallOption) (*DepartmentResponse, error)
}

type departmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentServiceClient(cc grpc.ClientConnInterface) DepartmentServiceClient {
	return &departmentServiceClient{cc}
}

func (c *departmentServiceClient) Create(ctx context.Context, in *Department, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	out := new(DepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) Update(ctx context.Context, in *Department, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	out := new(DepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) Delete(ctx context.Context, in *DepartmentIds, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	out := new(DepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) FindById(ctx context.Context, in *DepartmentIds, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	out := new(DepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) FindTree(ctx context.Context, in *global.Auth, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	out := new(DepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_FindTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentServiceServer is the server API for DepartmentService service.
// All implementations must embed UnimplementedDepartmentServiceServer
// for forward compatibility
type DepartmentServiceServer interface {
	// 新增部门
	Create(context.Context, *Department) (*DepartmentResponse, error)
	// 修改部门
	Update(context.Context, *Department) (*DepartmentResponse, error)
	// 查询部门
	Delete(context.Context, *DepartmentIds) (*DepartmentResponse, error)
	// 查询部门详情
	FindById(context.Context, *DepartmentIds) (*DepartmentResponse, error)
	// 查询部门数据
	FindTree(context.Context, *global.Auth) (*DepartmentResponse, error)
	mustEmbedUnimplementedDepartmentServiceServer()
}

// UnimplementedDepartmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepartmentServiceServer struct {
}

func (UnimplementedDepartmentServiceServer) Create(context.Context, *Department) (*DepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDepartmentServiceServer) Update(context.Context, *Department) (*DepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDepartmentServiceServer) Delete(context.Context, *DepartmentIds) (*DepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDepartmentServiceServer) FindById(context.Context, *DepartmentIds) (*DepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedDepartmentServiceServer) FindTree(context.Context, *global.Auth) (*DepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTree not implemented")
}
func (UnimplementedDepartmentServiceServer) mustEmbedUnimplementedDepartmentServiceServer() {}

// UnsafeDepartmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepartmentServiceServer will
// result in compilation errors.
type UnsafeDepartmentServiceServer interface {
	mustEmbedUnimplementedDepartmentServiceServer()
}

func RegisterDepartmentServiceServer(s grpc.ServiceRegistrar, srv DepartmentServiceServer) {
	s.RegisterService(&DepartmentService_ServiceDesc, srv)
}

func _DepartmentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Department)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).Create(ctx, req.(*Department))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Department)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).Update(ctx, req.(*Department))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepartmentIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).Delete(ctx, req.(*DepartmentIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepartmentIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).FindById(ctx, req.(*DepartmentIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_FindTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(global.Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).FindTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_FindTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).FindTree(ctx, req.(*global.Auth))
	}
	return interceptor(ctx, in, info, handler)
}

// DepartmentService_ServiceDesc is the grpc.ServiceDesc for DepartmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepartmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.DepartmentService",
	HandlerType: (*DepartmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DepartmentService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DepartmentService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DepartmentService_Delete_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _DepartmentService_FindById_Handler,
		},
		{
			MethodName: "FindTree",
			Handler:    _DepartmentService_FindTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/common/department.proto",
}
