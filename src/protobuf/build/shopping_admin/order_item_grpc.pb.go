// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/shopping_admin/order_item.proto

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
	OrderItemService_Create_FullMethodName       = "/shopping_admin.OrderItemService/Create"
	OrderItemService_Update_FullMethodName       = "/shopping_admin.OrderItemService/Update"
	OrderItemService_Delete_FullMethodName       = "/shopping_admin.OrderItemService/Delete"
	OrderItemService_FindById_FullMethodName     = "/shopping_admin.OrderItemService/FindById"
	OrderItemService_FindPageList_FullMethodName = "/shopping_admin.OrderItemService/FindPageList"
)

// OrderItemServiceClient is the client API for OrderItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderItemServiceClient interface {
	// 新增
	Create(ctx context.Context, in *OrderItem, opts ...grpc.CallOption) (*OrderItemResponse, error)
	// 修改
	Update(ctx context.Context, in *OrderItem, opts ...grpc.CallOption) (*OrderItemResponse, error)
	// 删除
	Delete(ctx context.Context, in *OrderItemIds, opts ...grpc.CallOption) (*OrderItemResponse, error)
	// 查询 详情
	FindById(ctx context.Context, in *OrderItemIds, opts ...grpc.CallOption) (*OrderItemResponse, error)
	// 查询 分页
	FindPageList(ctx context.Context, in *OrderItemPageAuthQuery, opts ...grpc.CallOption) (*OrderItemResponse, error)
}

type orderItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderItemServiceClient(cc grpc.ClientConnInterface) OrderItemServiceClient {
	return &orderItemServiceClient{cc}
}

func (c *orderItemServiceClient) Create(ctx context.Context, in *OrderItem, opts ...grpc.CallOption) (*OrderItemResponse, error) {
	out := new(OrderItemResponse)
	err := c.cc.Invoke(ctx, OrderItemService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemServiceClient) Update(ctx context.Context, in *OrderItem, opts ...grpc.CallOption) (*OrderItemResponse, error) {
	out := new(OrderItemResponse)
	err := c.cc.Invoke(ctx, OrderItemService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemServiceClient) Delete(ctx context.Context, in *OrderItemIds, opts ...grpc.CallOption) (*OrderItemResponse, error) {
	out := new(OrderItemResponse)
	err := c.cc.Invoke(ctx, OrderItemService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemServiceClient) FindById(ctx context.Context, in *OrderItemIds, opts ...grpc.CallOption) (*OrderItemResponse, error) {
	out := new(OrderItemResponse)
	err := c.cc.Invoke(ctx, OrderItemService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemServiceClient) FindPageList(ctx context.Context, in *OrderItemPageAuthQuery, opts ...grpc.CallOption) (*OrderItemResponse, error) {
	out := new(OrderItemResponse)
	err := c.cc.Invoke(ctx, OrderItemService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderItemServiceServer is the server API for OrderItemService service.
// All implementations must embed UnimplementedOrderItemServiceServer
// for forward compatibility
type OrderItemServiceServer interface {
	// 新增
	Create(context.Context, *OrderItem) (*OrderItemResponse, error)
	// 修改
	Update(context.Context, *OrderItem) (*OrderItemResponse, error)
	// 删除
	Delete(context.Context, *OrderItemIds) (*OrderItemResponse, error)
	// 查询 详情
	FindById(context.Context, *OrderItemIds) (*OrderItemResponse, error)
	// 查询 分页
	FindPageList(context.Context, *OrderItemPageAuthQuery) (*OrderItemResponse, error)
	mustEmbedUnimplementedOrderItemServiceServer()
}

// UnimplementedOrderItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderItemServiceServer struct {
}

func (UnimplementedOrderItemServiceServer) Create(context.Context, *OrderItem) (*OrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderItemServiceServer) Update(context.Context, *OrderItem) (*OrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderItemServiceServer) Delete(context.Context, *OrderItemIds) (*OrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOrderItemServiceServer) FindById(context.Context, *OrderItemIds) (*OrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedOrderItemServiceServer) FindPageList(context.Context, *OrderItemPageAuthQuery) (*OrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedOrderItemServiceServer) mustEmbedUnimplementedOrderItemServiceServer() {}

// UnsafeOrderItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderItemServiceServer will
// result in compilation errors.
type UnsafeOrderItemServiceServer interface {
	mustEmbedUnimplementedOrderItemServiceServer()
}

func RegisterOrderItemServiceServer(s grpc.ServiceRegistrar, srv OrderItemServiceServer) {
	s.RegisterService(&OrderItemService_ServiceDesc, srv)
}

func _OrderItemService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderItemService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemServiceServer).Create(ctx, req.(*OrderItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderItemService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemServiceServer).Update(ctx, req.(*OrderItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderItemService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemServiceServer).Delete(ctx, req.(*OrderItemIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderItemService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemServiceServer).FindById(ctx, req.(*OrderItemIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemPageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderItemService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemServiceServer).FindPageList(ctx, req.(*OrderItemPageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderItemService_ServiceDesc is the grpc.ServiceDesc for OrderItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shopping_admin.OrderItemService",
	HandlerType: (*OrderItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderItemService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderItemService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrderItemService_Delete_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _OrderItemService_FindById_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _OrderItemService_FindPageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/shopping_admin/order_item.proto",
}
