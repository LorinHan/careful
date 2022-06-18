package controller

import (
	"careful/model"
	"careful/pkg/define/params"
	"github.com/gin-gonic/gin"
	"github.com/zenfire-cn/webkit/rest"
)

var Folder = &folder{}

type folder struct {
}

func (f *folder) Create(ctx *gin.Context) {
	var req params.FolderCreateReq

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	folderModel := &model.Folder{Name: req.Name}
	if err := folderModel.Create(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (f *folder) List(ctx *gin.Context) {
	folderModel := &model.Folder{}
	list, err := folderModel.Query("id, name")
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, list)
}

func (f *folder) Update(ctx *gin.Context) {
	var req params.FolderUpdateReq
	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	folderModel := &model.Folder{
		ID:   req.ID,
		Name: req.Name,
	}
	if err := folderModel.Update(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (f *folder) Delete(ctx *gin.Context) {
	var req params.IDReq
	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	folderModel := &model.Folder{
		ID: req.ID,
	}
	if err := folderModel.Delete(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}
