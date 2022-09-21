package request

import "net/url"

// PddPopAuthTokenCreateRequest 用户通过code换获取access_token
// https://open.pinduoduo.com/application/document/api?id=pdd.pop.auth.token.create
type PddPopAuthTokenCreateRequest struct {
	Parameters *url.Values
}

func (req *PddPopAuthTokenCreateRequest) CheckParameters() {
}

// AddParameter 添加请求参数
func (req *PddPopAuthTokenCreateRequest) AddParameter(key, val string) {
	if req.Parameters == nil {
		req.Parameters = &url.Values{}
	}
	req.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (req *PddPopAuthTokenCreateRequest) GetApiName() string {
	return "pdd.pop.auth.token.create"
}

// GetParameters 返回请求参数
func (req *PddPopAuthTokenCreateRequest) GetParameters() url.Values {
	return *req.Parameters
}
