package pddddkgoodspromotionurlgenerate

import (
	json "encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.goods.promotion.url.generate（生成多多进宝推广链接非授权）
type Response struct {
	response.TopResponse
	GoodsPromotionURLGenerateResponse GoodsPromotionURLGenerateResponse `json:"goods_promotion_url_generate_response"`
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

type GoodsPromotionURLGenerateResponse struct {
	GoodsPromotionURLList []GoodsPromotionURLList `json:"goods_promotion_url_list"`
	RequestID             string                  `json:"request_id"`
}

type GoodsPromotionURLList struct {
	MobileURL       string                  `json:"mobile_url"`
	WeAppInfo       GoodsPromotionWeAppInfo `json:"we_app_info"`
	MobileShortURL  string                  `json:"mobile_short_url"`
	URL             string                  `json:"url"`
	ShortURL        string                  `json:"short_url"`
	SchemaUrl       string                  `json:"schema_url"`
	TzSchemaUrl     string                  `json:"tz_schema_url"`
	QQAppInfo       GoodsPromotionQQAppInfo `json:"qq_app_info"`
	WeixinCode      string                  `json:"weixin_code"`
	WeixinShortLink string                  `json:"weixin_short_link"`
}

type GoodsPromotionWeAppInfo struct {
	WeAppIconURL      string `json:"we_app_icon_url"`
	BannerUrl         string `json:"banner_url"`
	UserName          string `json:"user_name"`
	PagePath          string `json:"page_path"`
	SourceDisplayName string `json:"source_display_name"`
	Title             string `json:"title"`
	AppID             string `json:"app_id"`
	Desc              string `json:"desc"`
}

type GoodsPromotionQQAppInfo struct {
	AppID             string `json:"app_id"`
	BannerUrl         string `json:"banner_url"`
	Desc              string `json:"desc"`
	PagePath          string `json:"page_path"`
	QQAppIconUrl      string `json:"qq_app_icon_url"`
	SourceDisplayName string `json:"source_display_name"`
	Title             string `json:"title"`
	UserName          string `json:"user_name"`
}
