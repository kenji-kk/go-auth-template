package main

import (
	"github.com/kenji-kk/go-auth-template/internal/server"
)

func main() {
	// ポート番号 8080 で待ち受けを開始
	server := server.NewServer()
	server.Run()
}
