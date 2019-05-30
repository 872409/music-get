package handler

import (
	"github.com/winterssy/music-get/common"
	"github.com/winterssy/music-get/netease"
	"github.com/winterssy/music-get/tencent"
	"testing"
)

func TestParse(t *testing.T) {
	const (
		NeteaseSong     = "https://music.163.com/#/song?id=553310243"
		NeteaseArtist   = "https://music.163.com/#/artist?id=13193"
		NeteaseAlbum    = "https://music.163.com/#/album?id=38373053"
		NeteasePlaylist = "https://music.163.com/#/playlist?id=156934569"
		TencentSong     = "https://y.qq.com/n/yqq/song/002Zkt5S2z8JZx.html"
		TencentSinger   = "https://y.qq.com/n/yqq/singer/000Sp0Bz4JXH0o.html"
		TencentAlbum    = "https://y.qq.com/n/yqq/album/002fRO0N4FftzY.html"
		TencentPlaylist = "https://y.qq.com/n/yqq/playsquare/5474239760.html"
	)

	var req common.MusicRequest
	req, _ = Parse(NeteaseSong)
	if _, ok := req.(*netease.SongRequest); !ok {
		t.Errorf("failed to parse %q", NeteaseSong)
	}

	req, _ = Parse(NeteaseArtist)
	if _, ok := req.(*netease.ArtistRequest); !ok {
		t.Errorf("failed to parse %q", NeteaseArtist)
	}

	req, _ = Parse(NeteaseAlbum)
	if _, ok := req.(*netease.AlbumRequest); !ok {
		t.Errorf("failed to parse %q", NeteaseAlbum)
	}

	req, _ = Parse(NeteasePlaylist)
	if _, ok := req.(*netease.PlaylistRequest); !ok {
		t.Errorf("failed to parse %q", NeteasePlaylist)
	}

	req, _ = Parse(TencentSong)
	if _, ok := req.(*tencent.SongRequest); !ok {
		t.Errorf("failed to parse %q", TencentSong)
	}

	req, _ = Parse(TencentSinger)
	if _, ok := req.(*tencent.SingerRequest); !ok {
		t.Errorf("failed to parse %q", TencentSinger)
	}

	req, _ = Parse(TencentAlbum)
	if _, ok := req.(*tencent.AlbumRequest); !ok {
		t.Errorf("failed to parse %q", TencentAlbum)
	}

	req, _ = Parse(TencentPlaylist)
	if _, ok := req.(*tencent.PlaylistRequest); !ok {
		t.Errorf("failed to parse %q", TencentPlaylist)
	}
}
