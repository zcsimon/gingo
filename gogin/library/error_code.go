// 定义工具类，错误code 及提示
package library

const (
	ERRNOSUCCESS = 0
	ERRNOERROR   = 1
	ERRNOUNKNOW  = 2
)

var ErrToMsgMap = map[int32]string{
	ERRNOSUCCESS: "success",
	ERRNOERROR:   "failed",
	ERRNOUNKNOW:  "unknow",
}

func GetErrMsg(errNo int32) string {
	if msg, ok := ErrToMsgMap[errNo]; ok {
		return msg
	}
	return "Unknow error"
}
