package main

import "github.com/lxn/walk"

// InfoFlow InfoFlow窗口元素
type InfoFlow struct {
	*walk.MainWindow                // 主窗口
	lb               *walk.ListBox  // 信息列表
	wv               *walk.WebView  // 网页展示
	infoPanel        *walk.TextEdit // 信息详情

	infoListModel *InfoListModel // 信息列表数据模型

	zhihu    *walk.RadioButton // 单选框
	weibo    *walk.RadioButton
	bilibili *walk.RadioButton
	csdn     *walk.RadioButton

	keywords *walk.LineEdit   // 搜索框
	query    *walk.PushButton //搜索按钮
}

// Item 信息元素
type Item struct {
	name   string
	value  string
	url    string
	imgUrl string
}

// InfoListModel 信息数据模型
type InfoListModel struct {
	walk.ListModelBase
	items []Item
}
