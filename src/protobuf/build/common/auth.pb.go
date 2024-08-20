// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: protos/common/auth.proto

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

// 角色和菜单的关系
type RoleMenu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色id
	RoleId string `protobuf:"bytes,1,opt,name=roleId,proto3" json:"roleId,omitempty" gorm:"primary_key;<-:create"` // @gotags: gorm:"primary_key;<-:create"
	// 菜单id
	MenuId string `protobuf:"bytes,2,opt,name=menuId,proto3" json:"menuId,omitempty"`
	// 删除状态
	DelStatus bool `protobuf:"varint,3,opt,name=delStatus,proto3" json:"delStatus,omitempty" gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"` // @gotags: gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"
	// 创建人
	CreateBy string `protobuf:"bytes,4,opt,name=createBy,proto3" json:"createBy,omitempty" gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 修改人
	UpdateBy string `protobuf:"bytes,5,opt,name=updateBy,proto3" json:"updateBy,omitempty" gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 创建时间
	CreateTime string `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
	// 修改时间
	UpdateTime string `protobuf:"bytes,7,opt,name=updateTime,proto3" json:"updateTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
}

func (x *RoleMenu) Reset() {
	*x = RoleMenu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleMenu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleMenu) ProtoMessage() {}

func (x *RoleMenu) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleMenu.ProtoReflect.Descriptor instead.
func (*RoleMenu) Descriptor() ([]byte, []int) {
	return file_protos_common_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RoleMenu) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RoleMenu) GetMenuId() string {
	if x != nil {
		return x.MenuId
	}
	return ""
}

func (x *RoleMenu) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *RoleMenu) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *RoleMenu) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *RoleMenu) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *RoleMenu) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

// 用户关联角色
type UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色id
	RoleId string `protobuf:"bytes,1,opt,name=roleId,proto3" json:"roleId,omitempty" gorm:"primary_key;<-:create"` // @gotags: gorm:"primary_key;<-:create"
	// 用户id
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	// 用户id
	AppId string `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
	// 删除状态
	DelStatus bool `protobuf:"varint,4,opt,name=delStatus,proto3" json:"delStatus,omitempty" gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"` // @gotags: gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"
	// 创建人
	CreateBy string `protobuf:"bytes,5,opt,name=createBy,proto3" json:"createBy,omitempty" gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 修改人
	UpdateBy string `protobuf:"bytes,6,opt,name=updateBy,proto3" json:"updateBy,omitempty" gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 创建时间
	CreateTime string `protobuf:"bytes,7,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
	// 修改时间
	UpdateTime string `protobuf:"bytes,8,opt,name=updateTime,proto3" json:"updateTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_protos_common_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UserRole) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserRole) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *UserRole) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *UserRole) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *UserRole) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UserRole) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *UserRole) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

// 角色和功能的关系
type RoleFunc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色id
	RoleId string `protobuf:"bytes,1,opt,name=roleId,proto3" json:"roleId,omitempty" gorm:"primary_key;<-:create"` // @gotags: gorm:"primary_key;<-:create"
	// 功能id
	FuncId string `protobuf:"bytes,2,opt,name=funcId,proto3" json:"funcId,omitempty"`
	// 删除状态
	DelStatus bool `protobuf:"varint,3,opt,name=delStatus,proto3" json:"delStatus,omitempty" gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"` // @gotags: gorm:"-" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"
	// 创建人
	CreateBy string `protobuf:"bytes,4,opt,name=createBy,proto3" json:"createBy,omitempty" gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:create" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 修改人
	UpdateBy string `protobuf:"bytes,5,opt,name=updateBy,proto3" json:"updateBy,omitempty" gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"` // @gotags: gorm:"->;<-:update" swaggerignore:"true" minLength:"6" maxLength:"19"
	// 创建时间
	CreateTime string `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
	// 修改时间
	UpdateTime string `protobuf:"bytes,7,opt,name=updateTime,proto3" json:"updateTime,omitempty" gorm:"->" swaggerignore:"true"` // @gotags: gorm:"->" swaggerignore:"true"
}

func (x *RoleFunc) Reset() {
	*x = RoleFunc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFunc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFunc) ProtoMessage() {}

func (x *RoleFunc) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFunc.ProtoReflect.Descriptor instead.
func (*RoleFunc) Descriptor() ([]byte, []int) {
	return file_protos_common_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RoleFunc) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RoleFunc) GetFuncId() string {
	if x != nil {
		return x.FuncId
	}
	return ""
}

func (x *RoleFunc) GetDelStatus() bool {
	if x != nil {
		return x.DelStatus
	}
	return false
}

func (x *RoleFunc) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *RoleFunc) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *RoleFunc) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *RoleFunc) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

// 返回结果
type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 获取数据数量
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// 错误信息
	Msg string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	// 用户id列表
	UserIds []string `protobuf:"bytes,3,rep,name=userIds,proto3" json:"userIds,omitempty"`
	// 角色id列表
	RoleIds []string `protobuf:"bytes,4,rep,name=roleIds,proto3" json:"roleIds,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_protos_common_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AuthResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AuthResponse) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *AuthResponse) GetRoleIds() []string {
	if x != nil {
		return x.RoleIds
	}
	return nil
}

// 色关联菜单列表
type RoleMenuList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*RoleMenu `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *RoleMenuList) Reset() {
	*x = RoleMenuList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleMenuList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleMenuList) ProtoMessage() {}

func (x *RoleMenuList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleMenuList.ProtoReflect.Descriptor instead.
func (*RoleMenuList) Descriptor() ([]byte, []int) {
	return file_protos_common_auth_proto_rawDescGZIP(), []int{4}
}

func (x *RoleMenuList) GetList() []*RoleMenu {
	if x != nil {
		return x.List
	}
	return nil
}

// 角色与用户关联列表
type UserRoleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*UserRole `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserRoleList) Reset() {
	*x = UserRoleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleList) ProtoMessage() {}

func (x *UserRoleList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleList.ProtoReflect.Descriptor instead.
func (*UserRoleList) Descriptor() ([]byte, []int) {
	return file_protos_common_auth_proto_rawDescGZIP(), []int{5}
}

func (x *UserRoleList) GetList() []*UserRole {
	if x != nil {
		return x.List
	}
	return nil
}

// 角色与功能关联列表
type RoleFuncList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*RoleFunc `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *RoleFuncList) Reset() {
	*x = RoleFuncList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFuncList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFuncList) ProtoMessage() {}

func (x *RoleFuncList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFuncList.ProtoReflect.Descriptor instead.
func (*RoleFuncList) Descriptor() ([]byte, []int) {
	return file_protos_common_auth_proto_rawDescGZIP(), []int{6}
}

func (x *RoleFuncList) GetList() []*RoleFunc {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_common_auth_proto protoreflect.FileDescriptor

var file_protos_common_auth_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0,
	0x01, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0xe6, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x08, 0x52,
	0x6f, 0x6c, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x75, 0x6e, 0x63, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x75, 0x6e, 0x63, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6a, 0x0a,
	0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x34, 0x0a, 0x0c, 0x52, 0x6f, 0x6c,
	0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22,
	0x34, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x75, 0x6e,
	0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xf0, 0x01, 0x0a, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x3c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x75,
	0x6e, 0x63, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x46, 0x75, 0x6e, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0c, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e,
	0x5a, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_common_auth_proto_rawDescOnce sync.Once
	file_protos_common_auth_proto_rawDescData = file_protos_common_auth_proto_rawDesc
)

func file_protos_common_auth_proto_rawDescGZIP() []byte {
	file_protos_common_auth_proto_rawDescOnce.Do(func() {
		file_protos_common_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_auth_proto_rawDescData)
	})
	return file_protos_common_auth_proto_rawDescData
}

var file_protos_common_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_common_auth_proto_goTypes = []interface{}{
	(*RoleMenu)(nil),     // 0: common.RoleMenu
	(*UserRole)(nil),     // 1: common.UserRole
	(*RoleFunc)(nil),     // 2: common.RoleFunc
	(*AuthResponse)(nil), // 3: common.AuthResponse
	(*RoleMenuList)(nil), // 4: common.RoleMenuList
	(*UserRoleList)(nil), // 5: common.UserRoleList
	(*RoleFuncList)(nil), // 6: common.RoleFuncList
	(*global.Auth)(nil),  // 7: global.Auth
}
var file_protos_common_auth_proto_depIdxs = []int32{
	0, // 0: common.RoleMenuList.list:type_name -> common.RoleMenu
	1, // 1: common.UserRoleList.list:type_name -> common.UserRole
	2, // 2: common.RoleFuncList.list:type_name -> common.RoleFunc
	4, // 3: common.Auth.CreateRoleMenu:input_type -> common.RoleMenuList
	5, // 4: common.Auth.CreateUserRole:input_type -> common.UserRoleList
	6, // 5: common.Auth.CreateRoleFunc:input_type -> common.RoleFuncList
	7, // 6: common.Auth.GetToken:input_type -> global.Auth
	3, // 7: common.Auth.CreateRoleMenu:output_type -> common.AuthResponse
	3, // 8: common.Auth.CreateUserRole:output_type -> common.AuthResponse
	3, // 9: common.Auth.CreateRoleFunc:output_type -> common.AuthResponse
	3, // 10: common.Auth.GetToken:output_type -> common.AuthResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_common_auth_proto_init() }
func file_protos_common_auth_proto_init() {
	if File_protos_common_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_common_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleMenu); i {
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
		file_protos_common_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRole); i {
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
		file_protos_common_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFunc); i {
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
		file_protos_common_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_protos_common_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleMenuList); i {
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
		file_protos_common_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoleList); i {
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
		file_protos_common_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFuncList); i {
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
			RawDescriptor: file_protos_common_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_common_auth_proto_goTypes,
		DependencyIndexes: file_protos_common_auth_proto_depIdxs,
		MessageInfos:      file_protos_common_auth_proto_msgTypes,
	}.Build()
	File_protos_common_auth_proto = out.File
	file_protos_common_auth_proto_rawDesc = nil
	file_protos_common_auth_proto_goTypes = nil
	file_protos_common_auth_proto_depIdxs = nil
}
