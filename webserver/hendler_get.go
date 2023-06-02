package webserver

import "github.com/gin-gonic/gin"

func (w *WebServer) initGet() {
	w.getHomeData()
}

func (w *WebServer) getHomeData() {
	w.router.GET("/api/getHomeData", func(c *gin.Context) {
		c.String(200, "Hello, World!") // 返回"Hello, World!"響應
	})
}
