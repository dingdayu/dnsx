package router

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/dingdayu/dnsx/api/controller/health"
	v1 "github.com/dingdayu/dnsx/api/controller/v1"
	"github.com/dingdayu/dnsx/api/controller/v1/record"
	"github.com/dingdayu/dnsx/api/middleware"
	"github.com/dingdayu/dnsx/model/entity"
	"github.com/dingdayu/dnsx/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/miekg/dns"
)

// httpHandler http handler
var httpHandler *gin.Engine

// initHTTPHandler init http handler
func initHTTPHandler() {
	httpHandler = gin.New()

	// 正式环境不再在控制台请求输出日志
	if gin.Mode() != gin.ReleaseMode {
		httpHandler.Use(gin.Logger())
	}
	httpHandler.Use(gin.Recovery())

	httpHandler.GET("/", health.Hello)

	httpHandler.GET("/dns-query", httpDNS)
	httpHandler.POST("/dns-query", httpDNS)

	httpHandler.HEAD("/health", health.Hello)
	httpHandler.GET("/health", health.Hello)
	httpHandler.GET("/ping", health.Ping)
	httpHandler.GET("/metrics", health.Prometheus)

	apiv1 := httpHandler.Group("/api/v1")

	// 根据配置决定是否启用 api 请求日志
	if config.GetBool("log.request_log") {
		apiv1.Use(middleware.WriterLog())
	}
	apiv1.Use(middleware.Cors())

	// 服务路由
	apiv1.GET("/version", v1.Version)

	// domain
	apiv1.GET("/domains", v1.Version)
	apiv1.GET("/domains/:domain/records", v1.Version)

	// records
	apiv1.GET("/records", record.Lists)
	apiv1.POST("/records", record.Create)
}

// HTTPHandler export http gin.engine
func HTTPHandler() *gin.Engine {
	return httpHandler
}

// httpDNS http dns handle
func httpDNS(c *gin.Context) {
	// 转换参数
	var err error
	msg := new(dns.Msg)

	// Parsing the request packet
	switch true {
	case c.ContentType() == "application/dns-message":
		var dnsByte []byte
		if c.Request.Method == http.MethodPost {
			_, _ = c.Request.Body.Read(dnsByte)
		} else {
			if dnsByte, err = base64.RawURLEncoding.DecodeString(c.Query("dns")); err != nil {
				c.String(http.StatusBadRequest, "message parsing exception")
				return
			}
		}
		if err := msg.Unpack(dnsByte); err != nil {
			c.String(http.StatusBadRequest, "message parsing exception")
			return
		}
	case c.ContentType() == "application/dns-json" || c.Query("ct") == "application/dns-json":
		if v, ok := entity.StringToType[c.Query("type")]; !ok {
			c.String(http.StatusBadRequest, "arge `type` exception")
			return
		} else {
			msg.SetQuestion(dns.Fqdn(c.Query("name")), v)
		}
	}

	// Handle dns message
	if err := dnsHandler.Handle(msg); err != nil {
		msg.Rcode = dns.RcodeServerFailure
	}

	// Json response
	if strings.Contains(c.ContentType(), "json") {

	}

	// Binary response
	var resp []byte
	if resp, err = msg.Pack(); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	_, _ = c.Writer.Write(resp)
}
