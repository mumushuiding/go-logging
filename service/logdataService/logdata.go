package logdataService

import (
	"github.com/mumushuiding/go-logging/model"
	"github.com/mumushuiding/go-logging/util"
)

// LogdataReceiver LogdataReceiver
type LogdataReceiver struct {
	Title     string `json:"title"`
	Data      string `json:"data"`
	From      string `json:"from"`
	PageSize  int    `json:"pageSize"`
	PageIndex int    `json:"pageIndex"`
}

// Save 保存
func Save(data map[string]interface{}) error {
	err := model.SaveLogdata(data)
	return err
}

// Delete 根据id删除
func Delete(id int) error {
	err := model.DeleteLogdata(id)
	return err
}

// FindAll 分页查询
func (l *LogdataReceiver) FindAll() ([]*model.Logdata, int, error) {
	var page = model.Page{}
	page.PageRequest(l.PageIndex, l.PageSize)
	datas, count, err := model.FindLogdatas(page.PageIndex, page.PageSize, l.getMaps())
	return datas, count, err
}

// FindAllPageAsJSON 分页查询并返回json
func (l *LogdataReceiver) FindAllPageAsJSON() (string, error) {
	datas, count, err := l.FindAll()
	if err != nil {
		return "", err
	}
	return util.ToPageJSON(datas, count, l.PageIndex, l.PageSize)
}
func (l *LogdataReceiver) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if len(l.Title) > 0 {
		maps["title"] = l.Title
	}
	return maps
}

// // FindAll 分页查询
// func FindAll(data map[string]interface{}) ([]*model.Logdata, error) {
// 	var page = model.Page{}
// 	page.PageRequest(data["pageIndex"], data["pageSize"])
// 	datas, err := model.FindLogdatas(page.PageIndex, page.PageSize, data)
// 	return datas, err
// }
