package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/beetcb/oa-bot/extract"
	"github.com/beetcb/oa-bot/upload"
)

func main() {
	LoadRemoteEnv()

	if len(os.Args) <= 1 {
		panic("Usage: oa-bot {dirPath}")
	}

	dir := filepath.ToSlash(os.Args[1])
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(fmt.Sprintf("指定 %s 文件夹错误", dir))
	}

	for _, file := range files {
		inputPath := filepath.Join(dir, file.Name())
		text := extract.ExtractPdfText(inputPath)

		// 读取文档信息
		parse := extract.ParsePdfText(inputPath, text)
		fmt.Println(parse)

		url := upload.UploadToTemp(inputPath)

		// 上传文档到 MINGDAO
		uploadResult := upload.UploadToMingDao(url, parse)

		// 成功则删除本地文件
		if uploadResult {
			os.Remove(inputPath)
		}
	}
}
