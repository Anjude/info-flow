package bilibili

type PgcRank struct {
	Result struct {
		List []struct {
			Badge string `json:"badge"`
			Cover string `json:"cover"`
			NewEp struct {
				Cover     string `json:"cover"`
				IndexShow string `json:"index_show"`
			} `json:"new_ep"`
			Rating string `json:"rating"`
			Stat   struct {
				Danmaku      int `json:"danmaku"`
				Follow       int `json:"follow"`
				SeriesFollow int `json:"series_follow"`
				View         int `json:"view"`
			} `json:"stat"`
			Title string `json:"title"`
			Url   string `json:"url"`
		} `json:"list"`
		Note string `json:"note"`
	} `json:"result"`
}
type Mblog struct {
	Text     string   `json:"text"`
	PicNum   int      `json:"pic_num"`
	PicIds   []string `json:"pic_ids"`
	PageInfo struct {
		ObjectType string `json:"object_type"`
		PageTitle  string `json:"page_title"`
		PagePic    string `json:"page_pic"`
		PageUrl    string `json:"page_url"`
		MediaInfo  struct {
			Name      string `json:"name"`
			StreamUrl string `json:"stream_url"`
			H5Url     string `json:"h5_url"`
		} `json:"media_info"`
	} `json:"page_info"`
}
