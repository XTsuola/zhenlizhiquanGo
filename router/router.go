package router

import (
	"net/http"
	"os"
	"strings"
	"time"

	"go_project/config"
	"go_project/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var staticDirs = []string{
	"cardImg", "shenqiImg", "chongwuImg", "yijiImg", "skinImg", "heroImg",
}

func corsOrigins() cors.Config {
	cfg := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	if raw := os.Getenv("CORS_ORIGINS"); raw != "" {
		cfg.AllowOrigins = strings.Split(raw, ",")
		return cfg
	}
	cfg.AllowOriginFunc = func(origin string) bool { return true }
	return cfg
}

// Start 初始化路由并启动 HTTP 服务
func Start(addr string) error {
	r := controllers.R
	r.Use(cors.New(corsOrigins()))

	r.Static("/static", "./dist")
	for _, dir := range staticDirs {
		r.Static("/"+dir, "./static/"+dir)
	}
	r.GET("/ws", wsHandler)

	config.InitDB()
	controllers.InitController()

	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method != http.MethodGet || strings.Contains(c.Request.URL.Path, ".") {
			c.Status(http.StatusNotFound)
			return
		}
		c.File("./dist/index.html")
	})

	return r.Run(addr)
}
