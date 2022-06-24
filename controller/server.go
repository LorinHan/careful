package controller

import (
	"careful/model"
	"careful/pkg/define/params"
	"careful/pkg/zip"
	"github.com/gin-gonic/gin"
	"github.com/zenfire-cn/commkit/utility"
	"github.com/zenfire-cn/webkit/rest"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

var Server = &server{}

type server struct {
}

func (s *server) FileRead(ctx *gin.Context) {
	var (
		req params.DockerRunReq
	)

	p := ctx.Query("path")
	req.Path = p
	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if !utility.PathFileExists(req.Path) {
		rest.Error(ctx, "文件不存在")
		return
	}

	f, err := os.Open(req.Path)
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	defer f.Close()

	// 读取文件全部内容
	b, err := ioutil.ReadAll(f)
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, string(b))
}

func (s *server) FileWrite(ctx *gin.Context) {
	var (
		req params.ConfWriteReq
	)

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if !utility.PathFileExists(req.Path) {
		rest.Error(ctx, "文件不存在")
		return
	}

	f, err := os.OpenFile(req.Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	// offset
	// os.Truncate(filename, 0) //clear
	// n, err := f.Seek(0, os.SEEK_END)
	// if err != nil {
	// 	rest.Error(ctx, err.Error())
	// 	return
	// }
	_, err = f.WriteString(req.Content)
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	defer f.Close()

	rest.Success(ctx, nil)
}

func (s *server) Upload(ctx *gin.Context) {
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

func (s *server) Create(ctx *gin.Context) {
	var req params.ServerCreateReq

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	folderModel := &model.Folder{ID: req.FolderID}
	exist, err := folderModel.QueryOne("id")
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	if !exist {
		rest.Error(ctx, "文件夹不存在")
		return
	}

	serverModel := &model.Server{Name: req.Name, FolderID: req.FolderID}
	if err := serverModel.Create(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (s *server) Update(ctx *gin.Context) {
	var req params.ServerUpdateReq
	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	serverModel := &model.Server{
		ID:   req.ID,
		Name: req.Name,
	}
	if err := serverModel.Update(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (s *server) Delete(ctx *gin.Context) {
	var req params.IDReq
	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	serverModel := &model.Server{
		ID: req.ID,
	}
	if err := serverModel.Delete(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}
