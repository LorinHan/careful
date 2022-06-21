package controller

import (
	"careful/pkg/define/params"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zenfire-cn/commkit/utility"
	"github.com/zenfire-cn/webkit/rest"
	"go.uber.org/zap"
	"io/ioutil"
	"os/exec"
	"runtime"
	"time"
)

var Docker = &docker{}

type docker struct {
}

func (d *docker) List(ctx *gin.Context) {
	var (
		req = &params.DockerListReq{
			Name: ctx.QueryArray("name"),
		}
		dockerListArgs []filters.KeyValuePair
		resp           []params.DockerListResp
	)

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	dockerListArgs = make([]filters.KeyValuePair, len(req.Name))
	for i := range req.Name {
		dockerListArgs[i] = filters.Arg("name", req.Name[i])
	}

	dockerCtx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	defer func() {
		if err := cli.Close(); err != nil {
			zap.L().Error(err.Error(), zap.String("docker cli", "close error"))
		}
	}()

	containers, err := cli.ContainerList(dockerCtx, types.ContainerListOptions{
		Filters: filters.NewArgs(dockerListArgs...),
	})
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	resp = make([]params.DockerListResp, len(containers))
	for i, container := range containers {
		name := container.Names[0]
		if name[:1] == "/" {
			name = name[1:]
		}
		resp[i] = params.DockerListResp{
			ID:          container.ID,
			Name:        name,
			Image:       container.Image,
			State:       container.State,
			Status:      container.Status,
			Created:     container.Created,
			PrivatePort: container.Ports[0].PrivatePort,
			PublicPort:  container.Ports[0].PublicPort,
		}
	}

	rest.Success(ctx, resp)
}

func (d *docker) Stop(ctx *gin.Context) {
	var (
		req params.DockerNameReq
	)

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	dockerCtx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	defer func() {
		if err := cli.Close(); err != nil {
			zap.L().Error(err.Error(), zap.String("docker cli", "close error"))
		}
	}()

	containers, err := cli.ContainerList(dockerCtx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("name", req.Name)),
	})
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if len(containers) < 1 {
		rest.Error(ctx, "该容器不存在")
		return
	}

	duration := time.Second * 10
	if err := cli.ContainerStop(ctx, containers[0].ID, &duration); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (d *docker) Start(ctx *gin.Context) {
	var (
		req params.DockerStartReq
	)

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	dockerCtx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	defer func() {
		if err := cli.Close(); err != nil {
			zap.L().Error(err.Error(), zap.String("docker cli", "close error"))
		}
	}()

	containers, err := cli.ContainerList(dockerCtx, types.ContainerListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", req.Name)),
	})
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if len(containers) < 1 {
		rest.Error(ctx, "该容器不存在")
		return
	}

	if req.Restart {
		duration := time.Second * 10
		if err := cli.ContainerRestart(ctx, containers[0].ID, &duration); err != nil {
			rest.Error(ctx, err.Error())
			return
		}
	} else {
		if err := cli.ContainerStart(ctx, containers[0].ID, types.ContainerStartOptions{}); err != nil {
			rest.Error(ctx, err.Error())
			return
		}
	}

	rest.Success(ctx, nil)
}

func (d *docker) Delete(ctx *gin.Context) {
	var (
		req params.DockerNameReq
	)

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	dockerCtx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	defer func() {
		if err := cli.Close(); err != nil {
			zap.L().Error(err.Error(), zap.String("docker cli", "close error"))
		}
	}()

	containers, err := cli.ContainerList(dockerCtx, types.ContainerListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", req.Name)),
	})
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if len(containers) < 1 {
		rest.Error(ctx, "该容器不存在")
		return
	}

	if err := cli.ContainerRemove(dockerCtx, containers[0].ID, types.ContainerRemoveOptions{}); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, nil)
}

func (d *docker) Logs(ctx *gin.Context) {
	var (
		req = &params.DockerLogsReq{
			DockerNameReq: params.DockerNameReq{
				Name: ctx.Query("name"),
			},
			Tails: ctx.Query("tails"),
		}
	)

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	dockerCtx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	defer func() {
		if err := cli.Close(); err != nil {
			zap.L().Error(err.Error(), zap.String("docker cli", "close error"))
		}
	}()

	containers, err := cli.ContainerList(dockerCtx, types.ContainerListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", req.Name)),
	})
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if len(containers) < 1 {
		rest.Error(ctx, "该容器不存在")
		return
	}

	logs, err := cli.ContainerLogs(dockerCtx, containers[0].ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       req.Tails,
	})
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	defer logs.Close()

	logBytes, err := ioutil.ReadAll(logs)
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	rest.Success(ctx, string(logBytes))
}

func (d *docker) Info(ctx *gin.Context) {
	var (
		req = &params.DockerNameReq{
			Name: ctx.Query("name"),
		}
	)

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	dockerCtx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}
	defer func() {
		if err := cli.Close(); err != nil {
			zap.L().Error(err.Error(), zap.String("docker cli", "close error"))
		}
	}()

	containers, err := cli.ContainerList(dockerCtx, types.ContainerListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", req.Name)),
	})
	if err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if len(containers) < 1 {
		rest.Success(ctx, nil)
		return
	}

	rest.Success(ctx, containers[0])
}

func (d *docker) Run(ctx *gin.Context) {
	var (
		req params.DockerRunReq
	)

	if err := ctx.Bind(&req); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	if runtime.GOOS != "windows" && !utility.PathFileExists(req.Path) {
		rest.Error(ctx, "执行脚本不存在")
		return
	}

	if _, err := exec.Command("chmod", "755", req.Path).Output(); err != nil {
		rest.Error(ctx, err.Error())
		return
	}

	var (
		cmd = exec.Command("sh", "-c", req.Path)
	)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", req.Path)
	}

	buf, err := cmd.CombinedOutput()
	if err != nil {
		rest.Success(ctx, errors.WithMessage(err, string(buf)).Error())
		return
	}
	rest.Success(ctx, string(buf))
}
