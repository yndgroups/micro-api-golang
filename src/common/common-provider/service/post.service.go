package service

import (
	"context"
	"core/config"
	"core/db"
	"core/exception"
	"fmt"
	"protobuf/build/common"
	"protobuf/enum"
	"strings"
)

var PostService = &postService{}

type postService struct {
	common.PostServiceServer
}

// CreatePost 添加岗位
func (p postService) Create(ctx context.Context, post *common.Post) (*common.PostResponse, error) {
	return &common.PostResponse{Count: db.NewCrud(post, 1)}, nil
}

// UpdatePost 修改岗位
func (p postService) Update(ctx context.Context, post *common.Post) (*common.PostResponse, error) {
	return &common.PostResponse{Count: db.NewCrud(post, 2)}, nil
}

// DeletePost 删除岗位
func (p postService) Delete(ctx context.Context, ids *common.PostIds) (*common.PostResponse, error) {
	return &common.PostResponse{Count: db.GrpcBatchDeleteByIds(ids.PostIds, enum.SysPost.TableName, enum.SysPost.PrimaryKey, ids.UserId)}, nil
}

// FindById 查询岗位详情
func (p postService) FindById(ctx context.Context, ids *common.PostIds) (*common.PostResponse, error) {
	//TODO implement me
	panic("implement me")
}

// FindPageList 查询岗位分页列表
func (p postService) FindPageList(ctx context.Context, query *common.PostPageAuthQuery) (*common.PostResponse, error) {
	sql := `
SELECT 
	post_id,
	post_code,
	post_name,
	introduce,
	status,
	sort_by,
	del_status,
	create_by,
	update_by,
	create_time,
	update_time
FROM `
	sql += fmt.Sprintf(`%s WHERE del_status = 0`, enum.SysPost.TableName)
	list := make([]*common.Post, 0)
	res := db.DB.Raw(sql).Scan(&list)
	if res.Error != nil {
		return &common.PostResponse{Msg: exception.DbMsg(res.Error.Error(), config.Global.ServerConfig().CommonProvider)}, nil
	}
	return &common.PostResponse{List: list}, nil
}

// FindPostList 插叙岗位列表
func (p postService) FindList(ctx context.Context, query *common.PostPageAuthQuery) (*common.PostResponse, error) {
	if strings.Contains("x", "") {
		return &common.PostResponse{Count: 0, Msg: "fuck"}, nil
	}
	return &common.PostResponse{Count: 0, Msg: "fuck"}, nil
}
