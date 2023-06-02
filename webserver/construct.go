package webserver

import (
	"project_b_webserver/conf"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
	router *gin.Engine
}

func New(webConfig conf.WebServer) {
	w := new(WebServer)

	w.router = gin.Default()

	w.initGet()
	w.initPost()

	ip := webConfig.Host + ":" + strconv.Itoa((webConfig.Port))
	w.router.Run(ip)
}
