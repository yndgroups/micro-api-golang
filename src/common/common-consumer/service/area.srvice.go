package service

import (
	"context"
	"core/app"
	"core/config"
	"core/db"
	"core/logger"
	"core/model"
	"core/pager"
	"core/redis"
	"core/resp"
	"core/validate"
	"protobuf/build/common"

	"github.com/gin-gonic/gin"
)

var Area = &area{}

type area struct{}

// Create 地区新增信息
func (a *area) Create(ctx *gin.Context) any {
	area := common.Area{}
	if err := ctx.ShouldBindJSON(&area); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&area); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	area.AreaId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	area.CreateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewAreaServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &area); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count)
	}
}

// Update 更新地区信息
func (a *area) Update(ctx *gin.Context) any {
	area := common.Area{}
	if err := ctx.ShouldBindJSON(&area); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&area); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	area.UpdateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewAreaServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &area); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

// Delete 删除地区信息
func (a *area) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}

	if result, getResultError := common.NewAreaServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.AreaIds{
		AreaIds: ids,
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Delete(result.Count)
	}
}

// FindById 查询地区详情
func (a *area) FindById(ctx *gin.Context) any {

	if result, getResultError := common.NewAreaServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.AreaIds{
		AreaIds: []string{ctx.Param("id")},
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, result.Detail)
	}
}

// FindPageList 查询地区列表
func (a *area) FindPageList(ctx *gin.Context) any {
	request := common.AreaPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}

	if result, getResultError := common.NewAreaServiceClient(app.Nacos.CommonGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

// FindByAreaCode 根据地区编码查询地区详情
func (a *area) FindByAreaCode(ctx *gin.Context) any {
	var detail common.Area
	first := db.DB.First(&detail, common.Area{AreaCode: ctx.Param("areaCode")})
	if first.RowsAffected > 0 {
		return resp.Success.WithData(&detail)
	} else {
		return resp.Fail.WithMsg("不存在该areaCode相关信息！")
	}
}

// FindListByParentCode 根据地区父编码查询地区列表
func (a *area) FindListByParentCode(ctx *gin.Context) any {
	if parentCode, bl := ctx.GetQuery("parentCode"); !bl {
		return resp.Fail.WithMsg("请正确传入parentCode参数！")
	} else {
		areaList := make([]common.AreaListVo, 0)
		first := db.DB.Find(&areaList, common.Area{ParentCode: parentCode})
		if first.RowsAffected > 0 {
			return resp.Success.WithData(areaList)
		} else {
			return resp.Fail.WithMsg("不存在该areaCode相关信息！")
		}
	}
}

// FindTree 查询省市区三级数据树状结构
func (a *area) FindTree(ctx *gin.Context) any {
	if result, getResultError := common.NewAreaServiceClient(app.Nacos.CommonGrpcProvider()).FindTree(context.Background(), &common.AreaTags{AreaTag: 3}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址：%s,服务生产者【%v】发生错误,错误信息:%v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, a.treeRecursion(result.TreeList, "000000000000"))
	}
}

func (a *area) treeRecursion(list []*common.AreaTree, parent string) []*common.AreaTree {
	// 递归出传入的父级编码数据及非父级编码数据
	parentList := make([]*common.AreaTree, 0)
	childList := make([]*common.AreaTree, 0)
	for _, v := range list {
		if v.ParentCode == parent {
			parentList = append(parentList, v)
		} else {
			childList = append(childList, v)
		}
	}
	// 将非父级别数据进行分组
	childMapList := make(map[string][]*common.AreaTree)
	for _, v := range childList {
		t := childMapList[v.ParentCode]
		t = append(t, v)
		childMapList[v.ParentCode] = t
	}
	// 最总结果
	for i, v := range parentList {
		child := childMapList[v.AreaCode]
		if len(child) > 0 {
			t := a.treeRecursion(childList, v.AreaCode)
			parentList[i].Children = t
		} else {
			parentList[i].Children = child
		}
	}
	return parentList
}
