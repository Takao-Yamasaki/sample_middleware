package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})
	
	// ハンドラ:myMiddleware1をパス:/に登録している
	http.Handle("/", myMiddleware1(helloHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// helloHandlerを受け取って、新しいハンドラを返す関数
func myMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Pre-process1\n")
		
		// 元のハンドラを実行
		// ミドルウェアを導入 = 引数nextに渡す
		next.ServeHTTP(w, r)
		
		io.WriteString(w, "Post-process1\n")
	})
}