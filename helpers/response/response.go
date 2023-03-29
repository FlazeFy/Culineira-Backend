package response

type (
	Response struct {
		Status  int         `json:"status"`
		Message string      `json:"msg"`
		Data    interface{} `json:"data"`
	}
	ContentResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"msg"`
		Total   int         `json:"total"`
		Data    interface{} `json:"data"`
	}
)
