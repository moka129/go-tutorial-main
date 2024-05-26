package chapter1

import (
	"fmt"
	"net/http"
)

// revive:disable

func GetEcho(w http.ResponseWriter, r *http.Request) {
	// Getメソッドのアクセスか確認
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// パラメータをFormに変換する
	if err := r.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// パラメータを取得する
	type params struct {
		Name string
		Age  string
	}

	// パラメータをレスポンスに書き出す
	var p params
	p.Name = r.Form.Get("name")
	p.Age = r.Form.Get("age")

	// レスポンスコード設定
	_, err := w.Write([]byte(fmt.Sprintf("name: %s, age: %s\n", p.Name, p.Age)))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
