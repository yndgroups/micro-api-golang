package service

import (
	"context"
	"core/db"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
)

var UserAddressService = &userAddressService{}

type userAddressService struct {
	common.UserAddressServiceServer
}

// FindList 查询用户下的收货地址
func (u userAddressService) FindList(ctx context.Context, ids *common.UserAddressIds) (*common.UserAddressResponse, error) {
	list := make([]*common.UserAddress, 0)
	sql := ` SELECT * FROM ` + enum.SysUserAddress.TableName
	sql += fmt.Sprintf(` WHERE del_status = 0 And create_by = '%v'`, ids.UserId)
	count := db.DB.Raw(sql).Scan(&list).RowsAffected
	return &common.UserAddressResponse{List: list, Count: count}, nil
}

// Create 创建用户收货地址
func (u userAddressService) Create(ctx context.Context, address *common.UserAddress) (*common.UserAddressResponse, error) {
	if address.IsDefault == 1 {
		db.DB.Model(&common.UserAddress{}).Where("create_by = ?", address.UpdateBy).Update("is_default", "0")
	}
	return &common.UserAddressResponse{Count: db.NewCrud(&address, 1)}, nil
}

// Update 修改用户收货地址
func (u userAddressService) Update(ctx context.Context, address *common.UserAddress) (*common.UserAddressResponse, error) {
	if address.IsDefault == 1 {
		db.DB.Model(&common.UserAddress{}).Where("create_by = ?", address.UpdateBy).Update("is_default", "0")
	}
	return &common.UserAddressResponse{Count: db.NewCrud(&address, 2)}, nil
}

// Delete 删除收货地址
func (u userAddressService) Delete(ctx context.Context, ids *common.UserAddressIds) (*common.UserAddressResponse, error) {
	return &common.UserAddressResponse{Count: db.GrpcBatchDeleteByIds(ids.AddressIds, enum.SysUserAddress.TableName, enum.SysUserAddress.PrimaryKey, ids.UserId)}, nil
}

// FindById 查询详情
func (u userAddressService) FindById(ctx context.Context, ids *common.UserAddressIds) (*common.UserAddressResponse, error) {
	detail := common.UserAddress{}
	affected := db.GetDB().First(&detail, common.UserAddress{AddressId: ids.AddressIds[0]}).RowsAffected
	return &common.UserAddressResponse{Detail: &detail, Count: affected}, nil
}

// FindPageList 查询收货地址分页列表
func (u userAddressService) FindPageList(ctx context.Context, query *common.UserAddressPageAuthQuery) (*common.UserAddressResponse, error) {
	return &common.UserAddressResponse{}, nil
}
