package main

import (
	"InfoFlow/bilibili"
	"InfoFlow/csdn"
	"InfoFlow/weibo"
	"InfoFlow/zhihu"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"strconv"
)

func main() {
	var showAboutBoxAction *walk.Action

	// 获取软件窗口，初始化数据
	var mw = &InfoFlow{infoListModel: getZhihuList(nil)}

	if err := (MainWindow{
		Title:    "InfoFlow",
		MinSize:  Size{Width: 600, Height: 400},
		Icon:     "img/InfoFlowLogo.ico",
		Layout:   VBox{MarginsZero: true},
		AssignTo: &mw.MainWindow,
		// 顶部菜单
		MenuItems: []MenuItem{
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: mw.showAboutBoxAction,
					}, Separator{},
					Action{
						Text:        "Exit",
						OnTriggered: func() { _ = mw.Close() },
					},
				},
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{Action: &showAboutBoxAction},
		},
		// 主元素
		Children: []Widget{
			// 信息来源筛选
			Composite{
				MaxSize: Size{Height: 50},
				Layout:  HBox{},
				Children: []Widget{
					Label{Text: "来源: "},
					RadioButtonGroup{
						DataMember: "Src",
						Buttons: []RadioButton{
							{
								Name:     "zhihu",
								Text:     "知乎",
								Value:    "0",
								AssignTo: &mw.zhihu,
								OnMouseUp: func(x, y int, button walk.MouseButton) {
									go getZhihuList(mw)
								},
							},
							{
								Name:     "weibo",
								Text:     "微博",
								Value:    "1",
								AssignTo: &mw.weibo,
								OnMouseUp: func(x, y int, button walk.MouseButton) {
									go mw.getWeiboList()
								},
							},
							{
								Name:     "bilibili",
								Text:     "BiliBili番剧",
								Value:    "2",
								AssignTo: &mw.bilibili,
								OnMouseUp: func(x, y int, button walk.MouseButton) {
									go mw.getBilibiliList()
								},
							},
							{
								Name:     "csdn",
								Text:     "CSDN",
								Value:    "3",
								AssignTo: &mw.csdn,
								OnMouseUp: func(x, y int, button walk.MouseButton) {
									go mw.getCsdnList()
								},
							},
						},
					},
					Label{Text: "关键词: "},
					LineEdit{
						AssignTo: &mw.keywords,
						Text:     "互联网",
					},
					PushButton{
						AssignTo: &mw.query,
						Text:     "搜索",
					},
				},
			},
			// 内容区域
			HSplitter{
				Children: []Widget{
					VSplitter{
						StretchFactor: 10,
						Children: []Widget{
							ListBox{
								AssignTo:              &mw.lb,
								Model:                 mw.infoListModel,
								OnCurrentIndexChanged: mw.lbCurrentIndexChanged,
								OnItemActivated:       mw.lbItemActivated,
								Font:                  Font{PointSize: 12},
							},
							TextEdit{
								AssignTo: &mw.infoPanel,
								Font:     Font{PointSize: 12},
								ReadOnly: true,
							},
						},
					},
					WebView{
						//MinSize:  Size{1000, 0},
						AssignTo: &mw.wv,
					},
				},
			},
			// 底部链接
			LinkLabel{
				Text: `项目发布地址：<a id="项目发布地址" href="https://gitee.com/anjude/info-flow/releases">点击复制地址</a> or 联系我： <a id="公众号" href="豆小匠的编程日常">豆小匠的编程日常</a>.`,
				OnLinkActivated: func(link *walk.LinkLabelLink) {
					_ = walk.Clipboard().SetText(link.URL())
					b := link.Id() == "公众号"
					if b {
						walk.MsgBox(mw, "复制成功", "微信搜索：豆小匠的编程日常", walk.MsgBoxIconQuestion)
					} else {
						walk.MsgBox(mw, "复制成功", "链接复制成功，到浏览器访问即可！", walk.MsgBoxUserIcon)
					}
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}
	mw.query.Clicked().Attach(func() {
		mw.GetList()
	})

	//初始化
	//_ = mw.wv.SetURL("https://www.ithome.com/")
	_ = mw.wv.SetURL("https://gitee.com/anjude/public-resource/raw/md-img/20211118174102.jpeg")

	mw.zhihu.SetChecked(true)

	mw.Run()
}

// GetList 按选择和关键词搜索内容
func (mw *InfoFlow) GetList() {
	if mw.csdn.Checked() == true {
		go mw.searchCsdnByKeyword(mw.keywords)
	} else {
		walk.MsgBox(mw, "不支持", "仅支持 CSDN 搜索", walk.MsgBoxUserIcon)
	}
}

// 根据 keywords 搜索
func (mw *InfoFlow) searchCsdnByKeyword(keywords *walk.LineEdit) {
	if !mw.csdn.Checked() {
		return
	}
	searchRes, err := csdn.SearchCsdn(keywords.Text())
	if err != nil {
		return
	}
	m := &InfoListModel{items: make([]Item, len(searchRes.ResultVos))}

	for i, v := range searchRes.ResultVos {
		imgUrl := ""
		if len(v.PicList) != 0 {
			imgUrl = v.PicList[0]
		}
		m.items[i] = Item{strconv.Itoa(i+1) + "、" + v.Title +
			"\n", v.Body, v.Url,
			imgUrl}
	}
	model := &InfoListModel{items: m.items}
	_ = mw.lb.SetModel(model)
	mw.infoListModel = m
}

// 获取 csdn 热榜
func (mw *InfoFlow) getCsdnList() {
	if !mw.csdn.Checked() {
		return
	}
	hotRank, err := csdn.GetHotRank()
	if err != nil {
		return
	}
	m := &InfoListModel{items: make([]Item, len(hotRank.Data))}

	for i, v := range hotRank.Data {
		imgUrl := ""
		if len(v.PicList) != 0 {
			imgUrl = v.PicList[0]
		}
		m.items[i] = Item{strconv.Itoa(i+1) + "、" + v.ArticleTitle +
			"\n", "热度" + v.PcHotRankScore, v.ArticleDetailUrl,
			imgUrl}
	}
	model := &InfoListModel{items: m.items}
	_ = mw.lb.SetModel(model)
	mw.infoListModel = m
}

// 获取 bilibili 番剧热榜
func (mw *InfoFlow) getBilibiliList() {
	if !mw.bilibili.Checked() {
		return
	}
	pgcRank, err := bilibili.GetPgcRank()
	if err != nil {
		return
	}
	m := &InfoListModel{items: make([]Item, len(pgcRank.Result.List))}

	for i, v := range pgcRank.Result.List {
		m.items[i] = Item{strconv.Itoa(i+1) + "、" + v.Title +
			"\n", v.Badge + "," +
			v.Rating + "," + v.NewEp.IndexShow, v.Url,
			v.Cover}
	}
	model := &InfoListModel{items: m.items}
	_ = mw.lb.SetModel(model)
	mw.infoListModel = m
}

// 获取微博信息
func (mw *InfoFlow) getWeiboList() {
	if !mw.weibo.Checked() {
		return
	}
	hotBand, err := weibo.GetHotBand()
	if err != nil {
		return
	}
	m := &InfoListModel{items: make([]Item, len(hotBand.Data.BandList)+1)}
	i := 0
	if hotBand.Data.HotGov.IsHot == 1 {
		hotGov := hotBand.Data.HotGov
		m.items[i] = Item{strconv.Itoa(i+1) + "、" + hotGov.Name +
			"\n", hotGov.Mblog.Text, hotGov.Url, hotGov.Mblog.PageInfo.PagePic}
		i++
	}
	for _, v := range hotBand.Data.BandList {
		m.items[i] = Item{strconv.Itoa(i+1) + "、" + v.Word +
			"\n", v.Mblog.Text, v.Mblog.PageInfo.MediaInfo.H5Url, v.Mblog.PageInfo.PagePic}
		i++
	}
	model := &InfoListModel{items: m.items}
	_ = mw.lb.SetModel(model)
	mw.infoListModel = m
}

// 获取知乎热榜数据
func getZhihuList(mw *InfoFlow) *InfoListModel {
	if mw != nil && !mw.zhihu.Checked() {
		return nil
	}

	topStory, err := zhihu.GetTopStory()
	if err != nil {
		walk.MsgBox(mw, "网络错误", "请连接网络", walk.MsgBoxIconError)
		return nil
	}

	m := &InfoListModel{items: make([]Item, len(topStory.Data))}

	for i, v := range topStory.Data {
		name := v.Target.TitleArea.Text
		value := v.Target.ExcerptArea.Text
		url := v.Target.Link.Url
		imgUrl := v.Target.ImageArea.Url
		m.items[i] = Item{strconv.Itoa(i+1) + "、" +
			name + "\n", value, url, imgUrl}
	}
	if mw != nil {
		model := &InfoListModel{items: m.items}
		_ = mw.lb.SetModel(model)
		mw.infoListModel = m
	}
	return m
}

// 点击信息列
func (mw *InfoFlow) lbCurrentIndexChanged() {
	i := mw.lb.CurrentIndex()
	if i < 0 {
		return
	}
	item := &mw.infoListModel.items[i]
	_ = mw.infoPanel.SetText(item.value)
	_ = mw.wv.SetURL(item.imgUrl)
}

// 双击信息列
func (mw *InfoFlow) lbItemActivated() {
	i := mw.lb.CurrentIndex()
	if i < 0 {
		return
	}
	item := &mw.infoListModel.items[i]

	_ = mw.wv.SetURL(item.url)
}

// 关于作者
func (mw *InfoFlow) showAboutBoxAction() {
	walk.MsgBox(mw, "About", "InfoFlow-V1.0.0\n作者：豆小匠"+
		"\n本软件仅用于学习，不做商业用途!\n2021-11-30", walk.MsgBoxUserIcon)
}

func (m *InfoListModel) ItemCount() int {
	return len(m.items)
}

func (m *InfoListModel) Value(index int) interface{} {
	return m.items[index].name
}
