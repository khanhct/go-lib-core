package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Default = &response{}

const (
	RequestId = "X-Request-Id"
)

// GenerateMsgIDFromContext
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(RequestId)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(RequestId, requestId)
	}
	return requestId
}

func Error(c *gin.Context, code int, err error, msg string) {
	res := Default.Clone()
	if err != nil {
		res.SetMsg(err.Error())
	}
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(GenerateMsgIDFromContext(c))
	res.SetCode(code)
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(code, res)
}

func OK(c *gin.Context, data interface{}, msg string) {
	res := Default.Clone()
	res.SetData(data)
	res.SetSuccess(true)
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(GenerateMsgIDFromContext(c))
	res.SetCode(http.StatusOK)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}

func Custum(c *gin.Context, data gin.H) {
	data["requestId"] = GenerateMsgIDFromContext(c)
	c.Set("result", data)
	c.AbortWithStatusJSON(http.StatusOK, data)
}
