// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/shopping_admin/product_spec.proto

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

type ProductSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商品规格id
	SpecId string `protobuf:"bytes,1,opt,name=specId,proto3" json:"specId,omitempty" gorm:"primary_key;<-:create"` // @gotags: gorm:"primary_key;<-:create"
	// 规格名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"<-" validate:"required"` // @gotags: gorm:"<-" validate:"required"
	// 绑定分类
	ProductCategoryId string `protobuf:"bytes,3,opt,name=productCategoryId,proto3" json:"productCategoryId,omitempty" gorm:"<-" validate:"required"` // @gotags: gorm:"<-" validate:"required"
	// 可选项
	Options string `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty" gorm:"<-" validate:"required"` // @gotags: gorm:"<-" validate:"required"
	// 排序
	SortBy int64 `protobuf:"varint,5,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	// 删除状态
	DelStatus bool `protobuf:"varint,6,opt,name=delStatus,proto3" json:"delStatus,omitempty" gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"` // @gotags: gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"
	// 创建人
	CreateBy string `protobuf:"bytes,7,opt,name=createBy,proto3" json:"createBy,omitempty" gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 修改人
	UpdateBy string `protobuf:"bytes,8,opt,name=updateBy,proto3" json:"updateBy,omitempty" gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 创建时间
	CreateTime string `protobuf:"bytes,9,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
	// 修改时间
	UpdateTime string `protobuf:"bytes,10,opt,name=updateTime,proto3" json:"updateTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
}

func (x *ProductSpec) Reset() {
	*x = ProductSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSpec) ProtoMessage() {}

func (x *ProductSpec) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSpec.ProtoReflect.Descriptor instead.
func (*ProductSpec) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_product_spec_proto_rawDescGZIP(), []int{0}
}

func (x *ProductSpec) GetSpecId() string {
	if x != nil {
		return x.SpecId
	}
	return ""
}

func (x *ProductSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductSpec) GetProductCategoryId() string {
	if x != nil {
		return x.ProductCategoryId
	}
	return ""
}

func (x *ProductSpec) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

func (x *ProductSpec) GetSortBy() int64 {
	if x != nil {
		return x.SortBy
	}
	return 0
}

func (x *ProductSpec) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *ProductSpec) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *ProductSpec) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *ProductSpec) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ProductSpec) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type ProductSpecListVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商品规格id
	SpecId string `protobuf:"bytes,1,opt,name=specId,proto3" json:"specId,omitempty" gorm:"primary_key;<-:create"` // @gotags: gorm:"primary_key;<-:create"
	// 规格名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"<-" validate:"required"` // @gotags: gorm:"<-" validate:"required"
	// 绑定分类
	ProductCategoryId string `protobuf:"bytes,3,opt,name=productCategoryId,proto3" json:"productCategoryId,omitempty" gorm:"<-" validate:"required"` // @gotags: gorm:"<-" validate:"required"
	// 可选项
	Options string `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty" gorm:"<-" validate:"required"` // @gotags: gorm:"<-" validate:"required"
	// 排序
	SortBy int64 `protobuf:"varint,5,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
}

func (x *ProductSpecListVO) Reset() {
	*x = ProductSpecListVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSpecListVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSpecListVO) ProtoMessage() {}

func (x *ProductSpecListVO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSpecListVO.ProtoReflect.Descriptor instead.
func (*ProductSpecListVO) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_product_spec_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSpecListVO) GetSpecId() string {
	if x != nil {
		return x.SpecId
	}
	return ""
}

func (x *ProductSpecListVO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductSpecListVO) GetProductCategoryId() string {
	if x != nil {
		return x.ProductCategoryId
	}
	return ""
}

func (x *ProductSpecListVO) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

func (x *ProductSpecListVO) GetSortBy() int64 {
	if x != nil {
		return x.SortBy
	}
	return 0
}

type ProductSpecDetailVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商品规格id
	SpecId string `protobuf:"bytes,1,opt,name=specId,proto3" json:"specId,omitempty" gorm:"primary_key;<-:create"` // @gotags: gorm:"primary_key;<-:create"
	// 规格名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"<-" validate:"required"` // @gotags: gorm:"<-" validate:"required"
	// 绑定分类
	ProductCategoryId string `protobuf:"bytes,3,opt,name=productCategoryId,proto3" json:"productCategoryId,omitempty" gorm:"<-" validate:"required"` // @gotags: gorm:"<-" validate:"required"
	// 可选项
	Options string `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty" gorm:"<-" validate:"required"` // @gotags: gorm:"<-" validate:"required"
	// 排序
	SortBy int64 `protobuf:"varint,5,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
}

func (x *ProductSpecDetailVO) Reset() {
	*x = ProductSpecDetailVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSpecDetailVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSpecDetailVO) ProtoMessage() {}

func (x *ProductSpecDetailVO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSpecDetailVO.ProtoReflect.Descriptor instead.
func (*ProductSpecDetailVO) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_product_spec_proto_rawDescGZIP(), []int{2}
}

func (x *ProductSpecDetailVO) GetSpecId() string {
	if x != nil {
		return x.SpecId
	}
	return ""
}

func (x *ProductSpecDetailVO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductSpecDetailVO) GetProductCategoryId() string {
	if x != nil {
		return x.ProductCategoryId
	}
	return ""
}

func (x *ProductSpecDetailVO) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

func (x *ProductSpecDetailVO) GetSortBy() int64 {
	if x != nil {
		return x.SortBy
	}
	return 0
}

// ProductSpecParam 查询参数
type ProductSpecParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProductSpecParam) Reset() {
	*x = ProductSpecParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSpecParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSpecParam) ProtoMessage() {}

func (x *ProductSpecParam) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSpecParam.ProtoReflect.Descriptor instead.
func (*ProductSpecParam) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_product_spec_proto_rawDescGZIP(), []int{3}
}

// ProductSpecIds 规格id参数
type ProductSpecIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id集合
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	// 用户id
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	// 分类id
	ProductCategoryId string `protobuf:"bytes,3,opt,name=productCategoryId,proto3" json:"productCategoryId,omitempty"`
}

func (x *ProductSpecIds) Reset() {
	*x = ProductSpecIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSpecIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSpecIds) ProtoMessage() {}

func (x *ProductSpecIds) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSpecIds.ProtoReflect.Descriptor instead.
func (*ProductSpecIds) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_product_spec_proto_rawDescGZIP(), []int{4}
}

func (x *ProductSpecIds) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ProductSpecIds) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ProductSpecIds) GetProductCategoryId() string {
	if x != nil {
		return x.ProductCategoryId
	}
	return ""
}

// 分页查询加权参数
type ProductSpecPageAuthQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页索引
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	// 分页长度
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 查询参数
	Params *ProductSpecParam `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	// 用户登录授权信息
	Auth *global.Auth `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty" swaggerignore:"true"` // @gotags: swaggerignore:"true"`
}

func (x *ProductSpecPageAuthQuery) Reset() {
	*x = ProductSpecPageAuthQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSpecPageAuthQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSpecPageAuthQuery) ProtoMessage() {}

func (x *ProductSpecPageAuthQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSpecPageAuthQuery.ProtoReflect.Descriptor instead.
func (*ProductSpecPageAuthQuery) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_product_spec_proto_rawDescGZIP(), []int{5}
}

func (x *ProductSpecPageAuthQuery) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *ProductSpecPageAuthQuery) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ProductSpecPageAuthQuery) GetParams() *ProductSpecParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *ProductSpecPageAuthQuery) GetAuth() *global.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// 返回结果
type ProductSpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64          `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`  // 获取数据数量
	Msg    string         `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`       // 错误信息
	Detail *ProductSpec   `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
	List   []*ProductSpec `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`     // 列表
}

func (x *ProductSpecResponse) Reset() {
	*x = ProductSpecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSpecResponse) ProtoMessage() {}

func (x *ProductSpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_product_spec_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSpecResponse.ProtoReflect.Descriptor instead.
func (*ProductSpecResponse) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_product_spec_proto_rawDescGZIP(), []int{6}
}

func (x *ProductSpecResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ProductSpecResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ProductSpecResponse) GetDetail() *ProductSpec {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *ProductSpecResponse) GetList() []*ProductSpec {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_shopping_admin_product_spec_proto protoreflect.FileDescriptor

var file_protos_shopping_admin_product_spec_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xaf, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x65, 0x63,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6f,
	0x72, 0x74, 0x42, 0x79, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x56, 0x4f, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x70,
	0x65, 0x63, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x68, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x49, 0x64, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x50, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x20, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xc7, 0x03,
	0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x49, 0x64, 0x73, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x49, 0x64, 0x73, 0x1a, 0x1d, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c,
	0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x50, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x56, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x49, 0x64, 0x73, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_shopping_admin_product_spec_proto_rawDescOnce sync.Once
	file_protos_shopping_admin_product_spec_proto_rawDescData = file_protos_shopping_admin_product_spec_proto_rawDesc
)

func file_protos_shopping_admin_product_spec_proto_rawDescGZIP() []byte {
	file_protos_shopping_admin_product_spec_proto_rawDescOnce.Do(func() {
		file_protos_shopping_admin_product_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_shopping_admin_product_spec_proto_rawDescData)
	})
	return file_protos_shopping_admin_product_spec_proto_rawDescData
}

var file_protos_shopping_admin_product_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_shopping_admin_product_spec_proto_goTypes = []interface{}{
	(*ProductSpec)(nil),              // 0: template.ProductSpec
	(*ProductSpecListVO)(nil),        // 1: template.ProductSpecListVO
	(*ProductSpecDetailVO)(nil),      // 2: template.ProductSpecDetailVO
	(*ProductSpecParam)(nil),         // 3: template.ProductSpecParam
	(*ProductSpecIds)(nil),           // 4: template.ProductSpecIds
	(*ProductSpecPageAuthQuery)(nil), // 5: template.ProductSpecPageAuthQuery
	(*ProductSpecResponse)(nil),      // 6: template.ProductSpecResponse
	(*global.Auth)(nil),              // 7: global.Auth
}
var file_protos_shopping_admin_product_spec_proto_depIdxs = []int32{
	3,  // 0: template.ProductSpecPageAuthQuery.params:type_name -> template.ProductSpecParam
	7,  // 1: template.ProductSpecPageAuthQuery.auth:type_name -> global.Auth
	0,  // 2: template.ProductSpecResponse.detail:type_name -> template.ProductSpec
	0,  // 3: template.ProductSpecResponse.list:type_name -> template.ProductSpec
	0,  // 4: template.ProductSpecService.Create:input_type -> template.ProductSpec
	0,  // 5: template.ProductSpecService.Update:input_type -> template.ProductSpec
	4,  // 6: template.ProductSpecService.Delete:input_type -> template.ProductSpecIds
	4,  // 7: template.ProductSpecService.FindById:input_type -> template.ProductSpecIds
	5,  // 8: template.ProductSpecService.FindPageList:input_type -> template.ProductSpecPageAuthQuery
	4,  // 9: template.ProductSpecService.FindListByProductCategoryId:input_type -> template.ProductSpecIds
	6,  // 10: template.ProductSpecService.Create:output_type -> template.ProductSpecResponse
	6,  // 11: template.ProductSpecService.Update:output_type -> template.ProductSpecResponse
	6,  // 12: template.ProductSpecService.Delete:output_type -> template.ProductSpecResponse
	6,  // 13: template.ProductSpecService.FindById:output_type -> template.ProductSpecResponse
	6,  // 14: template.ProductSpecService.FindPageList:output_type -> template.ProductSpecResponse
	6,  // 15: template.ProductSpecService.FindListByProductCategoryId:output_type -> template.ProductSpecResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_protos_shopping_admin_product_spec_proto_init() }
func file_protos_shopping_admin_product_spec_proto_init() {
	if File_protos_shopping_admin_product_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_shopping_admin_product_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSpec); i {
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
		file_protos_shopping_admin_product_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSpecListVO); i {
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
		file_protos_shopping_admin_product_spec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSpecDetailVO); i {
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
		file_protos_shopping_admin_product_spec_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSpecParam); i {
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
		file_protos_shopping_admin_product_spec_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSpecIds); i {
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
		file_protos_shopping_admin_product_spec_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSpecPageAuthQuery); i {
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
		file_protos_shopping_admin_product_spec_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSpecResponse); i {
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
			RawDescriptor: file_protos_shopping_admin_product_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_shopping_admin_product_spec_proto_goTypes,
		DependencyIndexes: file_protos_shopping_admin_product_spec_proto_depIdxs,
		MessageInfos:      file_protos_shopping_admin_product_spec_proto_msgTypes,
	}.Build()
	File_protos_shopping_admin_product_spec_proto = out.File
	file_protos_shopping_admin_product_spec_proto_rawDesc = nil
	file_protos_shopping_admin_product_spec_proto_goTypes = nil
	file_protos_shopping_admin_product_spec_proto_depIdxs = nil
}