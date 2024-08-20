// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/shopping_client/ad.proto

package shopping_client

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdListVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 广告 id
	AdId string `protobuf:"bytes,1,opt,name=adId,proto3" json:"adId,omitempty"`
	// 广告类型(1:全网推广广告, 2:商家自定义广告)
	AdType int64 `protobuf:"varint,2,opt,name=adType,proto3" json:"adType,omitempty"`
	// 广告位置(1:首页, 2:列表页,3详情页)
	Position int64 `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	// 店铺id
	StoreId string `protobuf:"bytes,4,opt,name=storeId,proto3" json:"storeId,omitempty"`
	// 呈现类型(1:轮播图广告, 2:单图片广告)
	ViewType int64 `protobuf:"varint,5,opt,name=viewType,proto3" json:"viewType,omitempty"`
	// 广告信息json数组：[{"title":"商品标题","url":"广告图地址","goodId":"跳转的商品id"}]
	AdInfo string `protobuf:"bytes,6,opt,name=adInfo,proto3" json:"adInfo,omitempty"`
	// 删除状态
	DelStatus bool `protobuf:"varint,7,opt,name=delStatus,proto3" json:"delStatus,omitempty"`
	// 创建人
	CreateBy string `protobuf:"bytes,8,opt,name=createBy,proto3" json:"createBy,omitempty"`
	// 修改人
	UpdateBy string `protobuf:"bytes,9,opt,name=updateBy,proto3" json:"updateBy,omitempty"`
	// 创建时间
	CreateTime string `protobuf:"bytes,10,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// 修改时间
	UpdateTime string `protobuf:"bytes,11,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *AdListVo) Reset() {
	*x = AdListVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_client_ad_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdListVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdListVo) ProtoMessage() {}

func (x *AdListVo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_client_ad_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdListVo.ProtoReflect.Descriptor instead.
func (*AdListVo) Descriptor() ([]byte, []int) {
	return file_protos_shopping_client_ad_proto_rawDescGZIP(), []int{0}
}

func (x *AdListVo) GetAdId() string {
	if x != nil {
		return x.AdId
	}
	return ""
}

func (x *AdListVo) GetAdType() int64 {
	if x != nil {
		return x.AdType
	}
	return 0
}

func (x *AdListVo) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *AdListVo) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *AdListVo) GetViewType() int64 {
	if x != nil {
		return x.ViewType
	}
	return 0
}

func (x *AdListVo) GetAdInfo() string {
	if x != nil {
		return x.AdInfo
	}
	return ""
}

func (x *AdListVo) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *AdListVo) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *AdListVo) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *AdListVo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *AdListVo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

// 广告id参数
type AdIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdIds  []string `protobuf:"bytes,1,rep,name=adIds,proto3" json:"adIds,omitempty"`
	UserId string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AdIds) Reset() {
	*x = AdIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_client_ad_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdIds) ProtoMessage() {}

func (x *AdIds) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_client_ad_proto_msgTypes[1]
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
	return file_protos_shopping_client_ad_proto_rawDescGZIP(), []int{1}
}

func (x *AdIds) GetAdIds() []string {
	if x != nil {
		return x.AdIds
	}
	return nil
}

func (x *AdIds) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// 广告详情参数
type AdDetailVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 广告 id
	AdId string `protobuf:"bytes,1,opt,name=adId,proto3" json:"adId,omitempty"`
	// 广告类型(1:全网推广广告, 2:商家自定义广告)
	AdType int64 `protobuf:"varint,2,opt,name=adType,proto3" json:"adType,omitempty"`
	// 广告位置(1:首页, 2:列表页,3详情页)
	Position int64 `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	// 店铺id
	StoreId string `protobuf:"bytes,4,opt,name=storeId,proto3" json:"storeId,omitempty"`
	// 呈现类型(1:轮播图广告, 2:单图片广告)
	ViewType int64 `protobuf:"varint,5,opt,name=viewType,proto3" json:"viewType,omitempty"`
	// 广告信息json数组：[{"title":"商品标题","url":"广告图地址","goodId":"跳转的商品id"}]
	AdInfo string `protobuf:"bytes,6,opt,name=adInfo,proto3" json:"adInfo,omitempty"`
	// 删除状态
	DelStatus bool `protobuf:"varint,7,opt,name=delStatus,proto3" json:"delStatus,omitempty"`
	// 创建人
	CreateBy string `protobuf:"bytes,8,opt,name=createBy,proto3" json:"createBy,omitempty"`
	// 修改人
	UpdateBy string `protobuf:"bytes,9,opt,name=updateBy,proto3" json:"updateBy,omitempty"`
	// 创建时间
	CreateTime string `protobuf:"bytes,10,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// 修改时间
	UpdateTime string `protobuf:"bytes,11,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *AdDetailVo) Reset() {
	*x = AdDetailVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_client_ad_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdDetailVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdDetailVo) ProtoMessage() {}

func (x *AdDetailVo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_client_ad_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdDetailVo.ProtoReflect.Descriptor instead.
func (*AdDetailVo) Descriptor() ([]byte, []int) {
	return file_protos_shopping_client_ad_proto_rawDescGZIP(), []int{2}
}

func (x *AdDetailVo) GetAdId() string {
	if x != nil {
		return x.AdId
	}
	return ""
}

func (x *AdDetailVo) GetAdType() int64 {
	if x != nil {
		return x.AdType
	}
	return 0
}

func (x *AdDetailVo) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *AdDetailVo) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *AdDetailVo) GetViewType() int64 {
	if x != nil {
		return x.ViewType
	}
	return 0
}

func (x *AdDetailVo) GetAdInfo() string {
	if x != nil {
		return x.AdInfo
	}
	return ""
}

func (x *AdDetailVo) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *AdDetailVo) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *AdDetailVo) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *AdDetailVo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *AdDetailVo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type AddListParma struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 广告类型
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// 广告位置
	Position string `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *AddListParma) Reset() {
	*x = AddListParma{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_client_ad_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddListParma) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddListParma) ProtoMessage() {}

func (x *AddListParma) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_client_ad_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddListParma.ProtoReflect.Descriptor instead.
func (*AddListParma) Descriptor() ([]byte, []int) {
	return file_protos_shopping_client_ad_proto_rawDescGZIP(), []int{3}
}

func (x *AddListParma) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AddListParma) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

// 返回结果
type AdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`  // 获取数据数量
	Msg    string      `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`       // 错误信息
	Detail *AdDetailVo `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
	List   []*AdListVo `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`     // 列表
}

func (x *AdResponse) Reset() {
	*x = AdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shopping_client_ad_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdResponse) ProtoMessage() {}

func (x *AdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shopping_client_ad_proto_msgTypes[4]
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
	return file_protos_shopping_client_ad_proto_rawDescGZIP(), []int{4}
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

func (x *AdResponse) GetDetail() *AdDetailVo {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *AdResponse) GetList() []*AdListVo {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_shopping_client_ad_proto protoreflect.FileDescriptor

var file_protos_shopping_client_ad_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0xb6, 0x02, 0x0a, 0x08, 0x41, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x76, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x05, 0x41,
	0x64, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xb8, 0x02, 0x0a, 0x0a, 0x41, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x56,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x76, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3e, 0x0a,
	0x0c, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x6d, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x98, 0x01,
	0x0a, 0x0a, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x33, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x56,
	0x6f, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x96, 0x01, 0x0a, 0x09, 0x41, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x49, 0x64, 0x73, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x72, 0x6d, 0x61, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x17, 0x5a, 0x15, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_shopping_client_ad_proto_rawDescOnce sync.Once
	file_protos_shopping_client_ad_proto_rawDescData = file_protos_shopping_client_ad_proto_rawDesc
)

func file_protos_shopping_client_ad_proto_rawDescGZIP() []byte {
	file_protos_shopping_client_ad_proto_rawDescOnce.Do(func() {
		file_protos_shopping_client_ad_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_shopping_client_ad_proto_rawDescData)
	})
	return file_protos_shopping_client_ad_proto_rawDescData
}

var file_protos_shopping_client_ad_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_shopping_client_ad_proto_goTypes = []interface{}{
	(*AdListVo)(nil),     // 0: shopping_client.AdListVo
	(*AdIds)(nil),        // 1: shopping_client.AdIds
	(*AdDetailVo)(nil),   // 2: shopping_client.AdDetailVo
	(*AddListParma)(nil), // 3: shopping_client.AddListParma
	(*AdResponse)(nil),   // 4: shopping_client.AdResponse
}
var file_protos_shopping_client_ad_proto_depIdxs = []int32{
	2, // 0: shopping_client.AdResponse.detail:type_name -> shopping_client.AdDetailVo
	0, // 1: shopping_client.AdResponse.list:type_name -> shopping_client.AdListVo
	1, // 2: shopping_client.AdService.FindById:input_type -> shopping_client.AdIds
	3, // 3: shopping_client.AdService.FindAdList:input_type -> shopping_client.AddListParma
	4, // 4: shopping_client.AdService.FindById:output_type -> shopping_client.AdResponse
	4, // 5: shopping_client.AdService.FindAdList:output_type -> shopping_client.AdResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_shopping_client_ad_proto_init() }
func file_protos_shopping_client_ad_proto_init() {
	if File_protos_shopping_client_ad_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_shopping_client_ad_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdListVo); i {
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
		file_protos_shopping_client_ad_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_shopping_client_ad_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdDetailVo); i {
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
		file_protos_shopping_client_ad_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddListParma); i {
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
		file_protos_shopping_client_ad_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_protos_shopping_client_ad_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_shopping_client_ad_proto_goTypes,
		DependencyIndexes: file_protos_shopping_client_ad_proto_depIdxs,
		MessageInfos:      file_protos_shopping_client_ad_proto_msgTypes,
	}.Build()
	File_protos_shopping_client_ad_proto = out.File
	file_protos_shopping_client_ad_proto_rawDesc = nil
	file_protos_shopping_client_ad_proto_goTypes = nil
	file_protos_shopping_client_ad_proto_depIdxs = nil
}
