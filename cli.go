package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/beetcb/oa-bot/extract"
	"github.com/beetcb/oa-bot/upload"
)

type ActResult struct {
	inputPath string
	url       string
	parse     extract.ParseInfo
}

type ActResultArray []ActResult

func init() {
	if len(os.Args) <= 2 {
		panic("Usage: oa-bot {dirPath} {envPass}")
	}
	pass := os.Args[2]
	LoadRemoteEnv(pass)
	extract.FillLicense()
}

func main() {
	wg := &sync.WaitGroup{}
	dir := filepath.ToSlash(os.Args[1])
	files, err := ioutil.ReadDir(dir)
	acts := ActResultArray{}

	if err != nil {
		panic(fmt.Sprintf("指定 %s 文件夹错误", dir))
	}

	for _, file := range files {
		wg.Add(1)
		go func(file fs.FileInfo) {
			defer wg.Done()
			inputPath := filepath.Join(dir, file.Name())
			text := extract.ExtractPdfText(inputPath)
			parse := extract.ParsePdfText(inputPath, text)
			url := upload.UploadToTemp(inputPath)
			acts = append(acts, ActResult{inputPath, url, parse})
		}(file)
	}

	wg.Wait()

	for _, act := range acts {
		fmt.Println(act)
		up(act.inputPath, act.url, act.parse)
	}
}

// Sync upload process due to API limit
func up(inputPath string, url string, parse extract.ParseInfo) {
	// 上传文档到 MINGDAO
	uploadResult := upload.UploadToMingDao(url, parse)
	fmt.Println(uploadResult)

	// 成功则删除本地文件
	if strings.Contains(uploadResult, "true") {
		os.Remove(inputPath)
	}
}
