package main

import (
	"log"
	_ "main/matchers"
	"main/search"
	"os"
)

func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("American")

}
