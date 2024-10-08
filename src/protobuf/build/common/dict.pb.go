// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/common/dict.proto

package common

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

// 字典
type Dict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 字典id
	DictId string `protobuf:"bytes,1,opt,name=dictId,proto3" json:"dictId,omitempty" gorm:"primary_key;<-:create"` //  @gotags: gorm:"primary_key;<-:create"
	// 父id(默认为1)
	ParentId string `protobuf:"bytes,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	// 字典名称
	DictName string `protobuf:"bytes,3,opt,name=dictName,proto3" json:"dictName,omitempty"`
	// 字典解释/说明
	DictValue string `protobuf:"bytes,4,opt,name=dictValue,proto3" json:"dictValue,omitempty" validate:"required"` //  @gotags: validate:"required"
	// 字典解释/说明
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// 启用状态
	Enable int64 `protobuf:"varint,6,opt,name=Enable,proto3" json:"Enable,omitempty"`
	// 排序
	SortBy int64 `protobuf:"varint,7,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	// 删除状态
	DelStatus bool `protobuf:"varint,8,opt,name=delStatus,proto3" json:"delStatus,omitempty" gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"` // @gotags: gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"
	// 创建人
	CreateBy string `protobuf:"bytes,9,opt,name=createBy,proto3" json:"createBy,omitempty" gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 修改人
	UpdateBy string `protobuf:"bytes,10,opt,name=updateBy,proto3" json:"updateBy,omitempty" gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 创建时间
	CreateTime string `protobuf:"bytes,11,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
	// 修改时间
	UpdateTime string `protobuf:"bytes,17,opt,name=updateTime,proto3" json:"updateTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
}

func (x *Dict) Reset() {
	*x = Dict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_dict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dict) ProtoMessage() {}

func (x *Dict) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_dict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dict.ProtoReflect.Descriptor instead.
func (*Dict) Descriptor() ([]byte, []int) {
	return file_protos_common_dict_proto_rawDescGZIP(), []int{0}
}

func (x *Dict) GetDictId() string {
	if x != nil {
		return x.DictId
	}
	return ""
}

func (x *Dict) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *Dict) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *Dict) GetDictValue() string {
	if x != nil {
		return x.DictValue
	}
	return ""
}

func (x *Dict) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dict) GetEnable() int64 {
	if x != nil {
		return x.Enable
	}
	return 0
}

func (x *Dict) GetSortBy() int64 {
	if x != nil {
		return x.SortBy
	}
	return 0
}

func (x *Dict) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *Dict) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *Dict) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Dict) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Dict) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type DictParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DictParam) Reset() {
	*x = DictParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_dict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictParam) ProtoMessage() {}

func (x *DictParam) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_dict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictParam.ProtoReflect.Descriptor instead.
func (*DictParam) Descriptor() ([]byte, []int) {
	return file_protos_common_dict_proto_rawDescGZIP(), []int{1}
}

type DictIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 字典id
	DictIds []string `protobuf:"bytes,1,rep,name=dictIds,proto3" json:"dictIds,omitempty"`
	// 用户id
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	// id类型 1为字典id 2为字典parentId
	ParentId string `protobuf:"bytes,3,opt,name=parentId,proto3" json:"parentId,omitempty"`
}

func (x *DictIds) Reset() {
	*x = DictIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_dict_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictIds) ProtoMessage() {}

func (x *DictIds) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_dict_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictIds.ProtoReflect.Descriptor instead.
func (*DictIds) Descriptor() ([]byte, []int) {
	return file_protos_common_dict_proto_rawDescGZIP(), []int{2}
}

func (x *DictIds) GetDictIds() []string {
	if x != nil {
		return x.DictIds
	}
	return nil
}

func (x *DictIds) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DictIds) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

// 分页查询加权参数
type DictPageAuthQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页索引
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	// 分页长度
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 查询参数
	Params *DictParam `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	// 用户登录授权信息
	Auth *global.Auth `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty" swaggerignore:"true"` // @gotags: swaggerignore:"true"`
}

func (x *DictPageAuthQuery) Reset() {
	*x = DictPageAuthQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_dict_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictPageAuthQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictPageAuthQuery) ProtoMessage() {}

func (x *DictPageAuthQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_dict_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictPageAuthQuery.ProtoReflect.Descriptor instead.
func (*DictPageAuthQuery) Descriptor() ([]byte, []int) {
	return file_protos_common_dict_proto_rawDescGZIP(), []int{3}
}

func (x *DictPageAuthQuery) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *DictPageAuthQuery) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DictPageAuthQuery) GetParams() *DictParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *DictPageAuthQuery) GetAuth() *global.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// 返回结果
type DictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`  // 获取数据数量
	Msg    string  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`       // 错误信息
	Detail *Dict   `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
	List   []*Dict `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`     // 列表
}

func (x *DictResponse) Reset() {
	*x = DictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_dict_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictResponse) ProtoMessage() {}

func (x *DictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_dict_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictResponse.ProtoReflect.Descriptor instead.
func (*DictResponse) Descriptor() ([]byte, []int) {
	return file_protos_common_dict_proto_rawDescGZIP(), []int{4}
}

func (x *DictResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DictResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DictResponse) GetDetail() *Dict {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *DictResponse) GetList() []*Dict {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_common_dict_proto protoreflect.FileDescriptor

var file_protos_common_dict_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc,
	0x02, 0x0a, 0x04, 0x44, 0x69, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x63, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x0b, 0x0a,
	0x09, 0x44, 0x69, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x57, 0x0a, 0x07, 0x44, 0x69,
	0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x11, 0x44, 0x69, 0x63, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x63,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x20,
	0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x22, 0x7e, 0x0a, 0x0c, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x32, 0xcb, 0x02, 0x0a, 0x0b, 0x44, 0x69, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x44, 0x69, 0x63, 0x74, 0x49, 0x64, 0x73, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x49, 0x64, 0x73, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x44, 0x69, 0x63, 0x74, 0x49, 0x64, 0x73, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x0c, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x50, 0x61, 0x67, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e,
	0x5a, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_common_dict_proto_rawDescOnce sync.Once
	file_protos_common_dict_proto_rawDescData = file_protos_common_dict_proto_rawDesc
)

func file_protos_common_dict_proto_rawDescGZIP() []byte {
	file_protos_common_dict_proto_rawDescOnce.Do(func() {
		file_protos_common_dict_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_dict_proto_rawDescData)
	})
	return file_protos_common_dict_proto_rawDescData
}

var file_protos_common_dict_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_common_dict_proto_goTypes = []interface{}{
	(*Dict)(nil),              // 0: common.Dict
	(*DictParam)(nil),         // 1: common.DictParam
	(*DictIds)(nil),           // 2: common.DictIds
	(*DictPageAuthQuery)(nil), // 3: common.DictPageAuthQuery
	(*DictResponse)(nil),      // 4: common.DictResponse
	(*global.Auth)(nil),       // 5: global.Auth
}
var file_protos_common_dict_proto_depIdxs = []int32{
	1,  // 0: common.DictPageAuthQuery.params:type_name -> common.DictParam
	5,  // 1: common.DictPageAuthQuery.auth:type_name -> global.Auth
	0,  // 2: common.DictResponse.detail:type_name -> common.Dict
	0,  // 3: common.DictResponse.list:type_name -> common.Dict
	0,  // 4: common.DictService.Create:input_type -> common.Dict
	0,  // 5: common.DictService.Update:input_type -> common.Dict
	2,  // 6: common.DictService.Delete:input_type -> common.DictIds
	2,  // 7: common.DictService.FindById:input_type -> common.DictIds
	2,  // 8: common.DictService.FindListByParentId:input_type -> common.DictIds
	3,  // 9: common.DictService.FindPageList:input_type -> common.DictPageAuthQuery
	4,  // 10: common.DictService.Create:output_type -> common.DictResponse
	4,  // 11: common.DictService.Update:output_type -> common.DictResponse
	4,  // 12: common.DictService.Delete:output_type -> common.DictResponse
	4,  // 13: common.DictService.FindById:output_type -> common.DictResponse
	4,  // 14: common.DictService.FindListByParentId:output_type -> common.DictResponse
	4,  // 15: common.DictService.FindPageList:output_type -> common.DictResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_protos_common_dict_proto_init() }
func file_protos_common_dict_proto_init() {
	if File_protos_common_dict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_common_dict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dict); i {
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
		file_protos_common_dict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictParam); i {
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
		file_protos_common_dict_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictIds); i {
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
		file_protos_common_dict_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictPageAuthQuery); i {
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
		file_protos_common_dict_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictResponse); i {
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
			RawDescriptor: file_protos_common_dict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_common_dict_proto_goTypes,
		DependencyIndexes: file_protos_common_dict_proto_depIdxs,
		MessageInfos:      file_protos_common_dict_proto_msgTypes,
	}.Build()
	File_protos_common_dict_proto = out.File
	file_protos_common_dict_proto_rawDesc = nil
	file_protos_common_dict_proto_goTypes = nil
	file_protos_common_dict_proto_depIdxs = nil
}
