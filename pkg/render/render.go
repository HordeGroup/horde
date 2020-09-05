package render

import (
	"github.com/HordeGroup/horde/pkg/def"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type CodeMsgRender struct {
	Code     int
	Msg      string
	Callback func(s *jsoniter.Stream)
}

func (render CodeMsgRender) Render(w http.ResponseWriter) error {
	render.WriteContentType(w)
	return nil
}

func (render CodeMsgRender) WriteContentType(w http.ResponseWriter) {

}

func JSONWithError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusOK, &def.Resp{
		Code: def.CodeError,
		Data: nil,
		Msg:  err.Error(),
	})
}

func JSONSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &def.Resp{
		Code: def.CodeSuccess,
		Data: data,
		Msg:  def.MsgSuccess,
	})
}
