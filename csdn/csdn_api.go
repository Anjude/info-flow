package csdn

import "InfoFlow/utils"

const (
	hotRankUrl = "https://blog.csdn.net/phoenix/web/blog/hot-rank?page=0&pageSize=25"
	searchUrl  = "https://so.csdn.net/api/v3/search?q="
)

// SearchCsdn 搜索 csdn 内容
func SearchCsdn(keyword string) (*SearchRes, error) {
	res, err := utils.GetReq(searchUrl+keyword, "https://so.csdn.net/so/search")
	if err != nil {
		return nil, nil
	}
	searchRes := SearchRes{}
	err = res.ToJSON(&searchRes)
	if err != nil {
		return nil, err
	}

	return &searchRes, nil
}

// GetHotRank 获取知乎热榜数据
func GetHotRank() (*HotRank, error) {
	res, err := utils.GetReq(hotRankUrl, "https://www.csdn.net/")
	if err != nil {
		return nil, nil
	}
	hotRank := HotRank{}
	err = res.ToJSON(&hotRank)
	if err != nil {
		return nil, err
	}

	return &hotRank, nil
}
