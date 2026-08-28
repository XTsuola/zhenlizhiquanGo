package router

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"go_project/controllers"
)

// StartHTTPS 可选 HTTPS 启动（需 CERT_FILE / KEY_FILE，或默认 router 目录下证书）
func StartHTTPS(addr string) error {
	if addr == "" {
		addr = ":8002"
	}
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")
	if certFile == "" || keyFile == "" {
		_, currentFile, _, _ := runtime.Caller(0)
		dir := filepath.Dir(currentFile)
		if certFile == "" {
			certFile = filepath.Join(dir, "server.pem")
		}
		if keyFile == "" {
			keyFile = filepath.Join(dir, "server.key")
		}
	}
	log.Printf("HTTPS 服务启动: %s", addr)
	return controllers.R.RunTLS(addr, certFile, keyFile)
}
