package controller

import (
	"careful/model"
	"careful/pkg/define/params"
	"github.com/gin-gonic/gin"
	"github.com/zenfire-cn/webkit/rest"
	"strconv"
)

var ServerPoint = &serverPoint{}

type serverPoint struct {
}

func (sp *serverPoint) List(ctx *gin.Context) {
	folderIDStr := ctx.Query("folder_id")
	if folderIDStr == "" {
		rest.Error(ctx, "参数 folder_id 为空")
		return
	}

	folderID, err := strconv.ParseUint(folderIDStr, 10, 64)
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	serverPointM := &model.ServerPoint{}
	list, err := serverPointM.List(uint(folderID))
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, list)
}

func (sp *serverPoint) Create(ctx *gin.Context) {
	var req params.ServerPointCreateReq

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	nodeM := &model.Node{ID: req.NodeID}
	exist, err := nodeM.QueryOne("id")
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	if !exist {
		rest.Error(ctx, "网络节点不存在")
		return
	}

	serverM := &model.Server{ID: req.ServerID}
	exist, err = serverM.QueryOne("id")
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	if !exist {
		rest.Error(ctx, "服务不存在")
		return
	}

	serverPointM := &model.ServerPoint{
		Name:          req.Name,
		ContainerName: req.ContainerName,
		ServerID:      req.ServerID,
		NodeID:        req.NodeID,
	}
	if err := serverPointM.Create(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (sp *serverPoint) Update(ctx *gin.Context) {
	var req params.ServerPointUpdateReq

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if req.NodeID != 0 {
		nodeM := &model.Node{ID: req.NodeID}
		exist, err := nodeM.QueryOne("id")
		if err != nil {
			rest.Error(ctx, err.Error())
			return
		}
		if !exist {
			rest.Error(ctx, "网络节点不存在")
			return
		}
	}

	serverPointM := &model.ServerPoint{
		ID:            req.ID,
		Name:          req.Name,
		ContainerName: req.ContainerName,
		NodeID:        req.NodeID,
		ShPath:        req.ShPath,
		ConfPath:      req.ConfPath,
	}
	if err := serverPointM.Update(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (sp *serverPoint) Delete(ctx *gin.Context) {
	var req params.IDReq
	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	serverPointM := &model.ServerPoint{
		ID: req.ID,
	}
	if err := serverPointM.Delete(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}
