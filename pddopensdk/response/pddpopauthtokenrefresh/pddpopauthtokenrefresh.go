package pddpopauthtokenrefresh

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.pop.auth.token.create 用户通过code换获取access_token
type Response struct {
	response.TopResponse
	PopAuthTokenRefreshResponse PopAuthTokenRefreshResponse `json:"pop_auth_token_refresh_response"`
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

type PopAuthTokenRefreshResponse struct {
	AccessToken           string   `json:"access_token"`
	ExpiresAt             uint64   `json:"expires_at"`               // access_token过期时间点
	ExpiresIn             int      `json:"expires_in"`               // access_token过期时间段，10（表示10秒后过期）
	OwnerId               string   `json:"owner_id"`                 // 商家店铺id
	OwnerName             string   `json:"owner_name"`               // 商家账号名称
	R1ExpiresAt           uint64   `json:"r1_expires_at"`            // r1级别API或字段的访问过期时间点
	R1ExpiresIn           int      `json:"r1_expires_in"`            // r1级别API或字段的访问过期时间； 10（表示10秒后过期）
	R2ExpiresAt           uint64   `json:"r2_expires_at"`            // r2级别API或字段的访问过期时间点
	R2ExpiresIn           int      `json:"r2_expires_in"`            // r2级别API或字段的访问过期时间；10（表示10秒后过期）
	RefreshToken          string   `json:"refresh_token"`            // refresh token，可用来刷新access_token
	RefreshTokenExpiresAt uint64   `json:"refresh_token_expires_at"` // Refresh token过期时间点
	RefreshTokenExpiresIn uint64   `json:"refresh_token_expires_in"` // refresh_token过期时间段，10表示10秒后过期
	Scope                 []string `json:"scope"`                    // 接口列表
	W1ExpiresAt           uint64   `json:"w1_expires_at"`            // w1级别API或字段的访问过期时间点
	W1ExpiresIn           int      `json:"w1_expires_in"`            // w1级别API或字段的访问过期时间； 10（表示10秒后过期）
	W2ExpiresAt           uint64   `json:"w2_expires_at"`            // w2级别API或字段的访问过期时间点
	W2ExpiresIn           int      `json:"w2_expires_in"`            // w2级别API或字段的访问过期时间；10（表示10秒后过期）
}
