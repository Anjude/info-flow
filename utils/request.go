package utils

import (
	"github.com/imroc/req"
)

// GetReq get 请求函数
func GetReq(url string, referer string) (*req.Resp, error) {
	return req.Get(url, req.Header{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36",
		"Referer":    referer,
	}, nil)
}
