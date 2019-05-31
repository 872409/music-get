# Music-Get

![](https://travis-ci.org/winterssy/music-get.svg?branch=master)
![](https://img.shields.io/badge/golang-1.12-blue.svg)
![](https://img.shields.io/github/release/winterssy/music-get.svg)
![](https://img.shields.io/github/license/winterssy/music-get.svg)

[网易云音乐](https://music.163.com) | [QQ音乐](https://y.qq.com) 下载助手，支持一键下载单曲/专辑/歌单以及歌手热门歌曲，并自动更新音乐标签。

>本项目仅供学习研究使用。如侵犯你的权益，请 [联系作者](mailto:winterssy@foxmail.com) 删除。

## 下载安装

你可以前往 [Releases](https://github.com/winterssy/music-get/releases) 标签下载程序的最新版本，或者克隆项目源码自行编译。

## 如何使用？

直接将音乐地址作为命令行参数传入即可，如：

- 下载单曲：
```
$ music-get https://music.163.com/#/song?id=553310243
$ music-get https://y.qq.com/n/yqq/song/002Zkt5S2z8JZx.html
```

- 下载专辑：
```
$ music-get https://music.163.com/#/album?id=38373053
$ music-get https://y.qq.com/n/yqq/song/002Zkt5S2z8JZx.html
```

- 下载歌单：
```
$ music-get https://music.163.com/#/playlist?id=156934569
$ music-get https://y.qq.com/n/yqq/album/002fRO0N4FftzY.html
```

- 下载歌手热门歌曲：
```
$ music-get https://music.163.com/#/artist?id=13193
$ music-get https://y.qq.com/n/yqq/singer/000Sp0Bz4JXH0o.html
```

命令选项：
- `-br`：优先下载音质，可选128/192/320，默认128。
- `-o`：下载保存目录，默认为 `/home/用户名/Music-Get`  （Windows为 `C:\\Users\\用户名\\Music-Get` ）。
- `-n`：并发下载任务数，最大值16，默认1，即单任务下载。
- `-h`：获取命令帮助。

**注：** 命令选项必须先于其它命令行参数输入。

## 运行截图

- 单任务下载：

![](/screenshots/single-download.png)

- 多任务同时下载：

![](/screenshots/concurrent-download.png)

- 自动更新音乐标签（效果预览）：

![](/screenshots/tag-updated.png)

## License

GPLv3.
