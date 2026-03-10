package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_project/config"
	"go_project/controllers"
	"time"
)

func InitRouter() {
	controllers.R.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有域，生产环境建议改为具体域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	controllers.R.Static("/static", "./dist")
	controllers.R.Static("/cardImg", "./static/cardImg")
	controllers.R.Static("/shenqiImg", "./static/shenqiImg")
	controllers.R.Static("/chongwuImg", "./static/chongwuImg")
	controllers.R.Static("/yijiImg", "./static/yijiImg")
	controllers.R.Static("/skinImg", "./static/skinImg")
	controllers.R.Static("/heroImg", "./static/heroImg")
	controllers.R.NoRoute(func(c *gin.Context) { // 兼容 Vue/React 单页应用路由
		c.File("./dist/index.html")
	})
	controllers.R.GET("/ws", wsHandler) // WebSocket 连接地址
	config.InitDB()
	controllers.InitController()
	err := controllers.R.Run(":8002")
	if err != nil {
		return
	}

	//StartHTTPS()
}
