// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/common/dict.proto

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
	DictService_Create_FullMethodName             = "/common.DictService/Create"
	DictService_Update_FullMethodName             = "/common.DictService/Update"
	DictService_Delete_FullMethodName             = "/common.DictService/Delete"
	DictService_FindById_FullMethodName           = "/common.DictService/FindById"
	DictService_FindListByParentId_FullMethodName = "/common.DictService/FindListByParentId"
	DictService_FindPageList_FullMethodName       = "/common.DictService/FindPageList"
)

// DictServiceClient is the client API for DictService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictServiceClient interface {
	// 新增字典
	Create(ctx context.Context, in *Dict, opts ...grpc.CallOption) (*DictResponse, error)
	// 修改字典
	Update(ctx context.Context, in *Dict, opts ...grpc.CallOption) (*DictResponse, error)
	// 修改字典
	Delete(ctx context.Context, in *DictIds, opts ...grpc.CallOption) (*DictResponse, error)
	// 查询字典详情
	FindById(ctx context.Context, in *DictIds, opts ...grpc.CallOption) (*DictResponse, error)
	// 查询字典列表
	FindListByParentId(ctx context.Context, in *DictIds, opts ...grpc.CallOption) (*DictResponse, error)
	// 查询字典分页
	FindPageList(ctx context.Context, in *DictPageAuthQuery, opts ...grpc.CallOption) (*DictResponse, error)
}

type dictServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDictServiceClient(cc grpc.ClientConnInterface) DictServiceClient {
	return &dictServiceClient{cc}
}

func (c *dictServiceClient) Create(ctx context.Context, in *Dict, opts ...grpc.CallOption) (*DictResponse, error) {
	out := new(DictResponse)
	err := c.cc.Invoke(ctx, DictService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) Update(ctx context.Context, in *Dict, opts ...grpc.CallOption) (*DictResponse, error) {
	out := new(DictResponse)
	err := c.cc.Invoke(ctx, DictService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) Delete(ctx context.Context, in *DictIds, opts ...grpc.CallOption) (*DictResponse, error) {
	out := new(DictResponse)
	err := c.cc.Invoke(ctx, DictService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) FindById(ctx context.Context, in *DictIds, opts ...grpc.CallOption) (*DictResponse, error) {
	out := new(DictResponse)
	err := c.cc.Invoke(ctx, DictService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) FindListByParentId(ctx context.Context, in *DictIds, opts ...grpc.CallOption) (*DictResponse, error) {
	out := new(DictResponse)
	err := c.cc.Invoke(ctx, DictService_FindListByParentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) FindPageList(ctx context.Context, in *DictPageAuthQuery, opts ...grpc.CallOption) (*DictResponse, error) {
	out := new(DictResponse)
	err := c.cc.Invoke(ctx, DictService_FindPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictServiceServer is the server API for DictService service.
// All implementations must embed UnimplementedDictServiceServer
// for forward compatibility
type DictServiceServer interface {
	// 新增字典
	Create(context.Context, *Dict) (*DictResponse, error)
	// 修改字典
	Update(context.Context, *Dict) (*DictResponse, error)
	// 修改字典
	Delete(context.Context, *DictIds) (*DictResponse, error)
	// 查询字典详情
	FindById(context.Context, *DictIds) (*DictResponse, error)
	// 查询字典列表
	FindListByParentId(context.Context, *DictIds) (*DictResponse, error)
	// 查询字典分页
	FindPageList(context.Context, *DictPageAuthQuery) (*DictResponse, error)
	mustEmbedUnimplementedDictServiceServer()
}

// UnimplementedDictServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDictServiceServer struct {
}

func (UnimplementedDictServiceServer) Create(context.Context, *Dict) (*DictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDictServiceServer) Update(context.Context, *Dict) (*DictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDictServiceServer) Delete(context.Context, *DictIds) (*DictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDictServiceServer) FindById(context.Context, *DictIds) (*DictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedDictServiceServer) FindListByParentId(context.Context, *DictIds) (*DictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListByParentId not implemented")
}
func (UnimplementedDictServiceServer) FindPageList(context.Context, *DictPageAuthQuery) (*DictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPageList not implemented")
}
func (UnimplementedDictServiceServer) mustEmbedUnimplementedDictServiceServer() {}

// UnsafeDictServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictServiceServer will
// result in compilation errors.
type UnsafeDictServiceServer interface {
	mustEmbedUnimplementedDictServiceServer()
}

func RegisterDictServiceServer(s grpc.ServiceRegistrar, srv DictServiceServer) {
	s.RegisterService(&DictService_ServiceDesc, srv)
}

func _DictService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dict)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).Create(ctx, req.(*Dict))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dict)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).Update(ctx, req.(*Dict))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).Delete(ctx, req.(*DictIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).FindById(ctx, req.(*DictIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_FindListByParentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).FindListByParentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_FindListByParentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).FindListByParentId(ctx, req.(*DictIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_FindPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictPageAuthQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).FindPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_FindPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).FindPageList(ctx, req.(*DictPageAuthQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// DictService_ServiceDesc is the grpc.ServiceDesc for DictService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DictService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.DictService",
	HandlerType: (*DictServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DictService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DictService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DictService_Delete_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _DictService_FindById_Handler,
		},
		{
			MethodName: "FindListByParentId",
			Handler:    _DictService_FindListByParentId_Handler,
		},
		{
			MethodName: "FindPageList",
			Handler:    _DictService_FindPageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/common/dict.proto",
}