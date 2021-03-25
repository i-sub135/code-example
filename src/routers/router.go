package routers

import "github.com/gin-gonic/gin"

func (r *RestRouter) IndexRoute(eng *gin.Engine) {
	eng.GET("", Ctrl.IndexCtrl)
}

func (r *RestRouter) AuthRoute(g *gin.RouterGroup) {
	g.POST("", Ctrl.AuthReqCodeCtrl)
}
