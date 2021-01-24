package beego

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/xm-chentl/go-mvc"
	ctxprop "github.com/xm-chentl/go-mvc/ctx-prop"
	"github.com/xm-chentl/go-mvc/enum"
)

type beegoEx struct {
	handler     mvc.IHandler
	mode        enum.RouteMode
	routeFormat string
}

func (b *beegoEx) SetHandle(handler mvc.IHandler) mvc.IMvc {
	b.handler = handler
	return b
}

func (b beegoEx) Run(port int) {
	beego.Post("/:service/:action", func(ctx *context.Context) {
		b.handler.Exec(ctxprop.Context{
			enum.CTX:         route{ctx: ctx},
			enum.ServerName:  ctx.Input.Param(":server"),
			enum.ServiceName: ctx.Input.Param(":service"),
			enum.ActionName:  ctx.Input.Param(":action"),
			enum.RespFunc: func(code int, data interface{}) {
				res, _ := json.Marshal(data)
				ctx.Output.Body(res)
			},
		})
	})
	fmt.Println("端口: ", port)
	beego.Run(
		fmt.Sprintf(":%d", port),
	)
}

func (b *beegoEx) route() string {
	if b.routeFormat == "" {
		switch b.mode {
		case enum.ThreeMode:
			b.routeFormat = "/:server/:service/:action"
			break
		default:
			b.routeFormat = "/:service/:action"
		}
	}

	return b.routeFormat
}

// New 实例
func New() mvc.IMvc {
	return new(beegoEx)
}

// NewMode 实例新的模式
func NewMode(mode enum.RouteMode) mvc.IMvc {
	return &beegoEx{
		mode: mode,
	}
}
