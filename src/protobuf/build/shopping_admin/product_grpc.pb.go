// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/shopping_admin/product.proto

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
	ProductService_Create_FullMethodName               = "/shopping_admin.ProductService/Create"
	ProductService_Update_FullMethodName               = "/shopping_admin.ProductService/Update"
	ProductService_Delete_FullMethodName               = "/shopping_admin.ProductService/Delete"
	ProductService_FindById_FullMethodName             = "/shopping_admin.ProductService/FindById"
	ProductService_FindPageList_FullMethodName         = "/shopping_admin.ProductService/FindPageList"
	ProductService_FindListByProductIds_FullMethodName = "/shopping_admin.ProductService/FindListByProductIds"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	// 新增商品
	Create(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductResponse, error)
	// 新增商品
	Update(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductResponse, error)
	// 删除商品
	Delete(ctx context.Context, in *ProductIds, opts ...grpc.CallOption) (*ProductResponse, error)
	// 查询商品详情
	FindById(ctx context.Context, in *ProductIds, opts ...grpc.CallOption) (*ProductResponse, error)
	// 查询商品分页列表
	FindPageList(ctx context.Context, in *ProductPageAuthQuery, opts ...grpc.CallOption) (*ProductResponse, error)
	// 根据id查询商品列表（包括sku反查）
	FindListByProductIds(ctx context.Context, in *ProductRequestParamMicroList, opts ...grpc.CallOption) (*ProductParamsResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) Create(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, ProductService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Update(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, ProductService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Delete(ctx context.Context, in *ProductIds, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, ProductService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindById(ctx context.Context, in *ProductIds, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, ProductService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindPageList(ctx context.Context, in *ProductPageAuthQuery, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, ProductService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindListByProductIds(ctx context.Context, in *ProductRequestParamMicroList, opts ...grpc.CallOption) (*ProductParamsResponse, error) {
	out := new(ProductParamsResponse)
	err := c.cc.Invoke(ctx, ProductService_FindListByProductIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	// 新增商品
	Create(context.Context, *Product) (*ProductResponse, error)
	// 新增商品
	Update(context.Context, *Product) (*ProductResponse, error)
	// 删除商品
	Delete(context.Context, *ProductIds) (*ProductResponse, error)
	// 查询商品详情
	FindById(context.Context, *ProductIds) (*ProductResponse, error)
	// 查询商品分页列表
	FindPageList(context.Context, *ProductPageAuthQuery) (*ProductResponse, error)
	// 根据id查询商品列表（包括sku反查）
	FindListByProductIds(context.Context, *ProductRequestParamMicroList) (*ProductParamsResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) Create(context.Context, *Product) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductServiceServer) Update(context.Context, *Product) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductServiceServer) Delete(context.Context, *ProductIds) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProductServiceServer) FindById(context.Context, *ProductIds) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedProductServiceServer) FindPageList(context.Context, *ProductPageAuthQuery) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedProductServiceServer) FindListByProductIds(context.Context, *ProductRequestParamMicroList) (*ProductParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListByProductIds not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Create(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Update(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Delete(ctx, req.(*ProductIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindById(ctx, req.(*ProductIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindPageList(ctx, req.(*ProductPageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindListByProductIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequestParamMicroList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindListByProductIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindListByProductIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindListByProductIds(ctx, req.(*ProductRequestParamMicroList))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shopping_admin.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductService_Delete_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _ProductService_FindById_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _ProductService_FindPageList_Handler,
		},
		{
			MethodName: "FindListByProductIds",
			Handler:    _ProductService_FindListByProductIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/shopping_admin/product.proto",
}