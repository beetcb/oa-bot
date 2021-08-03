package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/beetcb/oa-bot/extract"
	"github.com/beetcb/oa-bot/upload"
)

type ActResult struct {
	inputPath string
	url       string
	parse     extract.ParseInfo
}

func init() {
	if len(os.Args) <= 3 {
		panic("Usage: oa-bot {dirPath} {envName} {envPass}")
	}
	envName, pass := os.Args[2], os.Args[3]
	LoadRemoteEnv(envName, pass)
	extract.FillLicense()
}

func main() {
	dir := filepath.ToSlash(os.Args[1])
	files, err := ioutil.ReadDir(dir)
	l := len(files)
	if err != nil || l == 0 {
		panic(fmt.Sprintf("指定 %s 文件夹不存在或为空", dir))
	}

	c := make(chan ActResult)

	for i, file := range files {
		go func(file fs.FileInfo, i int) {
			inputPath := filepath.Join(dir, file.Name())
			text := extract.ExtractPdfText(inputPath)
			parse := extract.ParsePdfText(inputPath, text)
			url := upload.UploadToTemp(inputPath)
			c <- ActResult{inputPath, url, parse}
		}(file, i)
	}

	for i := 0; i < l; i++ {
		act := <-c
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
