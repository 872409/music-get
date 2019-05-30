package netease

import (
	"fmt"
	"github.com/winterssy/music-get/common"
	"github.com/winterssy/music-get/utils"
	"strings"
	"time"
)

type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PicUrl      string `json:"picUrl"`
	PublishTime int64  `json:"publishTime"`
}

type SongUrl struct {
	Id   int    `json:"id"`
	Code int    `json:"code"`
	Url  string `json:"url"`
}

type Song struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Artist      []Artist `json:"ar"`
	Album       Album    `json:"al"`
	Position    int      `json:"no"`
	PublishTime int64    `json:"publishTime"`
}

type TrackId struct {
	Id int `json:"id"`
}

type Playlist struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	TrackIds []TrackId `json:"trackIds"`
}

func (s *Song) Extract() *common.MP3 {
	title, album := strings.TrimSpace(s.Name), strings.TrimSpace(s.Album.Name)
	publishTime := time.Unix(0, s.PublishTime*1000*1000)
	year, track := fmt.Sprintf("%d", publishTime.Year()), fmt.Sprintf("%d", s.Position)
	coverImage := s.Album.PicUrl

	artistList := make([]string, 0, len(s.Artist))
	for _, ar := range s.Artist {
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
		Tag:      tag,
		Origin:   common.NeteaseMusic,
	}
}
