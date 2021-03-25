package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (r *RestMiddleware) ReqValidator(res *gin.Context) {
	var (
		token string = res.GetHeader("token")
	)

	switch {
	case token == "":
		out := respon.RespBad(401, "Token can't be empty", map[string]string{})
		res.JSON(out.Code, out)
		res.Abort()
	default:
		_, err := addOns.ValidJwt(token)
		if err != nil {
			out := respon.RespBad(401, "Token Not Valid :: "+err.Error()+" [ Please Get Accses on Endpoint /auth ]", map[string]string{})
			res.JSON(out.Code, out)
			res.Abort()
			return
		}
		fmt.Println(res.Writer.Status())
		res.Next()
	}

}
