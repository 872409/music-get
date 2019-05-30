package netease

import (
	"encoding/json"
	"fmt"
	"github.com/winterssy/music-get/common"
	"github.com/winterssy/music-get/utils"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

const (
	WeAPI       = "https://music.163.com/weapi"
	SongUrlAPI  = WeAPI + "/song/enhance/player/url"
	SongAPI     = WeAPI + "/v3/song/detail"
	ArtistAPI   = WeAPI + "/v1/artist"
	AlbumAPI    = WeAPI + "/v1/album"
	PlaylistAPI = WeAPI + "/v3/playlist/detail"
)

type SongUrlParams struct {
	Ids string `json:"ids"`
	Br  int    `json:"br"`
}

type SongUrlResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []SongUrl `json:"data"`
}

type SongUrlRequest struct {
	Params   SongUrlParams
	Response SongUrlResponse
}

type SongParams struct {
	C string `json:"c"`
}

type SongResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Songs []Song `json:"songs"`
}

type SongRequest struct {
	Params   SongParams
	Response SongResponse
}

type ArtistParams struct{}

type ArtistResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Artist   Artist `json:"artist"`
	HotSongs []Song `json:"hotSongs"`
}

type ArtistRequest struct {
	Id       int
	Params   ArtistParams
	Response ArtistResponse
}

type AlbumParams struct{}

type AlbumResponse struct {
	SongResponse
	Album Album `json:"album"`
}

type AlbumRequest struct {
	Id       int
	Params   AlbumParams
	Response AlbumResponse
}

type PlaylistParams struct {
	Id int `json:"id"`
}

type PlaylistResponse struct {
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
	Playlist Playlist `json:"playlist"`
}

type PlaylistRequest struct {
	Params   PlaylistParams
	Response PlaylistResponse
}

func NewSongUrlRequest(ids ...int) *SongUrlRequest {
	br := common.MP3DownloadBr
	switch br {
	case 128, 192, 320:
		br *= 1000
		break
	default:
		br = 999 * 1000
	}
	enc, _ := json.Marshal(ids)
	return &SongUrlRequest{Params: SongUrlParams{Ids: string(enc), Br: br}}
}

func NewSongRequest(ids ...int) *SongRequest {
	c := make([]map[string]int, 0, len(ids))
	for _, id := range ids {
		c = append(c, map[string]int{"id": id})
	}

	enc, _ := json.Marshal(c)
	return &SongRequest{Params: SongParams{C: string(enc)}}
}

func NewArtistRequest(id int) *ArtistRequest {
	return &ArtistRequest{Id: id, Params: ArtistParams{}}
}

func NewAlbumRequest(id int) *AlbumRequest {
	return &AlbumRequest{Id: id, Params: AlbumParams{}}
}

func NewPlaylistRequest(id int) *PlaylistRequest {
	return &PlaylistRequest{Params: PlaylistParams{Id: id}}
}

func (s *SongUrlRequest) Do() error {
	enc, _ := json.Marshal(s.Params)
	params, encSecKey, err := Encrypt(enc)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	resp, err := common.Request("POST", SongUrlAPI, nil, strings.NewReader(form.Encode()), common.NeteaseMusic)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s.Response); err != nil {
		return err
	}
	if s.Response.Code != http.StatusOK {
		return fmt.Errorf("%s %s error: %d %s", resp.Request.Method, resp.Request.URL.String(), s.Response.Code, s.Response.Msg)
	}

	return nil
}

func (s *SongRequest) Do() error {
	enc, _ := json.Marshal(s.Params)
	params, encSecKey, err := Encrypt(enc)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	resp, err := common.Request("POST", SongAPI, nil, strings.NewReader(form.Encode()), common.NeteaseMusic)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s.Response); err != nil {
		return err
	}
	if s.Response.Code != http.StatusOK {
		return fmt.Errorf("%s %s error: %d %s", resp.Request.Method, resp.Request.URL.String(), s.Response.Code, s.Response.Msg)
	}

	return nil
}

func (s *SongRequest) Extract() ([]*common.MP3, error) {
	return ExtractMP3List(s.Response.Songs, ".")
}

func (a *ArtistRequest) Do() error {
	enc, _ := json.Marshal(a.Params)
	params, encSecKey, err := Encrypt(enc)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	resp, err := common.Request("POST", ArtistAPI+fmt.Sprintf("/%d", a.Id), nil, strings.NewReader(form.Encode()), common.NeteaseMusic)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&a.Response); err != nil {
		return err
	}
	if a.Response.Code != http.StatusOK {
		return fmt.Errorf("%s %s error: %d %s", resp.Request.Method, resp.Request.URL.String(), a.Response.Code, a.Response.Msg)
	}

	return nil
}

func (a *ArtistRequest) Extract() ([]*common.MP3, error) {
	ids := make([]int, 0, len(a.Response.HotSongs))
	for _, i := range a.Response.HotSongs {
		ids = append(ids, i.Id)
	}

	req := NewSongRequest(ids...)
	if err := req.Do(); err != nil {
		return nil, err
	}

	savePath := filepath.Join(".", utils.TrimInvalidFilePathChars(a.Response.Artist.Name))
	return ExtractMP3List(req.Response.Songs, savePath)
}

func (a *AlbumRequest) Do() error {
	enc, _ := json.Marshal(a.Params)
	params, encSecKey, err := Encrypt(enc)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	resp, err := common.Request("POST", AlbumAPI+fmt.Sprintf("/%d", a.Id), nil, strings.NewReader(form.Encode()), common.NeteaseMusic)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&a.Response); err != nil {
		return err
	}
	if a.Response.Code != http.StatusOK {
		return fmt.Errorf("%s %s error: %d %s", resp.Request.Method, resp.Request.URL.String(), a.Response.Code, a.Response.Msg)
	}

	return nil
}

func (a *AlbumRequest) Extract() ([]*common.MP3, error) {
	savePath := filepath.Join(".", utils.TrimInvalidFilePathChars(a.Response.Album.Name))
	for i := range a.Response.Songs {
		a.Response.Songs[i].PublishTime = a.Response.Album.PublishTime
	}
	return ExtractMP3List(a.Response.Songs, savePath)
}

func (p *PlaylistRequest) Do() error {
	enc, _ := json.Marshal(p.Params)
	params, encSecKey, err := Encrypt(enc)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	resp, err := common.Request("POST", PlaylistAPI, nil, strings.NewReader(form.Encode()), common.NeteaseMusic)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&p.Response); err != nil {
		return err
	}
	if p.Response.Code != http.StatusOK {
		return fmt.Errorf("%s %s error: %d %s", resp.Request.Method, resp.Request.URL.String(), p.Response.Code, p.Response.Msg)
	}

	return nil
}

func (p *PlaylistRequest) Extract() ([]*common.MP3, error) {
	ids := make([]int, 0, len(p.Response.Playlist.TrackIds))
	for _, i := range p.Response.Playlist.TrackIds {
		ids = append(ids, i.Id)
	}

	req := NewSongRequest(ids...)
	if err := req.Do(); err != nil {
		return nil, err
	}

	savePath := filepath.Join(".", utils.TrimInvalidFilePathChars(p.Response.Playlist.Name))
	return ExtractMP3List(req.Response.Songs, savePath)
}
