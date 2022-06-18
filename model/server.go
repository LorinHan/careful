package model

import (
	"time"
)

type Server struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FolderID  uint      `gorm:"not null;default:0;comment:所属文件夹;" json:"id"`
	NodeID    uint      `gorm:"not null;default:0;comment:所属节点;" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
