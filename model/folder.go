package model

import (
	"careful/pkg/utils"
	"github.com/zenfire-cn/commkit/database"
	"gorm.io/gorm"
)

type Folder struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"not null;default:'';comment:文件夹名称;" json:"name"`
	CreatedAt utils.Time `json:"created_at,omitempty"`
	UpdatedAt utils.Time `json:"updated_at,omitempty"`
}

func (f *Folder) Create() error {
	return database.GetDB().Create(f).Error
}

func (f *Folder) Update() error {
	return database.GetDB().Updates(f).Error
}

func (f *Folder) Delete() error {
	return database.GetDB().Delete(f).Error
}

func (f *Folder) Query(col ...string) ([]Folder, error) {
	var (
		list      []Folder
		selectCol = "*"
	)
	if len(col) > 0 {
		selectCol = col[0]
	}
	if err := database.GetDB().Select(selectCol).Where(f).Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (f *Folder) QueryOne(col ...string) (bool, error) {
	var selectCol = "*"
	if len(col) > 0 {
		selectCol = col[0]
	}

	if err := database.GetDB().Select(selectCol).Where(f).First(f).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
