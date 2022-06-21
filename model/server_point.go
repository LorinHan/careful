package model

import (
	"careful/pkg/utils"
	"github.com/zenfire-cn/commkit/database"
)

type ServerPoint struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	Name          string     `gorm:"not null;default:'';comment:服务节点名称;" json:"name"`
	ContainerName string     `gorm:"not null;default:'';comment:容器名称;" json:"container_name"`
	ServerID      uint       `gorm:"not null;comment:所属服务;" json:"server_id"`
	NodeID        uint       `gorm:"not null;comment:所属服务器节点;" json:"node_id"`
	ShPath        string     `gorm:"not null;default:'';comment:脚本存放路径;" json:"sh_path"`
	ConfPath      string     `gorm:"not null;default:'';comment:配置文件存放路径;" json:"conf_path"`
	CreatedAt     utils.Time `json:"created_at,omitempty"`
	UpdatedAt     utils.Time `json:"updated_at,omitempty"`
}

func (sp *ServerPoint) Create() error {
	return database.GetDB().Create(sp).Error
}

func (sp *ServerPoint) Update() error {
	return database.GetDB().Updates(sp).Error
}

func (sp *ServerPoint) Delete() error {
	return database.GetDB().Delete(sp).Error
}

func (sp *ServerPoint) List(folderID uint) ([]map[string]interface{}, error) {
	var list []map[string]interface{}
	if err := database.GetDB().Raw(`
		SELECT sp.*, n.name AS node_name, s.name AS server_name
		FROM server_points AS sp, nodes AS n, servers AS s
		WHERE sp.node_id = n.id AND sp.server_id = s.id AND s.folder_id = ?;`, folderID).
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
