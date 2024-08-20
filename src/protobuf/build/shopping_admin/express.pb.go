// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/shopping_admin/express.proto

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

type Express struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 快递公司id
	ExpressId string `protobuf:"bytes,1,opt,name=expressId,proto3" json:"expressId,omitempty" gorm:"primary_key;<-:create"` // @gotags: gorm:"primary_key;<-:create"
	// 快递公司全称
	FullName string `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 简称
	Abbreviation string `protobuf:"bytes,3,opt,name=abbreviation,proto3" json:"abbreviation,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 合伙人id
	PartnerId string `protobuf:"bytes,4,opt,name=partnerId,proto3" json:"partnerId,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 合伙人密钥
	PartnerKey string `protobuf:"bytes,5,opt,name=partnerKey,proto3" json:"partnerKey,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 是否需要取件网店
	Net int64 `protobuf:"varint,6,opt,name=Net,proto3" json:"Net,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 账号
	Account string `protobuf:"bytes,7,opt,name=account,proto3" json:"account,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 密钥
	SecretCode string `protobuf:"bytes,8,opt,name=secretCode,proto3" json:"secretCode,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 网点名称
	NetName string `protobuf:"bytes,9,opt,name=netName,proto3" json:"netName,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 是否显示(0:不显示 1:显示)
	Display int64 `protobuf:"varint,10,opt,name=display,proto3" json:"display,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 是否可用(0:不可用 1:可用)
	Enabled int64 `protobuf:"varint,11,opt,name=enabled,proto3" json:"enabled,omitempty" gorm:"<-"` // @gotags: gorm:"<-"
	// 排序
	SortBy int64 `protobuf:"varint,12,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	// 删除状态
	DelStatus bool `protobuf:"varint,13,opt,name=delStatus,proto3" json:"delStatus,omitempty" gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"` // @gotags: gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"
	// 创建人
	CreateBy string `protobuf:"bytes,14,opt,name=createBy,proto3" json:"createBy,omitempty" gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 修改人
	UpdateBy string `protobuf:"bytes,15,opt,name=updateBy,proto3" json:"updateBy,omitempty" gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 创建时间
	CreateTime string `protobuf:"bytes,16,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
	// 修改时间
	UpdateTime string `protobuf:"bytes,17,opt,name=updateTime,proto3" json:"updateTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
}

func (x *Express) Reset() {
	*x = Express{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_express_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Express) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Express) ProtoMessage() {}

func (x *Express) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_express_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Express.ProtoReflect.Descriptor instead.
func (*Express) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_express_proto_rawDescGZIP(), []int{0}
}

func (x *Express) GetExpressId() string {
	if x != nil {
		return x.ExpressId
	}
	return ""
}

func (x *Express) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Express) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *Express) GetPartnerId() string {
	if x != nil {
		return x.PartnerId
	}
	return ""
}

func (x *Express) GetPartnerKey() string {
	if x != nil {
		return x.PartnerKey
	}
	return ""
}

func (x *Express) GetNet() int64 {
	if x != nil {
		return x.Net
	}
	return 0
}

func (x *Express) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Express) GetSecretCode() string {
	if x != nil {
		return x.SecretCode
	}
	return ""
}

func (x *Express) GetNetName() string {
	if x != nil {
		return x.NetName
	}
	return ""
}

func (x *Express) GetDisplay() int64 {
	if x != nil {
		return x.Display
	}
	return 0
}

func (x *Express) GetEnabled() int64 {
	if x != nil {
		return x.Enabled
	}
	return 0
}

func (x *Express) GetSortBy() int64 {
	if x != nil {
		return x.SortBy
	}
	return 0
}

func (x *Express) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *Express) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *Express) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Express) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Express) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type ExpressParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string `protobuf:"bytes,1,opt,name=fullName,proto3" json:"fullName,omitempty"`
}

func (x *ExpressParam) Reset() {
	*x = ExpressParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_express_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressParam) ProtoMessage() {}

func (x *ExpressParam) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_express_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressParam.ProtoReflect.Descriptor instead.
func (*ExpressParam) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_express_proto_rawDescGZIP(), []int{1}
}

func (x *ExpressParam) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

type ExpressListVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExpressListVo) Reset() {
	*x = ExpressListVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_express_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressListVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressListVo) ProtoMessage() {}

func (x *ExpressListVo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_express_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressListVo.ProtoReflect.Descriptor instead.
func (*ExpressListVo) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_express_proto_rawDescGZIP(), []int{2}
}

type ExpressDetailVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExpressDetailVo) Reset() {
	*x = ExpressDetailVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_express_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressDetailVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressDetailVo) ProtoMessage() {}

func (x *ExpressDetailVo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_express_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressDetailVo.ProtoReflect.Descriptor instead.
func (*ExpressDetailVo) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_express_proto_rawDescGZIP(), []int{3}
}

type ExpressIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	UserId string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ExpressIds) Reset() {
	*x = ExpressIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_express_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressIds) ProtoMessage() {}

func (x *ExpressIds) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_express_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressIds.ProtoReflect.Descriptor instead.
func (*ExpressIds) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_express_proto_rawDescGZIP(), []int{4}
}

func (x *ExpressIds) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ExpressIds) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// 分页查询加权参数
type ExpressPageAuthQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页索引
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	// 分页长度
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 查询参数
	Params *ExpressParam `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	// 用户登录授权信息
	Auth *global.Auth `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty" swaggerignore:"true"` // @gotags: swaggerignore:"true"`
}

func (x *ExpressPageAuthQuery) Reset() {
	*x = ExpressPageAuthQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_express_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressPageAuthQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressPageAuthQuery) ProtoMessage() {}

func (x *ExpressPageAuthQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_express_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressPageAuthQuery.ProtoReflect.Descriptor instead.
func (*ExpressPageAuthQuery) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_express_proto_rawDescGZIP(), []int{5}
}

func (x *ExpressPageAuthQuery) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *ExpressPageAuthQuery) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ExpressPageAuthQuery) GetParams() *ExpressParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *ExpressPageAuthQuery) GetAuth() *global.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// 返回结果
type ExpressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`  // 获取数据数量
	Msg    string     `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`       // 错误信息
	Detail *Express   `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
	List   []*Express `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`     // 列表
}

func (x *ExpressResponse) Reset() {
	*x = ExpressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_express_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressResponse) ProtoMessage() {}

func (x *ExpressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_express_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressResponse.ProtoReflect.Descriptor instead.
func (*ExpressResponse) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_express_proto_rawDescGZIP(), []int{6}
}

func (x *ExpressResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ExpressResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ExpressResponse) GetDetail() *Express {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *ExpressResponse) GetList() []*Express {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_shopping_admin_express_proto protoreflect.FileDescriptor

var file_protos_shopping_admin_express_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a,
	0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x03, 0x0a, 0x07,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x4e, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x56, 0x6f, 0x22, 0x34, 0x0a, 0x0a, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0xa2, 0x01, 0x0a, 0x14, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50, 0x61, 0x67,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x8b, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x29, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x32, 0xc3, 0x02, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x1a, 0x19, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x19, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x73, 0x1a, 0x19, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x73, 0x1a, 0x19, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x50, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x19, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_shopping_admin_express_proto_rawDescOnce sync.Once
	file_protos_shopping_admin_express_proto_rawDescData = file_protos_shopping_admin_express_proto_rawDesc
)

func file_protos_shopping_admin_express_proto_rawDescGZIP() []byte {
	file_protos_shopping_admin_express_proto_rawDescOnce.Do(func() {
		file_protos_shopping_admin_express_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_shopping_admin_express_proto_rawDescData)
	})
	return file_protos_shopping_admin_express_proto_rawDescData
}

var file_protos_shopping_admin_express_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_shopping_admin_express_proto_goTypes = []interface{}{
	(*Express)(nil),              // 0: template.Express
	(*ExpressParam)(nil),         // 1: template.ExpressParam
	(*ExpressListVo)(nil),        // 2: template.ExpressListVo
	(*ExpressDetailVo)(nil),      // 3: template.ExpressDetailVo
	(*ExpressIds)(nil),           // 4: template.ExpressIds
	(*ExpressPageAuthQuery)(nil), // 5: template.ExpressPageAuthQuery
	(*ExpressResponse)(nil),      // 6: template.ExpressResponse
	(*global.Auth)(nil),          // 7: global.Auth
}
var file_protos_shopping_admin_express_proto_depIdxs = []int32{
	1, // 0: template.ExpressPageAuthQuery.params:type_name -> template.ExpressParam
	7, // 1: template.ExpressPageAuthQuery.auth:type_name -> global.Auth
	0, // 2: template.ExpressResponse.detail:type_name -> template.Express
	0, // 3: template.ExpressResponse.list:type_name -> template.Express
	0, // 4: template.ExpressService.Create:input_type -> template.Express
	0, // 5: template.ExpressService.Update:input_type -> template.Express
	4, // 6: template.ExpressService.Delete:input_type -> template.ExpressIds
	4, // 7: template.ExpressService.FindById:input_type -> template.ExpressIds
	5, // 8: template.ExpressService.FindPageList:input_type -> template.ExpressPageAuthQuery
	6, // 9: template.ExpressService.Create:output_type -> template.ExpressResponse
	6, // 10: template.ExpressService.Update:output_type -> template.ExpressResponse
	6, // 11: template.ExpressService.Delete:output_type -> template.ExpressResponse
	6, // 12: template.ExpressService.FindById:output_type -> template.ExpressResponse
	6, // 13: template.ExpressService.FindPageList:output_type -> template.ExpressResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_shopping_admin_express_proto_init() }
func file_protos_shopping_admin_express_proto_init() {
	if File_protos_shopping_admin_express_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_shopping_admin_express_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Express); i {
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
		file_protos_shopping_admin_express_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressParam); i {
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
		file_protos_shopping_admin_express_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressListVo); i {
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
		file_protos_shopping_admin_express_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressDetailVo); i {
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
		file_protos_shopping_admin_express_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressIds); i {
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
		file_protos_shopping_admin_express_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressPageAuthQuery); i {
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
		file_protos_shopping_admin_express_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressResponse); i {
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
			RawDescriptor: file_protos_shopping_admin_express_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_shopping_admin_express_proto_goTypes,
		DependencyIndexes: file_protos_shopping_admin_express_proto_depIdxs,
		MessageInfos:      file_protos_shopping_admin_express_proto_msgTypes,
	}.Build()
	File_protos_shopping_admin_express_proto = out.File
	file_protos_shopping_admin_express_proto_rawDesc = nil
	file_protos_shopping_admin_express_proto_goTypes = nil
	file_protos_shopping_admin_express_proto_depIdxs = nil
}