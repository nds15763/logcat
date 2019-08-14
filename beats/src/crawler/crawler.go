package crawler

import (
	"beats/src/config"
	"beats/src/model"
	"bufio"
	"fmt"
	"os"
	"time"
)

type Crawler struct {
	cfg     []config.LogType
	FileMap map[string]config.LogType
	beats   chan model.Beat
}

func NewCrawler(cfg *config.Config) *Crawler {

	fileMap := func() map[string]config.LogType {

		fileMap := make(map[string]config.LogType)
		for _, v := range cfg.Items.LogType {
			fileMap[v.Paths] = v
		}
		return fileMap
	}()
	//查找所有符合的文件，并记录文件名
	return &Crawler{
		cfg:     cfg.Items.LogType,
		FileMap: fileMap,
		beats:   make(chan model.Beat, len(cfg.Items.LogType)),
	}
}

func (this *Crawler) Run() chan chan model.Beat {
	this.getFileMap()
	return this.DoFetch()
}

// 逐行读取文件内容
func (this *Crawler) readLines(fpath string) []string {
	fd, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	var lines []string
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return lines
}

func (this *Crawler) getFileMap() {
	//先去找所有匹配到的log file
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			for k, v := range this.FileMap {
				fileIDlist := this.readLines(k)
				if len(fileIDlist) == 0 {
					continue
				}
				for _, v1 := range fileIDlist {
					this.FileMap[v1] = v //检查文件是否已经存在
				}
			}

		}
	}
}

func (this *Crawler) DoFetch() chan chan model.Beat {

	ch := make(chan chan model.Beat, len(this.FileMap))

	for fileName, cfg := range this.FileMap {
		// node := nodelist[r.Intn(len(nodelist))]
		// go func(node, fileID string) {
		// 	ch <- Download(*clustername, node, fileID)
		// }(node, fileID)
		go func(filename string) {
			beater := NewBeater(cfg, fileName)
			ch <- beater.Run()
		}(fileName)

	}

	return ch
}
