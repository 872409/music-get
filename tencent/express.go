package tencent

import (
	"encoding/json"
	"github.com/winterssy/music-get/common"
	"math/rand"
	"strconv"
	"time"
)

const (
	MusicExpressAPI = "http://base.music.qq.com/fcgi-bin/fcg_musicexpress.fcg"
)

type MusicExpress struct {
	Code    int      `json:"code"`
	SIP     []string `json:"sip"`
	ThirdIP []string `json:"thirdip"`
	Key     string   `json:"key"`
}

func createGuid() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Intn(10000000000-1000000000) + 1000000000)
}

func getVkey(guid string) (string, error) {
	query := map[string]string{
		"guid":   guid,
		"format": "json",
	}

	resp, err := common.Request("GET", MusicExpressAPI, query, nil, common.TencentMusic)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	var m MusicExpress
	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return "", nil
	}

	return m.Key, nil
}
