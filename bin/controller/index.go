package controller

import "github.com/TomSuzuki/TomSuzuki.github.io_new/bin/model"

// Index ...全体定義のファイル処理
func Index(tags map[string]interface{}) string {
	return model.HTMLTemplate("index.html", tags)
}
