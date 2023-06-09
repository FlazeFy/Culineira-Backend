package response

type (
	Response struct {
		Status  int         `json:"status"`
		Message string      `json:"msg"`
		Data    interface{} `json:"data"`
	}
	ContentResponse struct {
		Response
		Total int `json:"total"`
	}
	LoginResponse struct {
		Response
		Token string `json:"token"`
	}
)
