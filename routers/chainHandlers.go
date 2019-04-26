package routers

import (
	"net/http"

	"github.com/mumushuiding/go-logging/util"
)

// PostHandler 对所有post请求进行拦截
func PostHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			util.ResponseErr(w, "只支持Post方法！！")
			return
		}
		h(w, r)
	}
}
