package model

import (
	"github.com/zenfire-cn/commkit/database"
	"gorm.io/gorm"
	"time"
)

type Server struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;default:'';comment:服务名称;" json:"name"`
	FolderID  uint      `gorm:"not null;default:0;comment:所属文件夹;" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Server) Create() error {
	return database.GetDB().Create(s).Error
}

func (s *Server) Update() error {
	return database.GetDB().Updates(s).Error
}

func (s *Server) Delete() error {
	return database.GetDB().Delete(s).Error
}

func (s *Server) QueryOne(col ...string) (bool, error) {
	var selectCol = "*"
	if len(col) > 0 {
		selectCol = col[0]
	}

	if err := database.GetDB().Select(selectCol).Where(s).First(s).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
