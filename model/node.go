package model

import (
	"careful/pkg/utils"
	"github.com/zenfire-cn/commkit/database"
	"gorm.io/gorm"
)

type Node struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"not null;default:'';comment:节点名称;" json:"name"`
	IP        string     `gorm:"not null;default:'';comment:ip;" json:"ip"`
	CreatedAt utils.Time `json:"created_at,omitempty"`
	UpdatedAt utils.Time `json:"updated_at,omitempty"`
}

func (n *Node) Create() error {
	return database.GetDB().Create(n).Error
}

func (n *Node) Update() error {
	return database.GetDB().Updates(n).Error
}

func (n *Node) Delete() error {
	return database.GetDB().Delete(n).Error
}

func (n *Node) Query(col ...string) ([]Node, error) {
	var (
		list      []Node
		selectCol = "*"
	)
	if len(col) > 0 {
		selectCol = col[0]
	}
	if err := database.GetDB().Select(selectCol).Where(n).Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (n *Node) QueryOne(col ...string) (bool, error) {
	var selectCol = "*"
	if len(col) > 0 {
		selectCol = col[0]
	}

	if err := database.GetDB().Select(selectCol).Where(n).First(n).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
