package main

import (
	"fmt"
	_ "fmt"
	"io/ioutil"
	"os"

	"github.com/TomSuzuki/TomSuzuki.github.io_new/bin/controller"
	"github.com/TomSuzuki/TomSuzuki.github.io_new/bin/model"
)

// main ...実行時にこれが呼び出される。
func main() {
	// メッセージ
	fmt.Println("処理を開始します。")

	// 設定の読み込み
	setting := model.JsonLoad("./data/setting.json")

	// 全体定義の読み込み
	indexData := model.JsonLoad("./data/index.json")

	// パーツの処理（全体定義 → 読み込み → bodyのHTML）
	body := string(controller.Index(indexData["data"].(map[string]interface{})))

	// 古いデータの削除
	if f, err := os.Stat(setting["output_path"].(string)); !(os.IsNotExist(err) || !f.IsDir()) {
		os.RemoveAll(setting["output_path"].(string))
	}
	os.MkdirAll(setting["output_path"].(string), 0777)

	// 出力
	outputPath := setting["output_path"].(string) + "index.html"
	ioutil.WriteFile(outputPath, []byte(body), 0666)

	// コピー
	files := model.Dirwalk("assets", true)
	for _, f := range files {
		model.CopyFile(f, setting["output_path"].(string)+f)
	}

	// メッセージ
	fmt.Println("正常に終了しました。")
}
