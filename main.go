package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"slices"

	"github.com/urfave/cli/v2"
)

var version string

//go:embed LICENSE
var license string

//go:embed NOTICE
var notice string

const flagNameLicense = "license"

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
	go io.Copy(dst, src)

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

	// --license の存在確認
	if slices.Contains(os.Args, "--license") || slices.Contains(os.Args, "-license") {
		// ライセンスフラグが立っていればライセンスを表示して終了
		fmt.Println(license)
		fmt.Println()
		fmt.Println(notice)
		return
	}

	app := &cli.App{
		Name:                   "port-forwarder",
		Usage:                  "指定されたアドレス間でポートフォワーディングを行います",
		Version:                version,
		UseShortOptionHandling: true,
		HideHelpCommand:        true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:               flagNameLicense,
				Value:              false,
				DisableDefaultText: true,
				Usage:              "show licensesa.",
			},
			&cli.StringFlag{
				Name: "source",
				// source, local
				Aliases:  []string{"s", "l"},
				Usage:    "source port. (ex: 127.0.0.1:8080)",
				Required: true,
			},
			&cli.StringFlag{
				Name: "destination",
				// destination, forward
				Aliases:  []string{"d", "f"},
				Usage:    "destination port. (ex: example.com:443)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {

			// リッスンアドレスと転送先アドレスを取得
			listenAddr := c.String("source")
			forwardAddr := c.String("destination")

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
