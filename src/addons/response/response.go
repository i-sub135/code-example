package response

import "net/http"

// RespOK -- response Output OK
func (r *RestResponse) RespOK(data interface{}) Output {
	out := Output{
		Code:    http.StatusOK,
		Message: "OK",
		Result:  data,
	}

	return out
}

// RespBad -- response bad request
func (r *RestResponse) RespBad(errCode int, errMsg string, data interface{}) Output {
	out := Output{
		Code:      http.StatusBadRequest,
		Message:   "Bad Requests",
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
		Result:    data,
	}

	return out
}

// RespNotFound -- response 404
func (r *RestResponse) RespNotFound(errCode int, errMsg string, data interface{}) Output {
	out := Output{
		Code:      http.StatusNotFound,
		Message:   "Not Found",
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
		Result:    data,
	}
	return out
}

// RespNotFound -- response 404
func (r *RestResponse) RespUnAuth(errCode int, errMsg string, data interface{}) Output {
	out := Output{
		Code:      http.StatusUnauthorized,
		Message:   "Status Unauthorized",
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
		Result:    data,
	}
	return out
}
