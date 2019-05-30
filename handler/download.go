package handler

import (
	"fmt"
	"github.com/winterssy/music-get/common"
	"github.com/winterssy/music-get/utils/logger"
	"os"
	"path/filepath"
	"sync"
)

func SingleDownload(mp3List []*common.MP3) {
	total, success, failure, ignore := len(mp3List), 0, 0, 0
	var err error
	wg := &sync.WaitGroup{}
	for _, m := range mp3List {
		if !m.Playable {
			logger.Info.Printf("Ignore no coypright music: %s", m.Tag.Title)
			ignore++
			continue
		}
		logger.Info.Printf("Downloading: %s", m.FileName)
		if err = m.SingleDownload(); err != nil {
			failure++
			logger.Error.Printf("Download error: %s", err.Error())
			_ = os.Remove(filepath.Join(m.SavePath, m.FileName))
			continue
		}
		logger.Info.Print("Download complete")
		success++

		wg.Add(1)
		go m.UpdateTag(wg)
	}
	wg.Wait()

	fmt.Printf("\nDownload report --> total: %d, success: %d, failure: %d, ignore: %d\n", total, success, failure, ignore)
}

func ConcurrentDownload(mp3List []*common.MP3, n int) {
	total, success, failure, ignore := len(mp3List), 0, 0, 0

	taskList := make(chan common.DownloadTask, total)
	taskQueue := make(chan struct{}, n)
	wg := &sync.WaitGroup{}
	wg.Add(total)
	for _, m := range mp3List {
		taskQueue <- struct{}{}
		go m.ConcurrentDownload(taskList, taskQueue, wg)
	}
	wg.Wait()

	for range mp3List {
		task := <-taskList
		switch task.Status {
		case common.DownloadSuccess:
			success++
			wg.Add(1)
			go task.MP3.UpdateTag(wg)
			break
		case common.DownloadNoCopyrightError:
			ignore++
			break
		default:
			failure++
			_ = os.Remove(filepath.Join(task.MP3.SavePath, task.MP3.FileName))
		}
	}
	wg.Wait()

	fmt.Printf("\nDownload report --> total: %d, success: %d, failure: %d, ignore: %d\n", total, success, failure, ignore)
}
