package beego

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/context"
)

type route struct {
	ctx *context.Context
}

func (r route) Bind(arge interface{}) {
	// r.ctx.Input.Bind 失败，获取不到参数
	buff, _ := ioutil.ReadAll(r.ctx.Request.Body)
	json.Unmarshal(buff, arge)
	r.ctx.Request.Body.Close()
}

func (r route) Request() *http.Request {
	return r.ctx.Request
}

func (r route) Response(data interface{}) {
	resp, _ := json.Marshal(data)
	r.ctx.Output.Body(resp)
}
