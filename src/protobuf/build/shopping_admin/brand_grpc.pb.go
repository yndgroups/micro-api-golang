// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/shopping_admin/brand.proto

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
	BrandService_Create_FullMethodName                      = "/template.BrandService/Create"
	BrandService_Update_FullMethodName                      = "/template.BrandService/Update"
	BrandService_Delete_FullMethodName                      = "/template.BrandService/Delete"
	BrandService_FindById_FullMethodName                    = "/template.BrandService/FindById"
	BrandService_FindPageList_FullMethodName                = "/template.BrandService/FindPageList"
	BrandService_FindListByProductCategoryId_FullMethodName = "/template.BrandService/FindListByProductCategoryId"
)

// BrandServiceClient is the client API for BrandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrandServiceClient interface {
	// 新增
	Create(ctx context.Context, in *Brand, opts ...grpc.CallOption) (*BrandResponse, error)
	// 修改
	Update(ctx context.Context, in *Brand, opts ...grpc.CallOption) (*BrandResponse, error)
	// 删除
	Delete(ctx context.Context, in *BrandIds, opts ...grpc.CallOption) (*BrandResponse, error)
	// 查询 详情
	FindById(ctx context.Context, in *BrandIds, opts ...grpc.CallOption) (*BrandResponse, error)
	// 查询 分页
	FindPageList(ctx context.Context, in *BrandPageAuthQuery, opts ...grpc.CallOption) (*BrandResponse, error)
	// 根据分类查询品牌
	FindListByProductCategoryId(ctx context.Context, in *BrandIds, opts ...grpc.CallOption) (*BrandResponse, error)
}

type brandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrandServiceClient(cc grpc.ClientConnInterface) BrandServiceClient {
	return &brandServiceClient{cc}
}

func (c *brandServiceClient) Create(ctx context.Context, in *Brand, opts ...grpc.CallOption) (*BrandResponse, error) {
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, BrandService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) Update(ctx context.Context, in *Brand, opts ...grpc.CallOption) (*BrandResponse, error) {
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, BrandService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) Delete(ctx context.Context, in *BrandIds, opts ...grpc.CallOption) (*BrandResponse, error) {
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, BrandService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) FindById(ctx context.Context, in *BrandIds, opts ...grpc.CallOption) (*BrandResponse, error) {
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, BrandService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) FindPageList(ctx context.Context, in *BrandPageAuthQuery, opts ...grpc.CallOption) (*BrandResponse, error) {
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, BrandService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) FindListByProductCategoryId(ctx context.Context, in *BrandIds, opts ...grpc.CallOption) (*BrandResponse, error) {
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, BrandService_FindListByProductCategoryId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrandServiceServer is the server API for BrandService service.
// All implementations must embed UnimplementedBrandServiceServer
// for forward compatibility
type BrandServiceServer interface {
	// 新增
	Create(context.Context, *Brand) (*BrandResponse, error)
	// 修改
	Update(context.Context, *Brand) (*BrandResponse, error)
	// 删除
	Delete(context.Context, *BrandIds) (*BrandResponse, error)
	// 查询 详情
	FindById(context.Context, *BrandIds) (*BrandResponse, error)
	// 查询 分页
	FindPageList(context.Context, *BrandPageAuthQuery) (*BrandResponse, error)
	// 根据分类查询品牌
	FindListByProductCategoryId(context.Context, *BrandIds) (*BrandResponse, error)
	mustEmbedUnimplementedBrandServiceServer()
}

// UnimplementedBrandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBrandServiceServer struct {
}

func (UnimplementedBrandServiceServer) Create(context.Context, *Brand) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBrandServiceServer) Update(context.Context, *Brand) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBrandServiceServer) Delete(context.Context, *BrandIds) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBrandServiceServer) FindById(context.Context, *BrandIds) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedBrandServiceServer) FindPageList(context.Context, *BrandPageAuthQuery) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedBrandServiceServer) FindListByProductCategoryId(context.Context, *BrandIds) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListByProductCategoryId not implemented")
}
func (UnimplementedBrandServiceServer) mustEmbedUnimplementedBrandServiceServer() {}

// UnsafeBrandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrandServiceServer will
// result in compilation errors.
type UnsafeBrandServiceServer interface {
	mustEmbedUnimplementedBrandServiceServer()
}

func RegisterBrandServiceServer(s grpc.ServiceRegistrar, srv BrandServiceServer) {
	s.RegisterService(&BrandService_ServiceDesc, srv)
}

func _BrandService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Brand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BrandService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).Create(ctx, req.(*Brand))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Brand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BrandService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).Update(ctx, req.(*Brand))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BrandService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).Delete(ctx, req.(*BrandIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BrandService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).FindById(ctx, req.(*BrandIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandPageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BrandService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).FindPageList(ctx, req.(*BrandPageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_FindListByProductCategoryId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).FindListByProductCategoryId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BrandService_FindListByProductCategoryId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).FindListByProductCategoryId(ctx, req.(*BrandIds))
	}
	return interceptor(ctx, in, info, handler)
}

// BrandService_ServiceDesc is the grpc.ServiceDesc for BrandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template.BrandService",
	HandlerType: (*BrandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BrandService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BrandService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BrandService_Delete_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _BrandService_FindById_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _BrandService_FindPageList_Handler,
		},
		{
			MethodName: "FindListByProductCategoryId",
			Handler:    _BrandService_FindListByProductCategoryId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/shopping_admin/brand.proto",
}
