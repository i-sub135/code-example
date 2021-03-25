package response

type (
	Output struct {
		Code      int         `json:"Code"`
		Message   string      `json:"Message"`
		ErrorCode int         `json:"ErrorCode"`
		ErrorMsg  string      `json:"ErrorMsg"`
		Result    interface{} `json:"Result"`
	}
)
