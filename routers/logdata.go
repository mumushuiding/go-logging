package routers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mumushuiding/go-logging/service/logdataService"
	"github.com/mumushuiding/go-logging/util"
)

// SaveLogdata 保存日志
func SaveLogdata(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	map1, err := util.Body2Map(request)
	if err != nil {
		util.Response(writer, fmt.Sprintf("%s", err), false)
		return
	}
	if err := logdataService.Save(map1); err != nil {
		util.Response(writer, fmt.Sprintf("%s", err), false)
		return
	}
	util.Response(writer, "成功", true)
}

// DeleteLogdataByID 根据id删除日志
func DeleteLogdataByID(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm() // 这句话不能少，不然取不到值
	var ids []string = request.Form["id"]
	if len(ids) == 0 {
		util.Response(writer, "id 不存在", false)
		return
	}
	id, err := strconv.Atoi(ids[0])
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	err = logdataService.Delete(id)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	util.ResponseOk(writer)
}

// FindLogdatas 查询
func FindLogdatas(w http.ResponseWriter, r *http.Request) {
	var log = logdataService.LogdataReceiver{PageIndex: 1, PageSize: 10}
	err := util.Body2Struct(r, &log)
	datas, err := log.FindAllPageAsJSON()
	if err != nil {
		util.ResponseErr(w, err)
		return
	}
	fmt.Fprintf(w, "%s", datas)
}
