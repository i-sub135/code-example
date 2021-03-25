package routers

import "github.com/gin-gonic/gin"

func (r *RestRouter) IndexRoute(eng *gin.Engine) {
	eng.GET("", Ctrl.IndexCtrl)
}

func (r *RestRouter) AuthRoute(g *gin.RouterGroup) {
	g.POST("", Ctrl.AuthReqCodeCtrl)
	g.POST("/validate", Ctrl.AuthValidateCtrl)
}

func (r *RestRouter) ProdukRoute(g *gin.RouterGroup) {
	g.GET("", Ctrl.AllProductCtrl)
	g.POST("", Ctrl.AddProductCtrl)
	g.GET("/:product_id", Ctrl.DetailProductCtrl)
	g.PUT("/:product_id", Ctrl.UpdateProductCtrl)
	g.DELETE("/:product_id", Ctrl.DeleteProductCtrl)
}
