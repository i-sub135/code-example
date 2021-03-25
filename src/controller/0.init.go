package controller

import (
	"code-example/src/addons"
	"code-example/src/addons/omitField"
	"code-example/src/addons/response"
	"code-example/src/models"
)

type (
	RestController struct{}
)

var (
	respon        response.RestResponse
	adons         addons.RestAddons
	model         models.RestModels
	displayFields omitField.OmitFields
)
