package limit1

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RandomReject(c *gin.Context) {
	// 检测失败返回reject 502

	refuseRate := 200
	if refuseRate != 0 {
		temp := rand.Intn(1000)
		if temp <= refuseRate {
			c.String(http.StatusBadGateway, "reject")
			return
		}

	}
	c.String(http.StatusOK, "ok")
}
