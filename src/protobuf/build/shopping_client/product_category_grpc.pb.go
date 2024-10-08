// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/shopping_client/product_category.proto

package shopping_client

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
	ProductCategoryService_FindByParentId_FullMethodName     = "/shopping_client.ProductCategoryService/FindByParentId"
	ProductCategoryService_FindTreeByParentId_FullMethodName = "/shopping_client.ProductCategoryService/FindTreeByParentId"
)

// ProductCategoryServiceClient is the client API for ProductCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCategoryServiceClient interface {
	// 查询 详情
	FindByParentId(ctx context.Context, in *ProductCategoryAuthParams, opts ...grpc.CallOption) (*ProductCategoryResponse, error)
	// 查询 分页
	FindTreeByParentId(ctx context.Context, in *ProductCategoryAuthParams, opts ...grpc.CallOption) (*ProductCategoryResponse, error)
}

type productCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCategoryServiceClient(cc grpc.ClientConnInterface) ProductCategoryServiceClient {
	return &productCategoryServiceClient{cc}
}

func (c *productCategoryServiceClient) FindByParentId(ctx context.Context, in *ProductCategoryAuthParams, opts ...grpc.CallOption) (*ProductCategoryResponse, error) {
	out := new(ProductCategoryResponse)
	err := c.cc.Invoke(ctx, ProductCategoryService_FindByParentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCategoryServiceClient) FindTreeByParentId(ctx context.Context, in *ProductCategoryAuthParams, opts ...grpc.CallOption) (*ProductCategoryResponse, error) {
	out := new(ProductCategoryResponse)
	err := c.cc.Invoke(ctx, ProductCategoryService_FindTreeByParentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCategoryServiceServer is the server API for ProductCategoryService service.
// All implementations must embed UnimplementedProductCategoryServiceServer
// for forward compatibility
type ProductCategoryServiceServer interface {
	// 查询 详情
	FindByParentId(context.Context, *ProductCategoryAuthParams) (*ProductCategoryResponse, error)
	// 查询 分页
	FindTreeByParentId(context.Context, *ProductCategoryAuthParams) (*ProductCategoryResponse, error)
	mustEmbedUnimplementedProductCategoryServiceServer()
}

// UnimplementedProductCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductCategoryServiceServer struct {
}

func (UnimplementedProductCategoryServiceServer) FindByParentId(context.Context, *ProductCategoryAuthParams) (*ProductCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByParentId not implemented")
}
func (UnimplementedProductCategoryServiceServer) FindTreeByParentId(context.Context, *ProductCategoryAuthParams) (*ProductCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTreeByParentId not implemented")
}
func (UnimplementedProductCategoryServiceServer) mustEmbedUnimplementedProductCategoryServiceServer() {
}

// UnsafeProductCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCategoryServiceServer will
// result in compilation errors.
type UnsafeProductCategoryServiceServer interface {
	mustEmbedUnimplementedProductCategoryServiceServer()
}

func RegisterProductCategoryServiceServer(s grpc.ServiceRegistrar, srv ProductCategoryServiceServer) {
	s.RegisterService(&ProductCategoryService_ServiceDesc, srv)
}

func _ProductCategoryService_FindByParentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCategoryAuthParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).FindByParentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_FindByParentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).FindByParentId(ctx, req.(*ProductCategoryAuthParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCategoryService_FindTreeByParentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCategoryAuthParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).FindTreeByParentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_FindTreeByParentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).FindTreeByParentId(ctx, req.(*ProductCategoryAuthParams))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCategoryService_ServiceDesc is the grpc.ServiceDesc for ProductCategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shopping_client.ProductCategoryService",
	HandlerType: (*ProductCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByParentId",
			Handler:    _ProductCategoryService_FindByParentId_Handler,
		},
		{
			MethodName: "FindTreeByParentId",
			Handler:    _ProductCategoryService_FindTreeByParentId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/shopping_client/product_category.proto",
}
