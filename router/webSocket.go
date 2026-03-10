package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	my "go_project/config"
	"go_project/models"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// ArrToString 数组转字符串
func ArrToString[T any](arr []T) string {
	if len(arr) == 0 {
		return `[]`
	} else {
		jsonBytes, _ := json.Marshal(arr)
		jsonStr := string(jsonBytes)
		return jsonStr
	}
}

// StringToArr 字符串转数组
func StringToArr[T any](str string) []T {
	var arr []T
	err := json.Unmarshal([]byte(str), &arr)
	if err != nil || len(arr) == 0 {
		fmt.Println(err)
		arr = []T{}
	}
	return arr
}

// 重置游戏
func reset(userId int) {
	var mapObj models.QingshuMapParams
	result := my.DB.Table("qingshu").Where("id = ?", 1).First(&mapObj)
	if result.Error != nil {
		return
	}
	var data models.QingshuMapData
	data.UserData = StringToArr[models.QingshuUserData](mapObj.UserData)
	//cardBaseList := []int{3, 2, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7, 8}
	cardBaseList := []int{1, 1, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 7, 8}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cardBaseList), func(i, j int) {
		cardBaseList[i], cardBaseList[j] = cardBaseList[j], cardBaseList[i]
	})
	var params models.QingshuMapData
	params.CardPile = cardBaseList[2:13]
	params.DisPile = cardBaseList[13:16]
	var user1 models.QingshuUserData
	user1.ID = 1
	user1.UserName = data.UserData[0].UserName
	user1.HandCards = cardBaseList[0:1]
	user1.DisCards = make([]int, 0)
	user1.Status = 1
	var user2 models.QingshuUserData
	user2.ID = 2
	user2.UserName = data.UserData[1].UserName
	user2.HandCards = cardBaseList[1:2]
	user2.DisCards = make([]int, 0)
	user2.Status = 1
	params.UserData = append(params.UserData, user1)
	params.UserData = append(params.UserData, user2)
	params.Round = userId
	params.Status = 0
	params.Msg = ""
	result2 := my.DB.Table("qingshu").Where("id = ?", 1).Updates(map[string]interface{}{
		"round":    params.Round,
		"status":   params.Status,
		"msg":      params.Msg,
		"cardPile": ArrToString(params.CardPile),
		"disPile":  ArrToString(params.DisPile),
		"userData": ArrToString(params.UserData),
	})
	if result2.Error != nil {
		return
	}
}

// 摸牌
func moPai(userId int) {
	var obj models.QingshuMapParams
	result := my.DB.Table("qingshu").First(&obj, 1) // SELECT * FROM users WHERE id = ? LIMIT 1
	if result.Error != nil {
		fmt.Println("数据不存在:", result.Error)
		return
	}
	var params models.QingshuMapData
	params.CardPile = StringToArr[int](obj.CardPile)[1:]
	params.DisPile = StringToArr[int](obj.DisPile)
	params.Round = obj.Round
	params.Status = 1
	params.Msg = ""
	user1 := StringToArr[models.QingshuUserData](obj.UserData)[0]
	user2 := StringToArr[models.QingshuUserData](obj.UserData)[1]
	if userId == 1 {
		user1.HandCards = append(user1.HandCards, StringToArr[int](obj.CardPile)[0])
		user1.Status = 1
	} else {
		user2.HandCards = append(user2.HandCards, StringToArr[int](obj.CardPile)[0])
		user2.Status = 1
	}
	params.UserData = make([]models.QingshuUserData, 0)
	params.UserData = append(params.UserData, user1)
	params.UserData = append(params.UserData, user2)

	result2 := my.DB.Table("qingshu").Where("id = ?", 1).Updates(map[string]interface{}{
		"round":    params.Round,
		"status":   params.Status,
		"msg":      params.Msg,
		"cardPile": ArrToString(params.CardPile),
		"disPile":  ArrToString(params.DisPile),
		"userData": ArrToString(params.UserData),
	})
	if result2.Error != nil {
		fmt.Println("操作失败:", result.Error)
		return
	}
}

// 出牌
func chuPai(myId int, pai int, obj models.QingshuMapData, youPari int, index int) models.QingshuMapData {
	fmt.Println(myId, pai, youPari, index, "kkk")
	var params models.QingshuMapData
	params.CardPile = obj.CardPile
	params.DisPile = obj.DisPile
	params.Round = obj.Round + 1
	var user1 models.QingshuUserData
	var user2 models.QingshuUserData
	if myId == 1 {
		user1 = obj.UserData[0]
		user2 = obj.UserData[1]
	} else {
		user2 = obj.UserData[0]
		user1 = obj.UserData[1]
	}
	if user2.Status != 3 {
		if pai == 1 {
			if youPari == user2.HandCards[0] {
				user2.Status = 2
				user2.DisCards = append(user2.DisCards, user2.HandCards[0])
				user2.HandCards = make([]int, 0)
				params.Status = 2
			}
		} else if pai == 2 {
			params.Msg = strconv.Itoa(user2.HandCards[0])
		} else if pai == 3 {
			myCard := 0
			if index == 0 {
				myCard = user1.HandCards[1]
			} else {
				myCard = user1.HandCards[0]
			}
			if myCard > user2.HandCards[0] {
				user2.Status = 2
				params.Status = 2
			} else if myCard < user2.HandCards[0] {
				user1.Status = 2
				params.Status = 2
			}
		} else if pai == 5 {
			user2.DisCards = append(user2.DisCards, user2.HandCards[0])
			if user2.HandCards[0] == 8 {
				user2.Status = 2
				user2.HandCards = make([]int, 0)
				params.Status = 2
			} else {
				if len(params.CardPile) > 0 {
					user2.HandCards[0] = params.CardPile[0]
					params.CardPile = params.CardPile[1:]
				} else {
					user2.HandCards[0] = params.DisPile[0]
					params.DisPile = params.DisPile[1:]
				}
			}
		} else if pai == 6 {
			user2Card := user2.HandCards[0]
			if index == 0 {
				user2.HandCards[0] = user1.HandCards[1]
				user1.HandCards[1] = user2Card
			} else {
				user2.HandCards[0] = user1.HandCards[0]
				user1.HandCards[0] = user2Card
			}
		}
	}
	if pai == 4 {
		user1.Status = 3
	} else if pai == 8 {
		user1.Status = 2
		params.Status = 2
	}
	if index == 0 {
		user1.HandCards = user1.HandCards[1:]
	} else {
		user1.HandCards = user1.HandCards[0:1]
	}
	user1.DisCards = append(user1.DisCards, pai)
	if len(params.CardPile) == 0 {
		if user1.HandCards[0] > user2.HandCards[0] {
			user2.Status = 2
			params.Status = 2
		} else if user1.HandCards[0] < user2.HandCards[0] {
			user1.Status = 2
			params.Status = 2
		} else {
			user1.HandCards[0] = params.DisPile[0]
			user2.HandCards[0] = params.DisPile[1]
			params.DisPile = params.DisPile[2:]
			if user1.HandCards[0] > user2.HandCards[0] {
				user2.Status = 2
				params.Status = 2
			} else {
				user1.Status = 2
				params.Status = 2
			}
		}
	}
	params.UserData = make([]models.QingshuUserData, 0)
	if myId == 1 {
		params.UserData = append(params.UserData, user1)
		params.UserData = append(params.UserData, user2)
	} else {
		params.UserData = append(params.UserData, user2)
		params.UserData = append(params.UserData, user1)
	}
	if params.Status == 2 {
		params.Msg = "游戏结束"
	}
	return params
}

// 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域，生产环境要做域名校验
	},
}

// 全局用户连接表
var clients = make(map[string]*websocket.Conn)
var mu sync.Mutex

// WebSocket 连接处理
func wsHandler(c *gin.Context) {
	userId := c.Query("userId") // 通过 ?userId=1 来区分用户
	if userId == "" {
		c.String(http.StatusBadRequest, "userId 必传")
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("升级失败:", err)
		return
	}
	defer conn.Close()
	mu.Lock()
	clients[userId] = conn
	mu.Unlock()
	fmt.Printf("用户 %s 已连接\n", userId)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("读取错误:", err)
			break
		}
		var m models.Message
		if err2 := json.Unmarshal(msg, &m); err2 != nil {
			fmt.Println("JSON 解析失败:", err2)
			continue
		}
		var returnMessage models.ReturnMessage
		returnMessage.Code = 200
		if m.Type == 1 {
			moPai(m.UserId)
			returnMessage.Msg = "摸牌成功"
			broadcast(userId, returnMessage)
		}
		if m.Type == 9 {
			resetId, _ := strconv.Atoi(userId)
			reset(resetId)
			returnMessage.Msg = "重置成功"
			broadcast(userId, returnMessage)
		}
		if m.Type == 2 {
			var params models.QingshuMapParams
			result := my.DB.Table("qingshu").First(&params, 1) // SELECT * FROM users WHERE id = ? LIMIT 1
			if result.Error != nil {
				fmt.Println("数据不存在:", result.Error)
				return
			}
			var obj models.QingshuMapData
			obj.Round = params.Round
			obj.Status = params.Status
			obj.Msg = params.Msg
			obj.CardPile = StringToArr[int](params.CardPile)
			obj.DisPile = StringToArr[int](params.DisPile)
			obj.UserData = StringToArr[models.QingshuUserData](params.UserData)
			updateData := chuPai(m.UserId, m.Pai, obj, m.YourPai, m.Index)
			result2 := my.DB.Table("qingshu").Where("id = ?", 1).Updates(map[string]interface{}{
				"round":    updateData.Round,
				"status":   updateData.Status,
				"msg":      updateData.Msg,
				"cardPile": ArrToString(updateData.CardPile),
				"disPile":  ArrToString(updateData.DisPile),
				"userData": ArrToString(updateData.UserData),
			})
			if result2.Error != nil {
				fmt.Println("操作失败:", result.Error)
				return
			}
			returnMessage.Msg = "出牌成功"
			broadcast(userId, returnMessage)
			returnMessage.Msg = updateData.Msg
			validMsgs := map[string]bool{
				"2": true,
				"3": true,
				"4": true,
				"5": true,
				"6": true,
				"7": true,
				"8": true,
			}
			if validMsgs[returnMessage.Msg] {
				returnMessage.Code = 202
			}
			if _, ok := clients[userId]; ok {
				data, _ := json.Marshal(returnMessage)
				conn.WriteMessage(websocket.TextMessage, data)
			} else {
				fmt.Printf("用户 %s 不在线\n", userId)
			}
		}
	}
}

func broadcast(fromUser string, msgData models.ReturnMessage) {
	mu.Lock()
	defer mu.Unlock()
	data, err := json.Marshal(msgData)
	if err != nil {
		panic(err)
	}
	str := string(data)
	for uid, conn := range clients {
		// 给所有人发，包括自己
		err2 := conn.WriteMessage(websocket.TextMessage, []byte(str))
		if err2 != nil {
			fmt.Println("写入错误:", err2)
			conn.Close()
			delete(clients, uid)
		}
	}
}
