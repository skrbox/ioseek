package handler

import (
	"path"
	"strings"

	"github.com/gin-gonic/gin"

	c "github.com/skrbox/ioseek/pkg/conf"
)

const (
	contentType = "content-type"
	jsonStyle   = "application/json"
	errorHTML   = "error.html"
)

// 响应对象
type Response interface {
	Do()
}

// url 统一处理
func U(url string) string {
	if *c.MetaUrlPrefix == "" {
		*c.MetaUrlPrefix = "/"
	}
	return path.Join(*c.MetaUrlPrefix, url)
}

// 应用错误码和状态码关系转换
func toh(code int) (hcode int) {
	hcode = 200
	if x := code / 100; x > 1 && x < 5 {
		hcode = code
	}
	return
}

// json 响应
type json struct {
	Succeed bool        `json:"succeed"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 实例化 json 响应
func NewJsonResponse() *json {
	return &json{}
}

func (j json) Do(c *gin.Context) {
	c.JSON(toh(j.Code), j)
}

func (j *json) WithCode(code int) *json {
	j.Code = code
	return j
}

func (j *json) WithSucceed(s bool) *json {
	j.Succeed = s
	return j
}

func (j *json) WithMessage(msg string) *json {
	j.Message = msg
	return j
}

func (j *json) WithData(data interface{}) *json {
	j.Data = data
	j.Succeed = true
	return j
}

func (j *json) WithError(e ApiErr) *json {
	j.WithSucceed(false).WithCode(e.Code).WithMessage(e.Message)
	return j
}

// html 响应
type html struct {
	template string
	h        gin.H
	code     int
}

// 实例化 html 响应
func NewHTMLResponse() *html {
	return &html{}
}

func (h *html) Do(c *gin.Context) {
	c.HTML(toh(h.code), h.template, h.h)
}

func (h *html) WithTemplate(tmpl string) *html {
	h.template = tmpl
	return h
}

func (h *html) WithH(gh gin.H) *html {
	h.h = gh
	return h
}

func (h *html) WithCode(code int) *html {
	h.code = code
	return h
}

func (h *html) WithError(e ApiErr) *html {
	h.WithTemplate(errorHTML).WithCode(e.Code).WithH(gin.H{
		"title":   e.Code,
		"code":    e.Code,
		"message": e.Message,
	})
	return h
}

// 自定义 404 处理
func Handle404(c *gin.Context) {
	if strings.Contains(c.GetHeader(contentType), jsonStyle) {
		NewJsonResponse().WithError(HTTP404).Do(c)
		return
	}
	NewHTMLResponse().WithError(HTTP404).Do(c)
}

// 自定义 405 处理
func Handle405(c *gin.Context) {
	if strings.Contains(c.GetHeader(contentType), jsonStyle) {
		NewJsonResponse().WithError(HTTP405).Do(c)
	}
	NewHTMLResponse().WithError(HTTP405).Do(c)
}

// 健康状态反馈
func handlePing(c *gin.Context) {
	NewJsonResponse().WithSucceed(true).WithMessage("pong").Do(c)
}

// 业务指标暴露: 友链状态，攻击行为等
func handleMetrics(c *gin.Context) {
	// todo: ...
	NewHTMLResponse().WithError(NotCompleted).Do(c)
}

// 服务版本信息查询
func handleVersion(ctx *gin.Context) {
	NewJsonResponse().WithData(gin.H{
		"app":      c.MetaAppName,
		"version":  c.MetaVersion,
		"commitId": c.MetaCommitId,
		"branch":   c.MetaBranch,
		"buildAt":  c.MetaBuildAt,
		"platform": c.MetaPlatform,
		"webPage":  c.MetaWebPage,
	}).Do(ctx)
}

// 全局路由注册信息
func routerPaths(e *gin.Engine) gin.HandlerFunc {
	type Path struct {
		Path   string
		Method string
	}
	var (
		paths  = make([]Path, 0)
		routes = e.Routes()
	)
	for _, route := range routes {
		paths = append(paths, Path{
			Path:   route.Path,
			Method: route.Method,
		})
	}
	return func(ctx *gin.Context) {
		NewJsonResponse().WithData(paths).Do(ctx)
	}
}
