package middleware

import (
	"code-example/src/addons"
	"code-example/src/addons/response"
)

type RestMiddleware struct{}

var (
	addOns addons.RestAddons
	respon response.RestResponse
)
