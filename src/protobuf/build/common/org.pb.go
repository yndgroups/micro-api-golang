// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/common/org.proto

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

type Org struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组织id
	OrgId string `protobuf:"bytes,1,opt,name=orgId,proto3" json:"orgId,omitempty" gorm:"primary_key;<-:create"` // @gotags: gorm:"primary_key;<-:create"
	// 创建类型：1:后台管理员添加 2:自己注册
	CreateType int64 `protobuf:"varint,2,opt,name=createType,proto3" json:"createType,omitempty"` // @gotags gorm:"<-"
	// 企业名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 企业logo
	Logo string `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty"`
	// 统一社会信用代码
	CreditCode string `protobuf:"bytes,5,opt,name=creditCode,proto3" json:"creditCode,omitempty"`
	// 注册地址
	RegisterAddress string `protobuf:"bytes,6,opt,name=registerAddress,proto3" json:"registerAddress,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 企业注册号
	RegisterNo string `protobuf:"bytes,7,opt,name=registerNo,proto3" json:"registerNo,omitempty"`
	// 企业营业执照地址
	LicensePhotoUrl string `protobuf:"bytes,8,opt,name=licensePhotoUrl,proto3" json:"licensePhotoUrl,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 法人
	LegalPerson string `protobuf:"bytes,9,opt,name=legalPerson,proto3" json:"legalPerson,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 法人介绍地址
	LegalPersonUrl string `protobuf:"bytes,10,opt,name=legalPersonUrl,proto3" json:"legalPersonUrl,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 法人身份证正面地址
	LegalPersonCardPurl string `protobuf:"bytes,11,opt,name=legalPersonCardPurl,proto3" json:"legalPersonCardPurl,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 法人身份证反面地址
	LegalPersonCardBurl string `protobuf:"bytes,12,opt,name=legalPersonCardBurl,proto3" json:"legalPersonCardBurl,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 联系人
	ContactPerson string `protobuf:"bytes,13,opt,name=contactPerson,proto3" json:"contactPerson,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 联系电话
	ContactNumber string `protobuf:"bytes,14,opt,name=contactNumber,proto3" json:"contactNumber,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 联系邮箱
	ContactEmail string `protobuf:"bytes,15,opt,name=contactEmail,proto3" json:"contactEmail,omitempty"` // @gotags gorm:"<-" validate:"required"
	// 公司规模
	Scale string `protobuf:"bytes,16,opt,name=scale,proto3" json:"scale,omitempty"` // @gotags gorm:"<-"
	// 人员规模最小人数
	ScalePersonMin int64 `protobuf:"varint,17,opt,name=scalePersonMin,proto3" json:"scalePersonMin,omitempty"` // @gotags gorm:"<-"
	// 人工规模最大人数，-1 为以上
	ScalePersonMax int64 `protobuf:"varint,18,opt,name=scalePersonMax,proto3" json:"scalePersonMax,omitempty"` // @gotags gorm:"<-"
	// 企业性质 多个使用-分割
	Nature string `protobuf:"bytes,19,opt,name=nature,proto3" json:"nature,omitempty"` // @gotags gorm:"<-"
	// 企业浏览量
	Views uint64 `protobuf:"varint,20,opt,name=views,proto3" json:"views,omitempty"` // @gotags gorm:"->"
	// 行业类别id，多个使用-分割
	IndustryIds string `protobuf:"bytes,21,opt,name=industryIds,proto3" json:"industryIds,omitempty"` // @gotags gorm:"<-"
	// 公司福利 使用-分割
	Benefits string `protobuf:"bytes,22,opt,name=benefits,proto3" json:"benefits,omitempty"` // @gotags gorm:"<-"
	// 公司特点 使用-分割
	Characteristic string `protobuf:"bytes,23,opt,name=characteristic,proto3" json:"characteristic,omitempty"` // @gotags gorm:"<-"
	// 描述
	Description string `protobuf:"bytes,29,opt,name=description,proto3" json:"description,omitempty"` // @gotags gorm:"<-"
	// 岗位数量
	JobNum int64 `protobuf:"varint,30,opt,name=jobNum,proto3" json:"jobNum,omitempty"` // @gotags gorm:"<-"
	// 联系邮箱
	Email string `protobuf:"bytes,31,opt,name=email,proto3" json:"email,omitempty"` // @gotags gorm:"<-"
	// 省代码
	AreaCode string `protobuf:"bytes,32,opt,name=areaCode,proto3" json:"areaCode,omitempty"` // @gotags gorm:"<-"
	// 地质详情 比如多少栋楼多少号
	AddrDetail string `protobuf:"bytes,33,opt,name=addrDetail,proto3" json:"addrDetail,omitempty"` // @gotags gorm:"<-"
	// 公司经度
	Lng string `protobuf:"bytes,34,opt,name=lng,proto3" json:"lng,omitempty"` // @gotags gorm:"<-"
	// 公司纬度
	Lat string `protobuf:"bytes,35,opt,name=lat,proto3" json:"lat,omitempty"` // @gotags gorm:"<-"
	// 企业状态 1:待审核 2:审核失败 3:审核成功 4:再次提交
	AuditStatus int64 `protobuf:"varint,36,opt,name=auditStatus,proto3" json:"auditStatus,omitempty"` // @gotags gorm:"<-:update"
	// 审核消息
	AuditDescription string `protobuf:"bytes,37,opt,name=auditDescription,proto3" json:"auditDescription,omitempty"` // @gotags gorm:"<-"
	// 是否为推荐企业或组织 1是 2不是
	IsRecommend int64 `protobuf:"varint,38,opt,name=isRecommend,proto3" json:"isRecommend,omitempty"` // @gotags gorm:"<-"
	// 推荐企业显示排序
	RecommendSort int64 `protobuf:"varint,39,opt,name=recommendSort,proto3" json:"recommendSort,omitempty"` // @gotags gorm:"<-"
	// 是否为行业巨头或组织 1是 2不是
	IsIndustryGiant int64 `protobuf:"varint,40,opt,name=isIndustryGiant,proto3" json:"isIndustryGiant,omitempty"` // @gotags gorm:"<-"
	// 行业巨头显示排序
	IndustrySort int64 `protobuf:"varint,41,opt,name=industrySort,proto3" json:"industrySort,omitempty"` // @gotags gorm:"<-"
	// 是否为热门企业 1 是 0不是
	IsHot int64 `protobuf:"varint,42,opt,name=isHot,proto3" json:"isHot,omitempty"` // @gotags gorm:"<-"
	// 热门企业或组织显示排序
	HotSort int64 `protobuf:"varint,43,opt,name=hotSort,proto3" json:"hotSort,omitempty"` // @gotags gorm:"<-"
	// 删除状态
	DelStatus bool `protobuf:"varint,44,opt,name=delStatus,proto3" json:"delStatus,omitempty" gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"` // @gotags: gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"
	// 创建人
	CreateBy string `protobuf:"bytes,45,opt,name=createBy,proto3" json:"createBy,omitempty" gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 修改人
	UpdateBy string `protobuf:"bytes,46,opt,name=updateBy,proto3" json:"updateBy,omitempty" gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 创建时间
	CreateTime string `protobuf:"bytes,47,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
	// 修改时间
	UpdateTime string `protobuf:"bytes,48,opt,name=updateTime,proto3" json:"updateTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
}

func (x *Org) Reset() {
	*x = Org{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_org_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Org) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Org) ProtoMessage() {}

func (x *Org) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_org_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Org.ProtoReflect.Descriptor instead.
func (*Org) Descriptor() ([]byte, []int) {
	return file_protos_common_org_proto_rawDescGZIP(), []int{0}
}

func (x *Org) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Org) GetCreateType() int64 {
	if x != nil {
		return x.CreateType
	}
	return 0
}

func (x *Org) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Org) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Org) GetCreditCode() string {
	if x != nil {
		return x.CreditCode
	}
	return ""
}

func (x *Org) GetRegisterAddress() string {
	if x != nil {
		return x.RegisterAddress
	}
	return ""
}

func (x *Org) GetRegisterNo() string {
	if x != nil {
		return x.RegisterNo
	}
	return ""
}

func (x *Org) GetLicensePhotoUrl() string {
	if x != nil {
		return x.LicensePhotoUrl
	}
	return ""
}

func (x *Org) GetLegalPerson() string {
	if x != nil {
		return x.LegalPerson
	}
	return ""
}

func (x *Org) GetLegalPersonUrl() string {
	if x != nil {
		return x.LegalPersonUrl
	}
	return ""
}

func (x *Org) GetLegalPersonCardPurl() string {
	if x != nil {
		return x.LegalPersonCardPurl
	}
	return ""
}

func (x *Org) GetLegalPersonCardBurl() string {
	if x != nil {
		return x.LegalPersonCardBurl
	}
	return ""
}

func (x *Org) GetContactPerson() string {
	if x != nil {
		return x.ContactPerson
	}
	return ""
}

func (x *Org) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

func (x *Org) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *Org) GetScale() string {
	if x != nil {
		return x.Scale
	}
	return ""
}

func (x *Org) GetScalePersonMin() int64 {
	if x != nil {
		return x.ScalePersonMin
	}
	return 0
}

func (x *Org) GetScalePersonMax() int64 {
	if x != nil {
		return x.ScalePersonMax
	}
	return 0
}

func (x *Org) GetNature() string {
	if x != nil {
		return x.Nature
	}
	return ""
}

func (x *Org) GetViews() uint64 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *Org) GetIndustryIds() string {
	if x != nil {
		return x.IndustryIds
	}
	return ""
}

func (x *Org) GetBenefits() string {
	if x != nil {
		return x.Benefits
	}
	return ""
}

func (x *Org) GetCharacteristic() string {
	if x != nil {
		return x.Characteristic
	}
	return ""
}

func (x *Org) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Org) GetJobNum() int64 {
	if x != nil {
		return x.JobNum
	}
	return 0
}

func (x *Org) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Org) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *Org) GetAddrDetail() string {
	if x != nil {
		return x.AddrDetail
	}
	return ""
}

func (x *Org) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *Org) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Org) GetAuditStatus() int64 {
	if x != nil {
		return x.AuditStatus
	}
	return 0
}

func (x *Org) GetAuditDescription() string {
	if x != nil {
		return x.AuditDescription
	}
	return ""
}

func (x *Org) GetIsRecommend() int64 {
	if x != nil {
		return x.IsRecommend
	}
	return 0
}

func (x *Org) GetRecommendSort() int64 {
	if x != nil {
		return x.RecommendSort
	}
	return 0
}

func (x *Org) GetIsIndustryGiant() int64 {
	if x != nil {
		return x.IsIndustryGiant
	}
	return 0
}

func (x *Org) GetIndustrySort() int64 {
	if x != nil {
		return x.IndustrySort
	}
	return 0
}

func (x *Org) GetIsHot() int64 {
	if x != nil {
		return x.IsHot
	}
	return 0
}

func (x *Org) GetHotSort() int64 {
	if x != nil {
		return x.HotSort
	}
	return 0
}

func (x *Org) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *Org) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *Org) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Org) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Org) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type OrgParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组织或企业名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 法人
	LegalPerson string `protobuf:"bytes,2,opt,name=legalPerson,proto3" json:"legalPerson,omitempty"`
	// 联系人
	ContactPerson string `protobuf:"bytes,3,opt,name=contactPerson,proto3" json:"contactPerson,omitempty"`
}

func (x *OrgParam) Reset() {
	*x = OrgParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_org_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgParam) ProtoMessage() {}

func (x *OrgParam) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_org_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgParam.ProtoReflect.Descriptor instead.
func (*OrgParam) Descriptor() ([]byte, []int) {
	return file_protos_common_org_proto_rawDescGZIP(), []int{1}
}

func (x *OrgParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrgParam) GetLegalPerson() string {
	if x != nil {
		return x.LegalPerson
	}
	return ""
}

func (x *OrgParam) GetContactPerson() string {
	if x != nil {
		return x.ContactPerson
	}
	return ""
}

type OrgIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgIds []string `protobuf:"bytes,1,rep,name=OrgIds,proto3" json:"OrgIds,omitempty"`
	UserId string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *OrgIds) Reset() {
	*x = OrgIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_org_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgIds) ProtoMessage() {}

func (x *OrgIds) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_org_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgIds.ProtoReflect.Descriptor instead.
func (*OrgIds) Descriptor() ([]byte, []int) {
	return file_protos_common_org_proto_rawDescGZIP(), []int{2}
}

func (x *OrgIds) GetOrgIds() []string {
	if x != nil {
		return x.OrgIds
	}
	return nil
}

func (x *OrgIds) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// 分页查询加权参数
type OrgPageAuthQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页索引
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	// 分页长度
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 查询参数
	Params *OrgParam `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	// 用户登录授权信息
	Auth *global.Auth `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty" swaggerignore:"true"` // @gotags: swaggerignore:"true"`
}

func (x *OrgPageAuthQuery) Reset() {
	*x = OrgPageAuthQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_org_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgPageAuthQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgPageAuthQuery) ProtoMessage() {}

func (x *OrgPageAuthQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_org_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgPageAuthQuery.ProtoReflect.Descriptor instead.
func (*OrgPageAuthQuery) Descriptor() ([]byte, []int) {
	return file_protos_common_org_proto_rawDescGZIP(), []int{3}
}

func (x *OrgPageAuthQuery) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *OrgPageAuthQuery) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OrgPageAuthQuery) GetParams() *OrgParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *OrgPageAuthQuery) GetAuth() *global.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// 返回结果
type OrgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`  // 获取数据数量
	Msg    string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`       // 错误信息
	Detail *Org   `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
	List   []*Org `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`     // 列表
}

func (x *OrgResponse) Reset() {
	*x = OrgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_org_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgResponse) ProtoMessage() {}

func (x *OrgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_org_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgResponse.ProtoReflect.Descriptor instead.
func (*OrgResponse) Descriptor() ([]byte, []int) {
	return file_protos_common_org_proto_rawDescGZIP(), []int{4}
}

func (x *OrgResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *OrgResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OrgResponse) GetDetail() *Org {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *OrgResponse) GetList() []*Org {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_common_org_proto protoreflect.FileDescriptor

var file_protos_common_org_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x0a,
	0x0a, 0x03, 0x4f, 0x72, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x6f, 0x67, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x28, 0x0a,
	0x0f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x65, 0x67, 0x61, 0x6c,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x65,
	0x67, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x65, 0x67,
	0x61, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x55, 0x72,
	0x6c, 0x12, 0x30, 0x0a, 0x13, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x43, 0x61, 0x72, 0x64, 0x50, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x6c, 0x65, 0x67, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x50,
	0x75, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x13, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x42, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x43, 0x61, 0x72,
	0x64, 0x42, 0x75, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x4d, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x64,
	0x75, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x4e, 0x75, 0x6d, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6a, 0x6f, 0x62, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x64, 0x64, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x64, 0x64, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x6e, 0x67, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x61, 0x74, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x24,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x75, 0x64, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x18, 0x26, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x69, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x53, 0x6f, 0x72, 0x74,
	0x18, 0x27, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x79, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x69, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x53, 0x6f, 0x72, 0x74, 0x18,
	0x29, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x53,
	0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x18, 0x2a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x74,
	0x53, 0x6f, 0x72, 0x74, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x53,
	0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x2c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x2d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x30, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x08, 0x4f, 0x72, 0x67,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x65, 0x67,
	0x61, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6c, 0x65, 0x67, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x22, 0x38, 0x0a, 0x06, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x72, 0x67, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x72, 0x67,
	0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x10,
	0x4f, 0x72, 0x67, 0x50, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x7b, 0x0a, 0x0b, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x23, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x32, 0x83, 0x02, 0x0a, 0x0a, 0x4f, 0x72, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x67, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4f, 0x72, 0x67, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f,
	0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72,
	0x67, 0x49, 0x64, 0x73, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f,
	0x72, 0x67, 0x49, 0x64, 0x73, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f,
	0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x46, 0x69,
	0x6e, 0x64, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x67, 0x50, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_common_org_proto_rawDescOnce sync.Once
	file_protos_common_org_proto_rawDescData = file_protos_common_org_proto_rawDesc
)

func file_protos_common_org_proto_rawDescGZIP() []byte {
	file_protos_common_org_proto_rawDescOnce.Do(func() {
		file_protos_common_org_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_org_proto_rawDescData)
	})
	return file_protos_common_org_proto_rawDescData
}

var file_protos_common_org_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_common_org_proto_goTypes = []interface{}{
	(*Org)(nil),              // 0: common.Org
	(*OrgParam)(nil),         // 1: common.OrgParam
	(*OrgIds)(nil),           // 2: common.OrgIds
	(*OrgPageAuthQuery)(nil), // 3: common.OrgPageAuthQuery
	(*OrgResponse)(nil),      // 4: common.OrgResponse
	(*global.Auth)(nil),      // 5: global.Auth
}
var file_protos_common_org_proto_depIdxs = []int32{
	1, // 0: common.OrgPageAuthQuery.params:type_name -> common.OrgParam
	5, // 1: common.OrgPageAuthQuery.auth:type_name -> global.Auth
	0, // 2: common.OrgResponse.detail:type_name -> common.Org
	0, // 3: common.OrgResponse.list:type_name -> common.Org
	0, // 4: common.OrgService.Create:input_type -> common.Org
	0, // 5: common.OrgService.Update:input_type -> common.Org
	2, // 6: common.OrgService.Delete:input_type -> common.OrgIds
	2, // 7: common.OrgService.FindById:input_type -> common.OrgIds
	3, // 8: common.OrgService.FindPageList:input_type -> common.OrgPageAuthQuery
	4, // 9: common.OrgService.Create:output_type -> common.OrgResponse
	4, // 10: common.OrgService.Update:output_type -> common.OrgResponse
	4, // 11: common.OrgService.Delete:output_type -> common.OrgResponse
	4, // 12: common.OrgService.FindById:output_type -> common.OrgResponse
	4, // 13: common.OrgService.FindPageList:output_type -> common.OrgResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_common_org_proto_init() }
func file_protos_common_org_proto_init() {
	if File_protos_common_org_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_common_org_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Org); i {
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
		file_protos_common_org_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgParam); i {
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
		file_protos_common_org_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgIds); i {
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
		file_protos_common_org_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgPageAuthQuery); i {
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
		file_protos_common_org_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgResponse); i {
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
			RawDescriptor: file_protos_common_org_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_common_org_proto_goTypes,
		DependencyIndexes: file_protos_common_org_proto_depIdxs,
		MessageInfos:      file_protos_common_org_proto_msgTypes,
	}.Build()
	File_protos_common_org_proto = out.File
	file_protos_common_org_proto_rawDesc = nil
	file_protos_common_org_proto_goTypes = nil
	file_protos_common_org_proto_depIdxs = nil
}
