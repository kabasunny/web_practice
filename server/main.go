package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFS("/assets", http.Dir("../client/dist/assets"))
	// /assetsパスで静的ファイルを提供するようにGinルーターを設定
	// http.Dir("../client/dist/assets")は、静的ファイルが格納されているディレクトリを指定

	router.LoadHTMLFiles("../client/dist/index.html")
	// HTMLテンプレートファイル(index.html)を読み込み

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 存在しないルートへのリクエストがあった場合の処理を定義
	// どのルートにもマッチしないリクエストがあった場合に、ステータスコード200（OK）とともにindex.htmlを返す

	router.Run(":8080")
}
