package api

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/khanhct/go-lib-core/sdk/response"
	"github.com/khanhct/go-lib-core/sdk/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DefaultLanguage = "vi-VN"

type Api struct {
	Context *gin.Context
	Orm     *gorm.DB
	Errors  error
	Logger  *logrus.Logger
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

func (e *Api) MakeLogger(f func(c *gin.Context) (*logrus.Logger, error)) *Api {
	if f == nil {
		f = getLogger
	}

	logger, err := f(e.Context)
	if err != nil {
		e.AddError(err)
	}
	e.Logger = logger
	return e
}

func getLogger(c *gin.Context) (*logrus.Logger, error) {
	iLogger, exist := c.Get("logger")
	if !exist {
		return nil, errors.New("logger connect not exist")
	}
	switch idb := iLogger.(type) {
	case *logrus.Logger:
		return idb, nil
	default:
		return nil, errors.New("logger connect not exist")
	}
}

func (e *Api) MakeService(s *service.Service) *Api {
	s.Orm = e.Orm
	s.Logger = e.Logger
	return e
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

func (e Api) GetClient() string {
	ClientIP := e.Context.ClientIP()
	//fmt.Println("ClientIP:", ClientIP)
	RemoteIP, _ := e.Context.RemoteIP()
	//fmt.Println("RemoteIP:", RemoteIP)
	ip := e.Context.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = e.Context.Request.Header.Get("X-real-ip")
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	if RemoteIP.String() != "127.0.0.1" {
		ip = RemoteIP.String()
	}
	if ClientIP != "127.0.0.1" {
		ip = ClientIP
	}
	return ip
}
