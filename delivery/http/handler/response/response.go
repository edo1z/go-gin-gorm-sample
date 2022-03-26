package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status       int         `json:"status"`
	Error        int         `json:"error"`
	ErrorMessage string      `json:"error_message"`
	Result       interface{} `json:"result"`
}

func NewResponse(status int, result interface{}, err error) *Response {
	res := &Response{
		Status: status,
		Result: result,
	}
	if err != nil {
		res.Error = 1
		res.ErrorMessage = err.Error()
	}
	return res
}

func Json(status int, result interface{}, err error, c *gin.Context) {
	res := NewResponse(status, result, err)
	c.JSON(status, res)
}
