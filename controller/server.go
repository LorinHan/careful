package controller

import (
	"careful/pkg/zip"
	"github.com/gin-gonic/gin"
	"github.com/zenfire-cn/commkit/utility"
	"github.com/zenfire-cn/webkit/rest"
	"go.uber.org/zap"
	"os"
	"path"
	"path/filepath"
)

func TestCtrl(ctx *gin.Context) {
	logger := zap.L()
	num := ctx.Param("num")

	logger.Info("params", zap.String("num", num))

	rest.Success(ctx, "Hello webkit, param is "+num)
}

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	p := ctx.PostForm("path")
	if p == "" {
		rest.Error(ctx, "path 参数为空")
		return
	}

	if !utility.PathFileExists(p) {
		if err := os.MkdirAll(p, os.ModePerm); err != nil {
			rest.Error(ctx, err.Error())
			return
		}
	}
	dst := path.Join(p, file.Filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if filepath.Ext(file.Filename) == ".zip" {
		if err := zip.Unzip(dst, p); err != nil {
			rest.Error(ctx, err.Error())
			return
		}

		if err := os.Remove(dst); err != nil {
			rest.Error(ctx, err.Error())
			return
		}
	}

	rest.Success(ctx, nil)
}
