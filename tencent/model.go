package tencent

import (
	"fmt"
	"github.com/winterssy/music-get/common"
	"github.com/winterssy/music-get/utils"
	"strings"
	"time"
)

const (
	AlbumPicUrl = "https://y.gtimg.cn/music/photo_new/T002R300x300M000%s.jpg"
)

type Singer struct {
	Id   int    `json:"id"`
	Mid  string `json:"mid"`
	Name string `json:"name"`
}

type Album struct {
	Id   int    `json:"id"`
	Mid  string `json:"mid"`
	Name string `json:"name"`
}

type GetAlbumInfo struct {
	FAlbumId   string `json:"Falbum_id"`
	FAlbumMid  string `json:"Falbum_mid"`
	FAlbumName string `json:"Falbum_name"`
}

type Song struct {
	Id         int      `json:"id"`
	Mid        string   `json:"mid"`
	Name       string   `json:"name"`
	Singer     []Singer `json:"singer"`
	Album      Album    `json:"album"`
	IndexAlbum int      `json:"index_album"`
	TimePublic string   `json:"time_public"`
	Action     struct {
		Switch int `json:"switch"`
	} `json:"action"`
}

type CD struct {
	DissTid  string `json:"disstid"`
	DissName string `json:"dissname"`
	SongList []Song `json:"songlist"`
}

func (s *Song) Extract() *common.MP3 {
	title, album := strings.TrimSpace(s.Name), strings.TrimSpace(s.Album.Name)
	// QQ音乐支持下载无版权的音乐
	//playable := s.Action.Switch == 65537
	publishTime, _ := time.Parse("2006-01-02", s.TimePublic)
	year, track := fmt.Sprintf("%d", publishTime.Year()), fmt.Sprintf("%d", s.IndexAlbum)
	coverImage := fmt.Sprintf(AlbumPicUrl, s.Album.Mid)

	artistList := make([]string, 0, len(s.Singer))
	for _, ar := range s.Singer {
		artistList = append(artistList, strings.TrimSpace(ar.Name))
	}
	artist := strings.Join(artistList, "/")

	fileName := utils.TrimInvalidFilePathChars(fmt.Sprintf("%s - %s.mp3", strings.Join(artistList, " "), title))
	tag := common.Tag{
		Title:      title,
		Artist:     artist,
		Album:      album,
		Year:       year,
		Track:      track,
		CoverImage: coverImage,
	}

	return &common.MP3{
		FileName: fileName,
		Playable: true,
		Tag:      tag,
		Origin:   common.TencentMusic,
	}
}
