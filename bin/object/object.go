package object

// SettingData ...設定ファイルのデータリスト
type SettingData struct {
	OutputPath string `json:"output_path"`
}

// IndexData ...全体定義のデータ
type IndexData struct {
	Tags []string `json:"tags"`
}
