// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/order/order_detail.proto

package order

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

// 订单详情
type OrderDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 订单详情id
	DetailId string `protobuf:"bytes,1,opt,name=detailId,proto3" json:"detailId,omitempty" gorm:"primary_key;<-:create"` // @gotags: `gorm:"primary_key;<-:create"`
	// 商品id
	ProductId string `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 分类id
	CategoryId string `protobuf:"bytes,3,opt,name=categoryId,proto3" json:"categoryId,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 分类名称
	CategoryName string `protobuf:"bytes,4,opt,name=categoryName,proto3" json:"categoryName,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 商品名称
	ProductName string `protobuf:"bytes,5,opt,name=productName,proto3" json:"productName,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 商品简介
	Introduction string `protobuf:"bytes,6,opt,name=introduction,proto3" json:"introduction,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 成本价
	Price float64 `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 商品价格
	CostPrice float64 `protobuf:"fixed64,8,opt,name=costPrice,proto3" json:"costPrice,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 会员价格
	VipPrice float64 `protobuf:"fixed64,9,opt,name=vipPrice,proto3" json:"vipPrice,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 市场价
	MarketPrice float64 `protobuf:"fixed64,10,opt,name=marketPrice,proto3" json:"marketPrice,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 商品图片
	Image string `protobuf:"bytes,11,opt,name=image,proto3" json:"image,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 单位名
	UnitName string `protobuf:"bytes,12,opt,name=unitName,proto3" json:"unitName,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 运费模板id
	TempId string `protobuf:"bytes,13,opt,name=tempId,proto3" json:"tempId,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 规格(0:单规格 1:多规格)
	SpecType int64 `protobuf:"varint,14,opt,name=specType,proto3" json:"specType,omitempty" gorm:"<-:create"` // @gotags: `gorm:"<-:create"`
	// 是否虚拟商品(0:否 1:是)
	IsVirtual int64 `protobuf:"varint,15,opt,name=isVirtual,proto3" json:"isVirtual,omitempty"`
	// 商户id(0为超级员创建 非0时为商户)
	MerId string `protobuf:"bytes,16,opt,name=merId,proto3" json:"merId,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 商户是否代理 0:不可代理 1:可代理)
	MerUse int64 `protobuf:"varint,17,opt,name=merUse,proto3" json:"merUse,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 商品条码(条形码)
	BarCode string `protobuf:"bytes,18,opt,name=barCode,proto3" json:"barCode,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 是否优惠
	IsBenefit int64 `protobuf:"varint,19,opt,name=isBenefit,proto3" json:"isBenefit,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 是否精品
	IsBest int64 `protobuf:"varint,20,opt,name=isBest,proto3" json:"isBest,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 是否单独分佣
	IsSub int64 `protobuf:"varint,21,opt,name=isSub,proto3" json:"isSub,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 商品二维码地址(用户小程序海报)
	CodePath string `protobuf:"bytes,22,opt,name=codePath,proto3" json:"codePath,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 预售结束后几天内发货
	PresaleDay string `protobuf:"bytes,23,opt,name=presaleDay,proto3" json:"presaleDay,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 物流方式
	Logistics int64 `protobuf:"varint,24,opt,name=logistics,proto3" json:"logistics,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 运费设置
	Freight int64 `protobuf:"varint,25,opt,name=freight,proto3" json:"freight,omitempty" gorm:"-"` // @gotags: `gorm:"-"`
	// 订单id
	OrderId string `protobuf:"bytes,26,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *OrderDetail) Reset() {
	*x = OrderDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_order_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetail) ProtoMessage() {}

func (x *OrderDetail) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_order_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetail.ProtoReflect.Descriptor instead.
func (*OrderDetail) Descriptor() ([]byte, []int) {
	return file_protos_order_order_detail_proto_rawDescGZIP(), []int{0}
}

func (x *OrderDetail) GetDetailId() string {
	if x != nil {
		return x.DetailId
	}
	return ""
}

func (x *OrderDetail) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderDetail) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *OrderDetail) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *OrderDetail) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *OrderDetail) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *OrderDetail) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderDetail) GetCostPrice() float64 {
	if x != nil {
		return x.CostPrice
	}
	return 0
}

func (x *OrderDetail) GetVipPrice() float64 {
	if x != nil {
		return x.VipPrice
	}
	return 0
}

func (x *OrderDetail) GetMarketPrice() float64 {
	if x != nil {
		return x.MarketPrice
	}
	return 0
}

func (x *OrderDetail) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *OrderDetail) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *OrderDetail) GetTempId() string {
	if x != nil {
		return x.TempId
	}
	return ""
}

func (x *OrderDetail) GetSpecType() int64 {
	if x != nil {
		return x.SpecType
	}
	return 0
}

func (x *OrderDetail) GetIsVirtual() int64 {
	if x != nil {
		return x.IsVirtual
	}
	return 0
}

func (x *OrderDetail) GetMerId() string {
	if x != nil {
		return x.MerId
	}
	return ""
}

func (x *OrderDetail) GetMerUse() int64 {
	if x != nil {
		return x.MerUse
	}
	return 0
}

func (x *OrderDetail) GetBarCode() string {
	if x != nil {
		return x.BarCode
	}
	return ""
}

func (x *OrderDetail) GetIsBenefit() int64 {
	if x != nil {
		return x.IsBenefit
	}
	return 0
}

func (x *OrderDetail) GetIsBest() int64 {
	if x != nil {
		return x.IsBest
	}
	return 0
}

func (x *OrderDetail) GetIsSub() int64 {
	if x != nil {
		return x.IsSub
	}
	return 0
}

func (x *OrderDetail) GetCodePath() string {
	if x != nil {
		return x.CodePath
	}
	return ""
}

func (x *OrderDetail) GetPresaleDay() string {
	if x != nil {
		return x.PresaleDay
	}
	return ""
}

func (x *OrderDetail) GetLogistics() int64 {
	if x != nil {
		return x.Logistics
	}
	return 0
}

func (x *OrderDetail) GetFreight() int64 {
	if x != nil {
		return x.Freight
	}
	return 0
}

func (x *OrderDetail) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

var File_protos_order_order_detail_proto protoreflect.FileDescriptor

var file_protos_order_order_detail_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xe9, 0x05, 0x0a, 0x0b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x76, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6d, 0x70, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x65, 0x6d, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x70, 0x65, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69,
	0x73, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x42, 0x65,
	0x73, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x73, 0x42, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x53, 0x75, 0x62, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x69, 0x73, 0x53, 0x75, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x44, 0x61, 0x79,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x44,
	0x61, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x42, 0x0d, 0x5a, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_order_order_detail_proto_rawDescOnce sync.Once
	file_protos_order_order_detail_proto_rawDescData = file_protos_order_order_detail_proto_rawDesc
)

func file_protos_order_order_detail_proto_rawDescGZIP() []byte {
	file_protos_order_order_detail_proto_rawDescOnce.Do(func() {
		file_protos_order_order_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_order_order_detail_proto_rawDescData)
	})
	return file_protos_order_order_detail_proto_rawDescData
}

var file_protos_order_order_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_order_order_detail_proto_goTypes = []interface{}{
	(*OrderDetail)(nil), // 0: order.OrderDetail
}
var file_protos_order_order_detail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_order_order_detail_proto_init() }
func file_protos_order_order_detail_proto_init() {
	if File_protos_order_order_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_order_order_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetail); i {
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
			RawDescriptor: file_protos_order_order_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_order_order_detail_proto_goTypes,
		DependencyIndexes: file_protos_order_order_detail_proto_depIdxs,
		MessageInfos:      file_protos_order_order_detail_proto_msgTypes,
	}.Build()
	File_protos_order_order_detail_proto = out.File
	file_protos_order_order_detail_proto_rawDesc = nil
	file_protos_order_order_detail_proto_goTypes = nil
	file_protos_order_order_detail_proto_depIdxs = nil
}
