package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadRemoteEnv(envName, pass string) {
	res, err := http.Get(fmt.Sprintf("https://renv.deno.dev/%s?pass=%s", envName, pass))
	if err != nil {
		log.Fatal("无法获取环境变量信息")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal((fmt.Errorf("bad status: %s", res.Status)))
	}

	b, _ := ioutil.ReadAll(res.Body)
	env, _ := godotenv.Unmarshal(string(b))

	for k, v := range env {
		os.Setenv(k, v)
	}
}
