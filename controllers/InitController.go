package controllers

import (
	"github.com/gin-gonic/gin"
)

func newEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}

var R = newEngine()

func InitController() {
	R.GET("/", func(c *gin.Context) {
		c.String(200, "HTTPS服务运行在端口8002")
	})

	// card
	R.GET("/card/list", cardList)
	R.GET("/cardAll/list", cardAllList)
	R.POST("/card/add", cardAdd)
	R.POST("/card/gradeUpdate", cardGradeUpdate)
	R.POST("/card/gradeUpdateList", cardGradeUpdateList)
	R.POST("/card/tagUpdate", cardTagUpdate)
	R.POST("/card/tagUpdateList", cardTagUpdateList)

	// card diy
	R.GET("/cardDiy/list", cardDiyList)
	R.POST("/cardDiy/add", cardDiyAdd)
	R.POST("/cardDiy/addAll", cardDiyAddAll)
	R.POST("/cardDiy/update", cardDiyUpdate)
	R.POST("/cardDiy/updateTemp", cardDiyUpdateTemp)

	// shenqi / hero / shard
	R.GET("/shenqi/list", shenqiList)
	R.POST("/shenqi/add", shenqiAdd)
	R.GET("/hero/list", heroList)
	R.POST("/hero/add", heroAdd)
	R.GET("/hero/shardList", shardList)
	R.POST("/hero/shardUpdate", shardUpdate)
	R.POST("/hero/shardAdd", shardAdd)
	R.POST("/hero/agentAdd", agentAdd)

	// frequency
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

	// skin / skin diy
	R.GET("/skin/list", skinList)
	R.POST("/skin/add", skinAdd)
	R.GET("/skinDiy/list", skinDiyList)
	R.POST("/skinDiy/add", skinDiyAdd)
	R.POST("/skinDiy/addAll", skinDiyAddAll)
	R.POST("/skinDiy/update", skinDiyUpdate)
	R.POST("/skinDiy/updateTemp", skinDiyUpdateTemp)

	// note / qingshu
	R.GET("/note/list", noteList)
	R.POST("/note/add", noteAdd)
	R.DELETE("/note/delete", noteDelete)
	R.GET("/qingshu/getMap", qingshuGetMap)
	R.GET("/qingshu/reset", qingshuReset)
	R.POST("/qingshu/updateUsername", userNameUpdate)

	// shijiesai
	R.GET("/shijiesai/list", shijiesaiList)
	R.GET("/shijiesai/select", shijiesaiSelect)
	R.POST("/shijiesai/add", shijiesaiAdd)
	R.POST("/shijiesai/update", shijiesaiUpdate)
	R.DELETE("/shijiesai/delete", shijiesaiDelete)
	R.POST("/shijiesai/addList", shijiesaiAddList)

	// log / question / answer
	R.GET("/log/list", logList)
	R.GET("/log/add", logAdd)
	R.GET("/question/list", questionList)
	R.POST("/question/add", questionAdd)
	R.GET("/question/detail", questionDetail)
	R.POST("/question/addAll", questionAddAll)
	R.GET("/answer/list", answerList)
	R.POST("/answer/add", answerAdd)
	R.POST("/answer/addAll", answerAddAll)
	R.GET("/answer/allList", answerAllList)

	// member
	R.GET("/member/list", memberList)
	R.POST("/member/add", memberAdd)
	R.POST("/member/update/:id/", memberUpdate)
	R.POST("/member/addAll", memberAddAll)
	R.DELETE("/member/delete/:id/", memberDelete)
	R.GET("/memberReward/list/", memberRewardList)

	// AI（需配置环境变量 ARK_API_KEY）
	// R.POST("/ai/query", aiQueryHandler)
}
