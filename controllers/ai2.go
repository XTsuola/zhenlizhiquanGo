package controllers

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	my "go_project/config"
//	"io"
//	"net/http"
//	"strings"
//	"time"
//
//	"github.com/gin-gonic/gin"
//)
//
//const (
//	// 火山方舟
//	ARK_API_KEY = "595088bf-db1e-44ac-a389-cd5ffb88ab2c"
//	ARK_EP_ID   = "ep-20260312141043-k2bpn"
//	ARK_API_URL = "https://ark.cn-beijing.volces.com/api/v3/responses"
//)
//
//// =================================================================
//
//var (
//	httpClient = &http.Client{Timeout: 30 * time.Second}
//)
//
//// Responses接口 请求结构体
//type ArkRespReq struct {
//	Model string      `json:"model"`
//	Input []InputItem `json:"input"`
//}
//type InputItem struct {
//	Role    string        `json:"role"`
//	Content []ContentItem `json:"content"`
//}
//type ContentItem struct {
//	Type string `json:"type"`
//	Text string `json:"text"`
//}
//
//type OutputItem struct {
//	Type    string        `json:"type"`
//	Role    string        `json:"role,omitempty"`
//	Content []ContentItem `json:"content,omitempty"`
//}
//
//// Responses接口 返回结构体
//type ArkRespResult struct {
//	CreatedAt int64        `json:"created_at"`
//	ID        string       `json:"id"`
//	Object    string       `json:"object"`
//	Output    []OutputItem `json:"output"` // 重点：[]数组！
//	Status    string       `json:"status"`
//	Error     struct {
//		Message string `json:"message"`
//	} `json:"error"`
//}
//
//// 用户http请求
//type QueryRequest struct {
//	Question string `json:"question" binding:"required"`
//}
//
//// 调用火山方舟 Responses 接口
//func callArkAPI(prompt string) (string, error) {
//	req := ArkRespReq{
//		Model: ARK_EP_ID,
//		Input: []InputItem{
//			{
//				Role: "user",
//				Content: []ContentItem{
//					{
//						Type: "input_text",
//						Text: prompt,
//					},
//				},
//			},
//		},
//	}
//	bodyBytes, err := json.Marshal(req)
//	if err != nil {
//		return "", err
//	}
//
//	httpReq, err := http.NewRequestWithContext(context.Background(), "POST", ARK_API_URL, strings.NewReader(string(bodyBytes)))
//	if err != nil {
//		return "", err
//	}
//	httpReq.Header.Set("Authorization", "Bearer "+ARK_API_KEY)
//	httpReq.Header.Set("Content-Type", "application/json")
//
//	resp, err := httpClient.Do(httpReq)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	respBody, _ := io.ReadAll(resp.Body)
//	var result ArkRespResult
//	if err := json.Unmarshal(respBody, &result); err != nil {
//		return "", fmt.Errorf("解析失败 raw=%s, err=%v", string(respBody), err)
//	}
//
//	if result.Error.Message != "" {
//		return "", fmt.Errorf("方舟API错误：%s", result.Error.Message)
//	}
//	fmt.Println(result.Output[1].Content[0].Text)
//	return result.Output[1].Content[0].Text, nil
//}
//
//// AI生成SQL
//func genSQLByAI(question string) (string, error) {
//	prompt := fmt.Sprintf(`
//	你是专业MySQL SQL生成助手，严格遵守下面所有规则：
//	【数据库表结构】
//	表名：card
//	字段清单：
//	id int 卡牌编号
//	name varchar(255) 卡牌名称
//	zhenyin int 卡牌种族【重点映射规则】
//		帝国 = 1
//		隐秘 = 2
//		禅意 = 3
//		港口 = 4
//		炼狱 = 5
//		蛮石 = 6
//		隐秘 = 7
//	quality int 卡牌品质【重点映射规则】
//		白卡 = 1
//		蓝卡 = 2
//		紫卡 = 3
//		橙卡 = 4
//	cost int 卡牌费用
//	type int 卡牌类型【重点映射规则】
//		部下 = 1
//		法术 = 2
//		传记 = 3
//		符文 = 4
//	img varchar(255) 卡牌图标
//	grade varchar(255) 卡牌评级
//	tag varchar(255) 卡牌标签
//	data varchar(4096) 卡牌具体各个等级详情
//
//
//	强制约束：
//	1. 只允许输出SELECT语句，禁止任何INSERT/UPDATE/DELETE/DROP/ALTER/TRUNCATE等修改、删除语句
//	2. 只返回纯SQL文本，不要markdown、不要注释、不要解释、不要额外文字
//	3. 用户说“帝国”，条件写 zhenyin = 1；
//	   用户说“蛮石”，条件写 zhenyin = 5；
//	   用户说“橙卡”，条件写 quality = 4；
//       用户说“橙卡”，条件写 quality = 4；
//       其他字段以此类推
//	4. 禁止使用不存在的表名、不存在的字段
//	5. 若无明确要求，不要自动添加ORDER BY、LIMIT
//
//	【示例】
//	用户：查询一班学生
//	正确SQL：SELECT * FROM student WHERE class_id = 1;
//
//	用户问题：%s
//	`, question)
//	sqlStr, err := callArkAPI(prompt)
//	if err != nil {
//		return "", err
//	}
//	fmt.Println(111)
//	fmt.Println(sqlStr)
//	fmt.Println(222)
//	sqlStr = strings.TrimSpace(sqlStr)
//	lowerSQL := strings.ToLower(sqlStr)
//
//	// 安全校验
//	if !strings.HasPrefix(lowerSQL, "select") {
//		return "", fmt.Errorf("AI生成非查询语句")
//	}
//	blockList := []string{"insert", "update", "delete", "drop", "alter", "truncate"}
//	for _, keyword := range blockList {
//		if strings.Contains(lowerSQL, keyword) {
//			return "", fmt.Errorf("SQL包含危险操作关键字")
//		}
//	}
//	return sqlStr, nil
//}
//
//// 整理自然语言回答
//func formatAnswer(question, sqlStr string, data []map[string]interface{}) (string, error) {
//	prompt := fmt.Sprintf(`
//用户原始问题：%s
//执行SQL：%s
//数据库查询结果：
//%s
//字段清单：
//	id int 卡牌编号
//	name varchar(255) 卡牌名称
//	zhenyin int 卡牌种族【重点映射规则】
//		帝国 = 1
//		隐秘 = 2
//		禅意 = 3
//		港口 = 4
//		炼狱 = 5
//		蛮石 = 6
//		隐秘 = 7
//	quality int 卡牌品质【重点映射规则】
//		白卡 = 1
//		蓝卡 = 2
//		紫卡 = 3
//		橙卡 = 4
//	cost int 卡牌费用
//	type int 卡牌类型【重点映射规则】
//		部下 = 1
//		法术 = 2
//		传记 = 3
//		符文 = 4
//	img varchar(255) 卡牌图标
//	grade varchar(255) 卡牌评级
//	tag varchar(255) 卡牌标签
//	data varchar(4096) 卡牌具体各个等级详情
//
//请根据上面数据，用通俗易懂中文整理回答用户，简洁清晰。
//如果结果为空，友好提示没有找到对应数据。
//
//`, question, sqlStr, data)
//
//	return callArkAPI(prompt)
//}
//
//func aiQueryHandler(c *gin.Context) {
//	var req QueryRequest
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
//		return
//	}
//
//	sqlText, err := genSQLByAI(req.Question)
//	if err != nil {
//		c.JSON(500, gin.H{"code": 500, "msg": "生成SQL失败:" + err.Error()})
//		return
//	}
//
//	data, err := execQuery(sqlText)
//	//fmt.Print(data, 777)
//	if err != nil {
//		c.JSON(500, gin.H{"code": 500, "msg": fmt.Sprintf("执行SQL失败，sql:%s,err:%v", sqlText, err)})
//		return
//	}
//
//	answer, err := formatAnswer(req.Question, sqlText, data)
//	if err != nil {
//		fmt.Println(err, 555)
//		c.JSON(500, gin.H{"code": 500, "msg": "生成回答失败:" + err.Error()})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"code":   200,
//		"sql":    sqlText,
//		"data":   data,
//		"answer": answer,
//	})
//}
//
//func QueryRaw(sql string) ([]map[string]interface{}, error) {
//	rows, err := my.DB.Raw(sql).Rows()
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	columns, err := rows.Columns()
//	if err != nil {
//		return nil, err
//	}
//	colNum := len(columns)
//	var result []map[string]interface{}
//
//	for rows.Next() {
//		buf := make([][]byte, colNum)
//		ptr := make([]interface{}, colNum)
//		for i := range columns {
//			ptr[i] = &buf[i]
//		}
//		if err := rows.Scan(ptr...); err != nil {
//			return nil, err
//		}
//		row := make(map[string]interface{})
//		for idx, col := range columns {
//			if buf[idx] == nil {
//				row[col] = nil
//			} else {
//				row[col] = string(buf[idx])
//			}
//		}
//		result = append(result, row)
//	}
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//// execQuery 执行AI生成的SQL（gorm版本）
//func execQuery(sqlStr string) ([]map[string]interface{}, error) {
//	fmt.Println(sqlStr, 888)
//	return QueryRaw(sqlStr)
//}
