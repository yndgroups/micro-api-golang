// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/common/agreement.proto

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

// Agreement 协议
type Agreement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 协议id
	AgreeId string `protobuf:"bytes,1,opt,name=agreeId,proto3" json:"agreeId" gorm:"primary_key;<-:create"` // @gotags: gorm:"primary_key;<-:create" json:"agreeId"
	// 协议类型  1：会员协议,2:代理商协议
	Type uint64 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty" gorm:"<-" valid:"required"` // @gotags: gorm:"<-" valid:"required"
	// 协议名称
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty" gorm:"<-" valid:"required"` // @gotags: gorm:"<-" valid:"required"
	// 协议内容
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty" gorm:"<-" valid:"required"` // @gotags: gorm:"<-" valid:"required"
	// 排序
	SortBy uint64 `protobuf:"varint,5,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	// 1：显示：0：不显示
	Display uint64 `protobuf:"varint,6,opt,name=display,proto3" json:"display,omitempty" gorm:"<-" valid:"required"` // @gotags: gorm:"<-"  valid:"required"
	// 删除状态
	DelStatus bool `protobuf:"varint,7,opt,name=delStatus,proto3" json:"delStatus,omitempty" gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"` // @gotags: gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"
	// 创建人
	CreateBy string `protobuf:"bytes,8,opt,name=createBy,proto3" json:"createBy,omitempty" gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 修改人
	UpdateBy string `protobuf:"bytes,9,opt,name=updateBy,proto3" json:"updateBy,omitempty" gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 创建时间
	CreateTime string `protobuf:"bytes,10,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
	// 修改时间
	UpdateTime string `protobuf:"bytes,11,opt,name=updateTime,proto3" json:"updateTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
}

func (x *Agreement) Reset() {
	*x = Agreement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_agreement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agreement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agreement) ProtoMessage() {}

func (x *Agreement) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_agreement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agreement.ProtoReflect.Descriptor instead.
func (*Agreement) Descriptor() ([]byte, []int) {
	return file_protos_common_agreement_proto_rawDescGZIP(), []int{0}
}

func (x *Agreement) GetAgreeId() string {
	if x != nil {
		return x.AgreeId
	}
	return ""
}

func (x *Agreement) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Agreement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Agreement) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Agreement) GetSortBy() uint64 {
	if x != nil {
		return x.SortBy
	}
	return 0
}

func (x *Agreement) GetDisplay() uint64 {
	if x != nil {
		return x.Display
	}
	return 0
}

func (x *Agreement) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *Agreement) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *Agreement) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Agreement) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Agreement) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

// AgreementDetailVo 协议详情
type AgreementDetailVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 协议id
	AgreeId string `protobuf:"bytes,1,opt,name=agreeId,proto3" json:"agreeId"` // @gotags: json:"agreeId"
	// 协议类型
	Type uint64 `protobuf:"varint,2,opt,name=type,proto3" json:"type"` // @gotags: json:"type"
	// 协议类型名称 1：会员协议 2:代理商协议
	TypeName string `protobuf:"bytes,3,opt,name=typeName,proto3" json:"typeName"` // @gotags: json:"typeName"
	// 协议名称
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title"` //  @gotags: json:"title"
	// 协议内容
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content"` // @gotags: json:"content"
	// 排序
	SortBy uint64 `protobuf:"varint,6,opt,name=sortBy,proto3" json:"sortBy"` // @gotags: json:"sortBy"
	// 显示状态名称
	Display uint64 `protobuf:"varint,7,opt,name=display,proto3" json:"display"` // @gotags: json:"display"
	// 显示状态名称 1：显示：0：不显示
	DisplayName string `protobuf:"bytes,8,opt,name=displayName,proto3" json:"displayName"` // @gotags: json:"displayName"
	// 创建时间
	CreateTime string `protobuf:"bytes,9,opt,name=createTime,proto3" json:"createTime"` // @gotags: json:"createTime"
	// 修改时间
	UpdateTime string `protobuf:"bytes,10,opt,name=updateTime,proto3" json:"updateTime"` // @gotags: json:"updateTime"
}

func (x *AgreementDetailVo) Reset() {
	*x = AgreementDetailVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_agreement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgreementDetailVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgreementDetailVo) ProtoMessage() {}

func (x *AgreementDetailVo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_agreement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgreementDetailVo.ProtoReflect.Descriptor instead.
func (*AgreementDetailVo) Descriptor() ([]byte, []int) {
	return file_protos_common_agreement_proto_rawDescGZIP(), []int{1}
}

func (x *AgreementDetailVo) GetAgreeId() string {
	if x != nil {
		return x.AgreeId
	}
	return ""
}

func (x *AgreementDetailVo) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AgreementDetailVo) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *AgreementDetailVo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AgreementDetailVo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AgreementDetailVo) GetSortBy() uint64 {
	if x != nil {
		return x.SortBy
	}
	return 0
}

func (x *AgreementDetailVo) GetDisplay() uint64 {
	if x != nil {
		return x.Display
	}
	return 0
}

func (x *AgreementDetailVo) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AgreementDetailVo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *AgreementDetailVo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

// AgreementParam
type AgreementParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 协议名称
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// 协议内容
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AgreementParam) Reset() {
	*x = AgreementParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_agreement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgreementParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgreementParam) ProtoMessage() {}

func (x *AgreementParam) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_agreement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgreementParam.ProtoReflect.Descriptor instead.
func (*AgreementParam) Descriptor() ([]byte, []int) {
	return file_protos_common_agreement_proto_rawDescGZIP(), []int{2}
}

func (x *AgreementParam) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AgreementParam) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// 用户id集合
type AgreementIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	AgreeIds []string `protobuf:"bytes,1,rep,name=agreeIds,proto3" json:"agreeIds,omitempty"`
	UserId   string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AgreementIds) Reset() {
	*x = AgreementIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_agreement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgreementIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgreementIds) ProtoMessage() {}

func (x *AgreementIds) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_agreement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgreementIds.ProtoReflect.Descriptor instead.
func (*AgreementIds) Descriptor() ([]byte, []int) {
	return file_protos_common_agreement_proto_rawDescGZIP(), []int{3}
}

func (x *AgreementIds) GetAgreeIds() []string {
	if x != nil {
		return x.AgreeIds
	}
	return nil
}

func (x *AgreementIds) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// UserPageQuery 用户分页查询
type AgreementPageAuthQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页索引
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" default:"1"` // @gotags: `json:"pageIndex" default:"1"`
	// 分页长度
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" default:"10"` // @gotags: `json:"pageSize" default:"10"`
	// 查询参数
	Params *AgreementParam `protobuf:"bytes,3,opt,name=params,proto3" json:"params"` // @gotags: `json:"params"`
	// 用户登录授权信息
	Auth *global.Auth `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty" swaggerignore:"true"` // @gotags: swaggerignore:"true"
}

func (x *AgreementPageAuthQuery) Reset() {
	*x = AgreementPageAuthQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_agreement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgreementPageAuthQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgreementPageAuthQuery) ProtoMessage() {}

func (x *AgreementPageAuthQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_agreement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgreementPageAuthQuery.ProtoReflect.Descriptor instead.
func (*AgreementPageAuthQuery) Descriptor() ([]byte, []int) {
	return file_protos_common_agreement_proto_rawDescGZIP(), []int{4}
}

func (x *AgreementPageAuthQuery) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *AgreementPageAuthQuery) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AgreementPageAuthQuery) GetParams() *AgreementParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *AgreementPageAuthQuery) GetAuth() *global.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// 用户数据
type AgreementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64              `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`  // 获取数据数量
	Msg    string             `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`       // 错误信息
	Detail *AgreementDetailVo `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
	List   []*Agreement       `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`     // 列表
}

func (x *AgreementResponse) Reset() {
	*x = AgreementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_agreement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgreementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgreementResponse) ProtoMessage() {}

func (x *AgreementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_agreement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgreementResponse.ProtoReflect.Descriptor instead.
func (*AgreementResponse) Descriptor() ([]byte, []int) {
	return file_protos_common_agreement_proto_rawDescGZIP(), []int{5}
}

func (x *AgreementResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AgreementResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AgreementResponse) GetDetail() *AgreementDetailVo {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *AgreementResponse) GetList() []*Agreement {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_common_agreement_proto protoreflect.FileDescriptor

var file_protos_common_agreement_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x09, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x67, 0x72, 0x65, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x72, 0x65, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa1, 0x02, 0x0a, 0x11, 0x41, 0x67, 0x72, 0x65,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x56, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x67, 0x72, 0x65, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x67, 0x72, 0x65, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x0e, 0x41,
	0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x42, 0x0a,
	0x0c, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x67, 0x72, 0x65, 0x65, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x67, 0x72, 0x65, 0x65, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0xa4, 0x01, 0x0a, 0x16, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x41, 0x67, 0x72,
	0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x31, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x56,
	0x6f, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x32, 0xc5, 0x02, 0x0a, 0x10, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x67, 0x72, 0x65,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x73, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41,
	0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x73, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x67, 0x72, 0x65,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0c, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x19, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_common_agreement_proto_rawDescOnce sync.Once
	file_protos_common_agreement_proto_rawDescData = file_protos_common_agreement_proto_rawDesc
)

func file_protos_common_agreement_proto_rawDescGZIP() []byte {
	file_protos_common_agreement_proto_rawDescOnce.Do(func() {
		file_protos_common_agreement_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_agreement_proto_rawDescData)
	})
	return file_protos_common_agreement_proto_rawDescData
}

var file_protos_common_agreement_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_common_agreement_proto_goTypes = []interface{}{
	(*Agreement)(nil),              // 0: common.Agreement
	(*AgreementDetailVo)(nil),      // 1: common.AgreementDetailVo
	(*AgreementParam)(nil),         // 2: common.AgreementParam
	(*AgreementIds)(nil),           // 3: common.AgreementIds
	(*AgreementPageAuthQuery)(nil), // 4: common.AgreementPageAuthQuery
	(*AgreementResponse)(nil),      // 5: common.AgreementResponse
	(*global.Auth)(nil),            // 6: global.Auth
}
var file_protos_common_agreement_proto_depIdxs = []int32{
	2, // 0: common.AgreementPageAuthQuery.params:type_name -> common.AgreementParam
	6, // 1: common.AgreementPageAuthQuery.auth:type_name -> global.Auth
	1, // 2: common.AgreementResponse.detail:type_name -> common.AgreementDetailVo
	0, // 3: common.AgreementResponse.list:type_name -> common.Agreement
	0, // 4: common.AgreementService.Create:input_type -> common.Agreement
	0, // 5: common.AgreementService.Update:input_type -> common.Agreement
	3, // 6: common.AgreementService.Delete:input_type -> common.AgreementIds
	3, // 7: common.AgreementService.FindById:input_type -> common.AgreementIds
	4, // 8: common.AgreementService.FindPageList:input_type -> common.AgreementPageAuthQuery
	5, // 9: common.AgreementService.Create:output_type -> common.AgreementResponse
	5, // 10: common.AgreementService.Update:output_type -> common.AgreementResponse
	5, // 11: common.AgreementService.Delete:output_type -> common.AgreementResponse
	5, // 12: common.AgreementService.FindById:output_type -> common.AgreementResponse
	5, // 13: common.AgreementService.FindPageList:output_type -> common.AgreementResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_common_agreement_proto_init() }
func file_protos_common_agreement_proto_init() {
	if File_protos_common_agreement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_common_agreement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agreement); i {
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
		file_protos_common_agreement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgreementDetailVo); i {
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
		file_protos_common_agreement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgreementParam); i {
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
		file_protos_common_agreement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgreementIds); i {
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
		file_protos_common_agreement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgreementPageAuthQuery); i {
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
		file_protos_common_agreement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgreementResponse); i {
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
			RawDescriptor: file_protos_common_agreement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_common_agreement_proto_goTypes,
		DependencyIndexes: file_protos_common_agreement_proto_depIdxs,
		MessageInfos:      file_protos_common_agreement_proto_msgTypes,
	}.Build()
	File_protos_common_agreement_proto = out.File
	file_protos_common_agreement_proto_rawDesc = nil
	file_protos_common_agreement_proto_goTypes = nil
	file_protos_common_agreement_proto_depIdxs = nil
}