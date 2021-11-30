package zhihu

// TopStory 热榜数据结构体
type TopStory struct {
	Data []struct {
		Type         string `json:"type"`
		StyleType    string `json:"style_type"`
		Id           string `json:"id"`
		CardId       string `json:"card_id"`
		FeedSpecific struct {
			AnswerCount int `json:"answer_count"`
		} `json:"feed_specific"`
		Target struct {
			TitleArea struct {
				Text string `json:"text"`
			} `json:"title_area"`
			ExcerptArea struct {
				Text string `json:"text"`
			} `json:"excerpt_area"`
			ImageArea struct {
				Url string `json:"url"`
			} `json:"image_area"`

			MetricsArea struct {
				Text string `json:"text"`
			} `json:"metrics_area"`

			LabelArea struct {
				Type        string `json:"type"`
				Trend       int    `json:"trend"`
				NightColor  string `json:"night_color"`
				NormalColor string `json:"normal_color"`
			} `json:"label_area"`
			Link struct {
				Url string `json:"url"`
			} `json:"link"`
		}
		AttachedInfo string `json:"attached_info"`
	} `json:"data"`
}
