package apps

import (
	"github.com/Ferza17/Products-RESTAPI/utils/env"
	"github.com/Ferza17/Products-RESTAPI/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapURL()
	if err := router.Run(env.GetEnvironmentVariable("PORT")); err != nil {
		logger.Error("Cant connect to server ", err)
		panic(err)
	}
}
