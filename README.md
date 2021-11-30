# InfoFlow    信息流

#### 介绍

InfoFlow是一个由纯 Golang 制作的 windows 程序，主要用于学习 golang 语法和少量网络编程的知识。
具体功能是收集一些平台的热榜，在软件上可以方便的查看。（仅用于学习）


#### 软件架构
软件架构说明

├─.git
├─bilibili	// bilibili接口
├─csdn	// csdn接口
├─img	// 图片
├─utils	// 工具包
├─weibo	// 微博接口
└─zhihu	// 知乎接口


#### 安装教程

- 本地需要有 go 环境

1. 执行 go build -ldflags="-H windowsgui" 生成exe文件，运行

#### 使用说明
![截图](https://images.gitee.com/uploads/images/2021/1129/233855_fd8017ad_7377902.png "屏幕截图.png")

![image-20211129234433542](https://gitee.com/anjude/public-resource/raw/md-img/20211129234433.png)

1. 点击单选框选择信息源
2. 根据关键词搜索内容
3. 单击信息流展示详细信息和图片
4. 双击信息流显示 webview 页面，因为 使用的开源 gui 工具对 webview 支持不够，会报脚本错误。

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

#### 鸣谢

- gui工具：https://pkg.go.dev/github.com/lxn/walk
- 参考项目：https://studygolang.com/articles/10755

#### 关注我

![公众号二维码](https://gitee.com/anjude/public-resource/raw/md-img/20211129235003.jpg)
