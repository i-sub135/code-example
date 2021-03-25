package controller

import (
	"code-example/src/store"

	"github.com/gin-gonic/gin"
)

func (r *RestController) AuthReqCodeCtrl(res *gin.Context) {
	var (
		param store.AuthParams
	)
	res.BindJSON(&param)

	switch param.Phone {
	case "":
		out := respon.RespBad(1, "phone number can't be empty", map[string]string{})
		res.JSON(out.Code, out)
	default:

		out := respon.RespOK(param)
		res.JSON(out.Code, out)

	}

}
