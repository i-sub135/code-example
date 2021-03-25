package controller

import (
	"code-example/src/store"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (r *RestController) AuthReqCodeCtrl(res *gin.Context) {
	var (
		param store.AuthParams
	)
	_ = res.BindJSON(&param)

	switch param.Phone {
	case "":
		out := respon.RespBad(1, "phone number can't be empty", map[string]string{})
		res.JSON(out.Code, out)
	default:
		ranNumb := adons.EncodeToString(6)
		keys := fmt.Sprintf("OTP-%v", param.Phone)
		err := model.SetRedis(keys, ranNumb)
		if err != nil {
			out := respon.RespBad(1, err.Error(), map[string]string{})
			res.JSON(out.Code, out)
			return
		}
		out := respon.RespOK(map[string]interface{}{"OTP": ranNumb})
		res.JSON(out.Code, out)
	}

}

func (r *RestController) AuthValidateCtrl(res *gin.Context) {
	var (
		param store.AuthParams
	)
	_ = res.BindJSON(&param)

	switch {
	case param.Phone == "" || param.Otp == "":
		out := respon.RespBad(1, "phone number & otp can't be empty", map[string]string{})
		res.JSON(out.Code, out)
	default:
		key := fmt.Sprintf("OTP-%v", param.Phone)
		getOTP, err := model.GetOtp(key)
		if err != nil {
			out := respon.RespBad(1, "akses notfound request to endpoint /auth", map[string]string{})
			res.JSON(out.Code, out)
			return
		}
		if getOTP != param.Otp {
			out := respon.RespBad(1, "phone & otp not match", map[string]string{})
			res.JSON(out.Code, out)
			return
		}

		token, _ := adons.CreateJwt(param.Phone)
		out := respon.RespOK(map[string]string{
			"token": token,
		})
		res.JSON(out.Code, out)

	}

}
