package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/khanhct/go-lib-core/sdk/response"
	"gorm.io/gorm"
)

var DefaultLanguage = "vi-VN"

type Api struct {
	Context *gin.Context
	Orm     *gorm.DB
	Errors  error
}

func (e *Api) AddError(err error) {
	if err != nil {
		e.Errors = err
	}
}

func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	return e
}

func (e *Api) MakeOrm(f func(c *gin.Context) (*gorm.DB, error)) *Api {
	if f == nil {
		f = getOrm
	}

	db, err := f(e.Context)
	if err != nil {
		e.AddError(err)
	}
	e.Orm = db
	return e
}

func getOrm(c *gin.Context) (*gorm.DB, error) {
	idb, exist := c.Get("db")
	if !exist {
		return nil, errors.New("db connect not exist")
	}
	switch idb := idb.(type) {
	case *gorm.DB:
		return idb, nil
	default:
		return nil, errors.New("db connect not exist")
	}
}

func (e Api) Error(code int, err error, msg string) {
	response.Error(e.Context, code, err, msg)
}

func (e Api) OK(data interface{}, msg string) {
	response.OK(e.Context, data, msg)
}

func (e Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(e.Context, result, count, pageIndex, pageSize, msg)
}

func (e Api) Custom(data gin.H) {
	response.Custum(e.Context, data)
}
