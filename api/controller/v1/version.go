package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BuildTime build in time
var BuildTime string

// BuildVersion build in git last commit id
var BuildVersion string

// Version 获取接口版本
// @Produce  json
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Router /api/v1/version [get]
func Version(c *gin.Context) {
	resp := map[string]interface{}{}
	resp["code"] = 200
	resp["message"] = "success"
	resp["time"] = BuildTime
	resp["version"] = BuildVersion
	c.JSON(http.StatusOK, resp)
}
