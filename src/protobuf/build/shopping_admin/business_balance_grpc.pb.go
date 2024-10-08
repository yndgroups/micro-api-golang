// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/shopping_admin/business_balance.proto

package shopping_admin

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
	BusinessBalanceService_Create_FullMethodName       = "/template.BusinessBalanceService/Create"
	BusinessBalanceService_Update_FullMethodName       = "/template.BusinessBalanceService/Update"
	BusinessBalanceService_Delete_FullMethodName       = "/template.BusinessBalanceService/Delete"
	BusinessBalanceService_FindById_FullMethodName     = "/template.BusinessBalanceService/FindById"
	BusinessBalanceService_FindPageList_FullMethodName = "/template.BusinessBalanceService/FindPageList"
)

// BusinessBalanceServiceClient is the client API for BusinessBalanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusinessBalanceServiceClient interface {
	// 新增
	Create(ctx context.Context, in *BusinessBalance, opts ...grpc.CallOption) (*BusinessBalanceResponse, error)
	// 修改
	Update(ctx context.Context, in *BusinessBalance, opts ...grpc.CallOption) (*BusinessBalanceResponse, error)
	// 删除
	Delete(ctx context.Context, in *BusinessBalanceIds, opts ...grpc.CallOption) (*BusinessBalanceResponse, error)
	// 查询 详情
	FindById(ctx context.Context, in *BusinessBalanceIds, opts ...grpc.CallOption) (*BusinessBalanceResponse, error)
	// 查询 分页
	FindPageList(ctx context.Context, in *BusinessBalancePageAuthQuery, opts ...grpc.CallOption) (*BusinessBalanceResponse, error)
}

type businessBalanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessBalanceServiceClient(cc grpc.ClientConnInterface) BusinessBalanceServiceClient {
	return &businessBalanceServiceClient{cc}
}

func (c *businessBalanceServiceClient) Create(ctx context.Context, in *BusinessBalance, opts ...grpc.CallOption) (*BusinessBalanceResponse, error) {
	out := new(BusinessBalanceResponse)
	err := c.cc.Invoke(ctx, BusinessBalanceService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessBalanceServiceClient) Update(ctx context.Context, in *BusinessBalance, opts ...grpc.CallOption) (*BusinessBalanceResponse, error) {
	out := new(BusinessBalanceResponse)
	err := c.cc.Invoke(ctx, BusinessBalanceService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessBalanceServiceClient) Delete(ctx context.Context, in *BusinessBalanceIds, opts ...grpc.CallOption) (*BusinessBalanceResponse, error) {
	out := new(BusinessBalanceResponse)
	err := c.cc.Invoke(ctx, BusinessBalanceService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessBalanceServiceClient) FindById(ctx context.Context, in *BusinessBalanceIds, opts ...grpc.CallOption) (*BusinessBalanceResponse, error) {
	out := new(BusinessBalanceResponse)
	err := c.cc.Invoke(ctx, BusinessBalanceService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessBalanceServiceClient) FindPageList(ctx context.Context, in *BusinessBalancePageAuthQuery, opts ...grpc.CallOption) (*BusinessBalanceResponse, error) {
	out := new(BusinessBalanceResponse)
	err := c.cc.Invoke(ctx, BusinessBalanceService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessBalanceServiceServer is the server API for BusinessBalanceService service.
// All implementations must embed UnimplementedBusinessBalanceServiceServer
// for forward compatibility
type BusinessBalanceServiceServer interface {
	// 新增
	Create(context.Context, *BusinessBalance) (*BusinessBalanceResponse, error)
	// 修改
	Update(context.Context, *BusinessBalance) (*BusinessBalanceResponse, error)
	// 删除
	Delete(context.Context, *BusinessBalanceIds) (*BusinessBalanceResponse, error)
	// 查询 详情
	FindById(context.Context, *BusinessBalanceIds) (*BusinessBalanceResponse, error)
	// 查询 分页
	FindPageList(context.Context, *BusinessBalancePageAuthQuery) (*BusinessBalanceResponse, error)
	mustEmbedUnimplementedBusinessBalanceServiceServer()
}

// UnimplementedBusinessBalanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBusinessBalanceServiceServer struct {
}

func (UnimplementedBusinessBalanceServiceServer) Create(context.Context, *BusinessBalance) (*BusinessBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBusinessBalanceServiceServer) Update(context.Context, *BusinessBalance) (*BusinessBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBusinessBalanceServiceServer) Delete(context.Context, *BusinessBalanceIds) (*BusinessBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBusinessBalanceServiceServer) FindById(context.Context, *BusinessBalanceIds) (*BusinessBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedBusinessBalanceServiceServer) FindPageList(context.Context, *BusinessBalancePageAuthQuery) (*BusinessBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedBusinessBalanceServiceServer) mustEmbedUnimplementedBusinessBalanceServiceServer() {
}

// UnsafeBusinessBalanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusinessBalanceServiceServer will
// result in compilation errors.
type UnsafeBusinessBalanceServiceServer interface {
	mustEmbedUnimplementedBusinessBalanceServiceServer()
}

func RegisterBusinessBalanceServiceServer(s grpc.ServiceRegistrar, srv BusinessBalanceServiceServer) {
	s.RegisterService(&BusinessBalanceService_ServiceDesc, srv)
}

func _BusinessBalanceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessBalance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessBalanceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessBalanceService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessBalanceServiceServer).Create(ctx, req.(*BusinessBalance))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessBalanceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessBalance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessBalanceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessBalanceService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessBalanceServiceServer).Update(ctx, req.(*BusinessBalance))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessBalanceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessBalanceIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessBalanceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessBalanceService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessBalanceServiceServer).Delete(ctx, req.(*BusinessBalanceIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessBalanceService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessBalanceIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessBalanceServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessBalanceService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessBalanceServiceServer).FindById(ctx, req.(*BusinessBalanceIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessBalanceService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessBalancePageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessBalanceServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessBalanceService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessBalanceServiceServer).FindPageList(ctx, req.(*BusinessBalancePageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// BusinessBalanceService_ServiceDesc is the grpc.ServiceDesc for BusinessBalanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BusinessBalanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template.BusinessBalanceService",
	HandlerType: (*BusinessBalanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BusinessBalanceService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BusinessBalanceService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BusinessBalanceService_Delete_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _BusinessBalanceService_FindById_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _BusinessBalanceService_FindPageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/shopping_admin/business_balance.proto",
}
