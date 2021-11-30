package weibo

import "InfoFlow/utils"

const (
	hotBandUrl = "https://weibo.com/ajax/statuses/hot_band"
)

// GetHotBand 获取知乎热榜数据
func GetHotBand() (*HotBand, error) {
	res, err := utils.GetReq(hotBandUrl, "https://weibo.com/hot/search")
	if err != nil {
		return nil, nil
	}
	hotBand := HotBand{}
	err = res.ToJSON(&hotBand)
	if err != nil {
		return nil, err
	}

	return &hotBand, nil
}
