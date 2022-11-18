package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/jwcjf/go-project-base/logger"
	"github.com/jwcjf/go-project-base/sdk/pkg"
	"github.com/jwcjf/go-project-base/sdk/pkg/response"
	"github.com/jwcjf/go-project-base/sdk/service"
	"github.com/jwcjf/go-project-base/tools/language"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// DefaultLanguage ...
var DefaultLanguage = "zh-CN"

// Api ...
type Api struct {
	Context *gin.Context
	Logger  *logger.Helper
	Orm     *gorm.DB
	Errors  error
}

// AddError ...
func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Logger.Error(err)
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}

// MakeContext 设置http上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	e.Logger = GetRequestLogger(c)
	return e
}

// GetLogger 获取上下文提供的日志
func (e Api) GetLogger() *logger.Helper {
	return GetRequestLogger(e.Context)
}

// Bind 参数校验
func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = []binding.Binding{nil, binding.JSON}
	}
	needValidateNum := len(bindings) - 1
	for i := range bindings {
		if bindings[i] == nil {
			err = e.Context.ShouldBindUri(d)
		} else {
			err = e.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			e.Logger.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if ok && i < needValidateNum {
				err = nil
				continue
			}
			trans, errT := transInit(e.getAcceptLanguage())
			if errT != nil {
				err = fmt.Errorf(errT.Error()+", %w", err)
				e.AddError(err)
				return e
			}
			validatorErrs := errs.Translate(trans)
			strArr := make([]string, 0)
			for k, v := range validatorErrs {
				strArr = append(strArr, k+":"+v)
			}
			err = errors.New(strings.Join(strArr, ","))
			e.AddError(err)
			return e
		}
	}
	return e
}

// GetOrm 获取Orm DB
func (e Api) GetOrm() (*gorm.DB, error) {
	db, err := pkg.GetOrm(e.Context, "")
	if err != nil {
		e.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		return nil, err
	}
	return db, nil
}

// MakeOrm 设置Orm DB
func (e *Api) MakeOrm() *Api {
	var err error
	if e.Logger == nil {
		err = errors.New("at MakeOrm logger is nil")
		e.AddError(err)
		return e
	}
	db, err := pkg.GetOrm(e.Context, "")
	if err != nil {
		e.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		e.AddError(err)
	}
	e.Orm = db
	return e
}

// GetOrmByKey 获取Orm DB
func (e Api) GetOrmByKey(key string) (*gorm.DB, error) {
	db, err := pkg.GetOrm(e.Context, key)
	if err != nil {
		e.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		return nil, err
	}
	return db, nil
}

// MakeOrmByKey 设置Orm DB
func (e *Api) MakeOrmByKey(key string) *Api {
	var err error
	if e.Logger == nil {
		err = errors.New("at MakeOrm logger is nil")
		e.AddError(err)
		return e
	}
	db, err := pkg.GetOrm(e.Context, key)
	if err != nil {
		e.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		e.AddError(err)
	}
	e.Orm = db
	return e
}

// MakeService ...
func (e *Api) MakeService(c *service.Service) *Api {
	c.Log = e.Logger
	c.Orm = e.Orm
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, err error, msg string) {
	response.Error(e.Context, code, err, msg)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg string) {
	response.OK(e.Context, data, msg)
}

// PageOK 分页数据处理
func (e Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(e.Context, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (e Api) Custom(data gin.H) {
	response.Custum(e.Context, data)
}

// getAcceptLanguage 获取当前语言
func (e *Api) getAcceptLanguage() string {
	languages := language.ParseAcceptLanguage(e.Context.GetHeader("Accept-Language"), nil)
	if len(languages) == 0 {
		return DefaultLanguage
	}
	return languages[0]
}
