package main

import (
	"log"
	"os"

	"go_project/router"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8002"
	}

	log.Printf("真理之泉服务启动，监听 %s", addr)
	if err := router.Start(addr); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
