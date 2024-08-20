// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/shopping_admin/product_spec.proto

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
	ProductSpecService_Create_FullMethodName                      = "/template.ProductSpecService/Create"
	ProductSpecService_Update_FullMethodName                      = "/template.ProductSpecService/Update"
	ProductSpecService_Delete_FullMethodName                      = "/template.ProductSpecService/Delete"
	ProductSpecService_FindById_FullMethodName                    = "/template.ProductSpecService/FindById"
	ProductSpecService_FindPageList_FullMethodName                = "/template.ProductSpecService/FindPageList"
	ProductSpecService_FindListByProductCategoryId_FullMethodName = "/template.ProductSpecService/FindListByProductCategoryId"
)

// ProductSpecServiceClient is the client API for ProductSpecService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductSpecServiceClient interface {
	// 新增
	Create(ctx context.Context, in *ProductSpec, opts ...grpc.CallOption) (*ProductSpecResponse, error)
	// 修改
	Update(ctx context.Context, in *ProductSpec, opts ...grpc.CallOption) (*ProductSpecResponse, error)
	// 删除
	Delete(ctx context.Context, in *ProductSpecIds, opts ...grpc.CallOption) (*ProductSpecResponse, error)
	// 查询 详情
	FindById(ctx context.Context, in *ProductSpecIds, opts ...grpc.CallOption) (*ProductSpecResponse, error)
	// 查询 分页
	FindPageList(ctx context.Context, in *ProductSpecPageAuthQuery, opts ...grpc.CallOption) (*ProductSpecResponse, error)
	FindListByProductCategoryId(ctx context.Context, in *ProductSpecIds, opts ...grpc.CallOption) (*ProductSpecResponse, error)
}

type productSpecServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductSpecServiceClient(cc grpc.ClientConnInterface) ProductSpecServiceClient {
	return &productSpecServiceClient{cc}
}

func (c *productSpecServiceClient) Create(ctx context.Context, in *ProductSpec, opts ...grpc.CallOption) (*ProductSpecResponse, error) {
	out := new(ProductSpecResponse)
	err := c.cc.Invoke(ctx, ProductSpecService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSpecServiceClient) Update(ctx context.Context, in *ProductSpec, opts ...grpc.CallOption) (*ProductSpecResponse, error) {
	out := new(ProductSpecResponse)
	err := c.cc.Invoke(ctx, ProductSpecService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSpecServiceClient) Delete(ctx context.Context, in *ProductSpecIds, opts ...grpc.CallOption) (*ProductSpecResponse, error) {
	out := new(ProductSpecResponse)
	err := c.cc.Invoke(ctx, ProductSpecService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSpecServiceClient) FindById(ctx context.Context, in *ProductSpecIds, opts ...grpc.CallOption) (*ProductSpecResponse, error) {
	out := new(ProductSpecResponse)
	err := c.cc.Invoke(ctx, ProductSpecService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSpecServiceClient) FindPageList(ctx context.Context, in *ProductSpecPageAuthQuery, opts ...grpc.CallOption) (*ProductSpecResponse, error) {
	out := new(ProductSpecResponse)
	err := c.cc.Invoke(ctx, ProductSpecService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSpecServiceClient) FindListByProductCategoryId(ctx context.Context, in *ProductSpecIds, opts ...grpc.CallOption) (*ProductSpecResponse, error) {
	out := new(ProductSpecResponse)
	err := c.cc.Invoke(ctx, ProductSpecService_FindListByProductCategoryId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductSpecServiceServer is the server API for ProductSpecService service.
// All implementations must embed UnimplementedProductSpecServiceServer
// for forward compatibility
type ProductSpecServiceServer interface {
	// 新增
	Create(context.Context, *ProductSpec) (*ProductSpecResponse, error)
	// 修改
	Update(context.Context, *ProductSpec) (*ProductSpecResponse, error)
	// 删除
	Delete(context.Context, *ProductSpecIds) (*ProductSpecResponse, error)
	// 查询 详情
	FindById(context.Context, *ProductSpecIds) (*ProductSpecResponse, error)
	// 查询 分页
	FindPageList(context.Context, *ProductSpecPageAuthQuery) (*ProductSpecResponse, error)
	FindListByProductCategoryId(context.Context, *ProductSpecIds) (*ProductSpecResponse, error)
	mustEmbedUnimplementedProductSpecServiceServer()
}

// UnimplementedProductSpecServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductSpecServiceServer struct {
}

func (UnimplementedProductSpecServiceServer) Create(context.Context, *ProductSpec) (*ProductSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductSpecServiceServer) Update(context.Context, *ProductSpec) (*ProductSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductSpecServiceServer) Delete(context.Context, *ProductSpecIds) (*ProductSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProductSpecServiceServer) FindById(context.Context, *ProductSpecIds) (*ProductSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedProductSpecServiceServer) FindPageList(context.Context, *ProductSpecPageAuthQuery) (*ProductSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedProductSpecServiceServer) FindListByProductCategoryId(context.Context, *ProductSpecIds) (*ProductSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListByProductCategoryId not implemented")
}
func (UnimplementedProductSpecServiceServer) mustEmbedUnimplementedProductSpecServiceServer() {}

// UnsafeProductSpecServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductSpecServiceServer will
// result in compilation errors.
type UnsafeProductSpecServiceServer interface {
	mustEmbedUnimplementedProductSpecServiceServer()
}

func RegisterProductSpecServiceServer(s grpc.ServiceRegistrar, srv ProductSpecServiceServer) {
	s.RegisterService(&ProductSpecService_ServiceDesc, srv)
}

func _ProductSpecService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductSpecServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductSpecService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductSpecServiceServer).Create(ctx, req.(*ProductSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductSpecService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductSpecServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductSpecService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductSpecServiceServer).Update(ctx, req.(*ProductSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductSpecService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSpecIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductSpecServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductSpecService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductSpecServiceServer).Delete(ctx, req.(*ProductSpecIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductSpecService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSpecIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductSpecServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductSpecService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductSpecServiceServer).FindById(ctx, req.(*ProductSpecIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductSpecService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSpecPageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductSpecServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductSpecService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductSpecServiceServer).FindPageList(ctx, req.(*ProductSpecPageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductSpecService_FindListByProductCategoryId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSpecIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductSpecServiceServer).FindListByProductCategoryId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductSpecService_FindListByProductCategoryId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductSpecServiceServer).FindListByProductCategoryId(ctx, req.(*ProductSpecIds))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductSpecService_ServiceDesc is the grpc.ServiceDesc for ProductSpecService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductSpecService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template.ProductSpecService",
	HandlerType: (*ProductSpecServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductSpecService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductSpecService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductSpecService_Delete_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _ProductSpecService_FindById_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _ProductSpecService_FindPageList_Handler,
		},
		{
			MethodName: "FindListByProductCategoryId",
			Handler:    _ProductSpecService_FindListByProductCategoryId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/shopping_admin/product_spec.proto",
}
