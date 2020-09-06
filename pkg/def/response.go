package def

const (
	CodeError   = -1
	CodeSuccess = 0

	MsgSuccess = "success"
)

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}
