package netease

import (
	"github.com/winterssy/music-get/common"
)

func ExtractMP3List(songs []Song, savePath string) ([]*common.MP3, error) {
	n := len(songs)
	ids := make([]int, 0, n)
	for _, i := range songs {
		ids = append(ids, i.Id)
	}

	req := NewSongUrlRequest(ids...)
	if err := req.Do(); err != nil {
		return nil, err
	}

	codeMap, urlMap := make(map[int]int, n), make(map[int]string, n)
	for _, i := range req.Response.Data {
		codeMap[i.Id] = i.Code
		urlMap[i.Id] = i.Url
	}

	mp3List := make([]*common.MP3, 0, n)
	for _, i := range songs {
		mp3 := i.Extract()
		mp3.SavePath = savePath
		mp3.Playable = codeMap[i.Id] == 200
		mp3.DownloadUrl = urlMap[i.Id]
		mp3List = append(mp3List, mp3)
	}

	return mp3List, nil
}
