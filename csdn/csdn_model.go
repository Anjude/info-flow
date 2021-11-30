package csdn

// SearchRes csdn 搜索内容结构
type SearchRes struct {
	ResultVos []struct {
		Description       string        `json:"description"`
		Language          string        `json:"language"`
		Pic               string        `json:"pic"`
		Type              string        `json:"type"`
		Title             string        `json:"title"`
		Body              string        `json:"body"`
		View              string        `json:"view"`
		Nickname          string        `json:"nickname"`
		Digest            string        `json:"digest"`
		Author            string        `json:"author"`
		Url               string        `json:"url"`
		Tags              []interface{} `json:"tags"`
		SearchTag         []string      `json:"search_tag,omitempty"`
		CreateTimeStr     string        `json:"create_time_str"`
		SoType            string        `json:"so_type"`
		//Comment           string        `json:"comment"`
		RequestId         string        `json:"request_id"`
		UrlLocation       string        `json:"url_location"`
		EventView         bool          `json:"eventView"`
		UtmSource         string        `json:"utm_source"`
		PicList           []string      `json:"pic_list,omitempty"`
		Sourcesize        string        `json:"sourcesize,omitempty"`
		LecturerLevel     string        `json:"lecturer_level,omitempty"`
		LecturerLevelName string        `json:"lecturer_level_name,omitempty"`
		LecturerAvatar    string        `json:"lecturer_avatar,omitempty"`
		LecturerUrl       string        `json:"lecturer_url,omitempty"`
		CourseCount       string        `json:"course_count,omitempty"`
	} `json:"result_vos"`
}

// HotRank csdn 热榜数据结构
type HotRank struct {
	Data []struct {
		HotRankScore      string      `json:"hotRankScore"`
		PcHotRankScore    string      `json:"pcHotRankScore"`
		LoginUserIsFollow bool        `json:"loginUserIsFollow"`
		NickName          string      `json:"nickName"`
		AvatarUrl         string      `json:"avatarUrl"`
		UserName          string      `json:"userName"`
		ArticleTitle      string      `json:"articleTitle"`
		ArticleDetailUrl  string      `json:"articleDetailUrl"`
		CommentCount      string      `json:"commentCount"`
		FavorCount        string      `json:"favorCount"`
		ViewCount         string      `json:"viewCount"`
		HotComment        interface{} `json:"hotComment"`
		PicList           []string    `json:"picList"`
	} `json:"data"`
}
