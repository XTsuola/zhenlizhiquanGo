package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	my "go_project/config"

	"github.com/gin-gonic/gin"
)

const (
	defaultArkEPID = "ep-20260312141043-k2bpn"
	defaultArkURL  = "https://ark.cn-beijing.volces.com/api/v3/responses"
)

var httpClient = &http.Client{Timeout: 200 * time.Second}

type ArkRespReq struct {
	Model string      `json:"model"`
	Input []InputItem `json:"input"`
}

type InputItem struct {
	Role    string        `json:"role"`
	Content []ContentItem `json:"content"`
}

type ContentItem struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type OutputItem struct {
	Type    string        `json:"type"`
	Role    string        `json:"role,omitempty"`
	Content []ContentItem `json:"content,omitempty"`
}

type ArkRespResult struct {
	CreatedAt int64        `json:"created_at"`
	ID        string       `json:"id"`
	Object    string       `json:"object"`
	Output    []OutputItem `json:"output"`
	Status    string       `json:"status"`
	Error     struct {
		Message string `json:"message"`
	} `json:"error"`
}

type QueryRequest struct {
	Question string `json:"question" binding:"required"`
}

func arkConfig() (apiKey, epID, apiURL string, err error) {
	apiKey = os.Getenv("ARK_API_KEY")
	if apiKey == "" {
		return "", "", "", fmt.Errorf("未配置环境变量 ARK_API_KEY")
	}
	epID = os.Getenv("ARK_EP_ID")
	if epID == "" {
		epID = defaultArkEPID
	}
	apiURL = os.Getenv("ARK_API_URL")
	if apiURL == "" {
		apiURL = defaultArkURL
	}
	return apiKey, epID, apiURL, nil
}

func extractArkText(result ArkRespResult) (string, error) {
	for i := len(result.Output) - 1; i >= 0; i-- {
		item := result.Output[i]
		if item.Type != "message" && item.Role != "assistant" {
			continue
		}
		for _, content := range item.Content {
			if content.Text != "" {
				return content.Text, nil
			}
		}
	}
	for _, item := range result.Output {
		for _, content := range item.Content {
			if content.Text != "" {
				return content.Text, nil
			}
		}
	}
	return "", fmt.Errorf("方舟返回内容为空")
}

func callArkAPI(ctx context.Context, prompt string) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	apiKey, epID, apiURL, err := arkConfig()
	if err != nil {
		return "", err
	}

	req := ArkRespReq{
		Model: epID,
		Input: []InputItem{{
			Role: "user",
			Content: []ContentItem{{
				Type: "input_text",
				Text: prompt,
			}},
		}},
	}
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("方舟HTTP错误 status=%d body=%s", resp.StatusCode, string(respBody))
	}

	var result ArkRespResult
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("解析失败 raw=%s, err=%v", string(respBody), err)
	}
	if result.Error.Message != "" {
		return "", fmt.Errorf("方舟API错误：%s", result.Error.Message)
	}
	return extractArkText(result)
}

func genSQLByAI(ctx context.Context, question string) (string, error) {
	prompt := fmt.Sprintf(`
你是专业MySQL SQL生成助手，严格遵守下面所有规则：
【数据库表结构】
表名：shenqi
字段清单：
id int 神器编号
name varchar(255) 神器名称
zhenyin int 卡牌种族【重点映射规则】
	帝国 = 1
	隐秘 = 2
	禅意 = 3
	港口 = 4
	炼狱 = 5
	蛮石 = 6
	中立 = 7
quality int 神器品质【重点映射规则】
	蓝色 = 1
	紫色 = 2
	橙色 = 3
type int 神器类型【重点映射规则】
	武器 = 1
	宝物 = 2

强制约束：
1. 只允许输出SELECT语句，禁止任何INSERT/UPDATE/DELETE/DROP/ALTER/TRUNCATE等修改、删除语句
2. 只返回纯SQL文本，不要markdown、不要注释、不要解释、不要额外文字
3. 用户说“帝国”，条件写 zhenyin = 1；
   用户说“蛮石”，条件写 zhenyin = 6；
   用户说“橙色”，条件写 quality = 3；
   用户说“类型”，条件写 type = 1或者2；
   其他字段以此类推
4. 禁止使用不存在的表名、不存在的字段
5. 若无明确要求，不要自动添加ORDER BY、LIMIT
6. SQL 只能查询 shenqi 表

【示例】
用户：查询帝国神器
正确SQL：SELECT * FROM shenqi WHERE zhenyin = 1;

用户问题：%s
`, question)

	sqlStr, err := callArkAPI(ctx, prompt)
	if err != nil {
		return "", err
	}
	sqlStr = strings.TrimSpace(sqlStr)
	sqlStr = strings.TrimPrefix(sqlStr, "```sql")
	sqlStr = strings.TrimPrefix(sqlStr, "```")
	sqlStr = strings.TrimSuffix(sqlStr, "```")
	sqlStr = strings.TrimSpace(sqlStr)

	lowerSQL := strings.ToLower(sqlStr)
	if !strings.HasPrefix(lowerSQL, "select") {
		return "", fmt.Errorf("AI生成非查询语句")
	}
	if !strings.Contains(lowerSQL, "shenqi") {
		return "", fmt.Errorf("SQL仅允许查询 shenqi 表")
	}
	trimmed := strings.TrimSuffix(strings.TrimSpace(lowerSQL), ";")
	if strings.Contains(trimmed, ";") {
		return "", fmt.Errorf("SQL包含多条语句")
	}
	blockList := []string{" insert", " update", " delete", " drop", " alter", " truncate", "--", "/*"}
	padded := " " + trimmed
	for _, keyword := range blockList {
		if strings.Contains(padded, keyword) {
			return "", fmt.Errorf("SQL包含危险操作关键字")
		}
	}
	return sqlStr, nil
}

func formatAnswer(ctx context.Context, question, sqlStr string, data []map[string]interface{}) (string, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	prompt := fmt.Sprintf(`
用户原始问题：%s
生成的sql语句：%s
数据库查询结果：%s

请根据上面数据生成markdown返回
如果结果为空，友好提示没有找到对应数据。
`, question, sqlStr, string(dataJSON))
	return callArkAPI(ctx, prompt)
}

func aiQueryHandler(c *gin.Context) {
	var req QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误"})
		return
	}

	ctx := c.Request.Context()
	sqlText, err := genSQLByAI(ctx, req.Question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "生成SQL失败:" + err.Error()})
		return
	}

	data, err := execQuery(sqlText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": fmt.Sprintf("执行SQL失败，sql:%s,err:%v", sqlText, err)})
		return
	}

	answer, err := formatAnswer(ctx, req.Question, sqlText, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "生成回答失败:" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"sql":    sqlText,
		"answer": answer,
	})
}

func QueryRaw(sql string) ([]map[string]interface{}, error) {
	rows, err := my.DB.Raw(sql).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	colNum := len(columns)
	result := make([]map[string]interface{}, 0)

	for rows.Next() {
		buf := make([][]byte, colNum)
		ptr := make([]interface{}, colNum)
		for i := range columns {
			ptr[i] = &buf[i]
		}
		if err := rows.Scan(ptr...); err != nil {
			return nil, err
		}
		row := make(map[string]interface{}, colNum)
		for idx, col := range columns {
			if buf[idx] == nil {
				row[col] = nil
			} else {
				row[col] = string(buf[idx])
			}
		}
		result = append(result, row)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func execQuery(sqlStr string) ([]map[string]interface{}, error) {
	return QueryRaw(sqlStr)
}
