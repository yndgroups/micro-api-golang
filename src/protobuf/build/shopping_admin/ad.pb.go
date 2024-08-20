// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/shopping_admin/ad.proto

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

type Ad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId string `protobuf:"bytes,1,opt,name=adId,proto3" json:"adId,omitempty" gorm:"primary_key;<-:create" validate:"required"` // @gotags:  gorm:"primary_key;<-:create" validate:"required"
	// 广告类型(1:全网推广广告，2:商家自定义广告)
	AdType int64 `protobuf:"varint,2,opt,name=adType,proto3" json:"adType,omitempty" validate:"required"` // @gotags:  validate:"required"
	// 广告信息json数组：[{"title":"商品标题","url":"广告图地址","goodId":"跳转的商品id"}]
	AdInfo []*AdInfo `protobuf:"bytes,3,rep,name=adInfo,proto3" json:"adInfo,omitempty"`
	// 呈现类型(1:轮播图广告，2:单图片广告)
	ViewType int64 `protobuf:"varint,4,opt,name=viewType,proto3" json:"viewType,omitempty" validate:"required"` // @gotags:  validate:"required"
	// 店铺id
	StoreId string `protobuf:"bytes,5,opt,name=storeId,proto3" json:"storeId,omitempty" validate:"required"` // @gotags:  validate:"required"
	// 广告位置
	Position int64 `protobuf:"varint,6,opt,name=position,proto3" json:"position,omitempty" validate:"required"` // @gotags:  validate:"required"
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

func (x *Ad) Reset() {
	*x = Ad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_ad_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ad) ProtoMessage() {}

func (x *Ad) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_ad_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ad.ProtoReflect.Descriptor instead.
func (*Ad) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_ad_proto_rawDescGZIP(), []int{0}
}

func (x *Ad) GetAdId() string {
	if x != nil {
		return x.AdId
	}
	return ""
}

func (x *Ad) GetAdType() int64 {
	if x != nil {
		return x.AdType
	}
	return 0
}

func (x *Ad) GetAdInfo() []*AdInfo {
	if x != nil {
		return x.AdInfo
	}
	return nil
}

func (x *Ad) GetViewType() int64 {
	if x != nil {
		return x.ViewType
	}
	return 0
}

func (x *Ad) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Ad) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Ad) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *Ad) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *Ad) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Ad) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Ad) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type AdDetailVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId string `protobuf:"bytes,1,opt,name=adId,proto3" json:"adId,omitempty" gorm:"primary_key;<-:create" validate:"required"` // @gotags:  gorm:"primary_key;<-:create" validate:"required"
	// 广告类型(1:全网推广广告，2:商家自定义广告)
	AdType int64 `protobuf:"varint,2,opt,name=adType,proto3" json:"adType,omitempty" validate:"required"` // @gotags:  validate:"required"
	// 广告信息json数组：[{"title":"商品标题","url":"广告图地址","goodId":"跳转的商品id"}]
	AdInfo []*AdInfo `protobuf:"bytes,3,rep,name=adInfo,proto3" json:"adInfo,omitempty"`
	// 呈现类型(1:轮播图广告，2:单图片广告)
	ViewType int64 `protobuf:"varint,4,opt,name=viewType,proto3" json:"viewType,omitempty" validate:"required"` // @gotags:  validate:"required"
	// 店铺id
	StoreId string `protobuf:"bytes,5,opt,name=storeId,proto3" json:"storeId,omitempty" validate:"required"` // @gotags:  validate:"required"
	// 广告位置
	Position string `protobuf:"bytes,6,opt,name=position,proto3" json:"position,omitempty" gorm:"-"` // @gotags: gorm:"-"
	// 广告类型名称
	AdTypeName string `protobuf:"bytes,7,opt,name=adTypeName,proto3" json:"adTypeName,omitempty" gorm:"-"` // @gotags: gorm:"-"
	// 呈现类型名称
	ViewTypeName string `protobuf:"bytes,8,opt,name=viewTypeName,proto3" json:"viewTypeName,omitempty" gorm:"-"` // @gotags: gorm:"-"
	// 广告位置名称
	PositionName string `protobuf:"bytes,9,opt,name=positionName,proto3" json:"positionName,omitempty" gorm:"-"` // @gotags: gorm:"-"
	// 店铺名称
	StoreName string `protobuf:"bytes,10,opt,name=storeName,proto3" json:"storeName,omitempty" gorm:"-"` // @gotags: gorm:"-"
	// 商品名称
	ProductName string `protobuf:"bytes,11,opt,name=productName,proto3" json:"productName,omitempty" gorm:"-"` // @gotags: gorm:"-"
	// 创建时间
	CreateTime string `protobuf:"bytes,12,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"-"` // @gotags: gorm:"-"
}

func (x *AdDetailVO) Reset() {
	*x = AdDetailVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_ad_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdDetailVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdDetailVO) ProtoMessage() {}

func (x *AdDetailVO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_ad_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdDetailVO.ProtoReflect.Descriptor instead.
func (*AdDetailVO) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_ad_proto_rawDescGZIP(), []int{1}
}

func (x *AdDetailVO) GetAdId() string {
	if x != nil {
		return x.AdId
	}
	return ""
}

func (x *AdDetailVO) GetAdType() int64 {
	if x != nil {
		return x.AdType
	}
	return 0
}

func (x *AdDetailVO) GetAdInfo() []*AdInfo {
	if x != nil {
		return x.AdInfo
	}
	return nil
}

func (x *AdDetailVO) GetViewType() int64 {
	if x != nil {
		return x.ViewType
	}
	return 0
}

func (x *AdDetailVO) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *AdDetailVO) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *AdDetailVO) GetAdTypeName() string {
	if x != nil {
		return x.AdTypeName
	}
	return ""
}

func (x *AdDetailVO) GetViewTypeName() string {
	if x != nil {
		return x.ViewTypeName
	}
	return ""
}

func (x *AdDetailVO) GetPositionName() string {
	if x != nil {
		return x.PositionName
	}
	return ""
}

func (x *AdDetailVO) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *AdDetailVO) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *AdDetailVO) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type AdInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 广告标题
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// 广告图片地址
	ImageUrl string `protobuf:"bytes,2,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	// 跳转地址
	TargetUrl string `protobuf:"bytes,3,opt,name=targetUrl,proto3" json:"targetUrl,omitempty"`
}

func (x *AdInfo) Reset() {
	*x = AdInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_ad_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdInfo) ProtoMessage() {}

func (x *AdInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_ad_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdInfo.ProtoReflect.Descriptor instead.
func (*AdInfo) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_ad_proto_rawDescGZIP(), []int{2}
}

func (x *AdInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AdInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *AdInfo) GetTargetUrl() string {
	if x != nil {
		return x.TargetUrl
	}
	return ""
}

type AdParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdParam) Reset() {
	*x = AdParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_ad_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdParam) ProtoMessage() {}

func (x *AdParam) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_ad_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdParam.ProtoReflect.Descriptor instead.
func (*AdParam) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_ad_proto_rawDescGZIP(), []int{3}
}

type AdIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	UserId string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AdIds) Reset() {
	*x = AdIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_ad_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdIds) ProtoMessage() {}

func (x *AdIds) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_ad_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdIds.ProtoReflect.Descriptor instead.
func (*AdIds) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_ad_proto_rawDescGZIP(), []int{4}
}

func (x *AdIds) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *AdIds) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// 分页查询加权参数
type AdPageAuthQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页索引
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	// 分页长度
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 查询参数
	Params *AdParam `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	// 用户登录授权信息
	Auth *global.Auth `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty" swaggerignore:"true"` // @gotags: swaggerignore:"true"`
}

func (x *AdPageAuthQuery) Reset() {
	*x = AdPageAuthQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_ad_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdPageAuthQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdPageAuthQuery) ProtoMessage() {}

func (x *AdPageAuthQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_ad_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdPageAuthQuery.ProtoReflect.Descriptor instead.
func (*AdPageAuthQuery) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_ad_proto_rawDescGZIP(), []int{5}
}

func (x *AdPageAuthQuery) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *AdPageAuthQuery) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AdPageAuthQuery) GetParams() *AdParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *AdPageAuthQuery) GetAuth() *global.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// 返回结果
type AdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`  // 获取数据数量
	Msg    string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`       // 错误信息
	Detail *Ad    `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
	List   []*Ad  `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`     // 列表
}

func (x *AdResponse) Reset() {
	*x = AdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_admin_ad_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdResponse) ProtoMessage() {}

func (x *AdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_admin_ad_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdResponse.ProtoReflect.Descriptor instead.
func (*AdResponse) Descriptor() ([]byte, []int) {
	return file_protos_shopping_admin_ad_proto_rawDescGZIP(), []int{6}
}

func (x *AdResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AdResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AdResponse) GetDetail() *Ad {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *AdResponse) GetList() []*Ad {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_shopping_admin_ad_proto protoreflect.FileDescriptor

var file_protos_shopping_admin_ad_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x41, 0x64, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbc, 0x02, 0x0a, 0x02, 0x41, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69, 0x65, 0x77, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x76, 0x69, 0x65, 0x77, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xf6, 0x02, 0x0a, 0x0a, 0x41, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x56, 0x4f, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x61, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x64, 0x2e,
	0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x76, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x76, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x76, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x06, 0x41, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55,
	0x72, 0x6c, 0x22, 0x09, 0x0a, 0x07, 0x41, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x2f, 0x0a,
	0x05, 0x41, 0x64, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x92,
	0x01, 0x0a, 0x0f, 0x41, 0x64, 0x50, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41,
	0x64, 0x2e, 0x41, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x20, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x22, 0x70, 0x0a, 0x0a, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x06, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x64, 0x2e, 0x41,
	0x64, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xd0, 0x01, 0x0a, 0x09, 0x41, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x06, 0x2e,
	0x41, 0x64, 0x2e, 0x41, 0x64, 0x1a, 0x0e, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x06, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64, 0x1a, 0x0e, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x09, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64, 0x49, 0x64, 0x73, 0x1a, 0x0e, 0x2e, 0x41,
	0x64, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x09, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64,
	0x49, 0x64, 0x73, 0x1a, 0x0e, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64, 0x50, 0x61, 0x67, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0e, 0x2e, 0x41, 0x64, 0x2e, 0x41, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_shopping_admin_ad_proto_rawDescOnce sync.Once
	file_protos_shopping_admin_ad_proto_rawDescData = file_protos_shopping_admin_ad_proto_rawDesc
)

func file_protos_shopping_admin_ad_proto_rawDescGZIP() []byte {
	file_protos_shopping_admin_ad_proto_rawDescOnce.Do(func() {
		file_protos_shopping_admin_ad_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_shopping_admin_ad_proto_rawDescData)
	})
	return file_protos_shopping_admin_ad_proto_rawDescData
}

var file_protos_shopping_admin_ad_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_shopping_admin_ad_proto_goTypes = []interface{}{
	(*Ad)(nil),              // 0: Ad.Ad
	(*AdDetailVO)(nil),      // 1: Ad.AdDetailVO
	(*AdInfo)(nil),          // 2: Ad.AdInfo
	(*AdParam)(nil),         // 3: Ad.AdParam
	(*AdIds)(nil),           // 4: Ad.AdIds
	(*AdPageAuthQuery)(nil), // 5: Ad.AdPageAuthQuery
	(*AdResponse)(nil),      // 6: Ad.AdResponse
	(*global.Auth)(nil),     // 7: global.Auth
}
var file_protos_shopping_admin_ad_proto_depIdxs = []int32{
	2,  // 0: Ad.Ad.adInfo:type_name -> Ad.AdInfo
	2,  // 1: Ad.AdDetailVO.adInfo:type_name -> Ad.AdInfo
	3,  // 2: Ad.AdPageAuthQuery.params:type_name -> Ad.AdParam
	7,  // 3: Ad.AdPageAuthQuery.auth:type_name -> global.Auth
	0,  // 4: Ad.AdResponse.detail:type_name -> Ad.Ad
	0,  // 5: Ad.AdResponse.list:type_name -> Ad.Ad
	0,  // 6: Ad.AdService.Create:input_type -> Ad.Ad
	0,  // 7: Ad.AdService.Update:input_type -> Ad.Ad
	4,  // 8: Ad.AdService.Delete:input_type -> Ad.AdIds
	4,  // 9: Ad.AdService.FindById:input_type -> Ad.AdIds
	5,  // 10: Ad.AdService.FindPageList:input_type -> Ad.AdPageAuthQuery
	6,  // 11: Ad.AdService.Create:output_type -> Ad.AdResponse
	6,  // 12: Ad.AdService.Update:output_type -> Ad.AdResponse
	6,  // 13: Ad.AdService.Delete:output_type -> Ad.AdResponse
	6,  // 14: Ad.AdService.FindById:output_type -> Ad.AdResponse
	6,  // 15: Ad.AdService.FindPageList:output_type -> Ad.AdResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_protos_shopping_admin_ad_proto_init() }
func file_protos_shopping_admin_ad_proto_init() {
	if File_protos_shopping_admin_ad_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_shopping_admin_ad_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ad); i {
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
		file_protos_shopping_admin_ad_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdDetailVO); i {
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
		file_protos_shopping_admin_ad_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdInfo); i {
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
		file_protos_shopping_admin_ad_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdParam); i {
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
		file_protos_shopping_admin_ad_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdIds); i {
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
		file_protos_shopping_admin_ad_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdPageAuthQuery); i {
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
		file_protos_shopping_admin_ad_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdResponse); i {
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
			RawDescriptor: file_protos_shopping_admin_ad_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_shopping_admin_ad_proto_goTypes,
		DependencyIndexes: file_protos_shopping_admin_ad_proto_depIdxs,
		MessageInfos:      file_protos_shopping_admin_ad_proto_msgTypes,
	}.Build()
	File_protos_shopping_admin_ad_proto = out.File
	file_protos_shopping_admin_ad_proto_rawDesc = nil
	file_protos_shopping_admin_ad_proto_goTypes = nil
	file_protos_shopping_admin_ad_proto_depIdxs = nil
}