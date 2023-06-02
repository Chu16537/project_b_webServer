package webserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (w *WebServer) initPost() {
	w.login()
	w.getMemberData()
}

func (w *WebServer) login() {
	w.router.POST("/api/login", func(c *gin.Context) {
		req := new(LoginReq)

		if err := c.ShouldBindJSON(req); err != nil {
			fmt.Println(err, req)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// 做一些處理，例如保存到數據庫
		if req.Username != "a" || req.Password != "a" {
			c.JSON(401, gin.H{"error": "pass fail"})
			return
		}

		res := new(LoginRes)
		res.Token = "success Token"

		c.JSON(200, res)
	})
}

func (w *WebServer) getMemberData() {
	w.router.POST("/api/getMemberData", func(c *gin.Context) {
		req := new(MemberReq)

		if err := c.ShouldBindJSON(req); err != nil {
			fmt.Println(err, req)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		res := new(MemberRes)
		res.TodayBet = 1
		res.TodayPay = 2
		res.Status = 3
		res.TotalBet = 4
		res.TotalPay = 5

		c.JSON(200, res)
	})
}
