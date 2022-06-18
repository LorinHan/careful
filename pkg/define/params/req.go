package params

import "errors"

type NodeCreateReq struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

func (n *NodeCreateReq) Check() error {
	if n.Name == "" {
		return errors.New("参数 name 为空")
	}
	if n.IP == "" {
		return errors.New("参数 ip 为空")
	}
	return nil
}

type FolderCreateReq struct {
	Name string `json:"name"`
}

func (f *FolderCreateReq) Check() error {
	if f.Name == "" {
		return errors.New("参数 name 为空")
	}
	return nil
}

type FolderUpdateReq struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (f *FolderUpdateReq) Check() error {
	if f.ID == 0 {
		return errors.New("参数 id 为空")
	}
	if f.Name == "" {
		return errors.New("参数 name 为空")
	}
	return nil
}

type NodeUpdateReq struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}

func (n *NodeUpdateReq) Check() error {
	if n.ID == 0 {
		return errors.New("参数 id 为空")
	}
	if n.Name == "" {
		return errors.New("参数 name 为空")
	}
	if n.IP == "" {
		return errors.New("参数 ip 为空")
	}
	return nil
}

type IDReq struct {
	ID uint `json:"id"`
}

func (i *IDReq) Check() error {
	if i.ID == 0 {
		return errors.New("参数 id 为空")
	}
	return nil
}

type DockerListReq struct {
	Name []string `json:"name"`
}

func (d *DockerListReq) Check() error {
	if len(d.Name) == 0 {
		return errors.New("参数 name 为空")
	}
	return nil
}

type DockerNameReq struct {
	Name string `json:"name"`
}

func (d *DockerNameReq) Check() error {
	if len(d.Name) == 0 {
		return errors.New("参数 name 为空")
	}
	return nil
}

type DockerStartReq struct {
	DockerNameReq
	Restart bool `json:"restart"`
}

func (d *DockerStartReq) Check() error {
	return d.DockerNameReq.Check()
}

type DockerLogsReq struct {
	DockerNameReq
	Tails string `json:"tails"`
}

func (d *DockerLogsReq) Check() error {
	if d.Tails == "" {
		return errors.New("参数 tails 为空")
	}
	return d.DockerNameReq.Check()
}

type DockerRunReq struct {
	Path string `json:"path"`
}

func (d *DockerRunReq) Check() error {
	if d.Path == "" {
		return errors.New("参数 path 为空")
	}
	return nil
}

type ServerCreateReq struct {
	Name     string `json:"name"`
	FolderID uint   `json:"folder_id"`
}

func (s *ServerCreateReq) Check() error {
	if s.Name == "" {
		return errors.New("参数 name 为空")
	}
	if s.FolderID == 0 {
		return errors.New("参数 folder_id 为空")
	}
	return nil
}

type ServerUpdateReq struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s *ServerUpdateReq) Check() error {
	if s.ID == 0 {
		return errors.New("参数 id 为空")
	}
	if s.Name == "" {
		return errors.New("参数 name 为空")
	}
	return nil
}

type ServerPointCreateReq struct {
	Name          string `json:"name"`
	ContainerName string `json:"container_name"`
	ServerID      uint   `json:"server_id"`
	NodeID        uint   `json:"node_id"`
}

func (sp *ServerPointCreateReq) Check() error {
	if sp.Name == "" {
		return errors.New("参数 name 为空")
	}
	if sp.ContainerName == "" {
		return errors.New("参数 container_name 为空")
	}
	if sp.ServerID == 0 {
		return errors.New("参数 server_id 为空")
	}
	if sp.NodeID == 0 {
		return errors.New("参数 node_id 为空")
	}
	return nil
}

type ServerPointUpdateReq struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	ContainerName string `json:"container_name"`
	NodeID        uint   `json:"node_id"`
	ShPath        string `json:"sh_path"`
	ConfPath      string `json:"conf_path"`
}

func (sp *ServerPointUpdateReq) Check() error {
	if sp.ID == 0 {
		return errors.New("参数 id 为空")
	}
	if sp.Name == "" && sp.NodeID == 0 && sp.ShPath == "" && sp.ConfPath == "" && sp.ContainerName == "" {
		return errors.New("参数 name, container_name, node_id, sh_path, conf_path 都为空")
	}
	return nil
}
