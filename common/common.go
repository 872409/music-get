package common

import (
	"errors"
	"github.com/bogem/id3v2"
	"github.com/winterssy/music-get/utils"
	"github.com/winterssy/music-get/utils/logger"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	NeteaseMusic = 1000 + iota
	TencentMusic
)

const (
	DownloadSuccess = 2000 + iota
	DownloadNoCopyrightError
	DownloadBuildPathError
	DownloadHTTPRequestError
	DownloadBuildFileError
	DownloadFileTransferError
)

type Tag struct {
	Title      string
	Artist     string
	Album      string
	Year       string
	Track      string
	CoverImage string
}

type MP3 struct {
	FileName    string
	SavePath    string
	Playable    bool
	DownloadUrl string
	Tag         Tag
	Origin      int
}

type DownloadTask struct {
	MP3    *MP3
	Status int
}

func (m *MP3) UpdateTag(wg *sync.WaitGroup) {
	var err error
	defer func() {
		if err != nil {
			logger.Error.Printf("Update music tag error: %s: %s", m.FileName, err.Error())
		} else {
			logger.Info.Printf("Music tag updated: %s", m.FileName)
		}
		wg.Done()
	}()

	file := filepath.Join(m.SavePath, m.FileName)
	resp, err := Request("GET", m.Tag.CoverImage, nil, nil, NeteaseMusic)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
	if err != nil {
		return
	}
	defer tag.Close()

	tag.SetDefaultEncoding(id3v2.EncodingUTF8)
	pic := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/jpg",
		PictureType: id3v2.PTOther,
		Picture:     data,
	}
	tag.AddAttachedPicture(pic)
	tag.SetTitle(m.Tag.Title)
	tag.SetArtist(m.Tag.Artist)
	tag.SetAlbum(m.Tag.Album)
	tag.SetYear(m.Tag.Year)
	textFrame := id3v2.TextFrame{
		Encoding: id3v2.EncodingUTF8,
		Text:     m.Tag.Track,
	}
	tag.AddFrame(tag.CommonID("Track number/Position in set"), textFrame)

	err = tag.Save()
	return
}

func (m *MP3) SingleDownload() error {
	m.SavePath = filepath.Join(MP3DownloadDir, m.SavePath)
	if err := utils.BuildPathIfNotExist(m.SavePath); err != nil {
		return err
	}

	resp, err := Request("GET", m.DownloadUrl, nil, nil, m.Origin)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fPath := filepath.Join(m.SavePath, m.FileName)
	f, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer f.Close()

	bar := pb.New(int(resp.ContentLength)).SetUnits(pb.U_BYTES).SetRefreshRate(100 * time.Millisecond)
	bar.ShowSpeed = true
	bar.Start()
	reader := bar.NewProxyReader(resp.Body)
	n, err := io.Copy(f, reader)
	if err != nil {
		return err
	}
	if n != resp.ContentLength {
		return errors.New("file transfer interrupted")
	}

	bar.Finish()
	return nil
}

func (m *MP3) ConcurrentDownload(taskList chan DownloadTask, taskQueue chan struct{}, wg *sync.WaitGroup) {
	var err error
	task := DownloadTask{
		MP3: m,
	}

	defer func() {
		if err != nil {
			logger.Error.Printf("Download error: %s: %s", m.FileName, err.Error())
		}
		wg.Done()
		taskList <- task
		<-taskQueue
	}()

	if !m.Playable {
		logger.Info.Printf("Ignore no copyright music: %s", m.Tag.Title)
		task.Status = DownloadNoCopyrightError
		return
	}

	logger.Info.Printf("Downloading: %s", m.FileName)
	m.SavePath = filepath.Join(MP3DownloadDir, m.SavePath)
	if err = utils.BuildPathIfNotExist(m.SavePath); err != nil {
		task.Status = DownloadBuildPathError
		return
	}

	resp, err := Request("GET", m.DownloadUrl, nil, nil, m.Origin)
	if err != nil {
		task.Status = DownloadHTTPRequestError
		return
	}
	defer resp.Body.Close()

	fPath := filepath.Join(m.SavePath, m.FileName)
	f, err := os.Create(fPath)
	if err != nil {
		task.Status = DownloadBuildFileError
		return
	}
	defer f.Close()

	n, err := io.Copy(f, resp.Body)
	if err != nil {
		task.Status = DownloadFileTransferError
		return
	}
	if n != resp.ContentLength {
		task.Status = DownloadFileTransferError
		return
	}

	task.Status = DownloadSuccess
	logger.Info.Printf("Download complete: %s", m.FileName)
	return
}
