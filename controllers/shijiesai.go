package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"strconv"
	"time"
)

func shijiesaiList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNo, _ := strconv.Atoi(c.Query("pageNo"))
	gameType, _ := strconv.Atoi(c.Query("gameType"))
	offset := (pageNo - 1) * pageSize
	var list []models.ShijiesaiAll
	var total int64
	my.DB.Table("shijiesai").Where("no > ? AND no < ?", gameType*10000, (gameType+1)*10000).Count(&total)
	result := my.DB.Table("shijiesai").Where("no > ? AND no < ?", gameType*10000, (gameType+1)*10000).Order("no asc").
		Limit(pageSize).
		Offset(offset).
		Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	var data []models.ShijiesaiList
	for _, item := range list {
		var obj models.ShijiesaiList
		obj.ID = item.ID
		obj.No = item.No
		info := StringToArr[models.XuanshouInfo](item.Info)
		obj.AInfo = info[0]
		obj.BInfo = info[1]
		obj.ShengfuList = StringToArr[int](item.ShengfuList)
		data = append(data, obj)
	}
	SearchByPage[models.ShijiesaiList]("查询成功", c, data, total)
}

func shijiesaiAdd(c *gin.Context) {
	var params models.ShijiesaiAdd
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.ShijiesaiBase
	data.No = params.No
	var infoData []models.XuanshouInfo
	infoData = append(infoData, params.AInfo)
	infoData = append(infoData, params.BInfo)
	data.Info = ArrToString(infoData)
	data.ShengfuList = ArrToString(params.ShengfuList)
	var shijiesaiObj []models.ShijiesaiAll
	result := my.DB.Table("shijiesai").Where("no = ?", data.No).Find(&shijiesaiObj)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	if len(shijiesaiObj) != 0 {
		MyErr("该场次已经存在", c)
		return
	}
	result2 := my.DB.Table("shijiesai").Create(&data)
	if result2.Error != nil {
		MyErr(result2.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func shijiesaiUpdate(c *gin.Context) {
	var params models.ShijiesaiList
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var infoData []models.XuanshouInfo
	infoData = append(infoData, params.AInfo)
	infoData = append(infoData, params.BInfo)
	result := my.DB.Table("shijiesai").Where("id = ?", params.ID).Updates(map[string]interface{}{
		"no":          params.No,
		"info":        ArrToString(infoData),
		"shengfuList": ArrToString(params.ShengfuList),
	})
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}

func shijiesaiDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	result := my.DB.Table("shijiesai").Where("id = ?", id).Delete(nil)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "删除成功")
}

// 批量新增
func shijiesaiAddList(c *gin.Context) {
	var params models.ShijiesaiAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.ShijiesaiBase
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.No = item.No
		var infoData []models.XuanshouInfo
		infoData = append(infoData, item.AInfo)
		infoData = append(infoData, item.BInfo)
		data.Info = ArrToString(infoData)
		data.ShengfuList = ArrToString(item.ShengfuList)
		result := my.DB.Table("shijiesai").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func shijiesaiSelect(c *gin.Context) {
	gameType, _ := strconv.Atoi(c.Query("gameType"))
	var list []models.ShijiesaiAll
	result := my.DB.Table("shijiesai").Where("no > ? AND no < ?", gameType*10000, (gameType+1)*10000).Order("no asc").Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	var data []models.ShijiesaiList
	for _, item := range list {
		var obj models.ShijiesaiList
		obj.ID = item.ID
		obj.No = item.No
		info := StringToArr[models.XuanshouInfo](item.Info)
		obj.AInfo = info[0]
		obj.BInfo = info[1]
		data = append(data, obj)
	}
	SearchList[models.ShijiesaiList]("查询成功", c, data)
}
