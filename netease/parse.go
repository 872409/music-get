package netease

import (
	"fmt"
	"github.com/winterssy/music-get/common"
	"regexp"
	"strconv"
)

const (
	UrlPattern = "/(song|artist|album|playlist)\\?id=(\\d+)"
)

func Parse(url string) (req common.MusicRequest, err error) {
	re := regexp.MustCompile(UrlPattern)
	matched, ok := re.FindStringSubmatch(url), re.MatchString(url)
	if !ok || len(matched) < 3 {
		err = fmt.Errorf("could not parse the url: %s", url)
		return
	}

	id, err := strconv.Atoi(matched[2])
	if err != nil {
		return
	}

	switch matched[1] {
	case "song":
		req = NewSongRequest(id)
	case "artist":
		req = NewArtistRequest(id)
	case "album":
		req = NewAlbumRequest(id)
	case "playlist":
		req = NewPlaylistRequest(id)
	}

	return
}
