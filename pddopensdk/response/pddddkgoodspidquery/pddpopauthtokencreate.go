package pddddkgoodspidquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.pop.auth.token.create 用户通过code换获取access_token
type Response struct {
	response.TopResponse
	PIdQueryResponse PIdQueryResponse `json:"p_id_query_response"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = -1
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	}
}

type PIdQueryResponse struct {
	PIdList    []PIdList `json:"p_id_list"`
	TotalCount int64     `json:"total_count"`
}

// PIdList 多多进宝推广位对象
type PIdList struct {
	CreateTime int64  `json:"create_time"` // 推广位生成时间
	MediaId    int64  `json:"media_id"`    // 媒体id
	PidName    string `json:"pid_name"`    // 	推广位名称
	PId        string `json:"p_id"`        // 	推广位ID
	Status     int    `json:"status"`      // 推广位状态：0-正常，1-封禁
}
