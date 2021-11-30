package zhihu

import "InfoFlow/utils"

const (
	topStoryUrl = "https://www.zhihu.com/api/v3/feed/topstory/hot-list-web?limit=20&desktop=true"
)

// GetTopStory 获取知乎热榜数据
func GetTopStory() (*TopStory, error) {
	res, err := utils.GetReq(topStoryUrl, "https://zhuanlan.zhihu.com/")
	if err != nil {
		return nil, nil
	}
	topStory := TopStory{}
	err = res.ToJSON(&topStory)
	if err != nil {
		return nil, err
	}

	return &topStory, nil
}
