// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/shopping_admin/cart.proto

package shopping_admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	global "protobuf/build/global"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 删除状态
	DelStatus bool `protobuf:"varint,1,opt,name=delStatus,proto3" json:"delStatus,omitempty" gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"` // @gotags: gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"
	// 创建人
	CreateBy string `protobuf:"bytes,2,opt,name=createBy,proto3" json:"createBy,omitempty" gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 修改人
	UpdateBy string `protobuf:"bytes,3,opt,name=updateBy,proto3" json:"updateBy,omitempty" gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 创建时间
	CreateTime string `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
	// 修改时间
	UpdateTime string `protobuf:"bytes,5,opt,name=updateTime,proto3" json:"updateTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_cart_proto_rawDescGZIP(), []int{0}
}

func (x *Cart) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *Cart) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *Cart) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Cart) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Cart) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type CartParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartParam) Reset() {
	*x = CartParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartParam) ProtoMessage() {}

func (x *CartParam) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartParam.ProtoReflect.Descriptor instead.
func (*CartParam) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_cart_proto_rawDescGZIP(), []int{1}
}

type CartIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	UserId string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CartIds) Reset() {
	*x = CartIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartIds) ProtoMessage() {}

func (x *CartIds) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartIds.ProtoReflect.Descriptor instead.
func (*CartIds) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_cart_proto_rawDescGZIP(), []int{2}
}

func (x *CartIds) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *CartIds) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// 分页查询加权参数
type CartPageAuthQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页索引
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	// 分页长度
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 查询参数
	Params *CartParam `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	// 用户登录授权信息
	Auth *global.Auth `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty" swaggerignore:"true"` // @gotags: swaggerignore:"true"`
}

func (x *CartPageAuthQuery) Reset() {
	*x = CartPageAuthQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartPageAuthQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartPageAuthQuery) ProtoMessage() {}

func (x *CartPageAuthQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartPageAuthQuery.ProtoReflect.Descriptor instead.
func (*CartPageAuthQuery) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_cart_proto_rawDescGZIP(), []int{3}
}

func (x *CartPageAuthQuery) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *CartPageAuthQuery) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CartPageAuthQuery) GetParams() *CartParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *CartPageAuthQuery) GetAuth() *global.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// 返回结果
type CartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`  // 获取数据数量
	Msg    string  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`       // 错误信息
	Detail *Cart   `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
	List   []*Cart `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`     // 列表
}

func (x *CartResponse) Reset() {
	*x = CartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartResponse) ProtoMessage() {}

func (x *CartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartResponse.ProtoReflect.Descriptor instead.
func (*CartResponse) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_cart_proto_rawDescGZIP(), []int{4}
}

func (x *CartResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CartResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CartResponse) GetDetail() *Cart {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *CartResponse) GetList() []*Cart {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_shopping_admin_cart_proto protoreflect.FileDescriptor

var file_protos_shopping_admin_cart_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x0b, 0x0a,
	0x09, 0x43, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x31, 0x0a, 0x07, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa2, 0x01,
	0x0a, 0x11, 0x43, 0x61, 0x72, 0x74, 0x50, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x31, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x20, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x32, 0xde, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x64, 0x73, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x64, 0x73, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x50, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_shopping_admin_cart_proto_rawDescOnce sync.Once
	file_protos_shopping_admin_cart_proto_rawDescData = file_protos_shopping_admin_cart_proto_rawDesc
)

func file_protos_shopping_admin_cart_proto_rawDescGZIP() []byte {
	file_protos_shopping_admin_cart_proto_rawDescOnce.Do(func() {
		file_protos_shopping_admin_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_shopping_admin_cart_proto_rawDescData)
	})
	return file_protos_shopping_admin_cart_proto_rawDescData
}

var file_protos_shopping_admin_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_shopping_admin_cart_proto_goTypes = []interface{}{
	(*Cart)(nil),              // 0: shopping_admin.Cart
	(*CartParam)(nil),         // 1: shopping_admin.CartParam
	(*CartIds)(nil),           // 2: shopping_admin.CartIds
	(*CartPageAuthQuery)(nil), // 3: shopping_admin.CartPageAuthQuery
	(*CartResponse)(nil),      // 4: shopping_admin.CartResponse
	(*global.Auth)(nil),       // 5: global.Auth
}
var file_protos_shopping_admin_cart_proto_depIdxs = []int32{
	1, // 0: shopping_admin.CartPageAuthQuery.params:type_name -> shopping_admin.CartParam
	5, // 1: shopping_admin.CartPageAuthQuery.auth:type_name -> global.Auth
	0, // 2: shopping_admin.CartResponse.detail:type_name -> shopping_admin.Cart
	0, // 3: shopping_admin.CartResponse.list:type_name -> shopping_admin.Cart
	0, // 4: shopping_admin.CartService.Create:input_type -> shopping_admin.Cart
	0, // 5: shopping_admin.CartService.Update:input_type -> shopping_admin.Cart
	2, // 6: shopping_admin.CartService.Delete:input_type -> shopping_admin.CartIds
	2, // 7: shopping_admin.CartService.FindById:input_type -> shopping_admin.CartIds
	3, // 8: shopping_admin.CartService.FindPageList:input_type -> shopping_admin.CartPageAuthQuery
	4, // 9: shopping_admin.CartService.Create:output_type -> shopping_admin.CartResponse
	4, // 10: shopping_admin.CartService.Update:output_type -> shopping_admin.CartResponse
	4, // 11: shopping_admin.CartService.Delete:output_type -> shopping_admin.CartResponse
	4, // 12: shopping_admin.CartService.FindById:output_type -> shopping_admin.CartResponse
	4, // 13: shopping_admin.CartService.FindPageList:output_type -> shopping_admin.CartResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_shopping_admin_cart_proto_init() }
func file_protos_shopping_admin_cart_proto_init() {
	if File_protos_shopping_admin_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_shopping_admin_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_shopping_admin_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_shopping_admin_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartIds); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_shopping_admin_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartPageAuthQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_shopping_admin_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_shopping_admin_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_shopping_admin_cart_proto_goTypes,
		DependencyIndexes: file_protos_shopping_admin_cart_proto_depIdxs,
		MessageInfos:      file_protos_shopping_admin_cart_proto_msgTypes,
	}.Build()
	File_protos_shopping_admin_cart_proto = out.File
	file_protos_shopping_admin_cart_proto_rawDesc = nil
	file_protos_shopping_admin_cart_proto_goTypes = nil
	file_protos_shopping_admin_cart_proto_depIdxs = nil
}
