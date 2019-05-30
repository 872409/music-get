# Music-Get

[![Build Status](https://travis-ci.org/winterssy/music-get.svg?branch=master)](https://travis-ci.org/winterssy/music-get)
[![License GPL-3.0](https://img.shields.io/badge/license-GPLv3.0-blue.svg)](https://github.com/winterssy/music-get/blob/master/LICENSE)

[网易云音乐](https://music.163.com) | [QQ音乐](https://y.qq.com) 下载助手，支持一键下载单曲/专辑/歌单以及歌手热门歌曲，并自动更新音乐标签。

>本项目仅供学习研究使用。如侵犯你的权益，请 [联系作者](mailto:winterssy@foxmail.com) 删除。

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

## License

GPLv3.
