package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go_project/models"
	"go_project/qingshu"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	clients = make(map[string]*websocket.Conn)
	mu      sync.Mutex
)

func wsHandler(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		c.String(http.StatusBadRequest, "userId 必传")
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("ws 升级失败:", err)
		return
	}

	mu.Lock()
	if old, ok := clients[userId]; ok && old != conn {
		_ = old.Close()
	}
	clients[userId] = conn
	mu.Unlock()
	log.Printf("用户 %s 已连接", userId)

	defer func() {
		mu.Lock()
		if cur, ok := clients[userId]; ok && cur == conn {
			delete(clients, userId)
		}
		mu.Unlock()
		_ = conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("ws 读取错误:", err)
			break
		}
		var m models.Message
		if err := json.Unmarshal(msg, &m); err != nil {
			log.Println("ws JSON 解析失败:", err)
			continue
		}

		switch m.Type {
		case 1: // 摸牌
			if err := qingshu.Draw(m.UserId); err != nil {
				replyWS(conn, models.ReturnMessage{Code: 500, Msg: err.Error()})
				continue
			}
			broadcast(models.ReturnMessage{Code: 200, Msg: "摸牌成功"})

		case 9: // 重置
			resetId, _ := strconv.Atoi(userId)
			if resetId <= 0 {
				resetId = m.UserId
			}
			if err := qingshu.Reset(resetId); err != nil {
				replyWS(conn, models.ReturnMessage{Code: 500, Msg: err.Error()})
				continue
			}
			broadcast(models.ReturnMessage{Code: 200, Msg: "重置成功"})

		case 2: // 出牌
			updateData, err := qingshu.Play(m.UserId, m.Pai, m.YourPai, m.Index)
			if err != nil {
				replyWS(conn, models.ReturnMessage{Code: 500, Msg: err.Error()})
				continue
			}
			broadcast(models.ReturnMessage{Code: 200, Msg: "出牌成功"})

			extra := models.ReturnMessage{Code: 200, Msg: updateData.Msg}
			switch updateData.Msg {
			case "2", "3", "4", "5", "6", "7", "8":
				extra.Code = 202
			}
			replyWS(conn, extra)
		}
	}
}

func replyWS(conn *websocket.Conn, msg models.ReturnMessage) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Println("ws 序列化失败:", err)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Println("ws 写入错误:", err)
	}
}

func broadcast(msg models.ReturnMessage) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Println("ws 广播序列化失败:", err)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for uid, conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Println("ws 广播写入错误:", err)
			_ = conn.Close()
			delete(clients, uid)
		}
	}
}
