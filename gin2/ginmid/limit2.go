package limit2

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"asap/aredis"
)

func CountReject(c *gin.Context) {
	currentTime := time.Now().Unix()
	key := fmt.Sprintf("count:%d", currentTime)
	limitCount := 1
	fmt.Println(key)
	trafficCount, _ := aredis.GetRedis(aredis.BASEREDIS).Incr(key)
	if trafficCount == 1 {
		aredis.GetRedis(aredis.BASEREDIS).Expire(key, 86400)
	}
	if int(trafficCount) > limitCount {
		c.String(http.StatusOK, "reject")
		return
	}
	c.String(http.StatusOK, "ok")
}
