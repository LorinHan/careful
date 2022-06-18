package controller

import (
	"careful/model"
	"careful/pkg/define/params"
	"github.com/gin-gonic/gin"
	"github.com/zenfire-cn/webkit/rest"
)

var Node = &node{}

type node struct {
}

func (n *node) Create(ctx *gin.Context) {
	var req params.NodeCreateReq

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	nodeModel := &model.Node{Name: req.Name, IP: req.IP}
	if err := nodeModel.Create(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (n *node) List(ctx *gin.Context) {
	nodeModel := &model.Node{}
	list, err := nodeModel.Query("*")
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, list)
}

func (n *node) Update(ctx *gin.Context) {
	var req params.NodeUpdateReq
	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	nodeModel := &model.Node{
		ID:   req.ID,
		Name: req.Name,
		IP:   req.IP,
	}
	if err := nodeModel.Update(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (n *node) Delete(ctx *gin.Context) {
	var req params.IDReq
	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	nodeModel := &model.Node{
		ID: req.ID,
	}
	if err := nodeModel.Delete(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}
