package library

// 返回数据处理 TODO:: data数据及json数据可选处理
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponseBody() *ResponseBody {
	return &ResponseBody{
		Code:    ERRNOSUCCESS,
		Message: GetErrMsg(ERRNOSUCCESS),
		Data:    map[string]interface{}{},
	}
}

func (rep *ResponseBody) SetData(data interface{}) {
	log.Println("----", data)
	rep.Data = data
}

func (rep *ResponseBody) SetCode(code int32) {
	rep.Code = code
}

func (rep *ResponseBody) SetMessage(message string) {
	rep.Message = message
}

func RecoverResponse(c *gin.Context, responseBody *ResponseBody) {
	// panic
	if err := recover(); err != nil {
		responseBody.SetCode(ERRNOUNKNOW)
	}

	resp, err := json.Marshal(responseBody)
	if err != nil {

		c.Data(http.StatusOK, "application/json;charset=utf-8", []byte(`{"code":2,"message":"unknown"}`))
	} else {
		c.Data(http.StatusOK, "application/json;charset=utf-8", resp)
	}
	return

}
