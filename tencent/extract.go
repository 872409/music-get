package tencent

import (
	"fmt"
	"github.com/winterssy/music-get/common"
)

const (
	SelfSongDownloadUrl  = "http://dl.stream.qqmusic.qq.com/M500%s.mp3?guid=%s&vkey=%s&fromtag=1"
	ThirdSongDownloadAPI = "https://v1.itooi.cn/tencent/url?id=%s&quality=%d"
)

func ExtractMP3List(songs []Song, savePath string) ([]*common.MP3, error) {
	// 测试发现 guid 可以是随机字符串
	guid := "yqq"
	vkey, err := getVkey(guid)
	if err != nil {
		return nil, err
	}

	br := common.MP3DownloadBr
	mp3List := make([]*common.MP3, 0, len(songs))
	for _, i := range songs {
		mp3 := i.Extract()
		mp3.SavePath = savePath
		switch br {
		case 192, 320:
			mp3.DownloadUrl = fmt.Sprintf(ThirdSongDownloadAPI, i.Mid, br)
			break
		default:
			mp3.DownloadUrl = fmt.Sprintf(SelfSongDownloadUrl, i.Mid, guid, vkey)
		}
		mp3List = append(mp3List, mp3)
	}

	return mp3List, nil
}
