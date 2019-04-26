package model

import (
	"github.com/jinzhu/gorm"
)

// Logdata 日志存储类型
type Logdata struct {
	Model
	Title string `json:"title"`
	Data  string `gorm:"size:1020" json:"data"`
	From  string `json:"from"`
}

// FindLogdatas 分页查询
func FindLogdatas(pageIndex int, pageSize int, maps interface{}) ([]*Logdata, int, error) {
	var logdatas []*Logdata
	var count int
	err := db.Where(maps).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&logdatas).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	err = db.Model(&Logdata{}).Where(maps).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return logdatas, count, nil
}

// SaveLogdata 保存
func SaveLogdata(data map[string]interface{}) error {
	var logdata = Logdata{
		Title: data["title"].(string),
		Data:  data["data"].(string),
		From:  data["from"].(string),
	}
	err := db.Create(&logdata).Error
	return err
}

// DeleteLogdata 根据id删除
func DeleteLogdata(id int) error {
	err := db.Where("id = ?", id).Delete(Logdata{}).Error
	return err
}
