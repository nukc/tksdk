package request

import "net/url"

// PddDdkGoodsPidQueryRequest 查询已经生成的推广位信息
// https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.pid.query
type PddDdkGoodsPidQueryRequest struct {
	Parameters *url.Values
}

func (req *PddDdkGoodsPidQueryRequest) CheckParameters() {
}

// AddParameter 添加请求参数
func (req *PddDdkGoodsPidQueryRequest) AddParameter(key, val string) {
	if req.Parameters == nil {
		req.Parameters = &url.Values{}
	}
	req.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (req *PddDdkGoodsPidQueryRequest) GetApiName() string {
	return "pdd.ddk.goods.pid.query"
}

// GetParameters 返回请求参数
func (req *PddDdkGoodsPidQueryRequest) GetParameters() url.Values {
	return *req.Parameters
}
