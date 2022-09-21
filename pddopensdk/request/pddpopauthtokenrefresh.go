package request

import "net/url"

// PddPopAuthTokenRefreshRequest 根据refresh_token重新生成token
// https://open.pinduoduo.com/application/document/api?id=pdd.pop.auth.token.refresh
type PddPopAuthTokenRefreshRequest struct {
	Parameters *url.Values
}

func (req *PddPopAuthTokenRefreshRequest) CheckParameters() {
}

// AddParameter 添加请求参数
func (req *PddPopAuthTokenRefreshRequest) AddParameter(key, val string) {
	if req.Parameters == nil {
		req.Parameters = &url.Values{}
	}
	req.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (req *PddPopAuthTokenRefreshRequest) GetApiName() string {
	return "pdd.pop.auth.token.refresh"
}

// GetParameters 返回请求参数
func (req *PddPopAuthTokenRefreshRequest) GetParameters() url.Values {
	return *req.Parameters
}
