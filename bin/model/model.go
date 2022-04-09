package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

// JsonLoad ...jsonファイルの読み込みを行う
func JsonLoad(path string) map[string]interface{} {
	var result map[string]interface{}
	raw, _ := ioutil.ReadFile(path)
	json.Unmarshal(raw, &result)
	return result
}

// HTMLTemplate ...テンプレートを処理する（後で再帰を html/template の機能に置き換え → 変数指定できない？）
func HTMLTemplate(file string, data interface{}) string {
	f := template.FuncMap{
		"HTMLTemplate": func(file string, data interface{}) template.HTML {
			return template.HTML(HTMLTemplate(file, data))
		},
		"JsonLoad": func(path string) map[string]interface{} {
			return JsonLoad("./template/" + path)
		},
		"DataMap": func(data interface{}) map[string]interface{} {
			return data.(map[string]interface{})
		},
		"NilMap": func() map[string]interface{} {
			return map[string]interface{}{}
		},
		"DataString": func(data interface{}) string {
			fmt.Println(data)
			return "title" //data.(string)
		},
		"SafeHTML": func(text string) template.HTML {
			return template.HTML(text)
		},
		"Add": func(i int, j int) int {
			return i + j
		},
		"Sub": func(i int, j int) int {
			return i - j
		},
		"Mod": func(i int, j int) int {
			return i % j
		},
		"Div": func(i int, j int) int {
			return int(i / j)
		},
		"Len": func(i []interface{}) int {
			return len(i)
		},
	}
	var body bytes.Buffer
	parts, _ := filepath.Glob("./template/_*")
	files := append([]string{"./template/" + file}, parts...)
	tpl := template.New("")
	t, _ := tpl.New(file).Funcs(f).ParseFiles(files...)
	t.Execute(&body, data)
	return body.String()
}
