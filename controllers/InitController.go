package controllers

import (
	"github.com/gin-gonic/gin"
)

var R = gin.Default()

func InitController() {
	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode) // 或 gin.DebugMode
	// 定义路由处理函数
	R.GET("/", func(c *gin.Context) {
		c.String(200, "HTTPS服务运行在端口8002")
	})
	R.GET("/card/list", cardList)
	R.GET("/cardAll/list", cardAllList)
	R.POST("/card/add", cardAdd)
	R.GET("/shenqi/list", shenqiList)
	R.POST("/shenqi/add", shenqiAdd)
	R.GET("/hero/list", heroList)
	R.POST("/hero/add", heroAdd)
	R.GET("/hero/shardList", shardList)
	R.POST("/hero/shardUpdate", shardUpdate)
	R.GET("/frequency/cardsAll", frequencyCardsAll)
	R.GET("/frequency/cardsDetail", frequencyCardsDetail)
	R.POST("/frequency/cardsAdd", frequencyCardsAdd)
	R.POST("/frequency/cardsUpdate", frequencyCardsUpdate)
	R.POST("/frequency/cardsUpdateTemp", frequencyCardsUpdateTemp)
	R.DELETE("/frequency/cardsDelete", frequencyCardsDelete)
	R.POST("/frequency/cardsAddAll", frequencyCardsAddAll)
	R.POST("/frequency/passwordAdd", frequencyPasswordAdd)
	R.GET("/frequency/passwordList", frequencyPasswordList)
	R.DELETE("/frequency/passwordDelete", frequencyPasswordDelete)
	R.GET("/skin/list", skinList)
	R.POST("skin/add", skinAdd)
	R.POST("skin/together", skinTogether)
	R.POST("skin/togetherAll", skinTogetherAll)
	R.GET("/note/list", noteList)
	R.POST("/note/add", noteAdd)
	R.DELETE("/note/delete", noteDelete)
	R.GET("/qingshu/getMap", qingshuGetMap)
	R.GET("/qingshu/reset", qingshuReset)
	R.POST("/qingshu/updateUsername", userNameUpdate)
	R.POST("card/gradeUpdate", cardGradeUpdate)
	R.POST("card/gradeUpdateList", cardGradeUpdateList)
	R.POST("card/tagUpdate", cardTagUpdate)
	R.POST("card/tagUpdateList", cardTagUpdateList)
	R.GET("shijiesai/list", shijiesaiList)
	R.GET("shijiesai/select", shijiesaiSelect)
	R.POST("shijiesai/add", shijiesaiAdd)
	R.POST("shijiesai/update", shijiesaiUpdate)
	R.DELETE("/shijiesai/delete", shijiesaiDelete)
	R.POST("shijiesai/addList", shijiesaiAddList)
	R.GET("log/list", logList)
	R.GET("log/add", logAdd)
}
