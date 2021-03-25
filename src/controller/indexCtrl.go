package controller

import "github.com/gin-gonic/gin"

func (r *RestController) IndexCtrl(res *gin.Context) {
	out := respon.RespOK(map[string]string{
		"Pesan": "Welcome On Api Services",
	})
	res.JSON(out.Code, out)
}
