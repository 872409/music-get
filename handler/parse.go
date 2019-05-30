package handler

import (
	"fmt"
	"github.com/winterssy/music-get/common"
	"github.com/winterssy/music-get/netease"
	"github.com/winterssy/music-get/tencent"
	"regexp"
)

const (
	UrlPattern = "music.163.com|y.qq.com"
)

func Parse(url string) (req common.MusicRequest, err error) {
	re := regexp.MustCompile(UrlPattern)
	matched, ok := re.FindString(url), re.MatchString(url)
	if !ok {
		err = fmt.Errorf("could not parse the url: %s", url)
		return
	}

	switch matched {
	case "music.163.com":
		req, err = netease.Parse(url)
	case "y.qq.com":
		req, err = tencent.Parse(url)
	}

	return
}
