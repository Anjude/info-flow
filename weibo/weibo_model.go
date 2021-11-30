package weibo

type HotBand struct {
	Data struct {
		BandList []struct {
			LabelName string `json:"label_name"`
			Word      string `json:"word"`
			Category  string `json:"category"`
			Num       int    `json:"num"`
			RawHot    int    `json:"raw_hot"`
			Note      string `json:"note"`
			Mblog     `json:"mblog"`
		} `json:"band_list"`
		HotGov struct {
			IsHot int    `json:"is_hot"`
			IsGov int    `json:"is_gov"`
			Word  string `json:"word"`
			Url   string `json:"url"`
			Name  string `json:"name"`
			Mblog `json:"mblog"`
		}
	} `json:"data"`
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
