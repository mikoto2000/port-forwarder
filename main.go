package main

import (
	"io"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)

// ポートフォワーディングの主処理を行う。
// 転送先・転送元のデータ通信をリレーする。
func forward(src net.Conn, dstAddr string) {
	defer src.Close()

	// 転送先に接続
	dst, err := net.Dial("tcp", dstAddr)
	if err != nil {
		log.Printf("転送先への接続失敗: %v\n", err)
		return
	}
	defer dst.Close()

	// 双方向のデータ転送を行う
		// クライアント → 転送先
	go  io.Copy(dst, src)

	// 転送先 → クライアント
	io.Copy(src, dst)
}

// ポートフォワーディング開始
func startForwarding(listenAddr, forwardAddr string) {
	// listener 生成
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("リッスン失敗: %v\n", err)
	}
	defer listener.Close()

	log.Printf("リッスン中: %s -> %s\n", listenAddr, forwardAddr)

	for {
		// listen 開始
		client, err := listener.Accept()
		if err != nil {
			log.Printf("接続エラー: %v\n", err)
			continue
		}

		// 主処理(バケツリレー)を実行
		go forward(client, forwardAddr)
	}
}

func main() {
	app := &cli.App{
		Name:  "port-forwarder",
		Usage: "指定されたアドレス間でポートフォワーディングを行います",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "listen",
				Aliases:  []string{"l"},
				Usage:    "リッスンするアドレスとポート (例: 127.0.0.1:8080)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "forward",
				Aliases:  []string{"f"},
				Usage:    "転送先のアドレスとポート (例: example.com:443)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			// リッスンアドレスと転送先アドレスを取得
			listenAddr := c.String("listen")
			forwardAddr := c.String("forward")

			// ポートフォワーディングを開始
			startForwarding(listenAddr, forwardAddr)
			return nil
		},
	}

	// アプリ実行
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("アプリケーションエラー: %v\n", err)
	}
}